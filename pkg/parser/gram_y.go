// Code generated by goyacc -o gram_y.go gram.y. DO NOT EDIT.

//line gram.y:7
package parser

import __yyfmt__ "fmt"

//line gram.y:7

import (
	plast "github.com/GuanceCloud/platypus/pkg/ast"
	// pltoken "github.com/GuanceCloud/platypus/pkg/token"
)

//line gram.y:16
type yySymType struct {
	yys      int
	aststmts plast.Stmts
	astblock *plast.BlockStmt
	node     *plast.Node
	item     Item
}

const SEMICOLON = 57346
const COMMA = 57347
const COMMENT = 57348
const DOT = 57349
const EOF = 57350
const ERROR = 57351
const ID = 57352
const NUMBER = 57353
const LEFT_PAREN = 57354
const LEFT_BRACKET = 57355
const LEFT_BRACE = 57356
const RIGHT_BRACE = 57357
const RIGHT_PAREN = 57358
const RIGHT_BRACKET = 57359
const SPACE = 57360
const STRING = 57361
const QUOTED_STRING = 57362
const MULTILINE_STRING = 57363
const FOR = 57364
const IN = 57365
const WHILE = 57366
const BREAK = 57367
const CONTINUE = 57368
const RETURN = 57369
const EOL = 57370
const COLON = 57371
const STR = 57372
const INT = 57373
const FLOAT = 57374
const BOOL = 57375
const LIST = 57376
const MAP = 57377
const STRUCT = 57378
const ANY = 57379
const LET = 57380
const FN = 57381
const RET_SYMB = 57382
const NOT = 57383
const BitwiseXOR = 57384
const BitwiseOR = 57385
const BitwiseNOT = 57386
const BitwiseAND = 57387
const CONST = 57388
const operatorsStart = 57389
const ADD = 57390
const DIV = 57391
const GTE = 57392
const GT = 57393
const LT = 57394
const LTE = 57395
const MOD = 57396
const MUL = 57397
const NEQ = 57398
const EQ = 57399
const EQEQ = 57400
const SUB = 57401
const operatorsEnd = 57402
const keywordsStart = 57403
const TRUE = 57404
const FALSE = 57405
const IDENTIFIER = 57406
const AND = 57407
const OR = 57408
const NIL = 57409
const NULL = 57410
const IF = 57411
const ELIF = 57412
const ELSE = 57413
const keywordsEnd = 57414
const startSymbolsStart = 57415
const START_STMTS = 57416
const startSymbolsEnd = 57417
const UMINUS = 57418

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"SEMICOLON",
	"COMMA",
	"COMMENT",
	"DOT",
	"EOF",
	"ERROR",
	"ID",
	"NUMBER",
	"LEFT_PAREN",
	"LEFT_BRACKET",
	"LEFT_BRACE",
	"RIGHT_BRACE",
	"RIGHT_PAREN",
	"RIGHT_BRACKET",
	"SPACE",
	"STRING",
	"QUOTED_STRING",
	"MULTILINE_STRING",
	"FOR",
	"IN",
	"WHILE",
	"BREAK",
	"CONTINUE",
	"RETURN",
	"EOL",
	"COLON",
	"STR",
	"INT",
	"FLOAT",
	"BOOL",
	"LIST",
	"MAP",
	"STRUCT",
	"ANY",
	"LET",
	"FN",
	"RET_SYMB",
	"NOT",
	"BitwiseXOR",
	"BitwiseOR",
	"BitwiseNOT",
	"BitwiseAND",
	"CONST",
	"operatorsStart",
	"ADD",
	"DIV",
	"GTE",
	"GT",
	"LT",
	"LTE",
	"MOD",
	"MUL",
	"NEQ",
	"EQ",
	"EQEQ",
	"SUB",
	"operatorsEnd",
	"keywordsStart",
	"TRUE",
	"FALSE",
	"IDENTIFIER",
	"AND",
	"OR",
	"NIL",
	"NULL",
	"IF",
	"ELIF",
	"ELSE",
	"keywordsEnd",
	"startSymbolsStart",
	"START_STMTS",
	"startSymbolsEnd",
	"UMINUS",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line gram.y:640

//line yacctab:1
var yyExca = [...]int16{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 231,
	1, 32,
	4, 32,
	8, 32,
	28, 32,
	-2, 133,
}

const yyPrivate = 57344

const yyLast = 1772

var yyAct = [...]int16{
	26, 172, 253, 54, 256, 249, 142, 174, 115, 17,
	177, 8, 113, 17, 76, 74, 293, 323, 92, 71,
	286, 91, 192, 77, 94, 97, 81, 120, 3, 242,
	98, 96, 143, 285, 17, 95, 92, 93, 324, 91,
	317, 112, 94, 97, 292, 118, 97, 17, 98, 96,
	191, 98, 96, 95, 122, 303, 125, 238, 310, 129,
	45, 131, 132, 133, 134, 135, 136, 304, 126, 137,
	298, 316, 254, 296, 45, 143, 45, 144, 72, 146,
	326, 149, 139, 150, 151, 152, 153, 154, 155, 156,
	157, 158, 159, 160, 161, 162, 163, 164, 165, 166,
	2, 288, 73, 182, 170, 171, 110, 45, 179, 7,
	168, 56, 167, 70, 19, 111, 169, 181, 266, 259,
	45, 175, 183, 91, 199, 45, 94, 97, 119, 204,
	196, 207, 98, 96, 100, 264, 195, 95, 20, 287,
	32, 205, 211, 245, 102, 263, 226, 239, 229, 232,
	17, 282, 57, 209, 105, 234, 143, 186, 203, 106,
	104, 103, 189, 79, 138, 233, 105, 185, 202, 127,
	148, 106, 104, 103, 108, 128, 130, 123, 188, 109,
	107, 184, 114, 116, 117, 246, 141, 240, 187, 269,
	193, 102, 251, 4, 92, 93, 211, 91, 313, 211,
	94, 97, 83, 84, 87, 88, 98, 96, 89, 82,
	90, 95, 307, 291, 284, 279, 241, 86, 85, 101,
	211, 1, 6, 5, 18, 30, 190, 211, 271, 273,
	141, 68, 28, 247, 248, 15, 33, 52, 34, 278,
	277, 55, 116, 280, 35, 281, 173, 178, 215, 180,
	178, 37, 36, 140, 283, 39, 141, 216, 222, 231,
	290, 176, 255, 211, 38, 45, 214, 210, 9, 13,
	211, 299, 297, 14, 12, 220, 217, 218, 219, 213,
	69, 11, 221, 10, 224, 29, 16, 309, 21, 211,
	75, 212, 211, 314, 31, 27, 42, 211, 0, 0,
	223, 0, 0, 0, 0, 278, 319, 320, 0, 0,
	0, 0, 0, 0, 211, 0, 0, 178, 0, 0,
	0, 211, 0, 0, 329, 0, 0, 0, 211, 0,
	250, 0, 0, 213, 257, 0, 213, 0, 0, 231,
	231, 141, 0, 0, 141, 212, 252, 0, 212, 258,
	0, 0, 0, 0, 0, 0, 0, 213, 0, 0,
	0, 0, 0, 0, 213, 0, 141, 0, 0, 212,
	267, 0, 194, 0, 0, 0, 212, 268, 0, 0,
	0, 0, 231, 141, 0, 141, 0, 0, 0, 250,
	0, 0, 0, 0, 0, 0, 257, 0, 225, 141,
	213, 230, 0, 0, 0, 0, 0, 213, 0, 0,
	257, 141, 212, 294, 0, 0, 0, 0, 0, 212,
	295, 0, 0, 0, 141, 0, 213, 0, 0, 213,
	0, 0, 0, 141, 213, 0, 0, 0, 212, 308,
	0, 212, 312, 0, 0, 0, 212, 315, 0, 0,
	0, 213, 0, 0, 0, 0, 0, 0, 213, 0,
	0, 0, 59, 212, 321, 213, 46, 58, 53, 80,
	212, 325, 121, 0, 47, 45, 48, 212, 330, 0,
	0, 272, 274, 275, 0, 0, 276, 0, 0, 0,
	69, 0, 0, 0, 0, 0, 64, 0, 0, 65,
	61, 0, 0, 62, 0, 0, 0, 0, 289, 0,
	60, 0, 120, 0, 63, 0, 0, 49, 50, 0,
	0, 0, 51, 0, 300, 301, 0, 302, 19, 0,
	0, 59, 0, 0, 0, 46, 58, 53, 40, 0,
	0, 311, 0, 47, 45, 48, 22, 0, 0, 24,
	23, 25, 20, 318, 0, 0, 0, 0, 0, 69,
	41, 0, 66, 43, 0, 64, 322, 0, 65, 61,
	67, 0, 62, 0, 0, 328, 0, 0, 0, 60,
	0, 0, 0, 63, 0, 0, 49, 50, 0, 0,
	59, 51, 0, 44, 46, 58, 53, 40, 111, 0,
	0, 0, 47, 45, 48, 22, 0, 0, 24, 23,
	25, 0, 0, 0, 0, 0, 0, 0, 69, 41,
	0, 66, 43, 0, 64, 0, 0, 65, 61, 67,
	0, 62, 0, 0, 0, 0, 0, 0, 60, 0,
	0, 0, 63, 0, 0, 49, 50, 0, 0, 59,
	51, 0, 44, 46, 58, 53, 40, 99, 0, 0,
	0, 47, 45, 48, 22, 0, 0, 24, 23, 25,
	0, 0, 0, 0, 0, 0, 0, 69, 41, 0,
	66, 43, 0, 64, 0, 0, 65, 61, 67, 0,
	62, 0, 0, 0, 0, 0, 0, 60, 0, 0,
	0, 63, 0, 0, 49, 50, 0, 0, 59, 51,
	0, 44, 46, 58, 53, 40, 0, 0, 0, 0,
	47, 45, 48, 22, 0, 0, 24, 23, 25, 0,
	0, 0, 0, 0, 0, 0, 69, 41, 0, 66,
	43, 0, 64, 0, 0, 65, 61, 67, 0, 62,
	0, 0, 0, 0, 0, 0, 60, 0, 0, 0,
	63, 0, 0, 49, 50, 0, 227, 0, 51, 59,
	44, 0, 0, 46, 58, 53, 80, 0, 0, 0,
	0, 47, 45, 48, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 69, 0, 0,
	0, 0, 0, 64, 0, 0, 65, 61, 0, 0,
	62, 0, 0, 0, 0, 0, 0, 60, 0, 0,
	0, 63, 0, 0, 49, 50, 59, 0, 0, 51,
	46, 58, 53, 80, 111, 0, 0, 0, 47, 45,
	48, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 69, 0, 0, 0, 0, 0,
	64, 0, 0, 65, 61, 0, 0, 62, 0, 0,
	0, 0, 0, 0, 60, 0, 0, 0, 63, 0,
	0, 49, 50, 59, 0, 0, 51, 46, 58, 53,
	80, 0, 0, 198, 0, 47, 45, 48, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 69, 0, 0, 0, 0, 0, 64, 0, 0,
	65, 61, 0, 0, 62, 0, 0, 0, 0, 0,
	0, 60, 0, 0, 0, 63, 0, 147, 49, 50,
	59, 0, 0, 51, 46, 58, 53, 80, 0, 0,
	0, 0, 47, 45, 48, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 69, 0,
	0, 0, 0, 0, 64, 0, 0, 65, 61, 0,
	0, 62, 0, 0, 0, 0, 0, 0, 60, 0,
	0, 0, 63, 0, 0, 49, 50, 59, 0, 0,
	51, 46, 58, 53, 80, 124, 0, 0, 0, 47,
	45, 48, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 69, 0, 0, 0, 0,
	0, 64, 0, 0, 65, 61, 0, 0, 62, 0,
	0, 0, 0, 0, 0, 60, 0, 0, 0, 63,
	0, 78, 49, 50, 59, 0, 0, 51, 46, 58,
	53, 80, 0, 0, 0, 0, 47, 45, 48, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 69, 0, 0, 0, 0, 0, 64, 0,
	0, 65, 61, 0, 0, 62, 0, 0, 0, 0,
	0, 0, 60, 0, 0, 0, 63, 0, 0, 49,
	50, 59, 0, 0, 51, 46, 58, 53, 80, 0,
	0, 0, 0, 47, 45, 48, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 69,
	0, 0, 0, 0, 0, 64, 0, 0, 65, 61,
	0, 0, 62, 0, 0, 0, 0, 0, 0, 60,
	0, 0, 0, 63, 0, 0, 49, 50, 59, 0,
	0, 51, 46, 58, 53, 40, 0, 0, 0, 0,
	47, 45, 48, 0, 0, 0, 0, 0, 0, 262,
	0, 0, 0, 0, 0, 0, 69, 0, 0, 261,
	0, 0, 64, 0, 0, 65, 61, 0, 0, 62,
	0, 0, 0, 184, 0, 0, 60, 0, 0, 0,
	63, 236, 0, 49, 50, 0, 92, 93, 51, 91,
	0, 235, 94, 97, 83, 84, 87, 88, 98, 96,
	89, 0, 90, 95, 0, 184, 0, 0, 0, 86,
	85, 0, 186, 0, 0, 0, 0, 0, 92, 93,
	0, 91, 185, 0, 94, 97, 83, 84, 87, 88,
	98, 96, 89, 0, 90, 95, 184, 0, 0, 0,
	0, 86, 85, 201, 0, 0, 0, 0, 0, 92,
	93, 0, 91, 200, 0, 94, 97, 83, 84, 87,
	88, 98, 96, 89, 0, 90, 95, 184, 0, 0,
	0, 0, 86, 85, 0, 0, 0, 0, 0, 0,
	92, 93, 0, 91, 0, 0, 94, 97, 83, 84,
	87, 88, 98, 96, 89, 305, 90, 95, 0, 0,
	0, 0, 0, 86, 85, 92, 93, 306, 91, 0,
	0, 94, 97, 83, 84, 87, 88, 98, 96, 0,
	92, 93, 95, 91, 0, 0, 94, 97, 83, 84,
	87, 88, 98, 96, 89, 243, 90, 95, 0, 0,
	0, 0, 0, 86, 85, 0, 0, 244, 0, 0,
	0, 0, 0, 0, 327, 0, 0, 0, 0, 0,
	92, 93, 0, 91, 0, 0, 94, 97, 83, 84,
	87, 88, 98, 96, 89, 0, 90, 95, 143, 92,
	93, 0, 91, 86, 85, 94, 97, 83, 84, 87,
	88, 98, 96, 89, 270, 90, 95, 0, 0, 0,
	0, 0, 86, 85, 0, 0, 92, 93, 0, 91,
	0, 0, 94, 97, 83, 84, 87, 88, 98, 96,
	89, 0, 90, 95, 0, 0, 265, 0, 0, 86,
	85, 0, 92, 93, 0, 91, 0, 0, 94, 97,
	83, 84, 87, 88, 98, 96, 89, 260, 90, 95,
	0, 92, 93, 0, 91, 86, 85, 94, 97, 83,
	84, 87, 88, 98, 96, 89, 0, 90, 95, 0,
	0, 0, 0, 0, 86, 85, 0, 0, 237, 0,
	0, 0, 0, 0, 92, 93, 0, 91, 0, 0,
	94, 97, 83, 84, 87, 88, 98, 96, 89, 228,
	90, 95, 0, 92, 93, 0, 91, 86, 85, 94,
	97, 83, 84, 87, 88, 98, 96, 89, 0, 90,
	95, 0, 0, 0, 0, 208, 86, 85, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 92, 93, 0,
	91, 0, 0, 94, 97, 83, 84, 87, 88, 98,
	96, 89, 0, 90, 95, 206, 0, 0, 0, 0,
	86, 85, 92, 93, 0, 91, 0, 0, 94, 97,
	83, 84, 87, 88, 98, 96, 89, 197, 90, 95,
	0, 92, 93, 0, 91, 86, 85, 94, 97, 83,
	84, 87, 88, 98, 96, 89, 0, 90, 95, 0,
	0, 145, 0, 0, 86, 85, 0, 0, 0, 0,
	0, 0, 0, 0, 92, 93, 0, 91, 0, 0,
	94, 97, 83, 84, 87, 88, 98, 96, 89, 0,
	90, 95, 0, 0, 0, 0, 0, 86, 85, 92,
	93, 0, 91, 0, 0, 94, 97, 83, 84, 87,
	88, 98, 96, 89, 0, 90, 95, 0, 92, 93,
	0, 91, 86, 85, 94, 97, 83, 84, 87, 88,
	98, 96, 89, 82, 90, 95, 0, 92, 93, 0,
	91, 86, 85, 94, 97, 83, 84, 87, 88, 98,
	96, 89, 0, 90, 95, 0, 92, 93, 0, 91,
	86, 85, 94, 97, 83, 84, 87, 88, 98, 96,
	89, 0, 90, 95, 92, 93, 0, 91, 0, 86,
	94, 97, 83, 84, 87, 88, 98, 96, 89, 0,
	90, 95,
}

var yyPact = [...]int16{
	26, 185, 524, -1000, -1000, -1000, 701, 110, 74, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -56, 1047, -1000, -1000, 1104, 1656, 642, -1000, 214,
	-1000, -1000, 159, -1000, -1000, -1000, -1000, 167, -1000, -1000,
	583, 40, 40, 40, 1104, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 111, 455, 164, 990, -1000, 161, 1104, 163,
	1104, 1104, 1104, 1104, 1104, 1104, -1000, -1000, 1104, 151,
	110, 74, -1000, -1000, 142, -1000, 1104, 1637, 933, 147,
	819, 1675, 1104, 1104, 1104, 1104, 1104, 1104, 1104, 1104,
	1104, 1104, 1104, 1104, 1104, 1104, 1104, 1104, 1104, -1000,
	110, 40, -1000, 819, 1104, 40, 105, 1104, 40, 87,
	110, -1000, 152, 173, 148, -1000, -7, 178, 1404, -1000,
	119, -1000, 1612, 876, -1000, 1278, 153, -1000, 819, 1579,
	1104, -1000, -1000, -1000, -1000, -1000, -1000, 1560, 245, 74,
	-1000, -1000, -1000, 583, 1404, 762, 1535, 1161, 1104, 1247,
	1675, -6, -6, 1694, 1712, -6, -6, 1303, 1303, 78,
	-24, -6, -3, -3, -1000, -1000, -1000, 74, -1000, 150,
	1216, 1501, 28, -1000, 131, -1000, 40, 211, -28, 1358,
	-1000, 127, -1000, 74, 1104, -1000, -1000, -1000, -1000, 100,
	-1000, 1104, 245, 56, -1000, 245, 102, -1000, -1000, 1482,
	-1000, -1000, -1000, -1000, 1184, 130, -1000, 1449, -1000, 101,
	-1000, -1000, -1000, -1000, -1000, -1000, 245, -1000, -1000, -1000,
	-1000, -1000, -30, 245, 177, -1000, 1430, 1161, 1161, 1404,
	-1000, -1000, 1404, -1000, -1000, -1000, -1000, -1000, 1104, -1000,
	210, -1000, 1104, -1000, 1104, -1000, 1675, 136, 40, 209,
	4, 1675, -37, 123, 61, 40, 208, -13, -1000, 245,
	-1000, -1000, -1000, -1000, -1000, -1000, 245, 57, -1000, 54,
	1161, 1404, -1000, 1404, -1000, -1000, -1000, 38, 1675, -1000,
	1675, 1318, -1000, 207, -1000, 245, 1104, 18, 245, -1000,
	193, -1000, 1104, 245, -1000, -1000, -1000, 55, 0, 1404,
	-1000, -1000, -1000, -1000, 1104, -1000, 1104, -1000, -1000, 1675,
	245, -1000, 142, -1000, 1675, -40, -2, 245, -1000, 63,
	1377, 142, -1000, 1104, 245, -1000, -1000, -1000, -1000, 1675,
	-1000,
}

var yyPgo = [...]int16{
	0, 296, 253, 111, 6, 295, 294, 290, 288, 106,
	286, 285, 283, 281, 274, 273, 269, 268, 0, 1,
	152, 267, 153, 3, 266, 264, 262, 4, 2, 261,
	10, 7, 140, 255, 252, 251, 248, 244, 12, 241,
	238, 237, 236, 235, 5, 234, 233, 232, 231, 225,
	8, 224, 223, 222, 221, 11,
}

var yyR1 = [...]int8{
	0, 55, 55, 55, 55, 54, 54, 54, 52, 52,
	52, 53, 53, 53, 9, 9, 9, 9, 9, 9,
	9, 9, 9, 9, 18, 18, 18, 18, 18, 18,
	18, 18, 2, 2, 4, 5, 5, 3, 17, 17,
	6, 8, 8, 7, 13, 13, 13, 13, 13, 13,
	13, 13, 12, 14, 16, 15, 15, 51, 51, 51,
	43, 10, 49, 49, 49, 49, 27, 27, 27, 27,
	26, 26, 28, 28, 28, 11, 11, 1, 1, 50,
	50, 50, 50, 47, 47, 44, 44, 45, 45, 46,
	46, 46, 32, 22, 22, 22, 22, 22, 22, 22,
	24, 21, 21, 21, 21, 21, 23, 23, 20, 36,
	36, 36, 36, 42, 42, 42, 42, 42, 42, 40,
	40, 40, 41, 41, 41, 41, 48, 48, 48, 38,
	37, 37, 37, 37, 37, 37, 37, 37, 37, 37,
	37, 37, 39, 39, 39, 39, 39, 39, 39, 39,
	34, 30, 30, 29, 29, 31, 31, 31, 19, 19,
	35, 35, 35, 35, 35, 35, 35, 35, 35, 35,
	35, 35, 35, 25, 25, 25, 25, 25, 25, 33,
	33, 33, 33, 33, 33, 33, 33, 33, 33, 33,
	33, 33, 33, 33, 33,
}

var yyR2 = [...]int8{
	0, 1, 1, 2, 2, 2, 2, 1, 2, 1,
	1, 2, 1, 3, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 2, 3, 3, 2, 1, 3,
	3, 1, 2, 3, 7, 6, 6, 5, 6, 5,
	5, 4, 5, 1, 1, 1, 2, 1, 1, 1,
	1, 3, 8, 7, 6, 5, 1, 3, 3, 5,
	2, 3, 1, 2, 1, 2, 3, 1, 1, 1,
	3, 5, 3, 5, 3, 1, 3, 2, 3, 1,
	2, 1, 1, 1, 1, 1, 1, 1, 1, 3,
	2, 1, 1, 1, 1, 1, 4, 5, 5, 6,
	5, 4, 3, 1, 1, 1, 1, 1, 1, 2,
	2, 3, 1, 2, 2, 3, 3, 4, 3, 3,
	2, 3, 3, 1, 2, 2, 3, 3, 4, 4,
	4, 4, 3, 3, 4, 4, 4, 4, 3, 3,
	3, 1, 3, 2, 3, 1, 2, 1, 1, 0,
	4, 4, 3, 4, 3, 6, 8, 4, 3, 4,
	3, 6, 8, 2, 2, 2, 2, 2, 2, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3,
}

var yyChk = [...]int16{
	-1000, -54, 74, 2, 8, -52, -53, -9, -55, -17,
	-12, -13, -14, -16, -15, -43, -10, -4, -51, 4,
	28, -8, 22, 26, 25, 27, -18, -5, -47, -11,
	-49, -6, -32, -42, -40, -37, -34, -35, -25, -33,
	14, 36, -1, 39, 69, 20, 11, 19, 21, 62,
	63, 67, -41, 13, -23, -39, -3, -20, 12, 7,
	55, 45, 48, 59, 41, 44, 38, 46, -48, 35,
	-9, -55, 4, 28, 71, -7, 70, -18, 4, -32,
	14, -18, 57, 50, 51, 66, 65, 52, 53, 56,
	58, 45, 42, 43, 48, 59, 55, 49, 54, 15,
	-9, 5, -3, 14, 13, 7, 12, 13, 7, 12,
	-9, 15, -18, -38, -32, -50, -32, -32, -18, 17,
	57, 17, -18, 13, 15, -18, -38, -3, 14, -18,
	13, -18, -18, -18, -18, -18, -18, -18, 13, -55,
	-2, -3, -4, 14, -18, 4, -18, 4, 23, -18,
	-18, -18, -18, -18, -18, -18, -18, -18, -18, -18,
	-18, -18, -18, -18, -18, -18, -18, -55, -50, -38,
	-18, -18, -19, -32, -31, 16, -29, -30, -32, -18,
	-32, -31, 16, -55, 29, 15, 5, 15, 5, 14,
	-3, 57, 29, 12, -2, 17, 11, 5, 17, -18,
	15, 5, 15, 5, -18, -38, 16, -18, 5, -22,
	-21, -23, -20, -32, -24, -36, 12, 31, 32, 33,
	30, 37, 13, 55, 39, -2, -18, 4, 4, -18,
	-2, -3, -18, 15, 5, 15, 5, 17, 29, 16,
	-30, 5, 57, 17, 29, 16, -18, -46, -45, -44,
	-32, -18, -22, -28, 16, -26, -27, -32, -22, 17,
	5, 15, 5, 15, 5, 17, 17, -22, -22, 12,
	4, -18, -2, -18, -2, -2, -2, -19, -18, 5,
	-18, -18, 15, -44, 5, 29, 57, 16, 40, -2,
	-27, 5, 57, 29, -22, -22, 16, -28, 16, -18,
	-2, -2, -2, 17, 29, 17, 29, 5, -22, -18,
	40, -2, -22, 5, -18, -22, 16, 40, -2, -19,
	-18, -22, -2, 57, 40, -22, 17, 17, -2, -18,
	-22,
}

var yyDef = [...]int16{
	0, -2, 0, 7, 6, 5, 9, 10, 12, 14,
	15, 16, 17, 18, 19, 20, 21, 22, 23, 1,
	2, 38, 0, 53, 54, 55, 60, 0, 57, 58,
	59, 41, 24, 25, 26, 27, 28, 29, 30, 31,
	0, 0, 0, 0, 0, 92, 113, 114, 115, 116,
	117, 118, 0, 0, 0, 0, 133, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 77, 78, 122, 0,
	8, 11, 3, 4, 0, 42, 0, 0, 0, 24,
	0, 56, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 34,
	0, 0, 135, 0, 159, 0, 0, 0, 0, 0,
	0, 37, 60, 0, 0, 75, 79, 0, 0, 119,
	0, 120, 124, 0, 130, 0, 0, 134, 0, 0,
	0, 173, 174, 175, 176, 177, 178, 123, 0, 13,
	39, 32, 33, 0, 0, 0, 0, 0, 0, 0,
	61, 179, 180, 181, 182, 183, 184, 185, 186, 187,
	188, 189, 190, 191, 192, 193, 194, 36, 76, 0,
	0, 158, 0, 162, 0, 164, 155, 157, 151, 0,
	168, 0, 170, 35, 0, 136, 143, 137, 142, 0,
	84, 0, 0, 0, 40, 0, 0, 126, 121, 125,
	131, 149, 132, 148, 0, 0, 150, 0, 128, 0,
	93, 94, 95, 96, 97, 98, 0, 101, 102, 103,
	104, 105, 0, 0, 0, 43, 0, 0, 0, 0,
	51, -2, 0, 138, 146, 139, 147, 160, 159, 163,
	156, 153, 0, 167, 0, 169, 129, 0, 89, 91,
	85, 80, 82, 0, 0, 72, 74, 66, 106, 0,
	127, 140, 145, 141, 144, 161, 0, 0, 100, 0,
	0, 0, 47, 0, 49, 50, 52, 0, 158, 154,
	152, 0, 83, 90, 87, 0, 0, 0, 0, 65,
	73, 70, 0, 0, 107, 108, 99, 0, 112, 0,
	45, 46, 48, 165, 159, 171, 0, 88, 86, 81,
	0, 64, 0, 71, 67, 68, 111, 0, 44, 0,
	0, 0, 63, 0, 0, 110, 166, 172, 62, 69,
	109,
}

var yyTok1 = [...]int8{
	1,
}

var yyTok2 = [...]int8{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 64, 65, 66, 67, 68, 69, 70, 71,
	72, 73, 74, 75, 76,
}

var yyTok3 = [...]int8{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
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

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := int(yyPact[state])
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && int(yyChk[int(yyAct[n])]) == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || int(yyExca[i+1]) != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := int(yyExca[i])
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = int(yyTok1[0])
		goto out
	}
	if char < len(yyTok1) {
		token = int(yyTok1[char])
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = int(yyTok2[char-yyPrivate])
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = int(yyTok3[i+0])
		if token == char {
			token = int(yyTok3[i+1])
			goto out
		}
	}

out:
	if token == 0 {
		token = int(yyTok2[1]) /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
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
	yyn = int(yyPact[yystate])
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = int(yyAct[yyn])
	if int(yyChk[yyn]) == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = int(yyDef[yystate])
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && int(yyExca[xi+1]) == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = int(yyExca[xi+0])
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = int(yyExca[xi+1])
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = int(yyPact[yyS[yyp].yys]) + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = int(yyAct[yyn]) /* simulate a shift of "error" */
					if int(yyChk[yystate]) == yyErrCode {
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
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
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

	yyp -= int(yyR2[yyn])
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = int(yyR1[yyn])
	yyg := int(yyPgo[yyn])
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = int(yyAct[yyg])
	} else {
		yystate = int(yyAct[yyj])
		if int(yyChk[yystate]) != -yyn {
			yystate = int(yyAct[yyg])
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 5:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:139
		{
			yylex.(*parser).parseResult = yyDollar[2].aststmts
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:142
		{
			yylex.(*parser).unexpected("", "")
		}
	case 8:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:147
		{
			s := yyDollar[1].aststmts
			s = append(s, yyDollar[2].node)
			yyVAL.aststmts = s
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:154
		{
			yyVAL.aststmts = plast.Stmts{yyDollar[1].node}
		}
	case 11:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:158
		{
			yyVAL.aststmts = plast.Stmts{yyDollar[1].node}
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:160
		{
			yyVAL.aststmts = plast.Stmts{}
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:162
		{
			s := yyDollar[1].aststmts
			s = append(s, yyDollar[2].node)
			yyVAL.aststmts = s
		}
	case 34:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:202
		{
			yyVAL.node = &ast.Node{}
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:206
		{
			yyVAL.node = &ast.Node{}
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:208
		{
			yyVAL.node = &ast.Node{}
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:212
		{
			yyVAL.node = &ast.Node{}
		}
	case 38:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:216
		{
			yyVAL.node = &ast.Node{}
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:218
		{
			yyVAL.node = &ast.Node{}
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:222
		{
			yyVAL.node = &ast.Node{}
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:226
		{
			yyVAL.node = &ast.Node{}
		}
	case 42:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:228
		{
			yyVAL.node = &ast.Node{}
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:232
		{
			yyVAL.node = &ast.Node{}
		}
	case 44:
		yyDollar = yyS[yypt-7 : yypt+1]
//line gram.y:241
		{
			yyVAL.node = &ast.Node{}
		}
	case 45:
		yyDollar = yyS[yypt-6 : yypt+1]
//line gram.y:243
		{
			yyVAL.node = &ast.Node{}
		}
	case 46:
		yyDollar = yyS[yypt-6 : yypt+1]
//line gram.y:245
		{
			yyVAL.node = &ast.Node{}
		}
	case 47:
		yyDollar = yyS[yypt-5 : yypt+1]
//line gram.y:247
		{
			yyVAL.node = &ast.Node{}
		}
	case 48:
		yyDollar = yyS[yypt-6 : yypt+1]
//line gram.y:249
		{
			yyVAL.node = &ast.Node{}
		}
	case 49:
		yyDollar = yyS[yypt-5 : yypt+1]
//line gram.y:251
		{
			yyVAL.node = &ast.Node{}
		}
	case 50:
		yyDollar = yyS[yypt-5 : yypt+1]
//line gram.y:253
		{
			yyVAL.node = &ast.Node{}
		}
	case 51:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:255
		{
			yyVAL.node = &ast.Node{}
		}
	case 52:
		yyDollar = yyS[yypt-5 : yypt+1]
//line gram.y:259
		{
			yyVAL.node = &ast.Node{}
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:263
		{
			yyVAL.node = &ast.Node{}
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:268
		{
			yyVAL.node = &ast.Node{}
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:273
		{
			yyVAL.node = &ast.Node{}
		}
	case 56:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:275
		{
			yyVAL.node = &ast.Node{}
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:289
		{
			yyVAL.node = &ast.Node{}
		}
	case 62:
		yyDollar = yyS[yypt-8 : yypt+1]
//line gram.y:298
		{
			yyVAL.node = &ast.Node{}
		}
	case 63:
		yyDollar = yyS[yypt-7 : yypt+1]
//line gram.y:300
		{
			yyVAL.node = &ast.Node{}
		}
	case 64:
		yyDollar = yyS[yypt-6 : yypt+1]
//line gram.y:302
		{
			yyVAL.node = &ast.Node{}
		}
	case 65:
		yyDollar = yyS[yypt-5 : yypt+1]
//line gram.y:304
		{
			yyVAL.node = &ast.Node{}
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:309
		{
			yyVAL.node = &ast.Node{}
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:311
		{
			yyVAL.node = &ast.Node{}
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:313
		{
			yyVAL.node = &ast.Node{}
		}
	case 69:
		yyDollar = yyS[yypt-5 : yypt+1]
//line gram.y:315
		{
			yyVAL.node = &ast.Node{}
		}
	case 70:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:319
		{
			yyVAL.node = &ast.Node{}
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:321
		{
			yyVAL.node = &ast.Node{}
		}
	case 75:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:331
		{
			yyVAL.node = &ast.Node{}
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:333
		{
			yyVAL.node = &ast.Node{}
		}
	case 83:
		yyDollar = yyS[yypt-5 : yypt+1]
//line gram.y:347
		{
			yyVAL.node = &ast.Node{}
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:349
		{
			yyVAL.node = &ast.Node{}
		}
	case 92:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:375
		{
			yyVAL.node = &ast.Node{}
		}
	case 99:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:386
		{
			yyVAL.node = &ast.Node{}
		}
	case 100:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:390
		{
			yyVAL.node = &ast.Node{}
		}
	case 101:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:395
		{
			yyVAL.node = &ast.Node{}
		}
	case 102:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:397
		{
			yyVAL.node = &ast.Node{}
		}
	case 103:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:399
		{
			yyVAL.node = &ast.Node{}
		}
	case 104:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:401
		{
			yyVAL.node = &ast.Node{}
		}
	case 105:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:403
		{
			yyVAL.node = &ast.Node{}
		}
	case 106:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:408
		{
			yyVAL.node = &ast.Node{}
		}
	case 107:
		yyDollar = yyS[yypt-5 : yypt+1]
//line gram.y:410
		{
			yyVAL.node = &ast.Node{}
		}
	case 108:
		yyDollar = yyS[yypt-5 : yypt+1]
//line gram.y:414
		{
			yyVAL.node = &ast.Node{}
		}
	case 109:
		yyDollar = yyS[yypt-6 : yypt+1]
//line gram.y:418
		{
			yyVAL.node = &ast.Node{}
		}
	case 110:
		yyDollar = yyS[yypt-5 : yypt+1]
//line gram.y:420
		{
			yyVAL.node = &ast.Node{}
		}
	case 111:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:422
		{
			yyVAL.node = &ast.Node{}
		}
	case 112:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:424
		{
			yyVAL.node = &ast.Node{}
		}
	case 113:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:429
		{
			yyVAL.node = &ast.Node{}
		}
	case 114:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:431
		{
			yyVAL.node = &ast.Node{}
		}
	case 115:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:433
		{
			yyVAL.node = &ast.Node{}
		}
	case 116:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:435
		{
			yyVAL.node = &ast.Node{}
		}
	case 117:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:437
		{
			yyVAL.node = &ast.Node{}
		}
	case 118:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:439
		{
			yyVAL.node = &ast.Node{}
		}
	case 119:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:443
		{
			yyVAL.node = &ast.Node{}
		}
	case 120:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:445
		{
			yyVAL.node = &ast.Node{}
		}
	case 121:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:447
		{
			yyVAL.node = &ast.Node{}
		}
	case 122:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:451
		{
			yyVAL.node = &ast.Node{}
		}
	case 123:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:453
		{
			yyVAL.node = &ast.Node{}
		}
	case 124:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:455
		{
			yyVAL.node = &ast.Node{}
		}
	case 125:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:457
		{
			yyVAL.node = &ast.Node{}
		}
	case 126:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:461
		{
			yyVAL.node = &ast.Node{}
		}
	case 127:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:463
		{
			yyVAL.node = &ast.Node{}
		}
	case 128:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:465
		{
			yyVAL.node = &ast.Node{}
		}
	case 130:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:472
		{
			yyVAL.node = &ast.Node{}
		}
	case 131:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:474
		{
			yyVAL.node = &ast.Node{}
		}
	case 132:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:476
		{
			yyVAL.node = &ast.Node{}
		}
	case 133:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:478
		{
			yyVAL.node = &ast.Node{}
		}
	case 134:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:480
		{
			yyVAL.node = &ast.Node{}
		}
	case 135:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:482
		{
			yyVAL.node = &ast.Node{}
		}
	case 136:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:484
		{
			yyVAL.node = &ast.Node{}
		}
	case 137:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:486
		{
			yyVAL.node = &ast.Node{}
		}
	case 138:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:488
		{
			yyVAL.node = &ast.Node{}
		}
	case 139:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:490
		{
			yyVAL.node = &ast.Node{}
		}
	case 140:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:492
		{
			yyVAL.node = &ast.Node{}
		}
	case 141:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:494
		{
			yyVAL.node = &ast.Node{}
		}
	case 142:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:498
		{
			yyVAL.node = &ast.Node{}
		}
	case 143:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:500
		{
			yyVAL.node = &ast.Node{}
		}
	case 144:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:502
		{
			yyVAL.node = &ast.Node{}
		}
	case 145:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:504
		{
			yyVAL.node = &ast.Node{}
		}
	case 146:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:506
		{
			yyVAL.node = &ast.Node{}
		}
	case 147:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:508
		{
			yyVAL.node = &ast.Node{}
		}
	case 148:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:510
		{
			yyVAL.node = &ast.Node{}
		}
	case 149:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:512
		{
			yyVAL.node = &ast.Node{}
		}
	case 150:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:518
		{
			yyVAL.node = &ast.Node{}
		}
	case 151:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:524
		{
			yyVAL.node = &ast.Node{}
		}
	case 152:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:526
		{
			yyVAL.node = &ast.Node{}
		}
	case 153:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:530
		{
			yyVAL.node = &ast.Node{}
		}
	case 154:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:532
		{
			yyVAL.node = &ast.Node{}
		}
	case 159:
		yyDollar = yyS[yypt-0 : yypt+1]
//line gram.y:542
		{
			yyVAL.node = nil
		}
	case 160:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:549
		{
			yyVAL.node = &ast.Node{}
		}
	case 161:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:551
		{
			yyVAL.node = &ast.Node{}
		}
	case 162:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:555
		{
			yyVAL.node = &ast.Node{}
		}
	case 163:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:559
		{
			yyVAL.node = &ast.Node{}
		}
	case 164:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:561
		{
			yyVAL.node = &ast.Node{}
		}
	case 165:
		yyDollar = yyS[yypt-6 : yypt+1]
//line gram.y:565
		{
			yyVAL.node = &ast.Node{}
		}
	case 166:
		yyDollar = yyS[yypt-8 : yypt+1]
//line gram.y:567
		{
			yyVAL.node = &ast.Node{}
		}
	case 167:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:571
		{
			yyVAL.node = &ast.Node{}
		}
	case 168:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:575
		{
			yyVAL.node = &ast.Node{}
		}
	case 169:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:579
		{
			yyVAL.node = &ast.Node{}
		}
	case 170:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:581
		{
			yyVAL.node = &ast.Node{}
		}
	case 171:
		yyDollar = yyS[yypt-6 : yypt+1]
//line gram.y:585
		{
			yyVAL.node = &ast.Node{}
		}
	case 172:
		yyDollar = yyS[yypt-8 : yypt+1]
//line gram.y:587
		{
			yyVAL.node = &ast.Node{}
		}
	case 173:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:592
		{
			yyVAL.node = &ast.Node{}
		}
	case 174:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:594
		{
			yyVAL.node = &ast.Node{}
		}
	case 175:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:596
		{
			yyVAL.node = &ast.Node{}
		}
	case 176:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:598
		{
			yyVAL.node = &ast.Node{}
		}
	case 177:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:600
		{
			yyVAL.node = &ast.Node{}
		}
	case 178:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:602
		{
			yyVAL.node = &ast.Node{}
		}
	case 179:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:606
		{
			yyVAL.node = &ast.Node{}
		}
	case 180:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:608
		{
			yyVAL.node = &ast.Node{}
		}
	case 181:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:610
		{
			yyVAL.node = &ast.Node{}
		}
	case 182:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:612
		{
			yyVAL.node = &ast.Node{}
		}
	case 183:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:614
		{
			yyVAL.node = &ast.Node{}
		}
	case 184:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:616
		{
			yyVAL.node = &ast.Node{}
		}
	case 185:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:618
		{
			yyVAL.node = &ast.Node{}
		}
	case 186:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:620
		{
			yyVAL.node = &ast.Node{}
		}
	case 187:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:622
		{
			yyVAL.node = &ast.Node{}
		}
	case 188:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:624
		{
			yyVAL.node = &ast.Node{}
		}
	case 189:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:626
		{
			yyVAL.node = &ast.Node{}
		}
	case 190:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:628
		{
			yyVAL.node = &ast.Node{}
		}
	case 191:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:630
		{
			yyVAL.node = &ast.Node{}
		}
	case 192:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:632
		{
			yyVAL.node = &ast.Node{}
		}
	case 193:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:634
		{
			yyVAL.node = &ast.Node{}
		}
	case 194:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:636
		{
			yyVAL.node = &ast.Node{}
		}
	}
	goto yystack /* stack new state and value */
}
