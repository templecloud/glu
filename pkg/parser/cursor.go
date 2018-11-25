package parser

import (
	"github.com/templecloud/glu/pkg/token"
)

// Parser Cursor Functions ====================================================
//

func (p *Parser) consume(tt token.Type, message string) *token.Token {
	if p.check(tt) {
		return p.advance()
	}
	panic(NewError(p.peek(), message))
}

func (p *Parser) match(tts ...token.Type) bool {
	for _, tt := range tts {
		if p.check(tt) {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser) check(tt token.Type) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().Type == tt
}

func (p *Parser) advance() *token.Token {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}

func (p *Parser) isAtEnd() bool {
	return p.peek().Type == token.EOF || p.current == len(p.tokens)
}

func (p *Parser) peek() *token.Token {
	return p.tokens[p.current]
}

func (p *Parser) previous() *token.Token {
	return p.tokens[p.current-1]
}

// synchronize scans the cursor in the input stream until it finds a known
// point to start/continue parsing.
func (p *Parser) synchronize() {
	p.advance()

	for !p.isAtEnd() {
		if p.previous().Type == token.Semicolon {
			return
		}
		switch p.peek().Type {
		case token.Func:
		case token.Var:
		case token.For:
		case token.If:
		case token.While:
		case token.Log:
		case token.Return:
			return
		}
		p.advance()
	}
}
