grammar DCell;

/*****************************************************************************
  Parser rules
******************************************************************************/

program
  : expression EOF
  ;

expression
  : term                                                # termExpression
  | expression '.' invocation                           # invocationExpression
  | expression '[' index ']'                            # indexExpression
  | expression ('is' 'not' | 'is') type                 # isExpression
  | expression ('in' | 'not' 'in') expression           # containsExpression
  | '(' expression ')'                                  # parenthesisExpression
  | ('!' | 'not') expression                            # logicalNotExpression
  | '~' expression                                      # bitwiseNotExpression
  | ('+' | '-') expression                              # polarityExpression
  | expression '**' expression                          # exponentiationExpression
  | expression ('*' | '/' | '//' | '%') expression      # multiplicativeExpression
  | expression ('+' | '-') expression                   # additiveExpression
  | expression ('&&' | 'and') expression                # logicalAndExpression
  | expression ('||' | 'or') expression                 # logicalOrExpression
  | expression ('<->' | 'implies') expression           # implicationExpression
  | expression ('<<' | '>>') expression                 # shiftExpression
  | expression '&' expression                           # bitwiseAndExpression
  | expression ('^' | '|') expression                   # bitwiseOrExpression
  | expression ('<=' | '<' | '>' | '>=') expression     # inequalityExpression
  | expression ('==' | '!=') expression                 # equalityExpression
  | expression '?' expression ':' expression            # ternaryExpression
  | expression '?:' expression                          # elvisExpression
  | expression '??' expression                          # coalesceExpression
  | expression 'as' type                                # castExpression
  ;

term
  : literal                                            # literalTerm
  | invocation                                         # invocationTerm
  ;

invocation
  : identifier                                         # memberInvocation
  | '*'                                                # wildcardInvocation
  | identifier '(' parameterList? ')'                  # functionInvocation
  ;

parameterList
  : expression (',' expression)*
  ;

identifier
  : IDENTIFIER
  ;

index
  : (expression)? ':' (expression)?                    # sliceIndex
  | expression                                         # expressionIndex
  ;

literal
  : string                                             # stringLiteral
  | integer                                            # integerLiteral
  | float                                              # floatLiteral
  | ('true' | 'false')                                 # booleanLiteral
  | 'null'                                             # nullLiteral
  | list                                               # listLiteral
  ;

type
  : ('int' | 'float' | 'string' | 'bool')
  ;

list
  : '[' (literal (',' literal)*)? ']'
  ;

string
  : SINGLE_QUOTE_STRING                                # singleQuoteString
  | DOUBLE_QUOTE_STRING                                # doubleQuoteString
  | TRIPLE_QUOTE_STRING                                # tripleQuoteString
  ;

integer
  : DECIMAL_INTEGER                                    # decimalInteger
  | HEX_INTEGER                                        # hexInteger
  | OCTAL_INTEGER                                      # octalInteger
  | BINARY_INTEGER                                     # binaryInteger
  ;

float
  : SCIENTIFIC_FLOAT                                   # scientificFloat
  | DECIMAL_FLOAT                                      # decimalFloat
  ;

/*****************************************************************************
  Lexer rules
******************************************************************************/

IDENTIFIER       : [a-zA-Z_][a-zA-Z0-9_-]*([a-zA-Z0-9_])?;
DECIMAL_INTEGER  : ('-'? [1-9][0-9]* | '0') ;
HEX_INTEGER      : '0' [xX] [0-9a-fA-F]+ ;
OCTAL_INTEGER    : '0' [0-7]+ ;
BINARY_INTEGER   : '0' [bB] [01]+ ;
DECIMAL_FLOAT    : '-'? ('0' | [1-9][0-9]*) '.' [0-9]+ ;
SCIENTIFIC_FLOAT : '-'? ('0' | [1-9][0-9]*) ('.' [0-9]+)? [eE] [+-]? [1-9][0-9]* ;
SINGLE_QUOTE_STRING : '\'' (ESC | ~['\\\r\n])* '\'' ;
DOUBLE_QUOTE_STRING : '"' (ESC | ~["\\\r\n])* '"' ;
TRIPLE_QUOTE_STRING : '"""' (ESC | .)*? '"""' ;

// Pipe whitespace to the HIDDEN channel to support retrieving source text through the parser.
WS             : [ \t\r\n]+ -> channel(HIDDEN) ;
COMMENT        : '#' ~[\r\n]* -> channel(HIDDEN) ;

// Fragments

fragment ESC
  : '\\' ([`'\\/fnrt] | UNICODE)
  ;

fragment UNICODE
  : 'u' HEX HEX HEX HEX
  ;

fragment HEX
  : [0-9a-fA-F]
  ;
