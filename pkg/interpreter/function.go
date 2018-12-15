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
	Closure     *Environment
}

// NewGluFn represents a function.
func NewGluFn(declaration *ast.FnStmt, environment *Environment) *GluFn {
	return &GluFn{Declaration: declaration, Closure: environment}
}

// Arity returns the number of parameters the function has.
func (gf GluFn) Arity() int {
	return len(gf.Declaration.Params)
}

// Call / Invoke this GluFn.
func (gf GluFn) Call(
	interpreter *Interpreter,
	arguments []interface{},
) (result interface{}) {
	// Define a new function environment and set the parameters.
	environment := NewChildEnvironment(gf.Closure)
	for idx, argument := range arguments {
		environment.Define(gf.Declaration.Params[idx].Lexeme, argument)
	}
	// Set-up a defferred function to handle the dodgy panic based function
	// return.
	defer func() {
		if r := recover(); r != nil {
			switch res := r.(type) {
			case *Return:
				// Dodgy! Catch *Return type structs and return the value.
				result = res.value
			case *Error:
				panic(res)
			default:
				panic(res)
			}
		}
	}()
	// Execute the function block.
	interpreter.executeBlock(gf.Declaration.Body, environment)
	return
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
