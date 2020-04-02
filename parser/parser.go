package parser

import (
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

	curToken token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l:l}
	
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

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement() 
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {}