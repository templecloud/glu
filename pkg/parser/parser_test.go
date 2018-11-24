package parser

import (
	"testing"

	"github.com/templecloud/glu/pkg/ast"
	"github.com/templecloud/glu/pkg/lexer"
)

func TestParse_ExpressionStatement(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"123;", "(#es 123)"},
		{"-123;", "(#es (- 123))"},
		{"123 + 123;", "(#es (+ 123 123))"},
		{"123 - 123;", "(#es (- 123 123))"},
		{"123 * 123;", "(#es (* 123 123))"},
		{"123 / 123;", "(#es (/ 123 123))"},
		{"-123 * 123;", "(#es (* (- 123) 123))"},
		{"(-123 * 123);", "(#es (#g (* (- 123) 123)))"},
		{"(-123 * 123) / (123 - 123);", "(#es (/ (#g (* (- 123) 123)) (#g (- 123 123))))"},
	}
	for idx, tt := range tests {
		l := lexer.New(tt.input)
		tokens, _ := l.ScanTokens()
		p := New(tokens)
		expr := p.Parse()
		printer := ast.Printer{}
		actual := printer.Print(expr[0])
		if tt.expected != actual {
			t.Fatalf("test[%d] - Expected=%q, Actual=%q", idx, tt.expected, actual)
		}
	}
}

func TestParse_ExpressionStatementFailure(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"1 + 1", "Expect ';' after expression."},
		{"(1 + ", "Token failed to match any rule."},
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

func TestParse_LogStatement(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"log 1 + 1;", "(#ls (+ 1 1))"},
	}
	for idx, tt := range tests {
		l := lexer.New(tt.input)
		tokens, _ := l.ScanTokens()
		p := New(tokens)
		expr := p.Parse()
		printer := ast.Printer{}
		actual := printer.Print(expr[0])
		if tt.expected != actual {
			t.Fatalf("test[%d] - Expected=%q, Actual=%q", idx, tt.expected, actual)
		}
	}
}


func TestParse_LogStatementFailure(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"log 1 + 1", "Expect ';' after value."},
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