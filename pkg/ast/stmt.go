package ast

// Stmt =======================================================================
//

// Stmt represents a statement node in the AST tree.
type Stmt interface {
	Accept(visitor Visitor) interface{}
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

// PrintStmt ==================================================================
//

// PrintStmt statement node.
type PrintStmt struct {
	Expr
}

// NewPrintStmt constructor.
func NewPrintStmt(expr Expr) *PrintStmt {
	return &PrintStmt{Expr: expr}
}

// Accept a Vistor that can perform an operation on the node to return a result.
func (ps *PrintStmt) Accept(visitor Visitor) interface{} {
	return visitor.VisitPrintStmt(ps)
}
