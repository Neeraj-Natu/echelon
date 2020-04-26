package parser

import (
	"fmt"
	"github.com/Neeraj-Natu/shifu/ast"
	"github.com/Neeraj-Natu/shifu/lexer"
	"github.com/Neeraj-Natu/shifu/token"
	"strconv"
)

/*
Parers take source code as input (either as text or tokens)
and produce a data structure which represents this source
code. While building up the data structure, they unavoidably
analyse the input, checking that it conforms to the expected
structure. Thus the process of parsing is also called syntactic analysis.
The parser below is a recursive descent parser also called top down
operator precedence parser, sometimes called Pratt Parser.
Pratt parser's main idea is the association of parsing functions with
token types. Whenever a certain token type is encountered, the
parsing functions are called to parse the appropriate expression
and return an AST node that represents it. Here Each token type
can have upto two parsing functions associated with it, depending
on whether the otken is found in a prefix or an infix position.
*/

const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	CALL        // myFunction(X)
)

//Predence table that associates token types with their precedences
var precedences = map[token.TokenType]int{
	token.EQ:       EQUALS,
	token.NOT_EQ:   EQUALS,
	token.LT:       LESSGREATER,
	token.GT:       LESSGREATER,
	token.PLUS:     SUM,
	token.MINUS:    SUM,
	token.SLASH:    PRODUCT,
	token.ASTERISK: PRODUCT,
	token.LPAREN:   CALL,
}

// Parser has three fields, l is a pointer to an instance of lexer on which we repeatedly call NextToken() to get next token input.
// curToken and peekToken work exactly the same as position and readPosition but for tokens.
type Parser struct {
	l      *lexer.Lexer
	errors []string

	curToken  token.Token
	peekToken token.Token

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

// This function creates an instance of the parser and initializes it.
func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.VARIABLE, p.parseVariable)
	p.registerPrefix(token.INT, p.parseIntegerLiteral)
	p.registerPrefix(token.BANG, p.parsePrefixExpression)
	p.registerPrefix(token.MINUS, p.parsePrefixExpression)

	p.infixParseFns = make(map[token.TokenType]infixParseFn)
	p.registerInfix(token.PLUS,p.parseInfixExpression)
	p.registerInfix(token.MINUS,p.parseInfixExpression)
	p.registerInfix(token.SLASH,p.parseInfixExpression)
	p.registerInfix(token.ASTERISK,p.parseInfixExpression)
	p.registerInfix(token.EQ,p.parseInfixExpression)
	p.registerInfix(token.NOT_EQ,p.parseInfixExpression)
	p.registerInfix(token.GT,p.parseInfixExpression)
	p.registerInfix(token.LT,p.parseInfixExpression)

	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

// This is a helper function that reads the tokens from lexer which will be parsed into an AST.
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}


// This function returns if there were any errors during parsing
func (p *Parser) Errors() []string {
	return p.errors
}

// Add an error to the errors slice when peekToken doesnot match the expectation.
func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}


func (p *Parser) noPrefixParseFnError(t token.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

// Pretty common function amongst many parsers.
// This function enforce the correctness of the
// order of tokens by checking the type of the
// peekToken and only if the type is correct it
// advance to the tokens by calling nextToken()
// else it adds an error using the peekError
// function and returns false
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

// Function to register prefix function for the token
func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

// Function to register infix function for the token
func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

// This function takes in a parser pointer and parses the tokens stored within it into an AST.
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

// This function parses every statement there is,
// it selects which parser function should apply
// for which type of statement based on the Identifier token.
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

// Function to parse Let statments
func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmnt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.VARIABLE) {
		return nil
	}
	stmnt.Name = &ast.Variable{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	p.nextToken()

	stmnt.Value = p.parseExpression(LOWEST)

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmnt
}

// Function to parse return statements
func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}
	p.nextToken()

	stmt.ReturnValue = p.parseExpression(LOWEST)

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt

}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}
	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

// Parsing function for all Expressions
func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		p.noPrefixParseFnError(p.curToken.Type)
		return nil
	}
	leftExp := prefix()

	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}
		p.nextToken()

		leftExp = infix(leftExp)
	}

	return leftExp
}

// Helper function to peek the precedence of next token
func (p *Parser) peekPrecedence() int{
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}

// Helper function to get the precedence of current token
func (p *Parser) curPrecedence() int{
	if p,ok := precedences[p.curToken.Type]; ok {
		return p
	}
	return LOWEST
}


// Parsing function for variables
func (p *Parser) parseVariable() ast.Expression {
	return &ast.Variable{Token: p.curToken, Value: p.curToken.Literal}
}

// Parsing function for integers
func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.curToken}

	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("Could not parse %q as integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	lit.Value = value
	return lit
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token: p.curToken,
		Operator: p.curToken.Literal,
	}
	p.nextToken()
	expression.Right = p.parseExpression(PREFIX)

	return expression
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expression := &ast.InfixExpression{
		Token: p.curToken,
		Operator: p.curToken.Literal,
		Left: left,
	}

	precedence := p.curPrecedence()
	p.nextToken()
	expression.Right = p.parseExpression(precedence)
	return expression
}
