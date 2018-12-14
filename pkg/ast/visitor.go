package ast

// Visitor =====================================================================
//

// Visitor represents the GoF 'visitor' design pattern. Ideally, it would be
// parameterized, but, we will need to wait for golang generic types for that!
type Visitor interface {
	// expressions
	VisitAssignExpr(a *Assign) interface{}
	VisitBinaryExpr(b *Binary) interface{}
	VisitCallExpr(*Call) interface{}
	VisitGroupingExpr(g *Grouping) interface{}
	VisitLiteralExpr(l *Literal) interface{}
	VisitLogicalExpr(l *Logical) interface{}
	VisitReturnExpr(r *Return) interface{}
	VisitUnaryExpr(u *Unary) interface{}
	VisitVarExpr(ve *VarExpr) interface{}
	// statements
	VisitBlockStmt(bs *BlockStmt) interface{}
	VisitExprStmt(es *ExprStmt) interface{}
	VisitIfStmt(stmt *IfStmt) interface{}
	VisitFnStmt(fs *FnStmt) interface{}
	VisitLogStmt(ps *LogStmt) interface{}
	VisitVariableStmt(vs *VariableStmt) interface{}
	VisitWhileStmt(ws *WhileStmt) interface{}
}
