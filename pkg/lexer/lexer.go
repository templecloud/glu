package lexer

import (
	"fmt"

	"github.com/templecloud/glu/pkg/token"
)

// Lexer ======================================================================
//

// Lexer takes an input strings and attempts to decompose it into tokens.
type Lexer struct {
	config Config
	input  []rune
	// scan state
	origin  string
	start   int
	current int
	line    int
	column  int
}

// New creates a default instance of a Lexer for the specified input string.
func New(input string) *Lexer {
	return &Lexer{
		config: defaultConfig(),
		input:  []rune(input),
		start:  0, current: 0, column: 0}
}

// NewWithConfig creates a configured instance of a Lexer for the specified
// input string.
func NewWithConfig(input string, config Config) *Lexer {
	return &Lexer{
		config: config,
		input:  []rune(input),
		start:  0, current: 0, column: 0}
}

// Lexical Methods ============================================================
//

// ScanTokens tokenises the input returns a list of Tokens and Errors.
func (l *Lexer) ScanTokens() ([]*token.Token, []*Error) {
	tokenz := []*token.Token{}
	errors := []*Error{}
	for !l.isAtEnd() {
		l.start = l.current // start of next lexeme
		t, e := l.ScanNextToken()
		if e != nil {
			errors = append(errors, e)
		}
		if t != nil {
			tokenz = append(tokenz, t)
		}
	}
	if l.config.autoEOFToken {
		tokenz = append(tokenz, l.createToken(token.EOF, ""))
	}
	return tokenz, errors
}

// ScanNextToken attempts to scan the next token from the current position in
// the input. A token is returned if successful; else and error.
func (l *Lexer) ScanNextToken() (*token.Token, *Error) {
	var t *token.Token
	var e *Error
	c := l.advance()
	lexeme := string(c)
	switch c {
	// single char tokens
	case '(':
		t = l.createToken(token.LeftParen, lexeme)
	case ')':
		t = l.createToken(token.RightParen, lexeme)
	case '{':
		t = l.createToken(token.LeftBrace, lexeme)
	case '}':
		t = l.createToken(token.RightBrace, lexeme)
	case ',':
		t = l.createToken(token.Comma, lexeme)
	case '.':
		t = l.createToken(token.Dot, lexeme)
	case ';':
		t = l.createToken(token.Semicolon, lexeme)
	case '-':
		t = l.createToken(token.Minus, lexeme)
	case '+':
		t = l.createToken(token.Plus, lexeme)
	case '*':
		t = l.createToken(token.Star, lexeme)
	// dual char tokens
	case '/':
		if l.matches('/') {
			// consume '//' comments.
			for !l.isAtEnd() && l.peek() != nilByte {
				l.advance()
			}
		} else {
			t = l.createToken(token.ForwardSlash, lexeme)
		}
	case '!':
		if l.matches('=') {
			t = l.createToken(token.NotEqual, fmt.Sprintf("%s%s", lexeme, "="))
		} else {
			t = l.createToken(token.Not, lexeme)
		}
	case '=':
		if l.matches('=') {
			t = l.createToken(token.EqualEqual, fmt.Sprintf("%s%s", lexeme, "="))
		} else {
			t = l.createToken(token.Equal, lexeme)
		}
	case '<':
		if l.matches('=') {
			t = l.createToken(token.LessThanOrEqual, fmt.Sprintf("%s%s", lexeme, "="))
		} else {
			t = l.createToken(token.LessThan, lexeme)
		}
	case '>':
		if l.matches('=') {
			t = l.createToken(token.GreaterThanOrEqual, fmt.Sprintf("%s%s", lexeme, "="))
		} else {
			t = l.createToken(token.GreaterThan, lexeme)
		}
	// whitespace
	case ' ':
		t = nil
	case '\r':
		t = nil
	case '\t':
		t = nil
	case '\n':
		l.line++
		l.column = 0
		t = nil
	// escaped whitespace - repl mode
	case '\\':
		if l.matches('n') {
			l.line++
			l.column = 0
		} else if l.matches('r') {
		} else if l.matches('t') {
		} else {
			uc := l.advance()
			e = l.createError(fmt.Sprintf("Unexpected escape character: %c.", uc))
		}
	case '"':
		t, e = l.string()
	default:
		if isDigit(c) {
			t = l.number()
		} else if isAlpha(c) {
			t = l.identifier()
		} else {
			e = l.createError(fmt.Sprintf("Unexpected escape character: %c.", c))
		}
	}

	return t, e
}

// Attempt to consume a 'string' from the character stream.
func (l *Lexer) string() (*token.Token, *Error) {
	var t *token.Token
	var e *Error
	for !l.isAtEnd() && l.peek() != '"' {
		if l.peek() != '\n' {
			l.advance()
		}
	}
	// Unterminated string.
	if l.isAtEnd() {
		e = l.createError("Unterminated string.")
	} else {
		l.advance() // consume closing '"'
		t = l.createToken(token.String, string(l.input[l.start+1:l.current-1]))
	}
	return t, e
}

// Attempt to consume a 'number' from the character stream.
func (l *Lexer) number() *token.Token {
	// consume integer component.
	c1 := l.peek()
	for !l.isAtEnd() && isDigit(c1) {
		l.advance()
		c1 = l.peek()
	}
	// consume decimal point '.'.
	c2 := l.peekNext()
	if l.peek() == '.' && isDigit(c2) {
		l.advance()
	}
	// consume fractional component.
	c3 := l.peek()
	for !l.isAtEnd() && isDigit(c3) {
		l.advance()
		c3 = l.peek()
	}
	return l.createToken(token.Number, string(l.input[l.start:l.current]))
}

// Attempt to consume an 'identifier' from the character stream.
func (l *Lexer) identifier() *token.Token {
	c1 := l.peek()
	for !l.isAtEnd() && isAlphaNumeric(c1) {
		l.advance()
		c1 = l.peek()
	}
	identifier := string(l.input[l.start:l.current])
	// NOTE: Maybe turn into hashmap?
	var tt token.Type
	switch identifier {
	// Logical
	case "nil":
		tt = token.Nil
	case "true":
		tt = token.True
	case "false":
		tt = token.False
	case "and":
		tt = token.And
	case "or":
		tt = token.Or
	// Conditional
	case "if":
		tt = token.If
	case "else":
		tt = token.Else
	case "while":
		tt = token.While
	case "for":
		tt = token.For
	case "return":
		tt = token.Return
	// Declaration
	case "var":
		tt = token.Var
	case "func":
		tt = token.Func
	// Utility
	case "log":
		tt = token.Log
	// Identifier (non-keyword)
	default:
		tt = token.Identifier
	}
	return l.createToken(tt, string(identifier))
}

// Constructor Methods =========================================================
//

func (l *Lexer) createToken(tokenType token.Type, lexeme string) *token.Token {
	ll := len(lexeme)
	return token.New(tokenType, lexeme, l.origin, l.line, l.column-ll, ll)
}

func (l *Lexer) createError(message string) *Error {
	return &Error{
		message,
		token.Source{Origin: l.origin, Line: l.line, Column: l.column}}
}

