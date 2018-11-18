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

func TestParse_ExpressionFailure(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// {&{Type:EOF Lexeme: Source:{Origin: Line:0 Column:5 Length:0}}, Token failed to match any rule.}
		{"(1 + ", "Token failed to match any rule."},
		// {&{Type:EOF Lexeme: Source:{Origin: Line:0 Column:6 Length:0}}, Expect ')' after expression.}
		{"(1 + 1", "Expect ')' after expression."},
	}
	for idx, tt := range tests {
		l := lexer.New(tt.input)
		tokens, _ := l.ScanTokens()
		p := New(tokens)
		p.Parse()
		actualErrorMessage := p.Errors[0].message
		if tt.expected != actualErrorMessage {
			t.Fatalf("test[%d] - Expected=%q, Actual=%q", idx, tt.expected, actualErrorMessage)
		}
	}
}

