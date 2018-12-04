package parser

import (
	"testing"

	"github.com/templecloud/glu/pkg/ast"
	"github.com/templecloud/glu/pkg/lexer"
)

func TestParse_BlockStmt(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"{var a = 1; log a; { a = 2; log a; } log a;}",
			"(#bs (#vs a = 1) (#ls a) (#bs (#es (#as a = 2)) (#ls a)) (#ls a))"},
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

func TestParse_IfStmt(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"if (1 + 1 == 2) log \"wibble\";",
			"(#is (== (+ 1 1) 2) (#ls \"wibble\"))"},
		{"if (1 + 1 == 2) log \"wibble\"; else log \"wobble\";",
			"(#is (== (+ 1 1) 2) (#ls \"wibble\") (#ls \"wobble\"))"},
		{"if (1 + 1 == 2) { log \"wibble\"; }",
			"(#is (== (+ 1 1) 2) (#bs (#ls \"wibble\")))"},
		{"if (1 + 1 == 5) { log \"wibble\"; }",
			"(#is (== (+ 1 1) 5) (#bs (#ls \"wibble\")))"},
		{"if (1 + 1 == 2) { log \"wibble\"; } else { log \"wobble\"; }",
			"(#is (== (+ 1 1) 2) (#bs (#ls \"wibble\")) (#bs (#ls \"wobble\")))"},
		{"if (1 + 1 == 5) { log \"wibble\"; } else { log \"wobble\"; }",
			"(#is (== (+ 1 1) 5) (#bs (#ls \"wibble\")) (#bs (#ls \"wobble\")))"},
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

func TestParseError_IfStmt(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"if 1 + 1 == 2 { log \"wibble\"; }", "Expect '(' after if condition."},
		{"if (1 + 1 == 2 { log \"wibble\"; }", "Expect ')' after if condition."},
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

func TestParse_LogStmt(t *testing.T) {
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

func TestParseError_LogStmt(t *testing.T) {
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

func TestParse_VariableStmt(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"var x = 1 + 1;", "(#vs x = (+ 1 1))"},
		{"var x = 12345;", "(#vs x = 12345)"},
		{"var x = \"12345\";", "(#vs x = \"12345\")"},
		{"var x = \"test\";", "(#vs x = \"test\")"},
		{"var x = test;", "(#vs x = test)"},
		{"var x;", "(#vs x)"},
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

func TestParse_WhileStmt(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"{var x = 0; while (x > 0) { log \"*\"; x = x - 1; }}",
			"(#bs (#vs x = 0) (#ws (> x 0) (#bs (#ls \"*\") (#es (#as x = (- x 1))))))"},
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

func TestParseError_WhileStmt(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"{while x > 0) { log \"*\"; }}", "Expected '(' after while."},
		{"{while (x > 0 { log \"*\"; }}", "Expected ')' after while."},
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
