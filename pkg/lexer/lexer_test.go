package lexer

import (
	"testing"

	"github.com/templecloud/glu/pkg/token"
)

// Tests ============================================================
//

func TestScanTokens_Structural(t *testing.T) {
	input := "(){},.;"
	expected := []expectedToken{
		{token.LeftParen, "(", 0, 0, 1},
		{token.RightParen, ")", 0, 1, 1},
		{token.LeftBrace, "{", 0, 2, 1},
		{token.RightBrace, "}", 0, 3, 1},
		{token.Comma, ",", 0, 4, 1},
		{token.Dot, ".", 0, 5, 1},
		{token.Semicolon, ";", 0, 6, 1},
	}
	actual, _ := New(input).ScanTokens()
	validateTestTokens(t, expected, actual)
}

func TestScanTokens_Arithmetic(t *testing.T) {
	input := "-+*/"
	expected := []expectedToken{
		{token.Minus, "-", 0, 0, 1},
		{token.Plus, "+", 0, 1, 1},
		{token.Star, "*", 0, 2, 1},
		{token.ForwardSlash, "/", 0, 3, 1},
	}
	actual, _ := New(input).ScanTokens()
	validateTestTokens(t, expected, actual)
}

func TestScanTokens_Comparator(t *testing.T) {
	input := "! != = == > >= < <="
	expected := []expectedToken{
		{token.Not, "!", 0, 0, 1},
		{token.NotEqual, "!=", 0, 2, 2},
		{token.Equal, "=", 0, 5, 1},
		{token.EqualEqual, "==", 0, 7, 2},
		{token.GreaterThan, ">", 0, 10, 1},
		{token.GreaterThanOrEqual, ">=", 0, 12, 2},
		{token.LessThan, "<", 0, 15, 1},
		{token.LessThanOrEqual, "<=", 0, 17, 2},
	}
	actual, _ := New(input).ScanTokens()
	validateTestTokens(t, expected, actual)
}

func TestScanTokens_SingleLineComments(t *testing.T) {
	input := "// Commented out."
	actual, _ := New(input).ScanTokens()
	expectedNumTokens := 0
	actualNumTokens := len(actual)
	if expectedNumTokens != actualNumTokens {
		t.Fatalf("test[%d] - Wrong number of tokens. Expected=%q, Actual=%q",
			0, expectedNumTokens, actualNumTokens)
	}
}

func TestScanTokens_EscapedNewLine(t *testing.T) {
	input := "test\n test\n  test\n"
	expected := []expectedToken{
		{token.Identifier, "test", 0, 0, 4},
		{token.Identifier, "test", 1, 1, 4},
		{token.Identifier, "test", 2, 2, 4},
	}
	actual, _ := New(input).ScanTokens()
	validateTestTokens(t, expected, actual)
}

func TestScanTokens_WhiteSpace(t *testing.T) {
	input := "\ttest\r\n"
	expected := []expectedToken{
		{token.Identifier, "test", 0, 1, 4},
	}
	actual, _ := New(input).ScanTokens()
	validateTestTokens(t, expected, actual)
}

func TestScanTokens_ShellEscapedWhiteSpace(t *testing.T) {
	input := "\\ttest\\r\\n"
	expected := []expectedToken{
		{token.Identifier, "test", 0, 2, 4},
	}
	actual, _ := New(input).ScanTokens()
	validateTestTokens(t, expected, actual)
}

func TestScanTokens_BadShellEscapedWhiteSpace(t *testing.T) {
	input := "\\n\\n\\q"
	actual, errs := New(input).ScanTokens()
	expectedNumTokens := 0
	actualNumTokens := len(actual)
	if actualNumTokens != 0 {
		t.Fatalf("test[%d] - Wrong number of tokens. Expected=%q, Actual=%q",
			0, expectedNumTokens, actualNumTokens)
	}
	err := errs[0]
	expectedMessage := "Unexpected escape character: q."
	expectedLine := 2
	expectedColumn := 2
	if err.Message != expectedMessage {
		t.Fatalf("test[%d] - Wrong Error. Expected=%q, Actual=%q", 0, expectedMessage, err.Message)
	}
	if err.Source.Line != expectedLine {
		t.Fatalf("test[%d] - Wrong line. Expected=%d, Actual=%d", 0, expectedLine, err.Source.Line)
	}
	if err.Source.Column != expectedColumn {
		t.Fatalf("test[%d] - Wrong column. Expected=%d, Actual=%d", 0, expectedColumn, err.Source.Column)
	}
}

func TestScanTokens_String(t *testing.T) {
	input := "\"s1\" \"s2\" \"s3\""
	expected := []expectedToken{
		{token.String, "s1", 0, 2, 2},
		{token.String, "s2", 0, 7, 2},
		{token.String, "s3", 0, 12, 2},
	}
	actual, _ := New(input).ScanTokens()
	validateTestTokens(t, expected, actual)
}

func TestScanTokens_UnterminatedString(t *testing.T) {
	input := "\"s1\" \"s2"
	expected := []expectedToken{
		{token.String, "s1", 0, 2, 2},
	}
	actual, errs := New(input).ScanTokens()
	validateTestTokens(t, expected, actual)

	err := errs[0]
	expectedMessage := "Unterminated string."
	expectedLine := 0
	expectedColumn := 8
	if err.Message != expectedMessage {
		t.Fatalf("test[%d] - Wrong Error. Expected=%q, Actual=%q", 0, expectedMessage, err.Message)
	}
	if err.Source.Line != expectedLine {
		t.Fatalf("test[%d] - Wrong line. Expected=%d, Actual=%d", 0, expectedLine, err.Source.Line)
	}
	if err.Source.Column != expectedColumn {
		t.Fatalf("test[%d] - Wrong column. Expected=%d, Actual=%d", 0, expectedColumn, err.Source.Column)
	}
}

func TestScanTokens_Numeric(t *testing.T) {
	input := "12 12.34 .12"
	expected := []expectedToken{
		{token.Number, "12", 0, 0, 2},
		{token.Number, "12.34", 0, 3, 5},
		{token.Dot, ".", 0, 9, 1},
		{token.Number, "12", 0, 10, 2},
	}
	actual, _ := New(input).ScanTokens()
	validateTestTokens(t, expected, actual)
}

func TestScanTokens_Keyword_Logical(t *testing.T) {
	input := "nil true false and or"
	expected := []expectedToken{
		{token.Nil, "nil", 0, 0, 3},
		{token.True, "true", 0, 4, 4},
		{token.False, "false", 0, 9, 5},
		{token.And, "and", 0, 15, 3},
		{token.Or, "or", 0, 19, 2},
	}
	actual, _ := New(input).ScanTokens()
	validateTestTokens(t, expected, actual)
}

func TestScanTokens_Keyword_Declaration(t *testing.T) {
	input := "let func"
	expected := []expectedToken{
		{token.Let, "let", 0, 0, 3},
		{token.Func, "func", 0, 4, 4},
	}
	actual, _ := New(input).ScanTokens()
	validateTestTokens(t, expected, actual)
}

func TestScanTokens_Keyword_Utility(t *testing.T) {
	input := "log"
	expected := []expectedToken{
		{token.Log, "log", 0, 0, 3},
	}
	actual, _ := New(input).ScanTokens()
	validateTestTokens(t, expected, actual)
}

func TestScanTokens_Keyword_Identifier(t *testing.T) {
	input := "some_identifer a_ducker trousers59 _id _23"
	expected := []expectedToken{
		{token.Identifier, "some_identifer", 0, 0, 14},
		{token.Identifier, "a_ducker", 0, 15, 8},
		{token.Identifier, "trousers59", 0, 24, 10},
		{token.Identifier, "_id", 0, 35, 3},
		{token.Identifier, "_23", 0, 39, 3},
	}
	actual, _ := New(input).ScanTokens()
	validateTestTokens(t, expected, actual)
}

func TestScanTokens_SimpleFunction(t *testing.T) {
	input := "func somefunc() { let x = 5; return x; }"
	expected := []expectedToken{
		{token.Func, "func", 0, 0, 4},
		{token.Identifier, "somefunc", 0, 5, 8},
		{token.LeftParen, "(", 0, 13, 1},
		{token.RightParen, ")", 0, 14, 1},
		{token.LeftBrace, "{", 0, 16, 1},
		{token.Let, "let", 0, 18, 3},
		{token.Identifier, "x", 0, 22, 1},
		{token.Equal, "=", 0, 24, 1},
		{token.Number, "5", 0, 26, 1},
		{token.Semicolon, ";", 0, 27, 1},
		{token.Return, "return", 0, 29, 6},
		{token.Identifier, "x", 0, 36, 1},
		{token.Semicolon, ";", 0, 37, 1},
		{token.RightBrace, "}", 0, 39, 1},
	}
	actual, _ := New(input).ScanTokens()
	validateTestTokens(t, expected, actual)
}

// Support Functions ==========================================================
//

type expectedToken struct {
	expectedType   token.Type
	expectedLexeme string
	expectedLine   int
	expectedColumn int
	expectedLength int
}

func validateTestTokens(
	t *testing.T, expected []expectedToken, actual []*token.Token,
) {
	for idx, e := range expected {
		validateTestToken(t, string(idx), e, actual[idx])
	}
}

func validateTestToken(t *testing.T, name string, e expectedToken, a *token.Token) {
	if e.expectedType != a.Type {
		t.Fatalf("test[%s] - Wrong token.Type. Expected=%q, Actual=%q",
			name, e.expectedType, a.Type)
	}
	if e.expectedLexeme != a.Lexeme {
		t.Fatalf("test[%s] - Wrong lexeme. Expected=%q, Actual=%q",
			name, e.expectedLexeme, a.Lexeme)
	}
	if e.expectedLine != a.Source.Line {
		t.Fatalf("test[%s] - Wrong line. Expected=%d, Actual=%d",
			name, e.expectedLine, a.Source.Line)
	}
	if e.expectedColumn != a.Source.Column {
		t.Fatalf("test[%s] - Wrong column. Expected=%d, Actual=%d",
			name, e.expectedColumn, a.Source.Column)
	}
	if e.expectedLength != a.Source.Length {
		t.Fatalf("test[%s] - Wrong length. Expected=%d, Actual=%d",
			name, e.expectedLength, a.Source.Length)
	}
}
