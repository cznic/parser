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
	9, 251,
	21, 251,
	73, 251,
	-2, 41,
	-1, 43,
	68, 134,
	-2, 139,
	-1, 61,
	4, 158,
	-2, 190,
	-1, 62,
	4, 159,
	-2, 160,
	-1, 98,
	19, 241,
	27, 241,
	67, 241,
	69, 241,
	-2, 115,
	-1, 102,
	19, 242,
	27, 242,
	67, 242,
	69, 242,
	-2, 115,
	-1, 159,
	2, 190,
	4, 158,
	11, 190,
	17, 190,
	-2, 150,
	-1, 160,
	4, 159,
	11, 160,
	17, 160,
	-2, 151,
	-1, 167,
	67, 226,
	69, 226,
	-2, 0,
	-1, 211,
	67, 226,
	69, 226,
	-2, 0,
	-1, 221,
	19, 226,
	27, 226,
	67, 226,
	69, 226,
	-2, 0,
	-1, 248,
	67, 226,
	69, 226,
	-2, 0,
	-1, 275,
	67, 226,
	69, 226,
	-2, 0,
	-1, 292,
	43, 211,
	67, 211,
	69, 211,
	-2, 138,
	-1, 352,
	4, 153,
	11, 153,
	17, 153,
	-2, 144,
	-1, 353,
	4, 154,
	11, 154,
	17, 154,
	-2, 145,
	-1, 354,
	4, 155,
	11, 155,
	17, 155,
	-2, 146,
	-1, 355,
	4, 156,
	11, 156,
	17, 156,
	-2, 147,
	-1, 361,
	19, 226,
	27, 226,
	67, 226,
	69, 226,
	-2, 0,
	-1, 417,
	19, 226,
	27, 226,
	67, 226,
	69, 226,
	-2, 0,
	-1, 478,
	4, 157,
	11, 157,
	17, 157,
	-2, 148,
}

const yyNprod = 273
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1918

var yyAct = []int{

	35, 62, 343, 169, 80, 247, 263, 393, 401, 21,
	345, 258, 212, 346, 195, 271, 139, 198, 295, 285,
	270, 214, 299, 85, 269, 291, 257, 254, 199, 78,
	287, 217, 215, 317, 106, 140, 72, 81, 157, 372,
	470, 132, 250, 337, 469, 504, 430, 137, 137, 394,
	137, 86, 42, 358, 251, 156, 160, 142, 316, 501,
	367, 277, 276, 493, 492, 171, 155, 175, 83, 158,
	450, 441, 130, 135, 437, 361, 463, 439, 361, 361,
	412, 399, 197, 176, 131, 429, 164, 197, 464, 85,
	197, 338, 419, 197, 177, 288, 289, 342, 146, 147,
	148, 149, 150, 151, 152, 153, 132, 93, 223, 191,
	186, 224, 225, 226, 227, 228, 229, 230, 231, 232,
	233, 234, 235, 236, 237, 238, 239, 240, 241, 242,
	243, 106, 106, 246, 202, 462, 132, 292, 361, 208,
	465, 455, 413, 333, 156, 160, 261, 333, 255, 262,
	167, 334, 381, 76, 298, 334, 182, 183, 158, 361,
	284, 360, 369, 160, 267, 313, 309, 302, 300, 253,
	249, 268, 77, 71, 8, 197, 158, 497, 41, 6,
	197, 197, 486, 197, 6, 6, 483, 482, 74, 6,
	481, 6, 480, 416, 106, 478, 292, 331, 207, 468,
	265, 61, 197, 203, 190, 453, 106, 95, 95, 297,
	185, 103, 301, 445, 146, 153, 436, 427, 197, 294,
	426, 197, 197, 423, 197, 308, 411, 410, 397, 312,
	391, 390, 305, 321, 387, 323, 314, 385, 365, 315,
	264, 3, 322, 376, 288, 289, 373, 188, 319, 318,
	298, 106, 106, 375, 378, 260, 159, 91, 187, 137,
	87, 156, 160, 137, 347, 84, 73, 174, 353, 347,
	76, 359, 339, 324, 11, 158, 422, 76, 197, 197,
	328, 79, 196, 94, 354, 325, 292, 196, 197, 329,
	196, 336, 292, 196, 376, 355, 376, 6, 307, 6,
	362, 357, 6, 275, 106, 74, 368, 6, 6, 197,
	57, 106, 74, 97, 395, 6, 380, 406, 197, 374,
	377, 248, 56, 6, 384, 166, 252, 414, 415, 466,
	404, 165, 388, 407, 156, 160, 44, 386, 409, 156,
	160, 98, 102, 171, 408, 159, 428, 141, 158, 349,
	425, 211, 197, 158, 154, 433, 434, 168, 197, 181,
	418, 290, 296, 159, 379, 398, 166, 99, 99, 197,
	143, 424, 165, 310, 144, 196, 145, 272, 431, 303,
	196, 196, 438, 196, 273, 197, 281, 13, 180, 442,
	457, 383, 222, 220, 15, 163, 452, 197, 221, 172,
	451, 26, 196, 180, 58, 197, 446, 297, 180, 448,
	443, 180, 406, 406, 180, 454, 458, 447, 196, 459,
	24, 196, 196, 137, 196, 404, 404, 23, 407, 407,
	417, 347, 330, 474, 347, 347, 476, 477, 467, 408,
	408, 471, 461, 332, 472, 18, 12, 10, 9, 460,
	7, 4, 1, 213, 189, 206, 283, 104, 286, 197,
	395, 344, 159, 484, 487, 488, 156, 160, 352, 293,
	485, 171, 184, 210, 490, 335, 137, 425, 196, 196,
	158, 392, 491, 489, 406, 479, 14, 88, 196, 67,
	136, 496, 347, 53, 502, 500, 180, 404, 2, 406,
	407, 180, 180, 36, 180, 68, 503, 382, 25, 196,
	22, 408, 404, 134, 133, 407, 402, 405, 196, 60,
	494, 495, 63, 180, 90, 421, 408, 420, 59, 473,
	0, 0, 0, 0, 0, 159, 0, 105, 0, 180,
	159, 138, 180, 180, 0, 180, 0, 0, 290, 0,
	444, 0, 196, 0, 0, 0, 0, 0, 196, 0,
	296, 0, 100, 122, 118, 119, 5, 122, 0, 196,
	0, 0, 43, 126, 75, 0, 0, 126, 0, 0,
	82, 82, 89, 92, 0, 196, 0, 0, 0, 180,
	0, 96, 96, 0, 0, 96, 0, 196, 0, 180,
	180, 0, 127, 0, 0, 196, 127, 218, 0, 180,
	128, 0, 405, 405, 128, 0, 120, 121, 123, 124,
	125, 0, 123, 124, 125, 0, 0, 200, 179, 0,
	180, 0, 0, 64, 244, 245, 75, 0, 180, 180,
	0, 0, 82, 65, 0, 201, 0, 82, 219, 0,
	89, 0, 0, 0, 216, 0, 101, 0, 0, 196,
	0, 0, 0, 0, 70, 0, 0, 159, 66, 6,
	0, 0, 0, 180, 0, 0, 0, 69, 0, 180,
	0, 0, 0, 0, 405, 0, 0, 0, 192, 0,
	180, 0, 0, 204, 0, 0, 209, 304, 0, 405,
	0, 0, 0, 0, 0, 0, 180, 259, 0, 311,
	0, 54, 161, 47, 48, 0, 0, 64, 180, 0,
	0, 0, 0, 0, 0, 0, 180, 65, 0, 162,
	43, 0, 0, 180, 180, 0, 0, 0, 0, 0,
	101, 0, 0, 0, 0, 96, 96, 0, 70, 0,
	55, 0, 66, 6, 244, 245, 0, 82, 0, 0,
	0, 69, 0, 340, 0, 0, 51, 0, 0, 46,
	0, 0, 0, 0, 43, 0, 0, 49, 50, 0,
	180, 278, 0, 216, 43, 216, 279, 280, 0, 282,
	0, 0, 122, 118, 119, 0, 0, 0, 0, 117,
	114, 111, 126, 0, 0, 180, 0, 389, 306, 129,
	0, 43, 0, 0, 396, 0, 0, 112, 0, 0,
	180, 116, 0, 0, 0, 0, 0, 320, 0, 115,
	0, 127, 0, 0, 113, 110, 0, 0, 43, 128,
	0, 0, 0, 0, 366, 120, 121, 123, 124, 125,
	0, 200, 179, 0, 0, 0, 205, 64, 0, 0,
	358, 0, 0, 75, 0, 82, 0, 65, 0, 201,
	0, 0, 82, 0, 351, 0, 89, 0, 400, 216,
	101, 0, 0, 0, 363, 364, 0, 0, 70, 0,
	0, 0, 66, 6, 370, 0, 122, 118, 119, 0,
	0, 69, 0, 117, 114, 111, 126, 0, 0, 0,
	0, 0, 0, 129, 0, 351, 0, 0, 0, 0,
	194, 112, 0, 0, 43, 116, 0, 0, 0, 0,
	0, 54, 96, 115, 96, 127, 0, 64, 113, 110,
	0, 0, 449, 128, 96, 0, 216, 65, 0, 120,
	121, 123, 124, 125, 0, 0, 0, 0, 278, 0,
	101, 0, 0, 341, 435, 0, 216, 0, 70, 0,
	55, 0, 66, 6, 0, 440, 0, 0, 0, 0,
	43, 69, 0, 0, 0, 0, 0, 0, 403, 179,
	0, 0, 0, 274, 64, 54, 45, 47, 48, 0,
	0, 64, 275, 456, 65, 0, 201, 0, 0, 28,
	0, 65, 0, 52, 34, 29, 0, 101, 82, 31,
	216, 0, 27, 37, 101, 70, 30, 32, 40, 66,
	6, 0, 70, 0, 55, 0, 66, 6, 69, 0,
	0, 0, 33, 0, 39, 69, 38, 19, 17, 0,
	51, 0, 16, 46, 54, 45, 47, 48, 0, 0,
	64, 49, 50, 0, 0, 0, 0, 0, 28, 0,
	65, 0, 52, 34, 29, 0, 0, 0, 31, 0,
	0, 27, 37, 20, 0, 30, 32, 40, 0, 0,
	0, 70, 0, 55, 0, 66, 6, 0, 0, 0,
	0, 33, 0, 39, 69, 38, 19, 17, 0, 51,
	0, 0, 46, 0, 0, 0, 54, 45, 47, 48,
	49, 50, 64, 348, 0, 0, 0, 0, 0, 0,
	0, 0, 65, 0, 52, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 101, 0, 0, 0, 0,
	0, 0, 0, 70, 0, 55, 0, 66, 6, 0,
	122, 118, 119, 0, 0, 0, 69, 117, 114, 111,
	126, 51, 0, 0, 46, 0, 0, 129, 0, 0,
	0, 0, 49, 50, 0, 112, 0, 0, 0, 116,
	0, 0, 200, 179, 0, 0, 0, 115, 64, 127,
	0, 0, 113, 110, 0, 0, 0, 128, 65, 0,
	201, 0, 0, 120, 121, 123, 124, 125, 122, 118,
	119, 101, 0, 499, 0, 117, 114, 111, 126, 70,
	0, 0, 0, 66, 6, 129, 0, 0, 0, 0,
	0, 0, 69, 112, 0, 0, 0, 116, 0, 0,
	0, 0, 0, 0, 0, 115, 0, 127, 0, 0,
	113, 110, 0, 0, 0, 128, 0, 0, 0, 0,
	0, 120, 121, 123, 124, 125, 54, 45, 47, 48,
	0, 498, 64, 475, 0, 0, 0, 0, 0, 0,
	0, 0, 65, 0, 52, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 101, 0, 0, 0, 0,
	0, 0, 0, 70, 0, 55, 0, 66, 6, 0,
	122, 118, 119, 0, 0, 0, 69, 117, 114, 111,
	126, 51, 0, 0, 46, 0, 0, 129, 0, 0,
	0, 0, 49, 50, 0, 112, 0, 0, 0, 116,
	54, 45, 47, 48, 0, 0, 64, 115, 0, 127,
	0, 0, 113, 110, 0, 0, 65, 128, 52, 0,
	0, 0, 0, 120, 121, 123, 124, 125, 0, 101,
	0, 0, 0, 432, 0, 0, 0, 70, 0, 55,
	0, 66, 6, 0, 0, 0, 327, 0, 0, 0,
	69, 0, 0, 0, 0, 51, 0, 0, 46, 54,
	45, 47, 48, 0, 0, 64, 49, 50, 0, 0,
	0, 0, 0, 0, 0, 65, 0, 52, 0, 54,
	161, 47, 48, 0, 0, 64, 200, 179, 101, 0,
	0, 193, 64, 0, 0, 65, 70, 162, 55, 0,
	66, 6, 65, 0, 201, 326, 0, 0, 101, 69,
	0, 0, 0, 0, 51, 101, 70, 46, 55, 0,
	66, 6, 0, 70, 0, 49, 50, 66, 6, 69,
	0, 0, 0, 0, 51, 0, 69, 46, 0, 0,
	256, 54, 45, 47, 48, 49, 50, 64, 54, 45,
	47, 48, 0, 0, 64, 194, 0, 65, 0, 52,
	0, 0, 170, 0, 65, 0, 52, 0, 0, 0,
	101, 0, 0, 0, 0, 0, 0, 101, 70, 0,
	55, 0, 66, 6, 0, 70, 0, 55, 0, 66,
	6, 69, 0, 0, 0, 0, 51, 0, 69, 46,
	0, 0, 0, 51, 0, 0, 46, 49, 50, 0,
	54, 161, 47, 48, 49, 50, 64, 54, 45, 47,
	48, 0, 0, 64, 0, 0, 65, 0, 162, 0,
	0, 0, 0, 266, 0, 52, 0, 0, 0, 101,
	0, 0, 0, 0, 0, 0, 101, 70, 0, 55,
	0, 66, 6, 0, 70, 0, 55, 0, 66, 6,
	69, 0, 0, 0, 0, 51, 0, 69, 46, 0,
	0, 0, 51, 0, 0, 46, 49, 50, 122, 118,
	119, 0, 0, 49, 50, 117, 114, 111, 126, 107,
	0, 0, 0, 0, 0, 129, 0, 0, 0, 109,
	200, 179, 0, 112, 0, 0, 64, 116, 0, 0,
	0, 0, 0, 108, 0, 115, 65, 127, 201, 0,
	113, 110, 0, 0, 0, 128, 0, 0, 0, 101,
	0, 120, 121, 123, 124, 125, 0, 70, 122, 118,
	119, 66, 6, 0, 0, 117, 114, 111, 126, 0,
	69, 0, 0, 0, 0, 129, 0, 0, 0, 0,
	0, 0, 0, 112, 0, 0, 0, 116, 0, 371,
	200, 179, 0, 0, 0, 115, 64, 127, 0, 0,
	113, 110, 0, 0, 0, 128, 266, 0, 201, 0,
	0, 120, 121, 123, 124, 125, 122, 118, 119, 101,
	0, 0, 0, 117, 114, 111, 126, 70, 0, 0,
	0, 66, 6, 0, 0, 0, 0, 0, 0, 0,
	69, 112, 0, 0, 0, 116, 0, 0, 356, 179,
	0, 0, 0, 115, 64, 127, 0, 0, 113, 110,
	0, 0, 0, 128, 65, 0, 350, 0, 0, 120,
	121, 123, 124, 125, 122, 118, 119, 101, 0, 0,
	0, 117, 114, 111, 126, 70, 0, 0, 0, 66,
	6, 0, 0, 0, 0, 0, 0, 0, 69, 112,
	0, 0, 0, 116, 0, 0, 0, 122, 118, 119,
	0, 115, 0, 127, 117, 114, 113, 126, 0, 0,
	0, 128, 0, 0, 0, 0, 0, 120, 121, 123,
	124, 125, 112, 0, 0, 0, 116, 178, 179, 0,
	0, 0, 0, 64, 115, 0, 127, 0, 0, 113,
	0, 0, 0, 65, 128, 173, 0, 0, 0, 0,
	120, 121, 123, 124, 125, 0, 101, 0, 0, 0,
	0, 0, 0, 0, 70, 0, 0, 0, 66, 6,
	0, 0, 0, 0, 0, 0, 0, 69,
}
var yyPact = []int{

	192, -1000, -1000, 143, -1000, 107, -1000, 235, -1000, 1050,
	106, 262, 105, -1000, -1000, -1000, -1000, 277, 261, 256,
	253, -1000, -1000, -1000, -1000, -1000, 39, -1000, 143, 143,
	927, 927, 143, 1494, -1000, 1623, 63, 1494, 1494, 330,
	1494, -1000, -1000, -1000, 366, 1494, 1494, 1494, 1494, 1494,
	1494, 1494, 1494, 343, 1556, -1000, -1000, -1000, 391, 314,
	-1000, -1000, -1000, 355, 1487, 1863, 349, -1000, -1000, 314,
	314, -1000, -1000, 145, -1000, 215, 204, -1000, -1000, 139,
	1432, -1000, -1000, -1000, 138, 847, -1000, 133, 1188, -1000,
	340, 623, 389, -1000, -1000, -1000, -1000, -1000, -1000, 366,
	-1000, 388, -1000, -1000, -1000, -32, 1683, 1494, -1000, -1000,
	1494, 1494, 1494, 1494, 1494, 1494, 1494, 1494, 1494, 1494,
	1494, 1494, 1494, 1494, 1494, 1494, 1494, 1494, 1494, 1494,
	1494, 1494, 1494, -1000, 304, 103, -1000, -1000, 33, 309,
	102, -1000, 304, 1425, 251, 1494, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 175, 1683, -1000, -1000, -1000,
	-1000, 1556, 1563, 1494, -1000, -1000, -1000, 991, -1000, -10,
	-11, 1683, -1000, 1188, -1000, -1000, -1000, -1000, 1188, 1188,
	378, 1188, 91, 150, 101, -1000, -1000, -1000, -1000, 100,
	-1000, -1000, 370, 1494, 143, -1000, -1000, -1000, -1000, -1000,
	1188, 278, 99, -1000, 364, 1494, 98, -1000, -1000, -1000,
	-1000, 991, 174, -15, -1000, -1000, 623, -1000, -1000, 1188,
	623, 991, 623, 1683, 1799, 1832, 558, 558, 558, 558,
	558, 558, 562, 562, 562, 562, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 1741, -32, -32, 1683, -1000, 991, 1494,
	1405, 1346, -1000, 1494, 128, -1000, -1000, 18, -1000, -1000,
	707, 891, 29, 1112, 338, -1000, 1774, 787, 1112, 92,
	-1000, -1000, -1000, -1000, -1000, 991, 1188, 1188, -1000, 173,
	-1000, 143, -12, 95, -1000, -1000, 1646, 203, 248, 250,
	-1000, -1000, 356, 85, -1000, -1000, 387, -1000, 197, 172,
	269, 169, 143, 1494, -32, -1000, 166, 1188, 165, 143,
	1494, -32, 163, 143, 12, 984, 623, -1000, -1000, -1000,
	-1000, 162, -1000, 161, 11, 75, 1494, 1494, 124, -1000,
	-1000, -1000, -1000, 1556, 24, 247, 158, -20, 1556, 155,
	152, -1000, 1494, 16, -27, -1000, -1000, 1315, -1000, -1000,
	1716, -1000, -1000, -1000, -1000, -1000, 1188, 151, -1000, 5,
	-1000, 991, 8, -1000, -1000, -1000, -1000, 1188, 2, 240,
	203, 143, -1000, -1000, 148, 197, 356, 203, 197, 143,
	1, 246, -1000, 623, 140, -1000, -1000, -1000, -1000, -32,
	-1000, -1000, 74, -1000, -1000, 847, -32, -1000, -1000, -1000,
	386, -1000, -1000, 623, -1000, -1000, -1000, -1000, -1000, -1000,
	984, 984, -1000, 1494, 1683, 1683, -1000, 991, 67, -1000,
	-1000, -1000, 292, -1000, 134, -1000, -1000, -1000, -28, -1000,
	1112, -1000, 1272, 1112, 1112, 130, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 203, 127, -1000, 125, -1000,
	-1000, -1000, 122, -1000, 121, 143, 364, 623, 117, -1000,
	-1000, 71, -1000, 1494, 1494, 1556, 1494, -1000, -1000, -1000,
	1494, -1000, -1000, -1000, 1683, -1000, -5, -6, -1000, -1000,
	203, 203, 984, -1000, -1000, 112, -1000, 1213, 1155, 304,
	-13, 1112, -1000, -1000, -1000, -1000, -1000, 984, -1000, -1000,
	-1000, -1000, -24, -1000, -1000,
}
var yyPgo = []int{

	0, 13, 529, 528, 15, 49, 7, 37, 28, 527,
	525, 30, 0, 11, 524, 522, 519, 516, 201, 514,
	513, 510, 16, 508, 36, 507, 18, 505, 10, 322,
	32, 178, 607, 384, 3, 39, 283, 35, 1, 498,
	25, 336, 493, 310, 17, 490, 9, 19, 489, 562,
	51, 487, 52, 486, 2, 377, 481, 4, 475, 503,
	26, 473, 472, 469, 461, 5, 458, 457, 20, 24,
	456, 455, 29, 454, 21, 31, 453, 8, 12, 452,
	451, 450, 448, 447, 22, 446, 445, 443, 432, 430,
	27, 427, 420, 33, 404, 86, 6, 38, 401, 14,
	399, 398,
}
var yyR1 = []int{

	0, 80, 79, 39, 39, 81, 81, 83, 83, 83,
	24, 24, 24, 62, 62, 85, 85, 85, 85, 85,
	55, 55, 55, 55, 55, 55, 55, 55, 55, 55,
	86, 72, 72, 72, 5, 5, 6, 6, 6, 51,
	50, 46, 46, 46, 46, 46, 46, 87, 87, 87,
	87, 4, 89, 88, 90, 90, 65, 45, 45, 20,
	20, 20, 19, 21, 22, 22, 23, 10, 58, 58,
	9, 9, 91, 92, 12, 12, 12, 12, 12, 12,
	12, 12, 12, 12, 12, 12, 12, 12, 12, 12,
	12, 12, 12, 12, 12, 52, 52, 52, 52, 52,
	52, 52, 52, 52, 43, 43, 43, 42, 42, 42,
	42, 42, 42, 42, 42, 42, 42, 42, 42, 42,
	42, 96, 28, 1, 1, 2, 2, 41, 41, 13,
	13, 30, 95, 95, 31, 7, 36, 36, 49, 29,
	98, 75, 75, 32, 32, 32, 32, 32, 32, 97,
	97, 97, 97, 100, 100, 100, 100, 100, 94, 94,
	3, 17, 17, 17, 17, 17, 8, 8, 38, 38,
	38, 38, 38, 38, 38, 44, 99, 48, 48, 27,
	27, 53, 14, 14, 18, 61, 61, 77, 77, 77,
	15, 16, 16, 82, 82, 73, 73, 56, 56, 71,
	71, 70, 70, 63, 63, 47, 47, 47, 47, 47,
	47, 40, 40, 11, 26, 26, 26, 25, 74, 74,
	74, 74, 76, 76, 78, 78, 68, 68, 68, 68,
	68, 33, 33, 33, 33, 33, 101, 33, 33, 33,
	33, 33, 33, 33, 33, 69, 69, 66, 66, 57,
	57, 59, 59, 60, 60, 64, 64, 64, 64, 54,
	54, 84, 84, 93, 93, 34, 34, 67, 67, 37,
	37, 35, 35,
}
var yyR2 = []int{

	0, 0, 4, 0, 3, 0, 3, 2, 5, 3,
	1, 2, 2, 1, 3, 0, 1, 1, 1, 1,
	2, 5, 3, 2, 5, 7, 3, 2, 5, 3,
	1, 2, 4, 3, 4, 3, 1, 2, 1, 1,
	2, 1, 3, 3, 3, 2, 2, 3, 5, 5,
	2, 3, 0, 3, 0, 2, 3, 4, 4, 5,
	1, 1, 2, 2, 1, 3, 5, 4, 0, 2,
	0, 2, 5, 4, 1, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 1, 2, 2, 2, 2,
	2, 2, 2, 2, 3, 5, 6, 1, 1, 3,
	5, 5, 4, 6, 8, 1, 5, 5, 5, 7,
	1, 0, 3, 1, 4, 1, 4, 1, 3, 1,
	1, 1, 1, 1, 1, 1, 0, 1, 1, 1,
	1, 1, 2, 1, 1, 1, 1, 1, 3, 1,
	1, 1, 2, 1, 1, 1, 1, 3, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 3, 4, 4,
	2, 3, 5, 1, 1, 2, 3, 5, 3, 5,
	3, 3, 5, 8, 5, 0, 3, 0, 1, 3,
	1, 4, 2, 0, 3, 1, 3, 1, 3, 1,
	3, 1, 3, 1, 3, 3, 2, 4, 3, 5,
	5, 1, 3, 1, 2, 1, 3, 4, 1, 2,
	2, 1, 1, 3, 0, 2, 0, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 0, 4, 1, 2,
	2, 2, 2, 2, 2, 1, 3, 1, 3, 1,
	3, 1, 3, 1, 3, 1, 1, 3, 3, 0,
	2, 0, 1, 0, 1, 0, 1, 0, 1, 0,
	1, 0, 1,
}
var yyChk = []int{

	-1000, -79, -39, 49, -80, -49, 46, -81, 67, -82,
	-83, 39, -85, -55, -53, -33, 2, 57, -86, 56,
	33, -46, -21, -91, -92, -23, -98, 31, 18, 24,
	35, 28, 36, 51, 23, -12, -59, 32, 55, 53,
	37, -31, -52, -49, -41, 5, 62, 6, 7, 70,
	71, 59, 22, -42, 4, 43, -29, -43, -94, -3,
	-16, -18, -38, -15, 10, 20, 45, -48, -27, 54,
	41, 67, -24, 4, 43, -49, 8, 67, -72, 4,
	-57, -7, -49, -5, 4, -57, -50, 4, -51, -49,
	-14, 4, -49, 68, -36, -31, -49, -36, -43, -41,
	-49, 33, -43, -31, -67, -59, -12, 16, 40, 26,
	48, 14, 30, 47, 13, 42, 34, 12, 6, 7,
	58, 59, 5, 60, 61, 62, 15, 44, 52, 22,
	9, 21, 73, -19, -20, -37, -45, -46, -59, -22,
	-37, 17, -22, 4, 8, 10, -52, -52, -52, -52,
	-52, -52, -52, -52, 11, -13, -12, -97, -99, -18,
	-38, 5, 22, 4, -95, 17, 11, -95, 2, -34,
	25, -12, -100, 22, -18, -38, -44, -8, 4, 5,
	-29, 10, -95, -95, -62, 65, -24, 43, 43, -73,
	65, -72, -32, 9, 73, -99, -18, -38, -44, -8,
	4, 22, -5, 65, -32, 9, -71, 65, -50, -32,
	-61, 11, -78, -76, -74, -30, -49, -75, -32, 25,
	4, -101, 4, -12, -12, -12, -12, -12, -12, -12,
	-12, -12, -12, -12, -12, -12, -12, -12, -12, -12,
	-12, -12, -12, -12, -59, -59, -12, -65, 17, 67,
	9, 21, 17, 67, -90, -65, 65, -60, -13, -49,
	4, -12, -34, -96, 65, -97, 20, -12, -96, -69,
	-68, -4, -55, -33, 2, 11, 72, 72, -32, -32,
	-32, 8, -32, -70, 69, -47, -66, -11, 4, 5,
	-31, -40, 46, -63, 69, -26, -31, -40, 4, -84,
	67, -84, 67, 9, -59, -7, -32, 20, -84, 67,
	9, -59, -84, 67, -69, 65, 73, -93, -30, -75,
	-32, -78, -68, -78, -69, -37, 50, 50, -90, -37,
	-88, 69, -87, 19, 27, -58, -93, 25, 73, -13,
	56, 72, 68, -54, -64, -28, -1, -12, 11, 11,
	22, -32, -18, -38, -44, -8, 4, -93, 73, -54,
	69, 67, -69, -32, -32, 65, -49, 72, -84, 67,
	-32, 73, -35, 43, -11, 5, 46, -11, 4, 8,
	-84, 67, -25, 4, -40, 65, -24, 65, -72, -59,
	65, 65, -56, -6, -5, -57, -59, 65, -50, 69,
	-49, -77, -17, 4, -99, -18, -38, -44, -8, -74,
	65, 65, 69, 67, -12, -12, 69, -89, -60, 68,
	-9, -10, 29, 65, -93, -13, 65, 65, -34, 69,
	73, -93, 68, -96, -96, -32, 65, 69, -68, 69,
	-32, 69, -47, -35, -31, 65, -11, -35, -11, -49,
	69, -26, -78, 65, -84, 67, -32, 4, -78, -77,
	-37, -69, 68, 9, 21, 73, 37, -4, 65, 72,
	68, -28, -1, -2, -12, 11, -54, -54, 65, -35,
	65, 65, 65, 65, -6, -78, 65, -12, -12, -22,
	-34, -96, 69, 69, -35, -35, -77, 65, 68, 68,
	-65, 72, -54, -77, 69,
}
var yyDef = []int{

	3, -2, 1, 0, 5, 0, 138, 193, 4, -2,
	0, 0, 0, 16, 17, 18, 19, 0, 0, 0,
	0, 231, 232, 233, 234, 235, 0, 238, 136, 136,
	0, 0, 0, 267, 30, -2, 0, 269, 269, 0,
	269, 140, 74, -2, 95, 0, 0, 0, 0, 0,
	0, 0, 0, 127, 0, 107, 108, 115, 0, 0,
	120, -2, -2, 0, 265, 0, 0, 173, 174, 0,
	0, 6, 7, 0, 10, 0, 0, 194, 20, 0,
	0, 249, 135, 23, 0, 0, 27, 0, 0, 39,
	185, 224, 0, 236, 239, 137, 134, 240, -2, 0,
	139, 0, -2, 243, 244, 268, 251, 0, 45, 46,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 63, 0, 60, 61, 270, 0, 0,
	64, 54, 0, 0, 0, 265, 96, 97, 98, 99,
	100, 101, 102, 103, 121, 0, 129, 130, 149, -2,
	-2, 0, 0, 0, 121, 132, 133, -2, 192, 0,
	0, 266, 170, 0, 153, 154, 155, 156, 0, 0,
	166, 0, 0, 0, 261, 9, 13, 11, 12, 261,
	22, 195, 31, 0, 0, 143, 144, 145, 146, 147,
	0, 0, 261, 26, 0, 0, 261, 29, 199, 40,
	181, -2, 0, 263, 222, 218, 139, 221, 131, 141,
	224, -2, 224, 42, 75, 76, 77, 78, 79, 80,
	81, 82, 83, 84, 85, 86, 87, 88, 89, 90,
	91, 92, 93, 94, 43, 44, 252, 62, -2, 269,
	0, 0, 54, 269, 0, 68, 104, 263, 253, 109,
	0, 266, 0, 259, 128, 152, 0, 263, 259, 0,
	245, 227, 228, 229, 230, -2, 0, 0, 171, 0,
	175, 0, 0, 261, 178, 201, 0, 271, 0, 0,
	247, 213, -2, 261, 180, 203, 0, 215, 0, 0,
	262, 0, 262, 0, 33, 250, 0, 0, 0, 262,
	0, 35, 0, 262, 0, 187, 264, 225, 219, 220,
	142, 0, 237, 0, 0, 0, 0, 0, 0, 65,
	55, 73, 52, 0, 0, 70, 0, 263, 264, 0,
	0, 112, 265, 0, 263, 255, 256, 123, 121, 121,
	0, 176, -2, -2, -2, -2, 0, 0, 264, 0,
	191, -2, 0, 168, 169, 157, 167, 0, 0, 262,
	271, 0, 206, 272, 0, 0, 211, 271, 0, 0,
	0, 262, 214, 224, 0, 8, 14, 21, 196, 32,
	148, 24, 261, 197, 36, 38, 34, 28, 200, 186,
	139, 184, 188, 224, 161, 162, 163, 164, 165, 223,
	187, 187, 56, 269, 57, 58, 72, -2, 0, 50,
	66, 69, 0, 105, 0, 254, 110, 111, 0, 118,
	264, 260, 0, 259, 259, 0, 116, 117, 246, 51,
	172, 177, 202, 205, 248, 271, 0, 208, 0, 212,
	179, 204, 0, 216, 0, 262, 37, 224, 0, 182,
	59, 53, 47, 0, 0, 0, 269, 71, 106, 113,
	265, 257, 258, 122, 125, 121, 0, 0, -2, 207,
	271, 271, 187, 25, 198, 0, 189, 0, 0, 0,
	0, 259, 124, 119, 209, 210, 217, 187, 48, 49,
	67, 114, 0, 183, 126,
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
			panic(".y:263")
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
			panic(".y:275")
		}
	case 51:
		{ //280
			yyVAL.node = &CompoundStament{yyS[yypt-2].token.pos, yyS[yypt-1].list}
		}
	case 52:
		{ //286
			panic(".y:287")
		}
	case 53:
		{ //290
			panic(".y:291")
		}
	case 54:
		{ //295
			panic(".y:296")
		}
	case 55:
		{ //299
			panic(".y:300")
		}
	case 56:
		{ //305
			yyVAL.list = yyS[yypt-1].list
		}
	case 57:
		{ //311
			yyVAL.node = &ForStmt{Range: &Assignment{yyS[yypt-2].token.pos, token.ASSIGN, yyS[yypt-3].list, []Node{yyS[yypt-0].node}}}
		}
	case 58:
		{ //315
			yyVAL.node = &ForStmt{Range: &Assignment{yyS[yypt-2].token.pos, token.DEFINE, yyS[yypt-3].list, []Node{yyS[yypt-0].node}}}
		}
	case 59:
		{ //321
			yyVAL.node = &ForStmt{Init: yyS[yypt-4].node, Cond: yyS[yypt-2].node, Post: yyS[yypt-0].node}
		}
	case 60:
		{ //325
			yyVAL.node = &ForStmt{Cond: yyS[yypt-0].node}
		}
	case 61:
		yyVAL.node = yyS[yypt-0].node
	case 62:
		{ //335
			yyS[yypt-1].node.(*ForStmt).Body = yyS[yypt-0].list
			yyVAL.node = yyS[yypt-1].node
		}
	case 63:
		{ //341
			yyS[yypt-0].node.(*ForStmt).pos = yyS[yypt-1].token.pos
			yyVAL.node = yyS[yypt-0].node
		}
	case 64:
		{ //347
			yyVAL.node = &IfStmt{Cond: yyS[yypt-0].node}
		}
	case 65:
		{ //351
			yyVAL.node = &IfStmt{Init: yyS[yypt-2].node, Cond: yyS[yypt-0].node}
		}
	case 66:
		{ //357
			x := yyS[yypt-3].node.(*IfStmt)
			l := make([]*IfStmt, len(yyS[yypt-1].list))
			for i, v := range yyS[yypt-1].list {
				l[i] = v.(*IfStmt)
			}
			x.pos, x.Body, x.Elif, x.Else = yyS[yypt-4].token.pos, yyS[yypt-2].list, l, yyS[yypt-0].node.(*CompoundStament)
			yyVAL.node = x
		}
	case 67:
		{ //363
			x := yyS[yypt-1].node.(*IfStmt)
			x.pos, x.Body = yyS[yypt-2].token.pos, yyS[yypt-0].list
			yyVAL.node = x
		}
	case 68:
		{ //368
			yyVAL.list = nil
		}
	case 69:
		{ //372
			yyVAL.list = append(yyS[yypt-1].list, yyS[yypt-0].node)
		}
	case 70:
		{ //377
			yyVAL.node = (*CompoundStament)(nil)
		}
	case 71:
		{ //381
			yyVAL.node = yyS[yypt-0].node
		}
	case 72:
		{ //387
			panic(".y:388")
		}
	case 73:
		{ //393
			panic(".y:394")
		}
	case 74:
		yyVAL.node = yyS[yypt-0].node
	case 75:
		{ //403
			panic(".y:404")
		}
	case 76:
		{ //407
			panic(".y:408")
		}
	case 77:
		{ //411
			yyVAL.node = &BinOp{yyS[yypt-1].token.pos, token.EQL, yyS[yypt-2].node, yyS[yypt-0].node}
		}
	case 78:
		{ //415
			yyVAL.node = &BinOp{yyS[yypt-1].token.pos, token.NEQ, yyS[yypt-2].node, yyS[yypt-0].node}
		}
	case 79:
		{ //419
			yyVAL.node = &BinOp{yyS[yypt-1].token.pos, token.LSS, yyS[yypt-2].node, yyS[yypt-0].node}
		}
	case 80:
		{ //423
			panic(".y:424")
		}
	case 81:
		{ //427
			panic(".y:428")
		}
	case 82:
		{ //431
			yyVAL.node = &BinOp{yyS[yypt-1].token.pos, token.GTR, yyS[yypt-2].node, yyS[yypt-0].node}
		}
	case 83:
		{ //435
			yyVAL.node = &BinOp{yyS[yypt-1].token.pos, token.ADD, yyS[yypt-2].node, yyS[yypt-0].node}
		}
	case 84:
		{ //439
			yyVAL.node = &BinOp{yyS[yypt-1].token.pos, token.SUB, yyS[yypt-2].node, yyS[yypt-0].node}
		}
	case 85:
		{ //443
			panic(".y:444")
		}
	case 86:
		{ //447
			panic(".y:448")
		}
	case 87:
		{ //451
			yyVAL.node = &BinOp{yyS[yypt-1].token.pos, token.MUL, yyS[yypt-2].node, yyS[yypt-0].node}
		}
	case 88:
		{ //455
			panic(".y:456")
		}
	case 89:
		{ //459
			panic(".y:460")
		}
	case 90:
		{ //463
			panic(".y:464")
		}
	case 91:
		{ //467
			panic(".y:468")
		}
	case 92:
		{ //471
			yyVAL.node = &BinOp{yyS[yypt-1].token.pos, token.SHL, yyS[yypt-2].node, yyS[yypt-0].node}
		}
	case 93:
		{ //475
			panic(".y:476")
		}
	case 94:
		{ //479
			panic(".y:480")
		}
	case 95:
		yyVAL.node = yyS[yypt-0].node
	case 96:
		{ //489
			yyVAL.node = &UnOp{yyS[yypt-1].token.pos, token.MUL, yyS[yypt-0].node}
		}
	case 97:
		{ //493
			panic(".y:494")
		}
	case 98:
		{ //497
			panic(".y:498")
		}
	case 99:
		{ //501
			yyVAL.node = &UnOp{yyS[yypt-1].token.pos, token.SUB, yyS[yypt-0].node}
		}
	case 100:
		{ //505
			panic(".y:506")
		}
	case 101:
		{ //509
			panic(".y:510")
		}
	case 102:
		{ //513
			panic(".y:514")
		}
	case 103:
		{ //517
			panic(".y:518")
		}
	case 104:
		{ //523
			yyVAL.node = &CallOp{yyS[yypt-1].token.pos, yyS[yypt-2].node, nil}
		}
	case 105:
		{ //527
			yyVAL.node = &CallOp{yyS[yypt-3].token.pos, yyS[yypt-4].node, yyS[yypt-2].list}
		}
	case 106:
		{ //531
			panic(".y:532")
		}
	case 107:
		{ //537
			yyVAL.node = &Literal{yyS[yypt-0].token.pos, yyS[yypt-0].token.tok, yyS[yypt-0].token.lit}
		}
	case 108:
		yyVAL.node = yyS[yypt-0].node
	case 109:
		{ //545
			yyVAL.node = &SelectOp{yyS[yypt-1].token.pos, yyS[yypt-2].node, yyS[yypt-0].node.(*Ident)}
		}
	case 110:
		{ //549
			panic(".y:550")
		}
	case 111:
		{ //553
			panic(".y:554")
		}
	case 112:
		{ //557
			yyVAL.node = &IndexOp{yyS[yypt-2].token.pos, yyS[yypt-3].node, yyS[yypt-1].node}
		}
	case 113:
		{ //561
			panic(".y:562")
		}
	case 114:
		{ //565
			panic(".y:566")
		}
	case 115:
		yyVAL.node = yyS[yypt-0].node
	case 116:
		{ //573
			panic(".y:574")
		}
	case 117:
		{ //577
			yyVAL.node = &CompLit{pos(yyS[yypt-4].node.Pos()), yyS[yypt-4].node, elements(yyS[yypt-1].list)}
		}
	case 118:
		{ //581
			yyVAL.node = &CompLit{pos(yyS[yypt-4].node.Pos()), yyS[yypt-4].node, elements(yyS[yypt-1].list)}
		}
	case 119:
		{ //585
			panic(".y:586")
		}
	case 120:
		yyVAL.node = yyS[yypt-0].node
	case 121:
		{ //594
		}
	case 122:
		{ //600
			yyVAL.node = &Element{pos(yyS[yypt-2].node.Pos()), yyS[yypt-2].node, yyS[yypt-0].node}
		}
	case 123:
		{ //609
			yyVAL.node = &Element{pos(yyS[yypt-0].node.Pos()), nil, yyS[yypt-0].node}
		}
	case 124:
		{ //610
			panic(".y:611")
		}
	case 125:
		{ //616
			yyVAL.node = &Element{pos(yyS[yypt-0].node.Pos()), nil, yyS[yypt-0].node}
		}
	case 126:
		{ //620
			panic(".y:621")
		}
	case 127:
		yyVAL.node = yyS[yypt-0].node
	case 128:
		{ //630
			yyVAL.node = &Paren{yyS[yypt-2].token.pos, yyS[yypt-1].node}
		}
	case 129:
		yyVAL.node = yyS[yypt-0].node
	case 130:
		{ //640
			panic(".y:641")
		}
	case 131:
		yyVAL.node = yyS[yypt-0].node
	case 134:
		yyVAL.node = yyS[yypt-0].node
	case 135:
		yyVAL.node = yyS[yypt-0].node
	case 136:
		{ //673
			yyVAL.node = (*Ident)(nil)
		}
	case 137:
		yyVAL.node = yyS[yypt-0].node
	case 138:
		{ //683
			yyVAL.node = &Ident{yyS[yypt-0].token.pos, yyS[yypt-0].token.lit}
		}
	case 139:
		yyVAL.node = yyS[yypt-0].node
	case 140:
		{ //695
			panic(".y:696")
		}
	case 141:
		{ //701
			yy(yylex).errPos(yyS[yypt-0].token.tpos, "final argument in variadic function missing type")
			yyVAL.param = &Param{pos: yyS[yypt-0].token.pos, Ddd: true}
		}
	case 142:
		{ //705
			yyVAL.param = &Param{pos: yyS[yypt-1].token.pos, Ddd: true, Type: yyS[yypt-0].node}
		}
	case 143:
		{ //711
			panic(".y:712")
		}
	case 144:
		yyVAL.node = yyS[yypt-0].node
	case 145:
		yyVAL.node = yyS[yypt-0].node
	case 146:
		yyVAL.node = yyS[yypt-0].node
	case 147:
		{
			yyVAL.node = &NamedType{pos(yyS[yypt-0].node.Pos()), yyS[yypt-0].node.(*QualifiedIdent), nil, yyScope(yylex)}
		}
	case 148:
		{ //731
			yyVAL.node = &Paren{yyS[yypt-2].token.pos, yyS[yypt-1].node}
		}
	case 149:
		{ //737
			panic(".y:738")
		}
	case 150:
		{ //741
			panic(".y:742")
		}
	case 151:
		{ //745
			panic(".y:746")
		}
	case 152:
		{ //749
			panic(".y:750")
		}
	case 153:
		{ //755
			panic(".y:756")
		}
	case 154:
		{ //759
			panic(".y:760")
		}
	case 155:
		{ //763
			panic(".y:764")
		}
	case 156:
		{ //767
			panic(".y:768")
		}
	case 157:
		{ //771
			panic(".y:772")
		}
	case 158:
		{ //777
			panic(".y:778")
		}
	case 159:
		{ //781
			panic(".y:782")
		}
	case 160:
		yyVAL.node = yyS[yypt-0].node
	case 161:
		{ //793
			panic(".y:794")
		}
	case 162:
		{ //797
			panic(".y:798")
		}
	case 163:
		{ //801
			panic(".y:802")
		}
	case 164:
		{ //805
			panic(".y:806")
		}
	case 165:
		{ //790
			yyVAL.node = &NamedType{pos(yyS[yypt-0].node.Pos()), yyS[yypt-0].node.(*QualifiedIdent), nil, yyScope(yylex)}
		}
	case 166:
		{ //815
			yyVAL.node = &QualifiedIdent{pos(yyS[yypt-0].node.Pos()), nil, yyS[yypt-0].node.(*Ident)}
		}
	case 167:
		{ //819
			yyVAL.node = &QualifiedIdent{pos(yyS[yypt-2].node.Pos()), yyS[yypt-2].node.(*Ident), yyS[yypt-0].node.(*Ident)}
		}
	case 168:
		{ //825
			switch {
			case yyS[yypt-2].node != nil:
				yyVAL.node = &ArrayType{yyS[yypt-3].token.pos, yyS[yypt-2].node, yyS[yypt-0].node}
			default:
				yyVAL.node = &SliceType{yyS[yypt-3].token.pos, yyS[yypt-0].node}
			}
		}
	case 169:
		{ //829
			yyVAL.node = &ArrayType{yyS[yypt-3].token.pos, nil, yyS[yypt-0].node}
		}
	case 170:
		{ //833
			panic(".y:834")
		}
	case 171:
		{ //837
			panic(".y:838")
		}
	case 172:
		{ //841
			panic(".y:842")
		}
	case 173:
		yyVAL.node = yyS[yypt-0].node
	case 174:
		yyVAL.node = yyS[yypt-0].node
	case 175:
		{ //855
			yyVAL.node = &PtrType{yyS[yypt-1].token.pos, yyS[yypt-0].node}
		}
	case 176:
		{ //861
			panic(".y:862")
		}
	case 177:
		{ //867
			yyVAL.node = newStructType(yylex, yyS[yypt-4].token, yyS[yypt-2].list)
		}
	case 178:
		{ //871
			yyVAL.node = newStructType(yylex, yyS[yypt-2].token, nil)
		}
	case 179:
		{ //877
			x := newInterfaceType(yylex, yyS[yypt-2].list)
			x.pos = yyS[yypt-4].token.pos
			yyVAL.node = x
		}
	case 180:
		{ //881
			panic(".y:882")
		}
	case 181:
		{ //887
			x := yyS[yypt-1].node.(*FuncDecl)
			x.pos, x.Body = yyS[yypt-2].token.pos, yyS[yypt-0].list
			yyVAL.node = x
		}
	case 182:
		{ //893
			yyVAL.node = &FuncDecl{Name: (*Name)(yyS[yypt-4].node.(*Ident)), Type: newFuncType(yylex, yyS[yypt-3].token.pos, yyS[yypt-2].params, yyS[yypt-0].params)}
		}
	case 183:
		{ //897
			panic(".y:898")
		}
	case 184:
		{ //903
			yyVAL.node = newFuncType(yylex, yyS[yypt-4].token.pos, yyS[yypt-2].params, yyS[yypt-0].params)
		}
	case 185:
		{ //908
			yyVAL.list = nil
		}
	case 186:
		{ //912
			yyVAL.list = yyS[yypt-1].list
		}
	case 187:
		{ //918
			yyVAL.params = nil
		}
	case 188:
		{ //922
			yyVAL.params = []*Param{{pos: pos(yyS[yypt-0].node.Pos()), Type: yyS[yypt-0].node}}
		}
	case 189:
		{ //926
			yyVAL.params = yyS[yypt-1].params
		}
	case 190:
		yyVAL.node = yyS[yypt-0].node
	case 191:
		{ //938
			x := yyS[yypt-3].node.(*FuncType)
			yyVAL.node = &FuncLit{x.pos, x, yyS[yypt-1].list}
		}
	case 192:
		{ //942
			panic(".y:943")
		}
	case 195:
		{ //957
			yyVAL.list = yyS[yypt-0].list
		}
	case 196:
		{ //961
			yyVAL.list = append(yyS[yypt-2].list, yyS[yypt-0].list...)
		}
	case 197:
		{ //967
			yyVAL.list = []Node{yyS[yypt-0].node}
		}
	case 198:
		{ //971
			yyVAL.list = append(yyS[yypt-2].list, yyS[yypt-0].node)
		}
	case 199:
		{ //977
			yyVAL.list = []Node{yyS[yypt-0].node}
		}
	case 200:
		{ //981
			yyVAL.list = append(yyS[yypt-2].list, yyS[yypt-0].node)
		}
	case 201:
		{ //987
			yyVAL.list = []Node{yyS[yypt-0].node}
		}
	case 202:
		{ //991
			yyVAL.list = append(yyS[yypt-2].list, yyS[yypt-0].node)
		}
	case 203:
		{ //997
			yyVAL.list = []Node{yyS[yypt-0].node}
		}
	case 204:
		{ //1001
			yyVAL.list = append(yyS[yypt-2].list, yyS[yypt-0].node)
		}
	case 205:
		{ //1007
			yyVAL.node = newFields(yyS[yypt-2].list, false, yyS[yypt-1].node, yyS[yypt-0].node)
		}
	case 206:
		{ //1011
			q := yyS[yypt-1].node.(*QualifiedIdent)
			yyVAL.node = newFields([]Node{q.I}, true, &NamedType{q.pos, q, nil, yyScope(yylex)}, yyS[yypt-0].node)
		}
	case 207:
		{ //1015
			panic(".y:1016")
		}
	case 208:
		{ //1019
			q := yyS[yypt-1].node.(*QualifiedIdent)
			yyVAL.node = newFields([]Node{q.I}, true, &PtrType{yyS[yypt-2].token.pos, &NamedType{q.pos, q, nil, yyScope(yylex)}}, yyS[yypt-0].node)
		}
	case 209:
		{ //1023
			panic(".y:1024")
		}
	case 210:
		{ //1027
			panic(".y:1028")
		}
	case 211:
		{ //1033
			yyVAL.node = &QualifiedIdent{yyS[yypt-0].token.pos, nil, &Ident{yyS[yypt-0].token.pos, yyS[yypt-0].token.lit}}
		}
	case 212:
		{ //1037
			yyVAL.node = &QualifiedIdent{yyS[yypt-2].token.pos, &Ident{yyS[yypt-2].token.pos, yyS[yypt-2].token.lit}, yyS[yypt-0].node.(*Ident)}
		}
	case 213:
		yyVAL.node = yyS[yypt-0].node
	case 214:
		{ //1049
			yyVAL.node = &MethodSpec{pos(yyS[yypt-1].node.Pos()), yyS[yypt-1].node.(*Ident), yyS[yypt-0].node.(*FuncType)}
		}
	case 215:
		{ //1053
			panic(".y:1054")
		}
	case 216:
		{ //1057
			panic(".y:1058")
			//yyErrPos(yylex, $2, "cannot parenthesize embedded type");
		}
	case 217:
		{ //1063
			yyVAL.node = newFuncType(yylex, yyS[yypt-3].token.pos, yyS[yypt-2].params, yyS[yypt-0].params)
		}
	case 218:
		{ //1069
			yyVAL.param = &Param{pos: pos(yyS[yypt-0].node.Pos()), Type: yyS[yypt-0].node}
		}
	case 219:
		{ //1073
			yyVAL.param = &Param{pos: pos(yyS[yypt-1].node.Pos()), Name: yyS[yypt-1].node.(*Ident), Type: yyS[yypt-0].node}
		}
	case 220:
		{ //1077
			x := yyS[yypt-0].param
			x.Name, x.Ddd = yyS[yypt-1].node.(*Ident), true
			yyVAL.param = x
		}
	case 221:
		yyVAL.param = yyS[yypt-0].param
	case 222:
		{ //1087
			yyVAL.params = []*Param{yyS[yypt-0].param}
		}
	case 223:
		{ //1091
			yyVAL.params = append(yyS[yypt-2].params, yyS[yypt-0].param)
		}
	case 224:
		{ //1096
			yyVAL.params = nil
		}
	case 225:
		{
			yyVAL.params = yyS[yypt-1].params
		}
	case 226:
		{ //1105
			yyVAL.list = nil
		}
	case 227:
		{ //1109
			yyVAL.list = []Node{yyS[yypt-0].node}
		}
	case 228:
		yyVAL.list = yyS[yypt-0].list
	case 229:
		{ //1117
			yyVAL.list = []Node{yyS[yypt-0].node}
		}
	case 230:
		{ //1121
			yyVAL.list = nil
		}
	case 231:
		yyVAL.node = yyS[yypt-0].node
	case 232:
		yyVAL.node = yyS[yypt-0].node
	case 233:
		{ //1135
			panic(".y:1136")
		}
	case 234:
		{ //1139
			panic(".y:1140")
		}
	case 235:
		yyVAL.node = yyS[yypt-0].node
	case 236:
		{ //1147
			panic(".y:1148")
		}
	case 237:
		{ //1151
			panic(".y:1152")
		}
	case 238:
		{ //1155
			yyVAL.node = &FallthroughStmt{yyS[yypt-0].token.pos}
		}
	case 239:
		{ //1159
			yyVAL.node = &BreakStmt{yyS[yypt-1].token.pos, yyS[yypt-0].node.(*Ident)}
		}
	case 240:
		{ //1163
			yyVAL.node = &ContinueStmt{yyS[yypt-1].token.pos, yyS[yypt-0].node.(*Ident)}
		}
	case 241:
		{ //1167
			panic(".y:1168")
		}
	case 242:
		{ //1171
			panic(".y:1172")
		}
	case 243:
		{ //1175
			yyVAL.node = &GotoStmt{yyS[yypt-1].token.pos, yyS[yypt-0].node.(*Ident)}
		}
	case 244:
		{ //1179
			yyVAL.node = &ReturnStmt{yyS[yypt-1].token.pos, yyS[yypt-0].list}
		}
	case 245:
		yyVAL.list = yyS[yypt-0].list
	case 246:
		{ //1189
			yyVAL.list = append(yyS[yypt-2].list, yyS[yypt-0].list...)
		}
	case 247:
		{ //1195
			yyVAL.list = []Node{yyS[yypt-0].node}
		}
	case 248:
		{ //1199
			yyVAL.list = append(yyS[yypt-2].list, yyS[yypt-0].node)
		}
	case 249:
		{ //1205
			yyVAL.list = []Node{yyS[yypt-0].node}
		}
	case 250:
		{ //1209
			yyVAL.list = append(yyS[yypt-2].list, yyS[yypt-0].node)
		}
	case 251:
		{ //1215
			yyVAL.list = []Node{yyS[yypt-0].node}
		}
	case 252:
		{ //1219
			yyVAL.list = append(yyS[yypt-2].list, yyS[yypt-0].node)
		}
	case 253:
		{ //1225
			yyVAL.list = []Node{yyS[yypt-0].node}
		}
	case 254:
		{ //1229
			panic(".y:1230")
		}
	case 255:
		{ //1235
			yyVAL.list = []Node{yyS[yypt-0].node}
		}
	case 256:
		{
			yyVAL.list = []Node{yyS[yypt-0].node}
		}
	case 257:
		{ //1243
			yyVAL.list = append(yyS[yypt-2].list, yyS[yypt-0].node)
		}
	case 258:
		{ //1247
			yyVAL.list = append(yyS[yypt-2].list, yyS[yypt-0].node)
		}
	case 259:
		{ //1252
			yyVAL.list = nil
		}
	case 260:
		{ //1256
			yyVAL.list = yyS[yypt-1].list
		}
	case 265:
		{ //1279
			yyVAL.node = nil
		}
	case 266:
		yyVAL.node = yyS[yypt-0].node
	case 267:
		{ //1288
			yyVAL.list = nil
		}
	case 268:
		yyVAL.list = yyS[yypt-0].list
	case 269:
		{ //1297
			yyVAL.node = nil
		}
	case 270:
		yyVAL.node = yyS[yypt-0].node
	case 271:
		{ //1306
			yyVAL.node = (*Literal)(nil)
		}
	case 272:
		{ //1310
			panic(".y:1311")
		}
	}
	goto yystack /* stack new state and value */
}
