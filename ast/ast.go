package ast

import "github.com/komuw/khaled/token"

/*
 Every node in our AST has to implement the Node interface.
 TokenValue() is only used for debugging.
*/
type Node interface {
	TokenValue() string
}

type Statement interface {
	Node
	statementNode() // dummy method; used to guide compiler to differentiate Statement and Expression interfaces
}

type Expression interface {
	Node
	expressionNode() //dummy method
}

/*
Program implements the Node interface.
Program node will be the root node of any AST produced by the parser.
Every valid program is a series/slice of statements
*/
type Program struct {
	Statements []Statement
}

func (p *Program) TokenValue() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenValue()
	} else {
		return ""
	}
}

/*
LetStatement implements the Statement interface

Look at an example like;
	let x = 5;
The fields its(AST) should have are;
1. one for let token
2. another for name of identifier(x)
3. finally one for the expression(5)

Example of LetStatement struct for; let x = 5;
	&ast.LetStatement{
		Token: token.Token{Type: "LET", Value: "let"},
		Name: &ast.Identifier{
			Token: token.Token{Type: "IDENT", Value: "x", }, Value: "x"},
		Value: nil, // TODO: update this to reflect that Value should be 5 when we start parsing INTs
	}
*/
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenValue() string {
	return ls.Token.Value
}

/*
Identifier implements the Expression interface

It has fields;
1. one for the identifer(x) token
2. one for the expression(5)

Hower in `let x = 5;` 5 is not an expression it is a statement ****
TODO: come back here.
*/
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string      //
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Value }

/*
let x = 5;
can be represented by an AST like;

					*ast.Program
					Statements
						|
					*ast.LetStatement
	*ast.Identifier <--  Name
						Value --> *ast.Expression
*/
