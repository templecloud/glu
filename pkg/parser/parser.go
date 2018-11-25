package parser

import (
	"github.com/templecloud/glu/pkg/ast"
	"github.com/templecloud/glu/pkg/token"
)

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

// Statement Functions ========================================================
//

func (p *Parser) declaration() ast.Stmt {
	// trjl: synchronise here instead?
	if p.match(token.Let) {
		return p.varDeclaration()
	}
	return p.statement()
}

func (p *Parser) statement() ast.Stmt {
	if p.match(token.Log) {
		return p.printStatement()
	}
	return p.expressionStatement()
}

func (p *Parser) printStatement() ast.Stmt {
	value := p.expression()
	p.consume(token.Semicolon, "Expect ';' after value.")
	return ast.NewLogStmt(value)
}

func (p *Parser) expressionStatement() ast.Stmt {
	expr := p.expression()
	p.consume(token.Semicolon, "Expect ';' after expression.")
	return ast.NewExprStmt(expr)
}

func (p *Parser) varDeclaration() ast.Stmt {
	name := p.consume(token.Identifier, "Expected variable name.")
	var initialiser ast.Expr
	if p.match(token.Equal) {
		initialiser = p.expression()
	}
	p.consume(token.Semicolon, "Expected ';' after variable declaration.")
	return ast.NewVariableStmt(name, initialiser)
}

// Expression Functions =======================================================
//

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
	if p.match(token.Identifier) {
		return ast.NewVarExpr(p.previous())
	}
	if p.match(token.LeftParen) {
		expr := p.expression()
		p.consume(token.RightParen, "Expect ')' after expression.")
		return ast.NewGrouping(expr)
	}
	panic(NewError(p.tokens[p.current], "Token failed to match any rule."))
}

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

// Parser Synchronisation Functions ===========================================

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
		case token.Let:
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
