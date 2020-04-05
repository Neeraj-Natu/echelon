package parser

import (
	"fmt"
	"github.com/Neeraj-Natu/shifu/ast"
	"github.com/Neeraj-Natu/shifu/lexer"
	"github.com/Neeraj-Natu/shifu/token"
)

/*
Parers take source code as input (either as text or tokens)
and produce a data structure which represents this source
code. While building up the data structure, they unavoidably
analyse the input, checking that it conforms to the expected
structure. Thus the process of parsing is also called syntactic analysis.
The parser below is a recursive descent parser also called top down
operator precedence parser, sometimes called Pratt Parser.
*/


//Parser has three fields, l is a pointer to an instance of lexer on which we repeatedly call NextToken() to get next token input.
// curToken and peekToken work exactly the same as position and readPosition but for tokens.
type Parser struct {
	l *lexer.Lexer
	errors []string

	curToken token.Token
	peekToken token.Token
}

//This function creates an instance of the parser and initializes it.
func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:l,
		errors: []string{},	
	}
	
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

//This function parses every statement there is, 
//it selects which parser function should apply 
//for which type of statement based on the Identifier token.
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return nil	
	}
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

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmnt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.VARIABLE) {
		return nil
	}
	stmnt.Name = &ast.Variable{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: skipping expressions untill we encounter a semicolon
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmnt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}
	p.nextToken()

	// TODO: skipping expressions untill we encounter a semicolon
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt

}

func (p *Parser) curTokenIs (t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs (t token.TokenType) bool {
	return p.peekToken.Type == t
}

// Pretty common function amongst many parsers.
// This function enforce the correctnss of the 
// order of tokens by checking the type of the 
// peekToken and only if the type is correct it 
// advance to the tokens by calling nextToken()
// else it adds an error using the peekError 
// function and returns false
func (p *Parser) expectPeek (t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

