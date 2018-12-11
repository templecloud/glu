package interpreter

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/templecloud/glu/pkg/lexer"
	"github.com/templecloud/glu/pkg/parser"
)

func TestEvaluate_ExprStmt(t *testing.T) {
	tests := []struct {
		input         string
		expectedValue interface{}
		expectedType  string
	}{
		{"123;", float64(123), "float64"},
		{"-123;", float64(-123), "float64"},
		{"-123.5;", float64(-123.5), "float64"},
		{"123.5;", float64(123.5), "float64"},
		// TODO
		// {"123.45", float64(123.45)},
		{"1 + 1;", float64(2), "float64"},
		{"1 + 1 + 1;", float64(3), "float64"},
		{"2.5 + 2.5;", float64(5), "float64"},
		{"2.5 + 3.5 + 4.0;", float64(10), "float64"},

		{"1 - 1;", float64(0), "float64"},
		{"1 - 2;", float64(-1), "float64"},

		{"2 * 0;", float64(0), "float64"},
		{"2 * 1;", float64(2), "float64"},
		{"2 * 2;", float64(4), "float64"},
		{"2 * -2;", float64(-4), "float64"},

		{"2 / 2;", float64(1), "float64"},
		{"2 / 1;", float64(2), "float64"},
		{"2 / (1 / 2);", float64(4), "float64"},
		// TODO: +Inf
		// {"2 / 0", float64(4), "float64"},

		{" 1 == 1;", true, "bool"},
		{" 1 == 2;", false, "bool"},
		{" 1 != 1;", false, "bool"},
		{" 1 != 2;", true, "bool"},
		{" 2 > 1;", true, "bool"},
		{" 1 > 1;", false, "bool"},
		{" 1 >= 1;", true, "bool"},
		{" 1 < 2;", true, "bool"},
		{" 1 < 1;", false, "bool"},
		{" 1 < 2;", true, "bool"},
	}
	for idx, tt := range tests {
		l := lexer.New(tt.input)
		tokens, _ := l.ScanTokens()
		p := parser.New(tokens)
		expr := p.Parse()
		i := New()

		actual, _ := i.Eval(expr[0])
		var actualValue interface{}
		actualType := reflect.TypeOf(actual)
		switch actual.(type) {
		case float64:
			actualValue = actual.(float64)
		case bool:
			actualValue = actual.(bool)
		default:
			t.Fatalf("test[%d] - Unknown type: %s", idx, reflect.TypeOf(actual))
		}

		if tt.expectedValue != actualValue {
			t.Fatalf("test[%d] - Input=%s, ExpectedValue=%v, ActualValue=%v", idx, tt.input, tt.expectedValue, actualValue)
		}
		if tt.expectedType != fmt.Sprintf("%s", actualType) {
			t.Fatalf("test[%d] - ExpectedType=%s, ActualType=%s", idx, tt.expectedType, actualType)
		}
	}
}

func TestEvaluateError_ExprStmt(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"- \"test\";", "Operand must be a number."},
		{"1 + \"test\";", "Operands must both be numbers."},
	}
	for idx, tt := range tests {
		l := lexer.New(tt.input)
		tokens, _ := l.ScanTokens()
		p := parser.New(tokens)
		stmts := p.Parse()
		result, evalErr := New().Eval(stmts[0])
		if result != nil {
			t.Fatalf("test[%d] - Expected nil result. Expected=%v, Actual=%v",
				idx, nil, result)
		}
		if evalErr == nil {
			t.Fatalf("test[%d] - Expected error result. Expected=%s, Actual=%s",
				idx, tt.expected, "nil")
		}
		if tt.expected != evalErr.message {
			t.Fatalf("test[%d] - Expected=%q, Actual=%q", idx, tt.expected, evalErr)
		}
	}
}

func TestEvaluateError_CallExpr(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"\"some-func\"();", "Can only call functions."},
	}
	for idx, tt := range tests {
		l := lexer.New(tt.input)
		tokens, _ := l.ScanTokens()
		p := parser.New(tokens)
		stmts := p.Parse()
		result, evalErr := New().Eval(stmts[0])
		if result != nil {
			t.Fatalf("test[%d] - Expected nil result. Expected=%v, Actual=%v",
				idx, nil, result)
		}
		if evalErr == nil {
			t.Fatalf("test[%d] - Expected error result. Expected=%s, Actual=%s",
				idx, tt.expected, "nil")
		}
		if tt.expected != evalErr.message {
			t.Fatalf("test[%d] - Expected=%q, Actual=%q", idx, tt.expected, evalErr)
		}
	}
}

func TestEvaluate_LogStmt(t *testing.T) {
	tests := []struct {
		input string
	}{
		{"log 123;"},
		{"log 1 + 1;"},
	}
	for idx, tt := range tests {
		l := lexer.New(tt.input)
		tokens, _ := l.ScanTokens()
		p := parser.New(tokens)
		stmts := p.Parse()
		actual, _ := New().Eval(stmts[0])

		if actual != nil {
			t.Fatalf(
				"test[%d] Expected nil return from log statement - Input=%s, ExpectedValue=nil, ActualValue=%v",
				idx, tt.input, actual)
		}
	}
}

func TestEvaluateError_LogStmt(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"log bob;", "Undefined variable 'bob'."},
	}
	for idx, tt := range tests {
		l := lexer.New(tt.input)
		tokens, _ := l.ScanTokens()
		p := parser.New(tokens)
		stmts := p.Parse()
		result, evalErr := New().Eval(stmts[0])
		if result != nil {
			t.Fatalf("test[%d] - Expected nil result. Expected=%v, Actual=%v",
				idx, result, nil)
		}
		if evalErr == nil {
			t.Fatalf("test[%d] - Expected error result. Expected=%s, Actual=%s",
				idx, tt.expected, "nil")
		}
		if tt.expected != evalErr.message {
			t.Fatalf("test[%d] - Expected=%q, Actual=%q", idx, tt.expected, evalErr.message)
		}
	}
}

func TestEvaluate_LogicalExpr(t *testing.T) {
	tests := []struct {
		input         string
		expectedValue bool
	}{
		{"true and true;", true},
		{"true and true and true;", true},
		{"false and true and true;", false},
		{"false and true and true;", false},
		{"true or true;", true},
		{"true or true or true;", true},
		{"false or true or true;", true},
		{"false or true or true;", true},
		{"true or true and true;", true},
		{"true or false and true;", true},
		{"true or false and false;", true},
		{"false or false and false;", false},
		{"false or true and false;", false},
		{"false or true and true;", true},
	}
	for idx, tt := range tests {
		l := lexer.New(tt.input)
		tokens, _ := l.ScanTokens()
		p := parser.New(tokens)
		expr := p.Parse()
		actualValue, _ := New().Eval(expr[0])
		if tt.expectedValue != actualValue {
			t.Fatalf("test[%d] - Input=%s, ExpectedValue=%v, ActualValue=%v", idx, tt.input, tt.expectedValue, actualValue)
		}
	}
}
