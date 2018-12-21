package analysis

import (
	"github.com/templecloud/glu/pkg/ast"
	"github.com/templecloud/glu/pkg/interpreter"
	"github.com/templecloud/glu/pkg/parser"
	"github.com/templecloud/glu/pkg/token"
)

// TODO: Make Enum
const (
	Function = "Function"
	None     = "None"
)

type Resolver struct {
	Interpreter     *interpreter.Interpreter
	scopes          []map[string]bool
	currentFuncType string
}

func NewResolver(interpreter *interpreter.Interpreter) *Resolver {
	return &Resolver{
		Interpreter:     interpreter,
		scopes:          make([]map[string]bool, 1),
		currentFuncType: None,
	}
}

// trjl
func (r *Resolver) Resolve(stmts []ast.Stmt) {
	r.resolveStmts(stmts)
}

// Block
// VisitBlockStmt statically resolves variable usage in the node if required.
func (r *Resolver) VisitBlockStmt(bs *ast.BlockStmt) interface{} {
	r.beginScope()
	r.resolveStmts(bs.Stmts)
	r.endScope()
	return nil
}

func (r *Resolver) resolveStmts(stmts []ast.Stmt) {
	for _, stmt := range stmts {
		r.resolveStmt(stmt)
	}
}

func (r *Resolver) resolveStmt(stmt ast.Stmt) {
	stmt.Accept(r)
}

func (r *Resolver) beginScope() {
	r.scopes = append(r.scopes, make(map[string]bool))
}

func (r *Resolver) endScope() {
	r.scopes = r.scopes[:len(r.scopes)-1]
}

// VariableStmt

// VisitVariableStmt statically resolves variable usage in the node if required.
func (r *Resolver) VisitVariableStmt(vs *ast.VariableStmt) interface{} {
	r.declare(vs.Name)
	if vs.Initialiser != nil {
		r.resolveExpr(vs.Initialiser)
	}
	r.define(vs.Name)
	return nil
}

func (r *Resolver) resolveExpr(expr ast.Expr) {
	expr.Accept(r)
}

func (r *Resolver) declare(name *token.Token) {
	if len(r.scopes) == 0 {
		return
	}
	scope := r.scopes[len(r.scopes)-1]

	if _, ok := scope[name.Lexeme]; ok {
		panic(parser.NewError(name, "Variable with this name already declared in this scope."))
	}

	scope[name.Lexeme] = false
}

func (r *Resolver) define(name *token.Token) {
	if len(r.scopes) == 0 {
		return
	}
	scope := r.scopes[len(r.scopes)-1]
	scope[name.Lexeme] = true
}

// ---- VarExpr

// VisitVarExpr statically resolves variable usage in the node if required.
func (r *Resolver) VisitVarExpr(expr *ast.VarExpr) interface{} {
	if len(r.scopes) == 0 && r.scopes[len(r.scopes)-1][expr.Name.Lexeme] == false {
		panic(parser.NewError(expr.Name, "Cannot read local variable in its own initializer."))
	}
	r.resolveLocal(expr, expr.Name)
	return nil
}

func (r *Resolver) resolveLocal(expr ast.Expr, name *token.Token) {
	for i := len(r.scopes) - 1; i >= 0; i = i - 1 {
		if _, ok := r.scopes[i][name.Lexeme]; ok {
			r.Interpreter.Resolve(expr, len(r.scopes)-1-i)
			return
		}
	}
	// Not found . Assume global.
	return
}

// VisitAssignExpr statically resolves variable usage in the node if required.
func (r *Resolver) VisitAssignExpr(expr *ast.Assign) interface{} {
	r.resolveExpr(expr.Value)
	r.resolveLocal(expr, expr.Name)
	return nil
}

// VisitFnStmt statically resolves variable usage in the node if required.
func (r *Resolver) VisitFnStmt(stmt *ast.FnStmt) interface{} {
	r.declare(stmt.Name)
	r.define(stmt.Name)
	// trjl
	// r.resolveFuncStmt(stmt)
	r.resolveFuncStmt(stmt, Function)

	return nil
}

// VisitUnaryExpr statically resolves variable usage in the node if required.
func (r *Resolver) resolveFuncStmt(stmt *ast.FnStmt, funcType string) interface{} {

	// trjl
	enclosingFuncType := r.currentFuncType
	r.currentFuncType = funcType

	r.beginScope()
	for _, param := range stmt.Params {
		r.declare(param)
		r.define(param)
	}
	r.resolveStmts(stmt.Body)
	r.endScope()

	// trjl
	r.currentFuncType = enclosingFuncType

	return nil
}

// ----------- other nodes - stmts

// VisitExprStmt statically resolves variable usage in the node if required.
func (r *Resolver) VisitExprStmt(expr *ast.ExprStmt) interface{} {
	r.resolveStmt(expr.Expr)
	return nil
}

// VisitIfStmt statically resolves variable usage in the node if required.
func (r *Resolver) VisitIfStmt(stmt *ast.IfStmt) interface{} {
	r.resolveStmt(stmt.Condition)
	r.resolveStmt(stmt.ThenBranch)
	if stmt.ElseBranch != nil {
		r.resolveStmt(stmt.ElseBranch)
	}
	return nil
}

// VisitLogStmt statically resolves variable usage in the node if required.
func (r *Resolver) VisitLogStmt(stmt *ast.LogStmt) interface{} {
	r.resolveStmt(stmt.Expr)
	return nil
}

// VisitReturnExpr statically resolves variable usage in the node if required.
func (r *Resolver) VisitReturnExpr(expr *ast.Return) interface{} {	
	// trjl
	if r.currentFuncType == None {
		panic(parser.NewError(expr.Keyword, "Cannot return from top-level code."))
	}

	if expr.Value != nil {
		r.resolveStmt(expr.Value)
	}
	return nil
}

// VisitWhileStmt statically resolves variable usage in the node if required.
func (r *Resolver) VisitWhileStmt(stmt *ast.WhileStmt) interface{} {
	r.resolveStmt(stmt.Condition)
	r.resolveStmt(stmt.Body)
	return nil
}

// ----------- other nodes - expr

// VisitBinaryExpr statically resolves variable usage in the node if required.
func (r *Resolver) VisitBinaryExpr(expr *ast.Binary) interface{} {
	r.resolveExpr(expr.Left)
	r.resolveExpr(expr.Right)
	return nil
}

// VisitCallExpr statically resolves variable usage in the node if required.
func (r *Resolver) VisitCallExpr(expr *ast.Call) interface{} {
	r.resolveExpr(expr.Callee)
	for _, argument := range expr.Arguments {
		r.resolveExpr(argument)
	}
	return nil
}

// VisitGroupingExpr statically resolves variable usage in the node if required.
func (r *Resolver) VisitGroupingExpr(expr *ast.Grouping) interface{} {
	r.resolveExpr(expr.Expr)
	return nil
}

// VisitLiteralExpr statically resolves variable usage in the node if required.
func (r *Resolver) VisitLiteralExpr(expr *ast.Literal) interface{} {
	return nil
}

// VisitLogicalExpr statically resolves variable usage in the node if required.
func (r *Resolver) VisitLogicalExpr(expr *ast.Logical) interface{} {
	r.resolveExpr(expr.Left)
	r.resolveExpr(expr.Right)
	return nil
}

// VisitUnaryExpr statically resolves variable usage in the node if required.
func (r *Resolver) VisitUnaryExpr(expr *ast.Unary) interface{} {
	r.resolveExpr(expr.Right)
	return nil
}
