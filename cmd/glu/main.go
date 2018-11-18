package main

import (
	"fmt"
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
		fmt.Println("starting repl...")
		repl.Start(os.Stdin, os.Stdout)
	} else {
		exprStr := strings.Join(os.Args[1:], " ")
		repl.Exec(exprStr)
	}
}
