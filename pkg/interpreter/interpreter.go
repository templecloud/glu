package interpreter

import (
	"strconv"

	"github.com/templecloud/glu/pkg/ast"
	"github.com/templecloud/glu/pkg/token"
)

// Interpreter ================================================================
//

// Interpreter is an implementation of a Visitor that evaluates the statements
// and expressions it parses.
type Interpreter struct {
}

// Evaluate recursively traverses the specified Expr and returns a string
// representation.
func (i *Interpreter) Evaluate(expr ast.Expr) interface{} {
	return expr.Accept(i)
}

// VisitBinaryExpr evaluates the node.
func (i *Interpreter) VisitBinaryExpr(expr *ast.Binary) interface{} {
	left := i.Evaluate(expr.Left)
	right := i.Evaluate(expr.Right)

	switch expr.Operator.Type {
	// Compators
	case token.GreaterThan:
		return left.(float64) > right.(float64)
	case token.GreaterThanOrEqual:
		return left.(float64) >= right.(float64)
	case token.LessThan:
		return left.(float64) < right.(float64)
	case token.LessThanOrEqual:
		return left.(float64) <= right.(float64)
	// Equality
	case token.NotEqual:
		return !isEqual(left, right)
	case token.EqualEqual:
		return isEqual(left, right)
	// Arithmetic
	case token.Plus:
		// TODO: Handle Strings
		return left.(float64) + right.(float64)
	case token.Minus:
		return left.(float64) - right.(float64)
	case token.ForwardSlash:
		return left.(float64) / right.(float64)
	case token.Star:
		return left.(float64) * right.(float64)
	}
	// Unreachable.
	return nil
}

// VisitGroupingExpr evaluates the node.
func (i *Interpreter) VisitGroupingExpr(expr *ast.Grouping) interface{} {
	return i.Evaluate(expr.Expr)
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
	right := i.Evaluate(expr.Right)
	switch expr.Operator.Type {
	case token.Not:
		return !isTruthy(right)
	case token.Minus:
		return -right.(float64)
	}
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
