package parser

import (
	"fmt"

	"github.com/templecloud/glu/pkg/ast"
	"github.com/templecloud/glu/pkg/token"
)

// Expression Functions =======================================================
//

func (p *Parser) assignment() ast.Expr {
	expr := p.or()
	if p.match(token.Equal) {
		equals := p.previous()
		value := p.assignment()
		switch v := expr.(type) {
		case *ast.VarExpr:
			name := v.Name
			return ast.NewAssign(name, value)
		default:
			err := NewError(equals, "Invalid assignment target.")
			fmt.Printf("Parse Error: %+v\n", err)
		}
	}
	return expr
}

func (p *Parser) or() ast.Expr {
	expr := p.and()
	for p.match(token.Or) {
		operator := p.previous()
		right := p.and()
		expr = ast.NewLogical(expr, operator, right)
	}
	return expr
}

func (p *Parser) and() ast.Expr {
	expr := p.equality()
	for p.match(token.And) {
		operator := p.previous()
		right := p.equality()
		expr = ast.NewLogical(expr, operator, right)
	}
	return expr
}

// expression parses the root expression hierarchy.
func (p *Parser) expression() ast.Expr {
	return p.assignment()
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
		return ast.NewLiteral(p.previous().Type, false)
	}
	if p.match(token.True) {
		return ast.NewLiteral(p.previous().Type, true)
	}
	if p.match(token.Nil) {
		return ast.NewLiteral(p.previous().Type, nil)
	}
	if p.match(token.Number, token.String) {
		return ast.NewLiteral(p.previous().Type, p.previous().Lexeme)
	}
	if p.match(token.Identifier) {
		return ast.NewVarExpr(p.previous())
	}
	if p.match(token.LeftParen) {
		expr := p.expression()
		p.consume(token.RightParen, "Expected ')' after expression.")
		return ast.NewGrouping(expr)
	}
	panic(NewError(p.tokens[p.current], "Token failed to match any rule."))
}
