package parser

import (
	"fmt"
	"strconv"

	"github.com/komuw/khaled/ast"
	"github.com/komuw/khaled/lexer"
	"github.com/komuw/khaled/token"
)

/*
constants showing operator precendence of khaled language.

These constants let us answer:
does the * operator have a higher precedence than the == operator? etc
OpLowest is int 0, OpEqualsEquals is int 1 etc
*/
const (
	OpLowest       int = iota
	OpEqualsEquals     // ==
	OpLessGreater      // > or <
	OpPlus             // +
	OpMultiplier       // *
	OpPrefix           // -X or !X
	OpCall             // myFunction(X)
)

/*
A Pratt parser’s main idea is the association of parsing funcs with token types.
Each token type can have up to two parsing funcs associated with it, depending on whether the
token is found in a prefix or an infix position.

the infixParseFn takes an argument: another ast.Expression.
This arg is the left-side of the infix operator that’s being parsed.

All of our parsing funcs, prefixParseFn/infixParseFn, are going to follow this protocol:
start with curToken being the type of token you’re associated with and return with curToken being the last token that’s part of your
expression type.
Never advance the tokens too far.
*/
type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
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

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

func NewParser(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}
	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	/*
		create prefixParseFns map and register funcs for various tokens.
		if we encounter a type token.IDENT the parsing function to call is p.parseIdentifier
	*/
	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.prefixParseFns[token.IDENT] = p.parseIdentifier // equivalent to p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.prefixParseFns[token.INT] = p.parseIntegerLiteral
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

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}
func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
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
		Since the only two real statement types in khaled are let and return statements
		create a case for them, else parseExpression.
	*/
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
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

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}
	p.nextToken()

	// TODO: SKIP parsing expressions until we know how to parse them.
	// Continue until we encounter a semicolon
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

/*
parseExpressionStatement parses expressions
*/
func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}
	stmt.Expression = p.parseExpression(OpLowest)

	/*
		Unlike in the interpreter book, in khaled we stop on finding semicolon.
		In khaled we wont allow code like; 5+5 (it has to be 5+5;)
	*/
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	} else {
		p.peekError(token.SEMICOLON)
		return nil
	}
	return stmt
}

/*
parseExpression checks whether we have a parsing function associated with p.curToken.Type in the prefix position.
If we do, it calls this parsing function, else returns nil.
*/
func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		return nil
	}
	leftExp := prefix()
	return leftExp
}

/*
parseIdentifier returns a *ast.Identifier with the current
token in the Token field & value of the token in Value.

It doesn’t advance the tokens, it doesn’t call nextToken.
All of our parsing funcs, prefixParseFn/infixParseFn, are going to follow this protocol:
start with curToken being the type of token you’re associated with and return with curToken being the last token that’s part of your
expression type.
Never advance the tokens too far
*/
func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Value}
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.curToken}
	value, err := strconv.ParseInt(p.curToken.Value, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Value)
		p.errors = append(p.errors, msg)
		return nil
	}
	lit.Value = value
	return lit
}
