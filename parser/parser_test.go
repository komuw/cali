package parser

import (
	"testing"

	"github.com/komuw/cali/ast"
	"github.com/komuw/cali/lexer"
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
		t.Errorf("letStmt.Name.TokenValue() not '%s'. got=%s", name, letStmt.Name.TokenValue())
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

func TestReturnStatements(t *testing.T) {
	input := `
	return 5;
	return 10;
	return 993322;
	`
	l := lexer.NewLexer(input)
	p := NewParser(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}
	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.returnStatement. got=%T", stmt)
			continue
		}
		if returnStmt.TokenValue() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q", returnStmt.TokenValue())
		}
	}
}

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"

	l := lexer.NewLexer(input)
	p := NewParser(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	// check that the *ast.ExpressionStatement.Expression is an *ast.Identifier
	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("exp not *ast.Identifier. got=%T", stmt.Expression)
	}
	if ident.Value != "foobar" {
		t.Errorf("ident.Value not %s. got=%s", "foobar", ident.Value)
	}
	if ident.TokenValue() != "foobar" {
		t.Errorf("ident.TokenLiteral not %s. got=%s", "foobar", ident.TokenValue())
	}
}

func TestIdentifierExpressionParseError(t *testing.T) {
	input := "foobar"
	l := lexer.NewLexer(input)
	p := NewParser(l)
	_ = p.ParseProgram()

	errors := p.Errors()
	if len(errors) == 0 {
		t.Fatalf("\n No errors found. \ngot %#+v \nwanted %#+v", len(errors), 1)
	}
	if len(errors) != 1 {
		t.Fatalf("\n number of errors mismatch. \ngot %#+v \nwanted %#+v", len(errors), 1)
	}
	t.Logf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Logf("parser error: %q", msg)
	}
}

func TestIntegerLiteralExpression(t *testing.T) {
	/*
		other examples are;
			let x = 5;
			add(5, 10);
			5 + 5 + 5;
	*/
	input := "5;"
	l := lexer.NewLexer(input)
	p := NewParser(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d",
			len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
			program.Statements[0])
	}
	literal, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("exp not *ast.IntegerLiteral. got=%T", stmt.Expression)
	}
	if literal.Value != 5 {
		t.Errorf("literal.Value not %d. got=%d", 5, literal.Value)
	}
	if literal.TokenValue() != "5" {
		t.Errorf("literal.TokenLiteral not %s. got=%s", "5",
			literal.TokenValue())
	}
}
