package interpreter

// Return =====================================================================
//

// Return represents a return value encountered during evaluation.
type Return struct {
	value   interface{}
}

// NewReturn creates a Return Value.
func NewReturn(value interface{}) *Return {
	return &Return{value: value}
}
