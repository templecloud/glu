package interpreter

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/templecloud/glu/pkg/lexer"
	"github.com/templecloud/glu/pkg/parser"
)

func TestEvaluate_ExpressionStatemnt(t *testing.T) {
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

func TestEvaluate_ExpressionStatementFailure(t *testing.T) {
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

func TestEvaluate_LogStatement(t *testing.T) {
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
		i := Interpreter{}
		actual := i.Evaluate(stmts[0])

		if actual != nil {
			t.Fatalf(
				"test[%d] Expected nil return from log statement - Input=%s, ExpectedValue=nil, ActualValue=%v",
				idx, tt.input, actual)
		}
	}
}

func TestEvaluate_LogStatementFailure(t *testing.T) {
}

func TestEvaluateStdOut_LogStatement(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"log 1 + 1;", "2\n"},
	}
	pwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to initialise test: %v", err)
	}
	pwd = filepath.Dir(filepath.Dir(pwd))

	for idx, tt := range tests {
		cmd := fmt.Sprintf("%s/%s", pwd, "dist/glu")
		out, err := exec.Command(cmd, tt.input).Output()
		if err != nil {
			t.Fatalf(
				"test[%d] Expected no error - Input=%s, ExpectedValue=%v, Error=%v",
				idx, tt.input, tt.expected, err)
		}
		actual := string(out)
		if tt.expected != actual {
			t.Fatalf("test[%d] - Expected=%q, Actual=%q", idx, tt.expected, actual)
		}
	}
}
