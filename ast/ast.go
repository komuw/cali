package ast

import (
	"bytes"

	"github.com/komuw/cali/token"
)

/*
 Every node in our AST has to implement the Node interface.
 TokenValue() is only used for debugging.
*/
type Node interface {
	TokenValue() string
	String() string // makes debugging easy
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
func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
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
func (ls *LetStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ls.TokenValue() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")
	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")
	return out.String()
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
func (i *Identifier) String() string {
	return i.Value
}

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
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(rs.TokenValue() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

/*
ExpressionStatement implements the Statement interface
It is used to represent an expression statement like:
  x + 10;

The reason we have ExpressionStatement satisfying Statement interface is because,
Program node is the root node of any AST. And Program.Statements takes a slice of Statement Nodes
Since we would want Expressions to be added to []Statements, they need to satisfy Statement interface.
*/
type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression  // holds the expression
}

func (es *ExpressionStatement) statementNode()     {}
func (es *ExpressionStatement) TokenValue() string { return es.Token.Value }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

/*
IntegerLiteral satisfies Expression interface.
Value is an int64 and not a string.
This is the field containing the actual value the integer literal represents.
*/
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode()    {}
func (il *IntegerLiteral) TokenValue() string { return il.Token.Value }
func (il *IntegerLiteral) String() string     { return il.Token.Value }
