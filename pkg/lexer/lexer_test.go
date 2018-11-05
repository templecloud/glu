package lexer

import (
	"testing"

	"github.com/templecloud/glu/pkg/token"
)

func validateTestToken() {

}

func TestScanTokens_Structural(t *testing.T) {
	input := "(){},.;"
	tests := []struct {
		expectedType   token.Type
		expectedLexeme string
		expectedLine   int
		expectedColumn int
	}{
		{token.LeftParen, "(", 0, 0},
		{token.RightParen, ")", 0, 1},
		{token.LeftBrace, "{", 0, 2},
		{token.RightBrace, "}", 0, 3},
		{token.Comma, ",", 0, 4},
		{token.Dot, ".", 0, 5},
		{token.Semicolon, ";", 0, 6},
	}

	l := New(input)
	tks, _ := l.ScanTokens()
	for idx, tt := range tests {
		tk := tks[idx]
		if tk.Type != tt.expectedType {
			t.Fatalf("test[%d] - Wrong token.Type. Expected=%q, got=%q", idx, tt.expectedType, tk.Type)
		}
		if tk.Lexeme != tt.expectedLexeme {
			t.Fatalf("test[%d] - Wrong lexeme. Expected=%q, got=%q", idx, tt.expectedLexeme, tk.Lexeme)
		}
		if tk.Source.Line != tt.expectedLine {
			t.Fatalf("test[%d] - Wrong line. Expected=%d, got=%d", idx, tt.expectedLine, tk.Source.Line)
		}
		if tk.Source.Column != tt.expectedColumn {
			t.Fatalf("test[%d] - Wrong column. Expected=%d, got=%d", idx, tt.expectedColumn, tk.Source.Column)
		}
	}
}

func TestScanTokens_Arithmetic(t *testing.T) {
	input := "-+*/"
	tests := []struct {
		expectedType   token.Type
		expectedLexeme string
		expectedLine   int
		expectedColumn int
	}{
		{token.Minus, "-", 0, 0},
		{token.Plus, "+", 0, 1},
		{token.Star, "*", 0, 2},
		{token.ForwardSlash, "/", 0, 3},
	}

	l := New(input)
	tks, _ := l.ScanTokens()
	for idx, tt := range tests {
		tk := tks[idx]
		if tk.Type != tt.expectedType {
			t.Fatalf("test[%d] - Wrong token.Type. Expected=%q, got=%q", idx, tt.expectedType, tk.Type)
		}
		if tk.Lexeme != tt.expectedLexeme {
			t.Fatalf("test[%d] - Wrong lexeme. Expected=%q, got=%q", idx, tt.expectedLexeme, tk.Lexeme)
		}
		if tk.Source.Line != tt.expectedLine {
			t.Fatalf("test[%d] - Wrong line. Expected=%d, got=%d", idx, tt.expectedLine, tk.Source.Line)
		}
		if tk.Source.Column != tt.expectedColumn {
			t.Fatalf("test[%d] - Wrong column. Expected=%d, got=%d", idx, tt.expectedColumn, tk.Source.Column)
		}
	}
}

func TestScanTokens_Comparator(t *testing.T) {
	input := "! != = == > >= < <="
	tests := []struct {
		expectedType   token.Type
		expectedLexeme string
		expectedLine   int
		expectedColumn int
	}{
		{token.Not, "!", 0, 0},
		{token.NotEqual, "!=", 0, 2},
		{token.Equal, "=", 0, 5},
		{token.EqualEqual, "==", 0, 7},
		{token.GreaterThan, ">", 0, 10},
		{token.GreaterThanOrEqual, ">=", 0, 12},
		{token.LessThan, "<", 0, 15},
		{token.LessThanOrEqual, "<=", 0, 17},
	}

	l := New(input)
	tks, _ := l.ScanTokens()
	for idx, tt := range tests {
		tk := tks[idx]
		if tk.Type != tt.expectedType {
			t.Fatalf("test[%d] - Wrong token.Type. Expected=%q, got=%q", idx, tt.expectedType, tk.Type)
		}
		if tk.Lexeme != tt.expectedLexeme {
			t.Fatalf("test[%d] - Wrong lexeme. Expected=%q, got=%q", idx, tt.expectedLexeme, tk.Lexeme)
		}
		if tk.Source.Line != tt.expectedLine {
			t.Fatalf("test[%d] - Wrong line. Expected=%d, got=%d", idx, tt.expectedLine, tk.Source.Line)
		}
		if tk.Source.Column != tt.expectedColumn {
			t.Fatalf("test[%d] - Wrong column. Expected=%d, got=%d", idx, tt.expectedColumn, tk.Source.Column)
		}
	}
}

func TestScanTokens_SingleLineComments(t *testing.T) {
	input := "// Commented out."
	l := New(input)
	expectedNumTokens := 0
	tks, _ := l.ScanTokens()
	actualNumTokens := len(tks)

	if actualNumTokens != 0 {
		t.Fatalf("test[%d] - Wrong number of tokens. Expected=%q, got=%q",
			0, expectedNumTokens, actualNumTokens)
	}
}

func TestScanTokens_EscapedNewLine(t *testing.T) {
	input := "test\n test\n  test\n"
	tests := []struct {
		expectedType   token.Type
		expectedLexeme string
		expectedLine   int
		expectedColumn int
	}{
		{token.Identifier, "test", 0, 0},
		{token.Identifier, "test", 1, 1},
		{token.Identifier, "test", 2, 2},
	}

	l := New(input)
	tks, _ := l.ScanTokens()
	for idx, tt := range tests {
		tk := tks[idx]
		if tk.Type != tt.expectedType {
			t.Fatalf("test[%d] - Wrong token.Type. Expected=%q, got=%q", idx, tt.expectedType, tk.Type)
		}
		if tk.Lexeme != tt.expectedLexeme {
			t.Fatalf("test[%d] - Wrong lexeme. Expected=%q, got=%q", idx, tt.expectedLexeme, tk.Lexeme)
		}
		if tk.Source.Line != tt.expectedLine {
			t.Fatalf("test[%d] - Wrong line. Expected=%d, got=%d", idx, tt.expectedLine, tk.Source.Line)
		}
		if tk.Source.Column != tt.expectedColumn {
			t.Fatalf("test[%d] - Wrong column. Expected=%d, got=%d", idx, tt.expectedColumn, tk.Source.Column)
		}
	}
}

func TestScanTokens_WhiteSpace(t *testing.T) {
	input := "\ttest\r\n"
	tests := []struct {
		expectedType   token.Type
		expectedLexeme string
		expectedLine   int
		expectedColumn int
	}{
		{token.Identifier, "test", 0, 1},
	}

	l := New(input)
	tks, _ := l.ScanTokens()
	for idx, tt := range tests {
		tk := tks[idx]
		if tk.Type != tt.expectedType {
			t.Fatalf("test[%d] - Wrong token.Type. Expected=%q, got=%q", idx, tt.expectedType, tk.Type)
		}
		if tk.Lexeme != tt.expectedLexeme {
			t.Fatalf("test[%d] - Wrong lexeme. Expected=%q, got=%q", idx, tt.expectedLexeme, tk.Lexeme)
		}
		if tk.Source.Line != tt.expectedLine {
			t.Fatalf("test[%d] - Wrong line. Expected=%d, got=%d", idx, tt.expectedLine, tk.Source.Line)
		}
		if tk.Source.Column != tt.expectedColumn {
			t.Fatalf("test[%d] - Wrong column. Expected=%d, got=%d", idx, tt.expectedColumn, tk.Source.Column)
		}
	}
}

func TestScanTokens_ShellEscapedWhiteSpace(t *testing.T) {
	input := "\\ttest\\r\\n"
	tests := []struct {
		expectedType   token.Type
		expectedLexeme string
		expectedLine   int
		expectedColumn int
	}{
		{token.Identifier, "test", 0, 2},
	}

	l := New(input)
	tks, _ := l.ScanTokens()
	for idx, tt := range tests {
		tk := tks[idx]
		if tk.Type != tt.expectedType {
			t.Fatalf("test[%d] - Wrong token.Type. Expected=%q, got=%q", idx, tt.expectedType, tk.Type)
		}
		if tk.Lexeme != tt.expectedLexeme {
			t.Fatalf("test[%d] - Wrong lexeme. Expected=%q, got=%q", idx, tt.expectedLexeme, tk.Lexeme)
		}
		if tk.Source.Line != tt.expectedLine {
			t.Fatalf("test[%d] - Wrong line. Expected=%d, got=%d", idx, tt.expectedLine, tk.Source.Line)
		}
		if tk.Source.Column != tt.expectedColumn {
			t.Fatalf("test[%d] - Wrong column. Expected=%d, got=%d", idx, tt.expectedColumn, tk.Source.Column)
		}
	}
}

func TestScanTokens_BadShellEscapedWhiteSpace(t *testing.T) {
	input := "\\n\\n\\q"
	l := New(input)
	tks, errs := l.ScanTokens()

	expectedNumTokens := 0
	actualNumTokens := len(tks)
	if actualNumTokens != 0 {
		t.Fatalf("test[%d] - Wrong number of tokens. Expected=%q, got=%q",
			0, expectedNumTokens, actualNumTokens)
	}

	err := errs[0]
	expectedMessage := "Unexpected escape character: q."
	expectedLine := 2
	expectedColumn := 2
	if err.Message != expectedMessage {
		t.Fatalf("test[%d] - Wrong Error. Expected=%q, got=%q", 0, expectedMessage, err.Message)
	}
	if err.Source.Line != expectedLine {
		t.Fatalf("test[%d] - Wrong line. Expected=%d, got=%d", 0, expectedLine, err.Source.Line)
	}
	if err.Source.Column != expectedColumn {
		t.Fatalf("test[%d] - Wrong column. Expected=%d, got=%d", 0, expectedColumn, err.Source.Column)
	}
}

func TestScanTokens_String(t *testing.T) {
	input := "\"s1\" \"s2\" \"s3\""
	tests := []struct {
		expectedType   token.Type
		expectedLexeme string
		expectedLine   int
		expectedColumn int
	}{
		{token.String, "s1", 0, 2},
		{token.String, "s2", 0, 7},
		{token.String, "s3", 0, 12},
	}

	l := New(input)
	tks, _ := l.ScanTokens()
	for idx, tt := range tests {
		tk := tks[idx]
		if tk.Type != tt.expectedType {
			t.Fatalf("test[%d] - Wrong token.Type. Expected=%q, got=%q", idx, tt.expectedType, tk.Type)
		}
		if tk.Lexeme != tt.expectedLexeme {
			t.Fatalf("test[%d] - Wrong lexeme. Expected=%q, got=%q", idx, tt.expectedLexeme, tk.Lexeme)
		}
		if tk.Source.Line != tt.expectedLine {
			t.Fatalf("test[%d] - Wrong line. Expected=%d, got=%d", idx, tt.expectedLine, tk.Source.Line)
		}
		if tk.Source.Column != tt.expectedColumn {
			t.Fatalf("test[%d] - Wrong column. Expected=%d, got=%d", idx, tt.expectedColumn, tk.Source.Column)
		}
	}
}

func TestScanTokens_UnterminatedString(t *testing.T) {
	input := "\"s1\" \"s2"
	tests := []struct {
		expectedType   token.Type
		expectedLexeme string
		expectedLine   int
		expectedColumn int
	}{
		{token.String, "s1", 0, 2},
	}

	l := New(input)
	tks, errs := l.ScanTokens()
	for idx, tt := range tests {
		tk := tks[idx]
		if tk.Type != tt.expectedType {
			t.Fatalf("test[%d] - Wrong token.Type. Expected=%q, got=%q", idx, tt.expectedType, tk.Type)
		}
		if tk.Lexeme != tt.expectedLexeme {
			t.Fatalf("test[%d] - Wrong lexeme. Expected=%q, got=%q", idx, tt.expectedLexeme, tk.Lexeme)
		}
		if tk.Source.Line != tt.expectedLine {
			t.Fatalf("test[%d] - Wrong line. Expected=%d, got=%d", idx, tt.expectedLine, tk.Source.Line)
		}
		if tk.Source.Column != tt.expectedColumn {
			t.Fatalf("test[%d] - Wrong column. Expected=%d, got=%d", idx, tt.expectedColumn, tk.Source.Column)
		}
	}

	err := errs[0]
	expectedMessage := "Unterminated string."
	expectedLine := 0
	expectedColumn := 8
	if err.Message != expectedMessage {
		t.Fatalf("test[%d] - Wrong Error. Expected=%q, got=%q", 0, expectedMessage, err.Message)
	}
	if err.Source.Line != expectedLine {
		t.Fatalf("test[%d] - Wrong line. Expected=%d, got=%d", 0, expectedLine, err.Source.Line)
	}
	if err.Source.Column != expectedColumn {
		t.Fatalf("test[%d] - Wrong column. Expected=%d, got=%d", 0, expectedColumn, err.Source.Column)
	}
}

func TestScanTokens_Numeric(t *testing.T) {
	input := "12 12.34 .12"
	tests := []struct {
		expectedType   token.Type
		expectedLexeme string
		expectedLine   int
		expectedColumn int
	}{
		{token.Number, "12", 0, 0},
		{token.Number, "12.34", 0, 3},
		{token.Dot, ".", 0, 9},
		{token.Number, "12", 0, 10},
	}

	l := New(input)
	tks, _ := l.ScanTokens()
	for idx, tt := range tests {
		tk := tks[idx]
		if tk.Type != tt.expectedType {
			t.Fatalf("test[%d] - Wrong token.Type. Expected=%q, got=%q", idx, tt.expectedType, tk.Type)
		}
		if tk.Lexeme != tt.expectedLexeme {
			t.Fatalf("test[%d] - Wrong lexeme. Expected=%q, got=%q", idx, tt.expectedLexeme, tk.Lexeme)
		}
		if tk.Source.Line != tt.expectedLine {
			t.Fatalf("test[%d] - Wrong line. Expected=%d, got=%d", idx, tt.expectedLine, tk.Source.Line)
		}
		if tk.Source.Column != tt.expectedColumn {
			t.Fatalf("test[%d] - Wrong column. Expected=%d, got=%d", idx, tt.expectedColumn, tk.Source.Column)
		}
	}
}

func TestScanTokens_Keyword_Logical(t *testing.T) {
	input := "nil true false and or"
	tests := []struct {
		expectedType   token.Type
		expectedLexeme string
		expectedLine   int
		expectedColumn int
	}{
		{token.Nil, "nil", 0, 0},
		{token.True, "true", 0, 4},
		{token.False, "false", 0, 9},
		{token.And, "and", 0, 15},
		{token.Or, "or", 0, 19},
	}

	l := New(input)
	tks, _ := l.ScanTokens()
	for idx, tt := range tests {
		tk := tks[idx]
		if tk.Type != tt.expectedType {
			t.Fatalf("test[%d] - Wrong token.Type. Expected=%q, got=%q", idx, tt.expectedType, tk.Type)
		}
		if tk.Lexeme != tt.expectedLexeme {
			t.Fatalf("test[%d] - Wrong lexeme. Expected=%q, got=%q", idx, tt.expectedLexeme, tk.Lexeme)
		}
		if tk.Source.Line != tt.expectedLine {
			t.Fatalf("test[%d] - Wrong line. Expected=%d, got=%d", idx, tt.expectedLine, tk.Source.Line)
		}
		if tk.Source.Column != tt.expectedColumn {
			t.Fatalf("test[%d] - Wrong column. Expected=%d, got=%d", idx, tt.expectedColumn, tk.Source.Column)
		}
	}
}

func TestScanTokens_Keyword_Declaration(t *testing.T) {
	input := "let func"
	tests := []struct {
		expectedType   token.Type
		expectedLexeme string
		expectedLine   int
		expectedColumn int
	}{
		{token.Let, "let", 0, 0},
		{token.Func, "func", 0, 4},
	}

	l := New(input)
	tks, _ := l.ScanTokens()
	for idx, tt := range tests {
		tk := tks[idx]
		if tk.Type != tt.expectedType {
			t.Fatalf("test[%d] - Wrong token.Type. Expected=%q, got=%q", idx, tt.expectedType, tk.Type)
		}
		if tk.Lexeme != tt.expectedLexeme {
			t.Fatalf("test[%d] - Wrong lexeme. Expected=%q, got=%q", idx, tt.expectedLexeme, tk.Lexeme)
		}
		if tk.Source.Line != tt.expectedLine {
			t.Fatalf("test[%d] - Wrong line. Expected=%d, got=%d", idx, tt.expectedLine, tk.Source.Line)
		}
		if tk.Source.Column != tt.expectedColumn {
			t.Fatalf("test[%d] - Wrong column. Expected=%d, got=%d", idx, tt.expectedColumn, tk.Source.Column)
		}
	}
}

func TestScanTokens_Keyword_Utility(t *testing.T) {
	input := "log"
	tests := []struct {
		expectedType   token.Type
		expectedLexeme string
		expectedLine   int
		expectedColumn int
	}{
		{token.Log, "log", 0, 0},
	}

	l := New(input)
	tks, _ := l.ScanTokens()
	for idx, tt := range tests {
		tk := tks[idx]
		if tk.Type != tt.expectedType {
			t.Fatalf("test[%d] - Wrong token.Type. Expected=%q, got=%q", idx, tt.expectedType, tk.Type)
		}
		if tk.Lexeme != tt.expectedLexeme {
			t.Fatalf("test[%d] - Wrong lexeme. Expected=%q, got=%q", idx, tt.expectedLexeme, tk.Lexeme)
		}
		if tk.Source.Line != tt.expectedLine {
			t.Fatalf("test[%d] - Wrong line. Expected=%d, got=%d", idx, tt.expectedLine, tk.Source.Line)
		}
		if tk.Source.Column != tt.expectedColumn {
			t.Fatalf("test[%d] - Wrong column. Expected=%d, got=%d", idx, tt.expectedColumn, tk.Source.Column)
		}
	}
}

func TestScanTokens_Keyword_Identifier(t *testing.T) {
	input := "some_identifer a_ducker trousers59 _id _23"
	tests := []struct {
		expectedType   token.Type
		expectedLexeme string
		expectedLine   int
		expectedColumn int
	}{
		{token.Identifier, "some_identifer", 0, 0},
		{token.Identifier, "a_ducker", 0, 15},
		{token.Identifier, "trousers59", 0, 24},
		{token.Identifier, "_id", 0, 35},
		{token.Identifier, "_23", 0, 39},
	}

	l := New(input)
	tks, _ := l.ScanTokens()
	for idx, tt := range tests {
		tk := tks[idx]
		if tk.Type != tt.expectedType {
			t.Fatalf("test[%d] - Wrong token.Type. Expected=%q, got=%q", idx, tt.expectedType, tk.Type)
		}
		if tk.Lexeme != tt.expectedLexeme {
			t.Fatalf("test[%d] - Wrong lexeme. Expected=%q, got=%q", idx, tt.expectedLexeme, tk.Lexeme)
		}
		if tk.Source.Line != tt.expectedLine {
			t.Fatalf("test[%d] - Wrong line. Expected=%d, got=%d", idx, tt.expectedLine, tk.Source.Line)
		}
		if tk.Source.Column != tt.expectedColumn {
			t.Fatalf("test[%d] - Wrong column. Expected=%d, got=%d", idx, tt.expectedColumn, tk.Source.Column)
		}
	}
}

func TestScanTokens_SimpleFunction(t *testing.T) {
	input := "func somefunc() { let x = 5; return x; }"
	tests := []struct {
		expectedType   token.Type
		expectedLexeme string
		expectedLine   int
		expectedColumn int
	}{
		{token.Func, "func", 0, 0},
		{token.Identifier, "somefunc", 0, 5},
		{token.LeftParen, "(", 0, 13},
		{token.RightParen, ")", 0, 14},
		{token.LeftBrace, "{", 0, 16},
		{token.Let, "let", 0, 18},
		{token.Identifier, "x", 0, 22},
		{token.Equal, "=", 0, 24},
		{token.Number, "5", 0, 26},
		{token.Semicolon, ";", 0, 27},
		{token.Return, "return", 0, 29},
		{token.Identifier, "x", 0, 36},
		{token.Semicolon, ";", 0, 37},
		{token.RightBrace, "}", 0, 39},
	}
	l := New(input)
	tks, _ := l.ScanTokens()
	for idx, tt := range tests {
		tk := tks[idx]
		if tk.Type != tt.expectedType {
			t.Fatalf("test[%d] - Wrong token.Type. Expected=%q, got=%q", idx, tt.expectedType, tk.Type)
		}
		if tk.Lexeme != tt.expectedLexeme {
			t.Fatalf("test[%d] - Wrong lexeme. Expected=%q, got=%q", idx, tt.expectedLexeme, tk.Lexeme)
		}
		if tk.Source.Line != tt.expectedLine {
			t.Fatalf("test[%d] - Wrong line. Expected=%d, got=%d", idx, tt.expectedLine, tk.Source.Line)
		}
		if tk.Source.Column != tt.expectedColumn {
			t.Fatalf("test[%d] - Wrong column. Expected=%d, got=%d", idx, tt.expectedColumn, tk.Source.Column)
		}
	}
}
