package interpreter

import (
	"fmt"

	"github.com/templecloud/glu/pkg/token"
)

// Environment ================================================================
//

// Environment represents a scopted set of runtime variables.
type Environment struct {
	Parent *Environment
	Values map[string]interface{}
}

// NewGlobalEnvironment creates a new root map based environment.
func NewGlobalEnvironment() *Environment {
	return NewChildEnvironment(nil)
}

// NewChildEnvironment creates a new map based environment.
func NewChildEnvironment(environment *Environment) *Environment {
	values := make(map[string]interface{})
	return &Environment{Parent: environment, Values: values}
}

// Assign assigns a new value to an existing variable in the environment.
func (env *Environment) Assign(name *token.Token, value interface{}) {
	if _, ok := env.Values[name.Lexeme]; ok {
		env.Values[name.Lexeme] = value
	} else if env.Parent != nil {
		env.Parent.Assign(name, value)
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
	}
	if env.Parent != nil {
		return env.Parent.Get(name)
	}
	err := fmt.Sprintf("Undefined variable '%s'.", name.Lexeme)
	panic(NewError(name, err))
}
