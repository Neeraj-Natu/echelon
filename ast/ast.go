package ast

import (
	"bytes"
	"github.com/Neeraj-Natu/shifu/token"
	"strings"
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

	for _, s := range p.Statements {
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
	Token       token.Token // the return token
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

//BlockStatement is a series of Statements that are within curly braces.
type BlockStatement struct {
	Token      token.Token // the { token
	Statements []Statement
}

func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer
	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

// The token field what every statement has and the actual expression field that holds the entire expression.
type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
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
func (v *Variable) String() string       { return v.Value }

// This is to hold the integers in the expression statement. this implements the expression interface so it's an expression node.
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }

// This is to hold the Expressions that start with prefixes, the prefix could be '-' or '!'. this implements the expression interface so it's an expression node.
// Any Prefix Expression has 2 parts (<prefix> <Expression>) thus is also called unary operator as it has one Expression involved.
type PrefixExpression struct {
	Token    token.Token // The prefix token eg: '!'
	Operator string      // The string that contains the prefix '!' or '-'
	Right    Expression  // The expression to the right of the operator
}

func (pe *PrefixExpression) expressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")
	return out.String()
}

//This is to hold the Expressions that have an Infix operator. the Infix could be any operator '+,-,*,/,>,<,!=,=='. this implements the expression interface thus is an expression node.
// Any Infix Expression has 3 parts (<Left Expression> <Operator> <Right Expression>) thus this is called binary operator as it has two Expressions involved.
type InfixExpression struct {
	Token    token.Token // The infix operator token, eg: '+' or '-'
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode()      {}
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")
	return out.String()
}

type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) expressionNode()      {}
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (b *Boolean) String() string       { return b.Token.Literal }

// This is to hold the if else statement structure. Every such if statement has following format: (if (<Condition>) <Consequence> else <Alternative>).
// The below strct is to store such expressions.
type IfExpression struct {
	Token       token.Token // The 'if' token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode()      {}
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}
	return out.String()
}

//FunctionLiteral is to hold all the functions in the language. Every function can be represented as 'func <parameters> <block statement>'.
//Functions are firstclass citizens here which means these can be used as expression so shouldn't be a surprise when functionLiteral implements the expressionNode
type FunctionLiteral struct {
	Token      token.Token // The 'func' token
	Parameters []*Variable
	Body       *BlockStatement
}

func (fl *FunctionLiteral) expressionNode()      {}
func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ","))
	out.WriteString(")")
	out.WriteString(fl.Body.String())

	return out.String()
}
