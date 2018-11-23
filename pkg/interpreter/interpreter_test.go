package interpreter

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/templecloud/glu/pkg/lexer"
	"github.com/templecloud/glu/pkg/parser"
)

func TestEvaluate_Expression(t *testing.T) {
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
		i := Interpreter{}

		actual := i.Evaluate(expr[0])
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

func TestEvaluate_ExpressionFailure(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"1 + \"wibble\";", "Operands must both be numbers."},
	}
	for idx, tt := range tests {
		l := lexer.New(tt.input)
		tokens, _ := l.ScanTokens()
		p := parser.New(tokens)
		expr := p.Parse()
		i := Interpreter{}
		result := i.Evaluate(expr[0])
		if result != nil {
			t.Fatalf("test[%d] - Expected nil result. Expected=%v, Actual=%v",
				idx, result, nil)
		}
		if len(i.Errors) != 1 {
			t.Fatalf("test[%d] - Unexpected number of errors. Expected=%q, Actual=%q",
				idx, len(i.Errors), 1)
		}
		if tt.expected != i.Errors[0].message {
			t.Fatalf("test[%d] - Expected=%q, Actual=%q", idx, tt.expected, i.Errors[0].message)
		}
	}
}
