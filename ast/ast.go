package ast

import "github.com/komuw/cali/token"

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
			Token: token.Token{Type: "IDENT", Value: "x", },
			Value: "x"},
		Value: nil, // TODO: update this to reflect that Value should be 5 when we start parsing INTs
	}

let x = 5;
can be represented by an AST like;
				*ast.Program
				Statements
					|
				*ast.LetStatement
*ast.Identifier <--  Name
					Value --> *ast.Expression
*/
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

// statementNode satisfies the Statement interface
func (ls *LetStatement) statementNode() {}

// TokenValue satisfies the Node interface
func (ls *LetStatement) TokenValue() string {
	return ls.Token.Value
}

/*
Identifier implements the Expression interface

It has fields;
1. one for the identifer(x) token
2. one for the expression(5)

Example of Identifier struct for; let x = 5;
	&ast.Identifier{
			Token: token.Token{Type: "IDENT", Value: "x", },
			Value: "x"
		}
*/
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string      //
}

// expressionNode implements the Expression interface
func (i *Identifier) expressionNode() {}

// TokenValue implements the Node interface
func (i *Identifier) TokenValue() string { return i.Token.Value }

/*
ReturnStatement implements Statement interface

return 5;
looks lik;
return <expression>;
*/
type ReturnStatement struct {
	Token       token.Token // token.RETURN
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()     {}
func (rs *ReturnStatement) TokenValue() string { return rs.Token.Value }
