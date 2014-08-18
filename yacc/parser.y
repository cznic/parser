%{

/*

Copyright © 2001-2004 The IEEE and The Open Group, All Rights reserved.

Original source text: http://pubs.opengroup.org/onlinepubs/009695399/utilities/yacc.html

Modifications: Copyright 2014 The parser Authors. All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

Grammar for the input to yacc.

*/

// Package parser implements a parser for yacc source files.
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
	act    []*Act
	def    *Def
	defs   []*Def
	item   interface{}
	list   []interface{}
	nlist  []*Nmno
	nmno   *Nmno
	num    int
	number int
	pos    Pos
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

%token	_LEFT _RIGHT _NONASSOC _TOKEN _PREC _TYPE _START _UNION _ERR_VERBOSE

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
		lx(yylex).ast = &AST{Defs: $1, Rules: $3, Tail: $4}
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
|	defs def
	{
		$$ = append($1, $2)
	}

def:
   	_START _IDENTIFIER
	{
		s, ok := $2.(string)
		if !ok {
			lx := lx(yylex)
			lx.Error(fmt.Sprintf("%s:%d:%d expected name", lx.FName, $<pos>2.Line, $<pos>2.Col))
		}
		$$ = &Def{Pos: $<pos>1, Rword: Start, Tag: s}
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
			tok, _, _ := lx.Scan()
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
		$$ = &Def{Pos: $<pos>1, Rword: Union, Tag: s}
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
			tok, _, _ := lx.ScanRaw()
			if tok == scanner.RBRACE && last == scanner.REM && lx.Pos() == lpos+1 {
				lx.Mode(true)
				s := string(lx.src[off0+1:lpos-1])
				//dbg("----\n%q\n----\n", s)
				$$ = &Def{Pos: $<pos>1, Rword: Copy, Tag: s}
				break lcurl_loop
			}

			last, lpos = tok, lx.Pos()
		}
	}
|	_ERR_VERBOSE
	{
		$$ = &Def{Pos: $<pos>1, Rword: ErrVerbose}
	}
|	rword tag nlist
	{
		if $1 == Type {
			for _, v := range $3 {
				switch v.Identifier.(type) {
				case int:
					yylex.Error("literal invalid with %type.") // % is ok
					goto ret1
				}

				if v.Number > 0 {
					yylex.Error("number invalid with %type.") // % is ok
					goto ret1
				}
			}
		}

		$$ = &Def{Pos: $<pos>1, Rword: $1, Tag: $2, Nlist: $3}
	}

rword:
	_TOKEN
	{
		$<pos>$ = $<pos>1
		$$ = Token
	}
|	_LEFT
	{
		$<pos>$ = $<pos>1
		$$ = Left
	}
|	_RIGHT
	{
		$<pos>$ = $<pos>1
		$$ = Right
	}
|	_NONASSOC
	{
		$<pos>$ = $<pos>1
		$$ = Nonassoc
	}
|	_TYPE
	{
		$<pos>$ = $<pos>1
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
			lx.Error(fmt.Sprintf("%s:%d:%d expected name", lx.FName, $<pos>2.Line, $<pos>2.Col))
		}
		$<pos>$ = $<pos>2
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
		$$ = &Nmno{$<pos>1, $1, -1}
	}
|	_IDENTIFIER _NUMBER
	{
		$$ = &Nmno{$<pos>1, $1, $2}
	}

/* Rule section */

rules:
	_C_IDENTIFIER rbody prec
	{
		lx(yylex).rname = $1
		$$ = []*Rule{&Rule{Pos: $<pos>1, Name: $1, Body: $2, Prec: $3}}
	}
|	rules rule
	{
		$$ = append($1, $2)
	}

rule:
	_C_IDENTIFIER rbody prec
	{
		lx(yylex).rname = $1
		$$ = &Rule{Pos: $<pos>1, Name: $1, Body: $2, Prec: $3}
	}
|	'|' rbody prec
	{
		$$ = &Rule{Pos: $<pos>1, Name: lx(yylex).rname, Body: $2, Prec: $3}
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
		a := []*Act{}
		start := lx.Pos()
		n := 0
		line, col := -1, -1
	act_loop:
		for {
			tok, tag, num := lx.Scan()
			if line < 0 {
				line, col = lx.Line, lx.Col
			}
			tokStart := lx.Pos()-1
			switch tok {
			case scanner.DLR_DLR, scanner.DLR_NUM, scanner.DLR_TAG_DLR, scanner.DLR_TAG_NUM:
				s, ok := tag.(string)
				if !ok {
					s = ""
				}

				src := ""
				if start > 0 {
					src = string(lx.src[start:tokStart])
				}
				
				a = append(a, &Act{Pos: Pos{lx.Line, lx.Col}, Src: src, Tok: tok, Tag: s, Num: num})
				start = -1
			case scanner.LBRACE:
				n++
			case scanner.RBRACE:
				if n == 0 {
					if start < 0 {
						start = tokStart
						line, col = lx.Line, lx.Col
					}
					src := lx.src[start:tokStart]
					if len(src) != 0 {
						a = append(a, &Act{Pos: Pos{line, col}, Src: string(src)})
					}
					lx.Mode(true)
					break act_loop
				}

				n--
			case scanner.EOF:
				lx.Error("unexpected EOF")
				goto ret1
			default:
				if start < 0 {
					start = tokStart
					line, col = lx.Line, lx.Col
				}
			}
		}
		$$ = a
	}

prec:
	/* Empty */
	{
		$$ = nil
	}
|	_PREC _IDENTIFIER
	{
		$$ = &Prec{Pos: $<pos>1, Identifier: $2}
	}
|	_PREC _IDENTIFIER act
	{
		$$ = &Prec{Pos: $<pos>1, Identifier: $2, Act: $3}
	}
|	prec ';'
	{
		$$ = &Prec{}
	}

%%

// AST holds the parsed .y source.
type AST struct {
	Defs  []*Def  // Definitions
	Rules []*Rule // Rules
	Tail  string  // Optional rest of the file
}

// String implements fmt.Stringer.
func (s *AST) String() string {
	return str(s)
}

// Pos descibes a source line and column.
type Pos struct {
	Line int
	Col  int
}

// Def is the definition section definition entity
type Def struct {
	Pos
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
	Pos
	Name string
	Body []interface{}
	Prec *Prec
}

// String implements fmt.Stringer.
func (s *Rule) String() string {
	return str(s)
}

// Name-or-number is a definition section name list item. It's either a
// production name (type string), or a rune literal. Optional number associated
// with the name is in number, if non-negative.
type Nmno struct {
	Pos
	Identifier interface{}
	Number int
}

// String implements fmt.Stringer.
func (s *Nmno) String() string {
	return str(s)
}

// Prec defines the optional precedence of a rule.
type Prec struct {
	Pos
	Identifier interface{}
	Act []*Act
}

// String implements fmt.Stringer.
func (s *Prec) String() string {
	return str(s)
}

// Act captures the action optionally associated with a rule.
type Act struct{
	Pos
	Src string
	Tok scanner.Token       // github.com/cznic/scanner/yacc.DLR_* or zero
	Tag string              // DLR_TAG_*
	Num int                 // DLR_NUM, DLR_TAG_NUM
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
	Copy       // %{ ... %}
	ErrVerbose // %error-verbose
	Left       // %left
	Nonassoc   // %nonassoc
	Right      // %right
	Start      // %start
	Token      // %token
	Type       // %type
	Union      // %union
)

var rwords = map[Rword]string{
	Copy:       "Copy",
	ErrVerbose: "ErrorVerbose",
	Left:       "Left",
	Nonassoc:   "Nonassoc",
	Right:      "Right",
	Start:      "Start",
	Token:      "Token",
	Type:       "Type",
	Union:      "Union",
}

// String implements fmt.Stringer.
func (r Rword) String() string {
	if s := rwords[r]; s != "" {
		return s
	}

	return fmt.Sprintf("Rword(%d)", r)
}

type lexer struct {
	*scanner.Scanner
	ast    *AST
	closed bool
	rname  string // last rule name for '|' rules
	src    []byte
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
	scanner.ERR_VERBOSE:  _ERR_VERBOSE,

	scanner.EOF:          0,
	scanner.OR:           '|',
}

var todo = strings.ToUpper("todo")

func (l *lexer) Lex(lval *yySymType) (y int) {
	if l.closed {
		return 0
	}

	for {
		tok, val, num := l.Scan()
		lval.pos, lval.num = Pos{l.Line, l.Col}, num
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


// Parse parses src as a single yacc source file fname and returns the
// corresponding AST. If the source couldn't be read, the returned AST is nil
// and the error indicates all of the specific failures.
func Parse(fname string, src []byte) (s *AST, err error) {
	l := lexer{
		Scanner: scanner.New(fname, src),
		src:     src,
	}
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

	return l.ast, nil
}

func str(v interface{}) string {
	var buf bytes.Buffer
	f := strutil.IndentFormatter(&buf, "· ")
	g := func(interface{}){}
	g = func(v interface{}){
		switch x := v.(type) {
		case nil:
			f.Format("<nil>")
		case int:
			f.Format("'%c'\n", x)
		case string:
			f.Format("%q\n", x)
		case []*Act:
			f.Format("%T{%i\n", x)
			for _, v := range x {
				g(v)
			}
			f.Format("%u}\n")
		case *Act:
			f.Format("%T@%d:%d{%i\n", x, x.Line, x.Col)
			f.Format("Src: %q\n", x.Src)
			if x.Tok != 0 {
				f.Format("Tok: %s, Tag: %q, Num: %d\n", x.Tok, x.Tag, x.Num)
			}
			f.Format("%u}\n")
		case *Def:
			f.Format("%T@%d:%d{%i\n", x, x.Line, x.Col)
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
			f.Format("%T@%d:%d{Identifier: %s, Number: %d}\n", x, x.Line, x.Col, s, x.Number)
		case *Prec:
			var s string
			switch v := x.Identifier.(type) {
			case string:
				s = fmt.Sprintf("%q", v)
			case int:
				s = fmt.Sprintf("'%c'", v)
			}
			f.Format("%T@%d:%d{%i\n", x, x.Line, x.Col)
			f.Format("Identifier: %s\n", s)
			g(x.Act)
			f.Format("%u}\n")
		case *Rule:
			f.Format("%T@%d:%d{%i\n", x, x.Line, x.Col)
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
		case *AST:
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
