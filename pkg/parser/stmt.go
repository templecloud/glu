package parser

import (
	"fmt"

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
	p.consume(token.RightBrace, "Expected '}' after block.")
	return stmts
}

func (p *Parser) declaration() ast.Stmt {
	// trjl: synchronise here instead?
	if p.match(token.Func) {
		return p.fnStatement("function")
	}
	if p.match(token.Var) {
		return p.varDeclaration()
	}
	return p.statement()
}

func (p *Parser) expressionStatement() ast.Stmt {
	expr := p.expression()
	p.consume(token.Semicolon, "Expected ';' after expression.")
	return ast.NewExprStmt(expr)
}

func (p *Parser) fnStatement(kind string) ast.Stmt {
	// Consume function name.
	name := p.consume(token.Identifier, fmt.Sprintf("Expected kind %s.", kind))
	// Consume function parameters.
	p.consume(token.LeftParen, fmt.Sprintf("Expected '(' after kind %s.", kind))
	var parameters []*token.Token
	if !p.check(token.RightParen) {
		parameters = append(parameters, p.consume(token.Identifier, "Expected parameter name."))
		for p.match(token.Comma) {
			if len(parameters) >= 8 {
				err := NewError(p.peek(), "Cannot have more than 8 arguments.")
				fmt.Printf("Parse Error: %+v\n", err)
			}
			parameters = append(parameters, p.consume(token.Identifier, "Expected parameter name."))
		}
	}
	p.consume(token.RightParen, "Expected ')' after arguments.")
	// Consume function body.
	p.consume(token.LeftBrace, fmt.Sprintf("Expected '{' before kind %s body.", kind))
	body := p.blockStatement()
	return ast.NewFnStmt(name, parameters, body)
}

func (p *Parser) forStatement() ast.Stmt {
	p.consume(token.LeftParen, "Expected '(' after 'for'.")
	// initialiser
	var initializer ast.Stmt
	if p.match(token.Semicolon) {
		initializer = nil
	} else if p.match(token.Var) {
		initializer = p.varDeclaration()
	} else {
		initializer = p.expressionStatement()
	}
	// condition
	var condition ast.Expr
	if !p.check(token.Semicolon) {
		condition = p.expression()
	}
	p.consume(token.Semicolon, "Expected ';' after loop condition.")
	// incrementor
	var increment ast.Expr
	if !p.check(token.RightParen) {
		increment = p.expression()
	}
	p.consume(token.RightParen, "Expected ')' after if condition.")

	// de-sugared statement
	body := p.statement()
	if increment != nil {
		body = ast.NewBlockStmt([]ast.Stmt{body, ast.NewExprStmt(increment)})
	}
	if condition == nil {
		condition = ast.NewLiteral(token.True, true)
	}
	body = ast.NewWhileStmt(condition, body)
	if initializer != nil {
		body = ast.NewBlockStmt([]ast.Stmt{initializer, body})
	}
	return body
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
	if p.match(token.For) {
		return p.forStatement()
	}
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
