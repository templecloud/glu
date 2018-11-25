package interpreter

import (
	"fmt"
	"strconv"

	"github.com/templecloud/glu/pkg/ast"
	"github.com/templecloud/glu/pkg/token"
)

// Interpreter ================================================================
//

// Interpreter is an implementation of a Visitor that evaluates the statements
// and expressions it parses.
type Interpreter struct {
	*Environment
}

// New creates a Interpeter.
func New() *Interpreter {
	return &Interpreter{
		Environment: NewEnvironment(),
	}
}

// Eval recursively traverses the specified Stmt and returns the result or the
// first runtime error it encountered.
func (i *Interpreter) Eval(stmt ast.Stmt) (result interface{}, err *Error) {
	defer func() {
		if r := recover(); r != nil {
			switch e := r.(type) {
			case *Error:
				// If evaluation error is detected panic try and recover.
				err = e
			default:
				// Else, continue generic runtime error.
				panic(e)
			}
		}
	}()
	result = i.evaluate(stmt)
	return
}

// evaluate ===================================================================
//

// evaluate recursively traverses the specified Stmt and returns a string
// representation.
func (i *Interpreter) evaluate(stmt ast.Stmt) interface{} {
	return stmt.Accept(i)
}

// VisitBinaryExpr evaluates the node.
func (i *Interpreter) VisitBinaryExpr(expr *ast.Binary) interface{} {
	left := i.evaluate(expr.Left)
	right := i.evaluate(expr.Right)

	switch expr.Operator.Type {
	// Compators
	case token.GreaterThan:
		checkNumberOperands(expr.Operator, left, right)
		return left.(float64) > right.(float64)
	case token.GreaterThanOrEqual:
		checkNumberOperands(expr.Operator, left, right)
		return left.(float64) >= right.(float64)
	case token.LessThan:
		checkNumberOperands(expr.Operator, left, right)
		return left.(float64) < right.(float64)
	case token.LessThanOrEqual:
		checkNumberOperands(expr.Operator, left, right)
		return left.(float64) <= right.(float64)
	// Equality
	case token.NotEqual:
		return !isEqual(left, right)
	case token.EqualEqual:
		return isEqual(left, right)
	// Arithmetic
	case token.Plus:
		// TODO: Handle Strings
		checkNumberOperands(expr.Operator, left, right)
		return left.(float64) + right.(float64)
	case token.Minus:
		checkNumberOperands(expr.Operator, left, right)
		return left.(float64) - right.(float64)
	case token.ForwardSlash:
		checkNumberOperands(expr.Operator, left, right)
		return left.(float64) / right.(float64)
	case token.Star:
		checkNumberOperands(expr.Operator, left, right)
		return left.(float64) * right.(float64)
	}
	// Unreachable.
	return nil
}

// VisitGroupingExpr evaluates the node.
func (i *Interpreter) VisitGroupingExpr(expr *ast.Grouping) interface{} {
	return i.evaluate(expr.Expr)
}

// VisitLiteralExpr evaluates the node and terminates recursion to return
// a literal value.
func (i *Interpreter) VisitLiteralExpr(expr *ast.Literal) interface{} {
	if expr.Value == nil {
		return "nil"
	}
	number, err := strconv.ParseFloat(expr.Value.(string), 32)
	if err != nil {
		return expr.Value
	}
	return number
}

// VisitUnaryExpr evaluates the node.
func (i *Interpreter) VisitUnaryExpr(expr *ast.Unary) interface{} {
	right := i.evaluate(expr.Right)
	switch expr.Operator.Type {
	case token.Not:
		return !isTruthy(right)
	case token.Minus:
		checkNumberOperand(expr.Operator, right)
		return -right.(float64)
	}
	return nil
}

// VisitVarExpr evaluates the node.
func (i *Interpreter) VisitVarExpr(ve *ast.VarExpr) interface{} {
	return i.Environment.Get(ve.Name)
}

// Expr Runtime Error Functions ===============================================
//

func checkNumberOperand(operator *token.Token, operand interface{}) {
	switch operand.(type) {
	case float64:
		return
	default:
		panic(NewError(operator, "Operand must be a number."))
	}
}

func checkNumberOperands(
	operator *token.Token, leftOperand, rightOperand interface{}) {
	if _, ok := leftOperand.(float64); ok {
		if _, ok := rightOperand.(float64); ok {
			return
		}
	}

	panic(NewError(operator, "Operands must both be numbers."))
}

// Stmt Runtime Error Functions ===============================================
//

// VisitExprStmt evaluates the node.
func (i *Interpreter) VisitExprStmt(stmt *ast.ExprStmt) interface{} {
	return i.evaluate(stmt.Expr)
}

// VisitLogStmt evaluates the node.
func (i *Interpreter) VisitLogStmt(stmt *ast.LogStmt) interface{} {
	value := i.evaluate(stmt.Expr)
	fmt.Printf("%s", stringify(value))
	return nil
}

// VisitVariableStmt evaluates the node. See also VisitVarExpr.
func (i *Interpreter) VisitVariableStmt(stmt *ast.VariableStmt) interface{} {
	var value interface{}
	if stmt.Initialiser != nil {
		value = i.evaluate(stmt.Initialiser)
	}
	i.Environment.Define(stmt.Name.Lexeme, value)
	return nil
}

// Support Functions ==========================================================
//

// isTruthy defines the 'truthiness' semantics for Glu.
func isTruthy(thing interface{}) bool {
	if thing == nil {
		return false
	}
	switch thing.(type) {
	case bool:
		return thing.(bool)
	default:
		return true
	}
}

// isEqual defines the 'identity' semantics for Glu.
func isEqual(t1 interface{}, t2 interface{}) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil {
		return false
	}
	return t1 == t2
}

func stringify(value interface{}) string {
	if value == nil {
		return "nil"
	}
	switch value.(type) {
	case float64:
		return fmt.Sprintf("%v", value.(float64))
	case string:
		return value.(string)
	default:
		return fmt.Sprintf("%v", value)
	}
}
