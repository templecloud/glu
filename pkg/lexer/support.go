package lexer

import (
	"regexp"
	"unicode"
)

// Support Functions ===========================================================
//

var nilByte = '\000'
var newLine = '\n'
var alpha = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString
var digit = regexp.MustCompile(`^[0-9]+$`).MatchString

// Return true if the input is '_' or alphabetic; false otherwise.
func isAlpha(c rune) bool {
	return c == '_' || unicode.IsLetter(c)
}

// Return true if the input is numeric; false otherwise.
func isDigit(c rune) bool {
	return unicode.IsDigit(c)
}

// Return true if the input is alphanumeric; false otherwise.
func isAlphaNumeric(c rune) bool {
	return isDigit(c) || isAlpha(c)
}