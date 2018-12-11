package interpreter

import "time"

// GluCallable ================================================================
//

// GluCallable represents a callable bit of code.
type GluCallable interface {
	Arity() int
	Call(interpreter *Interpreter, arguments []interface{}) interface{}
}

func defineNativeFunctions() *Environment {
	native := NewGlobalEnvironment()
	native.Define("time", nowFn{})
	return native
}

// nowFn ======================================================================
//
type nowFn struct{}

func (fn nowFn) Arity() int { return 0 }
func (fn nowFn) Call(interpreter *Interpreter, arguments []interface{}) interface{} {
	return time.Now()
}
