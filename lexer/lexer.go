package lexer

import (
	"fmt"
	"strings"

	"github.com/komuw/khaled/token"
)

/*
The Lexer

It will take source code as input and output the tokens that rep the source code.
*/

type Lexer struct {
	input    string
	ch       byte // current char under examination
	position int  // position of ch
	// readPosition int  // current reading position in input (after current char)

}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	/*
		TODO: turn this to func
		eat whitespace, because it is of no use to khaled unlike python
	*/
	replacer := strings.NewReplacer(string(' '), "", string('\t'), "", string('\n'), "", string('\r'), "")
	l.input = replacer.Replace(l.input)

	/*
		TODO: tunr this to func
		the lexer only supports ASCII characters instead of full Unicode. This lets us keep things simple.
		To support Unicode/UTF-8 we would need to change l.ch from a byte to rune and change the way we read the
		next characters, since they could be multiple bytes wide now. Using l.input[l.readPosition]wouldn’t work anymore
	*/
	if l.position >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.position]
		l.position = l.position + 1
	}

	fmt.Println("l.ch", string(l.ch))

	switch l.ch {
	case '=':
		tok = token.NewToken(token.ASSIGN, l.ch)
	case ';':
		tok = token.NewToken(token.SEMICOLON, l.ch)
	case '(':
		tok = token.NewToken(token.LPAREN, l.ch)
	case ')':
		tok = token.NewToken(token.RPAREN, l.ch)
	case ',':
		tok = token.NewToken(token.COMMA, l.ch)
	case '+':
		tok = token.NewToken(token.PLUS, l.ch)
	case '{':
		tok = token.NewToken(token.LBRACE, l.ch)
	case '}':
		tok = token.NewToken(token.RBRACE, l.ch)
	case 0: // ASCII code for "NUL"
		tok.Value = ""
		tok.Type = token.EOF
	}
	return tok
}

// func (l *Lexer) readIdentifier() string {
// 	position := l.position
// 	for isLetter(l.ch) {
// 		l.readChar()
// 	}
// 	return l.input[position:l.position]
// }
// func isLetter(ch byte) bool {
// 	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
// }

// func (l *Lexer) readNumber() string {
// 	position := l.position
// 	for isDigit(l.ch) {
// 		l.readChar()
// 	}
// 	return l.input[position:l.position]
// }
// func isDigit(ch byte) bool {
// 	return '0' <= ch && ch <= '9'
// }
