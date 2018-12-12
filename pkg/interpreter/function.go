package interpreter

import (
	"time"

	"github.com/templecloud/glu/pkg/ast"
)

// GluCallable ================================================================
//

// GluCallable represents a callable bit of code.
type GluCallable interface {
	Arity() int
	Call(interpreter *Interpreter, arguments []interface{}) interface{}
}

// GluFunction ================================================================
//

// GluFn represents a function.
type GluFn struct {
	Declaration *ast.FnStmt
}

// NewGluFn represents a function.
func NewGluFn(declaration *ast.FnStmt) *GluFn {
	return &GluFn{Declaration: declaration}
}

// Arity returns the number of parameters the function has.
func (gf GluFn) Arity() int {
	return len(gf.Declaration.Params)
}

// Call / Invoke this GluFn.
func (gf GluFn) Call(
	interpreter *Interpreter,
	arguments []interface{},
) interface{} {
	environment := NewChildEnvironment(interpreter.Globals)
	for idx, argument := range arguments {
		environment.Define(gf.Declaration.Params[idx].Lexeme, argument)
	}
	interpreter.executeBlock(gf.Declaration.Body, environment)
	return nil
}

// Native Functions ===========================================================
//

func defineNativeFunctions() *Environment {
	native := NewGlobalEnvironment()
	native.Define("time", nowFn{})
	return native
}

// nowFn ------------------------------
//
type nowFn struct{}

func (fn nowFn) Arity() int { return 0 }
func (fn nowFn) Call(interpreter *Interpreter, arguments []interface{}) interface{} {
	return time.Now()
}
