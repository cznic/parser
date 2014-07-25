// Copyright (c) 2014 Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package scanner implements a scanner for N-Quads[0] source text.
package parser

import (
	"fmt"
	"strings"

	"github.com/cznic/scanner/nquads"
)

type lexer struct {
	*scanner.Scanner
	ast []*Statement
}

func (l *lexer) Lex(lval *yySymType) int {
	tok, val := l.Scan()
	lval.pos, lval.val = Pos{l.Line, l.Col}, val
	//dbg("%s:%d:%d %v %q", l.Fname, l.Line, l.Col, tok, val)
	switch tok {
	case scanner.EOF:
		return 0
	case scanner.DOT:
		return dot
	case scanner.DACCENT:
		return daccent
	case scanner.LABEL:
		return label
	case scanner.EOL:
		return eol
	case scanner.IRIREF:
		return iriref
	case scanner.LANGTAG:
		return langtag
	case scanner.STRING:
		return str
	default:
		panic("internal error")
	}
}

type Pos struct {
	Line int
	Col  int
}

type Tag int

const (
	_ Tag = iota
	IRIRef
	BlankNodeLabel
	Literal
	LangTag
)

type Statement struct {
	Pos
	*Subject
	*Predicate
	*Object
	*GraphLabel
}

type Subject struct {
	Pos
	Tag
	Value string
}

type Predicate struct {
	Pos
	Value string
}

type Object struct {
	Pos
	Tag
	Value  string
	Tag2   Tag
	Value2 string
}

type GraphLabel struct {
	Pos
	Tag
	Value string
}

// Parse parses src as a single N-Quads source file fname and returns the
// corresponding AST. If the source couldn't be parsed, the returned AST is
// nil and the error indicates all of the specific failures.
func Parse(fname string, src []byte) (ast []*Statement, err error) {
	l := lexer{
		Scanner: scanner.New(fname, src),
	}
	defer func() {
		if e := recover(); e != nil {
			l.Error(fmt.Sprintf("%v", e))
			ast, err = nil, errList(l.Errors)
		}
	}()
	if yyParse(&l) != 0 {
		return nil, errList(l.Errors)
	}

	return l.ast, nil
}

type errList []error

func (e errList) Error() string {
	a := []string{}
	for _, v := range e {
		a = append(a, v.Error())
	}
	return strings.Join(a, "\n")
}
