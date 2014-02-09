// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package parser (WIP) implements a parser for Go source files. Input may be
// provided in a variety of forms; the output is an abstract syntax tree (AST)
// representing the Go source.
package parser

import (
	"bytes"
	"go/scanner"
	"go/token"
	"io"
	"io/ioutil"
)

// ParseFile parses the source code of a single Go source file and returns the
// corresponding AST. The source code may be provided via the filename of the
// source file, or via the src parameter.
//
// If src != nil, ParseFile parses the source from src and the filename is only
// used when recording position information. The type of the argument for the
// src parameter must be string, []byte, or io.Reader. If src == nil, ParseFile
// parses the file specified by filename.
//
// If the source couldn't be read, the returned AST is nil and the error
// indicates the specific failure.
func ParseFile(fset *token.FileSet, filename string, src interface{}, packageScope *Scope /*TODO Opts*/) (ast []Node, err error) {
	var bsrc []byte
	switch x := src.(type) {
	case nil:
		if bsrc, err = ioutil.ReadFile(filename); err != nil {
			return
		}
	case string:
		bsrc = []byte(x)
	case []byte:
		bsrc = x
	case io.Reader:
		var buf bytes.Buffer
		if _, err := io.Copy(&buf, x); err != nil {
			return nil, err
		}

		bsrc = buf.Bytes()
	}

	file := fset.AddFile(filename, -1, len(bsrc))
	p := &parser{
		file:         file,
		fileScope:    NewScope(packageScope),
		fset:         fset,
		packageScope: packageScope,
	}
	p.currentScope = p.fileScope
	p.sc.Init(
		file,
		bsrc,
		func(pos token.Position, msg string) {
			p.errors.Add(pos, msg)
		},
		0,
	)
	if yyParse(p) != 0 || len(p.errors) != 0 {
		return nil, p.errors
	}

	return p.ast, nil
}

type parser struct {
	ast          []Node
	constExpr    []Node
	constIota    int
	constType    Node
	currentScope *Scope
	errors       scanner.ErrorList
	file         *token.File
	fileScope    *Scope
	fset         *token.FileSet
	packageScope *Scope
	pos          token.Pos
	sc           scanner.Scanner
	stack        []int
}

func (p *parser) Error(e string) {
	p.errPos(p.pos, e)
}

func (p *parser) errPos(pos token.Pos, e string) {
	p.errors.Add(p.file.Position(pos), e)
}

var xlat = map[token.Token]int{
	// Special tokens
	token.ILLEGAL: _IGNORE,
	token.EOF:     0,
	token.COMMENT: -1,

	// Identifiers and basic type literals
	// (these tokens stand for classes of literals)
	token.IDENT:  _NAME,    // main
	token.INT:    _LITERAL, // 12345
	token.FLOAT:  _LITERAL, // 123.45
	token.IMAG:   _LITERAL, // 123.45i
	token.CHAR:   _LITERAL, // 'a'
	token.STRING: _LITERAL, // "abc"

	// Operators and delimiters
	token.ADD: '+', // +
	token.SUB: '-', // -
	token.MUL: '*', // *
	token.QUO: '/', // /
	token.REM: '%', // %

	token.AND:     '&',     // &
	token.OR:      '|',     // |
	token.XOR:     '^',     // ^
	token.SHL:     _LSH,    // <<
	token.SHR:     _RSH,    // >>
	token.AND_NOT: _ANDNOT, // &^

	token.ADD_ASSIGN: _ASOP, // +=
	token.SUB_ASSIGN: _ASOP, // -=
	token.MUL_ASSIGN: _ASOP, // *=
	token.QUO_ASSIGN: _ASOP, // /=
	token.REM_ASSIGN: _ASOP, // %=

	token.AND_ASSIGN:     _ASOP, // &=
	token.OR_ASSIGN:      _ASOP, // |=
	token.XOR_ASSIGN:     _ASOP, // ^=
	token.SHL_ASSIGN:     _ASOP, // <<=
	token.SHR_ASSIGN:     _ASOP, // >>=
	token.AND_NOT_ASSIGN: _ASOP, // &^=

	token.LAND:  _ANDAND, // &&
	token.LOR:   _OROR,   // ||
	token.ARROW: _COMM,   // <-
	token.INC:   _INC,    // ++
	token.DEC:   _DEC,    // --

	token.EQL:    _EQ, // ==
	token.LSS:    '<', // <
	token.GTR:    '>', // >
	token.ASSIGN: '=', // =
	token.NOT:    '!', // !

	token.NEQ:      _NE,    // !=
	token.LEQ:      _LE,    // <=
	token.GEQ:      _GE,    // >=
	token.DEFINE:   _COLAS, // :=
	token.ELLIPSIS: _DDD,   // ...

	token.LPAREN: '(', // (
	token.LBRACK: '[', // [
	token.LBRACE: '{', // {
	token.COMMA:  ',', // ,
	token.PERIOD: '.', // .

	token.RPAREN:    ')', // )
	token.RBRACK:    ']', // ]
	token.RBRACE:    '}', // }
	token.SEMICOLON: ';', // ;
	token.COLON:     ':', // :

	// Keywords
	token.BREAK:    _BREAK,
	token.CASE:     _CASE,
	token.CHAN:     _CHAN,
	token.CONST:    _CONST,
	token.CONTINUE: _CONTINUE,

	token.DEFAULT:     _DEFAULT,
	token.DEFER:       _DEFER,
	token.ELSE:        _ELSE,
	token.FALLTHROUGH: _FALL,
	token.FOR:         _FOR,

	token.FUNC:   _FUNC,
	token.GO:     _GO,
	token.GOTO:   _GOTO,
	token.IF:     _IF,
	token.IMPORT: _IMPORT,

	token.INTERFACE: _INTERFACE,
	token.MAP:       _MAP,
	token.PACKAGE:   _PACKAGE,
	token.RANGE:     _RANGE,
	token.RETURN:    _RETURN,

	token.SELECT: _SELECT,
	token.STRUCT: _STRUCT,
	token.SWITCH: _SWITCH,
	token.TYPE:   _TYPE,
	token.VAR:    _VAR,
}

type tkn struct {
	lit  string
	pos  pos
	tok  token.Token
	tpos token.Pos
}

func (p *parser) Lex(lval *yySymType) (r int) {
	for {
		ps, tok, lit := p.sc.Scan()
		if r = xlat[tok]; r < 0 {
			continue
		}

		switch r {
		case '(':
			n := len(p.stack)
			if n == 0 {
				break
			}

			p.stack[n-1]++
		case ')':
			n := len(p.stack)
			if n == 0 {
				break
			}

			p.stack[n-1]--
		case '{':
			if n := len(p.stack); n != 0 && p.stack[n-1] == 0 {
				r = _BODY
				p.stack = p.stack[:n-1]
			}
		}

		p.pos, lval.token = ps, tkn{lit, pos(ps), tok, ps}
		return
	}
}

func idList(l []Node) (r []*Ident) {
	r = make([]*Ident, len(l))
	for i, v := range l {
		r[i] = v.(*Ident)
	}
	return
}

const dlrPkgName = "$pkgName"

// Special name(s):
// - "$package" in package scope holds parsed package name *Ident.
type Scope struct {
	Parent *Scope
	Names  map[string]Node
}

func NewScope(parent *Scope) *Scope { return &Scope{Parent: parent, Names: map[string]Node{}} }

func (s *Scope) New() *Scope {
	return NewScope(s)
}
