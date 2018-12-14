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

func TestParse_ForStmt(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"for (var i = 0; i < 10; i=i+1) { log i; }",
			"(#bs (#vs i = 0) (#ws (< i 10) (#bs (#bs (#ls i)) (#es (#as i = (+ i 1))))))"},
		{"for (var i = 10; i < 20; i=i+1) { log i; }",
			"(#bs (#vs i = 10) (#ws (< i 20) (#bs (#bs (#ls i)) (#es (#as i = (+ i 1))))))"},
		{"for (var i = 20; i >= 0; i=i-2) { log i; }",
			"(#bs (#vs i = 20) (#ws (>= i 0) (#bs (#bs (#ls i)) (#es (#as i = (- i 2))))))"},
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

func TestParseError_ForStmt(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"for var i = 20; i >= 0; i=i-2) { log i; }", "Expected '(' after 'for'."},
		{"for (var i = 20 i >= 0; i=i-2 { log i; }", "Expected ';' after variable declaration."},
		// ???
		// {"for (var i = 20; i >= 0; i=i-2 { log i; }", "Expected ')' after 'for'."},
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

func TestParse_FnStmt(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"func sayHi(name) { log \"Hello, \"; log name; }",
			"(#fn-stmt sayHi(name) { (#ls \"Hello, \"); (#ls name) })"},
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

func TestParseError_FnStmt(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// {"func sayHi(name) { log \"Hello, \"; log name; }", "Expect '(' after if condition."}, // TODO
		{"func (name) { log \"Hello, \"; log name; }", "Expected kind function."},
		{"func sayHi name) { log \"Hello, \"; log name; }", "Expected '(' after kind function."},
		{"func sayHi (name,) { log \"Hello, \"; log name; }", "Expected parameter name."},
		// {"func sayHi (a,b,c,d,e,f,g,h) { log \"Hello, \"; log name; }", "Cannot have more than 8 arguments."}, // TODO
		{"func sayHi (name { log \"Hello, \"; log name; }", "Expected ')' after arguments."},
		{"func sayHi (name) log \"Hello, \"; log name; }", "Expected '{' before kind function body."},
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

func TestParse_ReturnExpr(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"func add(a, b) { return a + b; }",
			"(#fn-stmt add(a, b) { (return (+ a b)) })"},
		{"{ func add(a, b) { return a + b; } var c = add(1, 2); log c;}",
			"(#bs (#fn-stmt add(a, b) { (return (+ a b)) }) (#vs c = (#call-expr add(1, 2))) (#ls c))"},
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


func TestParseError_ReturnExpr(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"func add(a, b) { return a + b }", "Expect ';' after value."},
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
