package parser

import (
	"fmt"
	"strings"

	"github.com/templecloud/glu/pkg/token"
)

type error interface {
	Error() string
}

// Error represents an error encounters during parsing.
type Error struct {
	token   *token.Token
	message string
}

// NewError create an parse error.
func NewError(token *token.Token, message string) *Error {
	return &Error{token: token, message: message}
}

func (e Error) Error() string {
	var builder strings.Builder

	if e.token.Source.Origin != "" {
		builder.WriteString(e.token.Source.Origin)
		builder.WriteString(" ")
	}

	builder.WriteString(fmt.Sprintf(e.message))
	builder.WriteString(" ")

	loc := fmt.Sprintf("At Line: %d, Column: %d", e.token.Source.Line+1, e.token.Source.Column+1)
	builder.WriteString(loc)
	builder.WriteString(", ")

	lex := fmt.Sprintf("Token: {%s: '%s'}.", e.token.Type, e.token.Lexeme)
	builder.WriteString(lex)
	builder.WriteString(" ")
	
	return builder.String()
}
