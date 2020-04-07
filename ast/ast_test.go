package ast

import (
	"github.com/Neeraj-Natu/shifu/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Variable{
					Token: token.Token{Type: token.VARIABLE, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Variable{
					Token: token.Token{Type: token.VARIABLE, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}