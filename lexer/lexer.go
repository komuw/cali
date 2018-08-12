package lexer

import (
	"github.com/komuw/khaled/token"
)

/*
The Lexer

It will take source code as input and output the tokens that rep the source code.
*/

type Lexer struct {
	input        string
	ch           byte   // current char under examination
	position     int    // position of Ch
	readPosition int    //  next reading position in input (after current char)
	ByteStream   []byte // value of Input as bytes

}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	for index := 0; index < len(l.input); index++ {
		l.ByteStream = append(l.ByteStream, l.input[index])

	}

	/*
		use readChar, so our *Lexer is in a fully working state b4 anyone calls NextToken()
	*/
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace()

	/*
		We look at the current character under
		examination (l.ch) and return a token depending on which character it is.
	*/
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
	case '-':
		tok = token.NewToken(token.MINUS, l.ch)
	case '!':
		tok = token.NewToken(token.BANG, l.ch)
	case '/':
		tok = token.NewToken(token.SLASH, l.ch)
	case '*':
		tok = token.NewToken(token.ASTERISK, l.ch)
	case '<':
		tok = token.NewToken(token.LT, l.ch)
	case '>':
		tok = token.NewToken(token.GT, l.ch)
	case 0: // ASCII code for "NUL"
		tok.Value = ""
		tok.Type = token.EOF
	default:
		/*
			Our lexer needs to  recognize whether the current character is a letter and if so,
			it needs to read the rest of the identifier/keyword until it encounters a non-letter-character.
		*/
		if isLetter(l.ch) {
			/*
			 if we encounter a letter; read until we encounter a non letter(ie read whole word)
			 eg if we encounter- let - then we need to read whole of it(let) as one token
			*/
			tok.Value = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Value)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Value = l.readNumber()
			return tok
		} else {
			tok = token.NewToken(token.ILLEGAL, l.ch)
		}

	}

	/*
		Before returning the token we advance our pointers into the input
		so when we call NextToken() again the l.ch field is already updated
	*/
	l.readChar()
	return tok
}

/* readChar gives us the next character and advance our position in the input string.
the lexer only supports ASCII characters instead of full Unicode. This lets us keep things simple.
To support Unicode/UTF-8 we would need to change l.ch from a byte to rune and change the way we read the
next characters, since they could be multiple bytes wide now. Using l.input[l.readPosition]wouldn’t work anymore
*/
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	/*
		In khaled we treat _ as a letter and allow it in identifiers and keywords.
		That means we can have var names like foo_bar. If you want to allow other chars in ua lang; add them here.
	*/
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

/*
eat whitespace. khaled doesnt require it unlike python.
This func is found in a lot of parsers. Sometimes it’s called eatWhitespace/consumeWhitespace.
Which chars these functions actually skip depends on the language being lexed.
*/
func (l *Lexer) skipWhitespace() {

	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

/*
We only read intergers. What about floats,hex notation, Octal notation?
We ignore em and say khaled doesn't support them.
*/
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
