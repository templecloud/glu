package parser

import (
	"github.com/templecloud/glu/pkg/ast"
	"github.com/templecloud/glu/pkg/token"
)

// Parser parses expression from a set of tokens.
type Parser struct {
	tokens  []*token.Token
	current int
}

// New creates a Parser from the specified set of tokens.
func New(tokens []*token.Token) *Parser {
	return &Parser{tokens: tokens, current: 0}
}

// Parse an expression from the Parser tokens.
func (p *Parser) Parse() ast.Expr {
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		err := r.(error)
	// 		fmt.Printf("Error: %+v\n", err)
	// 	}
	// }()
	return p.expression()
}

func (p *Parser) expression() ast.Expr {
	return p.equality()
}

func (p *Parser) equality() ast.Expr {
	expr := p.comparison()
	for p.match(token.EqualEqual, token.NotEqual) {
		operator := p.previous()
		right := p.comparison()
		expr = ast.NewBinary(expr, operator, right)
	}
	return expr
}

func (p *Parser) comparison() ast.Expr {
	expr := p.addition()
	for p.match(token.GreaterThan, token.GreaterThanOrEqual, token.LessThan, token.LessThanOrEqual) {
		operator := p.previous()
		right := p.addition()
		expr = ast.NewBinary(expr, operator, right)
	}
	return expr
}

func (p *Parser) addition() ast.Expr {
	expr := p.multiplication()
	for p.match(token.Minus, token.Plus) {
		operator := p.previous()
		right := p.multiplication()
		expr = ast.NewBinary(expr, operator, right)
	}
	return expr
}

func (p *Parser) multiplication() ast.Expr {
	expr := p.unary()
	for p.match(token.ForwardSlash, token.Star) {
		operator := p.previous()
		right := p.unary()
		expr = ast.NewBinary(expr, operator, right)
	}
	return expr
}

func (p *Parser) unary() ast.Expr {
	if p.match(token.Not, token.Minus) {
		operator := p.previous()
		right := p.unary()
		return ast.NewUnary(operator, right)
	}
	return p.primary()
}

func (p *Parser) primary() ast.Expr {
	if p.match(token.False) {
		return ast.NewLiteral(false)
	}
	if p.match(token.True) {
		return ast.NewLiteral(true)
	}
	if p.match(token.Nil) {
		return ast.NewLiteral(nil)
	}
	if p.match(token.Number, token.String) {
		return ast.NewLiteral(p.previous().Lexeme)
	}
	if p.match(token.LeftParen) {
		expr := p.expression()
		p.consume(token.RightParen, "Expect ')' after expression.")
		return ast.NewGrouping(expr)
	}
	// TODO: Panic or handle better?
	return nil
}

//=============================================================================

func (p *Parser) consume(tt token.Type, message string) *token.Token {
	if p.check(tt) {
		return p.advance()
	}
	return nil
	// panic(NewError(p.peek(), message))
}

// Parser Cursor Functions ====================================================

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
	return p.current == len(p.tokens)
	// TODO: Add EOF token.
	// return p.peek().Type == token.EOF
}

func (p *Parser) peek() *token.Token {
	return p.tokens[p.current]
}

func (p *Parser) previous() *token.Token {
	return p.tokens[p.current-1]
}
