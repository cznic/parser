/*

Copyright © 2001-2004 The IEEE and The Open Group, All Rights reserved.

Original source text: http://pubs.opengroup.org/onlinepubs/009695399/utilities/yacc.html

Modifications: Copyright 2014 The parser Authors. All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

Grammar for the input to yacc.

*/

// Package parser implements a parser for yacc source files.
//
// Changelog
//
// 2014-11-11: Include the opening and closing braces of semantic actions in
// Act.Src.
package parser

import __yyfmt__ "fmt"

import (
	"bytes"
	"fmt"
	"go/token"
	"strconv"
	"strings"

	"github.com/cznic/scanner/yacc"
	"github.com/cznic/strutil"
)

type yySymType struct {
	yys    int
	act    []*Act
	def    *Def
	defs   []*Def
	item   interface{}
	list   []interface{}
	nlist  []*Nmno
	nmno   *Nmno
	number int
	pos    token.Pos
	prec   *Prec
	rule   *Rule
	rules  []*Rule
	rword  Rword
	s      string
}

const illegal = 57346
const tkIdent = 57347
const tkCIdent = 57348
const tkNumber = 57349
const tkLeft = 57350
const tkRight = 57351
const tkNonAssoc = 57352
const tkToken = 57353
const tkPrec = 57354
const tkType = 57355
const tkStart = 57356
const tkUnion = 57357
const tkErrorVerbose = 57358
const tkMark = 57359
const tkLCurl = 57360
const tkRCurl = 57361

var yyToknames = []string{
	"illegal",
	"tkIdent",
	"tkCIdent",
	"tkNumber",
	"tkLeft",
	"tkRight",
	"tkNonAssoc",
	"tkToken",
	"tkPrec",
	"tkType",
	"tkStart",
	"tkUnion",
	"tkErrorVerbose",
	"tkMark",
	"tkLCurl",
	"tkRCurl",
}
var yyStatenames = []string{}

const yyEOFCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

// AST holds the parsed .y source.
type AST struct {
	Defs  []*Def  // Definitions
	Rules []*Rule // Rules
	Tail  string  // Optional rest of the file
	fset  *token.FileSet
}

// String implements fmt.Stringer.
func (s *AST) String() string {
	return str(s.fset, s)
}

// Def is the definition section definition entity
type Def struct {
	token.Pos
	Rword Rword
	Tag   string
	Nlist []*Nmno
}

// Rule is the rules section rule.
type Rule struct {
	token.Pos
	Name string
	Body []interface{}
	Prec *Prec
}

// Nmno (Name-or-number) is a definition section name list item. It's either a
// production name (type string), or a rune literal. Optional number associated
// with the name is in number, if non-negative.
type Nmno struct {
	token.Pos
	Identifier interface{}
	Number     int
}

// Prec defines the optional precedence of a rule.
type Prec struct {
	token.Pos
	Identifier interface{}
	Act        []*Act
}

// Act captures the action optionally associated with a rule.  The action parts
// are split at the yacc tokens $$, $num, $<tag>num, if present.
type Act struct {
	token.Pos
	Src string
	Tok scanner.Token // github.com/cznic/scanner/yacc.DLR_* or zero
	Tag string        // DLR_TAG_*
	Num int           // DLR_NUM, DLR_TAG_NUM
}

// Rword is a definition tag (Def.Rword).
type Rword int

// Values of Def.Rword
const (
	_ Rword = iota

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
	fset   *token.FileSet
	rname  string // last rule name for '|' rules
	src    []byte
}

var xlat = map[scanner.Token]int{
	scanner.LCURL:       tkLCurl,
	scanner.LEFT:        tkLeft,
	scanner.MARK:        tkMark,
	scanner.NONASSOC:    tkNonAssoc,
	scanner.PREC:        tkPrec,
	scanner.RCURL:       tkRCurl,
	scanner.RIGHT:       tkRight,
	scanner.START:       tkStart,
	scanner.TOKEN:       tkToken,
	scanner.TYPE:        tkType,
	scanner.UNION:       tkUnion,
	scanner.ERR_VERBOSE: tkErrorVerbose,

	scanner.COMMA: ',',
	scanner.EOF:   0,
	scanner.OR:    '|',
}

var todo = strings.ToUpper("todo")

func (l *lexer) Lex(lval *yySymType) (y int) {
	if l.closed {
		return 0
	}

	for {
		tok, val, _ := l.Scan()
		lval.pos = token.Pos(l.Pos())
		switch tok {
		case scanner.COMMENT:
			continue
		case scanner.C_IDENTIFIER:
			if s, ok := val.(string); ok {
				lval.s = s
			}
			return tkCIdent
		case scanner.IDENTIFIER:
			if s, ok := val.(string); ok {
				lval.item = s
			}
			return tkIdent
		case scanner.INT:
			if n, ok := val.(uint64); ok {
				lval.number = int(n)
			}
			return tkNumber
		case scanner.CHAR:
			if n, ok := val.(int32); ok {
				lval.item = int(n)
			}
			return tkIdent
		case scanner.ILLEGAL:
			if s, ok := val.(string); ok && s != "" {
				return int([]rune(s)[0])
			}
			return illegal
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
func Parse(fset *token.FileSet, fname string, src []byte) (s *AST, err error) {
	l := lexer{
		fset:    fset,
		Scanner: scanner.New(fset, fname, src),
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

	l.ast.fset = fset
	return l.ast, nil
}

func str(fset *token.FileSet, v interface{}) string {
	var buf bytes.Buffer
	f := strutil.IndentFormatter(&buf, "· ")
	g := func(interface{}) {}
	g = func(v interface{}) {
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
			f.Format("%T@%v{%i\n", x, fset.Position(x.Pos))
			f.Format("Src: %q\n", x.Src)
			if x.Tok != 0 {
				f.Format("Tok: %s, Tag: %q, Num: %d\n", x.Tok, x.Tag, x.Num)
			}
			f.Format("%u}\n")
		case *Def:
			f.Format("%T@%v{%i\n", x, fset.Position(x.Pos))
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
			f.Format("%T@%v{Identifier: %s, Number: %d}\n", x, fset.Position(x.Pos), s, x.Number)
		case *Prec:
			var s string
			switch v := x.Identifier.(type) {
			case string:
				s = fmt.Sprintf("%q", v)
			case int:
				s = fmt.Sprintf("'%c'", v)
			}
			f.Format("%T@%v{%i\n", x, fset.Position(x.Pos))
			f.Format("Identifier: %s\n", s)
			g(x.Act)
			f.Format("%u}\n")
		case *Rule:
			f.Format("%T@%v{%i\n", x, fset.Position(x.Pos))
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

var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 35
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 48

var yyAct = []int{

	34, 27, 32, 11, 12, 13, 10, 33, 14, 5,
	6, 8, 3, 7, 35, 43, 36, 25, 23, 28,
	40, 19, 39, 16, 28, 44, 36, 29, 37, 22,
	17, 1, 20, 41, 42, 24, 38, 18, 9, 15,
	45, 30, 31, 21, 26, 46, 2, 4,
}
var yyPact = []int{

	-1000, -1000, -5, 17, -1000, 25, -1000, -1000, -1000, 1,
	-1000, -1000, -1000, -1000, -1000, 12, -1000, -1000, 19, 22,
	-1000, -1000, -1000, -1000, -1000, 2, 14, -1000, 15, -1,
	2, 2, -10, -1000, -1000, 20, -1000, -1000, 19, -1000,
	-1000, -10, -10, -1000, -8, -1000, -1000,
}
var yyPgo = []int{

	0, 0, 47, 46, 17, 44, 1, 2, 43, 39,
	38, 37, 32, 31,
}
var yyR1 = []int{

	0, 13, 12, 12, 3, 3, 2, 2, 2, 2,
	2, 10, 10, 10, 10, 10, 11, 11, 5, 5,
	5, 6, 6, 9, 9, 8, 8, 4, 4, 4,
	1, 7, 7, 7, 7,
}
var yyR2 = []int{

	0, 4, 0, 1, 0, 2, 2, 1, 1, 1,
	3, 1, 1, 1, 1, 1, 0, 3, 1, 2,
	3, 1, 2, 3, 2, 3, 3, 0, 2, 2,
	1, 0, 2, 3, 2,
}
var yyChk = []int{

	-1000, -13, -3, 17, -2, 14, 15, 18, 16, -10,
	11, 8, 9, 10, 13, -9, 6, 5, -11, 20,
	-12, -8, 17, 6, 23, -4, -5, -6, 5, 5,
	-4, -4, -7, 5, -1, 12, 24, -6, 22, 7,
	21, -7, -7, 25, 5, -6, -1,
}
var yyDef = []int{

	4, -2, 0, 0, 5, 0, 7, 8, 9, 16,
	11, 12, 13, 14, 15, 2, 27, 6, 0, 0,
	1, 24, 3, 27, 27, 31, 10, 18, 21, 0,
	31, 31, 23, 28, 29, 0, 30, 19, 0, 22,
	17, 25, 26, 34, 32, 20, 33,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 22, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 25,
	20, 3, 21, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 24, 23,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19,
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
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(c), uint(char))
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
				__yyfmt__.Printf(" saw %s\n", yyTokname(yychar))
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
			if yychar == yyEOFCode {
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
			lx(yylex).ast = &AST{Defs: yyS[yypt-3].defs, Rules: yyS[yypt-1].rules, Tail: yyS[yypt-0].s}
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
				lx.Error(fmt.Sprintf("%v: expected name", yyS[yypt-0].pos))
			}
			yyVAL.def = &Def{Pos: yyS[yypt-1].pos, Rword: Start, Tag: s}
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
			yyVAL.def = &Def{Pos: yyS[yypt-0].pos, Rword: Union, Tag: s}
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
				tok, _, _ := lx.ScanRaw()
				if tok == scanner.RBRACE && last == scanner.REM && lx.Pos() == lpos+1 {
					lx.Mode(true)
					s := string(lx.src[off0+1 : lpos-1])
					//dbg("----\n%q\n----\n", s)
					yyVAL.def = &Def{Pos: yyS[yypt-0].pos, Rword: Copy, Tag: s}
					break lcurl_loop
				}

				last, lpos = tok, lx.Pos()
			}
		}
	case 9:
		{
			yyVAL.def = &Def{Pos: yyS[yypt-0].pos, Rword: ErrVerbose}
		}
	case 10:
		{
			if yyS[yypt-2].rword == Type {
				for _, v := range yyS[yypt-0].nlist {
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

			yyVAL.def = &Def{Pos: yyS[yypt-2].pos, Rword: yyS[yypt-2].rword, Tag: yyS[yypt-1].s, Nlist: yyS[yypt-0].nlist}
		}
	case 11:
		{
			yyVAL.pos = yyS[yypt-0].pos
			yyVAL.rword = Token
		}
	case 12:
		{
			yyVAL.pos = yyS[yypt-0].pos
			yyVAL.rword = Left
		}
	case 13:
		{
			yyVAL.pos = yyS[yypt-0].pos
			yyVAL.rword = Right
		}
	case 14:
		{
			yyVAL.pos = yyS[yypt-0].pos
			yyVAL.rword = Nonassoc
		}
	case 15:
		{
			yyVAL.pos = yyS[yypt-0].pos
			yyVAL.rword = Type
		}
	case 16:
		{
			yyVAL.s = ""
		}
	case 17:
		{
			lx := lx(yylex)
			s, ok := yyS[yypt-1].item.(string)
			if !ok {
				lx.Error(fmt.Sprintf("%v: expected name", yyS[yypt-1].pos))
			}
			yyVAL.pos = yyS[yypt-1].pos
			yyVAL.s = s
		}
	case 18:
		{
			yyVAL.nlist = []*Nmno{yyS[yypt-0].nmno}
		}
	case 19:
		{
			yyVAL.nlist = append(yyS[yypt-1].nlist, yyS[yypt-0].nmno)
		}
	case 20:
		{
			yyVAL.nlist = append(yyS[yypt-2].nlist, yyS[yypt-0].nmno)
		}
	case 21:
		{
			yyVAL.nmno = &Nmno{yyS[yypt-0].pos, yyS[yypt-0].item, -1}
		}
	case 22:
		{
			yyVAL.nmno = &Nmno{yyS[yypt-1].pos, yyS[yypt-1].item, yyS[yypt-0].number}
		}
	case 23:
		{
			lx(yylex).rname = yyS[yypt-2].s
			yyVAL.rules = []*Rule{&Rule{Pos: yyS[yypt-2].pos, Name: yyS[yypt-2].s, Body: yyS[yypt-1].list, Prec: yyS[yypt-0].prec}}
		}
	case 24:
		{
			yyVAL.rules = append(yyS[yypt-1].rules, yyS[yypt-0].rule)
		}
	case 25:
		{
			lx(yylex).rname = yyS[yypt-2].s
			yyVAL.rule = &Rule{Pos: yyS[yypt-2].pos, Name: yyS[yypt-2].s, Body: yyS[yypt-1].list, Prec: yyS[yypt-0].prec}
		}
	case 26:
		{
			yyVAL.rule = &Rule{Pos: yyS[yypt-2].pos, Name: lx(yylex).rname, Body: yyS[yypt-1].list, Prec: yyS[yypt-0].prec}
		}
	case 27:
		{
			yyVAL.list = []interface{}(nil)
		}
	case 28:
		{
			yyVAL.list = append(yyS[yypt-1].list, yyS[yypt-0].item)
		}
	case 29:
		{
			yyVAL.list = append(yyS[yypt-1].list, yyS[yypt-0].act)
		}
	case 30:
		{
			/* Copy action, translate $$, and so on. */
			lx := lx(yylex)
			lx.Mode(false)
			a := []*Act{}
			start := lx.Pos() - 1 // First '{' inclusive.
			var d int
			for lvl := 1; lvl > 0; {
				tok, tag, num := lx.Scan()
				s, _ := tag.(string)
				switch tok {
				case scanner.DLR_DLR:
					d = 1
				case scanner.DLR_NUM:
					d = len(strconv.Itoa(num))
				case scanner.DLR_TAG_DLR:
					d = len(s) + 3
				case scanner.DLR_TAG_NUM:
					d = len(s) + 2 + len(strconv.Itoa(num))
				}
				switch tok {
				case scanner.DLR_DLR, scanner.DLR_NUM, scanner.DLR_TAG_DLR, scanner.DLR_TAG_NUM:
					a = append(a, &Act{Pos: token.Pos(start + 1), Src: string(lx.src[start : lx.Pos()-1]), Tok: tok, Tag: s, Num: num})
					start = lx.Pos() + d
				case scanner.LBRACE:
					lvl++
				case scanner.RBRACE:
					lvl--
					if lvl == 0 {
						a = append(a, &Act{Pos: token.Pos(start + 1), Src: string(lx.src[start:lx.Pos()])})
						lx.Mode(true)
					}
				case scanner.EOF:
					lx.Error("unexpected EOF")
					goto ret1
				}
			}
			yyVAL.act = a
		}
	case 31:
		{
			yyVAL.prec = nil
		}
	case 32:
		{
			yyVAL.prec = &Prec{Pos: yyS[yypt-1].pos, Identifier: yyS[yypt-0].item}
		}
	case 33:
		{
			yyVAL.prec = &Prec{Pos: yyS[yypt-2].pos, Identifier: yyS[yypt-1].item, Act: yyS[yypt-0].act}
		}
	case 34:
		{
			yyVAL.prec = yyS[yypt-1].prec // Temporary workaround for issue #2
		}
	}
	goto yystack /* stack new state and value */
}
