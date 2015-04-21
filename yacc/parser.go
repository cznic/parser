// CAUTION: Generated file, DO NOT EDIT!

// Copyright © 2001-2004 The IEEE and The Open Group, All Rights reserved.
//
// Original source text:
// http://pubs.opengroup.org/onlinepubs/009695399/utilities/yacc.html
//
// Modifications: Copyright 2015 The parser Authors. All rights reserved.  Use
// of this source code is governed by a BSD-style license that can be found in
// the LICENSE file.
//
// Grammar for the input to yacc.
//
// CAUTION: Generated file (unless this is parser.y) - DO NOT EDIT!

package parser

import __yyfmt__ "fmt"

type yySymType struct {
	yys   int
	token *Token
	item  interface{}
}

type yyXError struct {
	state, xsym int
}

const (
	yyDefault     = 57364
	yyEofCode     = 57344
	COMMENT       = 57346
	C_IDENTIFIER  = 57347
	ERROR_VERBOSE = 57348
	IDENTIFIER    = 57349
	LCURL         = 57350
	LEFT          = 57351
	MARK          = 57352
	NONASSOC      = 57353
	NUMBER        = 57354
	PREC          = 57355
	PRECEDENCE    = 57356
	RCURL         = 57357
	RIGHT         = 57358
	START         = 57359
	TOKEN         = 57360
	TYPE          = 57361
	UNION         = 57362
	yyErrCode     = 57345

	yyMaxDepth = 200
	yyTabOfs   = -37
)

var (
	yyXLAT = map[int]int{
		57352: 0,  // MARK (30x)
		57349: 1,  // IDENTIFIER (28x)
		57344: 2,  // $end (20x)
		57347: 3,  // C_IDENTIFIER (18x)
		124:   4,  // '|' (17x)
		59:    5,  // ';' (15x)
		57348: 6,  // ERROR_VERBOSE (13x)
		57350: 7,  // LCURL (13x)
		57351: 8,  // LEFT (13x)
		57353: 9,  // NONASSOC (13x)
		57356: 10, // PRECEDENCE (13x)
		57358: 11, // RIGHT (13x)
		57359: 12, // START (13x)
		57360: 13, // TOKEN (13x)
		57361: 14, // TYPE (13x)
		57362: 15, // UNION (13x)
		123:   16, // '{' (10x)
		57355: 17, // PREC (9x)
		60:    18, // '<' (7x)
		44:    19, // ',' (6x)
		57365: 20, // Action (4x)
		57368: 21, // Name (3x)
		57370: 22, // Precedence (3x)
		57373: 23, // RuleItemList (3x)
		57357: 24, // RCURL (2x)
		57363: 25, // $@1 (1x)
		62:    26, // '>' (1x)
		125:   27, // '}' (1x)
		57366: 28, // Definition (1x)
		57367: 29, // DefinitionList (1x)
		57369: 30, // NameList (1x)
		57354: 31, // NUMBER (1x)
		57371: 32, // ReservedWord (1x)
		57372: 33, // Rule (1x)
		57374: 34, // RuleList (1x)
		57375: 35, // Specification (1x)
		57376: 36, // Tag (1x)
		57377: 37, // Tail (1x)
		57364: 38, // $default (0x)
		57346: 39, // COMMENT (0x)
		57345: 40, // error (0x)
	}

	yySymNames = []string{
		"MARK",
		"IDENTIFIER",
		"$end",
		"C_IDENTIFIER",
		"'|'",
		"';'",
		"ERROR_VERBOSE",
		"LCURL",
		"LEFT",
		"NONASSOC",
		"PRECEDENCE",
		"RIGHT",
		"START",
		"TOKEN",
		"TYPE",
		"UNION",
		"'{'",
		"PREC",
		"'<'",
		"','",
		"Action",
		"Name",
		"Precedence",
		"RuleItemList",
		"RCURL",
		"$@1",
		"'>'",
		"'}'",
		"Definition",
		"DefinitionList",
		"NameList",
		"NUMBER",
		"ReservedWord",
		"Rule",
		"RuleList",
		"Specification",
		"Tag",
		"Tail",
		"$default",
		"COMMENT",
		"error",
	}

	yyReductions = map[int]struct{ xsym, components int }{
		0:  {0, 1},
		1:  {20, 2},
		2:  {28, 2},
		3:  {28, 1},
		4:  {25, 0},
		5:  {28, 3},
		6:  {28, 3},
		7:  {28, 1},
		8:  {29, 0},
		9:  {29, 2},
		10: {21, 1},
		11: {21, 2},
		12: {30, 1},
		13: {30, 2},
		14: {30, 3},
		15: {22, 0},
		16: {22, 2},
		17: {22, 3},
		18: {22, 2},
		19: {32, 1},
		20: {32, 1},
		21: {32, 1},
		22: {32, 1},
		23: {32, 1},
		24: {32, 1},
		25: {33, 3},
		26: {33, 3},
		27: {23, 0},
		28: {23, 2},
		29: {23, 2},
		30: {34, 3},
		31: {34, 2},
		32: {35, 4},
		33: {36, 0},
		34: {36, 3},
		35: {37, 1},
		36: {37, 0},
	}

	yyXErrors = map[yyXError]string{
		yyXError{0, 2}:   "invalid empty input",
		yyXError{1, -1}:  "expected $end",
		yyXError{21, -1}: "expected $end",
		yyXError{22, -1}: "expected $end",
		yyXError{39, -1}: "expected '>'",
		yyXError{24, -1}: "expected '}'",
		yyXError{23, -1}: "expected Action or Precedence or one of [$end ';' '{' '|' C_IDENTIFIER IDENTIFIER MARK PREC]",
		yyXError{33, -1}: "expected Action or Precedence or one of [$end ';' '{' '|' C_IDENTIFIER IDENTIFIER MARK PREC]",
		yyXError{35, -1}: "expected Action or Precedence or one of [$end ';' '{' '|' C_IDENTIFIER IDENTIFIER MARK PREC]",
		yyXError{30, -1}: "expected Action or one of [$end ';' '{' '|' C_IDENTIFIER MARK]",
		yyXError{2, -1}:  "expected Definition or one of [ERROR_VERBOSE LCURL LEFT MARK NONASSOC PRECEDENCE RIGHT START TOKEN TYPE UNION]",
		yyXError{3, -1}:  "expected IDENTIFIER",
		yyXError{25, -1}: "expected IDENTIFIER",
		yyXError{38, -1}: "expected IDENTIFIER",
		yyXError{40, -1}: "expected IDENTIFIER",
		yyXError{46, -1}: "expected Name or IDENTIFIER",
		yyXError{41, -1}: "expected Name or one of [',' ERROR_VERBOSE IDENTIFIER LCURL LEFT MARK NONASSOC PRECEDENCE RIGHT START TOKEN TYPE UNION]",
		yyXError{37, -1}: "expected NameList or IDENTIFIER",
		yyXError{6, -1}:  "expected NameList or Tag or one of ['<' IDENTIFIER]",
		yyXError{16, -1}: "expected Precedence or RuleItemList or one of [$end ';' '{' '|' C_IDENTIFIER IDENTIFIER MARK PREC]",
		yyXError{18, -1}: "expected Precedence or RuleItemList or one of [$end ';' '{' '|' C_IDENTIFIER IDENTIFIER MARK PREC]",
		yyXError{19, -1}: "expected Precedence or RuleItemList or one of [$end ';' '{' '|' C_IDENTIFIER IDENTIFIER MARK PREC]",
		yyXError{5, -1}:  "expected RCURL",
		yyXError{48, -1}: "expected RCURL",
		yyXError{17, -1}: "expected Rule or Tail or one of [$end '|' C_IDENTIFIER MARK]",
		yyXError{15, -1}: "expected RuleList or C_IDENTIFIER",
		yyXError{0, -1}:  "expected Specification or one of [ERROR_VERBOSE LCURL LEFT MARK NONASSOC PRECEDENCE RIGHT START TOKEN TYPE UNION]",
		yyXError{27, -1}: "expected one of [$end ';' '{' '|' C_IDENTIFIER IDENTIFIER MARK PREC]",
		yyXError{28, -1}: "expected one of [$end ';' '{' '|' C_IDENTIFIER IDENTIFIER MARK PREC]",
		yyXError{32, -1}: "expected one of [$end ';' '{' '|' C_IDENTIFIER IDENTIFIER MARK PREC]",
		yyXError{26, -1}: "expected one of [$end ';' '|' C_IDENTIFIER MARK]",
		yyXError{29, -1}: "expected one of [$end ';' '|' C_IDENTIFIER MARK]",
		yyXError{31, -1}: "expected one of [$end ';' '|' C_IDENTIFIER MARK]",
		yyXError{34, -1}: "expected one of [$end ';' '|' C_IDENTIFIER MARK]",
		yyXError{36, -1}: "expected one of [$end ';' '|' C_IDENTIFIER MARK]",
		yyXError{20, -1}: "expected one of [$end '|' C_IDENTIFIER MARK]",
		yyXError{42, -1}: "expected one of [',' ERROR_VERBOSE IDENTIFIER LCURL LEFT MARK NONASSOC NUMBER PRECEDENCE RIGHT START TOKEN TYPE UNION]",
		yyXError{43, -1}: "expected one of [',' ERROR_VERBOSE IDENTIFIER LCURL LEFT MARK NONASSOC PRECEDENCE RIGHT START TOKEN TYPE UNION]",
		yyXError{44, -1}: "expected one of [',' ERROR_VERBOSE IDENTIFIER LCURL LEFT MARK NONASSOC PRECEDENCE RIGHT START TOKEN TYPE UNION]",
		yyXError{45, -1}: "expected one of [',' ERROR_VERBOSE IDENTIFIER LCURL LEFT MARK NONASSOC PRECEDENCE RIGHT START TOKEN TYPE UNION]",
		yyXError{47, -1}: "expected one of [',' ERROR_VERBOSE IDENTIFIER LCURL LEFT MARK NONASSOC PRECEDENCE RIGHT START TOKEN TYPE UNION]",
		yyXError{9, -1}:  "expected one of ['<' IDENTIFIER]",
		yyXError{10, -1}: "expected one of ['<' IDENTIFIER]",
		yyXError{11, -1}: "expected one of ['<' IDENTIFIER]",
		yyXError{12, -1}: "expected one of ['<' IDENTIFIER]",
		yyXError{13, -1}: "expected one of ['<' IDENTIFIER]",
		yyXError{14, -1}: "expected one of ['<' IDENTIFIER]",
		yyXError{4, -1}:  "expected one of [ERROR_VERBOSE LCURL LEFT MARK NONASSOC PRECEDENCE RIGHT START TOKEN TYPE UNION]",
		yyXError{7, -1}:  "expected one of [ERROR_VERBOSE LCURL LEFT MARK NONASSOC PRECEDENCE RIGHT START TOKEN TYPE UNION]",
		yyXError{8, -1}:  "expected one of [ERROR_VERBOSE LCURL LEFT MARK NONASSOC PRECEDENCE RIGHT START TOKEN TYPE UNION]",
		yyXError{49, -1}: "expected one of [ERROR_VERBOSE LCURL LEFT MARK NONASSOC PRECEDENCE RIGHT START TOKEN TYPE UNION]",
		yyXError{50, -1}: "expected one of [ERROR_VERBOSE LCURL LEFT MARK NONASSOC PRECEDENCE RIGHT START TOKEN TYPE UNION]",
	}

	yyParseTab = [51][]uint8{
		// 0
		{29, 6: 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29: 39, 35: 38},
		{2: 37},
		{52, 6: 44, 42, 47, 49, 51, 48, 40, 46, 50, 41, 28: 45, 32: 43},
		{1: 87},
		{34, 6: 34, 34, 34, 34, 34, 34, 34, 34, 34, 34},
		// 5
		{24: 33, 85},
		{1: 4, 18: 75, 36: 74},
		{30, 6: 30, 30, 30, 30, 30, 30, 30, 30, 30, 30},
		{28, 6: 28, 28, 28, 28, 28, 28, 28, 28, 28, 28},
		{1: 18, 18: 18},
		// 10
		{1: 17, 18: 17},
		{1: 16, 18: 16},
		{1: 15, 18: 15},
		{1: 14, 18: 14},
		{1: 13, 18: 13},
		// 15
		{3: 53, 34: 54},
		{10, 10, 10, 10, 10, 10, 16: 10, 10, 23: 72},
		{59, 2: 1, 55, 56, 33: 57, 37: 58},
		{10, 10, 10, 10, 10, 10, 16: 10, 10, 23: 70},
		{10, 10, 10, 10, 10, 10, 16: 10, 10, 23: 60},
		// 20
		{6, 2: 6, 6, 6},
		{2: 5},
		{2: 2},
		{22, 64, 22, 22, 22, 22, 16: 61, 62, 20: 65, 22: 63},
		{27: 69},
		// 25
		{1: 67},
		{11, 2: 11, 11, 11, 66},
		{9, 9, 9, 9, 9, 9, 16: 9, 9},
		{8, 8, 8, 8, 8, 8, 16: 8, 8},
		{19, 2: 19, 19, 19, 19},
		// 30
		{21, 2: 21, 21, 21, 21, 16: 61, 20: 68},
		{20, 2: 20, 20, 20, 20},
		{36, 36, 36, 36, 36, 36, 16: 36, 36},
		{22, 64, 22, 22, 22, 22, 16: 61, 62, 20: 65, 22: 71},
		{12, 2: 12, 12, 12, 66},
		// 35
		{22, 64, 22, 22, 22, 22, 16: 61, 62, 20: 65, 22: 73},
		{7, 2: 7, 7, 7, 66},
		{1: 79, 21: 80, 30: 78},
		{1: 76},
		{26: 77},
		// 40
		{1: 3},
		{31, 79, 6: 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 19: 83, 21: 82},
		{27, 27, 6: 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 19: 27, 31: 81},
		{25, 25, 6: 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 19: 25},
		{26, 26, 6: 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 19: 26},
		// 45
		{24, 24, 6: 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 19: 24},
		{1: 79, 21: 84},
		{23, 23, 6: 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 19: 23},
		{24: 86},
		{32, 6: 32, 32, 32, 32, 32, 32, 32, 32, 32, 32},
		// 50
		{35, 6: 35, 35, 35, 35, 35, 35, 35, 35, 35, 35},
	}
)

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyLexerEx interface {
	yyLexer
	Reduced(rule, state int, lval *yySymType) bool
}

func yySymName(c int) (s string) {
	x, ok := yyXLAT[c]
	if ok {
		return yySymNames[x]
	}

	return __yyfmt__.Sprintf("%d", c)
}

func yylex1(yylex yyLexer, lval *yySymType) (n int) {
	n = yylex.Lex(lval)
	if n <= 0 {
		n = yyEofCode
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("\nlex %s(%#x %d), lval.token: %v\n", yySymName(n), n, n, lval.token)
	}
	return n
}

func yyParse(yylex yyLexer) int {
	const yyError = 40

	yyEx, _ := yylex.(yyLexerEx)
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, 200)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yyerrok := func() {
		if yyDebug >= 2 {
			__yyfmt__.Printf("yyerrok()\n")
		}
		Errflag = 0
	}
	_ = yyerrok
	yystate := 0
	yychar := -1
	var yyxchar int
	var yyshift int
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
		var ok bool
		if yyxchar, ok = yyXLAT[yychar]; !ok {
			yyxchar = len(yySymNames) // > tab width
		}
	}
	if yyDebug >= 4 {
		var a []int
		for _, v := range yyS[:yyp+1] {
			a = append(a, v.yys)
		}
		__yyfmt__.Printf("state stack %v\n", a)
	}
	row := yyParseTab[yystate]
	yyn = 0
	if yyxchar < len(row) {
		if yyn = int(row[yyxchar]); yyn != 0 {
			yyn += yyTabOfs
		}
	}
	switch {
	case yyn > 0: // shift
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		yyshift = yyn
		if yyDebug >= 2 {
			__yyfmt__.Printf("shift, and goto state %d\n", yystate)
		}
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	case yyn < 0: // reduce
	case yystate == 1: // accept
		if yyDebug >= 2 {
			__yyfmt__.Println("accept")
		}
		goto ret0
	}

	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			if yyDebug >= 1 {
				__yyfmt__.Printf("no action for %s in state %d\n", yySymName(yychar), yystate)
			}
			msg, ok := yyXErrors[yyXError{yystate, yyxchar}]
			if !ok {
				msg, ok = yyXErrors[yyXError{yystate, -1}]
			}
			if !ok && yyshift != 0 {
				msg, ok = yyXErrors[yyXError{yyshift, yyxchar}]
			}
			if !ok {
				msg, ok = yyXErrors[yyXError{yyshift, -1}]
			}
			if !ok {
				msg = "syntax error"
			}
			if msg != "" {
				yylex.Error(msg)
			}
			Nerrs++
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				row := yyParseTab[yyS[yyp].yys]
				if yyError < len(row) {
					yyn = int(row[yyError]) + yyTabOfs
					if yyn > 0 { // hit
						if yyDebug >= 2 {
							__yyfmt__.Printf("error recovery found error shift in state %d\n", yyS[yyp].yys)
						}
						yystate = yyn /* simulate a shift of "error" */
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
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery failed\n")
			}
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yySymName(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}

			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	r := -yyn
	x0 := yyReductions[r]
	x, n := x0.xsym, x0.components
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= n
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	exState := yystate
	yystate = int(yyParseTab[yyS[yyp].yys][x]) + yyTabOfs
	/* reduction by production r */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce using rule %v (%s), and goto state %d\n", r, yySymNames[x], yystate)
	}

	switch r {
	case 1:
		{
			lx := yylex.(*lexer)
			lhs := &Action{
				Token:  yyS[yypt-1].token,
				Token2: yyS[yypt-0].token,
			}
			yyVAL.item = lhs
			//yy:field Pos token.Pos
			//yy:field Values []*ActionValue // For backward compatibility.
			lhs.Pos = lx.pos
			for i, v := range lx.values {
				a := lx.parseActionValue(lx.positions[i], v)
				if a != nil {
					lhs.Values = append(lhs.Values, a)
				}
			}
		}
	case 2:
		{
			lhs := &Definition{
				Token:  yyS[yypt-1].token,
				Token2: yyS[yypt-0].token,
			}
			yyVAL.item = lhs
			//yy:example "%start source\n\n%%"
			//yy:field Pos token.Pos
			//yy:field Value string
			//yy:field Nlist []*Name // For backward compatibility.
		}
	case 3:
		{
			lx := yylex.(*lexer)
			lhs := &Definition{
				Case:  1,
				Token: yyS[yypt-0].token,
			}
			yyVAL.item = lhs
			//yy:example "%union{\n        foo bar\n}\n\n%%"
			lhs.Pos = lx.pos
			lhs.Value = lx.value
		}
	case 4:
		{
			lx := yylex.(*lexer)
			lx.pos2 = lx.pos
			lx.value2 = lx.value
		}
	case 5:
		{
			lx := yylex.(*lexer)
			lhs := &Definition{
				Case:   2,
				Token:  yyS[yypt-2].token,
				Token2: yyS[yypt-0].token,
			}
			yyVAL.item = lhs
			lhs.Pos = lx.pos2
			lhs.Value = lx.value2
		}
	case 6:
		{
			lx := yylex.(*lexer)
			lhs := &Definition{
				Case:         3,
				ReservedWord: yyS[yypt-2].item.(*ReservedWord),
				Tag:          yyS[yypt-1].item.(*Tag),
				NameList:     yyS[yypt-0].item.(*NameList).reverse(),
			}
			yyVAL.item = lhs
			for n := lhs.NameList; n != nil; n = n.NameList {
				lhs.Nlist = append(lhs.Nlist, n.Name)
			}
			if lhs.ReservedWord.Token.Char.Rune == TYPE {
				for _, v := range lhs.Nlist {
					switch v.Identifier.(type) {
					case int:
						lx.err(v.Token.Pos(), "literal invalid with %type.")
					}

					if v.Number > 0 {
						lx.err(v.Token2.Pos(), "number invalid with %type.")
					}
				}
			}
		}
	case 7:
		{
			yyVAL.item = &Definition{
				Case:  4,
				Token: yyS[yypt-0].token,
			}
		}
	case 8:
		{
			yyVAL.item = (*DefinitionList)(nil)
		}
	case 9:
		{
			lx := yylex.(*lexer)
			lhs := &DefinitionList{
				Case:           1,
				DefinitionList: yyS[yypt-1].item.(*DefinitionList),
				Definition:     yyS[yypt-0].item.(*Definition),
			}
			yyVAL.item = lhs
			//yy:example "%left '+' '-'\n%left '*' '/'\n%%"
			lx.defs = append(lx.defs, lhs.Definition)
		}
	case 10:
		{
			lx := yylex.(*lexer)
			lhs := &Name{
				Token: yyS[yypt-0].token,
			}
			yyVAL.item = lhs
			//yy:field Identifier interface{} // For backward compatibility.
			//yy:field Number int             // For backward compatibility.
			lhs.Identifier = lx.ident(lhs.Token)
			lhs.Number = -1
		}
	case 11:
		{
			lx := yylex.(*lexer)
			lhs := &Name{
				Case:   1,
				Token:  yyS[yypt-1].token,
				Token2: yyS[yypt-0].token,
			}
			yyVAL.item = lhs
			lhs.Identifier = lx.ident(lhs.Token)
			lhs.Number = lx.number(lhs.Token2)
		}
	case 12:
		{
			yyVAL.item = &NameList{
				Name: yyS[yypt-0].item.(*Name),
			}
		}
	case 13:
		{
			yyVAL.item = &NameList{
				Case:     1,
				NameList: yyS[yypt-1].item.(*NameList),
				Name:     yyS[yypt-0].item.(*Name),
			}
		}
	case 14:
		{
			yyVAL.item = &NameList{
				Case:     2,
				NameList: yyS[yypt-2].item.(*NameList),
				Token:    yyS[yypt-1].token,
				Name:     yyS[yypt-0].item.(*Name),
			}
		}
	case 15:
		{
			lhs := (*Precedence)(nil)
			yyVAL.item = lhs
			//yy:field Identifier interface{} // Name string or literal int.
		}
	case 16:
		{
			lx := yylex.(*lexer)
			lhs := &Precedence{
				Case:   1,
				Token:  yyS[yypt-1].token,
				Token2: yyS[yypt-0].token,
			}
			yyVAL.item = lhs
			lhs.Identifier = lx.ident(lhs.Token2)
		}
	case 17:
		{
			lx := yylex.(*lexer)
			lhs := &Precedence{
				Case:   2,
				Token:  yyS[yypt-2].token,
				Token2: yyS[yypt-1].token,
				Action: yyS[yypt-0].item.(*Action),
			}
			yyVAL.item = lhs
			lhs.Identifier = lx.ident(lhs.Token2)
		}
	case 18:
		{
			yyVAL.item = &Precedence{
				Case:       3,
				Precedence: yyS[yypt-1].item.(*Precedence),
				Token:      yyS[yypt-0].token,
			}
		}
	case 19:
		{
			yyVAL.item = &ReservedWord{
				Token: yyS[yypt-0].token,
			}
		}
	case 20:
		{
			yyVAL.item = &ReservedWord{
				Case:  1,
				Token: yyS[yypt-0].token,
			}
		}
	case 21:
		{
			yyVAL.item = &ReservedWord{
				Case:  2,
				Token: yyS[yypt-0].token,
			}
		}
	case 22:
		{
			yyVAL.item = &ReservedWord{
				Case:  3,
				Token: yyS[yypt-0].token,
			}
		}
	case 23:
		{
			yyVAL.item = &ReservedWord{
				Case:  4,
				Token: yyS[yypt-0].token,
			}
		}
	case 24:
		{
			yyVAL.item = &ReservedWord{
				Case:  5,
				Token: yyS[yypt-0].token,
			}
		}
	case 25:
		{
			lx := yylex.(*lexer)
			lhs := &Rule{
				Token:        yyS[yypt-2].token,
				RuleItemList: yyS[yypt-1].item.(*RuleItemList).reverse(),
				Precedence:   yyS[yypt-0].item.(*Precedence),
			}
			yyVAL.item = lhs
			//yy:field Name *Token
			//yy:field Body []interface{} // For backward compatibility.
			//yy:example "%%\na:\nb:\n\t{\n\t\t//\n\t\tc\n\t}\n%%"
			lx.ruleName = lhs.Token
			lhs.Name = lhs.Token
		}
	case 26:
		{
			lx := yylex.(*lexer)
			lhs := &Rule{
				Case:         1,
				Token:        yyS[yypt-2].token,
				RuleItemList: yyS[yypt-1].item.(*RuleItemList).reverse(),
				Precedence:   yyS[yypt-0].item.(*Precedence),
			}
			yyVAL.item = lhs
			lhs.Name = lx.ruleName
		}
	case 27:
		{
			yyVAL.item = (*RuleItemList)(nil)
		}
	case 28:
		{
			yyVAL.item = &RuleItemList{
				Case:         1,
				RuleItemList: yyS[yypt-1].item.(*RuleItemList),
				Token:        yyS[yypt-0].token,
			}
		}
	case 29:
		{
			yyVAL.item = &RuleItemList{
				Case:         2,
				RuleItemList: yyS[yypt-1].item.(*RuleItemList),
				Action:       yyS[yypt-0].item.(*Action),
			}
		}
	case 30:
		{
			lx := yylex.(*lexer)
			lhs := &RuleList{
				Token:        yyS[yypt-2].token,
				RuleItemList: yyS[yypt-1].item.(*RuleItemList).reverse(),
				Precedence:   yyS[yypt-0].item.(*Precedence),
			}
			yyVAL.item = lhs
			lx.ruleName = lhs.Token
			rule := &Rule{
				Token:        yyS[yypt-2].token,
				Name:         yyS[yypt-2].token,
				RuleItemList: lhs.RuleItemList,
				Precedence:   yyS[yypt-0].item.(*Precedence),
			}
			rule.collect()
			lx.rules = append(lx.rules, rule)
		}
	case 31:
		{
			lx := yylex.(*lexer)
			lhs := &RuleList{
				Case:     1,
				RuleList: yyS[yypt-1].item.(*RuleList),
				Rule:     yyS[yypt-0].item.(*Rule),
			}
			yyVAL.item = lhs
			rule := lhs.Rule
			rule.collect()
			lx.rules = append(lx.rules, rule)
		}
	case 32:
		{
			lx := yylex.(*lexer)
			lhs := &Specification{
				DefinitionList: yyS[yypt-3].item.(*DefinitionList).reverse(),
				Token:          yyS[yypt-2].token,
				RuleList:       yyS[yypt-1].item.(*RuleList).reverse(),
				Tail:           yyS[yypt-0].item.(*Tail),
			}
			yyVAL.item = lhs
			//yy:field Defs  []*Definition // For backward compatibility.
			//yy:field Rules []*Rule       // For backward compatibility.
			lhs.Defs = lx.defs
			lhs.Rules = lx.rules
			lx.spec = lhs
		}
	case 33:
		{
			yyVAL.item = (*Tag)(nil)
		}
	case 34:
		{
			yyVAL.item = &Tag{
				Case:   1,
				Token:  yyS[yypt-2].token,
				Token2: yyS[yypt-1].token,
				Token3: yyS[yypt-0].token,
			}
		}
	case 35:
		{
			lx := yylex.(*lexer)
			lhs := &Tail{
				Token: yyS[yypt-0].token,
			}
			yyVAL.item = lhs
			//yy:field Value string
			lhs.Value = lx.value
		}
	case 36:
		{
			yyVAL.item = (*Tail)(nil)
		}

	}

	if yyEx != nil && yyEx.Reduced(r, exState, &yyVAL) {
		return -1
	}
	goto yystack /* stack new state and value */
}
