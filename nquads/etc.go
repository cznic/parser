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
	ast  []*Statement
	prev scanner.Token
}

func (l *lexer) Lex(lval *yySymType) int {
again:
	tok, val := l.Scan()
	if tok == scanner.EOL && l.prev == scanner.EOL {
		goto again
	}
	l.prev = tok

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

func (p Pos) String() string { return fmt.Sprintf("%d:%d", p.Line, p.Col) }

type Tag int

const (
	_ Tag = iota
	IRIRef
	BlankNodeLabel
	Literal
	LangTag
)

func (t Tag) String() string {
	switch t {
	case 0:
		return ""
	case IRIRef:
		return "IRIRef"
	case BlankNodeLabel:
		return "BlankNodeLabel"
	case Literal:
		return "Literal"
	case LangTag:
		return "LangTag"
	default:
		return fmt.Sprintf("%T(%d)", t, int(t))
	}
}

type Statement struct {
	Pos
	*Subject
	*Predicate
	*Object
	*GraphLabel
}

func (s *Statement) String() string {
	switch {
	case s.GraphLabel == nil:
		return fmt.Sprintf("stmt@%v{%v, %v, %v}", s.Pos, s.Subject, s.Predicate, s.Object)
	default:
		return fmt.Sprintf("stmt@%v{%v, %v, %v, %v}", s.Pos, s.Subject, s.Predicate, s.Object, s.GraphLabel)
	}
}

type Subject struct {
	Pos
	Tag
	Value string
}

func (s *Subject) String() string { return fmt.Sprintf("subj@%v{%v=%q}", s.Pos, s.Tag, s.Value) }

type Predicate struct {
	Pos
	Value string
}

func (p *Predicate) String() string { return fmt.Sprintf("pred@%v{%q}", p.Pos, p.Value) }

type Object struct {
	Pos
	Tag
	Value  string
	Tag2   Tag
	Value2 string
}

func (o *Object) String() string {
	switch {
	case o.Tag2 == 0:
		return fmt.Sprintf("obj@%v{%v=%q}", o.Pos, o.Tag, o.Value)
	default:
		return fmt.Sprintf("obj@%v{%v=%q, %v=%q}", o.Pos, o.Tag, o.Value, o.Tag2, o.Value2)
	}
}

type GraphLabel struct {
	Pos
	Tag
	Value string
}

func (g *GraphLabel) String() string { return fmt.Sprintf("graph@%v{%v=%q}", g.Pos, g.Tag, g.Value) }

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
