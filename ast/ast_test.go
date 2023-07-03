package ast

import (
	"monkey/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Stmts: []Stmt{
			&LetStmt{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Ident{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Ident{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}
	if program.String() != "var myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}