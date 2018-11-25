package ast

import (
	"testing"

	"github.com/templecloud/glu/pkg/token"
)

func TestAst_Printer(t *testing.T) {
	expr := NewBinary(
		NewUnary(
			&token.Token{Type: token.Minus, Lexeme: "-"},
			NewLiteral(token.Minus, "123"),
		),
		&token.Token{Type: token.Star, Lexeme: "*"},
		NewLiteral(token.Star, "123"),
	)
	printer := Printer{}
	output := printer.Print(expr)

	expected := "(* (- 123) 123)"
	if expected != output {
		t.Fatalf("test[1] - . Expected=%q, got=%q", expected, output)
	}
}
