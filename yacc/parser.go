//line parser.y:2

/*

Copyright © 2001-2004 The IEEE and The Open Group, All Rights reserved.

Original source text: http://pubs.opengroup.org/onlinepubs/009695399/utilities/yacc.html

Modifications: Copyright 2013 The Go Authors. All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

Grammar for the input to yacc.

*/

// (WIP:TODO)
package parser

import __yyfmt__ "fmt"

//line parser.y:18
import (
	"github.com/cznic/scanner/yacc"
)

//line parser.y:26
type yySymType struct {
	yys    int
	number int
}

const _IDENTIFIER = 57346
const _C_IDENTIFIER = 57347
const _NUMBER = 57348
const _LEFT = 57349
const _RIGHT = 57350
const _NONASSOC = 57351
const _TOKEN = 57352
const _PREC = 57353
const _TYPE = 57354
const _START = 57355
const _UNION = 57356
const _MARK = 57357
const _LCURL = 57358
const _RCURL = 57359

var yyToknames = []string{
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

//line parser.y:215

var xlat = map[scanner.Token]int{
	scanner.C_IDENTIFIER: _C_IDENTIFIER,
	scanner.IDENTIFIER:   _IDENTIFIER,
	scanner.INT:          _NUMBER,
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
}

type lexer struct {
	*scanner.Scanner
}

func (l *lexer) Lex(lval *yySymType) (y int) {
	tok, val := l.Scan()
	switch tok {
	case scanner.INT:
		if n, ok := val.(uint64); ok {
			lval.number = int(n)
		}
		return _NUMBER
	default:
		y = xlat[tok]
	}

	panic(".156")
}

//TODO docs
func Parse(fname string, src []byte) (y []interface{}, err error) {
	l := lexer{scanner.New(src)}
	l.Fname = fname
	if yyParse(&l) != 0 {
		return nil, l.Errors[0] //TODO
	}

	return
}

//line yacctab:1
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

	35, 33, 28, 25, 10, 11, 12, 9, 43, 13,
	5, 6, 3, 7, 34, 47, 37, 40, 19, 26,
	39, 36, 23, 15, 44, 29, 30, 31, 32, 16,
	38, 37, 22, 41, 42, 45, 21, 24, 27, 18,
	8, 17, 4, 20, 14, 46, 2, 1,
}
var yyPact = []int{

	-1000, -1000, -3, 18, -1000, 25, -1000, -1000, 0, -1000,
	-1000, -1000, -1000, -1000, 17, -1000, -1000, 2, 21, 22,
	-1000, -1000, -1000, -1000, -1000, 10, -1000, 21, -1000, 14,
	-2, 10, 10, -15, -1000, -1000, 20, -1000, -1000, -1000,
	-1000, -15, -15, -1000, -5, -7, -1000, -1000,
}
var yyPgo = []int{

	0, 47, 46, 44, 43, 42, 41, 40, 39, 38,
	2, 3, 1, 36, 0, 35,
}
var yyR1 = []int{

	0, 1, 4, 4, 2, 2, 5, 5, 6, 5,
	5, 7, 7, 7, 7, 7, 8, 8, 9, 9,
	10, 10, 3, 3, 13, 13, 11, 11, 11, 15,
	14, 12, 12, 12, 12,
}
var yyR2 = []int{

	0, 4, 0, 1, 0, 2, 2, 1, 0, 3,
	3, 1, 1, 1, 1, 1, 0, 3, 1, 2,
	1, 2, 3, 2, 3, 3, 0, 2, 2, 0,
	3, 0, 2, 3, 2,
}
var yyChk = []int{

	-1000, -1, -2, 15, -5, 13, 14, 16, -7, 10,
	7, 8, 9, 12, -3, 5, 4, -6, -8, 18,
	-4, -13, 15, 5, 20, -11, 17, -9, -10, 4,
	4, -11, -11, -12, 4, -14, 11, 21, -10, 6,
	19, -12, -12, 23, 4, -15, -14, 22,
}
var yyDef = []int{

	4, -2, 0, 0, 5, 0, 7, 8, 16, 11,
	12, 13, 14, 15, 2, 26, 6, 0, 0, 0,
	1, 23, 3, 26, 26, 31, 9, 10, 18, 20,
	0, 31, 31, 22, 27, 28, 0, 29, 19, 21,
	17, 24, 25, 34, 32, 0, 33, 30,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 23,
	18, 3, 19, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 21, 20, 22,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17,
}
var yyTok3 = []int{
	0,
}

//line yaccpar:1

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
		//line parser.y:56
		{
			panic(".y:56")
		}
	case 2:
		//line parser.y:62
		{
			panic(".y:62")
		}
	case 3:
		//line parser.y:66
		{
			/* In this action, set up the rest of the file. */
			panic(".y:67")
		}
	case 4:
		//line parser.y:72
		{
			panic(".y:73")
		}
	case 5:
		//line parser.y:76
		{
			panic(".y:77")
		}
	case 6:
		//line parser.y:82
		{
			panic(".y:83")
		}
	case 7:
		//line parser.y:86
		{
			/* Copy union definition to output. */
			panic(".y:88")
		}
	case 8:
		//line parser.y:90
		{
			/* Copy Go code to output file. */
			panic(".y:93")
		}
	case 9:
		//line parser.y:94
		{
			panic(".y:97")
		}
	case 10:
		//line parser.y:98
		{
			panic(".y:101")
		}
	case 11:
		//line parser.y:104
		{
			panic(".y:107")
		}
	case 12:
		//line parser.y:108
		{
			panic(".y:111")
		}
	case 13:
		//line parser.y:112
		{
			panic(".y:115")
		}
	case 14:
		//line parser.y:116
		{
			panic(".y:119")
		}
	case 15:
		//line parser.y:120
		{
			panic(".y:123")
		}
	case 16:
		//line parser.y:126
		{
			panic(".y:129")
		}
	case 17:
		//line parser.y:130
		{
			panic(".y:133")
		}
	case 18:
		//line parser.y:136
		{
			panic(".y:139")
		}
	case 19:
		//line parser.y:140
		{
			panic(".y:143")
		}
	case 20:
		//line parser.y:146
		{
			/* Note: literal invalid with % type. */
			panic(".y:150")
		}
	case 21:
		//line parser.y:150
		{
			/* Note: invalid with % type. */
			panic(".y:155")
		}
	case 22:
		//line parser.y:158
		{
			panic(".y:163")
		}
	case 23:
		//line parser.y:162
		{
			panic(".y:167")
		}
	case 24:
		//line parser.y:168
		{
			panic(".y:173")
		}
	case 25:
		//line parser.y:172
		{
			panic(".y:177")
		}
	case 26:
		//line parser.y:178
		{
			panic(".y:183")
		}
	case 27:
		//line parser.y:182
		{
			panic(".y:187")
		}
	case 28:
		//line parser.y:186
		{
			panic(".y:191")
		}
	case 29:
		//line parser.y:192
		{
			/* Copy action, translate $$, and so on. */
			panic(".y:198")
		}
	case 31:
		//line parser.y:199
		{
			panic(".y:205")
		}
	case 32:
		//line parser.y:203
		{
			panic(".y:209")
		}
	case 33:
		//line parser.y:207
		{
			panic(".y:213")
		}
	case 34:
		//line parser.y:211
		{
			panic(".y:217")
		}
	}
	goto yystack /* stack new state and value */
}
