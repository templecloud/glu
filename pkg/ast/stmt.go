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

// LogStmt ==================================================================
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
