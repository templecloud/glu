package interpreter

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestBinary_LogStmt(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"log 1 + 1;", "2\n"},
		{"log \"Hello\";", "Hello\n"},
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

func TestBinary_VarStmt(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"var x; log x;", "nil\n"},
		{"var x = 1 + 1; log x;", "2\n"},
		{"log x;", "runtime error: {&{Type:Identifier Lexeme:x Source:{Origin: Line:0 Column:4 Length:1}}, Undefined variable 'x'.}\n\n"},
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

func TestBinary_AssignExpr(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"var x;", ""},
		{"var x; x = 1;", "1\n"},
		{"var x; x = 1; log x;", "1\n"},
		{"var x = 1; var y = 2; log x + y;", "3\n"},
		{"var a; a = 1; var b = 3; log a + b;", "4\n"},
		{"1 + 2; 3 + 4;", "7\n"},
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
