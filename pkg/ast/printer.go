package ast

import (
	"fmt"
	"strings"

	"github.com/templecloud/glu/pkg/token"
)

// Printer ====================================================================
//

// Printer is an implementation of a Visitor that produces a string
// representation of the specified Expr.
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

// VisitCallExpr returns a string representation of the node.
func (p *Printer) VisitCallExpr(expr *Call) interface{} {
	var builder strings.Builder
	builder.WriteString("(")
	builder.WriteString("#call-expr")
	builder.WriteString(" ")
	builder.WriteString(expr.Callee.Accept(p).(string))
	builder.WriteString("(")
	for idx, a := range expr.Arguments {
		builder.WriteString(a.Accept(p).(string))
		if idx < len(expr.Arguments)-1 {
			builder.WriteString(", ")
		}
	}
	builder.WriteString(")")
	builder.WriteString(")")
	return builder.String()
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

// VisitLogicalExpr returns a string representation of the node.
func (p *Printer) VisitLogicalExpr(expr *Logical) interface{} {
	return p.parenthesize(expr.Operator.Lexeme, expr.Left, expr.Right)
}

// VisitReturnExpr returns a string representation of the node.
func (p *Printer) VisitReturnExpr(expr *Return) interface{} {
	return p.parenthesize(expr.Keyword.Lexeme, expr.Value)
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
func (p *Printer) VisitBlockStmt(stmt *BlockStmt) interface{} {
	var builder strings.Builder
	builder.WriteString("(")
	builder.WriteString("#bs")
	for _, s := range stmt.Stmts {
		builder.WriteString(" ")
		builder.WriteString(s.Accept(p).(string))
	}
	builder.WriteString(")")
	return builder.String()
}

// VisitExprStmt returns a string representation of the node.
func (p *Printer) VisitExprStmt(stmt *ExprStmt) interface{} {
	return p.parenthesize("#es", stmt.Expr)
}

// VisitFnStmt returns a string representation of the node.
func (p *Printer) VisitFnStmt(fn *FnStmt) interface{} {
	var builder strings.Builder
	builder.WriteString("(")
	builder.WriteString("#fn-stmt")
	builder.WriteString(" ")
	builder.WriteString(fn.Name.Lexeme)
	builder.WriteString("(")
	for idx, param := range fn.Params {
		builder.WriteString(param.Lexeme)
		if idx < len(fn.Params)-1 {
			builder.WriteString(", ")
		}
	}
	builder.WriteString(")")
	builder.WriteString(" { ")
	for idx, stmt := range fn.Body {
		builder.WriteString(stmt.Accept(p).(string))
		if idx < len(fn.Body)-1 {
			builder.WriteString("; ")
		}
	}
	builder.WriteString(" }")
	builder.WriteString(")")
	return builder.String()
}

// VisitIfStmt returns a string representation of the node.
func (p *Printer) VisitIfStmt(stmt *IfStmt) interface{} {
	var builder strings.Builder
	builder.WriteString("(")
	builder.WriteString("#is")
	builder.WriteString(" ")
	builder.WriteString(stmt.Condition.Accept(p).(string))
	builder.WriteString(" ")
	builder.WriteString(stmt.ThenBranch.Accept(p).(string))
	if stmt.ElseBranch != nil {
		builder.WriteString(" ")
		builder.WriteString(stmt.ElseBranch.Accept(p).(string))
	}
	builder.WriteString(")")
	return builder.String()
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

// VisitWhileStmt returns a string representation of the node.
func (p *Printer) VisitWhileStmt(stmt *WhileStmt) interface{} {
	var builder strings.Builder
	builder.WriteString("(")
	builder.WriteString("#ws")
	builder.WriteString(" ")
	builder.WriteString(stmt.Condition.Accept(p).(string))
	builder.WriteString(" ")
	builder.WriteString(stmt.Body.Accept(p).(string))
	builder.WriteString(")")
	return builder.String()
}

// Support Functions ==========================================================
//

// parenthesize adds grouping parenthese and recursively calls 'accept'.
func (p *Printer) parenthesize(name string, exprs ...Expr) string {
	var builder strings.Builder
	builder.WriteString("(")
	builder.WriteString(name)
	for _, expr := range exprs {
		if expr != nil {
			builder.WriteString(" ")
			builder.WriteString(expr.Accept(p).(string))
		}
	}
	builder.WriteString(")")
	return builder.String()
}
