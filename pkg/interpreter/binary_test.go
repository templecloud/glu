package interpreter

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestBinary_AssignExpr(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"var x;", ""},
		{"var x; x = 1;", ""},
		{"var x; x = 1; log x;", "1\n"},
		{"var x = 1; var y = 2; log x + y;", "3\n"},
		{"var a; a = 1; var b = 3; log a + b;", "4\n"},
		{"1 + 2; log 3 + 4;", "7\n"},
		{"1 + 2; 3 + 4;", ""},
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

func TestBinary_BlockExpr(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"{var a = 5; log a; { var a = 4; log a; } log a;}", "545"},
		{"{var a = 5; log a; { a = 4; log a; } log a;}", "544"},
		{"{var a = \"a\"; log a; { var a = \"a2\"; log a; } log a;}", "aa2a"},
		{"var a = 5; log a; { var a = 4; log a; } log a;", "5\n45\n"}, // ???
		{"var a = 5; log a; { a = 4; log a; } log a;", "5\n44\n"},     // ???
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

func TestBinaryError_BlockExpr(t *testing.T) {
	tests := []struct {
		input          string
		expectedResult string
		expectedError  string
	}{
		{"{var a = 5; log a; { var b = 2; log b; var a = 4; log a; } log a; log b;}", "524", "Undefined variable 'b'"},
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
				idx, tt.input, tt.expectedResult, err)
		}
		actual := string(out)
		if !strings.Contains(actual, tt.expectedResult) {
			t.Fatalf("test[%d] - Expected=%q, Actual=%q", idx, tt.expectedResult, actual)
		}
		if !strings.Contains(actual, tt.expectedError) {
			t.Fatalf("test[%d] - Expected=%q, Actual=%q", idx, tt.expectedError, actual)
		}
	}
}

func TestBinary_IfStmt(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"if (1 + 1 == 2) log \"test1\";", "test1"},
		{"if (1 + 1 == 5) log \"test1\";", ""},
		{"if (1 + 1 == 5) log \"test1\"; else log \"test2\";", "test2"},
		{"if (1 + 1 == 2) { log \"test1\"; }", "test1"},
		{"if (1 + 1 == 5) { log \"test1\"; }", ""},
		{"if (1 + 1 == 5) { log \"test1\"; } else { log \"test2\"; }", "test2"},
		{"if (1 + 1 == 2) { log \"test1\"; } else { log \"test2\"; }", "test1"},
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

func TestBinary_LogicalStmt(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"if (true and true) { log true; }", "true"},
		{"if (true or true) { log true; }", "true"},
		{"if (true and true or true) { log true; }", "true"},

		{"if (true and false) { log true; }else { log false; }", "false"},
		{"if (true or false) { log true; } else { log false; }", "true"},
		{"if (true and false or false) { log true; } else { log false; }", "false"},

		{"if (true and false) { log true; } else { log false; }", "false"},
		{"if (true or false) { log true; } else { log false; }", "true"},
		{"if (false and true or true) { log true; } else { log false; }", "true"},
		{"if (true and false or true) { log true; } else { log false; }", "true"},
		{"if (true and true or false) { log true; } else { log false; }", "true"},
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
		{"log x;", "runtime error: {&{Type:Identifier Lexeme:x Source:{Origin: Line:0 Column:4 Length:1}}, Undefined variable 'x'.}\n"},
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

func TestBinary_WhileStmt(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"var x = 0; while (x > 0) { log \"*\"; x = x - 1; }", ""},
		{"var x = 1; while (x > 0) { log \"*\"; x = x - 1; }", "*"},
		{"var x = 2; while (x > 0) { log \"*\"; x = x - 1; }", "**"},
		{"var x = 3; while (x > 0) { log \"*\"; x = x - 1; }", "***"},
		{"var x = 6; while (x > 0) { log \"*\"; x = x - 1; }", "******"},
		{"var x = 2 * 5; while (x > 0) { log \"*\"; x = x - 1; }", "**********"},
		{"var x = 10; while (x > 0) { log \"*\"; x = x - 1; }", "**********"},
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

func TestBinary_ForStmt(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"for (var i = 0; i < 10; i=i+1) { log i; }", "0123456789"},
		{"for (var i = 10; i < 20; i=i+1) { log i; }", "10111213141516171819"},
		{"for (var i = 20; i >= 0; i=i-2) { log i; }", "20181614121086420"},
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

