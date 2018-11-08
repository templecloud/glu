package parser

import (
	"testing"

	"github.com/templecloud/glu/pkg/ast"
	"github.com/templecloud/glu/pkg/lexer"
)

func TestParse_Expression(t *testing.T) {
	input := "-123 * 123"
	l := lexer.New(input)
	tokens, _ := l.ScanTokens()
	p := New(tokens)
	expr := p.Parse()
	printer := ast.Printer{}
	output := printer.Print(expr)
	expected := "(* (- 123) 123)"
	if expected != output {
		t.Fatalf("test[1] - . Expected=%q, got=%q", expected, output)
	}
}
