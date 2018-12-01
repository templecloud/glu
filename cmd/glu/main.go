package main

import (
	"os"
	"strings"

	"github.com/templecloud/glu/pkg/repl"
)

const (
	// Repl identifier.
	Repl = "repl"
)

func main() {
	if len(os.Args) == 2 && os.Args[1] == Repl {
		repl.New().Start(os.Stdin, os.Stdout)
	} else {
		exprStr := strings.Join(os.Args[1:], " ")
		repl.NewCmd().Exec(exprStr)
	}
}
