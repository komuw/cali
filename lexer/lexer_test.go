package lexer

import (
	"testing"

	"github.com/komuw/khaled/token"
)

func TestNextToken(t *testing.T) {
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
