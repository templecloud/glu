package ast

import (
	"fmt"
	"strings"

	"github.com/templecloud/glu/pkg/token"
)

// Printer ====================================================================
//

// Printer is an implementation of a Visitor that produces a string
// representation of the specified Expr..
type Printer struct {
}

// Print recursively traverses the specified Stmt/Expr and returns a string
// representation.
func (p *Printer) Print(stmt Stmt) string {
	return stmt.Accept(p).(string)
}

// Expr Functions =============================================================
//

// VisitAssignExpr returns a string representation of the node.
func (p *Printer) VisitAssignExpr(expr *Assign) interface{} {
	nfo := fmt.Sprintf("#as %s =", expr.Name.Lexeme)
	return p.parenthesize(nfo, expr.Value)
}

// VisitBinaryExpr returns a string representation of the node.
func (p *Printer) VisitBinaryExpr(expr *Binary) interface{} {
	return p.parenthesize(expr.Operator.Lexeme, expr.Left, expr.Right)
}

// VisitGroupingExpr returns a string representation of the node.
func (p *Printer) VisitGroupingExpr(expr *Grouping) interface{} {
	return p.parenthesize("#g", expr.Expr)
}

// VisitLiteralExpr returns a string representation of the node.
// Terminates recursion.
func (p *Printer) VisitLiteralExpr(expr *Literal) interface{} {
	if expr.Value == nil {
		return "nil"
	}
	if expr.TokenType == token.String {
		return fmt.Sprintf("\"%s\"", expr.Value.(string))
	}
	return fmt.Sprintf("%+v", expr.Value)
}

// VisitUnaryExpr returns a string representation of the node.
func (p *Printer) VisitUnaryExpr(expr *Unary) interface{} {
	return p.parenthesize(expr.Operator.Lexeme, expr.Right)
}

// VisitVarExpr returns a string representation of the node.
func (p *Printer) VisitVarExpr(expr *VarExpr) interface{} {
	return fmt.Sprintf("%s", expr.Name.Lexeme)
}

// Stmt Functions =============================================================
//

// VisitBlockStmt returns a string representation of the node.
func (p *Printer) VisitBlockStmt(bs *BlockStmt) interface{} {
	var builder strings.Builder
	builder.WriteString("(")
	builder.WriteString("#bs")
	for _, stmt := range bs.Stmts {
		builder.WriteString(" ")
		builder.WriteString(stmt.Accept(p).(string))
	}
	builder.WriteString(")")
	return builder.String()
}

// VisitExprStmt returns a string representation of the node.
func (p *Printer) VisitExprStmt(stmt *ExprStmt) interface{} {
	return p.parenthesize("#es", stmt.Expr)
}

// VisitLogStmt returns a string representation of the node.
func (p *Printer) VisitLogStmt(stmt *LogStmt) interface{} {
	return p.parenthesize("#ls", stmt.Expr)
}

// VisitVariableStmt returns a string representation of the node.
func (p *Printer) VisitVariableStmt(stmt *VariableStmt) interface{} {
	if stmt.Initialiser != nil {
		nfo := fmt.Sprintf("#vs %s =", stmt.Name.Lexeme)
		return p.parenthesize(nfo, stmt.Initialiser)
	}
	return fmt.Sprintf("(#vs %s)", stmt.Name.Lexeme)
}

// Support Functions ==========================================================
//

// parenthesize adds grouping parenthese and recursively calls 'accept'.
func (p *Printer) parenthesize(name string, exprs ...Expr) string {
	var builder strings.Builder
	builder.WriteString("(")
	builder.WriteString(name)
	for _, expr := range exprs {
		builder.WriteString(" ")
		builder.WriteString(expr.Accept(p).(string))
	}
	builder.WriteString(")")
	return builder.String()
}
