// Code generated from LEQLParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package leql // LEQLParser
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type LEQLParser struct {
	*antlr.BaseParser
}

var LEQLParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func leqlparserParserInit() {
	staticData := &LEQLParserParserStaticData
	staticData.LiteralNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "'!=='", "'!~'", "'!='", "'=~'",
		"'>='", "'<='", "'=='", "'='", "'>'", "'<'", "'!'", "'('", "')'", "'['",
		"']'", "','", "':'", "'.'", "'*'", "'#'",
	}
	staticData.SymbolicNames = []string{
		"", "SELECT", "WHERE", "GROUPBY", "HAVING", "CALCULATE", "SORT", "LIMIT",
		"TIMESLICE", "FROM", "LOOSE", "AND", "OR", "NOT", "ICONTAINS_ALL_OP",
		"ICONTAINS_ANY_OP", "CONTAINS_ALL_OP", "CONTAINS_ANY_OP", "ISTARTS_WITH_ANY_OP",
		"STARTS_WITH_ANY_OP", "ISTARTS_WITH_OP", "STARTS_WITH_OP", "ICONTAINS_OP",
		"CONTAINS_OP", "IIN_OP", "IN_OP", "NOCASE", "IP_FUNC", "ALL_FUNC", "STANDARDDEVIATION",
		"PERCENTILE", "AVERAGE", "UNIQUE", "COUNT", "BYTES", "PCTL", "SUM",
		"MIN", "MAX", "SD", "ASCENDING", "DESCENDING", "ASC", "DESC", "STRICT_NEQ",
		"REGEX_NEQ", "NEQ", "REGEX_EQ", "GTE", "LTE", "STRICT_EQ", "EQ", "GT",
		"LT", "BANG", "LPAREN", "RPAREN", "LBRACKET", "RBRACKET", "COMMA", "COLON",
		"DOT", "STAR", "HASH", "IP_CIDR", "REGEX", "TRIPLE_SINGLE_STRING", "TRIPLE_DOUBLE_STRING",
		"DOUBLE_STRING", "SINGLE_STRING", "NUMBER", "TIME_UNIT", "VARIABLE",
		"IDENTIFIER", "WS",
	}
	staticData.RuleNames = []string{
		"query", "fromClause", "selectClause", "selectFieldList", "selectField",
		"whereClause", "expression", "condition", "keywordSearch", "regexSearch",
		"fieldList", "allFieldList", "fieldName", "comparisonOp", "regexMatchOp",
		"stringOp", "setOp", "listStringOp", "value", "nocaseValue", "ipValue",
		"valueList", "valueListItem", "groupbyClause", "havingClause", "havingCondition",
		"calculateClause", "calcExpr", "calcFunction", "calcFunctionWithField",
		"percentileFunction", "sortClause", "sortSpec", "sortDirection", "limitClause",
		"timesliceClause",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 74, 419, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 1, 0, 3, 0,
		74, 8, 0, 1, 0, 3, 0, 77, 8, 0, 1, 0, 3, 0, 80, 8, 0, 1, 0, 3, 0, 83, 8,
		0, 1, 0, 3, 0, 86, 8, 0, 1, 0, 3, 0, 89, 8, 0, 1, 0, 3, 0, 92, 8, 0, 1,
		0, 3, 0, 95, 8, 0, 1, 0, 3, 0, 98, 8, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 5, 3, 115,
		8, 3, 10, 3, 12, 3, 118, 9, 3, 1, 4, 1, 4, 1, 4, 3, 4, 123, 8, 4, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 130, 8, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		3, 6, 149, 8, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6,
		159, 8, 6, 10, 6, 12, 6, 162, 9, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 261, 8, 7, 1, 8, 1, 8,
		3, 8, 265, 8, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 5, 10, 272, 8, 10, 10,
		10, 12, 10, 275, 9, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 5, 11, 282,
		8, 11, 10, 11, 12, 11, 285, 9, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1,
		13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20, 312,
		8, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 5, 21, 320, 8, 21, 10,
		21, 12, 21, 323, 9, 21, 1, 21, 1, 21, 1, 22, 1, 22, 3, 22, 329, 8, 22,
		1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 5, 23, 336, 8, 23, 10, 23, 12, 23, 339,
		9, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 3, 25, 356, 8, 25, 1, 26, 1, 26,
		1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 3, 27, 366, 8, 27, 1, 28, 1,
		28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30,
		3, 30, 380, 8, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 5, 31, 387, 8, 31,
		10, 31, 12, 31, 390, 9, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 3,
		32, 398, 8, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 5, 34,
		407, 8, 34, 10, 34, 12, 34, 410, 9, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1,
		35, 1, 35, 1, 35, 1, 35, 0, 1, 12, 36, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18,
		20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54,
		56, 58, 60, 62, 64, 66, 68, 70, 0, 12, 6, 0, 9, 10, 27, 28, 33, 34, 36,
		39, 68, 70, 73, 73, 3, 0, 44, 44, 46, 46, 48, 53, 2, 0, 45, 45, 47, 47,
		1, 0, 20, 23, 1, 0, 24, 25, 1, 0, 14, 19, 1, 0, 64, 73, 1, 0, 33, 34, 3,
		0, 29, 29, 31, 34, 36, 39, 2, 0, 30, 30, 35, 35, 1, 0, 40, 43, 1, 0, 70,
		71, 440, 0, 73, 1, 0, 0, 0, 2, 101, 1, 0, 0, 0, 4, 106, 1, 0, 0, 0, 6,
		111, 1, 0, 0, 0, 8, 119, 1, 0, 0, 0, 10, 124, 1, 0, 0, 0, 12, 148, 1, 0,
		0, 0, 14, 260, 1, 0, 0, 0, 16, 264, 1, 0, 0, 0, 18, 266, 1, 0, 0, 0, 20,
		268, 1, 0, 0, 0, 22, 276, 1, 0, 0, 0, 24, 288, 1, 0, 0, 0, 26, 290, 1,
		0, 0, 0, 28, 292, 1, 0, 0, 0, 30, 294, 1, 0, 0, 0, 32, 296, 1, 0, 0, 0,
		34, 298, 1, 0, 0, 0, 36, 300, 1, 0, 0, 0, 38, 302, 1, 0, 0, 0, 40, 307,
		1, 0, 0, 0, 42, 315, 1, 0, 0, 0, 44, 328, 1, 0, 0, 0, 46, 330, 1, 0, 0,
		0, 48, 342, 1, 0, 0, 0, 50, 355, 1, 0, 0, 0, 52, 357, 1, 0, 0, 0, 54, 365,
		1, 0, 0, 0, 56, 367, 1, 0, 0, 0, 58, 369, 1, 0, 0, 0, 60, 373, 1, 0, 0,
		0, 62, 381, 1, 0, 0, 0, 64, 393, 1, 0, 0, 0, 66, 399, 1, 0, 0, 0, 68, 401,
		1, 0, 0, 0, 70, 413, 1, 0, 0, 0, 72, 74, 3, 2, 1, 0, 73, 72, 1, 0, 0, 0,
		73, 74, 1, 0, 0, 0, 74, 76, 1, 0, 0, 0, 75, 77, 3, 4, 2, 0, 76, 75, 1,
		0, 0, 0, 76, 77, 1, 0, 0, 0, 77, 79, 1, 0, 0, 0, 78, 80, 3, 10, 5, 0, 79,
		78, 1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 82, 1, 0, 0, 0, 81, 83, 3, 46,
		23, 0, 82, 81, 1, 0, 0, 0, 82, 83, 1, 0, 0, 0, 83, 85, 1, 0, 0, 0, 84,
		86, 3, 52, 26, 0, 85, 84, 1, 0, 0, 0, 85, 86, 1, 0, 0, 0, 86, 88, 1, 0,
		0, 0, 87, 89, 3, 48, 24, 0, 88, 87, 1, 0, 0, 0, 88, 89, 1, 0, 0, 0, 89,
		91, 1, 0, 0, 0, 90, 92, 3, 62, 31, 0, 91, 90, 1, 0, 0, 0, 91, 92, 1, 0,
		0, 0, 92, 94, 1, 0, 0, 0, 93, 95, 3, 68, 34, 0, 94, 93, 1, 0, 0, 0, 94,
		95, 1, 0, 0, 0, 95, 97, 1, 0, 0, 0, 96, 98, 3, 70, 35, 0, 97, 96, 1, 0,
		0, 0, 97, 98, 1, 0, 0, 0, 98, 99, 1, 0, 0, 0, 99, 100, 5, 0, 0, 1, 100,
		1, 1, 0, 0, 0, 101, 102, 5, 9, 0, 0, 102, 103, 5, 55, 0, 0, 103, 104, 3,
		12, 6, 0, 104, 105, 5, 56, 0, 0, 105, 3, 1, 0, 0, 0, 106, 107, 5, 1, 0,
		0, 107, 108, 5, 55, 0, 0, 108, 109, 3, 6, 3, 0, 109, 110, 5, 56, 0, 0,
		110, 5, 1, 0, 0, 0, 111, 116, 3, 8, 4, 0, 112, 113, 5, 59, 0, 0, 113, 115,
		3, 8, 4, 0, 114, 112, 1, 0, 0, 0, 115, 118, 1, 0, 0, 0, 116, 114, 1, 0,
		0, 0, 116, 117, 1, 0, 0, 0, 117, 7, 1, 0, 0, 0, 118, 116, 1, 0, 0, 0, 119,
		122, 3, 24, 12, 0, 120, 121, 5, 73, 0, 0, 121, 123, 5, 73, 0, 0, 122, 120,
		1, 0, 0, 0, 122, 123, 1, 0, 0, 0, 123, 9, 1, 0, 0, 0, 124, 125, 5, 2, 0,
		0, 125, 126, 5, 55, 0, 0, 126, 129, 3, 12, 6, 0, 127, 128, 5, 59, 0, 0,
		128, 130, 5, 10, 0, 0, 129, 127, 1, 0, 0, 0, 129, 130, 1, 0, 0, 0, 130,
		131, 1, 0, 0, 0, 131, 132, 5, 56, 0, 0, 132, 11, 1, 0, 0, 0, 133, 134,
		6, 6, -1, 0, 134, 135, 5, 55, 0, 0, 135, 136, 3, 12, 6, 0, 136, 137, 5,
		56, 0, 0, 137, 149, 1, 0, 0, 0, 138, 139, 5, 2, 0, 0, 139, 140, 5, 55,
		0, 0, 140, 141, 3, 12, 6, 0, 141, 142, 5, 56, 0, 0, 142, 149, 1, 0, 0,
		0, 143, 144, 5, 13, 0, 0, 144, 149, 3, 12, 6, 7, 145, 149, 3, 14, 7, 0,
		146, 149, 3, 16, 8, 0, 147, 149, 3, 18, 9, 0, 148, 133, 1, 0, 0, 0, 148,
		138, 1, 0, 0, 0, 148, 143, 1, 0, 0, 0, 148, 145, 1, 0, 0, 0, 148, 146,
		1, 0, 0, 0, 148, 147, 1, 0, 0, 0, 149, 160, 1, 0, 0, 0, 150, 151, 10, 6,
		0, 0, 151, 152, 5, 11, 0, 0, 152, 159, 3, 12, 6, 7, 153, 154, 10, 5, 0,
		0, 154, 159, 3, 12, 6, 6, 155, 156, 10, 4, 0, 0, 156, 157, 5, 12, 0, 0,
		157, 159, 3, 12, 6, 5, 158, 150, 1, 0, 0, 0, 158, 153, 1, 0, 0, 0, 158,
		155, 1, 0, 0, 0, 159, 162, 1, 0, 0, 0, 160, 158, 1, 0, 0, 0, 160, 161,
		1, 0, 0, 0, 161, 13, 1, 0, 0, 0, 162, 160, 1, 0, 0, 0, 163, 164, 3, 20,
		10, 0, 164, 165, 3, 26, 13, 0, 165, 166, 3, 36, 18, 0, 166, 261, 1, 0,
		0, 0, 167, 168, 3, 20, 10, 0, 168, 169, 3, 28, 14, 0, 169, 170, 3, 36,
		18, 0, 170, 261, 1, 0, 0, 0, 171, 172, 3, 20, 10, 0, 172, 173, 3, 30, 15,
		0, 173, 174, 3, 36, 18, 0, 174, 261, 1, 0, 0, 0, 175, 176, 3, 20, 10, 0,
		176, 177, 3, 32, 16, 0, 177, 178, 3, 42, 21, 0, 178, 261, 1, 0, 0, 0, 179,
		180, 3, 20, 10, 0, 180, 181, 3, 34, 17, 0, 181, 182, 3, 42, 21, 0, 182,
		261, 1, 0, 0, 0, 183, 184, 5, 13, 0, 0, 184, 185, 3, 20, 10, 0, 185, 186,
		3, 26, 13, 0, 186, 187, 3, 36, 18, 0, 187, 261, 1, 0, 0, 0, 188, 189, 5,
		13, 0, 0, 189, 190, 3, 20, 10, 0, 190, 191, 3, 28, 14, 0, 191, 192, 3,
		36, 18, 0, 192, 261, 1, 0, 0, 0, 193, 194, 5, 13, 0, 0, 194, 195, 3, 20,
		10, 0, 195, 196, 3, 30, 15, 0, 196, 197, 3, 36, 18, 0, 197, 261, 1, 0,
		0, 0, 198, 199, 5, 13, 0, 0, 199, 200, 3, 20, 10, 0, 200, 201, 3, 32, 16,
		0, 201, 202, 3, 42, 21, 0, 202, 261, 1, 0, 0, 0, 203, 204, 5, 13, 0, 0,
		204, 205, 3, 20, 10, 0, 205, 206, 3, 34, 17, 0, 206, 207, 3, 42, 21, 0,
		207, 261, 1, 0, 0, 0, 208, 209, 3, 20, 10, 0, 209, 210, 5, 13, 0, 0, 210,
		211, 3, 26, 13, 0, 211, 212, 3, 36, 18, 0, 212, 261, 1, 0, 0, 0, 213, 214,
		3, 20, 10, 0, 214, 215, 5, 13, 0, 0, 215, 216, 3, 28, 14, 0, 216, 217,
		3, 36, 18, 0, 217, 261, 1, 0, 0, 0, 218, 219, 3, 20, 10, 0, 219, 220, 5,
		13, 0, 0, 220, 221, 3, 30, 15, 0, 221, 222, 3, 36, 18, 0, 222, 261, 1,
		0, 0, 0, 223, 224, 3, 20, 10, 0, 224, 225, 5, 13, 0, 0, 225, 226, 3, 32,
		16, 0, 226, 227, 3, 42, 21, 0, 227, 261, 1, 0, 0, 0, 228, 229, 3, 20, 10,
		0, 229, 230, 5, 13, 0, 0, 230, 231, 3, 34, 17, 0, 231, 232, 3, 42, 21,
		0, 232, 261, 1, 0, 0, 0, 233, 234, 3, 20, 10, 0, 234, 235, 3, 26, 13, 0,
		235, 236, 3, 38, 19, 0, 236, 261, 1, 0, 0, 0, 237, 238, 3, 20, 10, 0, 238,
		239, 3, 26, 13, 0, 239, 240, 3, 40, 20, 0, 240, 261, 1, 0, 0, 0, 241, 242,
		3, 20, 10, 0, 242, 243, 3, 32, 16, 0, 243, 244, 3, 42, 21, 0, 244, 261,
		1, 0, 0, 0, 245, 246, 3, 22, 11, 0, 246, 247, 3, 26, 13, 0, 247, 248, 3,
		36, 18, 0, 248, 261, 1, 0, 0, 0, 249, 250, 3, 22, 11, 0, 250, 251, 3, 32,
		16, 0, 251, 252, 3, 42, 21, 0, 252, 261, 1, 0, 0, 0, 253, 254, 3, 22, 11,
		0, 254, 255, 3, 26, 13, 0, 255, 256, 3, 38, 19, 0, 256, 261, 1, 0, 0, 0,
		257, 261, 5, 73, 0, 0, 258, 261, 5, 70, 0, 0, 259, 261, 5, 72, 0, 0, 260,
		163, 1, 0, 0, 0, 260, 167, 1, 0, 0, 0, 260, 171, 1, 0, 0, 0, 260, 175,
		1, 0, 0, 0, 260, 179, 1, 0, 0, 0, 260, 183, 1, 0, 0, 0, 260, 188, 1, 0,
		0, 0, 260, 193, 1, 0, 0, 0, 260, 198, 1, 0, 0, 0, 260, 203, 1, 0, 0, 0,
		260, 208, 1, 0, 0, 0, 260, 213, 1, 0, 0, 0, 260, 218, 1, 0, 0, 0, 260,
		223, 1, 0, 0, 0, 260, 228, 1, 0, 0, 0, 260, 233, 1, 0, 0, 0, 260, 237,
		1, 0, 0, 0, 260, 241, 1, 0, 0, 0, 260, 245, 1, 0, 0, 0, 260, 249, 1, 0,
		0, 0, 260, 253, 1, 0, 0, 0, 260, 257, 1, 0, 0, 0, 260, 258, 1, 0, 0, 0,
		260, 259, 1, 0, 0, 0, 261, 15, 1, 0, 0, 0, 262, 265, 5, 68, 0, 0, 263,
		265, 5, 69, 0, 0, 264, 262, 1, 0, 0, 0, 264, 263, 1, 0, 0, 0, 265, 17,
		1, 0, 0, 0, 266, 267, 5, 65, 0, 0, 267, 19, 1, 0, 0, 0, 268, 273, 3, 24,
		12, 0, 269, 270, 5, 59, 0, 0, 270, 272, 3, 24, 12, 0, 271, 269, 1, 0, 0,
		0, 272, 275, 1, 0, 0, 0, 273, 271, 1, 0, 0, 0, 273, 274, 1, 0, 0, 0, 274,
		21, 1, 0, 0, 0, 275, 273, 1, 0, 0, 0, 276, 277, 5, 28, 0, 0, 277, 278,
		5, 55, 0, 0, 278, 283, 3, 24, 12, 0, 279, 280, 5, 59, 0, 0, 280, 282, 3,
		24, 12, 0, 281, 279, 1, 0, 0, 0, 282, 285, 1, 0, 0, 0, 283, 281, 1, 0,
		0, 0, 283, 284, 1, 0, 0, 0, 284, 286, 1, 0, 0, 0, 285, 283, 1, 0, 0, 0,
		286, 287, 5, 56, 0, 0, 287, 23, 1, 0, 0, 0, 288, 289, 7, 0, 0, 0, 289,
		25, 1, 0, 0, 0, 290, 291, 7, 1, 0, 0, 291, 27, 1, 0, 0, 0, 292, 293, 7,
		2, 0, 0, 293, 29, 1, 0, 0, 0, 294, 295, 7, 3, 0, 0, 295, 31, 1, 0, 0, 0,
		296, 297, 7, 4, 0, 0, 297, 33, 1, 0, 0, 0, 298, 299, 7, 5, 0, 0, 299, 35,
		1, 0, 0, 0, 300, 301, 7, 6, 0, 0, 301, 37, 1, 0, 0, 0, 302, 303, 5, 26,
		0, 0, 303, 304, 5, 55, 0, 0, 304, 305, 3, 36, 18, 0, 305, 306, 5, 56, 0,
		0, 306, 39, 1, 0, 0, 0, 307, 308, 5, 27, 0, 0, 308, 311, 5, 55, 0, 0, 309,
		312, 5, 64, 0, 0, 310, 312, 3, 36, 18, 0, 311, 309, 1, 0, 0, 0, 311, 310,
		1, 0, 0, 0, 312, 313, 1, 0, 0, 0, 313, 314, 5, 56, 0, 0, 314, 41, 1, 0,
		0, 0, 315, 316, 5, 57, 0, 0, 316, 321, 3, 44, 22, 0, 317, 318, 5, 59, 0,
		0, 318, 320, 3, 44, 22, 0, 319, 317, 1, 0, 0, 0, 320, 323, 1, 0, 0, 0,
		321, 319, 1, 0, 0, 0, 321, 322, 1, 0, 0, 0, 322, 324, 1, 0, 0, 0, 323,
		321, 1, 0, 0, 0, 324, 325, 5, 58, 0, 0, 325, 43, 1, 0, 0, 0, 326, 329,
		3, 40, 20, 0, 327, 329, 3, 36, 18, 0, 328, 326, 1, 0, 0, 0, 328, 327, 1,
		0, 0, 0, 329, 45, 1, 0, 0, 0, 330, 331, 5, 3, 0, 0, 331, 332, 5, 55, 0,
		0, 332, 337, 3, 24, 12, 0, 333, 334, 5, 59, 0, 0, 334, 336, 3, 24, 12,
		0, 335, 333, 1, 0, 0, 0, 336, 339, 1, 0, 0, 0, 337, 335, 1, 0, 0, 0, 337,
		338, 1, 0, 0, 0, 338, 340, 1, 0, 0, 0, 339, 337, 1, 0, 0, 0, 340, 341,
		5, 56, 0, 0, 341, 47, 1, 0, 0, 0, 342, 343, 5, 4, 0, 0, 343, 344, 5, 55,
		0, 0, 344, 345, 3, 50, 25, 0, 345, 346, 5, 56, 0, 0, 346, 49, 1, 0, 0,
		0, 347, 348, 3, 56, 28, 0, 348, 349, 3, 26, 13, 0, 349, 350, 3, 36, 18,
		0, 350, 356, 1, 0, 0, 0, 351, 352, 3, 58, 29, 0, 352, 353, 3, 26, 13, 0,
		353, 354, 3, 36, 18, 0, 354, 356, 1, 0, 0, 0, 355, 347, 1, 0, 0, 0, 355,
		351, 1, 0, 0, 0, 356, 51, 1, 0, 0, 0, 357, 358, 5, 5, 0, 0, 358, 359, 5,
		55, 0, 0, 359, 360, 3, 54, 27, 0, 360, 361, 5, 56, 0, 0, 361, 53, 1, 0,
		0, 0, 362, 366, 3, 56, 28, 0, 363, 366, 3, 58, 29, 0, 364, 366, 3, 60,
		30, 0, 365, 362, 1, 0, 0, 0, 365, 363, 1, 0, 0, 0, 365, 364, 1, 0, 0, 0,
		366, 55, 1, 0, 0, 0, 367, 368, 7, 7, 0, 0, 368, 57, 1, 0, 0, 0, 369, 370,
		7, 8, 0, 0, 370, 371, 5, 60, 0, 0, 371, 372, 3, 24, 12, 0, 372, 59, 1,
		0, 0, 0, 373, 374, 7, 9, 0, 0, 374, 375, 5, 55, 0, 0, 375, 376, 5, 70,
		0, 0, 376, 379, 5, 56, 0, 0, 377, 378, 5, 60, 0, 0, 378, 380, 3, 24, 12,
		0, 379, 377, 1, 0, 0, 0, 379, 380, 1, 0, 0, 0, 380, 61, 1, 0, 0, 0, 381,
		382, 5, 6, 0, 0, 382, 383, 5, 55, 0, 0, 383, 388, 3, 64, 32, 0, 384, 385,
		5, 59, 0, 0, 385, 387, 3, 64, 32, 0, 386, 384, 1, 0, 0, 0, 387, 390, 1,
		0, 0, 0, 388, 386, 1, 0, 0, 0, 388, 389, 1, 0, 0, 0, 389, 391, 1, 0, 0,
		0, 390, 388, 1, 0, 0, 0, 391, 392, 5, 56, 0, 0, 392, 63, 1, 0, 0, 0, 393,
		397, 3, 66, 33, 0, 394, 395, 5, 63, 0, 0, 395, 398, 5, 73, 0, 0, 396, 398,
		5, 73, 0, 0, 397, 394, 1, 0, 0, 0, 397, 396, 1, 0, 0, 0, 397, 398, 1, 0,
		0, 0, 398, 65, 1, 0, 0, 0, 399, 400, 7, 10, 0, 0, 400, 67, 1, 0, 0, 0,
		401, 402, 5, 7, 0, 0, 402, 403, 5, 55, 0, 0, 403, 408, 5, 70, 0, 0, 404,
		405, 5, 59, 0, 0, 405, 407, 5, 70, 0, 0, 406, 404, 1, 0, 0, 0, 407, 410,
		1, 0, 0, 0, 408, 406, 1, 0, 0, 0, 408, 409, 1, 0, 0, 0, 409, 411, 1, 0,
		0, 0, 410, 408, 1, 0, 0, 0, 411, 412, 5, 56, 0, 0, 412, 69, 1, 0, 0, 0,
		413, 414, 5, 8, 0, 0, 414, 415, 5, 55, 0, 0, 415, 416, 7, 11, 0, 0, 416,
		417, 5, 56, 0, 0, 417, 71, 1, 0, 0, 0, 29, 73, 76, 79, 82, 85, 88, 91,
		94, 97, 116, 122, 129, 148, 158, 160, 260, 264, 273, 283, 311, 321, 328,
		337, 355, 365, 379, 388, 397, 408,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// LEQLParserInit initializes any static state used to implement LEQLParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewLEQLParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func LEQLParserInit() {
	staticData := &LEQLParserParserStaticData
	staticData.once.Do(leqlparserParserInit)
}

// NewLEQLParser produces a new parser instance for the optional input antlr.TokenStream.
func NewLEQLParser(input antlr.TokenStream) *LEQLParser {
	LEQLParserInit()
	this := new(LEQLParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &LEQLParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "LEQLParser.g4"

	return this
}

// LEQLParser tokens.
const (
	LEQLParserEOF                  = antlr.TokenEOF
	LEQLParserSELECT               = 1
	LEQLParserWHERE                = 2
	LEQLParserGROUPBY              = 3
	LEQLParserHAVING               = 4
	LEQLParserCALCULATE            = 5
	LEQLParserSORT                 = 6
	LEQLParserLIMIT                = 7
	LEQLParserTIMESLICE            = 8
	LEQLParserFROM                 = 9
	LEQLParserLOOSE                = 10
	LEQLParserAND                  = 11
	LEQLParserOR                   = 12
	LEQLParserNOT                  = 13
	LEQLParserICONTAINS_ALL_OP     = 14
	LEQLParserICONTAINS_ANY_OP     = 15
	LEQLParserCONTAINS_ALL_OP      = 16
	LEQLParserCONTAINS_ANY_OP      = 17
	LEQLParserISTARTS_WITH_ANY_OP  = 18
	LEQLParserSTARTS_WITH_ANY_OP   = 19
	LEQLParserISTARTS_WITH_OP      = 20
	LEQLParserSTARTS_WITH_OP       = 21
	LEQLParserICONTAINS_OP         = 22
	LEQLParserCONTAINS_OP          = 23
	LEQLParserIIN_OP               = 24
	LEQLParserIN_OP                = 25
	LEQLParserNOCASE               = 26
	LEQLParserIP_FUNC              = 27
	LEQLParserALL_FUNC             = 28
	LEQLParserSTANDARDDEVIATION    = 29
	LEQLParserPERCENTILE           = 30
	LEQLParserAVERAGE              = 31
	LEQLParserUNIQUE               = 32
	LEQLParserCOUNT                = 33
	LEQLParserBYTES                = 34
	LEQLParserPCTL                 = 35
	LEQLParserSUM                  = 36
	LEQLParserMIN                  = 37
	LEQLParserMAX                  = 38
	LEQLParserSD                   = 39
	LEQLParserASCENDING            = 40
	LEQLParserDESCENDING           = 41
	LEQLParserASC                  = 42
	LEQLParserDESC                 = 43
	LEQLParserSTRICT_NEQ           = 44
	LEQLParserREGEX_NEQ            = 45
	LEQLParserNEQ                  = 46
	LEQLParserREGEX_EQ             = 47
	LEQLParserGTE                  = 48
	LEQLParserLTE                  = 49
	LEQLParserSTRICT_EQ            = 50
	LEQLParserEQ                   = 51
	LEQLParserGT                   = 52
	LEQLParserLT                   = 53
	LEQLParserBANG                 = 54
	LEQLParserLPAREN               = 55
	LEQLParserRPAREN               = 56
	LEQLParserLBRACKET             = 57
	LEQLParserRBRACKET             = 58
	LEQLParserCOMMA                = 59
	LEQLParserCOLON                = 60
	LEQLParserDOT                  = 61
	LEQLParserSTAR                 = 62
	LEQLParserHASH                 = 63
	LEQLParserIP_CIDR              = 64
	LEQLParserREGEX                = 65
	LEQLParserTRIPLE_SINGLE_STRING = 66
	LEQLParserTRIPLE_DOUBLE_STRING = 67
	LEQLParserDOUBLE_STRING        = 68
	LEQLParserSINGLE_STRING        = 69
	LEQLParserNUMBER               = 70
	LEQLParserTIME_UNIT            = 71
	LEQLParserVARIABLE             = 72
	LEQLParserIDENTIFIER           = 73
	LEQLParserWS                   = 74
)

// LEQLParser rules.
const (
	LEQLParserRULE_query                 = 0
	LEQLParserRULE_fromClause            = 1
	LEQLParserRULE_selectClause          = 2
	LEQLParserRULE_selectFieldList       = 3
	LEQLParserRULE_selectField           = 4
	LEQLParserRULE_whereClause           = 5
	LEQLParserRULE_expression            = 6
	LEQLParserRULE_condition             = 7
	LEQLParserRULE_keywordSearch         = 8
	LEQLParserRULE_regexSearch           = 9
	LEQLParserRULE_fieldList             = 10
	LEQLParserRULE_allFieldList          = 11
	LEQLParserRULE_fieldName             = 12
	LEQLParserRULE_comparisonOp          = 13
	LEQLParserRULE_regexMatchOp          = 14
	LEQLParserRULE_stringOp              = 15
	LEQLParserRULE_setOp                 = 16
	LEQLParserRULE_listStringOp          = 17
	LEQLParserRULE_value                 = 18
	LEQLParserRULE_nocaseValue           = 19
	LEQLParserRULE_ipValue               = 20
	LEQLParserRULE_valueList             = 21
	LEQLParserRULE_valueListItem         = 22
	LEQLParserRULE_groupbyClause         = 23
	LEQLParserRULE_havingClause          = 24
	LEQLParserRULE_havingCondition       = 25
	LEQLParserRULE_calculateClause       = 26
	LEQLParserRULE_calcExpr              = 27
	LEQLParserRULE_calcFunction          = 28
	LEQLParserRULE_calcFunctionWithField = 29
	LEQLParserRULE_percentileFunction    = 30
	LEQLParserRULE_sortClause            = 31
	LEQLParserRULE_sortSpec              = 32
	LEQLParserRULE_sortDirection         = 33
	LEQLParserRULE_limitClause           = 34
	LEQLParserRULE_timesliceClause       = 35
)

// IQueryContext is an interface to support dynamic dispatch.
type IQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	FromClause() IFromClauseContext
	SelectClause() ISelectClauseContext
	WhereClause() IWhereClauseContext
	GroupbyClause() IGroupbyClauseContext
	CalculateClause() ICalculateClauseContext
	HavingClause() IHavingClauseContext
	SortClause() ISortClauseContext
	LimitClause() ILimitClauseContext
	TimesliceClause() ITimesliceClauseContext

	// IsQueryContext differentiates from other interfaces.
	IsQueryContext()
}

type QueryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQueryContext() *QueryContext {
	var p = new(QueryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_query
	return p
}

func InitEmptyQueryContext(p *QueryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_query
}

func (*QueryContext) IsQueryContext() {}

func NewQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryContext {
	var p = new(QueryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LEQLParserRULE_query

	return p
}

func (s *QueryContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryContext) EOF() antlr.TerminalNode {
	return s.GetToken(LEQLParserEOF, 0)
}

func (s *QueryContext) FromClause() IFromClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFromClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFromClauseContext)
}

func (s *QueryContext) SelectClause() ISelectClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectClauseContext)
}

func (s *QueryContext) WhereClause() IWhereClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhereClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhereClauseContext)
}

func (s *QueryContext) GroupbyClause() IGroupbyClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGroupbyClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGroupbyClauseContext)
}

func (s *QueryContext) CalculateClause() ICalculateClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICalculateClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICalculateClauseContext)
}

func (s *QueryContext) HavingClause() IHavingClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHavingClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHavingClauseContext)
}

func (s *QueryContext) SortClause() ISortClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISortClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISortClauseContext)
}

func (s *QueryContext) LimitClause() ILimitClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILimitClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILimitClauseContext)
}

func (s *QueryContext) TimesliceClause() ITimesliceClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITimesliceClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITimesliceClauseContext)
}

func (s *QueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterQuery(s)
	}
}

func (s *QueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitQuery(s)
	}
}

func (s *QueryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitQuery(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LEQLParser) Query() (localctx IQueryContext) {
	localctx = NewQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, LEQLParserRULE_query)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == LEQLParserFROM {
		{
			p.SetState(72)
			p.FromClause()
		}

	}
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == LEQLParserSELECT {
		{
			p.SetState(75)
			p.SelectClause()
		}

	}
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == LEQLParserWHERE {
		{
			p.SetState(78)
			p.WhereClause()
		}

	}
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == LEQLParserGROUPBY {
		{
			p.SetState(81)
			p.GroupbyClause()
		}

	}
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == LEQLParserCALCULATE {
		{
			p.SetState(84)
			p.CalculateClause()
		}

	}
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == LEQLParserHAVING {
		{
			p.SetState(87)
			p.HavingClause()
		}

	}
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == LEQLParserSORT {
		{
			p.SetState(90)
			p.SortClause()
		}

	}
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == LEQLParserLIMIT {
		{
			p.SetState(93)
			p.LimitClause()
		}

	}
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == LEQLParserTIMESLICE {
		{
			p.SetState(96)
			p.TimesliceClause()
		}

	}
	{
		p.SetState(99)
		p.Match(LEQLParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFromClauseContext is an interface to support dynamic dispatch.
type IFromClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FROM() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode

	// IsFromClauseContext differentiates from other interfaces.
	IsFromClauseContext()
}

type FromClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFromClauseContext() *FromClauseContext {
	var p = new(FromClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_fromClause
	return p
}

func InitEmptyFromClauseContext(p *FromClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_fromClause
}

func (*FromClauseContext) IsFromClauseContext() {}

func NewFromClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FromClauseContext {
	var p = new(FromClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LEQLParserRULE_fromClause

	return p
}

func (s *FromClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *FromClauseContext) FROM() antlr.TerminalNode {
	return s.GetToken(LEQLParserFROM, 0)
}

func (s *FromClauseContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(LEQLParserLPAREN, 0)
}

func (s *FromClauseContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FromClauseContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(LEQLParserRPAREN, 0)
}

func (s *FromClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FromClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FromClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterFromClause(s)
	}
}

func (s *FromClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitFromClause(s)
	}
}

func (s *FromClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitFromClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LEQLParser) FromClause() (localctx IFromClauseContext) {
	localctx = NewFromClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, LEQLParserRULE_fromClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(101)
		p.Match(LEQLParserFROM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(102)
		p.Match(LEQLParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(103)
		p.expression(0)
	}
	{
		p.SetState(104)
		p.Match(LEQLParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISelectClauseContext is an interface to support dynamic dispatch.
type ISelectClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SELECT() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	SelectFieldList() ISelectFieldListContext
	RPAREN() antlr.TerminalNode

	// IsSelectClauseContext differentiates from other interfaces.
	IsSelectClauseContext()
}

type SelectClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectClauseContext() *SelectClauseContext {
	var p = new(SelectClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_selectClause
	return p
}

func InitEmptySelectClauseContext(p *SelectClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_selectClause
}

func (*SelectClauseContext) IsSelectClauseContext() {}

func NewSelectClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectClauseContext {
	var p = new(SelectClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LEQLParserRULE_selectClause

	return p
}

func (s *SelectClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectClauseContext) SELECT() antlr.TerminalNode {
	return s.GetToken(LEQLParserSELECT, 0)
}

func (s *SelectClauseContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(LEQLParserLPAREN, 0)
}

func (s *SelectClauseContext) SelectFieldList() ISelectFieldListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectFieldListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectFieldListContext)
}

func (s *SelectClauseContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(LEQLParserRPAREN, 0)
}

func (s *SelectClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterSelectClause(s)
	}
}

func (s *SelectClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitSelectClause(s)
	}
}

func (s *SelectClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitSelectClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LEQLParser) SelectClause() (localctx ISelectClauseContext) {
	localctx = NewSelectClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, LEQLParserRULE_selectClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(106)
		p.Match(LEQLParserSELECT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(107)
		p.Match(LEQLParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(108)
		p.SelectFieldList()
	}
	{
		p.SetState(109)
		p.Match(LEQLParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISelectFieldListContext is an interface to support dynamic dispatch.
type ISelectFieldListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSelectField() []ISelectFieldContext
	SelectField(i int) ISelectFieldContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsSelectFieldListContext differentiates from other interfaces.
	IsSelectFieldListContext()
}

type SelectFieldListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectFieldListContext() *SelectFieldListContext {
	var p = new(SelectFieldListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_selectFieldList
	return p
}

func InitEmptySelectFieldListContext(p *SelectFieldListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_selectFieldList
}

func (*SelectFieldListContext) IsSelectFieldListContext() {}

func NewSelectFieldListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectFieldListContext {
	var p = new(SelectFieldListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LEQLParserRULE_selectFieldList

	return p
}

func (s *SelectFieldListContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectFieldListContext) AllSelectField() []ISelectFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISelectFieldContext); ok {
			len++
		}
	}

	tst := make([]ISelectFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISelectFieldContext); ok {
			tst[i] = t.(ISelectFieldContext)
			i++
		}
	}

	return tst
}

func (s *SelectFieldListContext) SelectField(i int) ISelectFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectFieldContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectFieldContext)
}

func (s *SelectFieldListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(LEQLParserCOMMA)
}

func (s *SelectFieldListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(LEQLParserCOMMA, i)
}

func (s *SelectFieldListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectFieldListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectFieldListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterSelectFieldList(s)
	}
}

func (s *SelectFieldListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitSelectFieldList(s)
	}
}

func (s *SelectFieldListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitSelectFieldList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LEQLParser) SelectFieldList() (localctx ISelectFieldListContext) {
	localctx = NewSelectFieldListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, LEQLParserRULE_selectFieldList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(111)
		p.SelectField()
	}
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == LEQLParserCOMMA {
		{
			p.SetState(112)
			p.Match(LEQLParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(113)
			p.SelectField()
		}

		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISelectFieldContext is an interface to support dynamic dispatch.
type ISelectFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FieldName() IFieldNameContext
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode

	// IsSelectFieldContext differentiates from other interfaces.
	IsSelectFieldContext()
}

type SelectFieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectFieldContext() *SelectFieldContext {
	var p = new(SelectFieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_selectField
	return p
}

func InitEmptySelectFieldContext(p *SelectFieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_selectField
}

func (*SelectFieldContext) IsSelectFieldContext() {}

func NewSelectFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectFieldContext {
	var p = new(SelectFieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LEQLParserRULE_selectField

	return p
}

func (s *SelectFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectFieldContext) FieldName() IFieldNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldNameContext)
}

func (s *SelectFieldContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(LEQLParserIDENTIFIER)
}

func (s *SelectFieldContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(LEQLParserIDENTIFIER, i)
}

func (s *SelectFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterSelectField(s)
	}
}

func (s *SelectFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitSelectField(s)
	}
}

func (s *SelectFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitSelectField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LEQLParser) SelectField() (localctx ISelectFieldContext) {
	localctx = NewSelectFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, LEQLParserRULE_selectField)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(119)
		p.FieldName()
	}
	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == LEQLParserIDENTIFIER {
		{
			p.SetState(120)
			p.Match(LEQLParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(121)
			p.Match(LEQLParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWhereClauseContext is an interface to support dynamic dispatch.
type IWhereClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WHERE() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode
	COMMA() antlr.TerminalNode
	LOOSE() antlr.TerminalNode

	// IsWhereClauseContext differentiates from other interfaces.
	IsWhereClauseContext()
}

type WhereClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhereClauseContext() *WhereClauseContext {
	var p = new(WhereClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_whereClause
	return p
}

func InitEmptyWhereClauseContext(p *WhereClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_whereClause
}

func (*WhereClauseContext) IsWhereClauseContext() {}

func NewWhereClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhereClauseContext {
	var p = new(WhereClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LEQLParserRULE_whereClause

	return p
}

func (s *WhereClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *WhereClauseContext) WHERE() antlr.TerminalNode {
	return s.GetToken(LEQLParserWHERE, 0)
}

func (s *WhereClauseContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(LEQLParserLPAREN, 0)
}

func (s *WhereClauseContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *WhereClauseContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(LEQLParserRPAREN, 0)
}

func (s *WhereClauseContext) COMMA() antlr.TerminalNode {
	return s.GetToken(LEQLParserCOMMA, 0)
}

func (s *WhereClauseContext) LOOSE() antlr.TerminalNode {
	return s.GetToken(LEQLParserLOOSE, 0)
}

func (s *WhereClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhereClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhereClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterWhereClause(s)
	}
}

func (s *WhereClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitWhereClause(s)
	}
}

func (s *WhereClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitWhereClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LEQLParser) WhereClause() (localctx IWhereClauseContext) {
	localctx = NewWhereClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, LEQLParserRULE_whereClause)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(124)
		p.Match(LEQLParserWHERE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(125)
		p.Match(LEQLParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(126)
		p.expression(0)
	}
	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == LEQLParserCOMMA {
		{
			p.SetState(127)
			p.Match(LEQLParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(128)
			p.Match(LEQLParserLOOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(131)
		p.Match(LEQLParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LEQLParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyAll(ctx *ExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NotExprContext struct {
	ExpressionContext
}

func NewNotExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotExprContext {
	var p = new(NotExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *NotExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotExprContext) NOT() antlr.TerminalNode {
	return s.GetToken(LEQLParserNOT, 0)
}

func (s *NotExprContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *NotExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterNotExpr(s)
	}
}

func (s *NotExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitNotExpr(s)
	}
}

func (s *NotExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitNotExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ConditionExprContext struct {
	ExpressionContext
}

func NewConditionExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConditionExprContext {
	var p = new(ConditionExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ConditionExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionExprContext) Condition() IConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *ConditionExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterConditionExpr(s)
	}
}

func (s *ConditionExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitConditionExpr(s)
	}
}

func (s *ConditionExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitConditionExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type OrExprContext struct {
	ExpressionContext
}

func NewOrExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrExprContext {
	var p = new(OrExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *OrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *OrExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *OrExprContext) OR() antlr.TerminalNode {
	return s.GetToken(LEQLParserOR, 0)
}

func (s *OrExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterOrExpr(s)
	}
}

func (s *OrExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitOrExpr(s)
	}
}

func (s *OrExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitOrExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type RegexExprContext struct {
	ExpressionContext
}

func NewRegexExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RegexExprContext {
	var p = new(RegexExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *RegexExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegexExprContext) RegexSearch() IRegexSearchContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegexSearchContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRegexSearchContext)
}

func (s *RegexExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterRegexExpr(s)
	}
}

func (s *RegexExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitRegexExpr(s)
	}
}

func (s *RegexExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitRegexExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ImplicitAndExprContext struct {
	ExpressionContext
}

func NewImplicitAndExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ImplicitAndExprContext {
	var p = new(ImplicitAndExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ImplicitAndExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImplicitAndExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ImplicitAndExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ImplicitAndExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterImplicitAndExpr(s)
	}
}

func (s *ImplicitAndExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitImplicitAndExpr(s)
	}
}

func (s *ImplicitAndExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitImplicitAndExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type NestedWhereExprContext struct {
	ExpressionContext
}

func NewNestedWhereExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NestedWhereExprContext {
	var p = new(NestedWhereExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *NestedWhereExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NestedWhereExprContext) WHERE() antlr.TerminalNode {
	return s.GetToken(LEQLParserWHERE, 0)
}

func (s *NestedWhereExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(LEQLParserLPAREN, 0)
}

func (s *NestedWhereExprContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *NestedWhereExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(LEQLParserRPAREN, 0)
}

func (s *NestedWhereExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterNestedWhereExpr(s)
	}
}

func (s *NestedWhereExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitNestedWhereExpr(s)
	}
}

func (s *NestedWhereExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitNestedWhereExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type KeywordExprContext struct {
	ExpressionContext
}

func NewKeywordExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *KeywordExprContext {
	var p = new(KeywordExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *KeywordExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeywordExprContext) KeywordSearch() IKeywordSearchContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeywordSearchContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeywordSearchContext)
}

func (s *KeywordExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterKeywordExpr(s)
	}
}

func (s *KeywordExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitKeywordExpr(s)
	}
}

func (s *KeywordExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitKeywordExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParenExprContext struct {
	ExpressionContext
}

func NewParenExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenExprContext {
	var p = new(ParenExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ParenExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(LEQLParserLPAREN, 0)
}

func (s *ParenExprContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParenExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(LEQLParserRPAREN, 0)
}

func (s *ParenExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterParenExpr(s)
	}
}

func (s *ParenExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitParenExpr(s)
	}
}

func (s *ParenExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitParenExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AndExprContext struct {
	ExpressionContext
}

func NewAndExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndExprContext {
	var p = new(AndExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *AndExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AndExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AndExprContext) AND() antlr.TerminalNode {
	return s.GetToken(LEQLParserAND, 0)
}

func (s *AndExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterAndExpr(s)
	}
}

func (s *AndExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitAndExpr(s)
	}
}

func (s *AndExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitAndExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LEQLParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *LEQLParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 12
	p.EnterRecursionRule(localctx, 12, LEQLParserRULE_expression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		localctx = NewParenExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(134)
			p.Match(LEQLParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(135)
			p.expression(0)
		}
		{
			p.SetState(136)
			p.Match(LEQLParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewNestedWhereExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(138)
			p.Match(LEQLParserWHERE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(139)
			p.Match(LEQLParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(140)
			p.expression(0)
		}
		{
			p.SetState(141)
			p.Match(LEQLParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewNotExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(143)
			p.Match(LEQLParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(144)
			p.expression(7)
		}

	case 4:
		localctx = NewConditionExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(145)
			p.Condition()
		}

	case 5:
		localctx = NewKeywordExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(146)
			p.KeywordSearch()
		}

	case 6:
		localctx = NewRegexExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(147)
			p.RegexSearch()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(160)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(158)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAndExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LEQLParserRULE_expression)
				p.SetState(150)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(151)
					p.Match(LEQLParserAND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(152)
					p.expression(7)
				}

			case 2:
				localctx = NewImplicitAndExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LEQLParserRULE_expression)
				p.SetState(153)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(154)
					p.expression(6)
				}

			case 3:
				localctx = NewOrExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LEQLParserRULE_expression)
				p.SetState(155)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(156)
					p.Match(LEQLParserOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(157)
					p.expression(5)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(162)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConditionContext is an interface to support dynamic dispatch.
type IConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsConditionContext differentiates from other interfaces.
	IsConditionContext()
}

type ConditionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionContext() *ConditionContext {
	var p = new(ConditionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_condition
	return p
}

func InitEmptyConditionContext(p *ConditionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_condition
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LEQLParserRULE_condition

	return p
}

func (s *ConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionContext) CopyAll(ctx *ConditionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IpSetConditionContext struct {
	ConditionContext
}

func NewIpSetConditionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IpSetConditionContext {
	var p = new(IpSetConditionContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *IpSetConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IpSetConditionContext) FieldList() IFieldListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldListContext)
}

func (s *IpSetConditionContext) SetOp() ISetOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISetOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISetOpContext)
}

func (s *IpSetConditionContext) ValueList() IValueListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueListContext)
}

func (s *IpSetConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterIpSetCondition(s)
	}
}

func (s *IpSetConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitIpSetCondition(s)
	}
}

func (s *IpSetConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitIpSetCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

type PostfixNegatedStringConditionContext struct {
	ConditionContext
}

func NewPostfixNegatedStringConditionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PostfixNegatedStringConditionContext {
	var p = new(PostfixNegatedStringConditionContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *PostfixNegatedStringConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PostfixNegatedStringConditionContext) FieldList() IFieldListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldListContext)
}

func (s *PostfixNegatedStringConditionContext) NOT() antlr.TerminalNode {
	return s.GetToken(LEQLParserNOT, 0)
}

func (s *PostfixNegatedStringConditionContext) StringOp() IStringOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringOpContext)
}

func (s *PostfixNegatedStringConditionContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *PostfixNegatedStringConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterPostfixNegatedStringCondition(s)
	}
}

func (s *PostfixNegatedStringConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitPostfixNegatedStringCondition(s)
	}
}

func (s *PostfixNegatedStringConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitPostfixNegatedStringCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

type NegatedSetConditionContext struct {
	ConditionContext
}

func NewNegatedSetConditionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NegatedSetConditionContext {
	var p = new(NegatedSetConditionContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *NegatedSetConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NegatedSetConditionContext) NOT() antlr.TerminalNode {
	return s.GetToken(LEQLParserNOT, 0)
}

func (s *NegatedSetConditionContext) FieldList() IFieldListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldListContext)
}

func (s *NegatedSetConditionContext) SetOp() ISetOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISetOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISetOpContext)
}

func (s *NegatedSetConditionContext) ValueList() IValueListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueListContext)
}

func (s *NegatedSetConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterNegatedSetCondition(s)
	}
}

func (s *NegatedSetConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitNegatedSetCondition(s)
	}
}

func (s *NegatedSetConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitNegatedSetCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

type AllFieldsNocaseConditionContext struct {
	ConditionContext
}

func NewAllFieldsNocaseConditionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AllFieldsNocaseConditionContext {
	var p = new(AllFieldsNocaseConditionContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *AllFieldsNocaseConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AllFieldsNocaseConditionContext) AllFieldList() IAllFieldListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAllFieldListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAllFieldListContext)
}

func (s *AllFieldsNocaseConditionContext) ComparisonOp() IComparisonOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparisonOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparisonOpContext)
}

func (s *AllFieldsNocaseConditionContext) NocaseValue() INocaseValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INocaseValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INocaseValueContext)
}

func (s *AllFieldsNocaseConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterAllFieldsNocaseCondition(s)
	}
}

func (s *AllFieldsNocaseConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitAllFieldsNocaseCondition(s)
	}
}

func (s *AllFieldsNocaseConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitAllFieldsNocaseCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

type FieldExistsConditionContext struct {
	ConditionContext
}

func NewFieldExistsConditionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FieldExistsConditionContext {
	var p = new(FieldExistsConditionContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *FieldExistsConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldExistsConditionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(LEQLParserIDENTIFIER, 0)
}

func (s *FieldExistsConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterFieldExistsCondition(s)
	}
}

func (s *FieldExistsConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitFieldExistsCondition(s)
	}
}

func (s *FieldExistsConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitFieldExistsCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

type VariableConditionContext struct {
	ConditionContext
}

func NewVariableConditionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VariableConditionContext {
	var p = new(VariableConditionContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *VariableConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableConditionContext) VARIABLE() antlr.TerminalNode {
	return s.GetToken(LEQLParserVARIABLE, 0)
}

func (s *VariableConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterVariableCondition(s)
	}
}

func (s *VariableConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitVariableCondition(s)
	}
}

func (s *VariableConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitVariableCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

type AllFieldsSetConditionContext struct {
	ConditionContext
}

func NewAllFieldsSetConditionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AllFieldsSetConditionContext {
	var p = new(AllFieldsSetConditionContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *AllFieldsSetConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AllFieldsSetConditionContext) AllFieldList() IAllFieldListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAllFieldListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAllFieldListContext)
}

func (s *AllFieldsSetConditionContext) SetOp() ISetOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISetOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISetOpContext)
}

func (s *AllFieldsSetConditionContext) ValueList() IValueListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueListContext)
}

func (s *AllFieldsSetConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterAllFieldsSetCondition(s)
	}
}

func (s *AllFieldsSetConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitAllFieldsSetCondition(s)
	}
}

func (s *AllFieldsSetConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitAllFieldsSetCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

type PostfixNegatedComparisonConditionContext struct {
	ConditionContext
}

func NewPostfixNegatedComparisonConditionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PostfixNegatedComparisonConditionContext {
	var p = new(PostfixNegatedComparisonConditionContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *PostfixNegatedComparisonConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PostfixNegatedComparisonConditionContext) FieldList() IFieldListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldListContext)
}

func (s *PostfixNegatedComparisonConditionContext) NOT() antlr.TerminalNode {
	return s.GetToken(LEQLParserNOT, 0)
}

func (s *PostfixNegatedComparisonConditionContext) ComparisonOp() IComparisonOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparisonOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparisonOpContext)
}

func (s *PostfixNegatedComparisonConditionContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *PostfixNegatedComparisonConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterPostfixNegatedComparisonCondition(s)
	}
}

func (s *PostfixNegatedComparisonConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitPostfixNegatedComparisonCondition(s)
	}
}

func (s *PostfixNegatedComparisonConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitPostfixNegatedComparisonCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

type ListStringConditionContext struct {
	ConditionContext
}

func NewListStringConditionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListStringConditionContext {
	var p = new(ListStringConditionContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *ListStringConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListStringConditionContext) FieldList() IFieldListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldListContext)
}

func (s *ListStringConditionContext) ListStringOp() IListStringOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListStringOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListStringOpContext)
}

func (s *ListStringConditionContext) ValueList() IValueListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueListContext)
}

func (s *ListStringConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterListStringCondition(s)
	}
}

func (s *ListStringConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitListStringCondition(s)
	}
}

func (s *ListStringConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitListStringCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

type NegatedStringConditionContext struct {
	ConditionContext
}

func NewNegatedStringConditionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NegatedStringConditionContext {
	var p = new(NegatedStringConditionContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *NegatedStringConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NegatedStringConditionContext) NOT() antlr.TerminalNode {
	return s.GetToken(LEQLParserNOT, 0)
}

func (s *NegatedStringConditionContext) FieldList() IFieldListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldListContext)
}

func (s *NegatedStringConditionContext) StringOp() IStringOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringOpContext)
}

func (s *NegatedStringConditionContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *NegatedStringConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterNegatedStringCondition(s)
	}
}

func (s *NegatedStringConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitNegatedStringCondition(s)
	}
}

func (s *NegatedStringConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitNegatedStringCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

type SetConditionContext struct {
	ConditionContext
}

func NewSetConditionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SetConditionContext {
	var p = new(SetConditionContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *SetConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetConditionContext) FieldList() IFieldListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldListContext)
}

func (s *SetConditionContext) SetOp() ISetOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISetOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISetOpContext)
}

func (s *SetConditionContext) ValueList() IValueListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueListContext)
}

func (s *SetConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterSetCondition(s)
	}
}

func (s *SetConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitSetCondition(s)
	}
}

func (s *SetConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitSetCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

type RegexMatchConditionContext struct {
	ConditionContext
}

func NewRegexMatchConditionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RegexMatchConditionContext {
	var p = new(RegexMatchConditionContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *RegexMatchConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegexMatchConditionContext) FieldList() IFieldListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldListContext)
}

func (s *RegexMatchConditionContext) RegexMatchOp() IRegexMatchOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegexMatchOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRegexMatchOpContext)
}

func (s *RegexMatchConditionContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *RegexMatchConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterRegexMatchCondition(s)
	}
}

func (s *RegexMatchConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitRegexMatchCondition(s)
	}
}

func (s *RegexMatchConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitRegexMatchCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

type NocaseConditionContext struct {
	ConditionContext
}

func NewNocaseConditionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NocaseConditionContext {
	var p = new(NocaseConditionContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *NocaseConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NocaseConditionContext) FieldList() IFieldListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldListContext)
}

func (s *NocaseConditionContext) ComparisonOp() IComparisonOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparisonOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparisonOpContext)
}

func (s *NocaseConditionContext) NocaseValue() INocaseValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INocaseValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INocaseValueContext)
}

func (s *NocaseConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterNocaseCondition(s)
	}
}

func (s *NocaseConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitNocaseCondition(s)
	}
}

func (s *NocaseConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitNocaseCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

type PostfixNegatedRegexMatchConditionContext struct {
	ConditionContext
}

func NewPostfixNegatedRegexMatchConditionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PostfixNegatedRegexMatchConditionContext {
	var p = new(PostfixNegatedRegexMatchConditionContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *PostfixNegatedRegexMatchConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PostfixNegatedRegexMatchConditionContext) FieldList() IFieldListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldListContext)
}

func (s *PostfixNegatedRegexMatchConditionContext) NOT() antlr.TerminalNode {
	return s.GetToken(LEQLParserNOT, 0)
}

func (s *PostfixNegatedRegexMatchConditionContext) RegexMatchOp() IRegexMatchOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegexMatchOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRegexMatchOpContext)
}

func (s *PostfixNegatedRegexMatchConditionContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *PostfixNegatedRegexMatchConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterPostfixNegatedRegexMatchCondition(s)
	}
}

func (s *PostfixNegatedRegexMatchConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitPostfixNegatedRegexMatchCondition(s)
	}
}

func (s *PostfixNegatedRegexMatchConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitPostfixNegatedRegexMatchCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

type NegatedListStringConditionContext struct {
	ConditionContext
}

func NewNegatedListStringConditionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NegatedListStringConditionContext {
	var p = new(NegatedListStringConditionContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *NegatedListStringConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NegatedListStringConditionContext) NOT() antlr.TerminalNode {
	return s.GetToken(LEQLParserNOT, 0)
}

func (s *NegatedListStringConditionContext) FieldList() IFieldListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldListContext)
}

func (s *NegatedListStringConditionContext) ListStringOp() IListStringOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListStringOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListStringOpContext)
}

func (s *NegatedListStringConditionContext) ValueList() IValueListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueListContext)
}

func (s *NegatedListStringConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterNegatedListStringCondition(s)
	}
}

func (s *NegatedListStringConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitNegatedListStringCondition(s)
	}
}

func (s *NegatedListStringConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitNegatedListStringCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

type NegatedComparisonConditionContext struct {
	ConditionContext
}

func NewNegatedComparisonConditionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NegatedComparisonConditionContext {
	var p = new(NegatedComparisonConditionContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *NegatedComparisonConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NegatedComparisonConditionContext) NOT() antlr.TerminalNode {
	return s.GetToken(LEQLParserNOT, 0)
}

func (s *NegatedComparisonConditionContext) FieldList() IFieldListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldListContext)
}

func (s *NegatedComparisonConditionContext) ComparisonOp() IComparisonOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparisonOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparisonOpContext)
}

func (s *NegatedComparisonConditionContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *NegatedComparisonConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterNegatedComparisonCondition(s)
	}
}

func (s *NegatedComparisonConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitNegatedComparisonCondition(s)
	}
}

func (s *NegatedComparisonConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitNegatedComparisonCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringConditionContext struct {
	ConditionContext
}

func NewStringConditionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringConditionContext {
	var p = new(StringConditionContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *StringConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringConditionContext) FieldList() IFieldListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldListContext)
}

func (s *StringConditionContext) StringOp() IStringOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringOpContext)
}

func (s *StringConditionContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *StringConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterStringCondition(s)
	}
}

func (s *StringConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitStringCondition(s)
	}
}

func (s *StringConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitStringCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

type AllFieldsComparisonConditionContext struct {
	ConditionContext
}

func NewAllFieldsComparisonConditionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AllFieldsComparisonConditionContext {
	var p = new(AllFieldsComparisonConditionContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *AllFieldsComparisonConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AllFieldsComparisonConditionContext) AllFieldList() IAllFieldListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAllFieldListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAllFieldListContext)
}

func (s *AllFieldsComparisonConditionContext) ComparisonOp() IComparisonOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparisonOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparisonOpContext)
}

func (s *AllFieldsComparisonConditionContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *AllFieldsComparisonConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterAllFieldsComparisonCondition(s)
	}
}

func (s *AllFieldsComparisonConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitAllFieldsComparisonCondition(s)
	}
}

func (s *AllFieldsComparisonConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitAllFieldsComparisonCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

type NegatedRegexMatchConditionContext struct {
	ConditionContext
}

func NewNegatedRegexMatchConditionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NegatedRegexMatchConditionContext {
	var p = new(NegatedRegexMatchConditionContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *NegatedRegexMatchConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NegatedRegexMatchConditionContext) NOT() antlr.TerminalNode {
	return s.GetToken(LEQLParserNOT, 0)
}

func (s *NegatedRegexMatchConditionContext) FieldList() IFieldListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldListContext)
}

func (s *NegatedRegexMatchConditionContext) RegexMatchOp() IRegexMatchOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegexMatchOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRegexMatchOpContext)
}

func (s *NegatedRegexMatchConditionContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *NegatedRegexMatchConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterNegatedRegexMatchCondition(s)
	}
}

func (s *NegatedRegexMatchConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitNegatedRegexMatchCondition(s)
	}
}

func (s *NegatedRegexMatchConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitNegatedRegexMatchCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

type PostfixNegatedSetConditionContext struct {
	ConditionContext
}

func NewPostfixNegatedSetConditionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PostfixNegatedSetConditionContext {
	var p = new(PostfixNegatedSetConditionContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *PostfixNegatedSetConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PostfixNegatedSetConditionContext) FieldList() IFieldListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldListContext)
}

func (s *PostfixNegatedSetConditionContext) NOT() antlr.TerminalNode {
	return s.GetToken(LEQLParserNOT, 0)
}

func (s *PostfixNegatedSetConditionContext) SetOp() ISetOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISetOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISetOpContext)
}

func (s *PostfixNegatedSetConditionContext) ValueList() IValueListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueListContext)
}

func (s *PostfixNegatedSetConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterPostfixNegatedSetCondition(s)
	}
}

func (s *PostfixNegatedSetConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitPostfixNegatedSetCondition(s)
	}
}

func (s *PostfixNegatedSetConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitPostfixNegatedSetCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

type ComparisonConditionContext struct {
	ConditionContext
}

func NewComparisonConditionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ComparisonConditionContext {
	var p = new(ComparisonConditionContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *ComparisonConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonConditionContext) FieldList() IFieldListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldListContext)
}

func (s *ComparisonConditionContext) ComparisonOp() IComparisonOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparisonOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparisonOpContext)
}

func (s *ComparisonConditionContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ComparisonConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterComparisonCondition(s)
	}
}

func (s *ComparisonConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitComparisonCondition(s)
	}
}

func (s *ComparisonConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitComparisonCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

type NumberKeywordConditionContext struct {
	ConditionContext
}

func NewNumberKeywordConditionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberKeywordConditionContext {
	var p = new(NumberKeywordConditionContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *NumberKeywordConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberKeywordConditionContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(LEQLParserNUMBER, 0)
}

func (s *NumberKeywordConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterNumberKeywordCondition(s)
	}
}

func (s *NumberKeywordConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitNumberKeywordCondition(s)
	}
}

func (s *NumberKeywordConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitNumberKeywordCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

type IpConditionContext struct {
	ConditionContext
}

func NewIpConditionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IpConditionContext {
	var p = new(IpConditionContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *IpConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IpConditionContext) FieldList() IFieldListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldListContext)
}

func (s *IpConditionContext) ComparisonOp() IComparisonOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparisonOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparisonOpContext)
}

func (s *IpConditionContext) IpValue() IIpValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIpValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIpValueContext)
}

func (s *IpConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterIpCondition(s)
	}
}

func (s *IpConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitIpCondition(s)
	}
}

func (s *IpConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitIpCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

type PostfixNegatedListStringConditionContext struct {
	ConditionContext
}

func NewPostfixNegatedListStringConditionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PostfixNegatedListStringConditionContext {
	var p = new(PostfixNegatedListStringConditionContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *PostfixNegatedListStringConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PostfixNegatedListStringConditionContext) FieldList() IFieldListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldListContext)
}

func (s *PostfixNegatedListStringConditionContext) NOT() antlr.TerminalNode {
	return s.GetToken(LEQLParserNOT, 0)
}

func (s *PostfixNegatedListStringConditionContext) ListStringOp() IListStringOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListStringOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListStringOpContext)
}

func (s *PostfixNegatedListStringConditionContext) ValueList() IValueListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueListContext)
}

func (s *PostfixNegatedListStringConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterPostfixNegatedListStringCondition(s)
	}
}

func (s *PostfixNegatedListStringConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitPostfixNegatedListStringCondition(s)
	}
}

func (s *PostfixNegatedListStringConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitPostfixNegatedListStringCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LEQLParser) Condition() (localctx IConditionContext) {
	localctx = NewConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, LEQLParserRULE_condition)
	p.SetState(260)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		localctx = NewComparisonConditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(163)
			p.FieldList()
		}
		{
			p.SetState(164)
			p.ComparisonOp()
		}
		{
			p.SetState(165)
			p.Value()
		}

	case 2:
		localctx = NewRegexMatchConditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(167)
			p.FieldList()
		}
		{
			p.SetState(168)
			p.RegexMatchOp()
		}
		{
			p.SetState(169)
			p.Value()
		}

	case 3:
		localctx = NewStringConditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(171)
			p.FieldList()
		}
		{
			p.SetState(172)
			p.StringOp()
		}
		{
			p.SetState(173)
			p.Value()
		}

	case 4:
		localctx = NewSetConditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(175)
			p.FieldList()
		}
		{
			p.SetState(176)
			p.SetOp()
		}
		{
			p.SetState(177)
			p.ValueList()
		}

	case 5:
		localctx = NewListStringConditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(179)
			p.FieldList()
		}
		{
			p.SetState(180)
			p.ListStringOp()
		}
		{
			p.SetState(181)
			p.ValueList()
		}

	case 6:
		localctx = NewNegatedComparisonConditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(183)
			p.Match(LEQLParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(184)
			p.FieldList()
		}
		{
			p.SetState(185)
			p.ComparisonOp()
		}
		{
			p.SetState(186)
			p.Value()
		}

	case 7:
		localctx = NewNegatedRegexMatchConditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(188)
			p.Match(LEQLParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(189)
			p.FieldList()
		}
		{
			p.SetState(190)
			p.RegexMatchOp()
		}
		{
			p.SetState(191)
			p.Value()
		}

	case 8:
		localctx = NewNegatedStringConditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(193)
			p.Match(LEQLParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(194)
			p.FieldList()
		}
		{
			p.SetState(195)
			p.StringOp()
		}
		{
			p.SetState(196)
			p.Value()
		}

	case 9:
		localctx = NewNegatedSetConditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(198)
			p.Match(LEQLParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(199)
			p.FieldList()
		}
		{
			p.SetState(200)
			p.SetOp()
		}
		{
			p.SetState(201)
			p.ValueList()
		}

	case 10:
		localctx = NewNegatedListStringConditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(203)
			p.Match(LEQLParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(204)
			p.FieldList()
		}
		{
			p.SetState(205)
			p.ListStringOp()
		}
		{
			p.SetState(206)
			p.ValueList()
		}

	case 11:
		localctx = NewPostfixNegatedComparisonConditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(208)
			p.FieldList()
		}
		{
			p.SetState(209)
			p.Match(LEQLParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(210)
			p.ComparisonOp()
		}
		{
			p.SetState(211)
			p.Value()
		}

	case 12:
		localctx = NewPostfixNegatedRegexMatchConditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(213)
			p.FieldList()
		}
		{
			p.SetState(214)
			p.Match(LEQLParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(215)
			p.RegexMatchOp()
		}
		{
			p.SetState(216)
			p.Value()
		}

	case 13:
		localctx = NewPostfixNegatedStringConditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(218)
			p.FieldList()
		}
		{
			p.SetState(219)
			p.Match(LEQLParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(220)
			p.StringOp()
		}
		{
			p.SetState(221)
			p.Value()
		}

	case 14:
		localctx = NewPostfixNegatedSetConditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(223)
			p.FieldList()
		}
		{
			p.SetState(224)
			p.Match(LEQLParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(225)
			p.SetOp()
		}
		{
			p.SetState(226)
			p.ValueList()
		}

	case 15:
		localctx = NewPostfixNegatedListStringConditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(228)
			p.FieldList()
		}
		{
			p.SetState(229)
			p.Match(LEQLParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(230)
			p.ListStringOp()
		}
		{
			p.SetState(231)
			p.ValueList()
		}

	case 16:
		localctx = NewNocaseConditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(233)
			p.FieldList()
		}
		{
			p.SetState(234)
			p.ComparisonOp()
		}
		{
			p.SetState(235)
			p.NocaseValue()
		}

	case 17:
		localctx = NewIpConditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(237)
			p.FieldList()
		}
		{
			p.SetState(238)
			p.ComparisonOp()
		}
		{
			p.SetState(239)
			p.IpValue()
		}

	case 18:
		localctx = NewIpSetConditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(241)
			p.FieldList()
		}
		{
			p.SetState(242)
			p.SetOp()
		}
		{
			p.SetState(243)
			p.ValueList()
		}

	case 19:
		localctx = NewAllFieldsComparisonConditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(245)
			p.AllFieldList()
		}
		{
			p.SetState(246)
			p.ComparisonOp()
		}
		{
			p.SetState(247)
			p.Value()
		}

	case 20:
		localctx = NewAllFieldsSetConditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(249)
			p.AllFieldList()
		}
		{
			p.SetState(250)
			p.SetOp()
		}
		{
			p.SetState(251)
			p.ValueList()
		}

	case 21:
		localctx = NewAllFieldsNocaseConditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 21)
		{
			p.SetState(253)
			p.AllFieldList()
		}
		{
			p.SetState(254)
			p.ComparisonOp()
		}
		{
			p.SetState(255)
			p.NocaseValue()
		}

	case 22:
		localctx = NewFieldExistsConditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 22)
		{
			p.SetState(257)
			p.Match(LEQLParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 23:
		localctx = NewNumberKeywordConditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 23)
		{
			p.SetState(258)
			p.Match(LEQLParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 24:
		localctx = NewVariableConditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 24)
		{
			p.SetState(259)
			p.Match(LEQLParserVARIABLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IKeywordSearchContext is an interface to support dynamic dispatch.
type IKeywordSearchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsKeywordSearchContext differentiates from other interfaces.
	IsKeywordSearchContext()
}

type KeywordSearchContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeywordSearchContext() *KeywordSearchContext {
	var p = new(KeywordSearchContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_keywordSearch
	return p
}

func InitEmptyKeywordSearchContext(p *KeywordSearchContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_keywordSearch
}

func (*KeywordSearchContext) IsKeywordSearchContext() {}

func NewKeywordSearchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeywordSearchContext {
	var p = new(KeywordSearchContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LEQLParserRULE_keywordSearch

	return p
}

func (s *KeywordSearchContext) GetParser() antlr.Parser { return s.parser }

func (s *KeywordSearchContext) CopyAll(ctx *KeywordSearchContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *KeywordSearchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeywordSearchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SingleQuotedKeywordContext struct {
	KeywordSearchContext
}

func NewSingleQuotedKeywordContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SingleQuotedKeywordContext {
	var p = new(SingleQuotedKeywordContext)

	InitEmptyKeywordSearchContext(&p.KeywordSearchContext)
	p.parser = parser
	p.CopyAll(ctx.(*KeywordSearchContext))

	return p
}

func (s *SingleQuotedKeywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleQuotedKeywordContext) SINGLE_STRING() antlr.TerminalNode {
	return s.GetToken(LEQLParserSINGLE_STRING, 0)
}

func (s *SingleQuotedKeywordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterSingleQuotedKeyword(s)
	}
}

func (s *SingleQuotedKeywordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitSingleQuotedKeyword(s)
	}
}

func (s *SingleQuotedKeywordContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitSingleQuotedKeyword(s)

	default:
		return t.VisitChildren(s)
	}
}

type QuotedKeywordContext struct {
	KeywordSearchContext
}

func NewQuotedKeywordContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *QuotedKeywordContext {
	var p = new(QuotedKeywordContext)

	InitEmptyKeywordSearchContext(&p.KeywordSearchContext)
	p.parser = parser
	p.CopyAll(ctx.(*KeywordSearchContext))

	return p
}

func (s *QuotedKeywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuotedKeywordContext) DOUBLE_STRING() antlr.TerminalNode {
	return s.GetToken(LEQLParserDOUBLE_STRING, 0)
}

func (s *QuotedKeywordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterQuotedKeyword(s)
	}
}

func (s *QuotedKeywordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitQuotedKeyword(s)
	}
}

func (s *QuotedKeywordContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitQuotedKeyword(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LEQLParser) KeywordSearch() (localctx IKeywordSearchContext) {
	localctx = NewKeywordSearchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, LEQLParserRULE_keywordSearch)
	p.SetState(264)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LEQLParserDOUBLE_STRING:
		localctx = NewQuotedKeywordContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(262)
			p.Match(LEQLParserDOUBLE_STRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case LEQLParserSINGLE_STRING:
		localctx = NewSingleQuotedKeywordContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(263)
			p.Match(LEQLParserSINGLE_STRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRegexSearchContext is an interface to support dynamic dispatch.
type IRegexSearchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsRegexSearchContext differentiates from other interfaces.
	IsRegexSearchContext()
}

type RegexSearchContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRegexSearchContext() *RegexSearchContext {
	var p = new(RegexSearchContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_regexSearch
	return p
}

func InitEmptyRegexSearchContext(p *RegexSearchContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_regexSearch
}

func (*RegexSearchContext) IsRegexSearchContext() {}

func NewRegexSearchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RegexSearchContext {
	var p = new(RegexSearchContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LEQLParserRULE_regexSearch

	return p
}

func (s *RegexSearchContext) GetParser() antlr.Parser { return s.parser }

func (s *RegexSearchContext) CopyAll(ctx *RegexSearchContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *RegexSearchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegexSearchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type RegexPatternContext struct {
	RegexSearchContext
}

func NewRegexPatternContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RegexPatternContext {
	var p = new(RegexPatternContext)

	InitEmptyRegexSearchContext(&p.RegexSearchContext)
	p.parser = parser
	p.CopyAll(ctx.(*RegexSearchContext))

	return p
}

func (s *RegexPatternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegexPatternContext) REGEX() antlr.TerminalNode {
	return s.GetToken(LEQLParserREGEX, 0)
}

func (s *RegexPatternContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterRegexPattern(s)
	}
}

func (s *RegexPatternContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitRegexPattern(s)
	}
}

func (s *RegexPatternContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitRegexPattern(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LEQLParser) RegexSearch() (localctx IRegexSearchContext) {
	localctx = NewRegexSearchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, LEQLParserRULE_regexSearch)
	localctx = NewRegexPatternContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(266)
		p.Match(LEQLParserREGEX)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFieldListContext is an interface to support dynamic dispatch.
type IFieldListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllFieldName() []IFieldNameContext
	FieldName(i int) IFieldNameContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsFieldListContext differentiates from other interfaces.
	IsFieldListContext()
}

type FieldListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldListContext() *FieldListContext {
	var p = new(FieldListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_fieldList
	return p
}

func InitEmptyFieldListContext(p *FieldListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_fieldList
}

func (*FieldListContext) IsFieldListContext() {}

func NewFieldListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldListContext {
	var p = new(FieldListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LEQLParserRULE_fieldList

	return p
}

func (s *FieldListContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldListContext) AllFieldName() []IFieldNameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldNameContext); ok {
			len++
		}
	}

	tst := make([]IFieldNameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldNameContext); ok {
			tst[i] = t.(IFieldNameContext)
			i++
		}
	}

	return tst
}

func (s *FieldListContext) FieldName(i int) IFieldNameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldNameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldNameContext)
}

func (s *FieldListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(LEQLParserCOMMA)
}

func (s *FieldListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(LEQLParserCOMMA, i)
}

func (s *FieldListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterFieldList(s)
	}
}

func (s *FieldListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitFieldList(s)
	}
}

func (s *FieldListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitFieldList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LEQLParser) FieldList() (localctx IFieldListContext) {
	localctx = NewFieldListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, LEQLParserRULE_fieldList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(268)
		p.FieldName()
	}
	p.SetState(273)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == LEQLParserCOMMA {
		{
			p.SetState(269)
			p.Match(LEQLParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(270)
			p.FieldName()
		}

		p.SetState(275)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAllFieldListContext is an interface to support dynamic dispatch.
type IAllFieldListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ALL_FUNC() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	AllFieldName() []IFieldNameContext
	FieldName(i int) IFieldNameContext
	RPAREN() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsAllFieldListContext differentiates from other interfaces.
	IsAllFieldListContext()
}

type AllFieldListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAllFieldListContext() *AllFieldListContext {
	var p = new(AllFieldListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_allFieldList
	return p
}

func InitEmptyAllFieldListContext(p *AllFieldListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_allFieldList
}

func (*AllFieldListContext) IsAllFieldListContext() {}

func NewAllFieldListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AllFieldListContext {
	var p = new(AllFieldListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LEQLParserRULE_allFieldList

	return p
}

func (s *AllFieldListContext) GetParser() antlr.Parser { return s.parser }

func (s *AllFieldListContext) ALL_FUNC() antlr.TerminalNode {
	return s.GetToken(LEQLParserALL_FUNC, 0)
}

func (s *AllFieldListContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(LEQLParserLPAREN, 0)
}

func (s *AllFieldListContext) AllFieldName() []IFieldNameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldNameContext); ok {
			len++
		}
	}

	tst := make([]IFieldNameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldNameContext); ok {
			tst[i] = t.(IFieldNameContext)
			i++
		}
	}

	return tst
}

func (s *AllFieldListContext) FieldName(i int) IFieldNameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldNameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldNameContext)
}

func (s *AllFieldListContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(LEQLParserRPAREN, 0)
}

func (s *AllFieldListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(LEQLParserCOMMA)
}

func (s *AllFieldListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(LEQLParserCOMMA, i)
}

func (s *AllFieldListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AllFieldListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AllFieldListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterAllFieldList(s)
	}
}

func (s *AllFieldListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitAllFieldList(s)
	}
}

func (s *AllFieldListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitAllFieldList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LEQLParser) AllFieldList() (localctx IAllFieldListContext) {
	localctx = NewAllFieldListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, LEQLParserRULE_allFieldList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(276)
		p.Match(LEQLParserALL_FUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(277)
		p.Match(LEQLParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(278)
		p.FieldName()
	}
	p.SetState(283)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == LEQLParserCOMMA {
		{
			p.SetState(279)
			p.Match(LEQLParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(280)
			p.FieldName()
		}

		p.SetState(285)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(286)
		p.Match(LEQLParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFieldNameContext is an interface to support dynamic dispatch.
type IFieldNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	DOUBLE_STRING() antlr.TerminalNode
	SINGLE_STRING() antlr.TerminalNode
	IP_FUNC() antlr.TerminalNode
	ALL_FUNC() antlr.TerminalNode
	COUNT() antlr.TerminalNode
	BYTES() antlr.TerminalNode
	MIN() antlr.TerminalNode
	MAX() antlr.TerminalNode
	SUM() antlr.TerminalNode
	SD() antlr.TerminalNode
	FROM() antlr.TerminalNode
	LOOSE() antlr.TerminalNode

	// IsFieldNameContext differentiates from other interfaces.
	IsFieldNameContext()
}

type FieldNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldNameContext() *FieldNameContext {
	var p = new(FieldNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_fieldName
	return p
}

func InitEmptyFieldNameContext(p *FieldNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_fieldName
}

func (*FieldNameContext) IsFieldNameContext() {}

func NewFieldNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldNameContext {
	var p = new(FieldNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LEQLParserRULE_fieldName

	return p
}

func (s *FieldNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(LEQLParserIDENTIFIER, 0)
}

func (s *FieldNameContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(LEQLParserNUMBER, 0)
}

func (s *FieldNameContext) DOUBLE_STRING() antlr.TerminalNode {
	return s.GetToken(LEQLParserDOUBLE_STRING, 0)
}

func (s *FieldNameContext) SINGLE_STRING() antlr.TerminalNode {
	return s.GetToken(LEQLParserSINGLE_STRING, 0)
}

func (s *FieldNameContext) IP_FUNC() antlr.TerminalNode {
	return s.GetToken(LEQLParserIP_FUNC, 0)
}

func (s *FieldNameContext) ALL_FUNC() antlr.TerminalNode {
	return s.GetToken(LEQLParserALL_FUNC, 0)
}

func (s *FieldNameContext) COUNT() antlr.TerminalNode {
	return s.GetToken(LEQLParserCOUNT, 0)
}

func (s *FieldNameContext) BYTES() antlr.TerminalNode {
	return s.GetToken(LEQLParserBYTES, 0)
}

func (s *FieldNameContext) MIN() antlr.TerminalNode {
	return s.GetToken(LEQLParserMIN, 0)
}

func (s *FieldNameContext) MAX() antlr.TerminalNode {
	return s.GetToken(LEQLParserMAX, 0)
}

func (s *FieldNameContext) SUM() antlr.TerminalNode {
	return s.GetToken(LEQLParserSUM, 0)
}

func (s *FieldNameContext) SD() antlr.TerminalNode {
	return s.GetToken(LEQLParserSD, 0)
}

func (s *FieldNameContext) FROM() antlr.TerminalNode {
	return s.GetToken(LEQLParserFROM, 0)
}

func (s *FieldNameContext) LOOSE() antlr.TerminalNode {
	return s.GetToken(LEQLParserLOOSE, 0)
}

func (s *FieldNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterFieldName(s)
	}
}

func (s *FieldNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitFieldName(s)
	}
}

func (s *FieldNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitFieldName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LEQLParser) FieldName() (localctx IFieldNameContext) {
	localctx = NewFieldNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, LEQLParserRULE_fieldName)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(288)
		_la = p.GetTokenStream().LA(1)

		if !(((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1056964609536) != 0) || ((int64((_la-68)) & ^0x3f) == 0 && ((int64(1)<<(_la-68))&39) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IComparisonOpContext is an interface to support dynamic dispatch.
type IComparisonOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EQ() antlr.TerminalNode
	NEQ() antlr.TerminalNode
	STRICT_EQ() antlr.TerminalNode
	STRICT_NEQ() antlr.TerminalNode
	GT() antlr.TerminalNode
	GTE() antlr.TerminalNode
	LT() antlr.TerminalNode
	LTE() antlr.TerminalNode

	// IsComparisonOpContext differentiates from other interfaces.
	IsComparisonOpContext()
}

type ComparisonOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonOpContext() *ComparisonOpContext {
	var p = new(ComparisonOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_comparisonOp
	return p
}

func InitEmptyComparisonOpContext(p *ComparisonOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_comparisonOp
}

func (*ComparisonOpContext) IsComparisonOpContext() {}

func NewComparisonOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonOpContext {
	var p = new(ComparisonOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LEQLParserRULE_comparisonOp

	return p
}

func (s *ComparisonOpContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparisonOpContext) EQ() antlr.TerminalNode {
	return s.GetToken(LEQLParserEQ, 0)
}

func (s *ComparisonOpContext) NEQ() antlr.TerminalNode {
	return s.GetToken(LEQLParserNEQ, 0)
}

func (s *ComparisonOpContext) STRICT_EQ() antlr.TerminalNode {
	return s.GetToken(LEQLParserSTRICT_EQ, 0)
}

func (s *ComparisonOpContext) STRICT_NEQ() antlr.TerminalNode {
	return s.GetToken(LEQLParserSTRICT_NEQ, 0)
}

func (s *ComparisonOpContext) GT() antlr.TerminalNode {
	return s.GetToken(LEQLParserGT, 0)
}

func (s *ComparisonOpContext) GTE() antlr.TerminalNode {
	return s.GetToken(LEQLParserGTE, 0)
}

func (s *ComparisonOpContext) LT() antlr.TerminalNode {
	return s.GetToken(LEQLParserLT, 0)
}

func (s *ComparisonOpContext) LTE() antlr.TerminalNode {
	return s.GetToken(LEQLParserLTE, 0)
}

func (s *ComparisonOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparisonOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterComparisonOp(s)
	}
}

func (s *ComparisonOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitComparisonOp(s)
	}
}

func (s *ComparisonOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitComparisonOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LEQLParser) ComparisonOp() (localctx IComparisonOpContext) {
	localctx = NewComparisonOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, LEQLParserRULE_comparisonOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(290)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17820884462993408) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRegexMatchOpContext is an interface to support dynamic dispatch.
type IRegexMatchOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	REGEX_EQ() antlr.TerminalNode
	REGEX_NEQ() antlr.TerminalNode

	// IsRegexMatchOpContext differentiates from other interfaces.
	IsRegexMatchOpContext()
}

type RegexMatchOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRegexMatchOpContext() *RegexMatchOpContext {
	var p = new(RegexMatchOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_regexMatchOp
	return p
}

func InitEmptyRegexMatchOpContext(p *RegexMatchOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_regexMatchOp
}

func (*RegexMatchOpContext) IsRegexMatchOpContext() {}

func NewRegexMatchOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RegexMatchOpContext {
	var p = new(RegexMatchOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LEQLParserRULE_regexMatchOp

	return p
}

func (s *RegexMatchOpContext) GetParser() antlr.Parser { return s.parser }

func (s *RegexMatchOpContext) REGEX_EQ() antlr.TerminalNode {
	return s.GetToken(LEQLParserREGEX_EQ, 0)
}

func (s *RegexMatchOpContext) REGEX_NEQ() antlr.TerminalNode {
	return s.GetToken(LEQLParserREGEX_NEQ, 0)
}

func (s *RegexMatchOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegexMatchOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RegexMatchOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterRegexMatchOp(s)
	}
}

func (s *RegexMatchOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitRegexMatchOp(s)
	}
}

func (s *RegexMatchOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitRegexMatchOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LEQLParser) RegexMatchOp() (localctx IRegexMatchOpContext) {
	localctx = NewRegexMatchOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, LEQLParserRULE_regexMatchOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(292)
		_la = p.GetTokenStream().LA(1)

		if !(_la == LEQLParserREGEX_NEQ || _la == LEQLParserREGEX_EQ) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStringOpContext is an interface to support dynamic dispatch.
type IStringOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONTAINS_OP() antlr.TerminalNode
	ICONTAINS_OP() antlr.TerminalNode
	STARTS_WITH_OP() antlr.TerminalNode
	ISTARTS_WITH_OP() antlr.TerminalNode

	// IsStringOpContext differentiates from other interfaces.
	IsStringOpContext()
}

type StringOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringOpContext() *StringOpContext {
	var p = new(StringOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_stringOp
	return p
}

func InitEmptyStringOpContext(p *StringOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_stringOp
}

func (*StringOpContext) IsStringOpContext() {}

func NewStringOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringOpContext {
	var p = new(StringOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LEQLParserRULE_stringOp

	return p
}

func (s *StringOpContext) GetParser() antlr.Parser { return s.parser }

func (s *StringOpContext) CONTAINS_OP() antlr.TerminalNode {
	return s.GetToken(LEQLParserCONTAINS_OP, 0)
}

func (s *StringOpContext) ICONTAINS_OP() antlr.TerminalNode {
	return s.GetToken(LEQLParserICONTAINS_OP, 0)
}

func (s *StringOpContext) STARTS_WITH_OP() antlr.TerminalNode {
	return s.GetToken(LEQLParserSTARTS_WITH_OP, 0)
}

func (s *StringOpContext) ISTARTS_WITH_OP() antlr.TerminalNode {
	return s.GetToken(LEQLParserISTARTS_WITH_OP, 0)
}

func (s *StringOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterStringOp(s)
	}
}

func (s *StringOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitStringOp(s)
	}
}

func (s *StringOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitStringOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LEQLParser) StringOp() (localctx IStringOpContext) {
	localctx = NewStringOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, LEQLParserRULE_stringOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(294)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&15728640) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISetOpContext is an interface to support dynamic dispatch.
type ISetOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IN_OP() antlr.TerminalNode
	IIN_OP() antlr.TerminalNode

	// IsSetOpContext differentiates from other interfaces.
	IsSetOpContext()
}

type SetOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySetOpContext() *SetOpContext {
	var p = new(SetOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_setOp
	return p
}

func InitEmptySetOpContext(p *SetOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_setOp
}

func (*SetOpContext) IsSetOpContext() {}

func NewSetOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SetOpContext {
	var p = new(SetOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LEQLParserRULE_setOp

	return p
}

func (s *SetOpContext) GetParser() antlr.Parser { return s.parser }

func (s *SetOpContext) IN_OP() antlr.TerminalNode {
	return s.GetToken(LEQLParserIN_OP, 0)
}

func (s *SetOpContext) IIN_OP() antlr.TerminalNode {
	return s.GetToken(LEQLParserIIN_OP, 0)
}

func (s *SetOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SetOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterSetOp(s)
	}
}

func (s *SetOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitSetOp(s)
	}
}

func (s *SetOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitSetOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LEQLParser) SetOp() (localctx ISetOpContext) {
	localctx = NewSetOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, LEQLParserRULE_setOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(296)
		_la = p.GetTokenStream().LA(1)

		if !(_la == LEQLParserIIN_OP || _la == LEQLParserIN_OP) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListStringOpContext is an interface to support dynamic dispatch.
type IListStringOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONTAINS_ANY_OP() antlr.TerminalNode
	ICONTAINS_ANY_OP() antlr.TerminalNode
	CONTAINS_ALL_OP() antlr.TerminalNode
	ICONTAINS_ALL_OP() antlr.TerminalNode
	STARTS_WITH_ANY_OP() antlr.TerminalNode
	ISTARTS_WITH_ANY_OP() antlr.TerminalNode

	// IsListStringOpContext differentiates from other interfaces.
	IsListStringOpContext()
}

type ListStringOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListStringOpContext() *ListStringOpContext {
	var p = new(ListStringOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_listStringOp
	return p
}

func InitEmptyListStringOpContext(p *ListStringOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_listStringOp
}

func (*ListStringOpContext) IsListStringOpContext() {}

func NewListStringOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListStringOpContext {
	var p = new(ListStringOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LEQLParserRULE_listStringOp

	return p
}

func (s *ListStringOpContext) GetParser() antlr.Parser { return s.parser }

func (s *ListStringOpContext) CONTAINS_ANY_OP() antlr.TerminalNode {
	return s.GetToken(LEQLParserCONTAINS_ANY_OP, 0)
}

func (s *ListStringOpContext) ICONTAINS_ANY_OP() antlr.TerminalNode {
	return s.GetToken(LEQLParserICONTAINS_ANY_OP, 0)
}

func (s *ListStringOpContext) CONTAINS_ALL_OP() antlr.TerminalNode {
	return s.GetToken(LEQLParserCONTAINS_ALL_OP, 0)
}

func (s *ListStringOpContext) ICONTAINS_ALL_OP() antlr.TerminalNode {
	return s.GetToken(LEQLParserICONTAINS_ALL_OP, 0)
}

func (s *ListStringOpContext) STARTS_WITH_ANY_OP() antlr.TerminalNode {
	return s.GetToken(LEQLParserSTARTS_WITH_ANY_OP, 0)
}

func (s *ListStringOpContext) ISTARTS_WITH_ANY_OP() antlr.TerminalNode {
	return s.GetToken(LEQLParserISTARTS_WITH_ANY_OP, 0)
}

func (s *ListStringOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListStringOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListStringOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterListStringOp(s)
	}
}

func (s *ListStringOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitListStringOp(s)
	}
}

func (s *ListStringOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitListStringOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LEQLParser) ListStringOp() (localctx IListStringOpContext) {
	localctx = NewListStringOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, LEQLParserRULE_listStringOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(298)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1032192) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DOUBLE_STRING() antlr.TerminalNode
	SINGLE_STRING() antlr.TerminalNode
	TRIPLE_SINGLE_STRING() antlr.TerminalNode
	TRIPLE_DOUBLE_STRING() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	TIME_UNIT() antlr.TerminalNode
	REGEX() antlr.TerminalNode
	IP_CIDR() antlr.TerminalNode
	VARIABLE() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LEQLParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) DOUBLE_STRING() antlr.TerminalNode {
	return s.GetToken(LEQLParserDOUBLE_STRING, 0)
}

func (s *ValueContext) SINGLE_STRING() antlr.TerminalNode {
	return s.GetToken(LEQLParserSINGLE_STRING, 0)
}

func (s *ValueContext) TRIPLE_SINGLE_STRING() antlr.TerminalNode {
	return s.GetToken(LEQLParserTRIPLE_SINGLE_STRING, 0)
}

func (s *ValueContext) TRIPLE_DOUBLE_STRING() antlr.TerminalNode {
	return s.GetToken(LEQLParserTRIPLE_DOUBLE_STRING, 0)
}

func (s *ValueContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(LEQLParserNUMBER, 0)
}

func (s *ValueContext) TIME_UNIT() antlr.TerminalNode {
	return s.GetToken(LEQLParserTIME_UNIT, 0)
}

func (s *ValueContext) REGEX() antlr.TerminalNode {
	return s.GetToken(LEQLParserREGEX, 0)
}

func (s *ValueContext) IP_CIDR() antlr.TerminalNode {
	return s.GetToken(LEQLParserIP_CIDR, 0)
}

func (s *ValueContext) VARIABLE() antlr.TerminalNode {
	return s.GetToken(LEQLParserVARIABLE, 0)
}

func (s *ValueContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(LEQLParserIDENTIFIER, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitValue(s)
	}
}

func (s *ValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LEQLParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, LEQLParserRULE_value)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(300)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&1023) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INocaseValueContext is an interface to support dynamic dispatch.
type INocaseValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NOCASE() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Value() IValueContext
	RPAREN() antlr.TerminalNode

	// IsNocaseValueContext differentiates from other interfaces.
	IsNocaseValueContext()
}

type NocaseValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNocaseValueContext() *NocaseValueContext {
	var p = new(NocaseValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_nocaseValue
	return p
}

func InitEmptyNocaseValueContext(p *NocaseValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_nocaseValue
}

func (*NocaseValueContext) IsNocaseValueContext() {}

func NewNocaseValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NocaseValueContext {
	var p = new(NocaseValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LEQLParserRULE_nocaseValue

	return p
}

func (s *NocaseValueContext) GetParser() antlr.Parser { return s.parser }

func (s *NocaseValueContext) NOCASE() antlr.TerminalNode {
	return s.GetToken(LEQLParserNOCASE, 0)
}

func (s *NocaseValueContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(LEQLParserLPAREN, 0)
}

func (s *NocaseValueContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *NocaseValueContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(LEQLParserRPAREN, 0)
}

func (s *NocaseValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NocaseValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NocaseValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterNocaseValue(s)
	}
}

func (s *NocaseValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitNocaseValue(s)
	}
}

func (s *NocaseValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitNocaseValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LEQLParser) NocaseValue() (localctx INocaseValueContext) {
	localctx = NewNocaseValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, LEQLParserRULE_nocaseValue)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(302)
		p.Match(LEQLParserNOCASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(303)
		p.Match(LEQLParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(304)
		p.Value()
	}
	{
		p.SetState(305)
		p.Match(LEQLParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIpValueContext is an interface to support dynamic dispatch.
type IIpValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IP_FUNC() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	IP_CIDR() antlr.TerminalNode
	Value() IValueContext

	// IsIpValueContext differentiates from other interfaces.
	IsIpValueContext()
}

type IpValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIpValueContext() *IpValueContext {
	var p = new(IpValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_ipValue
	return p
}

func InitEmptyIpValueContext(p *IpValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_ipValue
}

func (*IpValueContext) IsIpValueContext() {}

func NewIpValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IpValueContext {
	var p = new(IpValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LEQLParserRULE_ipValue

	return p
}

func (s *IpValueContext) GetParser() antlr.Parser { return s.parser }

func (s *IpValueContext) IP_FUNC() antlr.TerminalNode {
	return s.GetToken(LEQLParserIP_FUNC, 0)
}

func (s *IpValueContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(LEQLParserLPAREN, 0)
}

func (s *IpValueContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(LEQLParserRPAREN, 0)
}

func (s *IpValueContext) IP_CIDR() antlr.TerminalNode {
	return s.GetToken(LEQLParserIP_CIDR, 0)
}

func (s *IpValueContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *IpValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IpValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IpValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterIpValue(s)
	}
}

func (s *IpValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitIpValue(s)
	}
}

func (s *IpValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitIpValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LEQLParser) IpValue() (localctx IIpValueContext) {
	localctx = NewIpValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, LEQLParserRULE_ipValue)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(307)
		p.Match(LEQLParserIP_FUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(308)
		p.Match(LEQLParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(311)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(309)
			p.Match(LEQLParserIP_CIDR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		{
			p.SetState(310)
			p.Value()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(313)
		p.Match(LEQLParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValueListContext is an interface to support dynamic dispatch.
type IValueListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACKET() antlr.TerminalNode
	AllValueListItem() []IValueListItemContext
	ValueListItem(i int) IValueListItemContext
	RBRACKET() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsValueListContext differentiates from other interfaces.
	IsValueListContext()
}

type ValueListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueListContext() *ValueListContext {
	var p = new(ValueListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_valueList
	return p
}

func InitEmptyValueListContext(p *ValueListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_valueList
}

func (*ValueListContext) IsValueListContext() {}

func NewValueListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueListContext {
	var p = new(ValueListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LEQLParserRULE_valueList

	return p
}

func (s *ValueListContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueListContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(LEQLParserLBRACKET, 0)
}

func (s *ValueListContext) AllValueListItem() []IValueListItemContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValueListItemContext); ok {
			len++
		}
	}

	tst := make([]IValueListItemContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValueListItemContext); ok {
			tst[i] = t.(IValueListItemContext)
			i++
		}
	}

	return tst
}

func (s *ValueListContext) ValueListItem(i int) IValueListItemContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueListItemContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueListItemContext)
}

func (s *ValueListContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(LEQLParserRBRACKET, 0)
}

func (s *ValueListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(LEQLParserCOMMA)
}

func (s *ValueListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(LEQLParserCOMMA, i)
}

func (s *ValueListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterValueList(s)
	}
}

func (s *ValueListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitValueList(s)
	}
}

func (s *ValueListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitValueList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LEQLParser) ValueList() (localctx IValueListContext) {
	localctx = NewValueListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, LEQLParserRULE_valueList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(315)
		p.Match(LEQLParserLBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(316)
		p.ValueListItem()
	}
	p.SetState(321)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == LEQLParserCOMMA {
		{
			p.SetState(317)
			p.Match(LEQLParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(318)
			p.ValueListItem()
		}

		p.SetState(323)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(324)
		p.Match(LEQLParserRBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValueListItemContext is an interface to support dynamic dispatch.
type IValueListItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IpValue() IIpValueContext
	Value() IValueContext

	// IsValueListItemContext differentiates from other interfaces.
	IsValueListItemContext()
}

type ValueListItemContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueListItemContext() *ValueListItemContext {
	var p = new(ValueListItemContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_valueListItem
	return p
}

func InitEmptyValueListItemContext(p *ValueListItemContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_valueListItem
}

func (*ValueListItemContext) IsValueListItemContext() {}

func NewValueListItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueListItemContext {
	var p = new(ValueListItemContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LEQLParserRULE_valueListItem

	return p
}

func (s *ValueListItemContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueListItemContext) IpValue() IIpValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIpValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIpValueContext)
}

func (s *ValueListItemContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ValueListItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueListItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueListItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterValueListItem(s)
	}
}

func (s *ValueListItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitValueListItem(s)
	}
}

func (s *ValueListItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitValueListItem(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LEQLParser) ValueListItem() (localctx IValueListItemContext) {
	localctx = NewValueListItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, LEQLParserRULE_valueListItem)
	p.SetState(328)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LEQLParserIP_FUNC:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(326)
			p.IpValue()
		}

	case LEQLParserIP_CIDR, LEQLParserREGEX, LEQLParserTRIPLE_SINGLE_STRING, LEQLParserTRIPLE_DOUBLE_STRING, LEQLParserDOUBLE_STRING, LEQLParserSINGLE_STRING, LEQLParserNUMBER, LEQLParserTIME_UNIT, LEQLParserVARIABLE, LEQLParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(327)
			p.Value()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGroupbyClauseContext is an interface to support dynamic dispatch.
type IGroupbyClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GROUPBY() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	AllFieldName() []IFieldNameContext
	FieldName(i int) IFieldNameContext
	RPAREN() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsGroupbyClauseContext differentiates from other interfaces.
	IsGroupbyClauseContext()
}

type GroupbyClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroupbyClauseContext() *GroupbyClauseContext {
	var p = new(GroupbyClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_groupbyClause
	return p
}

func InitEmptyGroupbyClauseContext(p *GroupbyClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_groupbyClause
}

func (*GroupbyClauseContext) IsGroupbyClauseContext() {}

func NewGroupbyClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupbyClauseContext {
	var p = new(GroupbyClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LEQLParserRULE_groupbyClause

	return p
}

func (s *GroupbyClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupbyClauseContext) GROUPBY() antlr.TerminalNode {
	return s.GetToken(LEQLParserGROUPBY, 0)
}

func (s *GroupbyClauseContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(LEQLParserLPAREN, 0)
}

func (s *GroupbyClauseContext) AllFieldName() []IFieldNameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldNameContext); ok {
			len++
		}
	}

	tst := make([]IFieldNameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldNameContext); ok {
			tst[i] = t.(IFieldNameContext)
			i++
		}
	}

	return tst
}

func (s *GroupbyClauseContext) FieldName(i int) IFieldNameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldNameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldNameContext)
}

func (s *GroupbyClauseContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(LEQLParserRPAREN, 0)
}

func (s *GroupbyClauseContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(LEQLParserCOMMA)
}

func (s *GroupbyClauseContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(LEQLParserCOMMA, i)
}

func (s *GroupbyClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupbyClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupbyClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterGroupbyClause(s)
	}
}

func (s *GroupbyClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitGroupbyClause(s)
	}
}

func (s *GroupbyClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitGroupbyClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LEQLParser) GroupbyClause() (localctx IGroupbyClauseContext) {
	localctx = NewGroupbyClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, LEQLParserRULE_groupbyClause)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(330)
		p.Match(LEQLParserGROUPBY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(331)
		p.Match(LEQLParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(332)
		p.FieldName()
	}
	p.SetState(337)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == LEQLParserCOMMA {
		{
			p.SetState(333)
			p.Match(LEQLParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(334)
			p.FieldName()
		}

		p.SetState(339)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(340)
		p.Match(LEQLParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IHavingClauseContext is an interface to support dynamic dispatch.
type IHavingClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	HAVING() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	HavingCondition() IHavingConditionContext
	RPAREN() antlr.TerminalNode

	// IsHavingClauseContext differentiates from other interfaces.
	IsHavingClauseContext()
}

type HavingClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHavingClauseContext() *HavingClauseContext {
	var p = new(HavingClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_havingClause
	return p
}

func InitEmptyHavingClauseContext(p *HavingClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_havingClause
}

func (*HavingClauseContext) IsHavingClauseContext() {}

func NewHavingClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HavingClauseContext {
	var p = new(HavingClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LEQLParserRULE_havingClause

	return p
}

func (s *HavingClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *HavingClauseContext) HAVING() antlr.TerminalNode {
	return s.GetToken(LEQLParserHAVING, 0)
}

func (s *HavingClauseContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(LEQLParserLPAREN, 0)
}

func (s *HavingClauseContext) HavingCondition() IHavingConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHavingConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHavingConditionContext)
}

func (s *HavingClauseContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(LEQLParserRPAREN, 0)
}

func (s *HavingClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HavingClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HavingClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterHavingClause(s)
	}
}

func (s *HavingClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitHavingClause(s)
	}
}

func (s *HavingClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitHavingClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LEQLParser) HavingClause() (localctx IHavingClauseContext) {
	localctx = NewHavingClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, LEQLParserRULE_havingClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(342)
		p.Match(LEQLParserHAVING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(343)
		p.Match(LEQLParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(344)
		p.HavingCondition()
	}
	{
		p.SetState(345)
		p.Match(LEQLParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IHavingConditionContext is an interface to support dynamic dispatch.
type IHavingConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CalcFunction() ICalcFunctionContext
	ComparisonOp() IComparisonOpContext
	Value() IValueContext
	CalcFunctionWithField() ICalcFunctionWithFieldContext

	// IsHavingConditionContext differentiates from other interfaces.
	IsHavingConditionContext()
}

type HavingConditionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHavingConditionContext() *HavingConditionContext {
	var p = new(HavingConditionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_havingCondition
	return p
}

func InitEmptyHavingConditionContext(p *HavingConditionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_havingCondition
}

func (*HavingConditionContext) IsHavingConditionContext() {}

func NewHavingConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HavingConditionContext {
	var p = new(HavingConditionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LEQLParserRULE_havingCondition

	return p
}

func (s *HavingConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *HavingConditionContext) CalcFunction() ICalcFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICalcFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICalcFunctionContext)
}

func (s *HavingConditionContext) ComparisonOp() IComparisonOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparisonOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparisonOpContext)
}

func (s *HavingConditionContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *HavingConditionContext) CalcFunctionWithField() ICalcFunctionWithFieldContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICalcFunctionWithFieldContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICalcFunctionWithFieldContext)
}

func (s *HavingConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HavingConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HavingConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterHavingCondition(s)
	}
}

func (s *HavingConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitHavingCondition(s)
	}
}

func (s *HavingConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitHavingCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LEQLParser) HavingCondition() (localctx IHavingConditionContext) {
	localctx = NewHavingConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, LEQLParserRULE_havingCondition)
	p.SetState(355)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(347)
			p.CalcFunction()
		}
		{
			p.SetState(348)
			p.ComparisonOp()
		}
		{
			p.SetState(349)
			p.Value()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(351)
			p.CalcFunctionWithField()
		}
		{
			p.SetState(352)
			p.ComparisonOp()
		}
		{
			p.SetState(353)
			p.Value()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICalculateClauseContext is an interface to support dynamic dispatch.
type ICalculateClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CALCULATE() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	CalcExpr() ICalcExprContext
	RPAREN() antlr.TerminalNode

	// IsCalculateClauseContext differentiates from other interfaces.
	IsCalculateClauseContext()
}

type CalculateClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCalculateClauseContext() *CalculateClauseContext {
	var p = new(CalculateClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_calculateClause
	return p
}

func InitEmptyCalculateClauseContext(p *CalculateClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_calculateClause
}

func (*CalculateClauseContext) IsCalculateClauseContext() {}

func NewCalculateClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CalculateClauseContext {
	var p = new(CalculateClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LEQLParserRULE_calculateClause

	return p
}

func (s *CalculateClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *CalculateClauseContext) CALCULATE() antlr.TerminalNode {
	return s.GetToken(LEQLParserCALCULATE, 0)
}

func (s *CalculateClauseContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(LEQLParserLPAREN, 0)
}

func (s *CalculateClauseContext) CalcExpr() ICalcExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICalcExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICalcExprContext)
}

func (s *CalculateClauseContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(LEQLParserRPAREN, 0)
}

func (s *CalculateClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CalculateClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CalculateClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterCalculateClause(s)
	}
}

func (s *CalculateClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitCalculateClause(s)
	}
}

func (s *CalculateClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitCalculateClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LEQLParser) CalculateClause() (localctx ICalculateClauseContext) {
	localctx = NewCalculateClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, LEQLParserRULE_calculateClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(357)
		p.Match(LEQLParserCALCULATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(358)
		p.Match(LEQLParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(359)
		p.CalcExpr()
	}
	{
		p.SetState(360)
		p.Match(LEQLParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICalcExprContext is an interface to support dynamic dispatch.
type ICalcExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CalcFunction() ICalcFunctionContext
	CalcFunctionWithField() ICalcFunctionWithFieldContext
	PercentileFunction() IPercentileFunctionContext

	// IsCalcExprContext differentiates from other interfaces.
	IsCalcExprContext()
}

type CalcExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCalcExprContext() *CalcExprContext {
	var p = new(CalcExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_calcExpr
	return p
}

func InitEmptyCalcExprContext(p *CalcExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_calcExpr
}

func (*CalcExprContext) IsCalcExprContext() {}

func NewCalcExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CalcExprContext {
	var p = new(CalcExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LEQLParserRULE_calcExpr

	return p
}

func (s *CalcExprContext) GetParser() antlr.Parser { return s.parser }

func (s *CalcExprContext) CalcFunction() ICalcFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICalcFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICalcFunctionContext)
}

func (s *CalcExprContext) CalcFunctionWithField() ICalcFunctionWithFieldContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICalcFunctionWithFieldContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICalcFunctionWithFieldContext)
}

func (s *CalcExprContext) PercentileFunction() IPercentileFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPercentileFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPercentileFunctionContext)
}

func (s *CalcExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CalcExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CalcExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterCalcExpr(s)
	}
}

func (s *CalcExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitCalcExpr(s)
	}
}

func (s *CalcExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitCalcExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LEQLParser) CalcExpr() (localctx ICalcExprContext) {
	localctx = NewCalcExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, LEQLParserRULE_calcExpr)
	p.SetState(365)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(362)
			p.CalcFunction()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(363)
			p.CalcFunctionWithField()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(364)
			p.PercentileFunction()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICalcFunctionContext is an interface to support dynamic dispatch.
type ICalcFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COUNT() antlr.TerminalNode
	BYTES() antlr.TerminalNode

	// IsCalcFunctionContext differentiates from other interfaces.
	IsCalcFunctionContext()
}

type CalcFunctionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCalcFunctionContext() *CalcFunctionContext {
	var p = new(CalcFunctionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_calcFunction
	return p
}

func InitEmptyCalcFunctionContext(p *CalcFunctionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_calcFunction
}

func (*CalcFunctionContext) IsCalcFunctionContext() {}

func NewCalcFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CalcFunctionContext {
	var p = new(CalcFunctionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LEQLParserRULE_calcFunction

	return p
}

func (s *CalcFunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *CalcFunctionContext) COUNT() antlr.TerminalNode {
	return s.GetToken(LEQLParserCOUNT, 0)
}

func (s *CalcFunctionContext) BYTES() antlr.TerminalNode {
	return s.GetToken(LEQLParserBYTES, 0)
}

func (s *CalcFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CalcFunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CalcFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterCalcFunction(s)
	}
}

func (s *CalcFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitCalcFunction(s)
	}
}

func (s *CalcFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitCalcFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LEQLParser) CalcFunction() (localctx ICalcFunctionContext) {
	localctx = NewCalcFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, LEQLParserRULE_calcFunction)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(367)
		_la = p.GetTokenStream().LA(1)

		if !(_la == LEQLParserCOUNT || _la == LEQLParserBYTES) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICalcFunctionWithFieldContext is an interface to support dynamic dispatch.
type ICalcFunctionWithFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COLON() antlr.TerminalNode
	FieldName() IFieldNameContext
	SUM() antlr.TerminalNode
	AVERAGE() antlr.TerminalNode
	UNIQUE() antlr.TerminalNode
	MIN() antlr.TerminalNode
	MAX() antlr.TerminalNode
	STANDARDDEVIATION() antlr.TerminalNode
	SD() antlr.TerminalNode
	COUNT() antlr.TerminalNode
	BYTES() antlr.TerminalNode

	// IsCalcFunctionWithFieldContext differentiates from other interfaces.
	IsCalcFunctionWithFieldContext()
}

type CalcFunctionWithFieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCalcFunctionWithFieldContext() *CalcFunctionWithFieldContext {
	var p = new(CalcFunctionWithFieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_calcFunctionWithField
	return p
}

func InitEmptyCalcFunctionWithFieldContext(p *CalcFunctionWithFieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_calcFunctionWithField
}

func (*CalcFunctionWithFieldContext) IsCalcFunctionWithFieldContext() {}

func NewCalcFunctionWithFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CalcFunctionWithFieldContext {
	var p = new(CalcFunctionWithFieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LEQLParserRULE_calcFunctionWithField

	return p
}

func (s *CalcFunctionWithFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *CalcFunctionWithFieldContext) COLON() antlr.TerminalNode {
	return s.GetToken(LEQLParserCOLON, 0)
}

func (s *CalcFunctionWithFieldContext) FieldName() IFieldNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldNameContext)
}

func (s *CalcFunctionWithFieldContext) SUM() antlr.TerminalNode {
	return s.GetToken(LEQLParserSUM, 0)
}

func (s *CalcFunctionWithFieldContext) AVERAGE() antlr.TerminalNode {
	return s.GetToken(LEQLParserAVERAGE, 0)
}

func (s *CalcFunctionWithFieldContext) UNIQUE() antlr.TerminalNode {
	return s.GetToken(LEQLParserUNIQUE, 0)
}

func (s *CalcFunctionWithFieldContext) MIN() antlr.TerminalNode {
	return s.GetToken(LEQLParserMIN, 0)
}

func (s *CalcFunctionWithFieldContext) MAX() antlr.TerminalNode {
	return s.GetToken(LEQLParserMAX, 0)
}

func (s *CalcFunctionWithFieldContext) STANDARDDEVIATION() antlr.TerminalNode {
	return s.GetToken(LEQLParserSTANDARDDEVIATION, 0)
}

func (s *CalcFunctionWithFieldContext) SD() antlr.TerminalNode {
	return s.GetToken(LEQLParserSD, 0)
}

func (s *CalcFunctionWithFieldContext) COUNT() antlr.TerminalNode {
	return s.GetToken(LEQLParserCOUNT, 0)
}

func (s *CalcFunctionWithFieldContext) BYTES() antlr.TerminalNode {
	return s.GetToken(LEQLParserBYTES, 0)
}

func (s *CalcFunctionWithFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CalcFunctionWithFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CalcFunctionWithFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterCalcFunctionWithField(s)
	}
}

func (s *CalcFunctionWithFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitCalcFunctionWithField(s)
	}
}

func (s *CalcFunctionWithFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitCalcFunctionWithField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LEQLParser) CalcFunctionWithField() (localctx ICalcFunctionWithFieldContext) {
	localctx = NewCalcFunctionWithFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, LEQLParserRULE_calcFunctionWithField)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(369)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1063541276672) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(370)
		p.Match(LEQLParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(371)
		p.FieldName()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPercentileFunctionContext is an interface to support dynamic dispatch.
type IPercentileFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	PERCENTILE() antlr.TerminalNode
	PCTL() antlr.TerminalNode
	COLON() antlr.TerminalNode
	FieldName() IFieldNameContext

	// IsPercentileFunctionContext differentiates from other interfaces.
	IsPercentileFunctionContext()
}

type PercentileFunctionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPercentileFunctionContext() *PercentileFunctionContext {
	var p = new(PercentileFunctionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_percentileFunction
	return p
}

func InitEmptyPercentileFunctionContext(p *PercentileFunctionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_percentileFunction
}

func (*PercentileFunctionContext) IsPercentileFunctionContext() {}

func NewPercentileFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PercentileFunctionContext {
	var p = new(PercentileFunctionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LEQLParserRULE_percentileFunction

	return p
}

func (s *PercentileFunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *PercentileFunctionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(LEQLParserLPAREN, 0)
}

func (s *PercentileFunctionContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(LEQLParserNUMBER, 0)
}

func (s *PercentileFunctionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(LEQLParserRPAREN, 0)
}

func (s *PercentileFunctionContext) PERCENTILE() antlr.TerminalNode {
	return s.GetToken(LEQLParserPERCENTILE, 0)
}

func (s *PercentileFunctionContext) PCTL() antlr.TerminalNode {
	return s.GetToken(LEQLParserPCTL, 0)
}

func (s *PercentileFunctionContext) COLON() antlr.TerminalNode {
	return s.GetToken(LEQLParserCOLON, 0)
}

func (s *PercentileFunctionContext) FieldName() IFieldNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldNameContext)
}

func (s *PercentileFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PercentileFunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PercentileFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterPercentileFunction(s)
	}
}

func (s *PercentileFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitPercentileFunction(s)
	}
}

func (s *PercentileFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitPercentileFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LEQLParser) PercentileFunction() (localctx IPercentileFunctionContext) {
	localctx = NewPercentileFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, LEQLParserRULE_percentileFunction)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(373)
		_la = p.GetTokenStream().LA(1)

		if !(_la == LEQLParserPERCENTILE || _la == LEQLParserPCTL) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(374)
		p.Match(LEQLParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(375)
		p.Match(LEQLParserNUMBER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(376)
		p.Match(LEQLParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(379)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == LEQLParserCOLON {
		{
			p.SetState(377)
			p.Match(LEQLParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(378)
			p.FieldName()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISortClauseContext is an interface to support dynamic dispatch.
type ISortClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SORT() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	AllSortSpec() []ISortSpecContext
	SortSpec(i int) ISortSpecContext
	RPAREN() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsSortClauseContext differentiates from other interfaces.
	IsSortClauseContext()
}

type SortClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySortClauseContext() *SortClauseContext {
	var p = new(SortClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_sortClause
	return p
}

func InitEmptySortClauseContext(p *SortClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_sortClause
}

func (*SortClauseContext) IsSortClauseContext() {}

func NewSortClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SortClauseContext {
	var p = new(SortClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LEQLParserRULE_sortClause

	return p
}

func (s *SortClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *SortClauseContext) SORT() antlr.TerminalNode {
	return s.GetToken(LEQLParserSORT, 0)
}

func (s *SortClauseContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(LEQLParserLPAREN, 0)
}

func (s *SortClauseContext) AllSortSpec() []ISortSpecContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISortSpecContext); ok {
			len++
		}
	}

	tst := make([]ISortSpecContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISortSpecContext); ok {
			tst[i] = t.(ISortSpecContext)
			i++
		}
	}

	return tst
}

func (s *SortClauseContext) SortSpec(i int) ISortSpecContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISortSpecContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISortSpecContext)
}

func (s *SortClauseContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(LEQLParserRPAREN, 0)
}

func (s *SortClauseContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(LEQLParserCOMMA)
}

func (s *SortClauseContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(LEQLParserCOMMA, i)
}

func (s *SortClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SortClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SortClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterSortClause(s)
	}
}

func (s *SortClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitSortClause(s)
	}
}

func (s *SortClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitSortClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LEQLParser) SortClause() (localctx ISortClauseContext) {
	localctx = NewSortClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, LEQLParserRULE_sortClause)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(381)
		p.Match(LEQLParserSORT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(382)
		p.Match(LEQLParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(383)
		p.SortSpec()
	}
	p.SetState(388)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == LEQLParserCOMMA {
		{
			p.SetState(384)
			p.Match(LEQLParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(385)
			p.SortSpec()
		}

		p.SetState(390)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(391)
		p.Match(LEQLParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISortSpecContext is an interface to support dynamic dispatch.
type ISortSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SortDirection() ISortDirectionContext
	HASH() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode

	// IsSortSpecContext differentiates from other interfaces.
	IsSortSpecContext()
}

type SortSpecContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySortSpecContext() *SortSpecContext {
	var p = new(SortSpecContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_sortSpec
	return p
}

func InitEmptySortSpecContext(p *SortSpecContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_sortSpec
}

func (*SortSpecContext) IsSortSpecContext() {}

func NewSortSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SortSpecContext {
	var p = new(SortSpecContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LEQLParserRULE_sortSpec

	return p
}

func (s *SortSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *SortSpecContext) SortDirection() ISortDirectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISortDirectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISortDirectionContext)
}

func (s *SortSpecContext) HASH() antlr.TerminalNode {
	return s.GetToken(LEQLParserHASH, 0)
}

func (s *SortSpecContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(LEQLParserIDENTIFIER, 0)
}

func (s *SortSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SortSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SortSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterSortSpec(s)
	}
}

func (s *SortSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitSortSpec(s)
	}
}

func (s *SortSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitSortSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LEQLParser) SortSpec() (localctx ISortSpecContext) {
	localctx = NewSortSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, LEQLParserRULE_sortSpec)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(393)
		p.SortDirection()
	}
	p.SetState(397)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	switch p.GetTokenStream().LA(1) {
	case LEQLParserHASH:
		{
			p.SetState(394)
			p.Match(LEQLParserHASH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(395)
			p.Match(LEQLParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case LEQLParserIDENTIFIER:
		{
			p.SetState(396)
			p.Match(LEQLParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case LEQLParserRPAREN, LEQLParserCOMMA:

	default:
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISortDirectionContext is an interface to support dynamic dispatch.
type ISortDirectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ASC() antlr.TerminalNode
	ASCENDING() antlr.TerminalNode
	DESC() antlr.TerminalNode
	DESCENDING() antlr.TerminalNode

	// IsSortDirectionContext differentiates from other interfaces.
	IsSortDirectionContext()
}

type SortDirectionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySortDirectionContext() *SortDirectionContext {
	var p = new(SortDirectionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_sortDirection
	return p
}

func InitEmptySortDirectionContext(p *SortDirectionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_sortDirection
}

func (*SortDirectionContext) IsSortDirectionContext() {}

func NewSortDirectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SortDirectionContext {
	var p = new(SortDirectionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LEQLParserRULE_sortDirection

	return p
}

func (s *SortDirectionContext) GetParser() antlr.Parser { return s.parser }

func (s *SortDirectionContext) ASC() antlr.TerminalNode {
	return s.GetToken(LEQLParserASC, 0)
}

func (s *SortDirectionContext) ASCENDING() antlr.TerminalNode {
	return s.GetToken(LEQLParserASCENDING, 0)
}

func (s *SortDirectionContext) DESC() antlr.TerminalNode {
	return s.GetToken(LEQLParserDESC, 0)
}

func (s *SortDirectionContext) DESCENDING() antlr.TerminalNode {
	return s.GetToken(LEQLParserDESCENDING, 0)
}

func (s *SortDirectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SortDirectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SortDirectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterSortDirection(s)
	}
}

func (s *SortDirectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitSortDirection(s)
	}
}

func (s *SortDirectionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitSortDirection(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LEQLParser) SortDirection() (localctx ISortDirectionContext) {
	localctx = NewSortDirectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, LEQLParserRULE_sortDirection)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(399)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&16492674416640) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILimitClauseContext is an interface to support dynamic dispatch.
type ILimitClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LIMIT() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	AllNUMBER() []antlr.TerminalNode
	NUMBER(i int) antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsLimitClauseContext differentiates from other interfaces.
	IsLimitClauseContext()
}

type LimitClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLimitClauseContext() *LimitClauseContext {
	var p = new(LimitClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_limitClause
	return p
}

func InitEmptyLimitClauseContext(p *LimitClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_limitClause
}

func (*LimitClauseContext) IsLimitClauseContext() {}

func NewLimitClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LimitClauseContext {
	var p = new(LimitClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LEQLParserRULE_limitClause

	return p
}

func (s *LimitClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *LimitClauseContext) LIMIT() antlr.TerminalNode {
	return s.GetToken(LEQLParserLIMIT, 0)
}

func (s *LimitClauseContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(LEQLParserLPAREN, 0)
}

func (s *LimitClauseContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(LEQLParserNUMBER)
}

func (s *LimitClauseContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(LEQLParserNUMBER, i)
}

func (s *LimitClauseContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(LEQLParserRPAREN, 0)
}

func (s *LimitClauseContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(LEQLParserCOMMA)
}

func (s *LimitClauseContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(LEQLParserCOMMA, i)
}

func (s *LimitClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LimitClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LimitClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterLimitClause(s)
	}
}

func (s *LimitClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitLimitClause(s)
	}
}

func (s *LimitClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitLimitClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LEQLParser) LimitClause() (localctx ILimitClauseContext) {
	localctx = NewLimitClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, LEQLParserRULE_limitClause)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(401)
		p.Match(LEQLParserLIMIT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(402)
		p.Match(LEQLParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(403)
		p.Match(LEQLParserNUMBER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(408)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == LEQLParserCOMMA {
		{
			p.SetState(404)
			p.Match(LEQLParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(405)
			p.Match(LEQLParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(410)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(411)
		p.Match(LEQLParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITimesliceClauseContext is an interface to support dynamic dispatch.
type ITimesliceClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TIMESLICE() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	TIME_UNIT() antlr.TerminalNode

	// IsTimesliceClauseContext differentiates from other interfaces.
	IsTimesliceClauseContext()
}

type TimesliceClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimesliceClauseContext() *TimesliceClauseContext {
	var p = new(TimesliceClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_timesliceClause
	return p
}

func InitEmptyTimesliceClauseContext(p *TimesliceClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LEQLParserRULE_timesliceClause
}

func (*TimesliceClauseContext) IsTimesliceClauseContext() {}

func NewTimesliceClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimesliceClauseContext {
	var p = new(TimesliceClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LEQLParserRULE_timesliceClause

	return p
}

func (s *TimesliceClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *TimesliceClauseContext) TIMESLICE() antlr.TerminalNode {
	return s.GetToken(LEQLParserTIMESLICE, 0)
}

func (s *TimesliceClauseContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(LEQLParserLPAREN, 0)
}

func (s *TimesliceClauseContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(LEQLParserRPAREN, 0)
}

func (s *TimesliceClauseContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(LEQLParserNUMBER, 0)
}

func (s *TimesliceClauseContext) TIME_UNIT() antlr.TerminalNode {
	return s.GetToken(LEQLParserTIME_UNIT, 0)
}

func (s *TimesliceClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimesliceClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimesliceClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.EnterTimesliceClause(s)
	}
}

func (s *TimesliceClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LEQLParserListener); ok {
		listenerT.ExitTimesliceClause(s)
	}
}

func (s *TimesliceClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LEQLParserVisitor:
		return t.VisitTimesliceClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LEQLParser) TimesliceClause() (localctx ITimesliceClauseContext) {
	localctx = NewTimesliceClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, LEQLParserRULE_timesliceClause)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(413)
		p.Match(LEQLParserTIMESLICE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(414)
		p.Match(LEQLParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(415)
		_la = p.GetTokenStream().LA(1)

		if !(_la == LEQLParserNUMBER || _la == LEQLParserTIME_UNIT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(416)
		p.Match(LEQLParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *LEQLParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 6:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *LEQLParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
