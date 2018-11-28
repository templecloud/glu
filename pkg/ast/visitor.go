package ast

// Visitor =====================================================================
//

// Visitor represents the GoF 'visitor' design pattern. Ideally, it would be
// parameterized, but, we will need to wait for golang generic types for that!
type Visitor interface {
	// expressions
	VisitAssignExpr(a *Assign) interface{}
	VisitBinaryExpr(b *Binary) interface{}
	VisitGroupingExpr(g *Grouping) interface{}
	VisitLiteralExpr(l *Literal) interface{}
	VisitUnaryExpr(u *Unary) interface{}
	VisitVarExpr(ve *VarExpr) interface{}
	// statements
	VisitBlockStmt(bs *BlockStmt) interface{}
	VisitLogStmt(ps *LogStmt) interface{}
	VisitExprStmt(es *ExprStmt) interface{}
	VisitVariableStmt(vs *VariableStmt) interface{}
}
