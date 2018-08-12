package lexer

import (
	"testing"

	"github.com/komuw/khaled/token"
)

func TestNextTokenSimple(t *testing.T) {
	input := `=+(){},;`
	tests := []struct {
		expectedType  token.TokenType
		expectedValue string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}
	l := NewLexer(input)

	for _, v := range tests {
		tok := l.NextToken()
		if tok.Type != v.expectedType {
			t.Fatalf("\n Tokentype wrong. \nCalled l.NextToken() \ngot %#+v \nwanted %#+v", tok.Type, v.expectedType)
		}
		if tok.Value != v.expectedValue {
			t.Fatalf("\n Value wrong. \nCalled l.NextToken() \ngot %#+v \nwanted %#+v", tok.Value, v.expectedValue)
		}
	}
}

func TestNextTokenExtened(t *testing.T) {
	input := `let five = 5;
	let ten = 10;
	let add = fn(x, y) {
	x + y;
	};
	let result = add(five, ten);
	`
	tests := []struct {
		expectedType  token.TokenType
		expectedValue string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}
	l := NewLexer(input)

	for _, v := range tests {
		tok := l.NextToken()
		if tok.Type != v.expectedType {
			t.Fatalf("\n Tokentype wrong. \nCalled l.NextToken():%#+v  \ngot %#+v \nwanted %#+v", string(l.ch), tok.Type, v.expectedType)
		}
		if tok.Value != v.expectedValue {
			t.Fatalf("\n Value wrong. \nCalled l.NextToken():%#+v \ngot %#+v \nwanted %#+v", string(l.ch), tok.Value, v.expectedValue)
		}
	}
}

func TestNextTokenPage21(t *testing.T) {
	/*
		the input looks like khaled source code, with gibberish like !-/*5.
		That’s okay. The lexer’s job is not to tell us whether code makes sense.
		The lexer should only turn this input into tokens.
	*/
	input := `
	!-/*5;
	5 < 10 > 5;

	if (5 < 10) {
	return true;
	} else {
	return false;
	}

	10 == 10;
	10 != 9;
	`
	tests := []struct {
		expectedType  token.TokenType
		expectedValue string
	}{
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},

		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},

		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.NOT_EQ, "!="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},

		{token.EOF, ""},
	}
	l := NewLexer(input)

	for _, v := range tests {
		tok := l.NextToken()
		if tok.Type != v.expectedType {
			t.Fatalf("\n Tokentype wrong. \nCalled l.NextToken() \ngot type:%#+v of value:%#+v \nwanted %#+v", tok.Type, tok.Value, v.expectedType)
		}
		if tok.Value != v.expectedValue {
			t.Fatalf("\n Value wrong. \nCalled l.NextToken() \ngot %#+v \nwanted %#+v", tok.Value, v.expectedValue)
		}
	}
}
