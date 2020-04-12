package ast

import (
	"github.com/Neeraj-Natu/shifu/token"
	"bytes"
)

// The base Node interface
// The entire AST consists only of node, although each one of them is different based on which interface they implement
// Still at the end everything is a node, so all interfaces contain the Node interface.
type Node interface {
	TokenLiteral() string
	String() string
}

// All statement nodes implement this.
// as statements are a type of node.
type Statement interface {
	Node
	statementNode()
}

// All expression nodes implement this.
// as expressions are a type of nodes
type Expression interface {
	Node
	expressionNode()
}

//Program is a collection of numerous statements, also the root node for ast.
//Every valid program is a series of statements which are contained in the statements slice.
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// creates a buffer and writes the return value of each statement's String() method to it
// then returns the buffer as a string.
func (p *Program) String() string {
	var out bytes.Buffer

	for _,s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

// Any let statement has 3 parts first the token LET, second the variable name and third the expression that it points to on the right side of equals sign.
// Implements statement interface 
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Variable
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

// Any return statement has just 2 parts (return <expression>) the return keyword and an expression.
type ReturnStatement struct {
	Token token.Token // the return token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

// The token field what every statement has and the actual expression field that holds the entire expression.
type ExpressionStatement struct {
	Token token.Token  // the first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

// This is to hold the variable in the let statement. This implements the expression interface so it's an expression Node.
type Variable struct {
	Token token.Token // the token.VARIABLE token
	Value string
}

func (v *Variable) expressionNode()      {}
func (v *Variable) TokenLiteral() string { return v.Token.Literal }
func (v *Variable) String() string {return v.Value}

// This is to hold the integers in the expression statement. this implements the expression interface so it's also an expression node.
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }