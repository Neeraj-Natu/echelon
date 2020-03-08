package lexer

import (
	"testing"
	"github.com/Neeraj-Natu/shifu/token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){}[],;:`

	tests := []struct {
		expectedType token.TokenType
		expectedLiteral string
	} {
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LCBRACE, "{"},
		{token.RCBRACE, "}"},
		{token.LBRACE, "["},
		{token.RBRACE, "]"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.COLON, ":"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestAdvancedToken(t *testing.T) {
	input :=
	`
	let five = 5;
	let ten = 10;
	let add = func(x, y) {
  		x + y;
	};
	let result = add(five, ten);
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.VARIABLE, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.VARIABLE, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.VARIABLE, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "func"},
		{token.LPAREN, "("},
		{token.VARIABLE, "x"},
		{token.COMMA, ","},
		{token.VARIABLE, "y"},
		{token.RPAREN, ")"},
		{token.LCBRACE, "{"},
		{token.VARIABLE, "x"},
		{token.PLUS, "+"},
		{token.VARIABLE, "y"},
		{token.SEMICOLON, ";"},
		{token.RCBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.VARIABLE, "result"},
		{token.ASSIGN, "="},
		{token.VARIABLE, "add"},
		{token.LPAREN, "("},
		{token.VARIABLE, "five"},
		{token.COMMA, ","},
		{token.VARIABLE, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		t.Log("token to be tested")
		t.Log(tok.Literal)
		t.Log(tok.Type)
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
		
	}
}