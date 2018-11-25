package ast

import (
	"github.com/templecloud/glu/pkg/token"
)

// Visitor =====================================================================
//

// Visitor represents the GoF 'visitor' design pattern. Ideally, it would be
// parameterized, but, we will need to wait for golang generic types for that!
type Visitor interface {
	// expressions
	VisitBinaryExpr(b *Binary) interface{}
	VisitGroupingExpr(g *Grouping) interface{}
	VisitLiteralExpr(l *Literal) interface{}
	VisitUnaryExpr(u *Unary) interface{}
	VisitVarExpr(ve *VarExpr) interface{}
	// statements
	VisitLogStmt(ps *LogStmt) interface{}
	VisitExprStmt(es *ExprStmt) interface{}
	VisitVariableStmt(vs *VariableStmt) interface{}
}

// Expr =======================================================================
//

// Expr is (currently) the root abstract AST node type.
type Expr interface {
	Accept(visitor Visitor) interface{}
}

// Binary =====================================================================
//

// Binary expression node.
type Binary struct {
	Left     Expr
	Operator *token.Token
	Right    Expr
}

// NewBinary constructor.
func NewBinary(left Expr, operator *token.Token, right Expr) *Binary {
	return &Binary{Left: left, Operator: operator, Right: right}
}

// Accept a Vistor that can perform an operation on the node to return a result.
func (b *Binary) Accept(visitor Visitor) interface{} {
	return visitor.VisitBinaryExpr(b)
}

// Grouping ===================================================================
//

// Grouping expression node.
type Grouping struct {
	Expr
}

// NewGrouping constructor.
func NewGrouping(expr Expr) *Grouping {
	return &Grouping{Expr: expr}
}

// Accept a Vistor that can perform an operation on the node to return a result.
func (g *Grouping) Accept(visitor Visitor) interface{} {
	return visitor.VisitGroupingExpr(g)
}

// Literal ====================================================================
//

// Literal expression node.
type Literal struct {
	TokenType token.Type
	Value     interface{}
}

// NewLiteral constructor.
func NewLiteral(tokenType token.Type, value interface{}) *Literal {
	return &Literal{TokenType: tokenType, Value: value}
}

// Accept a Vistor that can perform an operation on the node to return a result.
func (l *Literal) Accept(visitor Visitor) interface{} {
	return visitor.VisitLiteralExpr(l)
}

// Unary ======================================================================
//

// Unary expression node.
type Unary struct {
	Operator *token.Token
	Right    Expr
}

// NewUnary constructor.
func NewUnary(operator *token.Token, right Expr) *Unary {
	return &Unary{Operator: operator, Right: right}
}

// Accept a Vistor that can perform an operation on the node to return a result.
func (u *Unary) Accept(visitor Visitor) interface{} {
	return visitor.VisitUnaryExpr(u)
}

// VarExpr ======================================================================
//

// VarExpr expression node.
type VarExpr struct {
	Name *token.Token
}

// NewVarExpr constructor.
func NewVarExpr(name *token.Token) *VarExpr {
	return &VarExpr{Name: name}
}

// Accept a Vistor that can perform an operation on the node to return a result.
func (ve *VarExpr) Accept(visitor Visitor) interface{} {
	return visitor.VisitVarExpr(ve)
}
