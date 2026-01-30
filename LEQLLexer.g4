lexer grammar LEQLLexer;

// Clauses (case-insensitive)
SELECT      : [Ss][Ee][Ll][Ee][Cc][Tt] ;
WHERE       : [Ww][Hh][Ee][Rr][Ee] ;
GROUPBY     : [Gg][Rr][Oo][Uu][Pp][Bb][Yy] | [Gg][Rr][Oo][Uu][Pp] '_' [Bb][Yy] ;
HAVING      : [Hh][Aa][Vv][Ii][Nn][Gg] ;
CALCULATE   : [Cc][Aa][Ll][Cc][Uu][Ll][Aa][Tt][Ee] ;
SORT        : [Ss][Oo][Rr][Tt] ;
LIMIT       : [Ll][Ii][Mm][Ii][Tt] ;
TIMESLICE   : [Tt][Ii][Mm][Ee][Ss][Ll][Ii][Cc][Ee] ;

// Logical operators (case-insensitive)
AND         : [Aa][Nn][Dd] ;
OR          : [Oo][Rr] ;
NOT         : [Nn][Oo][Tt] ;

// String comparison operators (case-insensitive, with hyphens)
// Order matters: longer tokens must come before shorter ones to avoid ambiguity
ICONTAINS_ALL_OP      : [Ii][Cc][Oo][Nn][Tt][Aa][Ii][Nn][Ss] '-' [Aa][Ll][Ll] ;
ICONTAINS_ANY_OP      : [Ii][Cc][Oo][Nn][Tt][Aa][Ii][Nn][Ss] '-' [Aa][Nn][Yy] ;
CONTAINS_ALL_OP       : [Cc][Oo][Nn][Tt][Aa][Ii][Nn][Ss] '-' [Aa][Ll][Ll] ;
CONTAINS_ANY_OP       : [Cc][Oo][Nn][Tt][Aa][Ii][Nn][Ss] '-' [Aa][Nn][Yy] ;
ISTARTS_WITH_ANY_OP   : [Ii][Ss][Tt][Aa][Rr][Tt][Ss] '-' [Ww][Ii][Tt][Hh] '-' [Aa][Nn][Yy] ;
STARTS_WITH_ANY_OP    : [Ss][Tt][Aa][Rr][Tt][Ss] '-' [Ww][Ii][Tt][Hh] '-' [Aa][Nn][Yy] ;
ISTARTS_WITH_OP       : [Ii][Ss][Tt][Aa][Rr][Tt][Ss] '-' [Ww][Ii][Tt][Hh] ;
STARTS_WITH_OP        : [Ss][Tt][Aa][Rr][Tt][Ss] '-' [Ww][Ii][Tt][Hh] ;
ICONTAINS_OP          : [Ii][Cc][Oo][Nn][Tt][Aa][Ii][Nn][Ss] ;
CONTAINS_OP           : [Cc][Oo][Nn][Tt][Aa][Ii][Nn][Ss] ;

// Set operators (case-insensitive) — IIN before IN to avoid prefix match
IIN_OP      : [Ii][Ii][Nn] ;
IN_OP       : [Ii][Nn] ;

// Special functions (case-insensitive)
NOCASE      : [Nn][Oo][Cc][Aa][Ss][Ee] ;
IP_FUNC     : [Ii][Pp] ;
ALL_FUNC    : [Aa][Ll][Ll] ;

// Calculate functions (case-insensitive)
// Longer tokens first to avoid prefix matching
STANDARDDEVIATION : [Ss][Tt][Aa][Nn][Dd][Aa][Rr][Dd][Dd][Ee][Vv][Ii][Aa][Tt][Ii][Oo][Nn] ;
PERCENTILE  : [Pp][Ee][Rr][Cc][Ee][Nn][Tt][Ii][Ll][Ee] ;
AVERAGE     : [Aa][Vv][Ee][Rr][Aa][Gg][Ee] ;
UNIQUE      : [Uu][Nn][Ii][Qq][Uu][Ee] ;
COUNT       : [Cc][Oo][Uu][Nn][Tt] ;
BYTES       : [Bb][Yy][Tt][Ee][Ss] ;
PCTL        : [Pp][Cc][Tt][Ll] ;
SUM         : [Ss][Uu][Mm] ;
MIN         : [Mm][Ii][Nn] ;
MAX         : [Mm][Aa][Xx] ;
SD          : [Ss][Dd] ;

// Sort directions (case-insensitive) — longer first
ASCENDING   : [Aa][Ss][Cc][Ee][Nn][Dd][Ii][Nn][Gg] ;
DESCENDING  : [Dd][Ee][Ss][Cc][Ee][Nn][Dd][Ii][Nn][Gg] ;
ASC         : [Aa][Ss][Cc] ;
DESC        : [Dd][Ee][Ss][Cc] ;

// Multi-char comparison operators (before single-char to avoid prefix issues)
STRICT_NEQ  : '!==' ;
NEQ         : '!=' ;
GTE         : '>=' ;
LTE         : '<=' ;
STRICT_EQ   : '==' ;
EQ          : '=' ;
GT          : '>' ;
LT          : '<' ;
BANG        : '!' ;

// Delimiters
LPAREN      : '(' ;
RPAREN      : ')' ;
LBRACKET    : '[' ;
RBRACKET    : ']' ;
COMMA       : ',' ;
COLON       : ':' ;
DOT         : '.' ;
STAR        : '*' ;
HASH        : '#' ;

// Regex literal: /pattern/ or /pattern/i
REGEX       : '/' REGEX_BODY '/' REGEX_FLAGS? ;
fragment REGEX_BODY : ( ~[/\\] | '\\' . )+ ;
fragment REGEX_FLAGS : [imsxU]+ ;

// String literals
TRIPLE_SINGLE_STRING  : '\'\'\'' .*? '\'\'\'' ;
TRIPLE_DOUBLE_STRING  : '"""' .*? '"""' ;
DOUBLE_STRING         : '"' ( ~["\\\r\n] | '\\' . )* '"' ;
SINGLE_STRING         : '\'' ( ~['\\\r\n] | '\\' . )* '\'' ;

// Numbers (integer, float, scientific notation)
NUMBER      : '-'? DIGIT+ ('.' DIGIT+)? ([eE] [+-]? DIGIT+)? ;
fragment DIGIT : [0-9] ;

// Time units for timeslice (e.g., 30m, 1h, 7d)
TIME_UNIT   : DIGIT+ [smhd] ;

// Identifiers: field names, including dotted paths (a.b.c) and wildcard (a.*)
// Also allows hyphens in field names and # prefix for log set references
IDENTIFIER  : (LETTER | '_' | '#') (LETTER | DIGIT | '_' | '-')* ('.' (LETTER | DIGIT | '_' | '-' | '*')+)* ;
fragment LETTER : [a-zA-Z] ;

// Whitespace
WS          : [ \t\r\n]+ -> skip ;
