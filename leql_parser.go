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
		"", "", "", "", "", "", "", "", "", "", "'!=='", "'!='", "'>='", "'<='",
		"'=='", "'='", "'>'", "'<'", "'!'", "'('", "')'", "'['", "']'", "','",
		"':'", "'.'", "'*'", "'#'",
	}
	staticData.SymbolicNames = []string{
		"", "SELECT", "WHERE", "GROUPBY", "HAVING", "CALCULATE", "SORT", "LIMIT",
		"TIMESLICE", "FROM", "LOOSE", "AND", "OR", "NOT", "ICONTAINS_ALL_OP",
		"ICONTAINS_ANY_OP", "CONTAINS_ALL_OP", "CONTAINS_ANY_OP", "ISTARTS_WITH_ANY_OP",
		"STARTS_WITH_ANY_OP", "ISTARTS_WITH_OP", "STARTS_WITH_OP", "ICONTAINS_OP",
		"CONTAINS_OP", "IIN_OP", "IN_OP", "NOCASE", "IP_FUNC", "ALL_FUNC", "STANDARDDEVIATION",
		"PERCENTILE", "AVERAGE", "UNIQUE", "COUNT", "BYTES", "PCTL", "SUM",
		"MIN", "MAX", "SD", "ASCENDING", "DESCENDING", "ASC", "DESC", "STRICT_NEQ",
		"NEQ", "GTE", "LTE", "STRICT_EQ", "EQ", "GT", "LT", "BANG", "LPAREN",
		"RPAREN", "LBRACKET", "RBRACKET", "COMMA", "COLON", "DOT", "STAR", "HASH",
		"IP_CIDR", "REGEX", "TRIPLE_SINGLE_STRING", "TRIPLE_DOUBLE_STRING",
		"DOUBLE_STRING", "SINGLE_STRING", "NUMBER", "TIME_UNIT", "VARIABLE",
		"IDENTIFIER", "WS",
	}
	staticData.RuleNames = []string{
		"query", "fromClause", "selectClause", "selectFieldList", "selectField",
		"whereClause", "expression", "condition", "keywordSearch", "regexSearch",
		"fieldList", "allFieldList", "fieldName", "comparisonOp", "stringOp",
		"setOp", "listStringOp", "value", "nocaseValue", "ipValue", "valueList",
		"valueListItem", "groupbyClause", "havingClause", "havingCondition",
		"calculateClause", "calcExpr", "calcFunction", "calcFunctionWithField",
		"percentileFunction", "sortClause", "sortSpec", "sortDirection", "limitClause",
		"timesliceClause",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 72, 401, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 1, 0, 3, 0, 72, 8, 0, 1,
		0, 3, 0, 75, 8, 0, 1, 0, 3, 0, 78, 8, 0, 1, 0, 3, 0, 81, 8, 0, 1, 0, 3,
		0, 84, 8, 0, 1, 0, 3, 0, 87, 8, 0, 1, 0, 3, 0, 90, 8, 0, 1, 0, 3, 0, 93,
		8, 0, 1, 0, 3, 0, 96, 8, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 5, 3, 113, 8, 3, 10, 3, 12,
		3, 116, 9, 3, 1, 4, 1, 4, 1, 4, 3, 4, 121, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 3, 5, 128, 8, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 147, 8, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 157, 8, 6, 10, 6,
		12, 6, 160, 9, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 3, 7, 245, 8, 7, 1, 8, 1, 8, 3, 8, 249, 8, 8, 1, 9, 1, 9, 1,
		10, 1, 10, 1, 10, 5, 10, 256, 8, 10, 10, 10, 12, 10, 259, 9, 10, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 5, 11, 266, 8, 11, 10, 11, 12, 11, 269, 9,
		11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15,
		1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1,
		19, 1, 19, 1, 19, 3, 19, 294, 8, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20,
		1, 20, 5, 20, 302, 8, 20, 10, 20, 12, 20, 305, 9, 20, 1, 20, 1, 20, 1,
		21, 1, 21, 3, 21, 311, 8, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 5, 22,
		318, 8, 22, 10, 22, 12, 22, 321, 9, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24,
		3, 24, 338, 8, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1,
		26, 3, 26, 348, 8, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 29,
		1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 3, 29, 362, 8, 29, 1, 30, 1, 30, 1,
		30, 1, 30, 1, 30, 5, 30, 369, 8, 30, 10, 30, 12, 30, 372, 9, 30, 1, 30,
		1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 3, 31, 380, 8, 31, 1, 32, 1, 32, 1,
		33, 1, 33, 1, 33, 1, 33, 1, 33, 5, 33, 389, 8, 33, 10, 33, 12, 33, 392,
		9, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 0, 1, 12,
		35, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
		36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 0,
		11, 6, 0, 9, 10, 27, 28, 33, 34, 36, 39, 66, 68, 71, 71, 1, 0, 44, 51,
		1, 0, 20, 23, 1, 0, 24, 25, 1, 0, 14, 19, 1, 0, 62, 71, 1, 0, 33, 34, 3,
		0, 29, 29, 31, 34, 36, 39, 2, 0, 30, 30, 35, 35, 1, 0, 40, 43, 1, 0, 68,
		69, 420, 0, 71, 1, 0, 0, 0, 2, 99, 1, 0, 0, 0, 4, 104, 1, 0, 0, 0, 6, 109,
		1, 0, 0, 0, 8, 117, 1, 0, 0, 0, 10, 122, 1, 0, 0, 0, 12, 146, 1, 0, 0,
		0, 14, 244, 1, 0, 0, 0, 16, 248, 1, 0, 0, 0, 18, 250, 1, 0, 0, 0, 20, 252,
		1, 0, 0, 0, 22, 260, 1, 0, 0, 0, 24, 272, 1, 0, 0, 0, 26, 274, 1, 0, 0,
		0, 28, 276, 1, 0, 0, 0, 30, 278, 1, 0, 0, 0, 32, 280, 1, 0, 0, 0, 34, 282,
		1, 0, 0, 0, 36, 284, 1, 0, 0, 0, 38, 289, 1, 0, 0, 0, 40, 297, 1, 0, 0,
		0, 42, 310, 1, 0, 0, 0, 44, 312, 1, 0, 0, 0, 46, 324, 1, 0, 0, 0, 48, 337,
		1, 0, 0, 0, 50, 339, 1, 0, 0, 0, 52, 347, 1, 0, 0, 0, 54, 349, 1, 0, 0,
		0, 56, 351, 1, 0, 0, 0, 58, 355, 1, 0, 0, 0, 60, 363, 1, 0, 0, 0, 62, 375,
		1, 0, 0, 0, 64, 381, 1, 0, 0, 0, 66, 383, 1, 0, 0, 0, 68, 395, 1, 0, 0,
		0, 70, 72, 3, 2, 1, 0, 71, 70, 1, 0, 0, 0, 71, 72, 1, 0, 0, 0, 72, 74,
		1, 0, 0, 0, 73, 75, 3, 4, 2, 0, 74, 73, 1, 0, 0, 0, 74, 75, 1, 0, 0, 0,
		75, 77, 1, 0, 0, 0, 76, 78, 3, 10, 5, 0, 77, 76, 1, 0, 0, 0, 77, 78, 1,
		0, 0, 0, 78, 80, 1, 0, 0, 0, 79, 81, 3, 44, 22, 0, 80, 79, 1, 0, 0, 0,
		80, 81, 1, 0, 0, 0, 81, 83, 1, 0, 0, 0, 82, 84, 3, 50, 25, 0, 83, 82, 1,
		0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 86, 1, 0, 0, 0, 85, 87, 3, 46, 23, 0,
		86, 85, 1, 0, 0, 0, 86, 87, 1, 0, 0, 0, 87, 89, 1, 0, 0, 0, 88, 90, 3,
		60, 30, 0, 89, 88, 1, 0, 0, 0, 89, 90, 1, 0, 0, 0, 90, 92, 1, 0, 0, 0,
		91, 93, 3, 66, 33, 0, 92, 91, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 95, 1,
		0, 0, 0, 94, 96, 3, 68, 34, 0, 95, 94, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0,
		96, 97, 1, 0, 0, 0, 97, 98, 5, 0, 0, 1, 98, 1, 1, 0, 0, 0, 99, 100, 5,
		9, 0, 0, 100, 101, 5, 53, 0, 0, 101, 102, 3, 12, 6, 0, 102, 103, 5, 54,
		0, 0, 103, 3, 1, 0, 0, 0, 104, 105, 5, 1, 0, 0, 105, 106, 5, 53, 0, 0,
		106, 107, 3, 6, 3, 0, 107, 108, 5, 54, 0, 0, 108, 5, 1, 0, 0, 0, 109, 114,
		3, 8, 4, 0, 110, 111, 5, 57, 0, 0, 111, 113, 3, 8, 4, 0, 112, 110, 1, 0,
		0, 0, 113, 116, 1, 0, 0, 0, 114, 112, 1, 0, 0, 0, 114, 115, 1, 0, 0, 0,
		115, 7, 1, 0, 0, 0, 116, 114, 1, 0, 0, 0, 117, 120, 3, 24, 12, 0, 118,
		119, 5, 71, 0, 0, 119, 121, 5, 71, 0, 0, 120, 118, 1, 0, 0, 0, 120, 121,
		1, 0, 0, 0, 121, 9, 1, 0, 0, 0, 122, 123, 5, 2, 0, 0, 123, 124, 5, 53,
		0, 0, 124, 127, 3, 12, 6, 0, 125, 126, 5, 57, 0, 0, 126, 128, 5, 10, 0,
		0, 127, 125, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 129,
		130, 5, 54, 0, 0, 130, 11, 1, 0, 0, 0, 131, 132, 6, 6, -1, 0, 132, 133,
		5, 53, 0, 0, 133, 134, 3, 12, 6, 0, 134, 135, 5, 54, 0, 0, 135, 147, 1,
		0, 0, 0, 136, 137, 5, 2, 0, 0, 137, 138, 5, 53, 0, 0, 138, 139, 3, 12,
		6, 0, 139, 140, 5, 54, 0, 0, 140, 147, 1, 0, 0, 0, 141, 142, 5, 13, 0,
		0, 142, 147, 3, 12, 6, 7, 143, 147, 3, 14, 7, 0, 144, 147, 3, 16, 8, 0,
		145, 147, 3, 18, 9, 0, 146, 131, 1, 0, 0, 0, 146, 136, 1, 0, 0, 0, 146,
		141, 1, 0, 0, 0, 146, 143, 1, 0, 0, 0, 146, 144, 1, 0, 0, 0, 146, 145,
		1, 0, 0, 0, 147, 158, 1, 0, 0, 0, 148, 149, 10, 6, 0, 0, 149, 150, 5, 11,
		0, 0, 150, 157, 3, 12, 6, 7, 151, 152, 10, 5, 0, 0, 152, 157, 3, 12, 6,
		6, 153, 154, 10, 4, 0, 0, 154, 155, 5, 12, 0, 0, 155, 157, 3, 12, 6, 5,
		156, 148, 1, 0, 0, 0, 156, 151, 1, 0, 0, 0, 156, 153, 1, 0, 0, 0, 157,
		160, 1, 0, 0, 0, 158, 156, 1, 0, 0, 0, 158, 159, 1, 0, 0, 0, 159, 13, 1,
		0, 0, 0, 160, 158, 1, 0, 0, 0, 161, 162, 3, 20, 10, 0, 162, 163, 3, 26,
		13, 0, 163, 164, 3, 34, 17, 0, 164, 245, 1, 0, 0, 0, 165, 166, 3, 20, 10,
		0, 166, 167, 3, 28, 14, 0, 167, 168, 3, 34, 17, 0, 168, 245, 1, 0, 0, 0,
		169, 170, 3, 20, 10, 0, 170, 171, 3, 30, 15, 0, 171, 172, 3, 40, 20, 0,
		172, 245, 1, 0, 0, 0, 173, 174, 3, 20, 10, 0, 174, 175, 3, 32, 16, 0, 175,
		176, 3, 40, 20, 0, 176, 245, 1, 0, 0, 0, 177, 178, 5, 13, 0, 0, 178, 179,
		3, 20, 10, 0, 179, 180, 3, 26, 13, 0, 180, 181, 3, 34, 17, 0, 181, 245,
		1, 0, 0, 0, 182, 183, 5, 13, 0, 0, 183, 184, 3, 20, 10, 0, 184, 185, 3,
		28, 14, 0, 185, 186, 3, 34, 17, 0, 186, 245, 1, 0, 0, 0, 187, 188, 5, 13,
		0, 0, 188, 189, 3, 20, 10, 0, 189, 190, 3, 30, 15, 0, 190, 191, 3, 40,
		20, 0, 191, 245, 1, 0, 0, 0, 192, 193, 5, 13, 0, 0, 193, 194, 3, 20, 10,
		0, 194, 195, 3, 32, 16, 0, 195, 196, 3, 40, 20, 0, 196, 245, 1, 0, 0, 0,
		197, 198, 3, 20, 10, 0, 198, 199, 5, 13, 0, 0, 199, 200, 3, 26, 13, 0,
		200, 201, 3, 34, 17, 0, 201, 245, 1, 0, 0, 0, 202, 203, 3, 20, 10, 0, 203,
		204, 5, 13, 0, 0, 204, 205, 3, 28, 14, 0, 205, 206, 3, 34, 17, 0, 206,
		245, 1, 0, 0, 0, 207, 208, 3, 20, 10, 0, 208, 209, 5, 13, 0, 0, 209, 210,
		3, 30, 15, 0, 210, 211, 3, 40, 20, 0, 211, 245, 1, 0, 0, 0, 212, 213, 3,
		20, 10, 0, 213, 214, 5, 13, 0, 0, 214, 215, 3, 32, 16, 0, 215, 216, 3,
		40, 20, 0, 216, 245, 1, 0, 0, 0, 217, 218, 3, 20, 10, 0, 218, 219, 3, 26,
		13, 0, 219, 220, 3, 36, 18, 0, 220, 245, 1, 0, 0, 0, 221, 222, 3, 20, 10,
		0, 222, 223, 3, 26, 13, 0, 223, 224, 3, 38, 19, 0, 224, 245, 1, 0, 0, 0,
		225, 226, 3, 20, 10, 0, 226, 227, 3, 30, 15, 0, 227, 228, 3, 40, 20, 0,
		228, 245, 1, 0, 0, 0, 229, 230, 3, 22, 11, 0, 230, 231, 3, 26, 13, 0, 231,
		232, 3, 34, 17, 0, 232, 245, 1, 0, 0, 0, 233, 234, 3, 22, 11, 0, 234, 235,
		3, 30, 15, 0, 235, 236, 3, 40, 20, 0, 236, 245, 1, 0, 0, 0, 237, 238, 3,
		22, 11, 0, 238, 239, 3, 26, 13, 0, 239, 240, 3, 36, 18, 0, 240, 245, 1,
		0, 0, 0, 241, 245, 5, 71, 0, 0, 242, 245, 5, 68, 0, 0, 243, 245, 5, 70,
		0, 0, 244, 161, 1, 0, 0, 0, 244, 165, 1, 0, 0, 0, 244, 169, 1, 0, 0, 0,
		244, 173, 1, 0, 0, 0, 244, 177, 1, 0, 0, 0, 244, 182, 1, 0, 0, 0, 244,
		187, 1, 0, 0, 0, 244, 192, 1, 0, 0, 0, 244, 197, 1, 0, 0, 0, 244, 202,
		1, 0, 0, 0, 244, 207, 1, 0, 0, 0, 244, 212, 1, 0, 0, 0, 244, 217, 1, 0,
		0, 0, 244, 221, 1, 0, 0, 0, 244, 225, 1, 0, 0, 0, 244, 229, 1, 0, 0, 0,
		244, 233, 1, 0, 0, 0, 244, 237, 1, 0, 0, 0, 244, 241, 1, 0, 0, 0, 244,
		242, 1, 0, 0, 0, 244, 243, 1, 0, 0, 0, 245, 15, 1, 0, 0, 0, 246, 249, 5,
		66, 0, 0, 247, 249, 5, 67, 0, 0, 248, 246, 1, 0, 0, 0, 248, 247, 1, 0,
		0, 0, 249, 17, 1, 0, 0, 0, 250, 251, 5, 63, 0, 0, 251, 19, 1, 0, 0, 0,
		252, 257, 3, 24, 12, 0, 253, 254, 5, 57, 0, 0, 254, 256, 3, 24, 12, 0,
		255, 253, 1, 0, 0, 0, 256, 259, 1, 0, 0, 0, 257, 255, 1, 0, 0, 0, 257,
		258, 1, 0, 0, 0, 258, 21, 1, 0, 0, 0, 259, 257, 1, 0, 0, 0, 260, 261, 5,
		28, 0, 0, 261, 262, 5, 53, 0, 0, 262, 267, 3, 24, 12, 0, 263, 264, 5, 57,
		0, 0, 264, 266, 3, 24, 12, 0, 265, 263, 1, 0, 0, 0, 266, 269, 1, 0, 0,
		0, 267, 265, 1, 0, 0, 0, 267, 268, 1, 0, 0, 0, 268, 270, 1, 0, 0, 0, 269,
		267, 1, 0, 0, 0, 270, 271, 5, 54, 0, 0, 271, 23, 1, 0, 0, 0, 272, 273,
		7, 0, 0, 0, 273, 25, 1, 0, 0, 0, 274, 275, 7, 1, 0, 0, 275, 27, 1, 0, 0,
		0, 276, 277, 7, 2, 0, 0, 277, 29, 1, 0, 0, 0, 278, 279, 7, 3, 0, 0, 279,
		31, 1, 0, 0, 0, 280, 281, 7, 4, 0, 0, 281, 33, 1, 0, 0, 0, 282, 283, 7,
		5, 0, 0, 283, 35, 1, 0, 0, 0, 284, 285, 5, 26, 0, 0, 285, 286, 5, 53, 0,
		0, 286, 287, 3, 34, 17, 0, 287, 288, 5, 54, 0, 0, 288, 37, 1, 0, 0, 0,
		289, 290, 5, 27, 0, 0, 290, 293, 5, 53, 0, 0, 291, 294, 5, 62, 0, 0, 292,
		294, 3, 34, 17, 0, 293, 291, 1, 0, 0, 0, 293, 292, 1, 0, 0, 0, 294, 295,
		1, 0, 0, 0, 295, 296, 5, 54, 0, 0, 296, 39, 1, 0, 0, 0, 297, 298, 5, 55,
		0, 0, 298, 303, 3, 42, 21, 0, 299, 300, 5, 57, 0, 0, 300, 302, 3, 42, 21,
		0, 301, 299, 1, 0, 0, 0, 302, 305, 1, 0, 0, 0, 303, 301, 1, 0, 0, 0, 303,
		304, 1, 0, 0, 0, 304, 306, 1, 0, 0, 0, 305, 303, 1, 0, 0, 0, 306, 307,
		5, 56, 0, 0, 307, 41, 1, 0, 0, 0, 308, 311, 3, 38, 19, 0, 309, 311, 3,
		34, 17, 0, 310, 308, 1, 0, 0, 0, 310, 309, 1, 0, 0, 0, 311, 43, 1, 0, 0,
		0, 312, 313, 5, 3, 0, 0, 313, 314, 5, 53, 0, 0, 314, 319, 3, 24, 12, 0,
		315, 316, 5, 57, 0, 0, 316, 318, 3, 24, 12, 0, 317, 315, 1, 0, 0, 0, 318,
		321, 1, 0, 0, 0, 319, 317, 1, 0, 0, 0, 319, 320, 1, 0, 0, 0, 320, 322,
		1, 0, 0, 0, 321, 319, 1, 0, 0, 0, 322, 323, 5, 54, 0, 0, 323, 45, 1, 0,
		0, 0, 324, 325, 5, 4, 0, 0, 325, 326, 5, 53, 0, 0, 326, 327, 3, 48, 24,
		0, 327, 328, 5, 54, 0, 0, 328, 47, 1, 0, 0, 0, 329, 330, 3, 54, 27, 0,
		330, 331, 3, 26, 13, 0, 331, 332, 3, 34, 17, 0, 332, 338, 1, 0, 0, 0, 333,
		334, 3, 56, 28, 0, 334, 335, 3, 26, 13, 0, 335, 336, 3, 34, 17, 0, 336,
		338, 1, 0, 0, 0, 337, 329, 1, 0, 0, 0, 337, 333, 1, 0, 0, 0, 338, 49, 1,
		0, 0, 0, 339, 340, 5, 5, 0, 0, 340, 341, 5, 53, 0, 0, 341, 342, 3, 52,
		26, 0, 342, 343, 5, 54, 0, 0, 343, 51, 1, 0, 0, 0, 344, 348, 3, 54, 27,
		0, 345, 348, 3, 56, 28, 0, 346, 348, 3, 58, 29, 0, 347, 344, 1, 0, 0, 0,
		347, 345, 1, 0, 0, 0, 347, 346, 1, 0, 0, 0, 348, 53, 1, 0, 0, 0, 349, 350,
		7, 6, 0, 0, 350, 55, 1, 0, 0, 0, 351, 352, 7, 7, 0, 0, 352, 353, 5, 58,
		0, 0, 353, 354, 3, 24, 12, 0, 354, 57, 1, 0, 0, 0, 355, 356, 7, 8, 0, 0,
		356, 357, 5, 53, 0, 0, 357, 358, 5, 68, 0, 0, 358, 361, 5, 54, 0, 0, 359,
		360, 5, 58, 0, 0, 360, 362, 3, 24, 12, 0, 361, 359, 1, 0, 0, 0, 361, 362,
		1, 0, 0, 0, 362, 59, 1, 0, 0, 0, 363, 364, 5, 6, 0, 0, 364, 365, 5, 53,
		0, 0, 365, 370, 3, 62, 31, 0, 366, 367, 5, 57, 0, 0, 367, 369, 3, 62, 31,
		0, 368, 366, 1, 0, 0, 0, 369, 372, 1, 0, 0, 0, 370, 368, 1, 0, 0, 0, 370,
		371, 1, 0, 0, 0, 371, 373, 1, 0, 0, 0, 372, 370, 1, 0, 0, 0, 373, 374,
		5, 54, 0, 0, 374, 61, 1, 0, 0, 0, 375, 379, 3, 64, 32, 0, 376, 377, 5,
		61, 0, 0, 377, 380, 5, 71, 0, 0, 378, 380, 5, 71, 0, 0, 379, 376, 1, 0,
		0, 0, 379, 378, 1, 0, 0, 0, 379, 380, 1, 0, 0, 0, 380, 63, 1, 0, 0, 0,
		381, 382, 7, 9, 0, 0, 382, 65, 1, 0, 0, 0, 383, 384, 5, 7, 0, 0, 384, 385,
		5, 53, 0, 0, 385, 390, 5, 68, 0, 0, 386, 387, 5, 57, 0, 0, 387, 389, 5,
		68, 0, 0, 388, 386, 1, 0, 0, 0, 389, 392, 1, 0, 0, 0, 390, 388, 1, 0, 0,
		0, 390, 391, 1, 0, 0, 0, 391, 393, 1, 0, 0, 0, 392, 390, 1, 0, 0, 0, 393,
		394, 5, 54, 0, 0, 394, 67, 1, 0, 0, 0, 395, 396, 5, 8, 0, 0, 396, 397,
		5, 53, 0, 0, 397, 398, 7, 10, 0, 0, 398, 399, 5, 54, 0, 0, 399, 69, 1,
		0, 0, 0, 29, 71, 74, 77, 80, 83, 86, 89, 92, 95, 114, 120, 127, 146, 156,
		158, 244, 248, 257, 267, 293, 303, 310, 319, 337, 347, 361, 370, 379, 390,
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
	LEQLParserNEQ                  = 45
	LEQLParserGTE                  = 46
	LEQLParserLTE                  = 47
	LEQLParserSTRICT_EQ            = 48
	LEQLParserEQ                   = 49
	LEQLParserGT                   = 50
	LEQLParserLT                   = 51
	LEQLParserBANG                 = 52
	LEQLParserLPAREN               = 53
	LEQLParserRPAREN               = 54
	LEQLParserLBRACKET             = 55
	LEQLParserRBRACKET             = 56
	LEQLParserCOMMA                = 57
	LEQLParserCOLON                = 58
	LEQLParserDOT                  = 59
	LEQLParserSTAR                 = 60
	LEQLParserHASH                 = 61
	LEQLParserIP_CIDR              = 62
	LEQLParserREGEX                = 63
	LEQLParserTRIPLE_SINGLE_STRING = 64
	LEQLParserTRIPLE_DOUBLE_STRING = 65
	LEQLParserDOUBLE_STRING        = 66
	LEQLParserSINGLE_STRING        = 67
	LEQLParserNUMBER               = 68
	LEQLParserTIME_UNIT            = 69
	LEQLParserVARIABLE             = 70
	LEQLParserIDENTIFIER           = 71
	LEQLParserWS                   = 72
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
	LEQLParserRULE_stringOp              = 14
	LEQLParserRULE_setOp                 = 15
	LEQLParserRULE_listStringOp          = 16
	LEQLParserRULE_value                 = 17
	LEQLParserRULE_nocaseValue           = 18
	LEQLParserRULE_ipValue               = 19
	LEQLParserRULE_valueList             = 20
	LEQLParserRULE_valueListItem         = 21
	LEQLParserRULE_groupbyClause         = 22
	LEQLParserRULE_havingClause          = 23
	LEQLParserRULE_havingCondition       = 24
	LEQLParserRULE_calculateClause       = 25
	LEQLParserRULE_calcExpr              = 26
	LEQLParserRULE_calcFunction          = 27
	LEQLParserRULE_calcFunctionWithField = 28
	LEQLParserRULE_percentileFunction    = 29
	LEQLParserRULE_sortClause            = 30
	LEQLParserRULE_sortSpec              = 31
	LEQLParserRULE_sortDirection         = 32
	LEQLParserRULE_limitClause           = 33
	LEQLParserRULE_timesliceClause       = 34
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
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == LEQLParserFROM {
		{
			p.SetState(70)
			p.FromClause()
		}

	}
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == LEQLParserSELECT {
		{
			p.SetState(73)
			p.SelectClause()
		}

	}
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == LEQLParserWHERE {
		{
			p.SetState(76)
			p.WhereClause()
		}

	}
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == LEQLParserGROUPBY {
		{
			p.SetState(79)
			p.GroupbyClause()
		}

	}
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == LEQLParserCALCULATE {
		{
			p.SetState(82)
			p.CalculateClause()
		}

	}
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == LEQLParserHAVING {
		{
			p.SetState(85)
			p.HavingClause()
		}

	}
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == LEQLParserSORT {
		{
			p.SetState(88)
			p.SortClause()
		}

	}
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == LEQLParserLIMIT {
		{
			p.SetState(91)
			p.LimitClause()
		}

	}
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == LEQLParserTIMESLICE {
		{
			p.SetState(94)
			p.TimesliceClause()
		}

	}
	{
		p.SetState(97)
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
		p.SetState(99)
		p.Match(LEQLParserFROM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(100)
		p.Match(LEQLParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(101)
		p.expression(0)
	}
	{
		p.SetState(102)
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
		p.SetState(104)
		p.Match(LEQLParserSELECT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(105)
		p.Match(LEQLParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(106)
		p.SelectFieldList()
	}
	{
		p.SetState(107)
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
		p.SetState(109)
		p.SelectField()
	}
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == LEQLParserCOMMA {
		{
			p.SetState(110)
			p.Match(LEQLParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
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
		p.SetState(117)
		p.FieldName()
	}
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == LEQLParserIDENTIFIER {
		{
			p.SetState(118)
			p.Match(LEQLParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(119)
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
		p.SetState(122)
		p.Match(LEQLParserWHERE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(123)
		p.Match(LEQLParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(124)
		p.expression(0)
	}
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == LEQLParserCOMMA {
		{
			p.SetState(125)
			p.Match(LEQLParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(126)
			p.Match(LEQLParserLOOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(129)
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
	p.SetState(146)
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
			p.SetState(132)
			p.Match(LEQLParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(133)
			p.expression(0)
		}
		{
			p.SetState(134)
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
			p.SetState(136)
			p.Match(LEQLParserWHERE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(137)
			p.Match(LEQLParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(138)
			p.expression(0)
		}
		{
			p.SetState(139)
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
			p.SetState(141)
			p.Match(LEQLParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(142)
			p.expression(7)
		}

	case 4:
		localctx = NewConditionExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(143)
			p.Condition()
		}

	case 5:
		localctx = NewKeywordExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(144)
			p.KeywordSearch()
		}

	case 6:
		localctx = NewRegexExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(145)
			p.RegexSearch()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(158)
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
			p.SetState(156)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAndExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LEQLParserRULE_expression)
				p.SetState(148)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(149)
					p.Match(LEQLParserAND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(150)
					p.expression(7)
				}

			case 2:
				localctx = NewImplicitAndExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LEQLParserRULE_expression)
				p.SetState(151)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(152)
					p.expression(6)
				}

			case 3:
				localctx = NewOrExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LEQLParserRULE_expression)
				p.SetState(153)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(154)
					p.Match(LEQLParserOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(155)
					p.expression(5)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(160)
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
	p.SetState(244)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		localctx = NewComparisonConditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(161)
			p.FieldList()
		}
		{
			p.SetState(162)
			p.ComparisonOp()
		}
		{
			p.SetState(163)
			p.Value()
		}

	case 2:
		localctx = NewStringConditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(165)
			p.FieldList()
		}
		{
			p.SetState(166)
			p.StringOp()
		}
		{
			p.SetState(167)
			p.Value()
		}

	case 3:
		localctx = NewSetConditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(169)
			p.FieldList()
		}
		{
			p.SetState(170)
			p.SetOp()
		}
		{
			p.SetState(171)
			p.ValueList()
		}

	case 4:
		localctx = NewListStringConditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(173)
			p.FieldList()
		}
		{
			p.SetState(174)
			p.ListStringOp()
		}
		{
			p.SetState(175)
			p.ValueList()
		}

	case 5:
		localctx = NewNegatedComparisonConditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(177)
			p.Match(LEQLParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(178)
			p.FieldList()
		}
		{
			p.SetState(179)
			p.ComparisonOp()
		}
		{
			p.SetState(180)
			p.Value()
		}

	case 6:
		localctx = NewNegatedStringConditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(182)
			p.Match(LEQLParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(183)
			p.FieldList()
		}
		{
			p.SetState(184)
			p.StringOp()
		}
		{
			p.SetState(185)
			p.Value()
		}

	case 7:
		localctx = NewNegatedSetConditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(187)
			p.Match(LEQLParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(188)
			p.FieldList()
		}
		{
			p.SetState(189)
			p.SetOp()
		}
		{
			p.SetState(190)
			p.ValueList()
		}

	case 8:
		localctx = NewNegatedListStringConditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(192)
			p.Match(LEQLParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(193)
			p.FieldList()
		}
		{
			p.SetState(194)
			p.ListStringOp()
		}
		{
			p.SetState(195)
			p.ValueList()
		}

	case 9:
		localctx = NewPostfixNegatedComparisonConditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(197)
			p.FieldList()
		}
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
			p.ComparisonOp()
		}
		{
			p.SetState(200)
			p.Value()
		}

	case 10:
		localctx = NewPostfixNegatedStringConditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(202)
			p.FieldList()
		}
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
			p.StringOp()
		}
		{
			p.SetState(205)
			p.Value()
		}

	case 11:
		localctx = NewPostfixNegatedSetConditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(207)
			p.FieldList()
		}
		{
			p.SetState(208)
			p.Match(LEQLParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(209)
			p.SetOp()
		}
		{
			p.SetState(210)
			p.ValueList()
		}

	case 12:
		localctx = NewPostfixNegatedListStringConditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(212)
			p.FieldList()
		}
		{
			p.SetState(213)
			p.Match(LEQLParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(214)
			p.ListStringOp()
		}
		{
			p.SetState(215)
			p.ValueList()
		}

	case 13:
		localctx = NewNocaseConditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(217)
			p.FieldList()
		}
		{
			p.SetState(218)
			p.ComparisonOp()
		}
		{
			p.SetState(219)
			p.NocaseValue()
		}

	case 14:
		localctx = NewIpConditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(221)
			p.FieldList()
		}
		{
			p.SetState(222)
			p.ComparisonOp()
		}
		{
			p.SetState(223)
			p.IpValue()
		}

	case 15:
		localctx = NewIpSetConditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(225)
			p.FieldList()
		}
		{
			p.SetState(226)
			p.SetOp()
		}
		{
			p.SetState(227)
			p.ValueList()
		}

	case 16:
		localctx = NewAllFieldsComparisonConditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(229)
			p.AllFieldList()
		}
		{
			p.SetState(230)
			p.ComparisonOp()
		}
		{
			p.SetState(231)
			p.Value()
		}

	case 17:
		localctx = NewAllFieldsSetConditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(233)
			p.AllFieldList()
		}
		{
			p.SetState(234)
			p.SetOp()
		}
		{
			p.SetState(235)
			p.ValueList()
		}

	case 18:
		localctx = NewAllFieldsNocaseConditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(237)
			p.AllFieldList()
		}
		{
			p.SetState(238)
			p.ComparisonOp()
		}
		{
			p.SetState(239)
			p.NocaseValue()
		}

	case 19:
		localctx = NewFieldExistsConditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(241)
			p.Match(LEQLParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 20:
		localctx = NewNumberKeywordConditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(242)
			p.Match(LEQLParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 21:
		localctx = NewVariableConditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 21)
		{
			p.SetState(243)
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
	p.SetState(248)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LEQLParserDOUBLE_STRING:
		localctx = NewQuotedKeywordContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(246)
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
			p.SetState(247)
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
		p.SetState(250)
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
		p.SetState(252)
		p.FieldName()
	}
	p.SetState(257)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == LEQLParserCOMMA {
		{
			p.SetState(253)
			p.Match(LEQLParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(254)
			p.FieldName()
		}

		p.SetState(259)
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
		p.SetState(260)
		p.Match(LEQLParserALL_FUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(261)
		p.Match(LEQLParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(262)
		p.FieldName()
	}
	p.SetState(267)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == LEQLParserCOMMA {
		{
			p.SetState(263)
			p.Match(LEQLParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(264)
			p.FieldName()
		}

		p.SetState(269)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(270)
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
		p.SetState(272)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-9)) & ^0x3f) == 0 && ((int64(1)<<(_la-9))&5620492337022763011) != 0) {
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
		p.SetState(274)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4486007441326080) != 0) {
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
	p.EnterRule(localctx, 28, LEQLParserRULE_stringOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(276)
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
	p.EnterRule(localctx, 30, LEQLParserRULE_setOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(278)
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
	p.EnterRule(localctx, 32, LEQLParserRULE_listStringOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(280)
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
	p.EnterRule(localctx, 34, LEQLParserRULE_value)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(282)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-62)) & ^0x3f) == 0 && ((int64(1)<<(_la-62))&1023) != 0) {
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
	p.EnterRule(localctx, 36, LEQLParserRULE_nocaseValue)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(284)
		p.Match(LEQLParserNOCASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(285)
		p.Match(LEQLParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(286)
		p.Value()
	}
	{
		p.SetState(287)
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
	p.EnterRule(localctx, 38, LEQLParserRULE_ipValue)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(289)
		p.Match(LEQLParserIP_FUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(290)
		p.Match(LEQLParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(293)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(291)
			p.Match(LEQLParserIP_CIDR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		{
			p.SetState(292)
			p.Value()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(295)
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
	p.EnterRule(localctx, 40, LEQLParserRULE_valueList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(297)
		p.Match(LEQLParserLBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(298)
		p.ValueListItem()
	}
	p.SetState(303)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == LEQLParserCOMMA {
		{
			p.SetState(299)
			p.Match(LEQLParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(300)
			p.ValueListItem()
		}

		p.SetState(305)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(306)
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
	p.EnterRule(localctx, 42, LEQLParserRULE_valueListItem)
	p.SetState(310)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LEQLParserIP_FUNC:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(308)
			p.IpValue()
		}

	case LEQLParserIP_CIDR, LEQLParserREGEX, LEQLParserTRIPLE_SINGLE_STRING, LEQLParserTRIPLE_DOUBLE_STRING, LEQLParserDOUBLE_STRING, LEQLParserSINGLE_STRING, LEQLParserNUMBER, LEQLParserTIME_UNIT, LEQLParserVARIABLE, LEQLParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(309)
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
	p.EnterRule(localctx, 44, LEQLParserRULE_groupbyClause)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(312)
		p.Match(LEQLParserGROUPBY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(313)
		p.Match(LEQLParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(314)
		p.FieldName()
	}
	p.SetState(319)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == LEQLParserCOMMA {
		{
			p.SetState(315)
			p.Match(LEQLParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(316)
			p.FieldName()
		}

		p.SetState(321)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(322)
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
	p.EnterRule(localctx, 46, LEQLParserRULE_havingClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(324)
		p.Match(LEQLParserHAVING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(325)
		p.Match(LEQLParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(326)
		p.HavingCondition()
	}
	{
		p.SetState(327)
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
	p.EnterRule(localctx, 48, LEQLParserRULE_havingCondition)
	p.SetState(337)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(329)
			p.CalcFunction()
		}
		{
			p.SetState(330)
			p.ComparisonOp()
		}
		{
			p.SetState(331)
			p.Value()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(333)
			p.CalcFunctionWithField()
		}
		{
			p.SetState(334)
			p.ComparisonOp()
		}
		{
			p.SetState(335)
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
	p.EnterRule(localctx, 50, LEQLParserRULE_calculateClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(339)
		p.Match(LEQLParserCALCULATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(340)
		p.Match(LEQLParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(341)
		p.CalcExpr()
	}
	{
		p.SetState(342)
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
	p.EnterRule(localctx, 52, LEQLParserRULE_calcExpr)
	p.SetState(347)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(344)
			p.CalcFunction()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(345)
			p.CalcFunctionWithField()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(346)
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
	p.EnterRule(localctx, 54, LEQLParserRULE_calcFunction)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(349)
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
	p.EnterRule(localctx, 56, LEQLParserRULE_calcFunctionWithField)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(351)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1063541276672) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(352)
		p.Match(LEQLParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(353)
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
	p.EnterRule(localctx, 58, LEQLParserRULE_percentileFunction)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(355)
		_la = p.GetTokenStream().LA(1)

		if !(_la == LEQLParserPERCENTILE || _la == LEQLParserPCTL) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(356)
		p.Match(LEQLParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(357)
		p.Match(LEQLParserNUMBER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(358)
		p.Match(LEQLParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(361)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == LEQLParserCOLON {
		{
			p.SetState(359)
			p.Match(LEQLParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(360)
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
	p.EnterRule(localctx, 60, LEQLParserRULE_sortClause)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(363)
		p.Match(LEQLParserSORT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(364)
		p.Match(LEQLParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(365)
		p.SortSpec()
	}
	p.SetState(370)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == LEQLParserCOMMA {
		{
			p.SetState(366)
			p.Match(LEQLParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(367)
			p.SortSpec()
		}

		p.SetState(372)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(373)
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
	p.EnterRule(localctx, 62, LEQLParserRULE_sortSpec)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(375)
		p.SortDirection()
	}
	p.SetState(379)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	switch p.GetTokenStream().LA(1) {
	case LEQLParserHASH:
		{
			p.SetState(376)
			p.Match(LEQLParserHASH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(377)
			p.Match(LEQLParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case LEQLParserIDENTIFIER:
		{
			p.SetState(378)
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
	p.EnterRule(localctx, 64, LEQLParserRULE_sortDirection)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(381)
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
	p.EnterRule(localctx, 66, LEQLParserRULE_limitClause)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(383)
		p.Match(LEQLParserLIMIT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(384)
		p.Match(LEQLParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(385)
		p.Match(LEQLParserNUMBER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(390)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == LEQLParserCOMMA {
		{
			p.SetState(386)
			p.Match(LEQLParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(387)
			p.Match(LEQLParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(392)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(393)
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
	p.EnterRule(localctx, 68, LEQLParserRULE_timesliceClause)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(395)
		p.Match(LEQLParserTIMESLICE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(396)
		p.Match(LEQLParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(397)
		_la = p.GetTokenStream().LA(1)

		if !(_la == LEQLParserNUMBER || _la == LEQLParserTIME_UNIT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(398)
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
