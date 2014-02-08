// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parser

import __yyfmt__ "fmt"

type yySymType struct {
	yys   int
	token tkn
	node  Node
	list  []Node
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
	" +",
	" -",
	" |",
	" ^",
	" *",
	" /",
	" %",
	" &",
	"notPackage",
	"notParen",
	" (",
	" )",
	"preferToRightParen",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

func yy(y yyLexer) *parser { return y.(*parser) }

func yyError(y yyLexer, msg string) {
	yy(y).Error(msg)
}

func yyTLD(y yyLexer, n Node) {
	p := yy(y)
	p.ast = append(p.ast, n)
}

func yyTLDs(y yyLexer, l []Node) {
	p := yy(y)
	p.ast = append(p.ast, l...)
}

var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 9,
	1, 2,
	63, 15,
	-2, 0,
	-1, 35,
	11, 251,
	65, 251,
	73, 251,
	-2, 41,
	-1, 43,
	66, 134,
	-2, 139,
	-1, 61,
	60, 158,
	-2, 190,
	-1, 62,
	60, 159,
	-2, 160,
	-1, 98,
	60, 115,
	64, 115,
	67, 115,
	71, 115,
	-2, 241,
	-1, 102,
	60, 115,
	64, 115,
	67, 115,
	71, 115,
	-2, 242,
	-1, 159,
	2, 190,
	7, 190,
	60, 158,
	67, 190,
	-2, 150,
	-1, 160,
	7, 160,
	60, 159,
	67, 160,
	-2, 151,
	-1, 167,
	63, 226,
	68, 226,
	-2, 0,
	-1, 211,
	63, 226,
	68, 226,
	-2, 0,
	-1, 221,
	9, 226,
	17, 226,
	63, 226,
	68, 226,
	-2, 0,
	-1, 248,
	63, 226,
	68, 226,
	-2, 0,
	-1, 275,
	63, 226,
	68, 226,
	-2, 0,
	-1, 292,
	34, 211,
	63, 211,
	68, 211,
	-2, 138,
	-1, 352,
	7, 153,
	60, 153,
	67, 153,
	-2, 144,
	-1, 353,
	7, 154,
	60, 154,
	67, 154,
	-2, 145,
	-1, 354,
	7, 155,
	60, 155,
	67, 155,
	-2, 146,
	-1, 355,
	7, 156,
	60, 156,
	67, 156,
	-2, 147,
	-1, 361,
	9, 226,
	17, 226,
	63, 226,
	68, 226,
	-2, 0,
	-1, 417,
	9, 226,
	17, 226,
	63, 226,
	68, 226,
	-2, 0,
	-1, 478,
	7, 157,
	60, 157,
	67, 157,
	-2, 148,
}

const yyNprod = 273
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1974

var yyAct = []int{

	35, 62, 343, 169, 80, 247, 263, 393, 401, 21,
	346, 258, 212, 345, 195, 271, 139, 199, 295, 285,
	270, 214, 299, 85, 269, 291, 257, 254, 198, 78,
	287, 217, 215, 317, 106, 140, 72, 81, 157, 372,
	470, 132, 430, 358, 316, 501, 469, 137, 137, 394,
	137, 86, 42, 367, 277, 156, 160, 142, 276, 143,
	181, 504, 292, 144, 493, 171, 155, 175, 83, 158,
	145, 361, 361, 135, 361, 361, 439, 412, 337, 399,
	360, 492, 197, 177, 298, 450, 441, 197, 251, 85,
	197, 437, 294, 197, 176, 429, 349, 211, 146, 147,
	148, 149, 150, 151, 152, 153, 154, 464, 223, 191,
	186, 224, 225, 226, 227, 228, 229, 230, 231, 232,
	233, 234, 235, 236, 237, 238, 239, 240, 241, 242,
	243, 106, 106, 246, 202, 131, 338, 292, 65, 208,
	201, 419, 250, 219, 156, 160, 261, 466, 255, 262,
	132, 101, 168, 289, 342, 93, 164, 165, 158, 288,
	70, 463, 462, 160, 267, 66, 6, 284, 165, 465,
	310, 268, 303, 74, 69, 197, 158, 6, 41, 333,
	197, 197, 179, 197, 333, 379, 275, 334, 200, 130,
	281, 457, 334, 361, 106, 455, 413, 132, 381, 64,
	265, 61, 197, 76, 369, 313, 106, 95, 95, 297,
	309, 103, 301, 497, 146, 153, 302, 166, 197, 300,
	167, 197, 197, 253, 197, 308, 182, 183, 166, 312,
	74, 249, 305, 321, 6, 323, 314, 77, 416, 71,
	74, 8, 322, 331, 6, 486, 483, 126, 319, 318,
	482, 106, 106, 6, 481, 480, 159, 185, 6, 137,
	76, 156, 160, 137, 347, 478, 73, 174, 353, 347,
	76, 359, 339, 324, 468, 158, 207, 127, 197, 197,
	328, 203, 196, 453, 355, 325, 128, 196, 197, 329,
	196, 336, 6, 196, 445, 354, 122, 123, 124, 125,
	362, 357, 436, 292, 106, 427, 368, 292, 426, 197,
	423, 106, 411, 410, 395, 190, 380, 406, 197, 374,
	377, 397, 56, 289, 384, 298, 383, 414, 415, 288,
	404, 391, 388, 408, 156, 160, 390, 386, 409, 156,
	160, 387, 385, 171, 407, 159, 428, 365, 158, 315,
	425, 376, 197, 158, 264, 433, 434, 6, 197, 222,
	418, 290, 296, 159, 6, 398, 220, 163, 6, 197,
	6, 424, 6, 378, 3, 196, 6, 376, 431, 260,
	196, 196, 438, 196, 376, 197, 91, 373, 180, 442,
	87, 188, 84, 375, 79, 187, 452, 197, 11, 422,
	451, 307, 196, 180, 248, 197, 446, 297, 180, 448,
	443, 180, 406, 406, 180, 454, 458, 447, 196, 459,
	57, 196, 196, 137, 196, 404, 404, 94, 408, 408,
	252, 347, 141, 474, 347, 347, 476, 477, 467, 407,
	407, 472, 461, 273, 471, 344, 272, 104, 221, 460,
	213, 98, 102, 15, 382, 44, 13, 97, 286, 197,
	395, 63, 159, 484, 487, 488, 156, 160, 352, 210,
	485, 171, 90, 293, 490, 283, 137, 425, 196, 196,
	158, 68, 491, 489, 406, 479, 99, 99, 196, 67,
	402, 496, 347, 172, 502, 500, 180, 404, 26, 406,
	408, 180, 180, 36, 180, 473, 503, 60, 59, 196,
	58, 407, 404, 53, 24, 408, 23, 405, 196, 421,
	494, 495, 420, 180, 335, 25, 407, 22, 133, 134,
	136, 417, 330, 332, 88, 159, 206, 105, 392, 180,
	159, 138, 180, 180, 18, 180, 189, 14, 290, 12,
	444, 184, 196, 10, 9, 7, 4, 1, 196, 2,
	296, 0, 100, 0, 0, 0, 5, 0, 0, 196,
	0, 0, 43, 0, 75, 0, 0, 0, 0, 0,
	82, 82, 89, 92, 0, 196, 0, 0, 0, 180,
	0, 96, 96, 0, 0, 96, 0, 196, 0, 180,
	180, 65, 0, 201, 0, 196, 0, 218, 0, 180,
	0, 0, 405, 405, 101, 0, 0, 0, 0, 0,
	0, 0, 0, 70, 0, 0, 0, 0, 66, 6,
	180, 0, 0, 65, 244, 245, 75, 69, 180, 180,
	0, 0, 82, 0, 0, 179, 101, 82, 0, 0,
	89, 200, 0, 0, 216, 70, 205, 55, 0, 196,
	66, 6, 64, 0, 194, 0, 0, 159, 0, 69,
	0, 0, 0, 180, 0, 0, 0, 0, 0, 180,
	0, 0, 0, 54, 405, 0, 0, 0, 192, 0,
	180, 0, 0, 204, 64, 0, 209, 304, 0, 405,
	65, 0, 162, 0, 0, 0, 180, 259, 0, 311,
	0, 0, 0, 101, 0, 0, 0, 0, 180, 126,
	0, 0, 70, 0, 55, 0, 180, 66, 6, 0,
	43, 0, 0, 180, 180, 0, 69, 0, 340, 0,
	47, 48, 0, 51, 161, 96, 96, 46, 0, 127,
	54, 0, 0, 0, 244, 245, 0, 82, 128, 49,
	50, 64, 0, 0, 118, 119, 120, 121, 122, 123,
	124, 125, 0, 0, 43, 0, 0, 0, 0, 0,
	180, 278, 0, 216, 43, 216, 279, 280, 0, 282,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 111, 126, 0, 0, 180, 0, 389, 306, 129,
	0, 43, 0, 0, 396, 0, 0, 112, 0, 0,
	180, 116, 0, 0, 117, 0, 0, 320, 0, 0,
	115, 0, 127, 114, 0, 0, 113, 110, 43, 0,
	0, 128, 0, 0, 366, 0, 0, 118, 119, 120,
	121, 122, 123, 124, 125, 0, 0, 0, 0, 0,
	0, 0, 0, 75, 0, 82, 0, 0, 0, 65,
	358, 52, 82, 0, 351, 0, 89, 0, 400, 216,
	0, 0, 101, 0, 363, 364, 0, 0, 0, 0,
	0, 70, 0, 55, 370, 0, 66, 6, 0, 0,
	0, 0, 0, 0, 0, 69, 0, 0, 0, 47,
	48, 0, 51, 45, 0, 351, 46, 0, 0, 54,
	0, 0, 0, 0, 43, 0, 348, 0, 49, 50,
	64, 0, 96, 0, 96, 65, 0, 201, 0, 0,
	0, 0, 449, 0, 96, 0, 216, 0, 101, 0,
	0, 0, 0, 0, 0, 0, 0, 70, 278, 0,
	0, 0, 66, 6, 435, 0, 216, 0, 0, 0,
	0, 69, 0, 0, 0, 440, 0, 0, 274, 179,
	43, 0, 0, 0, 28, 403, 65, 0, 52, 34,
	29, 0, 0, 0, 31, 0, 64, 27, 37, 101,
	0, 30, 32, 456, 40, 0, 0, 0, 70, 0,
	55, 0, 0, 66, 6, 0, 0, 0, 82, 33,
	216, 39, 69, 38, 19, 17, 47, 48, 0, 51,
	45, 0, 0, 46, 0, 0, 54, 0, 0, 0,
	0, 0, 16, 275, 0, 49, 50, 64, 28, 0,
	65, 0, 52, 34, 29, 0, 0, 0, 31, 0,
	0, 27, 37, 20, 0, 30, 32, 0, 40, 0,
	0, 0, 70, 0, 55, 0, 0, 66, 6, 0,
	0, 0, 0, 33, 0, 39, 69, 38, 19, 17,
	47, 48, 0, 51, 45, 0, 0, 46, 111, 126,
	54, 0, 0, 0, 0, 0, 129, 0, 0, 49,
	50, 64, 0, 0, 112, 0, 0, 0, 116, 0,
	0, 117, 0, 0, 0, 0, 0, 115, 0, 127,
	114, 0, 0, 113, 110, 0, 0, 0, 128, 0,
	65, 0, 52, 0, 118, 119, 120, 121, 122, 123,
	124, 125, 0, 101, 0, 0, 0, 0, 0, 0,
	0, 0, 70, 0, 55, 0, 341, 66, 6, 0,
	0, 0, 0, 0, 0, 0, 69, 0, 0, 0,
	47, 48, 0, 51, 45, 0, 0, 46, 0, 65,
	54, 52, 0, 0, 0, 0, 0, 475, 0, 49,
	50, 64, 101, 0, 0, 0, 0, 0, 0, 0,
	0, 70, 0, 55, 0, 0, 66, 6, 0, 0,
	0, 327, 0, 0, 0, 69, 0, 0, 0, 47,
	48, 0, 51, 45, 0, 0, 46, 0, 65, 54,
	52, 0, 0, 0, 0, 0, 0, 0, 49, 50,
	64, 101, 0, 0, 0, 0, 0, 0, 0, 0,
	70, 0, 55, 0, 0, 66, 6, 65, 0, 162,
	326, 0, 0, 0, 69, 0, 0, 0, 47, 48,
	101, 51, 45, 0, 0, 46, 0, 0, 54, 70,
	0, 55, 0, 0, 66, 6, 0, 49, 50, 64,
	0, 0, 0, 69, 0, 0, 0, 47, 48, 0,
	51, 161, 0, 0, 46, 0, 0, 54, 256, 65,
	0, 52, 0, 0, 170, 0, 49, 50, 64, 0,
	0, 0, 101, 0, 0, 0, 0, 0, 0, 0,
	0, 70, 0, 55, 0, 0, 66, 6, 65, 0,
	52, 0, 0, 0, 0, 69, 0, 0, 0, 47,
	48, 101, 51, 45, 0, 0, 46, 0, 0, 54,
	70, 0, 55, 0, 0, 66, 6, 0, 49, 50,
	64, 0, 0, 0, 69, 0, 0, 0, 47, 48,
	0, 51, 45, 0, 0, 46, 0, 65, 54, 162,
	0, 0, 0, 0, 0, 0, 0, 49, 50, 64,
	101, 0, 0, 0, 0, 0, 0, 0, 0, 70,
	0, 55, 0, 0, 66, 6, 266, 0, 52, 0,
	0, 0, 0, 69, 0, 0, 0, 47, 48, 101,
	51, 161, 0, 0, 46, 0, 0, 54, 70, 0,
	55, 0, 0, 66, 6, 0, 49, 50, 64, 0,
	0, 0, 69, 0, 0, 0, 47, 48, 0, 51,
	45, 0, 0, 46, 111, 126, 54, 0, 0, 0,
	0, 0, 129, 0, 0, 49, 50, 64, 0, 0,
	112, 0, 0, 0, 116, 0, 0, 117, 0, 0,
	0, 0, 0, 115, 0, 127, 114, 0, 0, 113,
	110, 0, 0, 0, 128, 0, 0, 0, 0, 0,
	118, 119, 120, 121, 122, 123, 124, 125, 0, 111,
	126, 0, 0, 0, 0, 0, 499, 129, 0, 0,
	0, 0, 0, 0, 0, 112, 0, 0, 0, 116,
	0, 0, 117, 0, 0, 0, 0, 0, 115, 0,
	127, 114, 0, 0, 113, 110, 0, 0, 0, 128,
	0, 0, 0, 0, 0, 118, 119, 120, 121, 122,
	123, 124, 125, 0, 111, 126, 0, 0, 0, 0,
	0, 498, 129, 0, 0, 0, 0, 0, 0, 0,
	112, 0, 0, 0, 116, 0, 0, 117, 0, 0,
	0, 0, 0, 115, 0, 127, 114, 0, 0, 113,
	110, 0, 0, 0, 128, 65, 0, 201, 0, 0,
	118, 119, 120, 121, 122, 123, 124, 125, 101, 0,
	0, 0, 0, 65, 0, 201, 432, 70, 0, 0,
	0, 0, 66, 6, 0, 65, 101, 201, 0, 0,
	0, 69, 0, 0, 0, 70, 0, 0, 101, 179,
	66, 6, 0, 0, 0, 200, 0, 70, 0, 69,
	193, 0, 66, 6, 0, 0, 64, 179, 194, 0,
	266, 69, 201, 200, 0, 0, 0, 0, 0, 179,
	65, 0, 350, 101, 64, 200, 371, 0, 0, 0,
	0, 0, 70, 101, 0, 0, 64, 66, 6, 65,
	0, 173, 70, 0, 0, 0, 69, 66, 6, 0,
	0, 0, 101, 0, 179, 0, 69, 0, 0, 0,
	200, 70, 0, 0, 179, 0, 66, 6, 0, 0,
	356, 64, 111, 126, 107, 69, 0, 0, 0, 0,
	129, 64, 0, 179, 109, 0, 0, 0, 112, 178,
	0, 0, 116, 0, 0, 117, 0, 0, 0, 108,
	64, 115, 0, 127, 114, 0, 0, 113, 110, 0,
	0, 0, 128, 0, 0, 0, 0, 0, 118, 119,
	120, 121, 122, 123, 124, 125, 111, 126, 0, 0,
	0, 0, 0, 0, 129, 0, 0, 0, 0, 0,
	0, 0, 112, 0, 0, 0, 116, 0, 0, 117,
	0, 0, 0, 0, 0, 115, 0, 127, 114, 0,
	0, 113, 110, 0, 111, 126, 128, 0, 0, 0,
	0, 0, 118, 119, 120, 121, 122, 123, 124, 125,
	112, 0, 0, 0, 116, 0, 0, 117, 0, 0,
	0, 0, 0, 115, 0, 127, 114, 0, 0, 113,
	110, 0, 111, 126, 128, 0, 0, 0, 0, 0,
	118, 119, 120, 121, 122, 123, 124, 125, 112, 0,
	0, 0, 116, 0, 0, 117, 0, 0, 0, 0,
	0, 115, 0, 127, 114, 0, 0, 113, 0, 0,
	0, 126, 128, 0, 0, 0, 0, 0, 118, 119,
	120, 121, 122, 123, 124, 125, 112, 0, 0, 0,
	116, 0, 0, 117, 0, 0, 0, 0, 0, 115,
	0, 127, 114, 0, 0, 113, 0, 0, 0, 0,
	128, 0, 0, 0, 0, 0, 118, 119, 120, 121,
	122, 123, 124, 125,
}
var yyPact = []int{

	333, -1000, -1000, 338, -1000, 178, -1000, 368, -1000, 1040,
	176, 206, 174, -1000, -1000, -1000, -1000, 334, 332, 330,
	326, -1000, -1000, -1000, -1000, -1000, 89, -1000, 338, 338,
	623, 623, 338, 1338, -1000, 1748, 124, 1338, 1338, 425,
	1338, -1000, -1000, -1000, -1, 1338, 1338, 1338, 1338, 1338,
	1338, 1338, 1338, 39, 1387, -1000, -1000, -1000, 307, 161,
	-1000, -1000, -1000, 150, 1309, 1709, -11, -1000, -1000, 161,
	161, -1000, -1000, 196, -1000, 361, 357, -1000, -1000, 254,
	1615, -1000, -1000, -1000, 220, 591, -1000, 215, 1645, -1000,
	30, 128, 306, -1000, -1000, -1000, -1000, -1000, -1000, -1,
	-1000, 299, -1000, -1000, -1000, -32, 1802, 1338, -1000, -1000,
	1338, 1338, 1338, 1338, 1338, 1338, 1338, 1338, 1338, 1338,
	1338, 1338, 1338, 1338, 1338, 1338, 1338, 1338, 1338, 1338,
	1338, 1338, 1338, -1000, 397, 168, -1000, -1000, 77, 423,
	160, -1000, 397, 1257, 319, 1338, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 293, 1802, -1000, -1000, -1000,
	-1000, 1387, 1416, 1338, -1000, -1000, -1000, 976, -1000, -14,
	-18, 1802, -1000, 1645, -1000, -1000, -1000, -1000, 1645, 1645,
	126, 1645, 99, 24, 156, -1000, -1000, -1000, -1000, 153,
	-1000, -1000, 107, 1338, 338, -1000, -1000, -1000, -1000, -1000,
	1645, 391, 147, -1000, 105, 1338, 142, -1000, -1000, -1000,
	-1000, 976, 288, -29, -1000, -1000, 128, -1000, -1000, 1645,
	128, 976, 128, 1802, 1878, 1916, 714, 714, 714, 714,
	714, 714, 242, 242, 242, 242, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 1840, -32, -32, 1802, -1000, 976, 1338,
	1228, 1179, -1000, 1338, 175, -1000, -1000, 63, -1000, -1000,
	690, 1094, 88, 859, 29, -1000, 1690, 797, 859, 12,
	-1000, -1000, -1000, -1000, -1000, 976, 1645, 1645, -1000, 286,
	-1000, 338, -19, 141, -1000, -1000, 1633, 353, 339, 313,
	-1000, -1000, 121, 135, -1000, -1000, 266, -1000, 346, 281,
	139, 280, 338, 1338, -32, -1000, 275, 1645, 270, 338,
	1338, -32, 260, 338, 11, 925, 128, -1000, -1000, -1000,
	-1000, 252, -1000, 251, 9, 133, 1338, 1338, 170, -1000,
	-1000, -1000, -1000, 1387, 75, 380, 249, -30, 1387, 247,
	244, -1000, 1338, 27, -31, -1000, -1000, 1580, -1000, -1000,
	1680, -1000, -1000, -1000, -1000, -1000, 1645, 241, -1000, 23,
	-1000, 976, 8, -1000, -1000, -1000, -1000, 1645, 18, 269,
	353, 338, -1000, -1000, 233, 346, 121, 353, 346, 338,
	17, 265, -1000, 128, 222, -1000, -1000, -1000, -1000, -32,
	-1000, -1000, 132, -1000, -1000, 591, -32, -1000, -1000, -1000,
	131, -1000, -1000, 128, -1000, -1000, -1000, -1000, -1000, -1000,
	925, 925, -1000, 1338, 1802, 1802, -1000, 976, 96, -1000,
	-1000, -1000, 119, -1000, 213, -1000, -1000, -1000, -26, -1000,
	859, -1000, 1130, 859, 859, 204, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 353, 194, -1000, 193, -1000,
	-1000, -1000, 189, -1000, 185, 338, 105, 128, 184, -1000,
	-1000, 130, -1000, 1338, 1338, 1387, 1338, -1000, -1000, -1000,
	1338, -1000, -1000, -1000, 1802, -1000, 13, -4, -1000, -1000,
	353, 353, 925, -1000, -1000, 152, -1000, 1525, 1470, 397,
	-27, 859, -1000, -1000, -1000, -1000, -1000, 925, -1000, -1000,
	-1000, -1000, -7, -1000, -1000,
}
var yyPgo = []int{

	0, 559, 562, 446, 557, 556, 555, 554, 553, 36,
	551, 22, 549, 547, 443, 29, 546, 544, 49, 538,
	51, 536, 4, 607, 503, 7, 534, 9, 0, 533,
	26, 15, 24, 532, 531, 27, 5, 530, 529, 35,
	528, 527, 16, 525, 524, 522, 519, 516, 514, 52,
	455, 420, 33, 513, 322, 11, 3, 510, 508, 156,
	6, 2, 507, 13, 505, 10, 38, 32, 178, 37,
	427, 498, 31, 14, 201, 1, 28, 17, 493, 490,
	489, 481, 475, 473, 472, 469, 12, 8, 461, 19,
	18, 458, 39, 30, 25, 454, 21, 450, 20, 448,
	447, 445,
}
var yyR1 = []int{

	0, 5, 4, 1, 1, 6, 6, 8, 8, 8,
	9, 9, 9, 10, 10, 12, 12, 12, 12, 12,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	17, 15, 15, 15, 18, 18, 25, 25, 25, 26,
	20, 27, 27, 27, 27, 27, 27, 29, 29, 29,
	29, 31, 34, 33, 35, 35, 36, 37, 37, 38,
	38, 38, 40, 41, 42, 42, 43, 46, 44, 44,
	45, 45, 47, 48, 28, 28, 28, 28, 28, 28,
	28, 28, 28, 28, 28, 28, 28, 28, 28, 28,
	28, 28, 28, 28, 28, 49, 49, 49, 49, 49,
	49, 49, 49, 49, 51, 51, 51, 53, 53, 53,
	53, 53, 53, 53, 53, 53, 53, 53, 53, 53,
	53, 60, 63, 65, 65, 64, 64, 50, 50, 55,
	55, 67, 59, 59, 68, 69, 70, 70, 2, 54,
	71, 72, 72, 23, 23, 23, 23, 23, 23, 66,
	66, 66, 66, 78, 78, 78, 78, 78, 57, 57,
	58, 79, 79, 79, 79, 79, 77, 77, 75, 75,
	75, 75, 75, 75, 75, 76, 73, 80, 80, 81,
	81, 13, 84, 84, 74, 85, 85, 87, 87, 87,
	88, 62, 62, 7, 7, 16, 16, 19, 19, 21,
	21, 82, 82, 83, 83, 89, 89, 89, 89, 89,
	89, 94, 94, 93, 90, 90, 90, 95, 96, 96,
	96, 96, 97, 97, 86, 86, 98, 98, 98, 98,
	98, 14, 14, 14, 14, 14, 99, 14, 14, 14,
	14, 14, 14, 14, 14, 32, 32, 91, 91, 22,
	22, 24, 24, 30, 30, 101, 101, 101, 101, 61,
	61, 11, 11, 52, 52, 56, 56, 100, 100, 39,
	39, 92, 92,
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

	-1000, -4, -1, 41, -5, -2, 38, -6, 63, -7,
	-8, 30, -12, -3, -13, -14, 2, 49, -17, 48,
	23, -27, -41, -47, -48, -43, -71, 21, 8, 14,
	25, 18, 26, 43, 13, -28, -24, 22, 47, 45,
	28, -68, -49, -2, -50, 54, 57, 50, 51, 69,
	70, 53, 12, -53, 60, 34, -54, -51, -57, -58,
	-62, -74, -75, -88, 71, 10, 37, -80, -81, 46,
	32, 63, -9, 60, 34, -2, 64, 63, -15, 60,
	-22, -69, -2, -18, 60, -22, -20, 60, -26, -2,
	-84, 60, -2, 66, -70, -68, -2, -70, -51, -50,
	-2, 23, -51, -68, -100, -24, -28, 6, 31, 16,
	40, 4, 20, 39, 36, 33, 24, 27, 50, 51,
	52, 53, 54, 55, 56, 57, 5, 35, 44, 12,
	65, 11, 73, -40, -38, -39, -37, -27, -24, -42,
	-39, 7, -42, 60, 64, 71, -49, -49, -49, -49,
	-49, -49, -49, -49, 67, -55, -28, -66, -73, -74,
	-75, 54, 12, 60, -59, 7, 67, -59, 2, -56,
	15, -28, -78, 12, -74, -75, -76, -77, 60, 54,
	-54, 71, -59, -59, -10, 61, -9, 34, 34, -16,
	61, -15, -23, 65, 73, -73, -74, -75, -76, -77,
	60, 12, -18, 61, -23, 65, -21, 61, -20, -23,
	-85, 67, -86, -97, -96, -67, -2, -72, -23, 15,
	60, -99, 60, -28, -28, -28, -28, -28, -28, -28,
	-28, -28, -28, -28, -28, -28, -28, -28, -28, -28,
	-28, -28, -28, -28, -24, -24, -28, -36, 7, 63,
	65, 11, 7, 63, -35, -36, 61, -30, -55, -2,
	60, -28, -56, -60, 61, -66, 10, -28, -60, -32,
	-98, -31, -3, -14, 2, 67, 72, 72, -23, -23,
	-23, 64, -23, -82, 68, -89, -91, -93, 60, 54,
	-68, -94, 38, -83, 68, -90, -68, -94, 60, -11,
	63, -11, 63, 65, -24, -69, -23, 10, -11, 63,
	65, -24, -11, 63, -32, 61, 73, -52, -67, -72,
	-23, -86, -98, -86, -32, -39, 42, 42, -35, -39,
	-33, 68, -29, 9, 17, -44, -52, 15, 73, -55,
	48, 72, 66, -61, -101, -63, -65, -28, 67, 67,
	12, -23, -74, -75, -76, -77, 60, -52, 73, -61,
	68, 63, -32, -23, -23, 61, -2, 72, -11, 63,
	-23, 73, -92, 34, -93, 54, 38, -93, 60, 64,
	-11, 63, -95, 60, -94, 61, -9, 61, -15, -24,
	61, 61, -19, -25, -18, -22, -24, 61, -20, 68,
	-2, -87, -79, 60, -73, -74, -75, -76, -77, -96,
	61, 61, 68, 63, -28, -28, 68, -34, -30, 66,
	-45, -46, 19, 61, -52, -55, 61, 61, -56, 68,
	73, -52, 66, -60, -60, -23, 61, 68, -98, 68,
	-23, 68, -89, -92, -68, 61, -93, -92, -93, -2,
	68, -90, -86, 61, -11, 63, -23, 60, -86, -87,
	-39, -32, 66, 65, 11, 73, 28, -31, 61, 72,
	66, -63, -65, -64, -28, 67, -61, -61, 61, -92,
	61, 61, 61, 61, -25, -86, 61, -28, -28, -42,
	-56, -60, 68, 68, -92, -92, -87, 61, 66, 66,
	-36, 72, -61, -87, 68,
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
	3, 3, 3, 69, 3, 3, 3, 56, 57, 3,
	60, 61, 54, 50, 73, 51, 64, 55, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 66, 63,
	3, 65, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 71, 3, 72, 53, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 67, 52, 68, 70,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 58, 59,
	62,
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
			yyError(yylex, "package statement must be first")
			goto ret1
		}
	case 4:
		{ //64
			yyVAL.node = &Package{yyS[yypt-2].token.pos, yyS[yypt-1].node.(*Ident)}
		}
	case 6:
		{ //73
			panic(".y:74")
		}
	case 7:
		{ //79
			panic(".y:80")
		}
	case 8:
		{ //83
			panic(".y:84")
		}
	case 9:
		{ //87
			panic(".y:88")
		}
	case 10:
		{ //93
			panic(".y:94")
		}
	case 11:
		{ //97
			panic(".y:98")
		}
	case 12:
		{ //101
			panic(".y:102")
		}
	case 13:
		{ //107
			panic(".y:108")
		}
	case 14:
		{ //111
			panic(".y:112")
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
			panic(".y:129")
		}
	case 19:
		{ //132
			panic(".y:133")
		}
	case 20:
		{ //138
			panic(".y:139")
		}
	case 21:
		{ //142
			panic(".y:143")
		}
	case 22:
		{ //146
			panic(".y:147")
		}
	case 23:
		{ //150
			panic(".y:151")
		}
	case 24:
		{ //154
			panic(".y:155")
		}
	case 25:
		{ //158
			panic(".y:159")
		}
	case 26:
		{ //162
			yyVAL.list = nil
		}
	case 27:
		{ //166
			panic(".y:167")
		}
	case 28:
		{ //170
			panic(".y:171")
		}
	case 29:
		{ //174
			panic(".y:175")
		}
	case 30:
		{ //180
			p := yy(yylex)
			p.constExpr, p.constIota, p.constType = nil, 0, nil
		}
	case 31:
		{ //186
			panic(".y:187")
		}
	case 32:
		{ //190
			panic(".y:191")
		}
	case 33:
		{ //194
			panic(".y:195")
		}
	case 34:
		{ //200
			panic(".y:201")
		}
	case 35:
		{ //204
			panic(".y:205")
		}
	case 36:
		{ //210
			panic(".y:211")
		}
	case 37:
		{ //214
			panic(".y:215")
		}
	case 38:
		{ //218
			panic(".y:219")
		}
	case 39:
		{ //224
			panic(".y:225")
		}
	case 40:
		{ //230
			panic(".y:231")
		}
	case 41:
		{ //236
			panic(".y:237")
		}
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
		{ //399
			panic(".y:400")
		}
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
			panic(".y:436")
		}
	case 84:
		{ //439
			panic(".y:440")
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
			panic(".y:452")
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
			panic(".y:472")
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
		{ //485
			panic(".y:486")
		}
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
			panic(".y:502")
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
			panic(".y:528")
		}
	case 106:
		{ //531
			panic(".y:532")
		}
	case 107:
		{ //537
			panic(".y:538")
		}
	case 108:
		{ //541
			panic(".y:542")
		}
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
			panic(".y:558")
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
		{ //569
			panic(".y:570")
		}
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
		{ //626
			panic(".y:627")
		}
	case 128:
		{ //630
			panic(".y:631")
		}
	case 129:
		{ //636
			panic(".y:637")
		}
	case 130:
		{ //640
			panic(".y:641")
		}
	case 131:
		{ //646
			panic(".y:647")
		}
	case 132:
		{ //652
			panic(".y:653")
		}
	case 133:
		{ //656
			panic(".y:657")
		}
	case 134:
		{ //662
			panic(".y:663")
		}
	case 135:
		{ //668
			panic(".y:669")
		}
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
		{ //689
			panic(".y:690")
		}
	case 140:
		{ //695
			panic(".y:696")
		}
	case 141:
		{ //701
			panic(".y:702")
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
		{ //715
			panic(".y:716")
		}
	case 145:
		{ //719
			panic(".y:720")
		}
	case 146:
		{ //723
			panic(".y:724")
		}
	case 147:
		{ //727
			panic(".y:728")
		}
	case 148:
		{ //731
			panic(".y:732")
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
		{ //809
			panic(".y:810")
		}
	case 166:
		{ //815
			panic(".y:816")
		}
	case 167:
		{ //819
			panic(".y:820")
		}
	case 168:
		{ //825
			panic(".y:826")
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
		{ //845
			panic(".y:846")
		}
	case 174:
		{ //849
			panic(".y:850")
		}
	case 175:
		{ //855
			panic(".y:856")
		}
	case 176:
		{ //861
			panic(".y:862")
		}
	case 177:
		{ //867
			panic(".y:868")
		}
	case 178:
		{ //871
			panic(".y:872")
		}
	case 179:
		{ //877
			panic(".y:878")
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
			panic(".y:904")
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
			panic(".y:919")
		}
	case 188:
		{ //922
			panic(".y:923")
		}
	case 189:
		{ //926
			panic(".y:927")
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
			panic(".y:958")
		}
	case 196:
		{ //961
			panic(".y:962")
		}
	case 197:
		{ //967
			panic(".y:968")
		}
	case 198:
		{ //971
			panic(".y:972")
		}
	case 199:
		{ //977
			panic(".y:978")
		}
	case 200:
		{ //981
			panic(".y:982")
		}
	case 201:
		{ //987
			panic(".y:988")
		}
	case 202:
		{ //991
			panic(".y:992")
		}
	case 203:
		{ //997
			panic(".y:998")
		}
	case 204:
		{ //1001
			panic(".y:1002")
		}
	case 205:
		{ //1007
			panic(".y:1008")
		}
	case 206:
		{ //1011
			panic(".y:1012")
		}
	case 207:
		{ //1015
			panic(".y:1016")
		}
	case 208:
		{ //1019
			panic(".y:1020")
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
			panic(".y:1034")
		}
	case 212:
		{ //1037
			panic(".y:1038")
		}
	case 213:
		{ //1043
			panic(".y:1044")
		}
	case 214:
		{ //1049
			panic(".y:1050")
		}
	case 215:
		{ //1053
			panic(".y:1054")
		}
	case 216:
		{ //1057
			panic(".y:1058")
		}
	case 217:
		{ //1063
			panic(".y:1064")
		}
	case 218:
		{ //1069
			panic(".y:1070")
		}
	case 219:
		{ //1073
			panic(".y:1074")
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
			panic(".y:1088")
		}
	case 223:
		{ //1091
			panic(".y:1092")
		}
	case 224:
		{ //1096
			panic(".y:1097")
		}
	case 225:
		{ //1100
			panic(".y:1101")
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
		{ //1127
			panic(".y:1128")
		}
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
			panic(".y:1196")
		}
	case 248:
		{ //1199
			panic(".y:1200")
		}
	case 249:
		{ //1205
			panic(".y:1206")
		}
	case 250:
		{ //1209
			panic(".y:1210")
		}
	case 251:
		{ //1215
			panic(".y:1216")
		}
	case 252:
		{ //1219
			panic(".y:1220")
		}
	case 253:
		{ //1225
			panic(".y:1226")
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
	case 261:
		{ //1261
			panic(".y:1262")
		}
	case 262:
		{ //1265
			panic(".y:1266")
		}
	case 263:
		{ //1270
			panic(".y:1271")
		}
	case 264:
		{ //1274
			panic(".y:1275")
		}
	case 265:
		{ //1279
			panic(".y:1280")
		}
	case 266:
		{ //1283
			panic(".y:1284")
		}
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
			panic(".y:1307")
		}
	case 272:
		{ //1310
			panic(".y:1311")
		}
	}
	goto yystack /* stack new state and value */
}
