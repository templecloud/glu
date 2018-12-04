package parser

import (
	"github.com/templecloud/glu/pkg/ast"
	"github.com/templecloud/glu/pkg/token"
)

// Statement Functions ========================================================
//

func (p *Parser) blockStatement() []ast.Stmt {
	var stmts []ast.Stmt
	for !p.check(token.RightBrace) && !p.isAtEnd() {
		stmts = append(stmts, p.declaration())
	}
	p.consume(token.RightBrace, "Expect '}' after block.")
	return stmts
}

func (p *Parser) declaration() ast.Stmt {
	// trjl: synchronise here instead?
	if p.match(token.Var) {
		return p.varDeclaration()
	}
	return p.statement()
}

func (p *Parser) expressionStatement() ast.Stmt {
	expr := p.expression()
	p.consume(token.Semicolon, "Expect ';' after expression.")
	return ast.NewExprStmt(expr)
}

func (p *Parser) ifStatement() ast.Stmt {
	p.consume(token.LeftParen, "Expect '(' after if condition.")
	condition := p.expression()
	p.consume(token.RightParen, "Expect ')' after if condition.")
	thenBranch := p.statement()
	var elseBranch ast.Stmt
	if p.match(token.Else) {
		elseBranch = p.statement()
	}
	return ast.NewIfStmt(condition, thenBranch, elseBranch)
}

func (p *Parser) printStatement() ast.Stmt {
	value := p.expression()
	p.consume(token.Semicolon, "Expect ';' after value.")
	return ast.NewLogStmt(value)
}

func (p *Parser) statement() ast.Stmt {
	if p.match(token.If) {
		return p.ifStatement()
	}
	if p.match(token.Log) {
		return p.printStatement()
	}
	if p.match(token.While) {
		return p.whileStatement()
	}
	if p.match(token.LeftBrace) {
		return ast.NewBlockStmt(p.blockStatement())
	}
	return p.expressionStatement()
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

func (p *Parser) whileStatement() ast.Stmt {
	p.consume(token.LeftParen, "Expected '(' after while.")
	condition := p.expression()
	p.consume(token.RightParen, "Expected ')' after while.")
	body := p.statement()
	return ast.NewWhileStmt(condition, body)
}
