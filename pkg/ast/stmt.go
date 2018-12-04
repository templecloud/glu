package ast

import "github.com/templecloud/glu/pkg/token"

// Stmt =======================================================================
//

// Stmt represents a statement node in the AST tree.
type Stmt interface {
	Accept(visitor Visitor) interface{}
}

// BlockStmt ==================================================================
//

// BlockStmt statement node.
type BlockStmt struct {
	Stmts []Stmt
}

// NewBlockStmt constructor.
func NewBlockStmt(stmts []Stmt) *BlockStmt {
	return &BlockStmt{Stmts: stmts}
}

// Accept a Vistor that can perform an operation on the node to return a result.
func (bs *BlockStmt) Accept(visitor Visitor) interface{} {
	return visitor.VisitBlockStmt(bs)
}

// ExprStmt ===================================================================
//

// ExprStmt statement node.
type ExprStmt struct {
	Expr
}

// NewExprStmt constructor.
func NewExprStmt(expr Expr) *ExprStmt {
	return &ExprStmt{Expr: expr}
}

// Accept a Vistor that can perform an operation on the node to return a result.
func (es *ExprStmt) Accept(visitor Visitor) interface{} {
	return visitor.VisitExprStmt(es)
}

// IfStmt ===================================================================
//

// IfStmt statement node.
type IfStmt struct {
	Condition  Expr
	ThenBranch Stmt
	ElseBranch Stmt
}

// NewIfStmt constructor.
func NewIfStmt(condition Expr, thenBranch Stmt, elseBranch Stmt) *IfStmt {
	return &IfStmt{Condition: condition, ThenBranch: thenBranch, ElseBranch: elseBranch}
}

// Accept a Vistor that can perform an operation on the node to return a result.
func (is *IfStmt) Accept(visitor Visitor) interface{} {
	return visitor.VisitIfStmt(is)
}

// LogStmt ====================================================================
//

// LogStmt statement node.
type LogStmt struct {
	Expr
}

// NewLogStmt constructor.
func NewLogStmt(expr Expr) *LogStmt {
	return &LogStmt{Expr: expr}
}

// Accept a Vistor that can perform an operation on the node to return a result.
func (ps *LogStmt) Accept(visitor Visitor) interface{} {
	return visitor.VisitLogStmt(ps)
}

// VariableStmt ====================================================================
//

// VariableStmt statement node.
type VariableStmt struct {
	Name        *token.Token
	Initialiser Expr
}

// NewVariableStmt constructor.
func NewVariableStmt(name *token.Token, initialiser Expr) *VariableStmt {
	return &VariableStmt{Name: name, Initialiser: initialiser}
}

// Accept a Vistor that can perform an operation on the node to return a result.
func (vs *VariableStmt) Accept(visitor Visitor) interface{} {
	return visitor.VisitVariableStmt(vs)
}

// WhileStmt ====================================================================
//

// WhileStmt statement node.
type WhileStmt struct {
	Condition Expr
	Body      Stmt
}

// NewWhileStmt constructor.
func NewWhileStmt(condition Expr, body Stmt) *WhileStmt {
	return &WhileStmt{Condition: condition, Body: body}
}

// Accept a Vistor that can perform an operation on the node to return a result.
func (vs *WhileStmt) Accept(visitor Visitor) interface{} {
	return visitor.VisitWhileStmt(vs)
}
