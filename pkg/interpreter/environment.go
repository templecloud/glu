package interpreter

import (
	"fmt"

	"github.com/templecloud/glu/pkg/token"
)

// Environment ================================================================
//

// Environment represents a scopted set of runtime variables.
type Environment struct {
	Values map[string]interface{}
}

// NewEnvironment creates a new map based environment.
func NewEnvironment() *Environment {
	values := make(map[string]interface{})
	return &Environment{Values: values}
}

// Assign assigns a new value to an existing variable in the environment.
func (env *Environment) Assign(name *token.Token, value interface{}) {
	if _, ok := env.Values[name.Lexeme]; ok {
		env.Values[name.Lexeme] = value
	} else {
		err := fmt.Sprintf("Undefined variable '%s'.", name.Lexeme)
		panic(NewError(name, err))
	}
}

// Define adds a new variable to the environment.
func (env *Environment) Define(name string, value interface{}) {
	env.Values[name] = value
}

// Get attempts to retrieve the specified environment variable.
func (env *Environment) Get(name *token.Token) interface{} {
	if lexeme, ok := env.Values[name.Lexeme]; ok {
		return lexeme
	} else {
		err := fmt.Sprintf("Undefined variable '%s'.", name.Lexeme)
		panic(NewError(name, err))
	}
}
