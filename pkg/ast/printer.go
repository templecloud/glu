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
