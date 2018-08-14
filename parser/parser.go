package parser

import (
	"fmt"

	"github.com/komuw/khaled/ast"
	"github.com/komuw/khaled/lexer"
	"github.com/komuw/khaled/token"
)

/*
curToken & peekToken do same job like the position and readPosition fields we had in the lexer.
They point to the current and the next token.
Think of a single line only containing 5;. Then curToken is a token.INT and we need peekToken to decide whether
we are at the end of the line or if we are at just the start of an arithmetic expression
*/
type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
	errors    []string
}

func NewParser(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}
	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()
	return p
}
func (p *Parser) Errors() []string {
	return p.errors
}
func (p *Parser) peekError(t token.TokenType) {
	// TODO: add line numbers to this errors
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

/*
ParseProgram :
1. constructs the root node of the AST, an *ast.Program
2. It iterates over every token in input until an token.EOF by repeatedly calling nextToken()
   nextToken() advances both p.curToken and p.peekToken
3. In every iteration it calls parseStatement, whose job it is to parse a statement.
If parseStatement returned something(not nil), its return value is added to Statements slice
4. When nothing is left to parse the *ast.Program root node is returned.
*/
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	/*
		We are currently only parsing LET statements.
		We are for example not parsing value like int(5), we'll do that later.
	*/
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

/*
parseLetStatement:
1. constructs an *ast.LetStatement node with the current token(token.LET)
2. advances the tokens while making assertions about the next token with calls to expectPeek.
  - First it expects a token.IDENT, which it then uses to construct an *ast.Identifier node
  - Then it expects an equal sign and finally it SKIPS over the expression following the equal sign until it encounters a semicolon.
  The skipping of expressions will be replaced, of course, as soon as we know how to parse expressions.
*/
func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}
	if !p.expectPeek(token.IDENT) {
		// if we get an UNexpected next token, we parseError, we do that in expectPeek()
		return nil
	}
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Value}
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}
	// TODO: SKIP parsing expressions until we know how to parse them.
	// Continue until we encounter a semicolon
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

/*
expectPeek() method is one of the "assertion funcs" that all parsers have.
They enforce the correctness of the order of tokens by checking the type of the next token.
Only if the type is correct does it advance the tokens by calling nextToken.
*/
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}
