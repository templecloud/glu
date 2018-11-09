package parser

import (
	"testing"

	"github.com/templecloud/glu/pkg/ast"
	"github.com/templecloud/glu/pkg/lexer"
)

func TestParse_Expression(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"123", "123"},
		{"-123", "(- 123)"},
		{"123 + 123", "(+ 123 123)"},
		{"123 - 123", "(- 123 123)"},
		{"123 * 123", "(* 123 123)"},
		{"123 / 123", "(/ 123 123)"},
		{"-123 * 123", "(* (- 123) 123)"},
		{"(-123 * 123)", "(#g (* (- 123) 123))"},
		{"(-123 * 123) / (123 - 123)", "(/ (#g (* (- 123) 123)) (#g (- 123 123)))"},
	}
	for idx, tt := range tests {
		l := lexer.New(tt.input)
		tokens, _ := l.ScanTokens()
		p := New(tokens)
		expr := p.Parse()
		printer := ast.Printer{}
		actual := printer.Print(expr)
		if tt.expected != actual {
			t.Fatalf("test[%d] - Expected=%q, Actual=%q", idx, tt.expected, actual)
		}
	}

}
