package lexer

import (
	"fmt"

	"github.com/sanity-io/litter"

	"github.com/komuw/khaled/token"
)

/*
The Lexer

It will take source code as input and output the tokens that rep the source code.
*/

type Lexer struct {
	Input      string
	Ch         byte   // current char under examination
	Position   int    // position of ch
	ByteStream []byte // value of Input as bytes
	// readPosition int  // current reading position in input (after current char)

}

func NewLexer(input string) *Lexer {
	l := &Lexer{Input: input}
	for index := 0; index < len(l.Input); index++ {
		l.ByteStream = append(l.ByteStream, l.Input[index])

	}
	return l
}

// skipWhiteSpace eats whitespace, because it is of no use to khaled unlike python
func (l *Lexer) skipWhiteSpace() {
	if l.Ch == ' ' || l.Ch == '\t' || l.Ch == '\n' || l.Ch == '\r' {
		// skip that char
		l.Position = l.Position + 1
		l.Ch = l.Input[l.Position]
	}
}

func (l *Lexer) NextToken() token.Token {
	litter.Dump("l1", l)
	var tok token.Token
	l.skipWhiteSpace()

	/*
		TODO: tunr this to func
		the lexer only supports ASCII characters instead of full Unicode. This lets us keep things simple.
		To support Unicode/UTF-8 we would need to change l.Ch from a byte to rune and change the way we read the
		next characters, since they could be multiple bytes wide now. Using l.Input[l.readPosition]wouldn’t work anymore
	*/
	if l.Position >= len(l.Input) {
		l.Ch = 0
	}

	litter.Dump("l2", l)
	fmt.Println("l.Ch", string(l.Ch))

	switch l.Ch {
	case '=':
		tok = token.NewToken(token.ASSIGN, l.Ch)
	case ';':
		tok = token.NewToken(token.SEMICOLON, l.Ch)
	case '(':
		tok = token.NewToken(token.LPAREN, l.Ch)
	case ')':
		tok = token.NewToken(token.RPAREN, l.Ch)
	case ',':
		tok = token.NewToken(token.COMMA, l.Ch)
	case '+':
		tok = token.NewToken(token.PLUS, l.Ch)
	case '{':
		tok = token.NewToken(token.LBRACE, l.Ch)
	case '}':
		tok = token.NewToken(token.RBRACE, l.Ch)
	case 0: // ASCII code for "NUL"
		tok.Value = ""
		tok.Type = token.EOF
		// default:
		// 	if isLetter(l.Ch) {
		// 		/*
		// 		 if we encounter a letter; read until we encounter a non letter(ie read whole word)
		// 		 eg if we encounter- let - then we need to read whole of it(let) as one token
		// 		*/
		// 		tok.Value = l.readIdentifier()
		// 		tok.Type = token.LookupIdent(tok.Value)
		// 		return tok
		// 	} else {
		// 		tok = token.NewToken(token.ILLEGAL, l.Ch)
		// 	}
	}

	// advance
	l.Position = l.Position + 1
	l.Ch = l.Input[l.Position]

	litter.Dump("l3", l)
	return tok
}

func (l *Lexer) readIdentifier() string {
	wordStartPos := l.Position
	for isLetter(l.Ch) {
		l.Ch = l.Input[l.Position]
		l.Position = l.Position + 1
	}
	fmt.Println("l.Input[wordStartPos:l.Position]", string(l.Input[wordStartPos:l.Position]))
	return l.Input[wordStartPos:l.Position]
}
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// func (l *Lexer) readNumber() string {
// 	position := l.Position
// 	for isDigit(l.Ch) {
// 		l.readChar()
// 	}
// 	return l.Input[position:l.Position]
// }
// func isDigit(ch byte) bool {
// 	return '0' <= ch && ch <= '9'
// }
