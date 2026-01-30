parser grammar LEQLParser;

options { tokenVocab = LEQLLexer; }

// Entry point: LEQL query is a sequence of clauses in a defined order
// All clauses are optional; a bare keyword/regex search has no clauses
query
    : selectClause? whereClause? groupbyClause? havingClause? calculateClause? sortClause? limitClause? timesliceClause? EOF
    ;

// SELECT clause: select(key1 as alias1, key2 as alias2, ...)
selectClause
    : SELECT LPAREN selectFieldList RPAREN
    ;

selectFieldList
    : selectField (COMMA selectField)*
    ;

selectField
    : fieldName (IDENTIFIER IDENTIFIER)?
    ;

// WHERE clause: where(condition)
whereClause
    : WHERE LPAREN expression RPAREN
    ;

// Boolean expressions with AND/OR/NOT and parenthesized grouping
expression
    : LPAREN expression RPAREN                          # parenExpr
    | NOT expression                                    # notExpr
    | expression AND expression                         # andExpr
    | expression OR expression                          # orExpr
    | condition                                         # conditionExpr
    | keywordSearch                                     # keywordExpr
    | regexSearch                                       # regexExpr
    ;

// Key-value pair conditions
condition
    : fieldList comparisonOp value                      # comparisonCondition
    | fieldList stringOp value                          # stringCondition
    | fieldList setOp valueList                         # setCondition
    | fieldList listStringOp valueList                  # listStringCondition
    | NOT fieldList comparisonOp value                  # negatedComparisonCondition
    | NOT fieldList stringOp value                      # negatedStringCondition
    | NOT fieldList setOp valueList                     # negatedSetCondition
    | NOT fieldList listStringOp valueList              # negatedListStringCondition
    | fieldList comparisonOp nocaseValue                # nocaseCondition
    | fieldList comparisonOp ipValue                    # ipCondition
    | fieldList setOp valueList                          # ipSetCondition
    | allFieldList comparisonOp value                   # allFieldsComparisonCondition
    | allFieldList setOp valueList                      # allFieldsSetCondition
    | fieldName                                         # fieldExistsCondition
    ;

// Keyword search: bare string or quoted phrase
keywordSearch
    : DOUBLE_STRING                                     # quotedKeyword
    | SINGLE_STRING                                     # singleQuotedKeyword
    ;

// Regex search: /pattern/ or /pattern/flags
regexSearch
    : REGEX                                             # regexPattern
    ;

// Field names (single or compound with comma = OR)
fieldList
    : fieldName (COMMA fieldName)*
    ;

// ALL(field1, field2) = AND logic for compound keys
allFieldList
    : ALL_FUNC LPAREN fieldName (COMMA fieldName)* RPAREN
    ;

// Field name: simple or dotted path (handled by IDENTIFIER token which includes dots)
fieldName
    : IDENTIFIER
    ;

// Comparison operators
comparisonOp
    : EQ | NEQ | STRICT_EQ | STRICT_NEQ | GT | GTE | LT | LTE
    ;

// String operators
stringOp
    : CONTAINS_OP | ICONTAINS_OP | STARTS_WITH_OP | ISTARTS_WITH_OP
    ;

// Set operators
setOp
    : IN_OP | IIN_OP
    ;

// List-based string operators
listStringOp
    : CONTAINS_ANY_OP | ICONTAINS_ANY_OP | CONTAINS_ALL_OP | ICONTAINS_ALL_OP
    | STARTS_WITH_ANY_OP | ISTARTS_WITH_ANY_OP
    ;

// Values
value
    : DOUBLE_STRING
    | SINGLE_STRING
    | TRIPLE_SINGLE_STRING
    | TRIPLE_DOUBLE_STRING
    | NUMBER
    | TIME_UNIT
    | REGEX
    | IDENTIFIER
    ;

// NOCASE(value) wrapper
nocaseValue
    : NOCASE LPAREN value RPAREN
    ;

// IP(cidr) wrapper
ipValue
    : IP_FUNC LPAREN value RPAREN
    ;

// Value list: [v1, v2, ...]
valueList
    : LBRACKET valueListItem (COMMA valueListItem)* RBRACKET
    ;

valueListItem
    : ipValue
    | value
    ;

// GROUPBY clause: groupby(field1, field2, ...)
groupbyClause
    : GROUPBY LPAREN fieldName (COMMA fieldName)* RPAREN
    ;

// HAVING clause: having(calcFunc operator value)
havingClause
    : HAVING LPAREN havingCondition RPAREN
    ;

havingCondition
    : calcFunction comparisonOp value
    | calcFunctionWithField comparisonOp value
    ;

// CALCULATE clause: calculate(function) or calculate(function:field)
calculateClause
    : CALCULATE LPAREN calcExpr RPAREN
    ;

calcExpr
    : calcFunction
    | calcFunctionWithField
    | percentileFunction
    ;

calcFunction
    : COUNT | BYTES
    ;

calcFunctionWithField
    : ( SUM | AVERAGE | UNIQUE | MIN | MAX | STANDARDDEVIATION | SD ) COLON fieldName
    ;

percentileFunction
    : ( PERCENTILE | PCTL ) LPAREN NUMBER RPAREN (COLON fieldName)?
    ;

// SORT clause: sort(asc) or sort(desc)
sortClause
    : SORT LPAREN sortDirection RPAREN
    ;

sortDirection
    : ASC | ASCENDING | DESC | DESCENDING
    ;

// LIMIT clause: limit(n) or limit(n1, n2)
limitClause
    : LIMIT LPAREN NUMBER (COMMA NUMBER)* RPAREN
    ;

// TIMESLICE clause: timeslice(n) or timeslice(30m)
timesliceClause
    : TIMESLICE LPAREN ( NUMBER | TIME_UNIT ) RPAREN
    ;
