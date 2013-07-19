%{

/*

Copyright © 2001-2004 The IEEE and The Open Group, All Rights reserved.

Original source text: http://pubs.opengroup.org/onlinepubs/009695399/utilities/yacc.html

Modifications: Copyright 2013 The Go Authors. All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

Grammar for the input to yacc.

*/

// Package parser (WIP:TODO) implements a parser for yacc source files.
package parser

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/cznic/scanner/yacc"
	"github.com/cznic/strutil"
)

%}

%union {
	act    *Act
	col    int
	def    *Def
	defs   []*Def
	item   interface{}
	line   int
	list   []interface{}
	nlist  []*Nmno
	nmno   *Nmno
	number int
	prec   *Prec
	rule   *Rule
	rules  []*Rule
	rword  Rword
	s      string
}

%token _ILLEGAL

/* Basic entries. The following are recognized by the lexical analyzer. */

%token	_IDENTIFIER      /* Includes identifiers and literals */
%token	_C_IDENTIFIER    /* identifier (but not literal)
                             followed by a :. */
%token	_NUMBER          /* [0-9][0-9]* */

/* Reserved words : %type=>_TYPE %left=>_LEFT, and so on */

%token	_LEFT _RIGHT _NONASSOC _TOKEN _PREC _TYPE _START _UNION

%token	_MARK            /* The %% mark. */
%token	_LCURL           /* The %{ mark. */
%token	_RCURL           /* The %} mark. */

%type	<item>	_IDENTIFIER
%type	<number> _NUMBER
%type	<s>	_C_IDENTIFIER

%type	<act>	act
%type	<def>	def
%type	<defs>	defs
%type	<list>	rbody
%type	<nlist>	nlist
%type	<nmno>	nmno
%type	<prec>	prec
%type	<rule>	rule
%type	<rules>	rules
%type	<rword>	rword
%type	<s>	tag tail

/* 8-bit character literals stand for themselves; */
/* tokens have to be defined for multi-byte characters. */

%start	spec

%%

spec:
	defs _MARK rules tail
	{
		lx(yylex).spec = &Spec{Defs: $1, Rules: $3, Tail: $4}
	}

tail:
	/* Empty; the second _MARK is optional. */
	{
		$$ = ""
	}
|	_MARK
	{
        	/* In this action, set up the rest of the file. */
		lx := lx(yylex)
		$$ = string(lx.src[lx.Pos()+1:])
		lx.closed = true
	}

defs:
	/* Empty. */
	{
		$$ = []*Def(nil)
	}
	| defs def
	{
		$$ = append($1, $2)
	}

def:
   	_START _IDENTIFIER
	{
		s, ok := $2.(string)
		if !ok {
			lx := lx(yylex)
			lx.Error(fmt.Sprintf("%s:%d:%d expected name", lx.fname, $<line>2, $<col>2))
		}
		$$ = &Def{Rword: Start, Tag: s}
	}
|	_UNION
	{
        	/* Copy union definition to output. */
		lx := lx(yylex)
		lx.Mode(false)
		off0 := lx.Pos()+5
		n := 0
	union_loop:
		for {
			tok, _ := lx.Scan()
			switch tok {
			case scanner.LBRACE:
				n++
			case scanner.RBRACE:
				n--
				if n == 0 {
					lx.Mode(true)
					break union_loop
				}
			}
		}
		s := string(lx.src[off0:lx.Pos()])
		$$ = &Def{Rword: Union, Tag: s}
	}
|	_LCURL
	{
		/* Copy Go code to output file. */
		lx := lx(yylex)
		off0, lpos := lx.Pos(), lx.Pos()
		lx.Mode(false)
		var last scanner.Token
	lcurl_loop:
		for {
			tok, _ := lx.ScanRaw()
			if tok == scanner.RBRACE && last == scanner.REM && lx.Pos() == lpos+1 {
				lx.Mode(true)
				s := string(lx.src[off0+1:lpos-1])
				//dbg("----\n%q\n----\n", s)
				$$ = &Def{Rword: Copy, Tag: s}
				break lcurl_loop
			}

			last, lpos = tok, lx.Pos()
		}
	}
|	rword tag nlist
	{
		$$ = &Def{Rword: $1, Tag: $2, Nlist: $3}
	}

rword:
	_TOKEN
	{
		$$ = Token
	}
|	_LEFT
	{
		$$ = Left
	}
|	_RIGHT
	{
		$$ = Right
	}
|	_NONASSOC
	{
		$$ = Nonassoc
	}
|	_TYPE
	{
		$$ = Type
	}

tag:
	/* Empty: union tag ID optional. */
	{
		$$ = ""
	}
|	'<' _IDENTIFIER '>'
	{
		lx := lx(yylex)
		s, ok := $2.(string)
		if ! ok {
			lx.Error(fmt.Sprintf("%s:%d:%d expected name", lx.fname, $<line>2, $<col>2))
		}
		$$ = s
	}

nlist:
	nmno
	{
		$$ = []*Nmno{$1}
	}
|	nlist nmno
	{
		$$ = append($1, $2)
	}

nmno:
	_IDENTIFIER
	{
		/*TODO Note: literal invalid with % type. */
		$$ = &Nmno{$1, -1}
	}
|	_IDENTIFIER _NUMBER
	{
		/*TODO Note: invalid with % type. */
		$$ = &Nmno{$1, $2}
	}

/* Rule section */

rules:
	_C_IDENTIFIER rbody prec
	{
		lx(yylex).rname = $1
		$$ = []*Rule{&Rule{Name: $1, Body: $2, Prec: $3}}
	}
|	rules rule
	{
		$$ = append($1, $2)
	}

rule:
	_C_IDENTIFIER rbody prec
	{
		lx(yylex).rname = $1
		$$ = &Rule{Name: $1, Body: $2, Prec: $3}
	}
|	'|' rbody prec
	{
		$$ = &Rule{Name: lx(yylex).rname, Body: $2, Prec: $3}
	}

rbody:
	/* empty */
	{
		$$ = []interface{}(nil)
	}
|	rbody _IDENTIFIER
	{
		$$ = append($1, $2)
	}
|	rbody act
	{
		$$ = append($1, $2)
	}

act:
	'{'
	{
		/* Copy action, translate $$, and so on. */
		lx := lx(yylex)
		lx.Mode(false)
		off0 := lx.Pos()
		n := 0
	act_loop:
		for {
			tok, _ := lx.Scan()
			switch tok {
			case scanner.LBRACE:
				n++
			case scanner.RBRACE:
				if n == 0 {
					lx.Mode(true)
					break act_loop
				}

				n--
			}
		}
		$$ = &Act{Src: string(lx.src[off0:lx.Pos()-1])}
	}

prec:
	/* Empty */
	{
		$$ = nil
	}
|	_PREC _IDENTIFIER
	{
		$$ = &Prec{Identifier: $2}
	}
|	_PREC _IDENTIFIER act
	{
		$$ = &Prec{Identifier: $2, Act: $3}
	}
|	prec ';'
	{
		$$ = &Prec{}
	}

%%

// Spec is the AST root entity.
type Spec struct {
	Defs  []*Def  // Definitions
	Rules []*Rule // Rules
	Tail  string  // Optional rest of the file
}

// String implements fmt.Stringer.
func (s *Spec) String() string {
	return str(s)
}

// Def is the definition section definition entity
type Def struct {
	Rword Rword
	Tag   string
	Nlist []*Nmno
}

// String implements fmt.Stringer.
func (s *Def) String() string {
	return str(s)
}

// Rule is the rules section rule.
type Rule struct{
	Name string
	Body []interface{}
	Prec *Prec
}

// String implements fmt.Stringer.
func (s *Rule) String() string {
	return str(s)
}

// Name-or-number is a definition section name list item. It's either a
// production name (type string), or a [rune] literal (type rune). Optional
// number associated with the name is in NUmber, if non-negative.
type Nmno struct {
	Identifier interface{}
	Number int
}

// String implements fmt.Stringer.
func (s *Nmno) String() string {
	return str(s)
}

// Prec defines precedenc of a rule.
type Prec struct {
	Identifier interface{}
	Act *Act
}

// String implements fmt.Stringer.
func (s *Prec) String() string {
	return str(s)
}

// Act captures the action optionally associated with a rule.
type Act struct{
	Src string
	//TODO process $$
}

// String implements fmt.Stringer.
func (s *Act) String() string {
	return str(s)
}


// Rword is a definition tag (Def.Rword).
type Rword int

const (
	_ Rword = iota

	// Values of Def.Rword
	Copy     // %{ ... %}
	Left     // %left
	Nonassoc // %nonassoc
	Right    // %right
	Start    // %start
	Token    // %token
	Type     // %type
	Union    // %union
)

var rwords = map[Rword]string{
	Copy:     "Copy",
	Left:     "Left",
	Nonassoc: "Nonassoc",
	Right:    "Right",
	Start:    "Start",
	Token:    "Token",
	Type:     "Type",
	Union:    "Union",
}

// String implements fmt.Stringer.
func (r Rword) String() string {
	if s := rwords[r]; s != "" {
		return s
	}

	return fmt.Sprintf("%T(%d)", r, r)
}

type lexer struct {
	*scanner.Scanner
	fname  string
	rname  string // last rule name for '|' rules
	spec   *Spec
	src    []byte
	closed bool
}

var xlat = map[scanner.Token]int{
	scanner.LCURL:        _LCURL,
	scanner.LEFT:         _LEFT,
	scanner.MARK:         _MARK,
	scanner.NONASSOC:     _NONASSOC,
	scanner.PREC:         _PREC,
	scanner.RCURL:        _RCURL,
	scanner.RIGHT:        _RIGHT,
	scanner.START:        _START,
	scanner.TOKEN:        _TOKEN,
	scanner.TYPE:         _TYPE,
	scanner.UNION:        _UNION,

	scanner.EOF:          0,
	scanner.OR:           '|',
}

var todo = strings.ToUpper("todo")

func (l *lexer) Lex(lval *yySymType) (y int) {
	if l.closed {
		return 0
	}

	for {
		tok, val := l.Scan()
		lval.line, lval.col = l.Line, l.Col
		//dbg("%s %T(%#v) %s:%d:%d", tok, val, val, l.fname, l.Line, l.Col)
		switch tok {
		case scanner.COMMENT:
			continue
		case scanner.C_IDENTIFIER:
			if s, ok := val.(string); ok {
				lval.s = s
			}
			return _C_IDENTIFIER
		case scanner.IDENTIFIER:
			if s, ok := val.(string); ok {
				lval.item = s
			}
			return _IDENTIFIER
		case scanner.INT:
			if n, ok := val.(uint64); ok {
				lval.number = int(n)
			}
			return _NUMBER
		case scanner.CHAR:
			if n, ok := val.(int32); ok {
				lval.item = int(n)
			}
			return _IDENTIFIER
		case scanner.ILLEGAL:
			if s, ok := val.(string); ok && s != "" {
				return int([]rune(s)[0])
			}
			return _ILLEGAL
		default:
			return xlat[tok]
		}
	}
}

type errList []error

func (e errList) Error() string {
	a := []string{}
	for _, v := range e {
		a = append(a, v.Error())
	}
	return strings.Join(a, "\n")
}

func lx(yylex yyLexer) *lexer {
	return yylex.(*lexer)
}


// Parse parses src a single yacc source file fname and returns the
// corresponding Spec. If the source couldn't be read, the returned Spec is nil
// and the error indicates all of the specific failures.
func Parse(fname string, src []byte) (s *Spec, err error) {
	l := lexer{
		Scanner: scanner.New(src),
		fname:   fname,
		src:     src,
	}
	l.Fname = fname
	defer func() {
		if e := recover(); e != nil {
			l.Error(fmt.Sprintf("%v", e))
			err = errList(l.Errors)
			return
		}
	}()
	if yyParse(&l) != 0 {
		return nil, errList(l.Errors)
	}

	return l.spec, nil
}

func str(v interface{}) string {
	var buf bytes.Buffer
	f := strutil.IndentFormatter(&buf, ". ")
	g := func(interface{}){}
	g = func(v interface{}){
		switch x := v.(type) {
		case nil:
			f.Format("<nil>")
		case int:
			f.Format("'%c'\n", x)
		case string:
			f.Format("%q\n", x)
		case *Act:
			f.Format("%T{", x)
			f.Format("Src: %q", x.Src)
			f.Format("}\n")
		case *Def:
			f.Format("%T{%i\n", x)
			f.Format("Rword: %s, ", x.Rword)
			f.Format("Tag: %q, ", x.Tag)
			f.Format("Nlist: %T{%i\n", x.Nlist)
			for _, v := range x.Nlist {
				g(v)
			}
			f.Format("%u}\n")
			f.Format("%u}\n")
		case *Nmno:
			var s string
			switch v := x.Identifier.(type) {
			case string:
				s = fmt.Sprintf("%q", v)
			case int:
				s = fmt.Sprintf("'%c'", v)
			}
			f.Format("%T{Identifier: %s, Number: %d}\n", x, s, x.Number)
		case *Prec:
			var s string
			switch v := x.Identifier.(type) {
			case string:
				s = fmt.Sprintf("%q", v)
			case int:
				s = fmt.Sprintf("'%c'", v)
			}
			f.Format("%T{Identifier: %s, Act: ", x, s)
			//TODO bypassing compiler bug? Should work w/o test for nil
			if x.Act != nil {
				g(x.Act)
			} else {
				f.Format("<nil>")
			}
			f.Format("}\n")
		case *Rule:
			f.Format("%T{%i\n", x)
			f.Format("Name: %q, ", x.Name)
			f.Format("Body: %T{%i\n", x.Body)
			for _, v := range x.Body {
				g(v)
			}
			f.Format("%u}\n")
			if x.Prec != nil {
				f.Format("Prec: ")
				g(x.Prec)
			}
			f.Format("%u}\n")
		case *Spec:
			f.Format("%T{%i\n", x)
			f.Format("Defs: %T{%i\n", x.Defs)
			for _, v := range x.Defs {
				g(v)
			}
			f.Format("%u}\n")
			f.Format("Rules: %T{%i\n", x.Rules)
			for _, v := range x.Rules {
				g(v)
			}
			f.Format("%u}\n")
			f.Format("Tail: %q\n", x.Tail)
			f.Format("%u}\n")
		default:
			f.Format("%s(str): %T(%#v)\n", todo, x, x)
		}
	}
	g(v)
	return buf.String()
}
