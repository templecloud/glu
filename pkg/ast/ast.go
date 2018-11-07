package ast

import (
	"github.com/templecloud/glu/pkg/token"
)

// Visitor =====================================================================
//

// Visitor represents the GoF 'visitor' design pattern. Ideally, it would be
// parameterized, but, we will need to wait for golang generic types for that!
type Visitor interface {
	visitBinaryExpr(b *Binary) interface{}
	visitGroupingExpr(g *Grouping) interface{}
	visitLiteralExpr(l *Literal) interface{}
	visitUnaryExpr(u *Unary) interface{}
}

// Expr =======================================================================
//

// Expr is (currently) the root abstract AST node type.
type Expr interface {
	accept(visitor Visitor) interface{}
}

// Binary =====================================================================
//

// Binary expression node.
type Binary struct {
	left     Expr
	operator token.Token
	right    Expr
}

// NewBinary constructor.
func NewBinary(left Expr, operator token.Token, right Expr) *Binary {
	return &Binary{left: left, operator: operator, right: right}
}

func (b *Binary) accept(visitor Visitor) interface{} {
	return visitor.visitBinaryExpr(b)
}

// Grouping ===================================================================
//

// Grouping expression node.
type Grouping struct {
	expr Expr
}

// NewGrouping constructor.
func NewGrouping(expr Expr) *Grouping {
	return &Grouping{expr: expr}
}

func (g *Grouping) accept(visitor Visitor) interface{} {
	return visitor.visitGroupingExpr(g)
}

// Literal ====================================================================
//

// Literal expression node.
type Literal struct {
	value interface{}
}

// NewLiteral constructor.
func NewLiteral(value interface{}) *Literal {
	return &Literal{value: value}
}

func (l *Literal) accept(visitor Visitor) interface{} {
	return visitor.visitLiteralExpr(l)
}

// Unary ======================================================================
//

// Unary expression node.
type Unary struct {
	operator token.Token
	right    Expr
}

// NewUnary constructor.
func NewUnary(operator token.Token, right Expr) *Unary {
	return &Unary{operator: operator, right: right}
}

func (u *Unary) accept(visitor Visitor) interface{} {
	return visitor.visitUnaryExpr(u)
}
