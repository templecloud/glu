package ast

import (
	"github.com/templecloud/glu/pkg/token"
)

// Expr =======================================================================
//

// Expr is (currently) the root abstract AST node type.
type Expr interface {
	Accept(visitor Visitor) interface{}
}

// Assign =====================================================================
//

// Assign expression node.
type Assign struct {
	Name  *token.Token
	Value Expr
}

// NewAssign constructor.
func NewAssign(name *token.Token, value Expr) *Assign {
	return &Assign{Name: name, Value: value}
}

// Accept a Vistor that can perform an operation on the node to return a result.
func (a *Assign) Accept(visitor Visitor) interface{} {
	return visitor.VisitAssignExpr(a)
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

// Call =======================================================================
//

// Call expression node.
type Call struct {
	Callee    Expr
	Paren     *token.Token
	Arguments []Expr
}

// NewCall constructor.
func NewCall(callee Expr, paren *token.Token, arguments []Expr) *Call {
	return &Call{Callee: callee, Paren: paren, Arguments: arguments}
}

// Accept a Vistor that can perform an operation on the node to return a result.
func (c *Call) Accept(visitor Visitor) interface{} {
	return visitor.VisitCallExpr(c)
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

// Logical ====================================================================
//

// Logical expression node.
type Logical struct {
	Left     Expr
	Operator *token.Token
	Right    Expr
}

// NewLogical constructor.
func NewLogical(left Expr, operator *token.Token, right Expr) *Logical {
	return &Logical{Left: left, Operator: operator, Right: right}
}

// Accept a Vistor that can perform an operation on the node to return a result.
func (l *Logical) Accept(visitor Visitor) interface{} {
	return visitor.VisitLogicalExpr(l)
}

// Return =====================================================================
//

// Return expression node.
type Return struct {
	Keyword *token.Token
	Value Expr	
}

// NewReturn constructor.
func NewReturn(keyword *token.Token, value Expr) *Return {
	return &Return{Keyword: keyword, Value: value}
}

// Accept a Vistor that can perform an operation on the node to return a result.
func (r *Return) Accept(visitor Visitor) interface{} {
	return visitor.VisitReturnExpr(r)
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
