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

import __yyfmt__ "fmt"

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/cznic/scanner/yacc"
	"github.com/cznic/strutil"
)

type yySymType struct {
	yys    int
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

const _ILLEGAL = 57346
const _IDENTIFIER = 57347
const _C_IDENTIFIER = 57348
const _NUMBER = 57349
const _LEFT = 57350
const _RIGHT = 57351
const _NONASSOC = 57352
const _TOKEN = 57353
const _PREC = 57354
const _TYPE = 57355
const _START = 57356
const _UNION = 57357
const _MARK = 57358
const _LCURL = 57359
const _RCURL = 57360

var yyToknames = []string{
	"_ILLEGAL",
	"_IDENTIFIER",
	"_C_IDENTIFIER",
	"_NUMBER",
	"_LEFT",
	"_RIGHT",
	"_NONASSOC",
	"_TOKEN",
	"_PREC",
	"_TYPE",
	"_START",
	"_UNION",
	"_MARK",
	"_LCURL",
	"_RCURL",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

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

type Def struct { //TODO docs
	Rword Rword
	Tag   string
	Nlist []*Nmno
}

// String implements fmt.Stringer.
func (s *Def) String() string {
	return str(s)
}

type Rule struct { //TODO docs
	Name string
	Body []interface{}
	Prec *Prec
}

// String implements fmt.Stringer.
func (s *Rule) String() string {
	return str(s)
}

type Nmno struct { //TODO docs
	Identifier interface{}
	Number     int
}

// String implements fmt.Stringer.
func (s *Nmno) String() string {
	return str(s)
}

type Prec struct { //TODO docs
	Identifier interface{}
	Act        *Act
}

// String implements fmt.Stringer.
func (s *Prec) String() string {
	return str(s)
}

type Act struct { //TODO docs
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
	scanner.LCURL:    _LCURL,
	scanner.LEFT:     _LEFT,
	scanner.MARK:     _MARK,
	scanner.NONASSOC: _NONASSOC,
	scanner.PREC:     _PREC,
	scanner.RCURL:    _RCURL,
	scanner.RIGHT:    _RIGHT,
	scanner.START:    _START,
	scanner.TOKEN:    _TOKEN,
	scanner.TYPE:     _TYPE,
	scanner.UNION:    _UNION,

	scanner.EOF: 0,
	scanner.OR:  '|',
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

//TODO docs
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
	g := func(interface{}) {}
	g = func(v interface{}) {
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

var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 33
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 44

var yyAct = []int{

	33, 31, 26, 10, 11, 12, 9, 32, 13, 5,
	6, 3, 7, 24, 34, 41, 35, 38, 18, 22,
	42, 37, 15, 27, 35, 28, 16, 1, 36, 21,
	19, 39, 40, 17, 23, 8, 29, 30, 14, 20,
	25, 2, 4, 43,
}
var yyPact = []int{

	-1000, -1000, -5, 16, -1000, 21, -1000, -1000, -1, -1000,
	-1000, -1000, -1000, -1000, 13, -1000, -1000, 18, 20, -1000,
	-1000, -1000, -1000, -1000, 2, 18, -1000, 14, -3, 2,
	2, -8, -1000, -1000, 15, -1000, -1000, -1000, -1000, -8,
	-8, -1000, -6, -1000,
}
var yyPgo = []int{

	0, 0, 42, 41, 13, 40, 2, 1, 39, 38,
	35, 33, 30, 27,
}
var yyR1 = []int{

	0, 13, 12, 12, 3, 3, 2, 2, 2, 2,
	10, 10, 10, 10, 10, 11, 11, 5, 5, 6,
	6, 9, 9, 8, 8, 4, 4, 4, 1, 7,
	7, 7, 7,
}
var yyR2 = []int{

	0, 4, 0, 1, 0, 2, 2, 1, 1, 3,
	1, 1, 1, 1, 1, 0, 3, 1, 2, 1,
	2, 3, 2, 3, 3, 0, 2, 2, 1, 0,
	2, 3, 2,
}
var yyChk = []int{

	-1000, -13, -3, 16, -2, 14, 15, 17, -10, 11,
	8, 9, 10, 13, -9, 6, 5, -11, 19, -12,
	-8, 16, 6, 21, -4, -5, -6, 5, 5, -4,
	-4, -7, 5, -1, 12, 22, -6, 7, 20, -7,
	-7, 23, 5, -1,
}
var yyDef = []int{

	4, -2, 0, 0, 5, 0, 7, 8, 15, 10,
	11, 12, 13, 14, 2, 25, 6, 0, 0, 1,
	22, 3, 25, 25, 29, 9, 17, 19, 0, 29,
	29, 21, 26, 27, 0, 28, 18, 20, 16, 23,
	24, 32, 30, 31,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 23,
	19, 3, 20, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 22, 21,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18,
}
var yyTok3 = []int{
	0,
}

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(yyToknames) {
		if yyToknames[c-4] != "" {
			return yyToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %U %s\n", uint(char), yyTokname(c))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yychar < 0 {
			yychar = yylex1(yylex, &yylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yychar {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf("saw %s\n", yyTokname(yychar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}
			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		{
			lx(yylex).spec = &Spec{Defs: yyS[yypt-3].defs, Rules: yyS[yypt-1].rules, Tail: yyS[yypt-0].s}
		}
	case 2:
		{
			yyVAL.s = ""
		}
	case 3:
		{
			/* In this action, set up the rest of the file. */
			lx := lx(yylex)
			yyVAL.s = string(lx.src[lx.Pos()+1:])
			lx.closed = true
		}
	case 4:
		{
			yyVAL.defs = []*Def(nil)
		}
	case 5:
		{
			yyVAL.defs = append(yyS[yypt-1].defs, yyS[yypt-0].def)
		}
	case 6:
		{
			s, ok := yyS[yypt-0].item.(string)
			if !ok {
				lx := lx(yylex)
				lx.Error(fmt.Sprintf("%s:%d:%d expected name", lx.fname, yyS[yypt-0].line, yyS[yypt-0].col))
			}
			yyVAL.def = &Def{Rword: Start, Tag: s}
		}
	case 7:
		{
			/* Copy union definition to output. */
			lx := lx(yylex)
			lx.Mode(false)
			off0 := lx.Pos() + 5
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
			yyVAL.def = &Def{Rword: Union, Tag: s}
		}
	case 8:
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
					s := string(lx.src[off0+1 : lpos-1])
					//dbg("----\n%q\n----\n", s)
					yyVAL.def = &Def{Rword: Copy, Tag: s}
					break lcurl_loop
				}

				last, lpos = tok, lx.Pos()
			}
		}
	case 9:
		{
			yyVAL.def = &Def{Rword: yyS[yypt-2].rword, Tag: yyS[yypt-1].s, Nlist: yyS[yypt-0].nlist}
		}
	case 10:
		{
			yyVAL.rword = Token
		}
	case 11:
		{
			yyVAL.rword = Left
		}
	case 12:
		{
			yyVAL.rword = Right
		}
	case 13:
		{
			yyVAL.rword = Nonassoc
		}
	case 14:
		{
			yyVAL.rword = Type
		}
	case 15:
		{
			yyVAL.s = ""
		}
	case 16:
		{
			lx := lx(yylex)
			s, ok := yyS[yypt-1].item.(string)
			if !ok {
				lx.Error(fmt.Sprintf("%s:%d:%d expected name", lx.fname, yyS[yypt-1].line, yyS[yypt-1].col))
			}
			yyVAL.s = s
		}
	case 17:
		{
			yyVAL.nlist = []*Nmno{yyS[yypt-0].nmno}
		}
	case 18:
		{
			yyVAL.nlist = append(yyS[yypt-1].nlist, yyS[yypt-0].nmno)
		}
	case 19:
		{
			/*TODO Note: literal invalid with % type. */
			yyVAL.nmno = &Nmno{yyS[yypt-0].item, -1}
		}
	case 20:
		{
			/*TODO Note: invalid with % type. */
			yyVAL.nmno = &Nmno{yyS[yypt-1].item, yyS[yypt-0].number}
		}
	case 21:
		{
			lx(yylex).rname = yyS[yypt-2].s
			yyVAL.rules = []*Rule{&Rule{Name: yyS[yypt-2].s, Body: yyS[yypt-1].list, Prec: yyS[yypt-0].prec}}
		}
	case 22:
		{
			yyVAL.rules = append(yyS[yypt-1].rules, yyS[yypt-0].rule)
		}
	case 23:
		{
			lx(yylex).rname = yyS[yypt-2].s
			yyVAL.rule = &Rule{Name: yyS[yypt-2].s, Body: yyS[yypt-1].list, Prec: yyS[yypt-0].prec}
		}
	case 24:
		{
			yyVAL.rule = &Rule{Name: lx(yylex).rname, Body: yyS[yypt-1].list, Prec: yyS[yypt-0].prec}
		}
	case 25:
		{
			yyVAL.list = []interface{}(nil)
		}
	case 26:
		{
			yyVAL.list = append(yyS[yypt-1].list, yyS[yypt-0].item)
		}
	case 27:
		{
			yyVAL.list = append(yyS[yypt-1].list, yyS[yypt-0].act)
		}
	case 28:
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
			yyVAL.act = &Act{Src: string(lx.src[off0 : lx.Pos()-1])}
		}
	case 29:
		{
			yyVAL.prec = nil
		}
	case 30:
		{
			yyVAL.prec = &Prec{Identifier: yyS[yypt-0].item}
		}
	case 31:
		{
			yyVAL.prec = &Prec{Identifier: yyS[yypt-1].item, Act: yyS[yypt-0].act}
		}
	case 32:
		{
			yyVAL.prec = &Prec{}
		}
	}
	goto yystack /* stack new state and value */
}
