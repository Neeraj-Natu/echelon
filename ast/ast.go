package ast

import (
	"github.com/Neeraj-Natu/shifu/token"
)

// The base Node interface
type Node interface {
	TokenLiteral() string
	// String() string
}

// All statement nodes implement this
type Statement interface {
	Node
	statementNode()
}

// All expression nodes implement this
type Expression interface {
	Node
	expressionNode()
}

//Program is a collection of numerous statements
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

// Any let statement has 3 parts first the token LET, second the variable name and third the expression that it points to on the right side of equals sign.
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Variable
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// This is to hold the variable in the let statement. This implements the expression interface.
type Variable struct {
	Token token.Token // the token.VARIABLE token
	Value string
}

func (i *Variable) expressionNode()      {}
func (i *Variable) TokenLiteral() string { return i.Token.Literal }
