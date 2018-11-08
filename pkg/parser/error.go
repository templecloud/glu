package parser

import (
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
