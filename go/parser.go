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
	" &",
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
	if d, ok := n.(Declaration); ok {
		p.pkgScope.declare(p, d.DeclName(), n)
	}
}

func yyTLDs(y yyLexer, l []Node) {
	p := yy(y)
	for _, v := range l {
		if d, ok := v.(Declaration); ok {
			p.pkgScope.declare(p, d.DeclName(), v)
		}
	}
	p.ast = append(p.ast, l...)
}

var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 9,
	1, 2,
	67, 15,
	-2, 0,
	-1, 35,
	9, 250,
	21, 250,
	73, 250,
	-2, 41,
	-1, 43,
	68, 133,
	-2, 138,
	-1, 61,
	4, 157,
	-2, 189,
	-1, 62,
	4, 158,
	-2, 159,
	-1, 98,
	19, 240,
	27, 240,
	67, 240,
	69, 240,
	-2, 114,
	-1, 102,
	19, 241,
	27, 241,
	67, 241,
	69, 241,
	-2, 114,
	-1, 159,
	2, 189,
	4, 157,
	11, 189,
	17, 189,
	-2, 149,
	-1, 160,
	4, 158,
	11, 159,
	17, 159,
	-2, 150,
	-1, 167,
	67, 225,
	69, 225,
	-2, 0,
	-1, 211,
	67, 225,
	69, 225,
	-2, 0,
	-1, 221,
	19, 225,
	27, 225,
	67, 225,
	69, 225,
	-2, 0,
	-1, 248,
	67, 225,
	69, 225,
	-2, 0,
	-1, 275,
	67, 225,
	69, 225,
	-2, 0,
	-1, 292,
	43, 210,
	67, 210,
	69, 210,
	-2, 137,
	-1, 332,
	19, 225,
	27, 225,
	67, 225,
	69, 225,
	-2, 0,
	-1, 352,
	4, 152,
	11, 152,
	17, 152,
	-2, 143,
	-1, 353,
	4, 153,
	11, 153,
	17, 153,
	-2, 144,
	-1, 354,
	4, 154,
	11, 154,
	17, 154,
	-2, 145,
	-1, 355,
	4, 155,
	11, 155,
	17, 155,
	-2, 146,
	-1, 361,
	19, 225,
	27, 225,
	67, 225,
	69, 225,
	-2, 0,
	-1, 477,
	4, 156,
	11, 156,
	17, 156,
	-2, 147,
}

const yyNprod = 272
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1862

var yyAct = []int{

	35, 62, 139, 401, 198, 247, 393, 199, 343, 345,
	346, 295, 317, 140, 195, 271, 80, 263, 86, 285,
	169, 257, 21, 72, 270, 372, 291, 462, 214, 254,
	78, 81, 217, 157, 106, 85, 215, 258, 269, 463,
	250, 132, 430, 142, 469, 130, 337, 358, 468, 394,
	316, 135, 251, 299, 500, 156, 160, 131, 367, 277,
	137, 137, 276, 137, 361, 171, 439, 175, 83, 158,
	176, 287, 164, 177, 361, 361, 412, 399, 361, 503,
	360, 492, 197, 491, 450, 441, 461, 197, 437, 212,
	197, 464, 155, 197, 338, 288, 289, 186, 429, 333,
	419, 85, 342, 93, 132, 361, 208, 334, 223, 132,
	191, 224, 225, 226, 227, 228, 229, 230, 231, 232,
	233, 234, 235, 236, 237, 238, 239, 240, 241, 242,
	243, 106, 106, 246, 202, 76, 167, 292, 42, 455,
	413, 381, 182, 183, 156, 160, 261, 333, 255, 416,
	369, 298, 313, 309, 302, 334, 300, 122, 158, 253,
	284, 249, 77, 160, 267, 71, 262, 126, 8, 496,
	74, 485, 482, 6, 481, 197, 158, 480, 479, 6,
	197, 197, 268, 197, 146, 147, 148, 149, 150, 151,
	152, 153, 185, 292, 106, 265, 127, 331, 207, 477,
	61, 6, 197, 467, 128, 453, 106, 445, 6, 436,
	297, 427, 123, 124, 125, 41, 294, 373, 197, 426,
	203, 197, 197, 423, 197, 411, 305, 190, 410, 397,
	391, 390, 387, 385, 365, 315, 264, 3, 6, 376,
	188, 288, 289, 301, 95, 95, 322, 187, 103, 319,
	314, 106, 106, 318, 73, 159, 308, 298, 76, 76,
	312, 156, 160, 325, 347, 378, 174, 329, 353, 347,
	336, 354, 137, 260, 355, 158, 137, 359, 197, 197,
	357, 196, 328, 292, 11, 422, 196, 324, 197, 196,
	375, 307, 196, 74, 74, 275, 6, 6, 339, 292,
	146, 153, 56, 248, 106, 252, 141, 376, 94, 197,
	321, 106, 323, 166, 362, 6, 349, 406, 197, 165,
	407, 465, 91, 408, 386, 384, 395, 414, 415, 181,
	404, 376, 398, 388, 156, 160, 87, 368, 97, 156,
	160, 57, 84, 171, 159, 409, 211, 380, 158, 272,
	424, 154, 197, 158, 310, 418, 303, 431, 197, 13,
	374, 377, 159, 428, 6, 79, 433, 434, 180, 197,
	379, 417, 98, 102, 196, 281, 425, 273, 6, 196,
	196, 457, 196, 180, 6, 197, 438, 15, 180, 442,
	383, 180, 222, 451, 180, 168, 443, 197, 290, 296,
	220, 196, 163, 447, 166, 197, 221, 6, 297, 172,
	165, 26, 406, 406, 459, 407, 407, 196, 408, 408,
	196, 196, 44, 196, 143, 404, 404, 460, 144, 58,
	145, 347, 24, 473, 347, 347, 137, 18, 466, 12,
	470, 471, 475, 476, 10, 9, 454, 446, 7, 4,
	448, 1, 213, 99, 99, 189, 206, 283, 104, 197,
	286, 159, 483, 486, 487, 156, 160, 352, 488, 344,
	171, 478, 395, 452, 293, 184, 180, 196, 196, 158,
	210, 180, 180, 406, 180, 495, 407, 196, 137, 408,
	489, 347, 490, 458, 499, 335, 404, 36, 406, 501,
	502, 407, 425, 180, 408, 493, 494, 392, 196, 14,
	88, 404, 23, 67, 136, 53, 405, 196, 2, 180,
	218, 68, 180, 180, 382, 180, 25, 22, 134, 133,
	402, 105, 60, 63, 159, 138, 90, 421, 420, 159,
	59, 472, 100, 330, 332, 0, 5, 484, 0, 200,
	179, 196, 43, 0, 75, 64, 0, 196, 0, 0,
	82, 82, 89, 92, 0, 65, 0, 201, 196, 180,
	219, 96, 96, 0, 0, 96, 0, 0, 101, 180,
	180, 0, 0, 0, 196, 290, 70, 444, 0, 180,
	66, 6, 0, 0, 0, 0, 196, 296, 0, 69,
	0, 192, 0, 0, 196, 0, 204, 0, 0, 209,
	180, 405, 405, 0, 0, 0, 75, 0, 180, 180,
	0, 0, 82, 0, 0, 0, 0, 82, 244, 245,
	89, 403, 179, 0, 216, 0, 0, 64, 54, 45,
	47, 48, 0, 0, 64, 348, 0, 65, 0, 201,
	0, 0, 0, 180, 65, 0, 52, 0, 196, 180,
	101, 0, 0, 0, 0, 159, 0, 101, 70, 0,
	180, 0, 66, 6, 0, 70, 0, 55, 0, 66,
	6, 69, 405, 0, 0, 0, 180, 259, 69, 0,
	0, 304, 0, 51, 278, 0, 46, 405, 180, 279,
	280, 0, 282, 311, 49, 50, 180, 0, 0, 0,
	43, 0, 0, 180, 180, 0, 54, 45, 47, 48,
	0, 306, 64, 474, 0, 96, 96, 0, 0, 0,
	0, 0, 65, 0, 52, 0, 0, 82, 0, 0,
	320, 0, 0, 0, 0, 101, 0, 0, 244, 245,
	0, 0, 0, 70, 43, 55, 0, 66, 6, 0,
	180, 0, 0, 216, 43, 216, 69, 0, 0, 0,
	0, 51, 0, 0, 46, 0, 0, 0, 0, 0,
	0, 0, 49, 50, 180, 0, 0, 351, 200, 179,
	0, 43, 0, 205, 64, 0, 0, 363, 364, 180,
	0, 389, 0, 0, 65, 54, 201, 370, 396, 0,
	0, 64, 0, 0, 0, 0, 0, 101, 43, 0,
	0, 65, 0, 0, 366, 70, 0, 0, 351, 66,
	6, 0, 0, 0, 101, 0, 0, 0, 69, 0,
	0, 0, 70, 75, 55, 82, 66, 6, 0, 0,
	0, 0, 82, 0, 0, 69, 89, 194, 400, 216,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 278, 0, 0, 0, 43, 274, 435, 54, 45,
	47, 48, 0, 0, 64, 275, 0, 0, 440, 0,
	0, 0, 28, 0, 65, 0, 52, 34, 29, 0,
	0, 0, 31, 0, 43, 27, 37, 101, 0, 30,
	32, 40, 96, 0, 96, 70, 456, 55, 0, 66,
	6, 0, 449, 0, 96, 33, 216, 39, 69, 38,
	19, 17, 0, 51, 0, 16, 46, 54, 45, 47,
	48, 0, 0, 64, 49, 50, 216, 0, 0, 0,
	0, 28, 0, 65, 0, 52, 34, 29, 0, 0,
	0, 31, 0, 0, 27, 37, 20, 0, 30, 32,
	40, 122, 118, 119, 70, 0, 55, 0, 66, 6,
	0, 126, 0, 0, 33, 0, 39, 69, 38, 19,
	17, 0, 51, 0, 0, 46, 0, 0, 82, 0,
	216, 0, 0, 49, 50, 122, 118, 119, 0, 0,
	127, 0, 117, 114, 111, 126, 0, 0, 128, 0,
	0, 0, 129, 0, 120, 121, 123, 124, 125, 0,
	112, 0, 0, 0, 116, 0, 0, 0, 0, 0,
	0, 0, 115, 0, 127, 0, 0, 113, 110, 0,
	0, 0, 128, 0, 0, 0, 0, 0, 120, 121,
	123, 124, 125, 0, 0, 0, 0, 122, 118, 119,
	0, 0, 0, 358, 117, 114, 111, 126, 0, 0,
	0, 0, 0, 0, 129, 0, 0, 0, 0, 0,
	0, 0, 112, 0, 0, 0, 116, 0, 0, 0,
	0, 0, 0, 0, 115, 0, 127, 0, 0, 113,
	110, 0, 0, 0, 128, 0, 0, 0, 0, 0,
	120, 121, 123, 124, 125, 0, 0, 0, 122, 118,
	119, 0, 0, 0, 341, 117, 114, 111, 126, 0,
	0, 0, 0, 0, 0, 129, 0, 0, 0, 0,
	0, 0, 0, 112, 0, 0, 0, 116, 0, 0,
	200, 179, 0, 0, 0, 115, 64, 127, 0, 0,
	113, 110, 0, 0, 0, 128, 65, 0, 201, 0,
	0, 120, 121, 123, 124, 125, 122, 118, 119, 101,
	0, 498, 0, 117, 114, 111, 126, 70, 0, 0,
	0, 66, 6, 129, 0, 0, 0, 0, 0, 0,
	69, 112, 0, 0, 0, 116, 0, 0, 200, 179,
	0, 0, 0, 115, 64, 127, 0, 0, 113, 110,
	0, 0, 0, 128, 266, 0, 201, 0, 0, 120,
	121, 123, 124, 125, 122, 118, 119, 101, 0, 497,
	0, 117, 114, 111, 126, 70, 0, 0, 0, 66,
	6, 129, 0, 0, 0, 0, 0, 0, 69, 112,
	0, 0, 0, 116, 54, 161, 47, 48, 0, 0,
	64, 115, 0, 127, 0, 0, 113, 110, 0, 0,
	65, 128, 162, 0, 0, 0, 0, 120, 121, 123,
	124, 125, 0, 101, 0, 0, 0, 432, 0, 0,
	0, 70, 0, 55, 0, 66, 6, 54, 45, 47,
	48, 0, 0, 64, 69, 0, 340, 0, 0, 51,
	0, 0, 46, 65, 0, 52, 0, 0, 0, 0,
	49, 50, 0, 0, 0, 0, 101, 0, 0, 0,
	0, 0, 0, 0, 70, 0, 55, 0, 66, 6,
	0, 0, 0, 327, 0, 0, 0, 69, 0, 0,
	0, 0, 51, 0, 0, 46, 54, 45, 47, 48,
	0, 0, 64, 49, 50, 0, 0, 0, 0, 0,
	0, 0, 65, 0, 52, 0, 54, 161, 47, 48,
	0, 0, 64, 200, 179, 101, 0, 0, 193, 64,
	0, 0, 65, 70, 162, 55, 0, 66, 6, 65,
	0, 201, 326, 0, 0, 101, 69, 0, 0, 0,
	0, 51, 101, 70, 46, 55, 0, 66, 6, 0,
	70, 0, 49, 50, 66, 6, 69, 0, 0, 0,
	0, 51, 0, 69, 46, 0, 0, 256, 54, 45,
	47, 48, 49, 50, 64, 54, 45, 47, 48, 0,
	0, 64, 194, 0, 65, 0, 52, 0, 0, 170,
	0, 65, 0, 52, 0, 0, 0, 101, 0, 0,
	0, 0, 0, 0, 101, 70, 0, 55, 0, 66,
	6, 0, 70, 0, 55, 0, 66, 6, 69, 0,
	0, 0, 0, 51, 0, 69, 46, 0, 0, 0,
	51, 0, 0, 46, 49, 50, 0, 54, 161, 47,
	48, 49, 50, 64, 54, 45, 47, 48, 0, 0,
	64, 0, 0, 65, 0, 162, 0, 0, 0, 0,
	266, 0, 52, 0, 0, 0, 101, 0, 0, 0,
	0, 0, 0, 101, 70, 0, 55, 0, 66, 6,
	0, 70, 0, 55, 0, 66, 6, 69, 0, 0,
	0, 0, 51, 0, 69, 46, 0, 0, 0, 51,
	0, 0, 46, 49, 50, 122, 118, 119, 0, 0,
	49, 50, 117, 114, 111, 126, 107, 0, 0, 0,
	0, 0, 129, 0, 0, 0, 109, 200, 179, 0,
	112, 0, 0, 64, 116, 0, 0, 0, 0, 0,
	108, 0, 115, 65, 127, 201, 0, 113, 110, 0,
	0, 0, 128, 0, 0, 0, 101, 0, 120, 121,
	123, 124, 125, 0, 70, 122, 118, 119, 66, 6,
	0, 0, 117, 114, 111, 126, 0, 69, 0, 0,
	0, 0, 129, 0, 0, 0, 0, 0, 0, 0,
	112, 0, 0, 0, 116, 0, 371, 356, 179, 0,
	0, 0, 115, 64, 127, 0, 0, 113, 110, 0,
	0, 0, 128, 65, 0, 350, 0, 0, 120, 121,
	123, 124, 125, 122, 118, 119, 101, 0, 0, 0,
	117, 114, 111, 126, 70, 0, 0, 0, 66, 6,
	0, 0, 0, 0, 0, 0, 0, 69, 112, 0,
	0, 0, 116, 0, 0, 178, 179, 0, 0, 0,
	115, 64, 127, 0, 0, 113, 110, 0, 0, 0,
	128, 65, 0, 173, 0, 0, 120, 121, 123, 124,
	125, 122, 118, 119, 101, 0, 0, 0, 117, 114,
	111, 126, 70, 0, 0, 0, 66, 6, 0, 0,
	0, 0, 0, 0, 0, 69, 112, 0, 0, 0,
	116, 0, 0, 0, 122, 118, 119, 0, 115, 0,
	127, 117, 114, 113, 126, 0, 0, 0, 128, 0,
	0, 0, 0, 0, 120, 121, 123, 124, 125, 112,
	0, 0, 0, 116, 0, 0, 0, 0, 0, 0,
	0, 115, 0, 127, 0, 0, 113, 0, 0, 0,
	0, 128, 0, 0, 0, 0, 0, 120, 121, 123,
	124, 125,
}
var yyPact = []int{

	188, -1000, -1000, 192, -1000, 101, -1000, 245, -1000, 933,
	98, 250, 95, -1000, -1000, -1000, -1000, 361, 338, 332,
	318, -1000, -1000, -1000, -1000, -1000, 35, -1000, 192, 192,
	801, 801, 192, 1461, -1000, 1590, 36, 1461, 1461, 289,
	1461, -1000, -1000, -1000, 420, 1461, 1461, 1461, 1461, 1461,
	1461, 1461, 1461, 340, 1523, -1000, -1000, -1000, 398, 302,
	-1000, -1000, -1000, 393, 1454, 1741, 319, -1000, -1000, 302,
	302, -1000, -1000, 127, -1000, 204, 197, -1000, -1000, 162,
	1399, -1000, -1000, -1000, 155, 784, -1000, 133, 1156, -1000,
	335, 545, 396, -1000, -1000, -1000, -1000, -1000, -1000, 420,
	-1000, 388, -1000, -1000, -1000, -32, 1650, 1461, -1000, -1000,
	1461, 1461, 1461, 1461, 1461, 1461, 1461, 1461, 1461, 1461,
	1461, 1461, 1461, 1461, 1461, 1461, 1461, 1461, 1461, 1461,
	1461, 1461, 1461, -1000, 286, 94, -1000, -1000, 31, 288,
	92, -1000, 286, 1392, 269, 1461, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 171, 1650, -1000, -1000, -1000,
	-1000, 1523, 1530, 1461, -1000, -1000, -1000, 874, -1000, -10,
	-13, 1650, -1000, 1156, -1000, -1000, -1000, -1000, 1156, 1156,
	367, 1156, 91, 147, 89, -1000, -1000, -1000, -1000, 87,
	-1000, -1000, 347, 1461, 192, -1000, -1000, -1000, -1000, -1000,
	1156, 271, 86, -1000, 345, 1461, 85, -1000, -1000, -1000,
	-1000, 874, 170, -23, -1000, -1000, 545, -1000, -1000, 1156,
	545, 874, 545, 1650, 1766, 1799, 966, 966, 966, 966,
	966, 966, 152, 152, 152, 152, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 1708, -32, -32, 1650, -1000, 874, 1461,
	1372, 1313, -1000, 1461, 128, -1000, -1000, 21, -1000, -1000,
	1270, 1062, 34, 634, 305, -1000, 1683, 1000, 634, 11,
	-1000, -1000, -1000, -1000, -1000, 874, 1156, 1156, -1000, 169,
	-1000, 192, -14, 83, -1000, -1000, 1613, 174, 285, 261,
	-1000, -1000, 362, 74, -1000, -1000, 386, -1000, 193, 168,
	251, 167, 192, 1461, -32, -1000, 166, 1156, 165, 192,
	1461, -32, 164, 192, 8, 627, 545, -1000, -1000, -1000,
	-1000, 163, -1000, 160, 7, 73, 1461, 1461, 80, -1000,
	-1000, -1000, 874, 1523, 32, 256, 158, -26, 1523, 154,
	146, -1000, 1461, 29, -31, -1000, -1000, 1239, -1000, -1000,
	1214, -1000, -1000, -1000, -1000, -1000, 1156, 144, -1000, 19,
	-1000, 874, -3, -1000, -1000, -1000, -1000, 1156, 16, 237,
	174, 192, -1000, -1000, 142, 193, 362, 174, 193, 192,
	15, 253, -1000, 545, 140, -1000, -1000, -1000, -1000, -32,
	-1000, -1000, 72, -1000, -1000, 784, -32, -1000, -1000, -1000,
	377, -1000, -1000, 545, -1000, -1000, -1000, -1000, -1000, -1000,
	627, 627, -1000, 1461, 1650, 1650, -1000, 38, 18, -1000,
	-1000, -1000, 284, -1000, 138, -1000, -1000, -1000, -24, -1000,
	634, -1000, 712, 634, 634, 134, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 174, 113, -1000, 112, -1000,
	-1000, -1000, 109, -1000, 107, 192, 345, 545, 106, -1000,
	-1000, -1000, 1461, 1461, 1523, 1461, -1000, -1000, -1000, 1461,
	-1000, -1000, -1000, 1650, -1000, 14, 12, -1000, -1000, 174,
	174, 627, -1000, -1000, 104, -1000, 1181, 1123, 286, -18,
	634, -1000, -1000, -1000, -1000, -1000, 627, -1000, -1000, -1000,
	-1000, 10, -1000, -1000,
}
var yyPgo = []int{

	0, 10, 544, 543, 541, 540, 15, 49, 6, 31,
	7, 538, 537, 71, 0, 37, 536, 533, 532, 530,
	200, 529, 528, 527, 2, 526, 23, 524, 11, 521,
	9, 302, 36, 215, 520, 377, 20, 25, 308, 13,
	1, 518, 26, 422, 515, 341, 4, 514, 22, 19,
	513, 512, 542, 18, 510, 138, 509, 8, 29, 349,
	507, 16, 495, 497, 21, 480, 475, 474, 469, 5,
	460, 458, 24, 38, 457, 456, 30, 455, 28, 32,
	452, 3, 89, 451, 449, 448, 445, 444, 53, 439,
	437, 432, 12, 429, 72, 17, 33, 411, 14, 409,
	406,
}
var yyR1 = []int{

	0, 84, 83, 41, 41, 85, 85, 87, 87, 87,
	26, 26, 26, 66, 66, 89, 89, 89, 89, 89,
	59, 59, 59, 59, 59, 59, 59, 59, 59, 59,
	90, 76, 76, 76, 7, 7, 8, 8, 8, 54,
	53, 48, 48, 48, 48, 48, 48, 2, 2, 2,
	2, 6, 3, 58, 58, 69, 47, 47, 22, 22,
	22, 21, 23, 24, 24, 25, 12, 62, 62, 11,
	11, 51, 91, 14, 14, 14, 14, 14, 14, 14,
	14, 14, 14, 14, 14, 14, 14, 14, 14, 14,
	14, 14, 14, 14, 55, 55, 55, 55, 55, 55,
	55, 55, 55, 45, 45, 45, 44, 44, 44, 44,
	44, 44, 44, 44, 44, 44, 44, 44, 44, 44,
	95, 30, 1, 1, 4, 4, 43, 43, 15, 15,
	32, 94, 94, 33, 9, 38, 38, 52, 31, 97,
	79, 79, 34, 34, 34, 34, 34, 34, 96, 96,
	96, 96, 99, 99, 99, 99, 99, 93, 93, 5,
	19, 19, 19, 19, 19, 10, 10, 40, 40, 40,
	40, 40, 40, 40, 46, 98, 50, 50, 29, 29,
	56, 16, 16, 20, 65, 65, 81, 81, 81, 17,
	18, 18, 86, 86, 77, 77, 60, 60, 75, 75,
	74, 74, 67, 67, 49, 49, 49, 49, 49, 49,
	42, 42, 13, 28, 28, 28, 27, 78, 78, 78,
	78, 80, 80, 82, 82, 72, 72, 72, 72, 72,
	35, 35, 35, 35, 35, 100, 35, 35, 35, 35,
	35, 35, 35, 35, 73, 73, 70, 70, 61, 61,
	63, 63, 64, 64, 68, 68, 68, 68, 57, 57,
	88, 88, 92, 92, 36, 36, 71, 71, 39, 39,
	37, 37,
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
	1, 1, 1, 1, 1, 0, 4, 1, 2, 2,
	2, 2, 2, 2, 1, 3, 1, 3, 1, 3,
	1, 3, 1, 3, 1, 1, 3, 3, 0, 2,
	0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
	0, 1,
}
var yyChk = []int{

	-1000, -83, -41, 49, -84, -52, 46, -85, 67, -86,
	-87, 39, -89, -59, -56, -35, 2, 57, -90, 56,
	33, -48, -23, -51, -91, -25, -97, 31, 18, 24,
	35, 28, 36, 51, 23, -14, -63, 32, 55, 53,
	37, -33, -55, -52, -43, 5, 62, 6, 7, 70,
	71, 59, 22, -44, 4, 43, -31, -45, -93, -5,
	-18, -20, -40, -17, 10, 20, 45, -50, -29, 54,
	41, 67, -26, 4, 43, -52, 8, 67, -76, 4,
	-61, -9, -52, -7, 4, -61, -53, 4, -54, -52,
	-16, 4, -52, 68, -38, -33, -52, -38, -45, -43,
	-52, 33, -45, -33, -71, -63, -14, 16, 40, 26,
	48, 14, 30, 47, 13, 42, 34, 12, 6, 7,
	58, 59, 5, 60, 61, 62, 15, 44, 52, 22,
	9, 21, 73, -21, -22, -39, -47, -48, -63, -24,
	-39, 17, -24, 4, 8, 10, -55, -55, -55, -55,
	-55, -55, -55, -55, 11, -15, -14, -96, -98, -20,
	-40, 5, 22, 4, -94, 17, 11, -94, 2, -36,
	25, -14, -99, 22, -20, -40, -46, -10, 4, 5,
	-31, 10, -94, -94, -66, 65, -26, 43, 43, -77,
	65, -76, -34, 9, 73, -98, -20, -40, -46, -10,
	4, 22, -7, 65, -34, 9, -75, 65, -53, -34,
	-65, 11, -82, -80, -78, -32, -52, -79, -34, 25,
	4, -100, 4, -14, -14, -14, -14, -14, -14, -14,
	-14, -14, -14, -14, -14, -14, -14, -14, -14, -14,
	-14, -14, -14, -14, -63, -63, -14, -69, 17, 67,
	9, 21, 17, 67, -58, -69, 65, -64, -15, -52,
	4, -14, -36, -95, 65, -96, 20, -14, -95, -73,
	-72, -6, -59, -35, 2, 11, 72, 72, -34, -34,
	-34, 8, -34, -74, 69, -49, -70, -13, 4, 5,
	-33, -42, 46, -67, 69, -28, -33, -42, 4, -88,
	67, -88, 67, 9, -63, -9, -34, 20, -88, 67,
	9, -63, -88, 67, -73, 65, 73, -92, -32, -79,
	-34, -82, -72, -82, -73, -39, 50, 50, -58, -39,
	-3, 69, -2, 19, 27, -62, -92, 25, 73, -15,
	56, 72, 68, -57, -68, -30, -1, -14, 11, 11,
	22, -34, -20, -40, -46, -10, 4, -92, 73, -57,
	69, 67, -73, -34, -34, 65, -52, 72, -88, 67,
	-34, 73, -37, 43, -13, 5, 46, -13, 4, 8,
	-88, 67, -27, 4, -42, 65, -26, 65, -76, -63,
	65, 65, -60, -8, -7, -61, -63, 65, -53, 69,
	-52, -81, -19, 4, -98, -20, -40, -46, -10, -78,
	65, 65, 69, 67, -14, -14, 69, -73, -64, 68,
	-11, -12, 29, 65, -92, -15, 65, 65, -36, 69,
	73, -92, 68, -95, -95, -34, 65, 69, -72, 69,
	-34, 69, -49, -37, -33, 65, -13, -37, -13, -52,
	69, -28, -82, 65, -88, 67, -34, 4, -82, -81,
	-39, 68, 9, 21, 73, 37, -6, 65, 72, 68,
	-30, -1, -4, -14, 11, -57, -57, 65, -37, 65,
	65, 65, 65, -8, -82, 65, -14, -14, -24, -36,
	-95, 69, 69, -37, -37, -81, 65, 68, 68, -69,
	72, -57, -81, 69,
}
var yyDef = []int{

	3, -2, 1, 0, 5, 0, 137, 192, 4, -2,
	0, 0, 0, 16, 17, 18, 19, 0, 0, 0,
	0, 230, 231, 232, 233, 234, 0, 237, 135, 135,
	0, 0, 0, 266, 30, -2, 0, 268, 268, 0,
	268, 139, 73, -2, 94, 0, 0, 0, 0, 0,
	0, 0, 0, 126, 0, 106, 107, 114, 0, 0,
	119, -2, -2, 0, 264, 0, 0, 172, 173, 0,
	0, 6, 7, 0, 10, 0, 0, 193, 20, 0,
	0, 248, 134, 23, 0, 0, 27, 0, 0, 39,
	184, 223, 0, 235, 238, 136, 133, 239, -2, 0,
	138, 0, -2, 242, 243, 267, 250, 0, 45, 46,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 62, 0, 59, 60, 269, 0, 0,
	63, 53, 0, 0, 0, 264, 95, 96, 97, 98,
	99, 100, 101, 102, 120, 0, 128, 129, 148, -2,
	-2, 0, 0, 0, 120, 131, 132, -2, 191, 0,
	0, 265, 169, 0, 152, 153, 154, 155, 0, 0,
	165, 0, 0, 0, 260, 9, 13, 11, 12, 260,
	22, 194, 31, 0, 0, 142, 143, 144, 145, 146,
	0, 0, 260, 26, 0, 0, 260, 29, 198, 40,
	180, -2, 0, 262, 221, 217, 138, 220, 130, 140,
	223, -2, 223, 42, 74, 75, 76, 77, 78, 79,
	80, 81, 82, 83, 84, 85, 86, 87, 88, 89,
	90, 91, 92, 93, 43, 44, 251, 61, -2, 268,
	0, 0, 53, 268, 0, 67, 103, 262, 252, 108,
	0, 265, 0, 258, 127, 151, 0, 262, 258, 0,
	244, 226, 227, 228, 229, -2, 0, 0, 170, 0,
	174, 0, 0, 260, 177, 200, 0, 270, 0, 0,
	246, 212, -2, 260, 179, 202, 0, 214, 0, 0,
	261, 0, 261, 0, 33, 249, 0, 0, 0, 261,
	0, 35, 0, 261, 0, 186, 263, 224, 218, 219,
	141, 0, 236, 0, 0, 0, 0, 0, 0, 64,
	54, 72, -2, 0, 0, 69, 0, 262, 263, 0,
	0, 111, 264, 0, 262, 254, 255, 122, 120, 120,
	0, 175, -2, -2, -2, -2, 0, 0, 263, 0,
	190, -2, 0, 167, 168, 156, 166, 0, 0, 261,
	270, 0, 205, 271, 0, 0, 210, 270, 0, 0,
	0, 261, 213, 223, 0, 8, 14, 21, 195, 32,
	147, 24, 260, 196, 36, 38, 34, 28, 199, 185,
	138, 183, 187, 223, 160, 161, 162, 163, 164, 222,
	186, 186, 55, 268, 56, 57, 71, 52, 0, 50,
	65, 68, 0, 104, 0, 253, 109, 110, 0, 117,
	263, 259, 0, 258, 258, 0, 115, 116, 245, 51,
	171, 176, 201, 204, 247, 270, 0, 207, 0, 211,
	178, 203, 0, 215, 0, 261, 37, 223, 0, 181,
	58, 47, 0, 0, 0, 268, 70, 105, 112, 264,
	256, 257, 121, 124, 120, 0, 0, -2, 206, 270,
	270, 186, 25, 197, 0, 188, 0, 0, 0, 0,
	258, 123, 118, 208, 209, 216, 186, 48, 49, 66,
	113, 0, 182, 125,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 70, 3, 3, 3, 61, 62, 3,
	4, 65, 5, 6, 73, 7, 8, 60, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 68, 67,
	13, 9, 12, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 10, 3, 72, 59, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 11, 58, 69, 71,
}
var yyTok2 = []int{

	2, 3, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 63, 64, 66,
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
			yyVAL.node = &TypeDecl{pos(yyS[yypt-1].node.Pos()), (*Name)(yyS[yypt-1].node.(*Ident)), yyS[yypt-0].node}
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
			panic(".y:494")
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
			panic(".y:574")
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
			panic(".y:611")
		}
	case 124:
		{ //616
			yyVAL.node = &Element{pos(yyS[yypt-0].node.Pos()), nil, yyS[yypt-0].node}
		}
	case 125:
		{ //620
			panic(".y:621")
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
		{ //695
			panic(".y:696")
		}
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
		{ //781
			panic(".y:782")
		}
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
			yyVAL.node = &FuncDecl{Name: (*Name)(yyS[yypt-4].node.(*Ident)), Type: newFuncType(yylex, yyS[yypt-3].token.pos, yyS[yypt-2].params, yyS[yypt-0].params)}
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
		{ //1147
			panic(".y:1148")
		}
	case 236:
		{ //1151
			panic(".y:1152")
		}
	case 237:
		{ //1155
			yyVAL.node = &FallthroughStmt{yyS[yypt-0].token.pos}
		}
	case 238:
		{ //1159
			yyVAL.node = &BreakStmt{yyS[yypt-1].token.pos, yyS[yypt-0].node.(*Ident)}
		}
	case 239:
		{ //1163
			yyVAL.node = &ContinueStmt{yyS[yypt-1].token.pos, yyS[yypt-0].node.(*Ident)}
		}
	case 240:
		{ //1167
			panic(".y:1168")
		}
	case 241:
		{ //1171
			panic(".y:1172")
		}
	case 242:
		{ //1175
			yyVAL.node = &GotoStmt{yyS[yypt-1].token.pos, yyS[yypt-0].node.(*Ident)}
		}
	case 243:
		{ //1179
			yyVAL.node = &ReturnStmt{yyS[yypt-1].token.pos, yyS[yypt-0].list}
		}
	case 244:
		yyVAL.list = yyS[yypt-0].list
	case 245:
		{ //1189
			yyVAL.list = append(yyS[yypt-2].list, yyS[yypt-0].list...)
		}
	case 246:
		{ //1195
			yyVAL.list = []Node{yyS[yypt-0].node}
		}
	case 247:
		{ //1199
			yyVAL.list = append(yyS[yypt-2].list, yyS[yypt-0].node)
		}
	case 248:
		{ //1205
			yyVAL.list = []Node{yyS[yypt-0].node}
		}
	case 249:
		{ //1209
			yyVAL.list = append(yyS[yypt-2].list, yyS[yypt-0].node)
		}
	case 250:
		{ //1215
			yyVAL.list = []Node{yyS[yypt-0].node}
		}
	case 251:
		{ //1219
			yyVAL.list = append(yyS[yypt-2].list, yyS[yypt-0].node)
		}
	case 252:
		{ //1225
			yyVAL.list = []Node{yyS[yypt-0].node}
		}
	case 253:
		{ //1229
			panic(".y:1230")
		}
	case 254:
		{ //1235
			yyVAL.list = []Node{yyS[yypt-0].node}
		}
	case 255:
		{
			yyVAL.list = []Node{yyS[yypt-0].node}
		}
	case 256:
		{ //1243
			yyVAL.list = append(yyS[yypt-2].list, yyS[yypt-0].node)
		}
	case 257:
		{ //1247
			yyVAL.list = append(yyS[yypt-2].list, yyS[yypt-0].node)
		}
	case 258:
		{ //1252
			yyVAL.list = nil
		}
	case 259:
		{ //1256
			yyVAL.list = yyS[yypt-1].list
		}
	case 264:
		{ //1279
			yyVAL.node = nil
		}
	case 265:
		yyVAL.node = yyS[yypt-0].node
	case 266:
		{ //1288
			yyVAL.list = nil
		}
	case 267:
		yyVAL.list = yyS[yypt-0].list
	case 268:
		{ //1297
			yyVAL.node = nil
		}
	case 269:
		yyVAL.node = yyS[yypt-0].node
	case 270:
		{ //1306
			yyVAL.node = (*Literal)(nil)
		}
	case 271:
		{ //1310
			panic(".y:1311")
		}
	}
	goto yystack /* stack new state and value */
}
