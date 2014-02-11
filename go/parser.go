// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parser

import __yyfmt__ "fmt"

import (
	"fmt"
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
	" (",
	" *",
	" +",
	" -",
	" .",
	" =",
	" [",
	" {",
	" >",
	" <",
	" :",
	" &",
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
	" |",
	" ^",
	" /",
	" %",
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

			p.errPos(n.Pos(), fmt.Sprintf("%s redeclared, previous declaration at %s", x.Name.Lit, p.fset.Position(ex.Pos())))
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
	68, 15,
	-2, 0,
	-1, 35,
	9, 249,
	23, 249,
	73, 249,
	-2, 41,
	-1, 43,
	14, 133,
	-2, 138,
	-1, 61,
	4, 157,
	-2, 189,
	-1, 62,
	4, 158,
	-2, 159,
	-1, 93,
	21, 225,
	29, 225,
	68, 225,
	69, 225,
	-2, 0,
	-1, 98,
	21, 239,
	29, 239,
	68, 239,
	69, 239,
	-2, 114,
	-1, 102,
	21, 240,
	29, 240,
	68, 240,
	69, 240,
	-2, 114,
	-1, 159,
	2, 189,
	4, 157,
	11, 189,
	19, 189,
	-2, 149,
	-1, 160,
	4, 158,
	11, 159,
	19, 159,
	-2, 150,
	-1, 167,
	68, 225,
	69, 225,
	-2, 0,
	-1, 211,
	68, 225,
	69, 225,
	-2, 0,
	-1, 226,
	68, 225,
	69, 225,
	-2, 0,
	-1, 253,
	68, 225,
	69, 225,
	-2, 0,
	-1, 292,
	45, 210,
	68, 210,
	69, 210,
	-2, 137,
	-1, 332,
	21, 225,
	29, 225,
	68, 225,
	69, 225,
	-2, 0,
	-1, 352,
	4, 152,
	11, 152,
	19, 152,
	-2, 143,
	-1, 353,
	4, 153,
	11, 153,
	19, 153,
	-2, 144,
	-1, 354,
	4, 154,
	11, 154,
	19, 154,
	-2, 145,
	-1, 355,
	4, 155,
	11, 155,
	19, 155,
	-2, 146,
	-1, 361,
	21, 225,
	29, 225,
	68, 225,
	69, 225,
	-2, 0,
	-1, 476,
	4, 156,
	11, 156,
	19, 156,
	-2, 147,
}

const yyNprod = 271
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 2037

var yyAct = []int{

	35, 62, 400, 169, 198, 252, 268, 285, 199, 343,
	345, 80, 139, 21, 346, 371, 212, 291, 86, 392,
	263, 214, 140, 222, 295, 287, 262, 78, 195, 299,
	85, 274, 217, 259, 106, 81, 72, 317, 215, 468,
	132, 157, 337, 430, 275, 461, 36, 358, 316, 393,
	460, 137, 137, 142, 137, 156, 160, 499, 366, 462,
	135, 277, 276, 361, 412, 171, 502, 175, 83, 491,
	176, 255, 130, 490, 177, 155, 361, 410, 361, 398,
	105, 449, 197, 158, 138, 256, 131, 197, 338, 164,
	197, 333, 333, 197, 361, 360, 85, 467, 440, 334,
	334, 437, 429, 288, 289, 361, 208, 191, 228, 463,
	186, 229, 230, 231, 232, 233, 234, 235, 236, 237,
	238, 239, 240, 241, 242, 243, 244, 245, 246, 247,
	248, 106, 106, 251, 202, 132, 132, 454, 221, 416,
	331, 413, 298, 380, 156, 160, 266, 292, 260, 267,
	368, 313, 309, 167, 302, 122, 118, 119, 41, 182,
	183, 300, 258, 160, 272, 125, 254, 126, 284, 77,
	71, 273, 158, 8, 6, 197, 495, 249, 250, 6,
	197, 197, 6, 197, 484, 481, 292, 95, 95, 480,
	158, 103, 207, 479, 106, 478, 127, 203, 476, 466,
	190, 297, 197, 270, 128, 122, 106, 294, 452, 61,
	120, 121, 123, 124, 444, 125, 298, 126, 197, 301,
	436, 197, 197, 427, 426, 423, 411, 409, 396, 197,
	305, 390, 308, 389, 386, 384, 312, 321, 364, 315,
	304, 269, 3, 314, 323, 377, 127, 76, 265, 319,
	374, 6, 311, 91, 128, 318, 106, 106, 322, 76,
	292, 375, 123, 124, 159, 372, 156, 160, 137, 347,
	188, 187, 137, 353, 347, 174, 354, 325, 197, 197,
	355, 329, 307, 359, 74, 324, 339, 6, 197, 375,
	196, 328, 6, 375, 158, 196, 74, 6, 196, 6,
	336, 196, 249, 250, 106, 185, 11, 288, 289, 197,
	357, 106, 422, 367, 373, 376, 383, 405, 197, 226,
	406, 394, 73, 379, 407, 87, 76, 414, 415, 84,
	387, 94, 397, 79, 156, 160, 56, 385, 408, 156,
	160, 290, 296, 171, 403, 253, 428, 464, 57, 257,
	388, 292, 197, 159, 141, 433, 434, 395, 197, 425,
	418, 97, 158, 74, 417, 419, 6, 158, 197, 6,
	342, 159, 93, 6, 349, 424, 441, 6, 168, 98,
	102, 211, 431, 196, 197, 442, 166, 166, 196, 196,
	154, 196, 446, 181, 165, 165, 197, 223, 297, 451,
	445, 310, 180, 447, 197, 450, 438, 13, 303, 378,
	196, 405, 458, 405, 406, 281, 406, 180, 407, 457,
	407, 453, 180, 456, 382, 180, 196, 137, 180, 196,
	196, 347, 224, 472, 347, 347, 459, 196, 403, 44,
	403, 469, 15, 474, 475, 470, 465, 143, 227, 220,
	163, 144, 172, 145, 24, 218, 18, 12, 197, 10,
	477, 9, 485, 486, 156, 160, 394, 7, 4, 171,
	99, 99, 488, 483, 482, 159, 1, 487, 137, 213,
	489, 352, 405, 494, 425, 406, 196, 196, 189, 407,
	347, 206, 158, 498, 492, 493, 196, 405, 501, 500,
	406, 283, 104, 286, 407, 344, 293, 184, 210, 403,
	180, 335, 391, 14, 88, 180, 180, 196, 180, 23,
	67, 136, 53, 2, 403, 404, 196, 290, 26, 443,
	68, 381, 25, 22, 134, 133, 192, 180, 401, 296,
	60, 204, 63, 159, 209, 90, 421, 420, 159, 58,
	59, 471, 330, 180, 332, 0, 180, 180, 0, 0,
	196, 0, 0, 0, 180, 0, 196, 0, 0, 100,
	0, 0, 0, 5, 0, 0, 196, 0, 0, 43,
	0, 75, 0, 0, 0, 0, 0, 82, 82, 89,
	92, 0, 196, 0, 0, 0, 0, 0, 96, 96,
	0, 0, 96, 0, 196, 0, 0, 0, 180, 0,
	0, 0, 196, 180, 180, 0, 0, 0, 0, 404,
	0, 404, 0, 180, 0, 0, 0, 0, 0, 278,
	0, 0, 42, 0, 279, 280, 0, 282, 0, 0,
	0, 0, 0, 75, 180, 0, 0, 0, 0, 82,
	0, 0, 180, 180, 82, 0, 306, 89, 0, 0,
	0, 216, 0, 43, 0, 0, 196, 0, 0, 0,
	0, 0, 0, 159, 0, 320, 0, 0, 146, 147,
	148, 149, 150, 151, 152, 153, 0, 180, 0, 0,
	404, 200, 179, 180, 0, 0, 205, 64, 0, 0,
	0, 0, 0, 180, 0, 404, 0, 0, 0, 65,
	0, 201, 0, 0, 264, 0, 0, 0, 54, 180,
	0, 0, 101, 0, 64, 0, 0, 351, 0, 0,
	70, 180, 362, 363, 66, 6, 65, 43, 0, 180,
	0, 0, 369, 69, 0, 0, 180, 0, 180, 101,
	0, 0, 96, 96, 0, 0, 0, 70, 0, 55,
	194, 66, 6, 351, 82, 0, 0, 0, 0, 0,
	69, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 43, 200, 179, 0, 0, 0, 0, 64, 0,
	216, 0, 0, 180, 146, 153, 43, 216, 0, 0,
	65, 0, 201, 0, 0, 219, 278, 0, 122, 118,
	119, 0, 435, 101, 0, 117, 114, 180, 125, 111,
	126, 70, 439, 43, 0, 66, 6, 129, 0, 0,
	0, 0, 180, 0, 69, 112, 0, 0, 0, 116,
	0, 0, 0, 0, 0, 0, 0, 115, 0, 127,
	455, 365, 113, 110, 0, 0, 0, 128, 0, 0,
	0, 0, 0, 120, 121, 123, 124, 0, 0, 0,
	75, 0, 82, 0, 402, 179, 358, 0, 0, 82,
	64, 0, 0, 89, 0, 399, 216, 0, 122, 118,
	119, 0, 65, 0, 201, 117, 114, 0, 125, 111,
	126, 107, 43, 0, 0, 101, 0, 129, 0, 0,
	0, 109, 0, 70, 0, 112, 0, 66, 6, 116,
	0, 0, 0, 0, 0, 108, 69, 115, 0, 127,
	0, 43, 113, 110, 0, 0, 0, 128, 96, 0,
	96, 0, 0, 120, 121, 123, 124, 0, 448, 0,
	96, 225, 216, 54, 45, 47, 48, 0, 0, 64,
	226, 0, 0, 0, 46, 0, 0, 0, 0, 28,
	0, 65, 216, 52, 34, 29, 0, 0, 0, 31,
	0, 0, 27, 37, 101, 0, 30, 32, 40, 0,
	0, 0, 70, 0, 55, 0, 66, 6, 0, 0,
	0, 0, 33, 0, 39, 69, 38, 19, 17, 16,
	51, 54, 45, 47, 48, 0, 0, 64, 0, 49,
	50, 0, 46, 0, 82, 0, 216, 28, 0, 65,
	0, 52, 34, 29, 0, 0, 0, 31, 0, 0,
	27, 37, 20, 0, 30, 32, 40, 0, 0, 0,
	70, 0, 55, 0, 66, 6, 0, 0, 0, 0,
	33, 0, 39, 69, 38, 19, 17, 0, 51, 0,
	0, 0, 122, 118, 119, 0, 0, 49, 50, 117,
	114, 0, 125, 111, 126, 0, 0, 0, 0, 0,
	0, 129, 0, 0, 54, 45, 47, 48, 0, 112,
	64, 348, 0, 116, 0, 46, 0, 0, 0, 0,
	0, 115, 65, 127, 52, 0, 113, 110, 0, 0,
	0, 128, 0, 0, 0, 101, 0, 120, 121, 123,
	124, 0, 0, 70, 0, 55, 0, 66, 6, 341,
	0, 0, 0, 0, 0, 0, 69, 54, 45, 47,
	48, 51, 0, 64, 473, 0, 0, 0, 46, 0,
	49, 50, 0, 0, 0, 65, 0, 52, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 101, 0,
	0, 0, 0, 0, 0, 0, 70, 0, 55, 0,
	66, 6, 54, 161, 47, 48, 0, 0, 64, 69,
	0, 0, 0, 46, 51, 0, 0, 0, 0, 0,
	65, 0, 162, 49, 50, 0, 0, 0, 0, 0,
	0, 0, 0, 101, 0, 0, 0, 0, 0, 0,
	0, 70, 0, 55, 0, 66, 6, 54, 45, 47,
	48, 0, 0, 64, 69, 0, 340, 0, 46, 51,
	0, 0, 0, 0, 0, 65, 0, 52, 49, 50,
	0, 0, 0, 0, 0, 0, 0, 0, 101, 0,
	0, 0, 0, 0, 0, 0, 70, 0, 55, 0,
	66, 6, 0, 0, 0, 327, 0, 0, 0, 69,
	54, 45, 47, 48, 51, 0, 64, 0, 0, 0,
	0, 46, 0, 49, 50, 0, 0, 0, 65, 0,
	52, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 101, 0, 0, 0, 0, 0, 0, 0, 70,
	0, 55, 0, 66, 6, 0, 0, 0, 326, 0,
	0, 0, 69, 54, 161, 47, 48, 51, 0, 64,
	0, 0, 0, 0, 46, 0, 49, 50, 0, 0,
	0, 65, 0, 162, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 101, 0, 0, 0, 0, 0,
	0, 0, 70, 0, 55, 0, 66, 6, 54, 45,
	47, 48, 0, 0, 64, 69, 0, 0, 0, 46,
	51, 0, 0, 0, 0, 261, 65, 0, 52, 49,
	50, 170, 0, 0, 0, 0, 0, 0, 0, 101,
	0, 0, 0, 0, 0, 0, 0, 70, 0, 55,
	0, 66, 6, 54, 45, 47, 48, 0, 0, 64,
	69, 0, 0, 0, 46, 51, 0, 0, 0, 0,
	0, 65, 0, 52, 49, 50, 0, 0, 0, 0,
	0, 0, 0, 0, 101, 0, 0, 0, 0, 0,
	0, 0, 70, 0, 55, 0, 66, 6, 54, 161,
	47, 48, 0, 0, 64, 69, 0, 0, 0, 46,
	51, 0, 0, 0, 0, 0, 65, 0, 162, 49,
	50, 0, 0, 0, 0, 0, 0, 0, 0, 101,
	0, 0, 0, 0, 0, 0, 0, 70, 0, 55,
	0, 66, 6, 54, 45, 47, 48, 0, 0, 64,
	69, 0, 0, 0, 46, 51, 0, 0, 0, 0,
	0, 271, 0, 52, 49, 50, 0, 0, 0, 0,
	0, 0, 200, 179, 101, 0, 0, 193, 64, 0,
	0, 0, 70, 0, 55, 0, 66, 6, 0, 0,
	65, 0, 201, 0, 0, 69, 0, 0, 0, 0,
	51, 0, 0, 101, 0, 0, 0, 0, 0, 49,
	50, 70, 0, 0, 0, 66, 6, 0, 122, 118,
	119, 0, 0, 0, 69, 117, 114, 497, 125, 111,
	126, 0, 0, 0, 0, 0, 0, 129, 0, 0,
	0, 194, 0, 0, 0, 112, 0, 0, 0, 116,
	0, 0, 0, 0, 0, 0, 0, 115, 0, 127,
	0, 0, 113, 110, 0, 0, 0, 128, 0, 0,
	122, 118, 119, 120, 121, 123, 124, 117, 114, 496,
	125, 111, 126, 0, 0, 0, 0, 0, 0, 129,
	0, 0, 0, 0, 0, 0, 0, 112, 0, 0,
	0, 116, 0, 0, 0, 0, 0, 0, 0, 115,
	0, 127, 0, 0, 113, 110, 0, 0, 0, 128,
	0, 0, 122, 118, 119, 120, 121, 123, 124, 117,
	114, 432, 125, 111, 126, 0, 0, 0, 0, 0,
	0, 129, 200, 179, 0, 0, 0, 0, 64, 112,
	0, 0, 0, 116, 0, 0, 0, 0, 0, 0,
	65, 115, 201, 127, 0, 0, 113, 110, 0, 0,
	0, 128, 0, 101, 0, 0, 0, 120, 121, 123,
	124, 70, 0, 0, 0, 66, 6, 0, 122, 118,
	119, 0, 0, 0, 69, 117, 114, 0, 125, 111,
	126, 0, 0, 0, 0, 0, 0, 129, 0, 0,
	0, 370, 0, 0, 0, 112, 0, 0, 0, 116,
	0, 0, 0, 0, 0, 0, 0, 115, 0, 127,
	0, 0, 113, 110, 0, 0, 0, 128, 0, 0,
	122, 118, 119, 120, 121, 123, 124, 117, 114, 0,
	125, 111, 126, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 112, 0, 0,
	0, 116, 0, 0, 0, 0, 0, 0, 0, 115,
	0, 127, 0, 0, 113, 110, 0, 0, 0, 128,
	0, 0, 122, 118, 119, 120, 121, 123, 124, 117,
	114, 0, 125, 111, 126, 0, 0, 0, 0, 0,
	0, 0, 200, 179, 0, 0, 0, 0, 64, 112,
	0, 0, 0, 116, 0, 0, 0, 122, 118, 119,
	65, 115, 201, 127, 117, 114, 113, 125, 0, 126,
	0, 128, 0, 101, 0, 0, 0, 120, 121, 123,
	124, 70, 0, 0, 112, 66, 6, 0, 116, 200,
	179, 0, 0, 0, 69, 64, 115, 0, 127, 356,
	179, 113, 0, 0, 0, 64, 128, 271, 0, 201,
	0, 0, 120, 121, 123, 124, 0, 65, 0, 350,
	101, 0, 0, 0, 0, 0, 0, 0, 70, 0,
	101, 0, 66, 6, 178, 179, 0, 0, 70, 0,
	64, 69, 66, 6, 0, 0, 0, 0, 0, 0,
	0, 69, 65, 0, 173, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 101, 0, 0, 0, 0,
	0, 0, 0, 70, 0, 0, 0, 66, 6, 0,
	0, 0, 0, 0, 0, 0, 69,
}
var yyPact = []int{

	191, -1000, -1000, 203, -1000, 105, -1000, 265, -1000, 1007,
	102, 318, 101, -1000, -1000, -1000, -1000, 329, 325, 321,
	249, -1000, -1000, -1000, -1000, -1000, 358, -1000, 203, 203,
	714, 714, 203, 1429, -1000, 883, 63, 1429, 1429, 335,
	1429, -1000, -1000, -1000, 443, 1429, 1429, 1429, 1429, 1429,
	1429, 1429, 1429, 379, 1474, -1000, -1000, -1000, 446, 375,
	-1000, -1000, -1000, 376, 1384, 1980, 383, -1000, -1000, 375,
	375, -1000, -1000, 239, -1000, 226, 225, -1000, -1000, 134,
	1548, -1000, -1000, -1000, 131, 687, -1000, 126, 1888, -1000,
	370, 778, 445, 949, -1000, -1000, -1000, -1000, -1000, 443,
	-1000, 444, -1000, -1000, -1000, -33, 1763, 1429, -1000, -1000,
	1429, 1429, 1429, 1429, 1429, 1429, 1429, 1429, 1429, 1429,
	1429, 1429, 1429, 1429, 1429, 1429, 1429, 1429, 1429, 1429,
	1429, 1429, 1429, -1000, 326, 98, -1000, -1000, 62, 330,
	94, -1000, 326, 1339, 244, 1429, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 175, 1763, -1000, -1000, -1000,
	-1000, 1474, 1519, 1429, -1000, -1000, -1000, 949, -1000, -10,
	-11, 1763, -1000, 1888, -1000, -1000, -1000, -1000, 1888, 1888,
	407, 1888, 99, 138, 93, -1000, -1000, -1000, -1000, 86,
	-1000, -1000, 399, 1429, 203, -1000, -1000, -1000, -1000, -1000,
	1888, 260, 84, -1000, 392, 1429, 83, -1000, -1000, -1000,
	-1000, 949, 173, -25, -1000, -1000, 778, -1000, -1000, 1888,
	778, -1000, -1000, -1000, -1000, -1000, 949, 778, 1763, 1867,
	1902, 150, 150, 150, 150, 150, 150, 200, 200, 200,
	200, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 1815, -33,
	-33, 1763, -1000, 949, 1429, 1286, 1233, -1000, 1429, 71,
	-1000, -1000, 15, -1000, -1000, 1188, 1067, 356, 1090, 363,
	-1000, 1945, 803, 1090, 26, -1000, 1888, 1888, -1000, 172,
	-1000, 203, -14, 82, -1000, -1000, 1718, 220, 245, 241,
	-1000, -1000, 401, 75, -1000, -1000, 420, -1000, 213, 169,
	251, 168, 203, 1429, -33, -1000, 167, 1888, 165, 203,
	1429, -33, 162, 203, 10, 870, 778, -1000, -1000, -1000,
	-1000, 161, 8, 160, -5, 73, 1429, 1429, 70, -1000,
	-1000, -1000, 949, 1474, 351, 281, 159, -26, 1474, 158,
	157, -1000, 1429, 33, -30, -1000, -1000, 1697, -1000, -1000,
	1935, -1000, -1000, -1000, -1000, -1000, 1888, 154, -1000, 32,
	-1000, 949, -1000, -1000, -1000, -1000, 1888, 29, 303, 220,
	203, -1000, -1000, 148, 213, 401, 220, 213, 203, 12,
	212, -1000, 778, 142, -1000, -1000, -1000, -1000, -33, -1000,
	-1000, 69, -1000, -1000, 687, -33, -1000, -1000, -1000, 419,
	-1000, -1000, 778, -1000, -1000, -1000, -1000, -1000, -1000, 870,
	-1000, 870, -1000, 1429, 1763, 1763, -1000, 37, 36, -1000,
	-1000, -1000, 308, -1000, 133, -1000, -1000, -1000, 25, -1000,
	1090, -1000, 1143, 1090, 1090, 132, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 220, 129, -1000, 127, -1000, -1000,
	-1000, 123, -1000, 119, 203, 392, 778, 118, -1000, -1000,
	-1000, 1429, 1429, 1474, 1429, -1000, -1000, -1000, 1429, -1000,
	-1000, -1000, 1763, -1000, 4, 0, -1000, -1000, 220, 220,
	870, -1000, -1000, 110, -1000, 1645, 1593, 326, -15, 1090,
	-1000, -1000, -1000, -1000, -1000, 870, -1000, -1000, -1000, -1000,
	-3, -1000, -1000,
}
var yyPgo = []int{

	0, 14, 554, 552, 551, 550, 23, 49, 19, 549,
	35, 8, 547, 546, 25, 0, 20, 545, 542, 540,
	538, 209, 535, 534, 533, 12, 532, 36, 531, 24,
	530, 10, 528, 336, 38, 158, 455, 432, 3, 15,
	331, 22, 1, 523, 17, 439, 522, 348, 4, 521,
	13, 7, 520, 519, 569, 18, 514, 632, 513, 9,
	33, 397, 512, 11, 511, 46, 26, 508, 507, 506,
	505, 5, 503, 502, 44, 31, 501, 491, 27, 488,
	21, 32, 479, 2, 16, 476, 468, 467, 461, 459,
	29, 457, 456, 454, 37, 89, 6, 41, 28, 452,
}
var yyR1 = []int{

	0, 86, 85, 43, 43, 87, 87, 89, 89, 89,
	27, 27, 27, 68, 68, 91, 91, 91, 91, 91,
	61, 61, 61, 61, 61, 61, 61, 61, 61, 61,
	92, 78, 78, 78, 7, 7, 8, 8, 8, 56,
	55, 50, 50, 50, 50, 50, 50, 2, 2, 2,
	2, 6, 3, 60, 60, 71, 49, 49, 23, 23,
	23, 22, 24, 25, 25, 26, 13, 64, 64, 12,
	12, 53, 93, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 57, 57, 57, 57, 57, 57,
	57, 57, 57, 47, 47, 47, 46, 46, 46, 46,
	46, 46, 46, 46, 46, 46, 46, 46, 46, 46,
	96, 31, 1, 1, 4, 4, 45, 45, 16, 16,
	34, 95, 95, 35, 10, 40, 40, 54, 33, 32,
	81, 81, 36, 36, 36, 36, 36, 36, 97, 97,
	97, 97, 99, 99, 99, 99, 99, 9, 9, 5,
	20, 20, 20, 20, 20, 11, 11, 42, 42, 42,
	42, 42, 42, 42, 48, 98, 52, 52, 30, 30,
	58, 17, 17, 21, 67, 67, 83, 83, 83, 18,
	19, 19, 88, 88, 79, 79, 62, 62, 77, 77,
	76, 76, 69, 69, 51, 51, 51, 51, 51, 51,
	44, 44, 14, 29, 29, 29, 28, 80, 80, 80,
	80, 82, 82, 84, 84, 74, 74, 74, 74, 74,
	37, 37, 37, 37, 37, 37, 37, 37, 37, 37,
	37, 37, 37, 75, 75, 72, 72, 63, 63, 65,
	65, 66, 66, 70, 70, 70, 70, 59, 59, 90,
	90, 94, 94, 38, 38, 73, 73, 41, 41, 39,
	39,
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

	-1000, -85, -43, 51, -86, -54, 48, -87, 68, -88,
	-89, 41, -91, -61, -58, -37, 2, 59, -92, 58,
	35, -50, -24, -53, -93, -26, -32, 33, 20, 26,
	37, 30, 38, 53, 25, -15, -65, 34, 57, 55,
	39, -35, -57, -54, -45, 5, 15, 6, 7, 70,
	71, 61, 24, -46, 4, 45, -33, -47, -9, -5,
	-19, -21, -42, -18, 10, 22, 47, -52, -30, 56,
	43, 68, -27, 4, 45, -54, 8, 68, -78, 4,
	-63, -10, -54, -7, 4, -63, -55, 4, -56, -54,
	-17, 4, -54, 14, -40, -35, -54, -40, -47, -45,
	-54, 35, -47, -35, -73, -65, -15, 18, 42, 28,
	50, 16, 32, 49, 13, 44, 36, 12, 6, 7,
	60, 61, 5, 62, 63, 15, 17, 46, 54, 24,
	9, 23, 73, -22, -23, -41, -49, -50, -65, -25,
	-41, 19, -25, 4, 8, 10, -57, -57, -57, -57,
	-57, -57, -57, -57, 11, -16, -15, -97, -98, -21,
	-42, 5, 24, 4, -95, 19, 11, -95, 2, -38,
	27, -15, -99, 24, -21, -42, -48, -11, 4, 5,
	-33, 10, -95, -95, -68, 66, -27, 45, 45, -79,
	66, -78, -36, 9, 73, -98, -21, -42, -48, -11,
	4, 24, -7, 66, -36, 9, -77, 66, -55, -36,
	-67, 11, -84, -82, -80, -34, -54, -81, -36, 27,
	4, -74, -6, -61, -37, 2, 11, 4, -15, -15,
	-15, -15, -15, -15, -15, -15, -15, -15, -15, -15,
	-15, -15, -15, -15, -15, -15, -15, -15, -15, -65,
	-65, -15, -71, 19, 68, 9, 23, 19, 68, -60,
	-71, 66, -66, -16, -54, 4, -15, -38, -96, 66,
	-97, 22, -15, -96, -75, -74, 72, 72, -36, -36,
	-36, 8, -36, -76, 69, -51, -72, -14, 4, 5,
	-35, -44, 48, -69, 69, -29, -35, -44, 4, -90,
	68, -90, 68, 9, -65, -10, -36, 22, -90, 68,
	9, -65, -90, 68, -75, 66, 73, -94, -34, -81,
	-36, -84, -75, -84, -75, -41, 52, 52, -60, -41,
	-3, 69, -2, 21, 29, -64, -94, 27, 73, -16,
	58, 72, 14, -59, -70, -31, -1, -15, 11, 11,
	24, -36, -21, -42, -48, -11, 4, -94, 73, -59,
	69, 68, -36, -36, 66, -54, 72, -90, 68, -36,
	73, -39, 45, -14, 5, 48, -14, 4, 8, -90,
	68, -28, 4, -44, 66, -27, 66, -78, -65, 66,
	66, -62, -8, -7, -63, -65, 66, -55, 69, -54,
	-83, -20, 4, -98, -21, -42, -48, -11, -80, 66,
	69, 66, 69, 68, -15, -15, 69, -75, -66, 14,
	-12, -13, 31, 66, -94, -16, 66, 66, -38, 69,
	73, -94, 14, -96, -96, -36, 66, 69, -74, -36,
	69, -51, -39, -35, 66, -14, -39, -14, -54, 69,
	-29, -84, 66, -90, 68, -36, 4, -84, -83, -41,
	14, 9, 23, 73, 39, -6, 66, 72, 14, -31,
	-1, -4, -15, 11, -59, -59, 66, -39, 66, 66,
	66, 66, -8, -84, 66, -15, -15, -25, -38, -96,
	69, 69, -39, -39, -83, 66, 14, 14, -71, 72,
	-59, -83, 69,
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
	3, 3, 3, 70, 3, 3, 3, 63, 15, 3,
	4, 66, 5, 6, 73, 7, 8, 62, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 14, 68,
	13, 9, 12, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 10, 3, 72, 61, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 11, 60, 69, 71,
}
var yyTok2 = []int{

	2, 3, 16, 17, 18, 19, 20, 21, 22, 23,
	24, 25, 26, 27, 28, 29, 30, 31, 32, 33,
	34, 35, 36, 37, 38, 39, 40, 41, 42, 43,
	44, 45, 46, 47, 48, 49, 50, 51, 52, 53,
	54, 55, 56, 57, 58, 59, 64, 65, 67,
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
			yyVAL.node = &Package{yyS[yypt-2].token.pos, yyS[yypt-1].node.(*Ident)}
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
			yyVAL.node = newImport(yylex, &Ident{yyS[yypt-1].token.pos, "."}, newLiteral(yyS[yypt-0].token))
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
			panic(".y:215")
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
			yyVAL.node = &TypeDecl{pos(yyS[yypt-1].node.Pos()), yyS[yypt-1].node.(*Ident), yyS[yypt-0].node}
		}
	case 41:
		yyVAL.node = yyS[yypt-0].node
	case 42:
		{ //240
			yyVAL.node = &Assignment{yyS[yypt-1].token.pos, yyS[yypt-1].token.tok, []Node{yyS[yypt-2].node}, []Node{yyS[yypt-0].node}}
		}
	case 43:
		{ //244
			yyVAL.node = &Assignment{yyS[yypt-1].token.pos, yyS[yypt-1].token.tok, yyS[yypt-2].list, yyS[yypt-0].list}
		}
	case 44:
		{ //248
			yyVAL.node = &ShortVarDecl{yyS[yypt-1].token.pos, yyS[yypt-2].list, yyS[yypt-0].list}
		}
	case 45:
		{ //252
			yyVAL.node = &IncDecStmt{yyS[yypt-0].token.pos, yyS[yypt-1].node, yyS[yypt-0].token.tok}
		}
	case 46:
		{ //256
			yyVAL.node = &IncDecStmt{yyS[yypt-0].token.pos, yyS[yypt-1].node, yyS[yypt-0].token.tok}
		}
	case 47:
		{ //262
			yyVAL.node = &SwitchCase{yyS[yypt-2].token.pos, yyS[yypt-1].list, nil}
		}
	case 48:
		{ //266
			panic(".y:267")
		}
	case 49:
		{ //270
			panic(".y:271")
		}
	case 50:
		{ //274
			yyVAL.node = &SwitchCase{pos: yyS[yypt-1].token.pos}
		}
	case 51:
		{ //280
			yyVAL.node = &CompoundStament{yyS[yypt-2].token.pos, yyS[yypt-1].list}
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
			yyVAL.node = &ForStmt{Range: &Assignment{yyS[yypt-2].token.pos, token.ASSIGN, yyS[yypt-3].list, []Node{yyS[yypt-0].node}}}
		}
	case 57:
		{ //315
			yyVAL.node = &ForStmt{Range: &Assignment{yyS[yypt-2].token.pos, token.DEFINE, yyS[yypt-3].list, []Node{yyS[yypt-0].node}}}
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
			yyS[yypt-0].node.(*ForStmt).pos = yyS[yypt-1].token.pos
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
			x.pos, x.Body, x.Elif, x.Else = yyS[yypt-4].token.pos, yyS[yypt-2].list, l, yyS[yypt-0].node.(*CompoundStament)
			yyVAL.node = x
		}
	case 66:
		{ //363
			x := yyS[yypt-1].node.(*IfStmt)
			x.pos, x.Body = yyS[yypt-2].token.pos, yyS[yypt-0].list
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
			yyVAL.node = &SwitchStmt{yyS[yypt-4].token.pos, x.Init, x.Cond, l}
		}
	case 72:
		{ //393
			panic(".y:394")
		}
	case 73:
		yyVAL.node = yyS[yypt-0].node
	case 74:
		{ //403
			panic(".y:404")
		}
	case 75:
		{ //407
			panic(".y:408")
		}
	case 76:
		{ //411
			yyVAL.node = &BinOp{yyS[yypt-1].token.pos, token.EQL, yyS[yypt-2].node, yyS[yypt-0].node}
		}
	case 77:
		{ //415
			yyVAL.node = &BinOp{yyS[yypt-1].token.pos, token.NEQ, yyS[yypt-2].node, yyS[yypt-0].node}
		}
	case 78:
		{ //419
			yyVAL.node = &BinOp{yyS[yypt-1].token.pos, token.LSS, yyS[yypt-2].node, yyS[yypt-0].node}
		}
	case 79:
		{ //423
			yyVAL.node = &BinOp{yyS[yypt-1].token.pos, token.LEQ, yyS[yypt-2].node, yyS[yypt-0].node}
		}
	case 80:
		{ //427
			panic(".y:428")
		}
	case 81:
		{ //431
			yyVAL.node = &BinOp{yyS[yypt-1].token.pos, token.GTR, yyS[yypt-2].node, yyS[yypt-0].node}
		}
	case 82:
		{ //435
			yyVAL.node = &BinOp{yyS[yypt-1].token.pos, token.ADD, yyS[yypt-2].node, yyS[yypt-0].node}
		}
	case 83:
		{ //439
			yyVAL.node = &BinOp{yyS[yypt-1].token.pos, token.SUB, yyS[yypt-2].node, yyS[yypt-0].node}
		}
	case 84:
		{ //443
			panic(".y:444")
		}
	case 85:
		{ //447
			panic(".y:448")
		}
	case 86:
		{ //451
			yyVAL.node = &BinOp{yyS[yypt-1].token.pos, token.MUL, yyS[yypt-2].node, yyS[yypt-0].node}
		}
	case 87:
		{ //455
			panic(".y:456")
		}
	case 88:
		{ //459
			panic(".y:460")
		}
	case 89:
		{ //463
			panic(".y:464")
		}
	case 90:
		{ //467
			panic(".y:468")
		}
	case 91:
		{ //471
			yyVAL.node = &BinOp{yyS[yypt-1].token.pos, token.SHL, yyS[yypt-2].node, yyS[yypt-0].node}
		}
	case 92:
		{ //475
			panic(".y:476")
		}
	case 93:
		{ //479
			panic(".y:480")
		}
	case 94:
		yyVAL.node = yyS[yypt-0].node
	case 95:
		{ //489
			yyVAL.node = &UnOp{yyS[yypt-1].token.pos, token.MUL, yyS[yypt-0].node}
		}
	case 96:
		{ //493
			yyVAL.node = &UnOp{yyS[yypt-1].token.pos, token.AND, yyS[yypt-0].node}
		}
	case 97:
		{ //497
			panic(".y:498")
		}
	case 98:
		{ //501
			yyVAL.node = &UnOp{yyS[yypt-1].token.pos, token.SUB, yyS[yypt-0].node}
		}
	case 99:
		{ //505
			panic(".y:506")
		}
	case 100:
		{ //509
			panic(".y:510")
		}
	case 101:
		{ //513
			panic(".y:514")
		}
	case 102:
		{ //517
			panic(".y:518")
		}
	case 103:
		{ //523
			yyVAL.node = &CallOp{yyS[yypt-1].token.pos, yyS[yypt-2].node, nil}
		}
	case 104:
		{ //527
			yyVAL.node = &CallOp{yyS[yypt-3].token.pos, yyS[yypt-4].node, yyS[yypt-2].list}
		}
	case 105:
		{ //531
			panic(".y:532")
		}
	case 106:
		{ //537
			yyVAL.node = &Literal{yyS[yypt-0].token.pos, yyS[yypt-0].token.tok, yyS[yypt-0].token.lit}
		}
	case 107:
		yyVAL.node = yyS[yypt-0].node
	case 108:
		{ //545
			yyVAL.node = &SelectOp{yyS[yypt-1].token.pos, yyS[yypt-2].node, yyS[yypt-0].node.(*Ident)}
		}
	case 109:
		{ //549
			panic(".y:550")
		}
	case 110:
		{ //553
			panic(".y:554")
		}
	case 111:
		{ //557
			yyVAL.node = &IndexOp{yyS[yypt-2].token.pos, yyS[yypt-3].node, yyS[yypt-1].node}
		}
	case 112:
		{ //561
			panic(".y:562")
		}
	case 113:
		{ //565
			panic(".y:566")
		}
	case 114:
		yyVAL.node = yyS[yypt-0].node
	case 115:
		{ //573
			yyVAL.node = &ConvOp{yyS[yypt-3].token.pos, yyS[yypt-4].node, yyS[yypt-2].node}
		}
	case 116:
		{ //577
			yyVAL.node = &CompLit{pos(yyS[yypt-4].node.Pos()), yyS[yypt-4].node, elements(yyS[yypt-1].list)}
		}
	case 117:
		{ //581
			yyVAL.node = &CompLit{pos(yyS[yypt-4].node.Pos()), yyS[yypt-4].node, elements(yyS[yypt-1].list)}
		}
	case 118:
		{ //585
			panic(".y:586")
		}
	case 119:
		yyVAL.node = yyS[yypt-0].node
	case 120:
		{ //594
		}
	case 121:
		{ //600
			yyVAL.node = &Element{pos(yyS[yypt-2].node.Pos()), yyS[yypt-2].node, yyS[yypt-0].node}
		}
	case 122:
		{ //609
			yyVAL.node = &Element{pos(yyS[yypt-0].node.Pos()), nil, yyS[yypt-0].node}
		}
	case 123:
		{ //610
			yyVAL.node = &Element{yyS[yypt-3].token.pos, nil, &CompLit{yyS[yypt-3].token.pos, nil, elements(yyS[yypt-1].list)}}
		}
	case 124:
		{ //616
			yyVAL.node = &Element{pos(yyS[yypt-0].node.Pos()), nil, yyS[yypt-0].node}
		}
	case 125:
		{ //620
			yyVAL.node = &Element{yyS[yypt-3].token.pos, nil, &CompLit{yyS[yypt-3].token.pos, nil, elements(yyS[yypt-1].list)}}
		}
	case 126:
		yyVAL.node = yyS[yypt-0].node
	case 127:
		{ //630
			yyVAL.node = &Paren{yyS[yypt-2].token.pos, yyS[yypt-1].node}
		}
	case 128:
		yyVAL.node = yyS[yypt-0].node
	case 129:
		{ //640
			panic(".y:641")
		}
	case 130:
		yyVAL.node = yyS[yypt-0].node
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
			yyVAL.node = &Ident{yyS[yypt-0].token.pos, yyS[yypt-0].token.lit}
		}
	case 138:
		yyVAL.node = yyS[yypt-0].node
	case 139:
		yyVAL.node = yyS[yypt-0].node
	case 140:
		{ //701
			yy(yylex).errPos(yyS[yypt-0].token.tpos, "final argument in variadic function missing type")
			yyVAL.param = &Param{pos: yyS[yypt-0].token.pos, Ddd: true}
		}
	case 141:
		{ //705
			yyVAL.param = &Param{pos: yyS[yypt-1].token.pos, Ddd: true, Type: yyS[yypt-0].node}
		}
	case 142:
		{ //711
			panic(".y:712")
		}
	case 143:
		yyVAL.node = yyS[yypt-0].node
	case 144:
		yyVAL.node = yyS[yypt-0].node
	case 145:
		yyVAL.node = yyS[yypt-0].node
	case 146:
		{
			yyVAL.node = &NamedType{pos(yyS[yypt-0].node.Pos()), yyS[yypt-0].node.(*QualifiedIdent), nil, yyScope(yylex)}
		}
	case 147:
		{ //731
			yyVAL.node = &Paren{yyS[yypt-2].token.pos, yyS[yypt-1].node}
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
		{ //745
			panic(".y:746")
		}
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
		{ //767
			panic(".y:768")
		}
	case 156:
		{ //771
			panic(".y:772")
		}
	case 157:
		{ //777
			panic(".y:778")
		}
	case 158:
		yyVAL.node = yyS[yypt-0].node
	case 159:
		yyVAL.node = yyS[yypt-0].node
	case 160:
		{ //793
			panic(".y:794")
		}
	case 161:
		{ //797
			panic(".y:798")
		}
	case 162:
		{ //801
			panic(".y:802")
		}
	case 163:
		{ //805
			panic(".y:806")
		}
	case 164:
		{ //790
			yyVAL.node = &NamedType{pos(yyS[yypt-0].node.Pos()), yyS[yypt-0].node.(*QualifiedIdent), nil, yyScope(yylex)}
		}
	case 165:
		{ //815
			yyVAL.node = &QualifiedIdent{pos(yyS[yypt-0].node.Pos()), nil, yyS[yypt-0].node.(*Ident)}
		}
	case 166:
		{ //819
			yyVAL.node = &QualifiedIdent{pos(yyS[yypt-2].node.Pos()), yyS[yypt-2].node.(*Ident), yyS[yypt-0].node.(*Ident)}
		}
	case 167:
		{ //825
			switch {
			case yyS[yypt-2].node != nil:
				yyVAL.node = &ArrayType{yyS[yypt-3].token.pos, yyS[yypt-2].node, yyS[yypt-0].node}
			default:
				yyVAL.node = &SliceType{yyS[yypt-3].token.pos, yyS[yypt-0].node}
			}
		}
	case 168:
		{ //829
			yyVAL.node = &ArrayType{yyS[yypt-3].token.pos, nil, yyS[yypt-0].node}
		}
	case 169:
		{ //833
			panic(".y:834")
		}
	case 170:
		{ //837
			panic(".y:838")
		}
	case 171:
		{ //841
			panic(".y:842")
		}
	case 172:
		yyVAL.node = yyS[yypt-0].node
	case 173:
		yyVAL.node = yyS[yypt-0].node
	case 174:
		{ //855
			yyVAL.node = &PtrType{yyS[yypt-1].token.pos, yyS[yypt-0].node}
		}
	case 175:
		{ //861
			panic(".y:862")
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
			x.pos = yyS[yypt-4].token.pos
			yyVAL.node = x
		}
	case 179:
		{ //881
			panic(".y:882")
		}
	case 180:
		{ //887
			x := yyS[yypt-1].node.(*FuncDecl)
			x.pos, x.Body = yyS[yypt-2].token.pos, yyS[yypt-0].list
			yyVAL.node = x
		}
	case 181:
		{ //893
			yyVAL.node = &FuncDecl{Name: yyS[yypt-4].node.(*Ident), Type: newFuncType(yylex, yyS[yypt-3].token.pos, yyS[yypt-2].params, yyS[yypt-0].params)}
		}
	case 182:
		{ //897
			panic(".y:898")
		}
	case 183:
		{ //903
			yyVAL.node = newFuncType(yylex, yyS[yypt-4].token.pos, yyS[yypt-2].params, yyS[yypt-0].params)
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
			yyVAL.params = []*Param{{pos: pos(yyS[yypt-0].node.Pos()), Type: yyS[yypt-0].node}}
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
			yyVAL.node = &FuncLit{x.pos, x, yyS[yypt-1].list}
		}
	case 191:
		{ //942
			panic(".y:943")
		}
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
			yyVAL.node = newFields([]Node{q.I}, true, &NamedType{q.pos, q, nil, yyScope(yylex)}, yyS[yypt-0].node)
		}
	case 206:
		{ //1015
			panic(".y:1016")
		}
	case 207:
		{ //1019
			q := yyS[yypt-1].node.(*QualifiedIdent)
			yyVAL.node = newFields([]Node{q.I}, true, &PtrType{yyS[yypt-2].token.pos, &NamedType{q.pos, q, nil, yyScope(yylex)}}, yyS[yypt-0].node)
		}
	case 208:
		{ //1023
			panic(".y:1024")
		}
	case 209:
		{ //1027
			panic(".y:1028")
		}
	case 210:
		{ //1033
			yyVAL.node = &QualifiedIdent{yyS[yypt-0].token.pos, nil, &Ident{yyS[yypt-0].token.pos, yyS[yypt-0].token.lit}}
		}
	case 211:
		{ //1037
			yyVAL.node = &QualifiedIdent{yyS[yypt-2].token.pos, &Ident{yyS[yypt-2].token.pos, yyS[yypt-2].token.lit}, yyS[yypt-0].node.(*Ident)}
		}
	case 212:
		yyVAL.node = yyS[yypt-0].node
	case 213:
		{ //1049
			yyVAL.node = &MethodSpec{pos(yyS[yypt-1].node.Pos()), yyS[yypt-1].node.(*Ident), yyS[yypt-0].node.(*FuncType)}
		}
	case 214:
		{ //1053
			panic(".y:1054")
		}
	case 215:
		{ //1057
			panic(".y:1058")
			//yyErrPos(yylex, $2, "cannot parenthesize embedded type");
		}
	case 216:
		{ //1063
			yyVAL.node = newFuncType(yylex, yyS[yypt-3].token.pos, yyS[yypt-2].params, yyS[yypt-0].params)
		}
	case 217:
		{ //1069
			yyVAL.param = &Param{pos: pos(yyS[yypt-0].node.Pos()), Type: yyS[yypt-0].node}
		}
	case 218:
		{ //1073
			yyVAL.param = &Param{pos: pos(yyS[yypt-1].node.Pos()), Name: yyS[yypt-1].node.(*Ident), Type: yyS[yypt-0].node}
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
		{ //1139
			panic(".y:1140")
		}
	case 234:
		yyVAL.node = yyS[yypt-0].node
	case 235:
		{ //1151
			yyVAL.node = &LabeledStmt{yyS[yypt-1].token.pos, yyS[yypt-2].node.(*Ident), yyS[yypt-0].list}
		}
	case 236:
		{ //1155
			yyVAL.node = &FallthroughStmt{yyS[yypt-0].token.pos}
		}
	case 237:
		{ //1159
			yyVAL.node = &BreakStmt{yyS[yypt-1].token.pos, yyS[yypt-0].node.(*Ident)}
		}
	case 238:
		{ //1163
			yyVAL.node = &ContinueStmt{yyS[yypt-1].token.pos, yyS[yypt-0].node.(*Ident)}
		}
	case 239:
		{ //1167
			panic(".y:1168")
		}
	case 240:
		{ //1171
			panic(".y:1172")
		}
	case 241:
		{ //1175
			yyVAL.node = &GotoStmt{yyS[yypt-1].token.pos, yyS[yypt-0].node.(*Ident)}
		}
	case 242:
		{ //1179
			yyVAL.node = &ReturnStmt{yyS[yypt-1].token.pos, yyS[yypt-0].list}
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
		{ //1310
			panic(".y:1311")
		}
	}
	goto yystack /* stack new state and value */
}
