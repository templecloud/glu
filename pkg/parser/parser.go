package parser

import (
	"github.com/templecloud/glu/pkg/ast"
	"github.com/templecloud/glu/pkg/token"
)

// Parser =====================================================================
//

// Parser parses expression from a set of tokens.
type Parser struct {
	tokens  []*token.Token
	Errors  []*Error
	current int
}

// New creates a Parser from the specified set of tokens.
func New(tokens []*token.Token) *Parser {
	return &Parser{tokens: tokens, current: 0}
}

// Parse an expression from the Parser tokens.
func (p *Parser) Parse() []ast.Stmt {
	var stmts []ast.Stmt
	for !p.isAtEnd() {
		defer func() {
			if r := recover(); r != nil {
				switch err := r.(type) {
				case *Error:
					// If parser error detected panic try and recover
					p.Errors = append(p.Errors, err)
					p.synchronize()
				default:
					// Else, continue generic runtime error.
					panic(err)
				}
			}
		}()
		stmt := p.declaration()
		stmts = append(stmts, stmt)
	}
	return stmts
}
