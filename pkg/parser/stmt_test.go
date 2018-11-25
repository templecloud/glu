package parser

import (
	"testing"

	"github.com/templecloud/glu/pkg/ast"
	"github.com/templecloud/glu/pkg/lexer"
)

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

func TestParse_VarStatement(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"var x = 1 + 1;", "(#vs x = (+ 1 1))"},
		{"var x = 12345;", "(#vs x = 12345)"},
		{"var x = \"12345\";", "(#vs x = \"12345\")"},
		{"var x = \"test\";", "(#vs x = \"test\")"},
		{"var x = test;", "(#vs x = test)"},
	}
	for idx, tt := range tests {
		l := lexer.New(tt.input)
		tokens, _ := l.ScanTokens()
		p := New(tokens)
		expr := p.Parse()
		if len(expr) < 1 {
			t.Fatalf("test[%d] - Expected=%q, Actual=%v", idx, tt.expected, nil)
		}
		printer := ast.Printer{}
		actual := printer.Print(expr[0])
		if tt.expected != actual {
			t.Fatalf("test[%d] - Expected=%q, Actual=%q", idx, tt.expected, actual)
		}
	}
}
