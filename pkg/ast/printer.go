package ast

import (
	"fmt"
	"strings"
)

// Printer ====================================================================
//

// Printer is an implementation of a Visitor that produces a string
// representation of the specified Expr..
type Printer struct {
}

// Expr Functions =============================================================
//

// Print recursively traverses the specified Expr and returns a string
// representation.
func (p *Printer) Print(expr Expr) string {
	return expr.Accept(p).(string)
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

// VisitLogStmt returns a string representation of the node.
func (p *Printer) VisitLogStmt(stmt *LogStmt) interface{} {
	return p.parenthesize("#ls", stmt.Expr)
}

// VisitExprStmt returns a string representation of the node.
func (p *Printer) VisitExprStmt(stmt *ExprStmt) interface{} {
	return p.parenthesize("#es", stmt.Expr)
}

// VisitVariableStmt returns a string representation of the node.
func (p *Printer) VisitVariableStmt(stmt *VariableStmt) interface{} {
	// trjl TODO
	return "VisitVariableStmt"
	// return p.parenthesize("#vs "+stmt.Name.Lexeme, stmt.Initialiser)
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
