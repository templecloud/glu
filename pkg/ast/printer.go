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
	return expr.accept(p).(string)
}

func (p *Printer) visitBinaryExpr(expr *Binary) interface{} {
	return p.parenthesize(expr.operator.Lexeme, expr.left, expr.right)
}

func (p *Printer) visitGroupingExpr(expr *Grouping) interface{} {
	return p.parenthesize("#g", expr.expr)
}

// visitLiteralExpr terminates recursion.
func (p *Printer) visitLiteralExpr(expr *Literal) interface{} {
	if expr.value == nil {
		return "nil"
	}
	return fmt.Sprintf("%+v", expr.value)
}

func (p *Printer) visitUnaryExpr(expr *Unary) interface{} {
	return p.parenthesize(expr.operator.Lexeme, expr.right)
}

// parenthesize adds grouping parenthese and recursively calls 'accept'.
func (p *Printer) parenthesize(name string, exprs ...Expr) string {
	var builder strings.Builder
	builder.WriteString("(")
	builder.WriteString(name)
	for _, expr := range exprs {
		builder.WriteString(" ")
		builder.WriteString(expr.accept(p).(string))
	}
	builder.WriteString(")")

	return builder.String()
}
