// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parser

import __yyfmt__ "fmt"

import (
	"go/token"
)

type yySymType struct {
	yys    int
	token  tkn
	node   Node
	list   []Node
	param  *Param
	params []*Param
}

const _ANDAND = 57346
const _ANDNOT = 57347
const _ASOP = 57348
const _BODY = 57349
const _BREAK = 57350
const _CASE = 57351
const _CHAN = 57352
const _COLAS = 57353
const _COMM = 57354
const _CONST = 57355
const _CONTINUE = 57356
const _DDD = 57357
const _DEC = 57358
const _DEFAULT = 57359
const _DEFER = 57360
const _ELSE = 57361
const _EQ = 57362
const _FALL = 57363
const _FOR = 57364
const _FUNC = 57365
const _GE = 57366
const _GO = 57367
const _GOTO = 57368
const _IF = 57369
const _IGNORE = 57370
const _IMPORT = 57371
const _INC = 57372
const _INTERFACE = 57373
const _LE = 57374
const _LITERAL = 57375
const _LSH = 57376
const _MAP = 57377
const _NAME = 57378
const _NE = 57379
const _OROR = 57380
const _PACKAGE = 57381
const _RANGE = 57382
const _RETURN = 57383
const _RSH = 57384
const _SELECT = 57385
const _STRUCT = 57386
const _SWITCH = 57387
const _TYPE = 57388
const _VAR = 57389
const notPackage = 57390
const notParen = 57391
const preferToRightParen = 57392

var yyToknames = []string{
	" !",
	" %",
	" &",
	" (",
	" *",
	" +",
	" -",
	" .",
	" /",
	" :",
	" <",
	" =",
	" >",
	" [",
	" ^",
	" {",
	" |",
	" ~",
	"_ANDAND",
	"_ANDNOT",
	"_ASOP",
	"_BODY",
	"_BREAK",
	"_CASE",
	"_CHAN",
	"_COLAS",
	"_COMM",
	"_CONST",
	"_CONTINUE",
	"_DDD",
	"_DEC",
	"_DEFAULT",
	"_DEFER",
	"_ELSE",
	"_EQ",
	"_FALL",
	"_FOR",
	"_FUNC",
	"_GE",
	"_GO",
	"_GOTO",
	"_IF",
	"_IGNORE",
	"_IMPORT",
	"_INC",
	"_INTERFACE",
	"_LE",
	"_LITERAL",
	"_LSH",
	"_MAP",
	"_NAME",
	"_NE",
	"_OROR",
	"_PACKAGE",
	"_RANGE",
	"_RETURN",
	"_RSH",
	"_SELECT",
	"_STRUCT",
	"_SWITCH",
	"_TYPE",
	"_VAR",
	"notPackage",
	"notParen",
	" )",
	"preferToRightParen",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

func yy(y yyLexer) *parser                   { return y.(*parser) }
func yyErr(y yyLexer, msg string)            { yy(y).Error(msg) }
func yyErrPos(y yyLexer, n Node, msg string) { yy(y).errPos(n.Pos(), msg) }
func yyFset(y yyLexer) *token.FileSet        { return yy(y).fset }
func yyFScope(y yyLexer) *Scope              { return yy(y).pkgScope }
func yyScope(y yyLexer) *Scope               { return yy(y).currentScope }

func yyTLD(y yyLexer, n Node) {
	p := yy(y)
	p.ast = append(p.ast, n)
	switch x := n.(type) {
	case *ConstDecl:
		p.pkgScope.declare(p, x.Name.Lit, n)
	case *VarDecl:
		p.pkgScope.declare(p, x.Name.Lit, n)
	case *TypeDecl:
		p.pkgScope.declare(p, x.Name.Lit, n)
	case *FuncDecl:
		p.pkgScope.declare(p, x.Name.Lit, n)
	case *Package:
		switch ex0, ok := p.pkgScope.Names[dlrPkgName]; {
		case ok:
			ex := ex0.(*Package)
			if ex.Name.Lit == x.Name.Lit {
				break
			}

			p.errPos(n.Pos(), "%s redeclared, previous declaration at %s", x.Name.Lit, p.fset.Position(ex.Pos()))
		default:
			p.pkgScope.declare(p, dlrPkgName, n)
		}
	case *Import:
		// not handled here
	default:
		panic("internal error")
	}
}

func yyTLDs(y yyLexer, l []Node) {
	for _, v := range l {
		yyTLD(y, v)
	}
}

var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 9,
	1, 2,
	70, 15,
	-2, 0,
	-1, 35,
	15, 249,
	29, 249,
	73, 249,
	-2, 41,
	-1, 43,
	13, 133,
	-2, 138,
	-1, 61,
	7, 157,
	-2, 189,
	-1, 62,
	7, 158,
	-2, 159,
	-1, 93,
	27, 225,
	35, 225,
	70, 225,
	71, 225,
	-2, 0,
	-1, 98,
	27, 239,
	35, 239,
	70, 239,
	71, 239,
	-2, 114,
	-1, 102,
	27, 240,
	35, 240,
	70, 240,
	71, 240,
	-2, 114,
	-1, 159,
	2, 189,
	7, 157,
	19, 189,
	25, 189,
	-2, 149,
	-1, 160,
	7, 158,
	19, 159,
	25, 159,
	-2, 150,
	-1, 167,
	70, 225,
	71, 225,
	-2, 0,
	-1, 211,
	70, 225,
	71, 225,
	-2, 0,
	-1, 226,
	70, 225,
	71, 225,
	-2, 0,
	-1, 253,
	70, 225,
	71, 225,
	-2, 0,
	-1, 292,
	51, 210,
	70, 210,
	71, 210,
	-2, 137,
	-1, 332,
	27, 225,
	35, 225,
	70, 225,
	71, 225,
	-2, 0,
	-1, 352,
	7, 152,
	19, 152,
	25, 152,
	-2, 143,
	-1, 353,
	7, 153,
	19, 153,
	25, 153,
	-2, 144,
	-1, 354,
	7, 154,
	19, 154,
	25, 154,
	-2, 145,
	-1, 355,
	7, 155,
	19, 155,
	25, 155,
	-2, 146,
	-1, 361,
	27, 225,
	35, 225,
	70, 225,
	71, 225,
	-2, 0,
	-1, 476,
	7, 156,
	19, 156,
	25, 156,
	-2, 147,
}

const yyNprod = 271
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1798

var yyAct = []int{

	35, 62, 400, 169, 198, 252, 268, 285, 199, 343,
	345, 80, 139, 21, 346, 371, 212, 291, 86, 392,
	263, 214, 140, 222, 295, 287, 262, 78, 195, 299,
	85, 274, 217, 259, 106, 81, 72, 317, 215, 337,
	468, 157, 132, 460, 275, 461, 36, 255, 430, 393,
	358, 137, 137, 142, 137, 156, 160, 316, 333, 462,
	135, 256, 499, 361, 412, 171, 334, 175, 83, 366,
	176, 361, 410, 502, 177, 155, 277, 130, 491, 338,
	105, 276, 197, 158, 138, 490, 164, 197, 361, 398,
	197, 131, 333, 197, 449, 298, 85, 440, 361, 467,
	334, 3, 416, 463, 437, 132, 208, 191, 228, 429,
	186, 229, 230, 231, 232, 233, 234, 235, 236, 237,
	238, 239, 240, 241, 242, 243, 244, 245, 246, 247,
	248, 106, 106, 251, 202, 132, 331, 454, 221, 361,
	360, 413, 292, 380, 156, 160, 266, 368, 260, 267,
	167, 288, 289, 76, 313, 309, 182, 183, 41, 294,
	302, 300, 258, 160, 272, 254, 77, 71, 8, 495,
	6, 273, 158, 484, 481, 197, 480, 249, 250, 479,
	197, 197, 478, 197, 207, 6, 6, 95, 95, 476,
	158, 103, 466, 74, 106, 452, 6, 444, 292, 203,
	190, 297, 197, 270, 200, 179, 106, 436, 427, 61,
	185, 426, 205, 423, 64, 284, 411, 409, 197, 301,
	396, 197, 197, 390, 389, 65, 386, 201, 384, 197,
	305, 364, 308, 315, 269, 298, 312, 321, 101, 6,
	304, 374, 377, 314, 323, 265, 70, 375, 11, 319,
	66, 6, 311, 76, 372, 318, 106, 106, 322, 69,
	91, 73, 188, 187, 159, 76, 156, 160, 137, 347,
	194, 226, 137, 353, 347, 174, 354, 325, 197, 197,
	355, 329, 292, 359, 87, 324, 339, 375, 197, 375,
	196, 328, 6, 74, 158, 196, 6, 464, 196, 422,
	336, 196, 249, 250, 106, 74, 307, 6, 6, 197,
	357, 106, 253, 367, 373, 376, 383, 405, 197, 84,
	406, 394, 79, 379, 407, 288, 289, 414, 415, 257,
	387, 6, 397, 94, 156, 160, 56, 385, 408, 156,
	160, 290, 296, 171, 403, 166, 428, 57, 349, 44,
	388, 165, 197, 159, 211, 433, 434, 395, 197, 425,
	418, 154, 158, 97, 417, 141, 6, 158, 197, 6,
	168, 159, 292, 181, 310, 424, 441, 303, 98, 102,
	99, 99, 431, 196, 197, 442, 419, 166, 196, 196,
	342, 196, 446, 165, 93, 378, 197, 143, 297, 451,
	445, 144, 180, 447, 197, 450, 438, 145, 281, 456,
	196, 405, 458, 405, 406, 382, 406, 180, 407, 457,
	407, 453, 180, 227, 220, 180, 196, 137, 180, 196,
	196, 347, 223, 472, 347, 347, 459, 196, 403, 224,
	403, 469, 13, 474, 475, 470, 465, 163, 18, 15,
	12, 10, 9, 7, 4, 218, 1, 213, 197, 189,
	477, 206, 485, 486, 156, 160, 394, 283, 104, 171,
	286, 344, 488, 483, 482, 159, 293, 487, 137, 184,
	489, 352, 405, 494, 425, 406, 196, 196, 210, 407,
	347, 335, 158, 498, 492, 493, 196, 405, 501, 500,
	406, 391, 14, 88, 407, 23, 67, 24, 136, 403,
	180, 53, 200, 179, 2, 180, 180, 196, 180, 172,
	193, 26, 64, 68, 403, 404, 196, 290, 54, 443,
	381, 25, 22, 65, 134, 201, 192, 180, 64, 296,
	133, 204, 401, 159, 209, 60, 101, 63, 159, 65,
	90, 421, 420, 180, 70, 58, 180, 180, 66, 6,
	196, 59, 101, 471, 180, 330, 196, 69, 332, 100,
	70, 0, 55, 5, 66, 6, 196, 0, 194, 43,
	0, 75, 0, 69, 200, 179, 0, 82, 82, 89,
	92, 0, 196, 0, 64, 0, 0, 0, 96, 96,
	0, 0, 96, 0, 196, 65, 0, 201, 180, 0,
	0, 0, 196, 180, 180, 0, 0, 0, 101, 404,
	0, 404, 0, 180, 0, 0, 70, 0, 0, 278,
	66, 6, 42, 0, 279, 280, 0, 282, 0, 69,
	0, 0, 0, 75, 180, 0, 0, 0, 0, 82,
	370, 0, 180, 180, 82, 0, 306, 89, 0, 0,
	0, 216, 0, 43, 0, 0, 196, 0, 0, 0,
	0, 0, 0, 159, 0, 320, 0, 0, 146, 147,
	148, 149, 150, 151, 152, 153, 0, 180, 402, 179,
	404, 0, 0, 180, 200, 179, 0, 0, 64, 0,
	0, 0, 0, 180, 64, 404, 0, 0, 0, 65,
	0, 201, 0, 0, 264, 65, 0, 201, 0, 180,
	219, 0, 101, 0, 0, 0, 0, 351, 101, 0,
	70, 180, 362, 363, 66, 6, 70, 43, 0, 180,
	66, 6, 369, 69, 0, 0, 180, 0, 180, 69,
	0, 0, 96, 96, 0, 124, 125, 0, 122, 118,
	119, 0, 123, 351, 82, 0, 0, 0, 121, 0,
	120, 124, 125, 126, 122, 118, 119, 0, 123, 0,
	114, 43, 117, 0, 121, 0, 120, 0, 111, 126,
	216, 0, 0, 180, 146, 153, 43, 216, 0, 0,
	0, 0, 127, 0, 112, 0, 278, 0, 116, 0,
	128, 0, 435, 0, 0, 0, 115, 180, 127, 200,
	179, 113, 439, 43, 0, 0, 128, 0, 0, 64,
	0, 49, 180, 46, 54, 161, 47, 48, 0, 0,
	65, 0, 201, 0, 64, 51, 0, 0, 50, 0,
	455, 365, 0, 101, 0, 65, 0, 162, 0, 0,
	0, 70, 0, 0, 0, 66, 6, 0, 101, 0,
	75, 0, 82, 0, 69, 0, 70, 0, 55, 82,
	66, 6, 0, 89, 0, 399, 216, 0, 0, 69,
	0, 0, 0, 0, 0, 261, 0, 0, 0, 0,
	0, 0, 43, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 225, 0, 49,
	0, 46, 54, 45, 47, 48, 0, 0, 0, 0,
	0, 43, 64, 51, 226, 0, 50, 0, 96, 0,
	96, 28, 0, 65, 0, 52, 34, 29, 448, 0,
	96, 31, 216, 0, 27, 37, 101, 0, 30, 32,
	40, 0, 0, 0, 70, 0, 55, 0, 66, 6,
	0, 0, 216, 0, 33, 0, 39, 69, 38, 19,
	17, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 16, 0, 49, 0, 46, 54, 45, 47, 48,
	0, 0, 0, 0, 0, 0, 64, 51, 0, 0,
	50, 0, 0, 0, 0, 28, 0, 65, 0, 52,
	34, 29, 0, 0, 82, 31, 216, 0, 27, 37,
	20, 0, 30, 32, 40, 0, 0, 0, 70, 0,
	55, 0, 66, 6, 0, 0, 0, 0, 33, 0,
	39, 69, 38, 19, 17, 124, 125, 0, 122, 118,
	119, 0, 123, 0, 114, 0, 117, 0, 121, 0,
	120, 0, 111, 126, 0, 0, 0, 0, 124, 125,
	129, 122, 0, 0, 0, 123, 0, 0, 112, 0,
	0, 0, 116, 0, 0, 0, 126, 0, 0, 0,
	115, 0, 127, 0, 0, 113, 110, 0, 0, 0,
	128, 124, 125, 0, 122, 118, 119, 0, 123, 0,
	114, 0, 117, 358, 121, 127, 120, 0, 111, 126,
	0, 0, 0, 128, 124, 125, 129, 122, 118, 119,
	0, 123, 0, 114, 112, 117, 0, 121, 116, 120,
	0, 111, 126, 107, 0, 0, 115, 0, 127, 129,
	0, 113, 110, 109, 0, 0, 128, 112, 0, 0,
	0, 116, 0, 0, 0, 0, 0, 108, 341, 115,
	0, 127, 0, 0, 113, 110, 0, 0, 49, 128,
	46, 54, 161, 47, 48, 0, 0, 0, 0, 0,
	0, 64, 51, 0, 0, 50, 0, 0, 0, 0,
	0, 0, 65, 0, 162, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 101, 0, 0, 0, 0,
	0, 0, 0, 70, 0, 55, 0, 66, 6, 0,
	0, 0, 0, 0, 0, 0, 69, 0, 340, 124,
	125, 0, 122, 118, 119, 0, 123, 497, 114, 0,
	117, 0, 121, 0, 120, 0, 111, 126, 0, 0,
	0, 0, 124, 125, 129, 122, 118, 119, 0, 123,
	496, 114, 112, 117, 0, 121, 116, 120, 0, 111,
	126, 0, 0, 0, 115, 0, 127, 129, 0, 113,
	110, 0, 0, 0, 128, 112, 0, 0, 0, 116,
	0, 0, 0, 0, 200, 179, 0, 115, 0, 127,
	0, 0, 113, 110, 64, 124, 125, 128, 122, 118,
	119, 0, 123, 432, 114, 271, 117, 201, 121, 0,
	120, 0, 111, 126, 356, 179, 0, 0, 101, 0,
	129, 0, 0, 0, 64, 0, 70, 0, 112, 0,
	66, 6, 116, 0, 0, 65, 0, 350, 0, 69,
	115, 0, 127, 0, 0, 113, 110, 0, 101, 49,
	128, 46, 54, 45, 47, 48, 70, 0, 0, 0,
	66, 6, 64, 51, 348, 0, 50, 124, 125, 69,
	122, 118, 119, 65, 123, 52, 114, 0, 117, 0,
	121, 0, 120, 0, 111, 126, 101, 0, 0, 0,
	0, 0, 129, 0, 70, 0, 55, 0, 66, 6,
	112, 0, 0, 0, 116, 0, 0, 69, 0, 0,
	178, 179, 115, 0, 127, 0, 0, 113, 110, 0,
	64, 49, 128, 46, 54, 45, 47, 48, 0, 0,
	0, 65, 0, 173, 64, 51, 473, 49, 50, 46,
	54, 45, 47, 48, 101, 65, 0, 52, 0, 0,
	64, 51, 70, 0, 50, 0, 66, 6, 101, 0,
	0, 65, 0, 52, 0, 69, 70, 0, 55, 0,
	66, 6, 0, 0, 101, 0, 0, 0, 0, 69,
	0, 0, 70, 0, 55, 0, 66, 6, 0, 0,
	0, 327, 0, 0, 49, 69, 46, 54, 45, 47,
	48, 0, 0, 0, 0, 0, 0, 64, 51, 0,
	0, 50, 0, 0, 0, 0, 0, 0, 65, 0,
	52, 49, 0, 46, 54, 45, 47, 48, 0, 0,
	0, 101, 0, 0, 64, 51, 0, 0, 50, 70,
	0, 55, 0, 66, 6, 65, 0, 52, 326, 0,
	170, 0, 69, 0, 0, 0, 0, 0, 101, 49,
	0, 46, 54, 45, 47, 48, 70, 0, 55, 0,
	66, 6, 64, 51, 0, 0, 50, 0, 0, 69,
	0, 0, 0, 65, 0, 52, 49, 0, 46, 54,
	161, 47, 48, 0, 0, 0, 101, 0, 0, 64,
	51, 0, 0, 50, 70, 0, 55, 0, 66, 6,
	65, 0, 162, 0, 0, 0, 0, 69, 0, 0,
	0, 0, 0, 101, 0, 0, 0, 0, 0, 0,
	0, 70, 0, 55, 0, 66, 6, 0, 124, 125,
	0, 122, 118, 119, 69, 123, 0, 114, 0, 117,
	0, 121, 0, 120, 0, 111, 126, 0, 0, 0,
	49, 0, 46, 54, 45, 47, 48, 0, 0, 0,
	0, 112, 0, 64, 51, 116, 0, 50, 0, 0,
	0, 0, 0, 115, 271, 127, 52, 0, 113, 110,
	0, 0, 0, 128, 0, 0, 0, 101, 0, 0,
	0, 0, 0, 0, 0, 70, 0, 55, 0, 66,
	6, 0, 124, 125, 0, 122, 118, 119, 69, 123,
	0, 114, 0, 117, 0, 121, 0, 120, 0, 0,
	126, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 112, 0, 0, 0, 116,
	0, 0, 0, 0, 0, 0, 0, 115, 0, 127,
	0, 0, 113, 0, 0, 0, 0, 128,
}
var yyPact = []int{

	44, -1000, -1000, 185, -1000, 98, -1000, 201, -1000, 989,
	97, 254, 96, -1000, -1000, -1000, -1000, 315, 312, 277,
	253, -1000, -1000, -1000, -1000, -1000, 381, -1000, 185, 185,
	521, 521, 185, 1585, -1000, 1129, 62, 1585, 1585, 346,
	1585, -1000, -1000, -1000, 390, 1585, 1585, 1585, 1585, 1585,
	1585, 1585, 1585, 342, 1612, -1000, -1000, -1000, 440, 326,
	-1000, -1000, -1000, 368, 1547, 1433, 356, -1000, -1000, 326,
	326, -1000, -1000, 142, -1000, 212, 211, -1000, -1000, 132,
	505, -1000, -1000, -1000, 131, 197, -1000, 116, 812, -1000,
	335, 687, 417, 915, -1000, -1000, -1000, -1000, -1000, 390,
	-1000, 416, -1000, -1000, -1000, -31, 1392, 1585, -1000, -1000,
	1585, 1585, 1585, 1585, 1585, 1585, 1585, 1585, 1585, 1585,
	1585, 1585, 1585, 1585, 1585, 1585, 1585, 1585, 1585, 1585,
	1585, 1585, 1585, -1000, 287, 95, -1000, -1000, 32, 304,
	92, -1000, 287, 827, 238, 1585, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 166, 1392, -1000, -1000, -1000,
	-1000, 1612, 1686, 1585, -1000, -1000, -1000, 915, -1000, 9,
	4, 1392, -1000, 812, -1000, -1000, -1000, -1000, 812, 812,
	397, 812, 144, 88, 91, -1000, -1000, -1000, -1000, 90,
	-1000, -1000, 362, 1585, 185, -1000, -1000, -1000, -1000, -1000,
	812, 278, 85, -1000, 359, 1585, 84, -1000, -1000, -1000,
	-1000, 915, 165, -16, -1000, -1000, 687, -1000, -1000, 812,
	687, -1000, -1000, -1000, -1000, -1000, 915, 687, 1392, 766,
	1737, 750, 750, 750, 750, 750, 750, 1073, 1073, 1073,
	1073, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 1663, -31,
	-31, 1392, -1000, 915, 1585, 1520, 1463, -1000, 1585, 65,
	-1000, -1000, 6, -1000, -1000, 1184, 1106, 377, 1375, 329,
	-1000, 1337, 1050, 1375, 69, -1000, 812, 812, -1000, 163,
	-1000, 185, -3, 77, -1000, -1000, 577, 203, 233, 235,
	-1000, -1000, 384, 73, -1000, -1000, 408, -1000, 193, 160,
	242, 158, 185, 1585, -31, -1000, 156, 812, 155, 185,
	1585, -31, 152, 185, 18, 681, 687, -1000, -1000, -1000,
	-1000, 149, 1, 148, -7, 71, 1585, 1585, 31, -1000,
	-1000, -1000, 915, 1612, 373, 262, 145, -23, 1612, 143,
	140, -1000, 1585, 38, -25, -1000, -1000, 1320, -1000, -1000,
	1307, -1000, -1000, -1000, -1000, -1000, 812, 139, -1000, 33,
	-1000, 915, -1000, -1000, -1000, -1000, 812, 26, 318, 203,
	185, -1000, -1000, 129, 193, 384, 203, 193, 185, 23,
	228, -1000, 687, 127, -1000, -1000, -1000, -1000, -31, -1000,
	-1000, 67, -1000, -1000, 197, -31, -1000, -1000, -1000, 402,
	-1000, -1000, 687, -1000, -1000, -1000, -1000, -1000, -1000, 681,
	-1000, 681, -1000, 1585, 1392, 1392, -1000, 28, 30, -1000,
	-1000, -1000, 252, -1000, 124, -1000, -1000, -1000, 27, -1000,
	1375, -1000, 1447, 1375, 1375, 121, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 203, 114, -1000, 111, -1000, -1000,
	-1000, 108, -1000, 106, 185, 359, 687, 105, -1000, -1000,
	-1000, 1585, 1585, 1612, 1585, -1000, -1000, -1000, 1585, -1000,
	-1000, -1000, 1392, -1000, 14, 7, -1000, -1000, 203, 203,
	681, -1000, -1000, 101, -1000, 1267, 1244, 287, -10, 1375,
	-1000, -1000, -1000, -1000, -1000, 681, -1000, -1000, -1000, -1000,
	2, -1000, -1000,
}
var yyPgo = []int{

	0, 14, 568, 565, 563, 561, 23, 49, 19, 555,
	35, 8, 552, 551, 25, 0, 20, 550, 547, 545,
	542, 209, 540, 534, 532, 12, 531, 36, 530, 24,
	523, 10, 521, 86, 336, 38, 158, 455, 439, 41,
	519, 3, 15, 333, 22, 1, 514, 17, 349, 511,
	347, 4, 508, 28, 507, 13, 7, 506, 505, 569,
	18, 503, 632, 502, 9, 33, 432, 501, 11, 491,
	46, 26, 488, 479, 476, 471, 5, 470, 468, 44,
	31, 467, 461, 27, 459, 21, 32, 457, 2, 16,
	456, 454, 453, 452, 451, 29, 450, 448, 37, 6,
}
var yyR1 = []int{

	0, 91, 90, 46, 46, 92, 92, 94, 94, 94,
	27, 27, 27, 73, 73, 96, 96, 96, 96, 96,
	66, 66, 66, 66, 66, 66, 66, 66, 66, 66,
	97, 83, 83, 83, 7, 7, 8, 8, 8, 61,
	60, 55, 55, 55, 55, 55, 55, 2, 2, 2,
	2, 6, 3, 65, 65, 76, 52, 52, 23, 23,
	23, 22, 24, 25, 25, 26, 13, 69, 69, 12,
	12, 58, 54, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 62, 62, 62, 62, 62, 62,
	62, 62, 62, 50, 50, 50, 49, 49, 49, 49,
	49, 49, 49, 49, 49, 49, 49, 49, 49, 49,
	99, 31, 1, 1, 4, 4, 48, 48, 16, 16,
	35, 33, 33, 36, 10, 43, 43, 59, 34, 32,
	86, 86, 37, 37, 37, 37, 37, 37, 39, 39,
	39, 39, 40, 40, 40, 40, 40, 9, 9, 5,
	20, 20, 20, 20, 20, 11, 11, 45, 45, 45,
	45, 45, 45, 45, 51, 53, 57, 57, 30, 30,
	63, 17, 17, 21, 72, 72, 88, 88, 88, 18,
	19, 19, 93, 93, 84, 84, 67, 67, 82, 82,
	81, 81, 74, 74, 56, 56, 56, 56, 56, 56,
	47, 47, 14, 29, 29, 29, 28, 85, 85, 85,
	85, 87, 87, 89, 89, 79, 79, 79, 79, 79,
	38, 38, 38, 38, 38, 38, 38, 38, 38, 38,
	38, 38, 38, 80, 80, 77, 77, 68, 68, 70,
	70, 71, 71, 75, 75, 75, 75, 64, 64, 95,
	95, 98, 98, 41, 41, 78, 78, 44, 44, 42,
	42,
}
var yyR2 = []int{

	0, 0, 4, 0, 3, 0, 3, 2, 5, 3,
	1, 2, 2, 1, 3, 0, 1, 1, 1, 1,
	2, 5, 3, 2, 5, 7, 3, 2, 5, 3,
	1, 2, 4, 3, 4, 3, 1, 2, 1, 1,
	2, 1, 3, 3, 3, 2, 2, 3, 5, 5,
	2, 3, 2, 0, 2, 3, 4, 4, 5, 1,
	1, 2, 2, 1, 3, 5, 4, 0, 2, 0,
	2, 5, 4, 1, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 1, 2, 2, 2, 2, 2,
	2, 2, 2, 3, 5, 6, 1, 1, 3, 5,
	5, 4, 6, 8, 1, 5, 5, 5, 7, 1,
	0, 3, 1, 4, 1, 4, 1, 3, 1, 1,
	1, 1, 1, 1, 1, 0, 1, 1, 1, 1,
	1, 2, 1, 1, 1, 1, 1, 3, 1, 1,
	1, 2, 1, 1, 1, 1, 3, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 3, 4, 4, 2,
	3, 5, 1, 1, 2, 3, 5, 3, 5, 3,
	3, 5, 8, 5, 0, 3, 0, 1, 3, 1,
	4, 2, 0, 3, 1, 3, 1, 3, 1, 3,
	1, 3, 1, 3, 3, 2, 4, 3, 5, 5,
	1, 3, 1, 2, 1, 3, 4, 1, 2, 2,
	1, 1, 3, 0, 2, 0, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 3, 1, 2, 2, 2,
	2, 2, 2, 1, 3, 1, 3, 1, 3, 1,
	3, 1, 3, 1, 1, 3, 3, 0, 2, 0,
	1, 0, 1, 0, 1, 0, 1, 0, 1, 0,
	1,
}
var yyChk = []int{

	-1000, -90, -46, 57, -91, -59, 54, -92, 70, -93,
	-94, 47, -96, -66, -63, -38, 2, 65, -97, 64,
	41, -55, -24, -58, -54, -26, -32, 39, 26, 32,
	43, 36, 44, 59, 31, -15, -70, 40, 63, 61,
	45, -36, -62, -59, -48, 8, 6, 9, 10, 4,
	21, 18, 30, -49, 7, 51, -34, -50, -9, -5,
	-19, -21, -45, -18, 17, 28, 53, -57, -30, 62,
	49, 70, -27, 7, 51, -59, 11, 70, -83, 7,
	-68, -10, -59, -7, 7, -68, -60, 7, -61, -59,
	-17, 7, -59, 13, -43, -36, -59, -43, -50, -48,
	-59, 41, -50, -36, -78, -70, -15, 24, 48, 34,
	56, 22, 38, 55, 14, 50, 42, 16, 9, 10,
	20, 18, 8, 12, 5, 6, 23, 52, 60, 30,
	15, 29, 73, -22, -23, -44, -52, -55, -70, -25,
	-44, 19, -25, 7, 11, 17, -62, -62, -62, -62,
	-62, -62, -62, -62, 19, -16, -15, -39, -53, -21,
	-45, 8, 30, 7, -33, 25, 19, -33, 2, -41,
	33, -15, -40, 30, -21, -45, -51, -11, 7, 8,
	-34, 17, -33, -33, -73, 68, -27, 51, 51, -84,
	68, -83, -37, 15, 73, -53, -21, -45, -51, -11,
	7, 30, -7, 68, -37, 15, -82, 68, -60, -37,
	-72, 19, -89, -87, -85, -35, -59, -86, -37, 33,
	7, -79, -6, -66, -38, 2, 19, 7, -15, -15,
	-15, -15, -15, -15, -15, -15, -15, -15, -15, -15,
	-15, -15, -15, -15, -15, -15, -15, -15, -15, -70,
	-70, -15, -76, 25, 70, 15, 29, 25, 70, -65,
	-76, 68, -71, -16, -59, 7, -15, -41, -99, 68,
	-39, 28, -15, -99, -80, -79, 72, 72, -37, -37,
	-37, 11, -37, -81, 71, -56, -77, -14, 7, 8,
	-36, -47, 54, -74, 71, -29, -36, -47, 7, -95,
	70, -95, 70, 15, -70, -10, -37, 28, -95, 70,
	15, -70, -95, 70, -80, 68, 73, -98, -35, -86,
	-37, -89, -80, -89, -80, -44, 58, 58, -65, -44,
	-3, 71, -2, 27, 35, -69, -98, 33, 73, -16,
	64, 72, 13, -64, -75, -31, -1, -15, 19, 19,
	30, -37, -21, -45, -51, -11, 7, -98, 73, -64,
	71, 70, -37, -37, 68, -59, 72, -95, 70, -37,
	73, -42, 51, -14, 8, 54, -14, 7, 11, -95,
	70, -28, 7, -47, 68, -27, 68, -83, -70, 68,
	68, -67, -8, -7, -68, -70, 68, -60, 71, -59,
	-88, -20, 7, -53, -21, -45, -51, -11, -85, 68,
	71, 68, 71, 70, -15, -15, 71, -80, -71, 13,
	-12, -13, 37, 68, -98, -16, 68, 68, -41, 71,
	73, -98, 13, -99, -99, -37, 68, 71, -79, -37,
	71, -56, -42, -36, 68, -14, -42, -14, -59, 71,
	-29, -89, 68, -95, 70, -37, 7, -89, -88, -44,
	13, 15, 29, 73, 45, -6, 68, 72, 13, -31,
	-1, -4, -15, 19, -64, -64, 68, -42, 68, 68,
	68, 68, -8, -89, 68, -15, -15, -25, -41, -99,
	71, 71, -42, -42, -88, 68, 13, 13, -76, 72,
	-64, -88, 71,
}
var yyDef = []int{

	3, -2, 1, 0, 5, 0, 137, 192, 4, -2,
	0, 0, 0, 16, 17, 18, 19, 0, 0, 0,
	0, 230, 231, 232, 233, 234, 0, 236, 135, 135,
	0, 0, 0, 265, 30, -2, 0, 267, 267, 0,
	267, 139, 73, -2, 94, 0, 0, 0, 0, 0,
	0, 0, 0, 126, 0, 106, 107, 114, 0, 0,
	119, -2, -2, 0, 263, 0, 0, 172, 173, 0,
	0, 6, 7, 0, 10, 0, 0, 193, 20, 0,
	0, 247, 134, 23, 0, 0, 27, 0, 0, 39,
	184, 223, 0, -2, 237, 136, 133, 238, -2, 0,
	138, 0, -2, 241, 242, 266, 249, 0, 45, 46,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 62, 0, 59, 60, 268, 0, 0,
	63, 53, 0, 0, 0, 263, 95, 96, 97, 98,
	99, 100, 101, 102, 120, 0, 128, 129, 148, -2,
	-2, 0, 0, 0, 120, 131, 132, -2, 191, 0,
	0, 264, 169, 0, 152, 153, 154, 155, 0, 0,
	165, 0, 0, 0, 259, 9, 13, 11, 12, 259,
	22, 194, 31, 0, 0, 142, 143, 144, 145, 146,
	0, 0, 259, 26, 0, 0, 259, 29, 198, 40,
	180, -2, 0, 261, 221, 217, 138, 220, 130, 140,
	223, 235, 226, 227, 228, 229, -2, 223, 42, 74,
	75, 76, 77, 78, 79, 80, 81, 82, 83, 84,
	85, 86, 87, 88, 89, 90, 91, 92, 93, 43,
	44, 250, 61, -2, 267, 0, 0, 53, 267, 0,
	67, 103, 261, 251, 108, 0, 264, 0, 257, 127,
	151, 0, 261, 257, 0, 243, 0, 0, 170, 0,
	174, 0, 0, 259, 177, 200, 0, 269, 0, 0,
	245, 212, -2, 259, 179, 202, 0, 214, 0, 0,
	260, 0, 260, 0, 33, 248, 0, 0, 0, 260,
	0, 35, 0, 260, 0, 186, 262, 224, 218, 219,
	141, 0, 0, 0, 0, 0, 0, 0, 0, 64,
	54, 72, -2, 0, 0, 69, 0, 261, 262, 0,
	0, 111, 263, 0, 261, 253, 254, 122, 120, 120,
	0, 175, -2, -2, -2, -2, 0, 0, 262, 0,
	190, -2, 167, 168, 156, 166, 0, 0, 260, 269,
	0, 205, 270, 0, 0, 210, 269, 0, 0, 0,
	260, 213, 223, 0, 8, 14, 21, 195, 32, 147,
	24, 259, 196, 36, 38, 34, 28, 199, 185, 138,
	183, 187, 223, 160, 161, 162, 163, 164, 222, 186,
	51, 186, 55, 267, 56, 57, 71, 52, 0, 50,
	65, 68, 0, 104, 0, 252, 109, 110, 0, 117,
	262, 258, 0, 257, 257, 0, 115, 116, 244, 171,
	176, 201, 204, 246, 269, 0, 207, 0, 211, 178,
	203, 0, 215, 0, 260, 37, 223, 0, 181, 58,
	47, 0, 0, 0, 267, 70, 105, 112, 263, 255,
	256, 121, 124, 120, 0, 0, -2, 206, 269, 269,
	186, 25, 197, 0, 188, 0, 0, 0, 0, 257,
	123, 118, 208, 209, 216, 186, 48, 49, 66, 113,
	0, 182, 125,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 4, 3, 3, 3, 5, 6, 3,
	7, 68, 8, 9, 73, 10, 11, 12, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 13, 70,
	14, 15, 16, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 17, 3, 72, 18, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 19, 20, 71, 21,
}
var yyTok2 = []int{

	2, 3, 22, 23, 24, 25, 26, 27, 28, 29,
	30, 31, 32, 33, 34, 35, 36, 37, 38, 39,
	40, 41, 42, 43, 44, 45, 46, 47, 48, 49,
	50, 51, 52, 53, 54, 55, 56, 57, 58, 59,
	60, 61, 62, 63, 64, 65, 66, 67, 69,
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
		{ //46
			yyTLD(yylex, yyS[yypt-0].node)
		}
	case 3:
		{ //60
			yyErr(yylex, "package statement must be first")
			goto ret1
		}
	case 4:
		{ //64
			yyVAL.node = &Package{yyS[yypt-2].token.p(), yyS[yypt-1].node.(*Ident)}
		}
	case 7:
		{
			yyTLD(yylex, yyS[yypt-0].node)
		}
	case 8:
		{ //83
			yyTLDs(yylex, yyS[yypt-2].list)
		}
	case 10:
		{ //93
			yyVAL.node = newImport(yylex, (*Ident)(nil), newLiteral(yyS[yypt-0].token))
		}
	case 11:
		{ //97
			yyVAL.node = newImport(yylex, yyS[yypt-1].node, newLiteral(yyS[yypt-0].token))
		}
	case 12:
		{ //101
			yyVAL.node = newImport(yylex, &Ident{yyS[yypt-1].token.p(), "."}, newLiteral(yyS[yypt-0].token))
		}
	case 13:
		{ //107
			yyVAL.list = []Node{yyS[yypt-0].node}
		}
	case 14:
		{ //111
			yyVAL.list = append(yyS[yypt-2].list, yyS[yypt-0].node)
		}
	case 16:
		{ //120
			yyTLDs(yylex, yyS[yypt-0].list)
		}
	case 17:
		{ //124
			yyTLD(yylex, yyS[yypt-0].node)
		}
	case 18:
		{ //128
			yyErrPos(yylex, yyS[yypt-0].node, "non-declaration statement outside function body")
		}
	case 20:
		{ //138
			yyVAL.list = yyS[yypt-0].list
		}
	case 21:
		{ //142
			yyVAL.list = yyS[yypt-2].list
		}
	case 22:
		{ //146
			yyVAL.list = nil
		}
	case 23:
		{ //150
			yyVAL.list = newConstDecls(yylex, []Node{yyS[yypt-0].node})
		}
	case 24:
		{ //154
			yyVAL.list = newConstDecls(yylex, []Node{yyS[yypt-2].node})
		}
	case 25:
		{ //158
			yyVAL.list = newConstDecls(yylex, append([]Node{yyS[yypt-4].node}, yyS[yypt-2].list...))
		}
	case 26:
		{ //162
			yyVAL.list = nil
		}
	case 27:
		{ //166
			yyVAL.list = []Node{yyS[yypt-0].node}
		}
	case 28:
		{ //170
			yyVAL.list = yyS[yypt-2].list
		}
	case 29:
		{ //174
			yyVAL.list = nil
		}
	case 30:
		{ //180
			p := yy(yylex)
			p.constExpr, p.constIota, p.constType = nil, 0, nil
		}
	case 31:
		{ //186
			yyVAL.list = newVarDecls(yyS[yypt-1].list, yyS[yypt-0].node, nil)
		}
	case 32:
		{ //190
			yyVAL.list = newVarDecls(yyS[yypt-3].list, yyS[yypt-2].node, yyS[yypt-0].list)
		}
	case 33:
		{ //194
			yyVAL.list = newVarDecls(yyS[yypt-2].list, nil, yyS[yypt-0].list)
		}
	case 34:
		{ //200
			yyVAL.node = newConstSpec(yylex, yyS[yypt-3].list, yyS[yypt-2].node, yyS[yypt-0].list)
		}
	case 35:
		{ //204
			yyVAL.node = newConstSpec(yylex, yyS[yypt-2].list, nil, yyS[yypt-0].list)
		}
	case 36:
		yyVAL.node = yyS[yypt-0].node
	case 37:
		{ //214
			panic(".y:215") //TODO ???
			//yyErrPos(yylex, $2, "const declaration cannot have type without expression")
		}
	case 38:
		{ //218
			yyVAL.node = newConstSpec(yylex, yyS[yypt-0].list, nil, nil)
		}
	case 39:
		yyVAL.node = yyS[yypt-0].node
	case 40:
		{ //230
			yyVAL.node = &TypeDecl{yyS[yypt-1].node.p(), yyS[yypt-1].node.(*Ident), yyS[yypt-0].node}
		}
	case 41:
		yyVAL.node = yyS[yypt-0].node
	case 42:
		{ //240
			yyVAL.node = &Assignment{yyS[yypt-1].token.p(), yyS[yypt-1].token.tok, []Node{yyS[yypt-2].node}, []Node{yyS[yypt-0].node}}
		}
	case 43:
		{ //244
			yyVAL.node = &Assignment{yyS[yypt-1].token.p(), yyS[yypt-1].token.tok, yyS[yypt-2].list, yyS[yypt-0].list}
		}
	case 44:
		{ //248
			yyVAL.node = &ShortVarDecl{yyS[yypt-1].token.p(), yyS[yypt-2].list, yyS[yypt-0].list}
		}
	case 45:
		{ //252
			yyVAL.node = &IncDecStmt{yyS[yypt-0].token.p(), yyS[yypt-1].node, yyS[yypt-0].token.tok}
		}
	case 46:
		{ //256
			yyVAL.node = &IncDecStmt{yyS[yypt-0].token.p(), yyS[yypt-1].node, yyS[yypt-0].token.tok}
		}
	case 47:
		{ //262
			yyVAL.node = &SwitchCase{yyS[yypt-2].token.p(), yyS[yypt-1].list, nil}
		}
	case 48:
		{ //266
			yyVAL.node = &SwitchCase{yyS[yypt-4].token.p(), []Node{&Assignment{yyS[yypt-2].token.p(), yyS[yypt-2].token.tok, yyS[yypt-3].list, []Node{yyS[yypt-1].node}}}, nil}
		}
	case 49:
		{ //270
			yyVAL.node = &SwitchCase{yyS[yypt-4].token.p(), []Node{&Assignment{yyS[yypt-2].token.p(), yyS[yypt-2].token.tok, yyS[yypt-3].list, []Node{yyS[yypt-1].node}}}, nil}
		}
	case 50:
		{ //274
			yyVAL.node = &SwitchCase{pos: yyS[yypt-1].token.p()}
		}
	case 51:
		{ //280
			yyVAL.node = &CompoundStament{yyS[yypt-2].token.p(), yyS[yypt-1].list}
		}
	case 52:
		{ //290
			yyS[yypt-1].node.(*SwitchCase).Body = yyS[yypt-0].list
			yyVAL.node = yyS[yypt-1].node
		}
	case 53:
		{ //295
			yyVAL.list = nil
		}
	case 54:
		{ //299
			yyVAL.list = append(yyS[yypt-1].list, yyS[yypt-0].node)
		}
	case 55:
		{ //305
			yyVAL.list = yyS[yypt-1].list
		}
	case 56:
		{ //311
			yyVAL.node = &ForStmt{Range: &Assignment{yyS[yypt-2].token.p(), yyS[yypt-2].token.tok, yyS[yypt-3].list, []Node{yyS[yypt-0].node}}}
		}
	case 57:
		{ //315
			yyVAL.node = &ForStmt{Range: &Assignment{yyS[yypt-2].token.p(), yyS[yypt-2].token.tok, yyS[yypt-3].list, []Node{yyS[yypt-0].node}}}
		}
	case 58:
		{ //321
			yyVAL.node = &ForStmt{Init: yyS[yypt-4].node, Cond: yyS[yypt-2].node, Post: yyS[yypt-0].node}
		}
	case 59:
		{ //325
			yyVAL.node = &ForStmt{Cond: yyS[yypt-0].node}
		}
	case 60:
		yyVAL.node = yyS[yypt-0].node
	case 61:
		{ //335
			yyS[yypt-1].node.(*ForStmt).Body = yyS[yypt-0].list
			yyVAL.node = yyS[yypt-1].node
		}
	case 62:
		{ //341
			yyS[yypt-0].node.(*ForStmt).pos = yyS[yypt-1].token.p()
			yyVAL.node = yyS[yypt-0].node
		}
	case 63:
		{ //347
			yyVAL.node = &IfStmt{Cond: yyS[yypt-0].node}
		}
	case 64:
		{ //351
			yyVAL.node = &IfStmt{Init: yyS[yypt-2].node, Cond: yyS[yypt-0].node}
		}
	case 65:
		{ //357
			x := yyS[yypt-3].node.(*IfStmt)
			l := make([]*IfStmt, len(yyS[yypt-1].list))
			for i, v := range yyS[yypt-1].list {
				l[i] = v.(*IfStmt)
			}
			x.pos, x.Body, x.Elif, x.Else = yyS[yypt-4].token.p(), yyS[yypt-2].list, l, yyS[yypt-0].node.(*CompoundStament)
			yyVAL.node = x
		}
	case 66:
		{ //363
			x := yyS[yypt-1].node.(*IfStmt)
			x.pos, x.Body = yyS[yypt-2].token.p(), yyS[yypt-0].list
			yyVAL.node = x
		}
	case 67:
		{ //368
			yyVAL.list = nil
		}
	case 68:
		{ //372
			yyVAL.list = append(yyS[yypt-1].list, yyS[yypt-0].node)
		}
	case 69:
		{ //377
			yyVAL.node = (*CompoundStament)(nil)
		}
	case 70:
		{ //381
			yyVAL.node = yyS[yypt-0].node
		}
	case 71:
		{ //387
			l := make([]*SwitchCase, len(yyS[yypt-1].list))
			for i, v := range yyS[yypt-1].list {
				l[i] = v.(*SwitchCase)
			}
			x := yyS[yypt-3].node.(*IfStmt)
			yyVAL.node = &SwitchStmt{yyS[yypt-4].token.p(), x.Init, x.Cond, l}
		}
	case 72:
		{ //393
			x := &SelectStmt{pos: yyS[yypt-3].token.p()}
			for _, v := range yyS[yypt-1].list {
				l := v.(*SwitchCase).Expr
				if len(l) != 1 {
					yyErrPos(yylex, l[1], "select cases cannot be lists")
					continue
				}

				v0 := l[0]
				switch t := v0.(type) {
				case *Assignment:
					if t.Op != token.ASSIGN && t.Op != token.DEFINE {
						break
					}

					if len(t.L) > 2 || len(t.R) != 1 {
						break
					}

					if y, ok := t.R[0].(*UnOp); ok && y.Op == token.ARROW {
						x.Cases = append(x.Cases, &CommCase{v0.p(), v0})
						continue
					}
				case *BinOp:
					if t.Op != token.ARROW {
						break
					}

					x.Cases = append(x.Cases, &CommCase{v0.p(), v0})
					continue
				case *UnOp:
					if t.Op != token.ARROW {
						break
					}

					x.Cases = append(x.Cases, &CommCase{v0.p(), v0})
					continue
				}
				yyErrPos(yylex, v0, "select case must be receive, send or assign recv")
			}
			yyVAL.node = x
		}
	case 73:
		yyVAL.node = yyS[yypt-0].node
	case 74:
		{ //403
			yyVAL.node = &BinOp{yyS[yypt-1].token.p(), yyS[yypt-1].token.tok, yyS[yypt-2].node, yyS[yypt-0].node}
		}
	case 75:
		{ //407
			yyVAL.node = &BinOp{yyS[yypt-1].token.p(), yyS[yypt-1].token.tok, yyS[yypt-2].node, yyS[yypt-0].node}
		}
	case 76:
		{ //411
			yyVAL.node = &BinOp{yyS[yypt-1].token.p(), yyS[yypt-1].token.tok, yyS[yypt-2].node, yyS[yypt-0].node}
		}
	case 77:
		{ //415
			yyVAL.node = &BinOp{yyS[yypt-1].token.p(), yyS[yypt-1].token.tok, yyS[yypt-2].node, yyS[yypt-0].node}
		}
	case 78:
		{ //419
			yyVAL.node = &BinOp{yyS[yypt-1].token.p(), yyS[yypt-1].token.tok, yyS[yypt-2].node, yyS[yypt-0].node}
		}
	case 79:
		{ //423
			yyVAL.node = &BinOp{yyS[yypt-1].token.p(), yyS[yypt-1].token.tok, yyS[yypt-2].node, yyS[yypt-0].node}
		}
	case 80:
		{ //427
			yyVAL.node = &BinOp{yyS[yypt-1].token.p(), yyS[yypt-1].token.tok, yyS[yypt-2].node, yyS[yypt-0].node}
		}
	case 81:
		{ //431
			yyVAL.node = &BinOp{yyS[yypt-1].token.p(), yyS[yypt-1].token.tok, yyS[yypt-2].node, yyS[yypt-0].node}
		}
	case 82:
		{ //435
			yyVAL.node = &BinOp{yyS[yypt-1].token.p(), yyS[yypt-1].token.tok, yyS[yypt-2].node, yyS[yypt-0].node}
		}
	case 83:
		{ //439
			yyVAL.node = &BinOp{yyS[yypt-1].token.p(), yyS[yypt-1].token.tok, yyS[yypt-2].node, yyS[yypt-0].node}
		}
	case 84:
		{ //443
			yyVAL.node = &BinOp{yyS[yypt-1].token.p(), yyS[yypt-1].token.tok, yyS[yypt-2].node, yyS[yypt-0].node}
		}
	case 85:
		{ //447
			yyVAL.node = &BinOp{yyS[yypt-1].token.p(), yyS[yypt-1].token.tok, yyS[yypt-2].node, yyS[yypt-0].node}
		}
	case 86:
		{ //451
			yyVAL.node = &BinOp{yyS[yypt-1].token.p(), yyS[yypt-1].token.tok, yyS[yypt-2].node, yyS[yypt-0].node}
		}
	case 87:
		{ //455
			yyVAL.node = &BinOp{yyS[yypt-1].token.p(), yyS[yypt-1].token.tok, yyS[yypt-2].node, yyS[yypt-0].node}
		}
	case 88:
		{ //459
			yyVAL.node = &BinOp{yyS[yypt-1].token.p(), yyS[yypt-1].token.tok, yyS[yypt-2].node, yyS[yypt-0].node}
		}
	case 89:
		{ //463
			yyVAL.node = &BinOp{yyS[yypt-1].token.p(), yyS[yypt-1].token.tok, yyS[yypt-2].node, yyS[yypt-0].node}
		}
	case 90:
		{ //467
			yyVAL.node = &BinOp{yyS[yypt-1].token.p(), yyS[yypt-1].token.tok, yyS[yypt-2].node, yyS[yypt-0].node}
		}
	case 91:
		{ //471
			yyVAL.node = &BinOp{yyS[yypt-1].token.p(), yyS[yypt-1].token.tok, yyS[yypt-2].node, yyS[yypt-0].node}
		}
	case 92:
		{ //475
			yyVAL.node = &BinOp{yyS[yypt-1].token.p(), yyS[yypt-1].token.tok, yyS[yypt-2].node, yyS[yypt-0].node}
		}
	case 93:
		{ //479
			yyVAL.node = &BinOp{yyS[yypt-1].token.p(), yyS[yypt-1].token.tok, yyS[yypt-2].node, yyS[yypt-0].node}
		}
	case 94:
		yyVAL.node = yyS[yypt-0].node
	case 95:
		{ //489
			yyVAL.node = &UnOp{yyS[yypt-1].token.p(), yyS[yypt-1].token.tok, yyS[yypt-0].node}
		}
	case 96:
		{ //493
			yyVAL.node = &UnOp{yyS[yypt-1].token.p(), yyS[yypt-1].token.tok, yyS[yypt-0].node}
		}
	case 97:
		{ //497
			yyVAL.node = &UnOp{yyS[yypt-1].token.p(), yyS[yypt-1].token.tok, yyS[yypt-0].node}
		}
	case 98:
		{ //501
			yyVAL.node = &UnOp{yyS[yypt-1].token.p(), yyS[yypt-1].token.tok, yyS[yypt-0].node}
		}
	case 99:
		{ //505
			yyVAL.node = &UnOp{yyS[yypt-1].token.p(), yyS[yypt-1].token.tok, yyS[yypt-0].node}
		}
	case 100:
		{ //509
			yyVAL.node = &UnOp{yyS[yypt-1].token.p(), yyS[yypt-1].token.tok, yyS[yypt-0].node}
		}
	case 101:
		{ //513
			yyVAL.node = &UnOp{yyS[yypt-1].token.p(), yyS[yypt-1].token.tok, yyS[yypt-0].node}
		}
	case 102:
		{ //517
			yyVAL.node = &UnOp{yyS[yypt-1].token.p(), yyS[yypt-1].token.tok, yyS[yypt-0].node}
		}
	case 103:
		{ //523
			yyVAL.node = &CallOp{yyS[yypt-1].token.p(), yyS[yypt-2].node, nil, false}
		}
	case 104:
		{ //527
			yyVAL.node = &CallOp{yyS[yypt-3].token.p(), yyS[yypt-4].node, yyS[yypt-2].list, false}
		}
	case 105:
		{ //531
			yyVAL.node = &CallOp{yyS[yypt-4].token.p(), yyS[yypt-5].node, yyS[yypt-3].list, true}
		}
	case 106:
		{ //537
			yyVAL.node = &Literal{yyS[yypt-0].token.p(), yyS[yypt-0].token.tok, yyS[yypt-0].token.lit}
		}
	case 107:
		yyVAL.node = yyS[yypt-0].node
	case 108:
		{ //545
			yyVAL.node = &SelectOp{yyS[yypt-1].token.p(), yyS[yypt-2].node, yyS[yypt-0].node.(*Ident)}
		}
	case 109:
		{ //549
			yyVAL.node = &TypeAssertion{yyS[yypt-2].token.p(), yyS[yypt-4].node, yyS[yypt-1].node}
		}
	case 110:
		{ //553
			yyVAL.node = &TypeSwitch{yyS[yypt-1].token.p(), yyS[yypt-4].node}
		}
	case 111:
		{ //557
			yyVAL.node = &IndexOp{yyS[yypt-2].token.p(), yyS[yypt-3].node, yyS[yypt-1].node}
		}
	case 112:
		{ //561
			yyVAL.node = &SliceOp{yyS[yypt-4].token.p(), yyS[yypt-5].node, yyS[yypt-3].node, yyS[yypt-1].node, nil}
		}
	case 113:
		{ //565
			if yyS[yypt-3].node == nil {
				yyErrPos(yylex, yyS[yypt-4].token, "middle index required in 3-index slice")
			}
			if yyS[yypt-1].node == nil {
				yyErrPos(yylex, yyS[yypt-2].token, "final index required in 3-index slice")
			}
			yyVAL.node = &SliceOp{yyS[yypt-6].token.p(), yyS[yypt-7].node, yyS[yypt-5].node, yyS[yypt-3].node, yyS[yypt-1].node}
		}
	case 114:
		yyVAL.node = yyS[yypt-0].node
	case 115:
		{ //573
			yyVAL.node = &ConvOp{yyS[yypt-3].token.p(), yyS[yypt-4].node, yyS[yypt-2].node}
		}
	case 116:
		{ //577
			yyVAL.node = &CompLit{yyS[yypt-3].node.p(), yyS[yypt-4].node, elements(yyS[yypt-1].list)}
		}
	case 117:
		{ //581
			yyVAL.node = &CompLit{yyS[yypt-3].token.p(), yyS[yypt-4].node, elements(yyS[yypt-1].list)}
		}
	case 118:
		{ //585
			yyVAL.node = &CompLit{yyS[yypt-3].token.p(), yyS[yypt-5].node, elements(yyS[yypt-1].list)}
		}
	case 119:
		yyVAL.node = yyS[yypt-0].node
	case 120:
		{ //594
		}
	case 121:
		{ //600
			yyVAL.node = &Element{yyS[yypt-2].node.p(), yyS[yypt-2].node, yyS[yypt-0].node}
		}
	case 122:
		{ //609
			yyVAL.node = &Element{yyS[yypt-0].node.p(), nil, yyS[yypt-0].node}
		}
	case 123:
		{ //610
			yyVAL.node = &Element{yyS[yypt-3].token.p(), nil, &CompLit{yyS[yypt-3].token.p(), nil, elements(yyS[yypt-1].list)}}
		}
	case 124:
		{ //616
			yyVAL.node = &Element{yyS[yypt-0].node.p(), nil, yyS[yypt-0].node}
		}
	case 125:
		{ //620
			yyVAL.node = &Element{yyS[yypt-3].token.p(), nil, &CompLit{yyS[yypt-3].token.p(), nil, elements(yyS[yypt-1].list)}}
		}
	case 126:
		yyVAL.node = yyS[yypt-0].node
	case 127:
		{ //630
			yyVAL.node = &Paren{yyS[yypt-2].token.p(), yyS[yypt-1].node}
		}
	case 128:
		yyVAL.node = yyS[yypt-0].node
	case 129:
		yyVAL.node = yyS[yypt-0].node
	case 130:
		yyVAL.node = yyS[yypt-0].node
	case 131:
		{
			yyVAL.node = Node(yyS[yypt-0].token)
		}
	case 132:
		{
			yyVAL.node = Node(yyS[yypt-0].token)
		}
	case 133:
		yyVAL.node = yyS[yypt-0].node
	case 134:
		yyVAL.node = yyS[yypt-0].node
	case 135:
		{ //673
			yyVAL.node = (*Ident)(nil)
		}
	case 136:
		yyVAL.node = yyS[yypt-0].node
	case 137:
		{ //683
			yyVAL.node = &Ident{yyS[yypt-0].token.p(), yyS[yypt-0].token.lit}
		}
	case 138:
		yyVAL.node = yyS[yypt-0].node
	case 139:
		yyVAL.node = yyS[yypt-0].node
	case 140:
		{ //701
			yy(yylex).errPos(yyS[yypt-0].token.tpos, "final argument in variadic function missing type")
			yyVAL.param = &Param{pos: yyS[yypt-0].token.p(), Ddd: true}
		}
	case 141:
		{ //705
			yyVAL.param = &Param{pos: yyS[yypt-1].token.p(), Ddd: true, Type: yyS[yypt-0].node}
		}
	case 142:
		yyVAL.node = yyS[yypt-0].node
	case 143:
		yyVAL.node = yyS[yypt-0].node
	case 144:
		yyVAL.node = yyS[yypt-0].node
	case 145:
		yyVAL.node = yyS[yypt-0].node
	case 146:
		{
			yyVAL.node = &NamedType{yyS[yypt-0].node.p(), yyS[yypt-0].node.(*QualifiedIdent), nil, yyScope(yylex)}
		}
	case 147:
		{ //731
			yyVAL.node = &Paren{yyS[yypt-2].token.p(), yyS[yypt-1].node}
		}
	case 148:
		{ //737
			panic(".y:738")
		}
	case 149:
		{ //741
			panic(".y:742")
		}
	case 150:
		yyVAL.node = yyS[yypt-0].node
	case 151:
		{ //749
			panic(".y:750")
		}
	case 152:
		{ //755
			panic(".y:756")
		}
	case 153:
		{ //759
			panic(".y:760")
		}
	case 154:
		{ //763
			panic(".y:764")
		}
	case 155:
		yyVAL.node = yyS[yypt-0].node
	case 156:
		{ //771
			panic(".y:772")
		}
	case 157:
		yyVAL.node = yyS[yypt-0].node
	case 158:
		yyVAL.node = yyS[yypt-0].node
	case 159:
		yyVAL.node = yyS[yypt-0].node
	case 160:
		yyVAL.node = yyS[yypt-0].node
	case 161:
		yyVAL.node = yyS[yypt-0].node
	case 162:
		yyVAL.node = yyS[yypt-0].node
	case 163:
		yyVAL.node = yyS[yypt-0].node
	case 164:
		{ //790
			yyVAL.node = &NamedType{yyS[yypt-0].node.p(), yyS[yypt-0].node.(*QualifiedIdent), nil, yyScope(yylex)}
		}
	case 165:
		{ //815
			yyVAL.node = &QualifiedIdent{yyS[yypt-0].node.p(), nil, yyS[yypt-0].node.(*Ident)}
		}
	case 166:
		{ //819
			yyVAL.node = &QualifiedIdent{yyS[yypt-2].node.p(), yyS[yypt-2].node.(*Ident), yyS[yypt-0].node.(*Ident)}
		}
	case 167:
		{ //825
			switch {
			case yyS[yypt-2].node != nil:
				yyVAL.node = &ArrayType{yyS[yypt-3].token.p(), yyS[yypt-2].node, yyS[yypt-0].node}
			default:
				yyVAL.node = &SliceType{yyS[yypt-3].token.p(), yyS[yypt-0].node}
			}
		}
	case 168:
		{ //829
			yyVAL.node = &ArrayType{yyS[yypt-3].token.p(), nil, yyS[yypt-0].node}
		}
	case 169:
		{ //833
			yyVAL.node = &ChannelType{yyS[yypt-1].token.p(), BidirectionalChannel, yyS[yypt-0].node}
		}
	case 170:
		{ //837
			yyVAL.node = &ChannelType{yyS[yypt-1].token.p(), SendOnlyChannel, yyS[yypt-0].node}
		}
	case 171:
		{ //841
			yyVAL.node = &MapType{yyS[yypt-4].token.p(), yyS[yypt-2].node, yyS[yypt-0].node}
		}
	case 172:
		yyVAL.node = yyS[yypt-0].node
	case 173:
		yyVAL.node = yyS[yypt-0].node
	case 174:
		{ //855
			yyVAL.node = &PtrType{yyS[yypt-1].token.p(), yyS[yypt-0].node}
		}
	case 175:
		{ //861
			yyVAL.node = &ChannelType{yyS[yypt-2].token.p(), ReadOnlyChannel, yyS[yypt-0].node}
		}
	case 176:
		{ //867
			yyVAL.node = newStructType(yylex, yyS[yypt-4].token, yyS[yypt-2].list)
		}
	case 177:
		{ //871
			yyVAL.node = newStructType(yylex, yyS[yypt-2].token, nil)
		}
	case 178:
		{ //877
			x := newInterfaceType(yylex, yyS[yypt-2].list)
			x.pos = yyS[yypt-4].token.p()
			yyVAL.node = x
		}
	case 179:
		{ //881
			yyVAL.node = &InterfaceType{pos: yyS[yypt-2].token.p()}
		}
	case 180:
		{ //887
			x := yyS[yypt-1].node.(*FuncDecl)
			x.pos, x.Body = yyS[yypt-2].token.p(), yyS[yypt-0].list
			yyVAL.node = x
		}
	case 181:
		{ //893
			yyVAL.node = &FuncDecl{Name: yyS[yypt-4].node.(*Ident), Type: newFuncType(yylex, yyS[yypt-3].token.p(), nil, yyS[yypt-2].params, yyS[yypt-0].params)}
		}
	case 182:
		{ //897
			yyVAL.node = &FuncDecl{Name: yyS[yypt-4].node.(*Ident), Type: newFuncType(yylex, yyS[yypt-7].token.p(), yyS[yypt-6].params, yyS[yypt-2].params, yyS[yypt-0].params)}
		}
	case 183:
		{ //903
			yyVAL.node = newFuncType(yylex, yyS[yypt-4].token.p(), nil, yyS[yypt-2].params, yyS[yypt-0].params)
		}
	case 184:
		{ //908
			yyVAL.list = nil
		}
	case 185:
		{ //912
			yyVAL.list = yyS[yypt-1].list
		}
	case 186:
		{ //918
			yyVAL.params = nil
		}
	case 187:
		{ //922
			yyVAL.params = []*Param{{pos: yyS[yypt-0].node.p(), Type: yyS[yypt-0].node}}
		}
	case 188:
		{ //926
			yyVAL.params = yyS[yypt-1].params
		}
	case 189:
		yyVAL.node = yyS[yypt-0].node
	case 190:
		{ //938
			x := yyS[yypt-3].node.(*FuncType)
			yyVAL.node = &FuncLit{x.p(), x, yyS[yypt-1].list}
		}
	case 191:
		yyVAL.node = yyS[yypt-0].node
	case 194:
		{ //957
			yyVAL.list = yyS[yypt-0].list
		}
	case 195:
		{ //961
			yyVAL.list = append(yyS[yypt-2].list, yyS[yypt-0].list...)
		}
	case 196:
		{ //967
			yyVAL.list = []Node{yyS[yypt-0].node}
		}
	case 197:
		{ //971
			yyVAL.list = append(yyS[yypt-2].list, yyS[yypt-0].node)
		}
	case 198:
		{ //977
			yyVAL.list = []Node{yyS[yypt-0].node}
		}
	case 199:
		{ //981
			yyVAL.list = append(yyS[yypt-2].list, yyS[yypt-0].node)
		}
	case 200:
		{ //987
			yyVAL.list = []Node{yyS[yypt-0].node}
		}
	case 201:
		{ //991
			yyVAL.list = append(yyS[yypt-2].list, yyS[yypt-0].node)
		}
	case 202:
		{ //997
			yyVAL.list = []Node{yyS[yypt-0].node}
		}
	case 203:
		{ //1001
			yyVAL.list = append(yyS[yypt-2].list, yyS[yypt-0].node)
		}
	case 204:
		{ //1007
			yyVAL.node = newFields(yyS[yypt-2].list, false, yyS[yypt-1].node, yyS[yypt-0].node)
		}
	case 205:
		{ //1011
			q := yyS[yypt-1].node.(*QualifiedIdent)
			yyVAL.node = newFields([]Node{q.I}, true, &NamedType{q.p(), q, nil, yyScope(yylex)}, yyS[yypt-0].node)
		}
	case 206:
		{ //1015
			yyErrPos(yylex, yyS[yypt-3].token.p(), "cannot parenthesize embedded type")
			yyVAL.node = &fields{}
		}
	case 207:
		{ //1019
			q := yyS[yypt-1].node.(*QualifiedIdent)
			yyVAL.node = newFields([]Node{q.I}, true, &PtrType{yyS[yypt-2].token.p(), &NamedType{q.p(), q, nil, yyScope(yylex)}}, yyS[yypt-0].node)
		}
	case 208:
		{ //1023
			yyErrPos(yylex, yyS[yypt-4].token.p(), "cannot parenthesize embedded type")
			yyVAL.node = &fields{}
		}
	case 209:
		{ //1027
			yyErrPos(yylex, yyS[yypt-4].token.p(), "cannot parenthesize embedded type")
			yyVAL.node = &fields{}
		}
	case 210:
		{ //1033
			yyVAL.node = &QualifiedIdent{yyS[yypt-0].token.p(), nil, &Ident{yyS[yypt-0].token.p(), yyS[yypt-0].token.lit}}
		}
	case 211:
		{ //1037
			yyVAL.node = &QualifiedIdent{yyS[yypt-2].token.p(), &Ident{yyS[yypt-2].token.p(), yyS[yypt-2].token.lit}, yyS[yypt-0].node.(*Ident)}
		}
	case 212:
		yyVAL.node = yyS[yypt-0].node
	case 213:
		{ //1049
			yyVAL.node = &MethodSpec{yyS[yypt-1].node.p(), &QualifiedIdent{yyS[yypt-1].node.p(), nil, yyS[yypt-1].node.(*Ident)}, yyS[yypt-0].node.(*FuncType)}
		}
	case 214:
		{ //1053
			yyVAL.node = &MethodSpec{yyS[yypt-0].node.p(), yyS[yypt-0].node.(*QualifiedIdent), nil}
		}
	case 215:
		{ //1057
			yyVAL.node = &MethodSpec{yyS[yypt-1].node.p(), yyS[yypt-1].node.(*QualifiedIdent), nil}
			yyErrPos(yylex, yyS[yypt-1].node, "cannot parenthesize embedded type")
		}
	case 216:
		{ //1063
			yyVAL.node = newFuncType(yylex, yyS[yypt-3].token.p(), nil, yyS[yypt-2].params, yyS[yypt-0].params)
		}
	case 217:
		{ //1069
			yyVAL.param = &Param{pos: yyS[yypt-0].node.p(), Type: yyS[yypt-0].node}
		}
	case 218:
		{ //1073
			yyVAL.param = &Param{pos: yyS[yypt-1].node.p(), Name: yyS[yypt-1].node.(*Ident), Type: yyS[yypt-0].node}
		}
	case 219:
		{ //1077
			x := yyS[yypt-0].param
			x.Name, x.Ddd = yyS[yypt-1].node.(*Ident), true
			yyVAL.param = x
		}
	case 220:
		yyVAL.param = yyS[yypt-0].param
	case 221:
		{ //1087
			yyVAL.params = []*Param{yyS[yypt-0].param}
		}
	case 222:
		{ //1091
			yyVAL.params = append(yyS[yypt-2].params, yyS[yypt-0].param)
		}
	case 223:
		{ //1096
			yyVAL.params = nil
		}
	case 224:
		{
			yyVAL.params = yyS[yypt-1].params
		}
	case 225:
		{ //1105
			yyVAL.list = nil
		}
	case 226:
		{ //1109
			yyVAL.list = []Node{yyS[yypt-0].node}
		}
	case 227:
		yyVAL.list = yyS[yypt-0].list
	case 228:
		{ //1117
			yyVAL.list = []Node{yyS[yypt-0].node}
		}
	case 229:
		{ //1121
			yyVAL.list = nil
		}
	case 230:
		yyVAL.node = yyS[yypt-0].node
	case 231:
		yyVAL.node = yyS[yypt-0].node
	case 232:
		yyVAL.node = yyS[yypt-0].node
	case 233:
		yyVAL.node = yyS[yypt-0].node
	case 234:
		yyVAL.node = yyS[yypt-0].node
	case 235:
		{ //1151
			yyVAL.node = &LabeledStmt{yyS[yypt-1].token.p(), yyS[yypt-2].node.(*Ident), yyS[yypt-0].list}
		}
	case 236:
		{ //1155
			yyVAL.node = &FallthroughStmt{yyS[yypt-0].token.p()}
		}
	case 237:
		{ //1159
			yyVAL.node = &BreakStmt{yyS[yypt-1].token.p(), yyS[yypt-0].node.(*Ident)}
		}
	case 238:
		{ //1163
			yyVAL.node = &ContinueStmt{yyS[yypt-1].token.p(), yyS[yypt-0].node.(*Ident)}
		}
	case 239:
		{ //1167
			yyVAL.node = &GoStmt{yyS[yypt-1].token.p(), yyS[yypt-0].node.(*CallOp)}
		}
	case 240:
		{ //1171
			yyVAL.node = &DeferStmt{yyS[yypt-1].token.p(), yyS[yypt-0].node.(*CallOp)}
		}
	case 241:
		{ //1175
			yyVAL.node = &GotoStmt{yyS[yypt-1].token.p(), yyS[yypt-0].node.(*Ident)}
		}
	case 242:
		{ //1179
			yyVAL.node = &ReturnStmt{yyS[yypt-1].token.p(), yyS[yypt-0].list}
		}
	case 243:
		yyVAL.list = yyS[yypt-0].list
	case 244:
		{ //1189
			yyVAL.list = append(yyS[yypt-2].list, yyS[yypt-0].list...)
		}
	case 245:
		{ //1195
			yyVAL.list = []Node{yyS[yypt-0].node}
		}
	case 246:
		{ //1199
			yyVAL.list = append(yyS[yypt-2].list, yyS[yypt-0].node)
		}
	case 247:
		{ //1205
			yyVAL.list = []Node{yyS[yypt-0].node}
		}
	case 248:
		{ //1209
			yyVAL.list = append(yyS[yypt-2].list, yyS[yypt-0].node)
		}
	case 249:
		{ //1215
			yyVAL.list = []Node{yyS[yypt-0].node}
		}
	case 250:
		{ //1219
			yyVAL.list = append(yyS[yypt-2].list, yyS[yypt-0].node)
		}
	case 251:
		{ //1225
			yyVAL.list = []Node{yyS[yypt-0].node}
		}
	case 252:
		{ //1229
			yyVAL.list = append(yyS[yypt-2].list, yyS[yypt-0].node)
		}
	case 253:
		{ //1235
			yyVAL.list = []Node{yyS[yypt-0].node}
		}
	case 254:
		{
			yyVAL.list = []Node{yyS[yypt-0].node}
		}
	case 255:
		{ //1243
			yyVAL.list = append(yyS[yypt-2].list, yyS[yypt-0].node)
		}
	case 256:
		{ //1247
			yyVAL.list = append(yyS[yypt-2].list, yyS[yypt-0].node)
		}
	case 257:
		{ //1252
			yyVAL.list = nil
		}
	case 258:
		{ //1256
			yyVAL.list = yyS[yypt-1].list
		}
	case 263:
		{ //1279
			yyVAL.node = nil
		}
	case 264:
		yyVAL.node = yyS[yypt-0].node
	case 265:
		{ //1288
			yyVAL.list = nil
		}
	case 266:
		yyVAL.list = yyS[yypt-0].list
	case 267:
		{ //1297
			yyVAL.node = nil
		}
	case 268:
		yyVAL.node = yyS[yypt-0].node
	case 269:
		{ //1306
			yyVAL.node = (*Literal)(nil)
		}
	case 270:
		{
			yyVAL.node = newLiteral(yyS[yypt-0].token)
		}
	}
	goto yystack /* stack new state and value */
}
