package interpreter

// GluCallable ================================================================
//

// GluCallable represents a callable bit of code.
type GluCallable interface {
    Call(interpreter *Interpreter, arguments []interface{}) interface{}
}
