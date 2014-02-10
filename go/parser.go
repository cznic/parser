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
const _GT = 57369
const _IF = 57370
const _IGNORE = 57371
const _IMPORT = 57372
const _INC = 57373
const _INTERFACE = 57374
const _LE = 57375
const _LITERAL = 57376
const _LSH = 57377
const _LT = 57378
const _MAP = 57379
const _NAME = 57380
const _NE = 57381
const _OROR = 57382
const _PACKAGE = 57383
const _RANGE = 57384
const _RETURN = 57385
const _RSH = 57386
const _SELECT = 57387
const _STRUCT = 57388
const _SWITCH = 57389
const _TYPE = 57390
const _VAR = 57391
const notPackage = 57392
const notParen = 57393
const preferToRightParen = 57394

var yyToknames = []string{
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
	"_GT",
	"_IF",
	"_IGNORE",
	"_IMPORT",
	"_INC",
	"_INTERFACE",
	"_LE",
	"_LITERAL",
	"_LSH",
	"_LT",
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
	" .",
	" -",
	" *",
	" [",
	" (",
	" +",
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
	65, 15,
	-2, 0,
	-1, 35,
	11, 251,
	66, 251,
	73, 251,
	-2, 41,
	-1, 43,
	67, 134,
	-2, 139,
	-1, 61,
	54, 158,
	-2, 190,
	-1, 62,
	54, 159,
	-2, 160,
	-1, 98,
	50, 115,
	53, 115,
	54, 115,
	68, 115,
	-2, 241,
	-1, 102,
	50, 115,
	53, 115,
	54, 115,
	68, 115,
	-2, 242,
	-1, 159,
	2, 190,
	7, 190,
	54, 158,
	68, 190,
	-2, 150,
	-1, 160,
	7, 160,
	54, 159,
	68, 160,
	-2, 151,
	-1, 167,
	65, 226,
	69, 226,
	-2, 0,
	-1, 211,
	65, 226,
	69, 226,
	-2, 0,
	-1, 221,
	9, 226,
	17, 226,
	65, 226,
	69, 226,
	-2, 0,
	-1, 248,
	65, 226,
	69, 226,
	-2, 0,
	-1, 275,
	65, 226,
	69, 226,
	-2, 0,
	-1, 292,
	34, 211,
	65, 211,
	69, 211,
	-2, 138,
	-1, 352,
	7, 153,
	54, 153,
	68, 153,
	-2, 144,
	-1, 353,
	7, 154,
	54, 154,
	68, 154,
	-2, 145,
	-1, 354,
	7, 155,
	54, 155,
	68, 155,
	-2, 146,
	-1, 355,
	7, 156,
	54, 156,
	68, 156,
	-2, 147,
	-1, 361,
	9, 226,
	17, 226,
	65, 226,
	69, 226,
	-2, 0,
	-1, 417,
	9, 226,
	17, 226,
	65, 226,
	69, 226,
	-2, 0,
	-1, 478,
	7, 157,
	54, 157,
	68, 157,
	-2, 148,
}

const yyNprod = 273
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1888

var yyAct = []int{

	35, 62, 343, 169, 80, 247, 263, 393, 401, 21,
	346, 258, 212, 345, 195, 271, 139, 198, 295, 285,
	270, 214, 299, 85, 269, 291, 257, 254, 199, 78,
	287, 217, 215, 317, 106, 140, 72, 81, 157, 372,
	470, 132, 430, 358, 316, 469, 501, 137, 137, 394,
	137, 86, 42, 367, 277, 156, 160, 142, 276, 361,
	361, 292, 504, 439, 412, 171, 155, 175, 83, 158,
	361, 493, 361, 135, 399, 289, 360, 288, 492, 450,
	464, 441, 197, 176, 437, 251, 429, 197, 131, 85,
	197, 164, 284, 197, 177, 349, 337, 211, 146, 147,
	148, 149, 150, 151, 152, 153, 154, 419, 223, 191,
	186, 224, 225, 226, 227, 228, 229, 230, 231, 232,
	233, 234, 235, 236, 237, 238, 239, 240, 241, 242,
	243, 106, 106, 246, 202, 463, 462, 111, 126, 208,
	250, 310, 465, 130, 156, 160, 261, 132, 255, 262,
	132, 168, 165, 112, 338, 167, 165, 116, 158, 333,
	117, 182, 183, 160, 267, 466, 115, 334, 127, 114,
	342, 268, 113, 110, 333, 197, 158, 128, 41, 292,
	197, 197, 334, 197, 119, 122, 93, 303, 118, 120,
	121, 123, 124, 125, 106, 298, 361, 455, 497, 413,
	265, 61, 197, 381, 369, 275, 106, 95, 95, 297,
	294, 103, 301, 166, 146, 153, 313, 166, 197, 416,
	309, 197, 197, 302, 197, 308, 300, 253, 249, 312,
	6, 77, 305, 321, 331, 323, 314, 74, 71, 8,
	6, 6, 322, 486, 483, 6, 482, 481, 319, 318,
	480, 106, 106, 76, 478, 207, 159, 468, 453, 137,
	445, 156, 160, 137, 347, 203, 185, 174, 353, 347,
	190, 359, 339, 324, 436, 158, 427, 426, 197, 197,
	328, 423, 196, 411, 354, 325, 410, 196, 197, 329,
	196, 336, 397, 196, 391, 355, 390, 387, 385, 365,
	362, 357, 315, 264, 106, 457, 368, 292, 383, 197,
	292, 106, 222, 376, 395, 6, 380, 406, 197, 374,
	377, 289, 56, 288, 384, 6, 298, 414, 415, 378,
	404, 260, 388, 407, 156, 160, 220, 386, 409, 156,
	160, 91, 163, 171, 408, 159, 428, 181, 158, 379,
	425, 281, 197, 158, 3, 433, 434, 6, 197, 6,
	418, 290, 296, 159, 144, 398, 376, 145, 143, 197,
	74, 424, 11, 87, 6, 196, 6, 6, 431, 373,
	196, 196, 438, 196, 376, 197, 76, 188, 180, 442,
	73, 187, 84, 79, 422, 307, 452, 197, 375, 248,
	451, 252, 196, 180, 141, 197, 446, 297, 180, 448,
	443, 180, 406, 406, 180, 454, 458, 447, 196, 459,
	57, 196, 196, 137, 196, 404, 404, 344, 407, 407,
	104, 347, 221, 474, 347, 347, 476, 477, 467, 408,
	408, 472, 461, 94, 471, 63, 272, 210, 273, 460,
	90, 98, 102, 74, 172, 44, 13, 6, 15, 197,
	395, 26, 159, 484, 487, 488, 156, 160, 352, 76,
	485, 171, 473, 97, 490, 126, 137, 425, 196, 196,
	158, 60, 491, 489, 406, 479, 99, 99, 196, 59,
	58, 496, 347, 24, 502, 500, 180, 404, 23, 406,
	407, 180, 180, 36, 180, 127, 503, 421, 420, 196,
	335, 408, 404, 25, 128, 407, 22, 405, 196, 133,
	494, 495, 122, 180, 134, 136, 408, 417, 123, 124,
	125, 330, 332, 18, 14, 159, 12, 105, 10, 180,
	159, 138, 180, 180, 9, 180, 7, 4, 290, 1,
	444, 213, 196, 189, 206, 283, 286, 293, 196, 184,
	296, 392, 100, 88, 67, 53, 5, 2, 68, 196,
	382, 402, 43, 0, 75, 0, 0, 0, 0, 0,
	82, 82, 89, 92, 0, 196, 0, 0, 0, 180,
	0, 96, 96, 0, 0, 96, 0, 196, 0, 180,
	180, 65, 0, 201, 0, 196, 0, 218, 0, 180,
	0, 0, 405, 405, 101, 0, 0, 0, 0, 0,
	0, 0, 0, 70, 0, 0, 0, 0, 66, 6,
	180, 0, 0, 0, 244, 245, 75, 69, 180, 180,
	0, 0, 82, 179, 64, 200, 0, 82, 0, 0,
	89, 0, 0, 0, 216, 0, 0, 205, 65, 196,
	201, 0, 0, 0, 194, 0, 0, 159, 0, 0,
	0, 101, 0, 180, 0, 0, 0, 0, 0, 180,
	70, 0, 0, 0, 405, 66, 6, 0, 192, 0,
	180, 0, 0, 204, 69, 65, 209, 304, 0, 405,
	179, 64, 200, 0, 0, 0, 180, 259, 101, 311,
	0, 0, 111, 126, 193, 0, 0, 70, 180, 55,
	129, 194, 66, 6, 0, 0, 180, 0, 112, 0,
	43, 69, 116, 180, 180, 117, 0, 0, 64, 54,
	0, 115, 0, 127, 114, 96, 96, 113, 110, 0,
	0, 0, 128, 0, 244, 245, 0, 82, 0, 119,
	122, 0, 0, 118, 120, 121, 123, 124, 125, 0,
	0, 0, 0, 0, 43, 499, 0, 0, 0, 0,
	180, 278, 0, 216, 43, 216, 279, 280, 0, 282,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 111, 126, 0, 0, 180, 0, 389, 306, 129,
	0, 43, 0, 0, 396, 0, 0, 112, 0, 0,
	180, 116, 0, 0, 117, 0, 0, 320, 0, 0,
	115, 0, 127, 114, 0, 0, 113, 110, 43, 0,
	0, 128, 0, 0, 366, 0, 0, 0, 119, 122,
	0, 0, 118, 120, 121, 123, 124, 125, 0, 0,
	0, 0, 0, 75, 0, 82, 0, 0, 0, 65,
	358, 52, 82, 0, 351, 0, 89, 0, 400, 216,
	0, 0, 101, 0, 363, 364, 0, 0, 0, 0,
	0, 70, 0, 55, 370, 0, 66, 6, 0, 0,
	0, 0, 0, 0, 0, 69, 0, 0, 0, 0,
	48, 45, 64, 54, 47, 351, 51, 0, 0, 46,
	0, 0, 0, 0, 43, 0, 0, 348, 0, 49,
	50, 0, 96, 0, 96, 0, 0, 0, 0, 0,
	0, 0, 449, 0, 96, 0, 216, 0, 0, 65,
	0, 201, 0, 0, 0, 0, 0, 0, 278, 0,
	0, 0, 101, 0, 435, 0, 216, 0, 0, 0,
	0, 70, 0, 0, 0, 440, 66, 6, 274, 0,
	43, 0, 0, 0, 28, 69, 65, 0, 52, 34,
	29, 179, 64, 200, 31, 0, 0, 27, 37, 101,
	0, 30, 32, 456, 40, 0, 0, 0, 70, 0,
	55, 0, 371, 66, 6, 0, 0, 0, 82, 33,
	216, 39, 69, 38, 19, 17, 0, 48, 45, 64,
	54, 47, 0, 51, 0, 0, 46, 0, 0, 0,
	0, 0, 16, 0, 275, 0, 49, 50, 28, 0,
	65, 0, 52, 34, 29, 0, 0, 0, 31, 0,
	0, 27, 37, 20, 0, 30, 32, 0, 40, 0,
	0, 0, 70, 0, 55, 0, 0, 66, 6, 0,
	0, 0, 0, 33, 0, 39, 69, 38, 19, 17,
	0, 48, 45, 64, 54, 47, 0, 51, 111, 126,
	46, 0, 0, 0, 0, 0, 129, 0, 0, 0,
	49, 50, 0, 0, 112, 0, 0, 0, 116, 0,
	0, 117, 0, 0, 0, 0, 0, 115, 0, 127,
	114, 0, 0, 113, 110, 65, 0, 52, 128, 0,
	0, 0, 0, 0, 0, 119, 122, 0, 101, 118,
	120, 121, 123, 124, 125, 0, 0, 70, 0, 55,
	0, 0, 66, 6, 0, 0, 341, 0, 0, 0,
	65, 69, 162, 0, 0, 0, 48, 45, 64, 54,
	47, 0, 51, 101, 0, 46, 0, 0, 0, 0,
	0, 0, 70, 475, 55, 49, 50, 66, 6, 0,
	0, 0, 0, 0, 0, 65, 69, 52, 340, 0,
	0, 48, 161, 64, 54, 47, 0, 51, 101, 0,
	46, 0, 0, 0, 0, 0, 0, 70, 0, 55,
	49, 50, 66, 6, 0, 0, 0, 327, 0, 0,
	65, 69, 52, 0, 0, 0, 48, 45, 64, 54,
	47, 0, 51, 101, 0, 46, 0, 0, 0, 0,
	0, 0, 70, 0, 55, 49, 50, 66, 6, 0,
	0, 0, 326, 0, 0, 65, 69, 162, 0, 0,
	0, 48, 45, 64, 54, 47, 0, 51, 101, 0,
	46, 0, 0, 0, 0, 0, 0, 70, 0, 55,
	49, 50, 66, 6, 0, 0, 0, 0, 0, 0,
	65, 69, 52, 0, 0, 170, 48, 161, 64, 54,
	47, 0, 51, 101, 0, 46, 0, 0, 256, 0,
	0, 0, 70, 0, 55, 49, 50, 66, 6, 0,
	0, 0, 0, 0, 0, 0, 69, 0, 0, 0,
	0, 48, 45, 64, 54, 47, 0, 51, 111, 126,
	46, 0, 0, 0, 0, 0, 129, 0, 0, 0,
	49, 50, 0, 0, 112, 0, 0, 0, 116, 0,
	0, 117, 0, 0, 0, 0, 0, 115, 0, 127,
	114, 0, 0, 113, 110, 65, 0, 52, 128, 0,
	0, 0, 0, 0, 0, 119, 122, 0, 101, 118,
	120, 121, 123, 124, 125, 0, 0, 70, 0, 55,
	0, 498, 66, 6, 0, 0, 0, 0, 0, 0,
	65, 69, 162, 0, 0, 0, 48, 45, 64, 54,
	47, 0, 51, 101, 0, 46, 0, 0, 0, 0,
	0, 0, 70, 0, 55, 49, 50, 66, 6, 0,
	0, 0, 0, 0, 0, 0, 69, 0, 0, 0,
	0, 48, 161, 64, 54, 47, 0, 51, 111, 126,
	46, 0, 0, 0, 0, 0, 129, 0, 0, 0,
	49, 50, 0, 0, 112, 0, 0, 0, 116, 0,
	0, 117, 0, 0, 0, 0, 0, 115, 0, 127,
	114, 0, 0, 113, 110, 266, 0, 52, 128, 0,
	0, 0, 0, 0, 0, 119, 122, 0, 101, 118,
	120, 121, 123, 124, 125, 0, 0, 70, 0, 55,
	0, 432, 66, 6, 0, 0, 0, 0, 0, 0,
	0, 69, 0, 0, 0, 0, 48, 45, 64, 54,
	47, 0, 51, 0, 0, 46, 111, 126, 107, 0,
	0, 0, 0, 0, 129, 49, 50, 0, 109, 0,
	0, 0, 112, 0, 0, 0, 116, 0, 0, 117,
	0, 0, 0, 108, 0, 115, 0, 127, 114, 0,
	0, 113, 110, 0, 0, 0, 128, 111, 126, 0,
	0, 0, 0, 119, 122, 129, 0, 118, 120, 121,
	123, 124, 125, 112, 0, 0, 0, 116, 0, 0,
	117, 0, 0, 0, 0, 0, 115, 0, 127, 114,
	0, 0, 113, 110, 0, 0, 0, 128, 111, 126,
	0, 0, 0, 0, 119, 122, 0, 0, 118, 120,
	121, 123, 124, 125, 112, 0, 0, 0, 116, 0,
	0, 117, 0, 0, 0, 0, 0, 115, 0, 127,
	114, 0, 0, 113, 0, 0, 0, 0, 128, 0,
	126, 0, 0, 0, 0, 119, 122, 0, 0, 118,
	120, 121, 123, 124, 125, 112, 0, 0, 126, 116,
	0, 0, 117, 0, 0, 0, 0, 0, 115, 0,
	127, 114, 0, 0, 113, 0, 0, 0, 0, 128,
	0, 0, 0, 0, 0, 0, 119, 122, 127, 0,
	118, 120, 121, 123, 124, 125, 0, 128, 0, 0,
	0, 65, 0, 201, 119, 122, 219, 0, 118, 120,
	121, 123, 124, 125, 101, 0, 0, 0, 0, 65,
	0, 201, 0, 70, 0, 0, 0, 0, 66, 6,
	0, 0, 101, 0, 0, 0, 0, 69, 65, 0,
	201, 70, 0, 179, 64, 200, 66, 6, 0, 0,
	0, 101, 0, 0, 0, 69, 266, 0, 201, 0,
	70, 179, 64, 403, 0, 66, 6, 0, 0, 101,
	0, 0, 0, 0, 69, 65, 0, 350, 70, 0,
	179, 64, 200, 66, 6, 0, 0, 0, 101, 0,
	0, 0, 69, 65, 0, 173, 0, 70, 179, 64,
	200, 0, 66, 6, 0, 0, 101, 0, 0, 0,
	0, 69, 0, 0, 0, 70, 0, 179, 64, 356,
	66, 6, 0, 0, 0, 0, 0, 0, 0, 69,
	0, 0, 0, 0, 0, 179, 64, 178,
}
var yyPact = []int{

	313, -1000, -1000, 321, -1000, 174, -1000, 342, -1000, 1040,
	173, 336, 166, -1000, -1000, -1000, -1000, 339, 338, 319,
	287, -1000, -1000, -1000, -1000, -1000, 119, -1000, 321, 321,
	685, 685, 321, 1385, -1000, 1562, 77, 1385, 1385, 397,
	1385, -1000, -1000, -1000, 314, 1385, 1385, 1385, 1385, 1385,
	1385, 1385, 1385, 38, 1420, -1000, -1000, -1000, 288, 145,
	-1000, -1000, -1000, 149, 1300, 1833, 294, -1000, -1000, 145,
	145, -1000, -1000, 203, -1000, 357, 353, -1000, -1000, 207,
	648, -1000, -1000, -1000, 202, 591, -1000, 192, 1778, -1000,
	29, 1741, 282, -1000, -1000, -1000, -1000, -1000, -1000, 314,
	-1000, 258, -1000, -1000, -1000, -32, 1603, 1385, -1000, -1000,
	1385, 1385, 1385, 1385, 1385, 1385, 1385, 1385, 1385, 1385,
	1385, 1385, 1385, 1385, 1385, 1385, 1385, 1385, 1385, 1385,
	1385, 1385, 1385, -1000, 392, 163, -1000, -1000, 74, 394,
	162, -1000, 392, 1265, 277, 1385, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 240, 1603, -1000, -1000, -1000,
	-1000, 1420, 1505, 1385, -1000, -1000, -1000, 976, -1000, -14,
	-18, 1603, -1000, 1778, -1000, -1000, -1000, -1000, 1778, 1778,
	301, 1778, 23, 141, 161, -1000, -1000, -1000, -1000, 158,
	-1000, -1000, 121, 1385, 321, -1000, -1000, -1000, -1000, -1000,
	1778, 385, 155, -1000, 75, 1385, 151, -1000, -1000, -1000,
	-1000, 976, 239, -29, -1000, -1000, 1741, -1000, -1000, 1778,
	1741, 976, 1741, 1603, 1644, 1685, 1703, 1703, 1703, 1703,
	1703, 1703, 470, 470, 470, 470, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 133, -32, -32, 1603, -1000, 976, 1385,
	1230, 1195, -1000, 1385, 165, -1000, -1000, 81, -1000, -1000,
	1160, 1094, 103, 859, 27, -1000, 1815, 797, 859, 7,
	-1000, -1000, -1000, -1000, -1000, 976, 1778, 1778, -1000, 236,
	-1000, 321, -19, 139, -1000, -1000, 939, 345, 346, 275,
	-1000, -1000, 299, 138, -1000, -1000, 254, -1000, 328, 235,
	419, 234, 321, 1385, -32, -1000, 233, 1778, 231, 321,
	1385, -32, 229, 321, 5, 1759, 1741, -1000, -1000, -1000,
	-1000, 223, -1000, 220, -5, 134, 1385, 1385, 150, -1000,
	-1000, -1000, -1000, 1420, 40, 375, 218, -30, 1420, 214,
	213, -1000, 1385, 17, -31, -1000, -1000, 1474, -1000, -1000,
	1796, -1000, -1000, -1000, -1000, -1000, 1778, 211, -1000, 15,
	-1000, 976, -6, -1000, -1000, -1000, -1000, 1778, 12, 269,
	345, 321, -1000, -1000, 197, 328, 299, 345, 328, 321,
	10, 272, -1000, 1741, 195, -1000, -1000, -1000, -1000, -32,
	-1000, -1000, 132, -1000, -1000, 591, -32, -1000, -1000, -1000,
	251, -1000, -1000, 1741, -1000, -1000, -1000, -1000, -1000, -1000,
	1759, 1759, -1000, 1385, 1603, 1603, -1000, 976, 69, -1000,
	-1000, -1000, 137, -1000, 194, -1000, -1000, -1000, -27, -1000,
	859, -1000, 1125, 859, 859, 191, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 345, 187, -1000, 184, -1000,
	-1000, -1000, 183, -1000, 181, 321, 75, 1741, 180, -1000,
	-1000, 131, -1000, 1385, 1385, 1420, 1385, -1000, -1000, -1000,
	1385, -1000, -1000, -1000, 1603, -1000, 9, 2, -1000, -1000,
	345, 345, 1759, -1000, -1000, 135, -1000, 1354, 708, 392,
	-26, 859, -1000, -1000, -1000, -1000, -1000, 1759, -1000, -1000,
	-1000, -1000, -7, -1000, -1000,
}
var yyPgo = []int{

	0, 49, 7, 37, 28, 30, 0, 11, 571, 201,
	36, 570, 18, 568, 322, 32, 178, 607, 448, 3,
	39, 1, 567, 25, 455, 565, 420, 17, 9, 19,
	564, 562, 51, 563, 52, 446, 561, 4, 503, 26,
	559, 557, 556, 555, 554, 29, 553, 21, 551, 8,
	12, 549, 547, 546, 544, 538, 22, 536, 534, 533,
	532, 15, 24, 531, 527, 27, 5, 525, 524, 35,
	519, 516, 16, 513, 510, 508, 507, 498, 493, 33,
	490, 489, 91, 6, 2, 481, 13, 472, 10, 38,
	443, 461, 31, 14, 454, 450, 447, 445, 20, 432,
	430, 427,
}
var yyR1 = []int{

	0, 52, 51, 22, 22, 53, 53, 55, 55, 55,
	10, 10, 10, 40, 40, 57, 57, 57, 57, 57,
	35, 35, 35, 35, 35, 35, 35, 35, 35, 35,
	59, 45, 45, 45, 1, 1, 2, 2, 2, 33,
	32, 28, 28, 28, 28, 28, 28, 60, 60, 60,
	60, 61, 64, 63, 65, 65, 66, 67, 67, 68,
	68, 68, 70, 71, 72, 72, 73, 76, 74, 74,
	75, 75, 77, 78, 6, 6, 6, 6, 6, 6,
	6, 6, 6, 6, 6, 6, 6, 6, 6, 6,
	6, 6, 6, 6, 6, 34, 34, 34, 34, 34,
	34, 34, 34, 34, 26, 26, 26, 25, 25, 25,
	25, 25, 25, 25, 25, 25, 25, 25, 25, 25,
	25, 83, 86, 88, 88, 87, 87, 24, 24, 7,
	7, 15, 82, 82, 16, 3, 90, 90, 31, 14,
	91, 92, 92, 17, 17, 17, 17, 17, 17, 89,
	89, 89, 89, 94, 94, 94, 94, 94, 80, 80,
	81, 8, 8, 8, 8, 8, 4, 4, 21, 21,
	21, 21, 21, 21, 21, 27, 93, 30, 30, 13,
	13, 58, 95, 95, 9, 96, 96, 49, 49, 49,
	97, 85, 85, 54, 54, 46, 46, 36, 36, 44,
	44, 43, 43, 41, 41, 29, 29, 29, 29, 29,
	29, 23, 23, 5, 12, 12, 12, 11, 47, 47,
	47, 47, 48, 48, 50, 50, 98, 98, 98, 98,
	98, 18, 18, 18, 18, 18, 99, 18, 18, 18,
	18, 18, 18, 18, 18, 62, 62, 42, 42, 37,
	37, 38, 38, 39, 39, 101, 101, 101, 101, 84,
	84, 56, 56, 79, 79, 19, 19, 100, 100, 69,
	69, 20, 20,
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

	-1000, -51, -22, 41, -52, -31, 38, -53, 65, -54,
	-55, 30, -57, -35, -58, -18, 2, 49, -59, 48,
	23, -28, -71, -77, -78, -73, -91, 21, 8, 14,
	25, 18, 26, 43, 13, -6, -38, 22, 47, 45,
	28, -16, -34, -31, -24, 52, 60, 55, 51, 70,
	71, 57, 12, -25, 54, 34, -14, -26, -80, -81,
	-85, -9, -21, -97, 53, 10, 37, -30, -13, 46,
	32, 65, -10, 54, 34, -31, 50, 65, -45, 54,
	-37, -3, -31, -1, 54, -37, -32, 54, -33, -31,
	-95, 54, -31, 67, -90, -16, -31, -90, -26, -24,
	-31, 23, -26, -16, -100, -38, -6, 6, 31, 16,
	40, 4, 20, 39, 36, 33, 24, 27, 55, 51,
	56, 57, 52, 58, 59, 60, 5, 35, 44, 12,
	66, 11, 73, -70, -68, -69, -67, -28, -38, -72,
	-69, 7, -72, 54, 50, 53, -34, -34, -34, -34,
	-34, -34, -34, -34, 68, -7, -6, -89, -93, -9,
	-21, 52, 12, 54, -82, 7, 68, -82, 2, -19,
	15, -6, -94, 12, -9, -21, -27, -4, 54, 52,
	-14, 53, -82, -82, -40, 63, -10, 34, 34, -46,
	63, -45, -17, 66, 73, -93, -9, -21, -27, -4,
	54, 12, -1, 63, -17, 66, -44, 63, -32, -17,
	-96, 68, -50, -48, -47, -15, -31, -92, -17, 15,
	54, -99, 54, -6, -6, -6, -6, -6, -6, -6,
	-6, -6, -6, -6, -6, -6, -6, -6, -6, -6,
	-6, -6, -6, -6, -38, -38, -6, -66, 7, 65,
	66, 11, 7, 65, -65, -66, 63, -39, -7, -31,
	54, -6, -19, -83, 63, -89, 10, -6, -83, -62,
	-98, -61, -35, -18, 2, 68, 72, 72, -17, -17,
	-17, 50, -17, -43, 69, -29, -42, -5, 54, 52,
	-16, -23, 38, -41, 69, -12, -16, -23, 54, -56,
	65, -56, 65, 66, -38, -3, -17, 10, -56, 65,
	66, -38, -56, 65, -62, 63, 73, -79, -15, -92,
	-17, -50, -98, -50, -62, -69, 42, 42, -65, -69,
	-63, 69, -60, 9, 17, -74, -79, 15, 73, -7,
	48, 72, 67, -84, -101, -86, -88, -6, 68, 68,
	12, -17, -9, -21, -27, -4, 54, -79, 73, -84,
	69, 65, -62, -17, -17, 63, -31, 72, -56, 65,
	-17, 73, -20, 34, -5, 52, 38, -5, 54, 50,
	-56, 65, -11, 54, -23, 63, -10, 63, -45, -38,
	63, 63, -36, -2, -1, -37, -38, 63, -32, 69,
	-31, -49, -8, 54, -93, -9, -21, -27, -4, -47,
	63, 63, 69, 65, -6, -6, 69, -64, -39, 67,
	-75, -76, 19, 63, -79, -7, 63, 63, -19, 69,
	73, -79, 67, -83, -83, -17, 63, 69, -98, 69,
	-17, 69, -29, -20, -16, 63, -5, -20, -5, -31,
	69, -12, -50, 63, -56, 65, -17, 54, -50, -49,
	-69, -62, 67, 66, 11, 73, 28, -61, 63, 72,
	67, -86, -88, -87, -6, 68, -84, -84, 63, -20,
	63, 63, 63, 63, -2, -50, 63, -6, -6, -72,
	-19, -83, 69, 69, -20, -20, -49, 63, 67, 67,
	-66, 72, -84, -49, 69,
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
	3, 3, 3, 70, 3, 3, 3, 59, 60, 3,
	54, 63, 52, 55, 73, 51, 50, 58, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 67, 65,
	3, 66, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 53, 3, 72, 57, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 68, 56, 69, 71,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 61, 62,
	64,
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
			panic(".y:125")
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
			panic(".y:147")
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
			panic(".y:241")
		}
	case 43:
		{ //244
			panic(".y:245")
		}
	case 44:
		{ //248
			panic(".y:249")
		}
	case 45:
		{ //252
			panic(".y:253")
		}
	case 46:
		{ //256
			panic(".y:257")
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
			panic(".y:281")
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
			panic(".y:306")
		}
	case 57:
		{ //311
			panic(".y:312")
		}
	case 58:
		{ //315
			panic(".y:316")
		}
	case 59:
		{ //321
			panic(".y:322")
		}
	case 60:
		{ //325
			panic(".y:326")
		}
	case 61:
		{ //329
			panic(".y:330")
		}
	case 62:
		{ //335
			panic(".y:336")
		}
	case 63:
		{ //341
			panic(".y:342")
		}
	case 64:
		{ //347
			panic(".y:348")
		}
	case 65:
		{ //351
			panic(".y:352")
		}
	case 66:
		{ //357
			panic(".y:358")
		}
	case 67:
		{ //363
			panic(".y:364")
		}
	case 68:
		{ //368
			panic(".y:369")
		}
	case 69:
		{ //372
			panic(".y:373")
		}
	case 70:
		{ //377
			panic(".y:378")
		}
	case 71:
		{ //381
			panic(".y:382")
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
			panic(".y:412")
		}
	case 78:
		{ //415
			panic(".y:416")
		}
	case 79:
		{ //419
			panic(".y:420")
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
			panic(".y:432")
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
			panic(".y:490")
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
			panic(".y:524")
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
			panic(".y:546")
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
			panic(".y:578")
		}
	case 118:
		{ //581
			panic(".y:582")
		}
	case 119:
		{ //585
			panic(".y:586")
		}
	case 120:
		{ //589
			panic(".y:590")
		}
	case 121:
		{ //594
			panic(".y:595")
		}
	case 122:
		{ //600
			panic(".y:601")
		}
	case 123:
		{ //606
			panic(".y:607")
		}
	case 124:
		{ //610
			panic(".y:611")
		}
	case 125:
		{ //616
			panic(".y:617")
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
	case 132:
		{ //652
			panic(".y:653")
		}
	case 134:
		yyVAL.node = yyS[yypt-0].node
	case 135:
		yyVAL.node = yyS[yypt-0].node
	case 136:
		{ //673
			panic(".y:674")
		}
	case 137:
		{ //677
			panic(".y:678")
		}
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
			panic(".y:702")
			//yy(yylex).errPos($1.tpos, "final argument in variadic function missing type")
		}
	case 142:
		{ //705
			panic(".y:706")
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
		{ //787
			panic(".y:788")
		}
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
			switch { //TODO + resolve scope
			case yyS[yypt-2].node != nil:
				yyVAL.node = &ArrayType{yyS[yypt-3].token.pos, yyS[yypt-2].node, yyS[yypt-0].node}
			default:
				yyVAL.node = &SliceType{yyS[yypt-3].token.pos, yyS[yypt-0].node}
			}
		}
	case 169:
		{ //829
			panic(".y:830")
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
			panic(".y:888")
		}
	case 182:
		{ //893
			panic(".y:894")
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
			panic(".y:909")
		}
	case 186:
		{ //912
			panic(".y:913")
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
		{ //932
			panic(".y:933")
		}
	case 191:
		{ //938
			panic(".y:939")
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
			panic(".y:1078")
		}
	case 221:
		{ //1081
			panic(".y:1082")
		}
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
			panic(".y:1106")
		}
	case 227:
		{ //1109
			panic(".y:1110")
		}
	case 228:
		{ //1113
			panic(".y:1114")
		}
	case 229:
		{ //1117
			panic(".y:1118")
		}
	case 230:
		{ //1121
			panic(".y:1122")
		}
	case 231:
		yyVAL.node = yyS[yypt-0].node
	case 232:
		{ //1131
			panic(".y:1132")
		}
	case 233:
		{ //1135
			panic(".y:1136")
		}
	case 234:
		{ //1139
			panic(".y:1140")
		}
	case 235:
		{ //1143
			panic(".y:1144")
		}
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
			panic(".y:1156")
		}
	case 239:
		{ //1159
			panic(".y:1160")
		}
	case 240:
		{ //1163
			panic(".y:1164")
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
			panic(".y:1176")
		}
	case 244:
		{ //1179
			panic(".y:1180")
		}
	case 245:
		{ //1185
			panic(".y:1186")
		}
	case 246:
		{ //1189
			panic(".y:1190")
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
			panic(".y:1236")
		}
	case 256:
		{ //1239
			panic(".y:1240")
		}
	case 257:
		{ //1243
			panic(".y:1244")
		}
	case 258:
		{ //1247
			panic(".y:1248")
		}
	case 259:
		{ //1252
			panic(".y:1253")
		}
	case 260:
		{ //1256
			panic(".y:1257")
		}
	case 264:
		{ //1274
			panic(".y:1275")
		}
	case 265:
		{ //1279
			yyVAL.node = nil
		}
	case 266:
		yyVAL.node = yyS[yypt-0].node
	case 267:
		{ //1288
			panic(".y:1289")
		}
	case 268:
		{ //1292
			panic(".y:1293")
		}
	case 269:
		{ //1297
			panic(".y:1298")
		}
	case 270:
		{ //1301
			panic(".y:1302")
		}
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
