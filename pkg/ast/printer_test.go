package ast

import (
	"testing"

	"github.com/templecloud/glu/pkg/token"
)

func TestAst_Printer(t *testing.T) {
	expr := NewBinary(
		NewUnary(
			token.New(token.Minus, "-", "", 0, 0),
			NewLiteral("123"),
		),
		token.New(token.Star, "*", "", 0, 0),
		NewLiteral("123"),
	)
	printer := Printer{}
	output := printer.Print(expr)

	expected := "(* (- 123) 123)"
	if expected != output {
		t.Fatalf("test[1] - . Expected=%q, got=%q", expected, output)
	}
}
