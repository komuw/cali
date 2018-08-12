package token

/*
LEXING

1. Lexical analysis


                   lexing/tokenization/scanning                  parsing
                   (lexer/tokenizer/scannner)                    (parser)
Source-code       ------------------------------>   Tokens   ---------------> AST
"let x = 5 + 5;"                                    [
                                                    LET,
                                                    IDENTIFIER("x"),
                                                    EQUAL_SIGN,
                                                    INTEGER(5),
                                                    PLUS_SIGN,
                                                    INTEGER(5),
                                                    SEMICOLON
                                                    ]

what exactly constitutes a "token" varies between different lexer implementations.
A production-ready lexer might also attach the line number, column number and filename to
a token. A reason may be so that we can output useful error messages in the parsing stage:
  "error: expected semicolon token. line 42, column 23, program.khaled"

2. Define tokens
	let five = 5; //1. numbers like 5
	let add = fn(x, y) { //2. keywords like fn, let
	x + y; //3. vars like x, y, add, result
	};
	let result = add(five, ten); //4. special chars like (, ), {, }, ;

*/

// TokenType is a string. But using an int or byte would have better performance
type TokenType string

/*
Token needs
1. “type” attribute, so we can distinguish between “integers” and “right bracket” for example.
2. field that holds the literal value of the token, so we can reuse it later and the info whether a “number” token is a 5 or a 10 doesn’t get lost.
*/
type Token struct {
	Type  TokenType
	Value string
}

// NewToken creates new token
func NewToken(tokenType TokenType, ch byte) Token {
	return Token{Type: tokenType, Value: string(ch)}
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"
	EQ       = "==" // EQ and NOT_EQ require the use of lookAhead
	NOT_EQ   = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
