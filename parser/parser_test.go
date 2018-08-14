package parser

import (
	"testing"

	"github.com/komuw/khaled/ast"
	"github.com/komuw/khaled/lexer"
)

/*
We are providing source code as input instead of tokens,
since that makes the tests much more readable and understandable.
Ofcourse this means that bugs in the lexer would lead to failures of parser tests; coupling.
*/
func TestLetStatementsSimple(t *testing.T) {
	input := `
		let x = 5;
		`
	l := lexer.NewLexer(input)
	p := NewParser(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 1 {
		t.Fatalf("\n No of program.Statements. \ngot %#+v \nwanted %#+v", len(program.Statements), 1)
	}
	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
	}
	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}
func TestLetStatementsExtended(t *testing.T) {
	input := `
		let x = 5;
		let y = 10;
		let foobar = 838383;
		`
	l := lexer.NewLexer(input)
	p := NewParser(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("\n No of program.Statements. \ngot %#+v \nwanted %#+v", len(program.Statements), 3)
	}
	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}
	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func TestLetStatementsParseError(t *testing.T) {
	input := `
		let = 5;
		`
	l := lexer.NewLexer(input)
	p := NewParser(l)
	_ = p.ParseProgram()

	errors := p.Errors()
	if len(errors) == 0 {
		t.Fatalf("\n No errors found. \ngot %#+v \nwanted %#+v", len(errors), 1)
	}
	t.Logf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Logf("parser error: %q", msg)
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	/*
		We get an AST(s argument) like;
			&ast.LetStatement{
				Token: token.Token{Type: "LET", Value: "let"},
				Name: &ast.Identifier{
					Token: token.Token{Type: "IDENT", Value: "x"},
					Value: "x"},
				Value: nil,
				}

		In this test helper func we are not checking s.Value
		This is because, that value is currently nil(instead of 5 for input like let x = 5;)
		It is nil because the only thing our parser currently is doing is only parsing let statements.
		It doesnt yet know how to parse expressions(the thing after the equal sign in let x = 5;)
		see documentation for parseLetStatement func.
	*/
	if s.TokenValue() != "let" {
		t.Errorf("s.TokenValue not 'let'. got=%q", s.TokenValue())
		return false
	}
	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}
	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.Value)
		return false
	}
	if letStmt.Name.TokenValue() != name {
		t.Errorf("letStmt.Name.TokenValue() not '%s'. got=%s",
			name, letStmt.Name.TokenValue())
		return false
	}
	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}
