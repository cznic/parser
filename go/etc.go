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
func ParseFile(fset *token.FileSet, filename string, src interface{}) (ast interface{}, err error) {
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
	p := &parser{file: file}
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
	file   *token.File
	sc     scanner.Scanner
	errors scanner.ErrorList
	ast    []Node
	pos    token.Pos
}

func (p *parser) Error(e string) {
	p.errors.Add(p.file.Position(p.pos), e)
}

var xlat = map[token.Token]int{
	// Special tokens
	token.ILLEGAL: -1,
	token.EOF:     0,
	token.COMMENT: -1,

	// Identifiers and basic type literals
	// (these tokens stand for classes of literals)
	token.IDENT:  NAME,    // main
	token.INT:    LITERAL, // 12345
	token.FLOAT:  LITERAL, // 123.45
	token.IMAG:   LITERAL, // 123.45i
	token.CHAR:   LITERAL, // 'a'
	token.STRING: LITERAL, // "abc"

	// Operators and delimiters
	token.ADD: '+', // +
	token.SUB: '-', // -
	token.MUL: '*', // *
	token.QUO: '/', // /
	token.REM: '%', // %

	token.AND:     '&',    // &
	token.OR:      '|',    // |
	token.XOR:     '^',    // ^
	token.SHL:     LSH,    // <<
	token.SHR:     RSH,    // >>
	token.AND_NOT: ANDNOT, // &^

	token.ADD_ASSIGN: ASOP, // +=
	token.SUB_ASSIGN: ASOP, // -=
	token.MUL_ASSIGN: ASOP, // *=
	token.QUO_ASSIGN: ASOP, // /=
	token.REM_ASSIGN: ASOP, // %=

	token.AND_ASSIGN:     ASOP, // &=
	token.OR_ASSIGN:      ASOP, // |=
	token.XOR_ASSIGN:     ASOP, // ^=
	token.SHL_ASSIGN:     ASOP, // <<=
	token.SHR_ASSIGN:     ASOP, // >>=
	token.AND_NOT_ASSIGN: ASOP, // &^=

	token.LAND:  ANDAND, // &&
	token.LOR:   ORO,    // ||
	token.ARROW: COMM,   // <-
	token.INC:   INC,    // ++
	token.DEC:   DEC,    // --

	token.EQL:    EQ,  // ==
	token.LSS:    '<', // <
	token.GTR:    '>', // >
	token.ASSIGN: '=', // =
	token.NOT:    '!', // !

	token.NEQ:      NE,    // !=
	token.LEQ:      LE,    // <=
	token.GEQ:      GE,    // >=
	token.DEFINE:   COLAS, // :=
	token.ELLIPSIS: DDD,   // ...

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
	token.BREAK:    BREAK,
	token.CASE:     CASE,
	token.CHAN:     CHAN,
	token.CONST:    COMST,
	token.CONTINUE: CONTINUE,

	token.DEFAULT:     DEFAULT,
	token.DEFER:       DEFER,
	token.ELSE:        ELSE,
	token.FALLTHROUGH: FALL,
	token.FOR:         FOR,

	token.FUNC:   FUNC,
	token.GO:     GO,
	token.GOTO:   GOTO,
	token.IF:     IF,
	token.IMPORT: IMPORT,

	token.INTERFACE: INTERFACE,
	token.MAP:       MAP,
	token.PACKAGE:   PACKAGE,
	token.RANGE:     RANGE,
	token.RETURN:    RETURN,

	token.SELECT: SELECT,
	token.STRUCT: STRUCT,
	token.SWITCH: SWITCH,
	token.TYPE:   TYPE,
	token.VAR:    VAR,
}

func (p *parser) Lex(lval *yySymType) int {
	var tok token.Token
	p.pos, tok, lval.lit = p.sc.Scan()
	return int(tok) //TODO
}
