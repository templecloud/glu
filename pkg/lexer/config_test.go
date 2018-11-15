package lexer

import (
	"testing"

	"github.com/templecloud/glu/pkg/token"
)

// Config Tests =====================================================
//

func TestScanTokens_AutoEOFToken_True(t *testing.T) {
	input := "test"
	expected := expectedToken{token.Identifier, input, 0, 0, 4}
	expected2 := expectedToken{token.EOF, "", 0, 4, 0}
	expectedNumTokens := 2
	actual, _ := New(input).ScanTokens()
	actualNumTokens := len(actual)
	if expectedNumTokens != actualNumTokens {
		t.Fatalf("test[%s] - Wrong number of tokens. Expected=%q, Actual=%q",
			"0", expectedNumTokens, actualNumTokens)
	}
	validateTestToken(t, "1", expected, actual[0])
	validateTestToken(t, "2", expected2, actual[1])
}

func TestScanTokens_AutoEOFToken_False(t *testing.T) {
	input := "test"
	expected := expectedToken{token.Identifier, input, 0, 0, 4}
	expectedNumTokens := 1
	actual, _ := NewWithConfig(
		input, Config{autoEOFToken: false}).ScanTokens()
	actualNumTokens := len(actual)
	if expectedNumTokens != actualNumTokens {
		t.Fatalf("test[%s] - Wrong number of tokens. Expected=%q, Actual=%q",
			"0", expectedNumTokens, actualNumTokens)
	}
	validateTestToken(t, "2", expected, actual[0])
}
