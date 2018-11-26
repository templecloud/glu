package lexer

import (
	"github.com/templecloud/glu/pkg/token"
)

// Lexer Error ================================================================
//

// Error represents a lexical error in a source file or stream.
type Error struct {
	Message string
	token.Source
}