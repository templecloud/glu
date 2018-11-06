package main

import (
	"os"

	"github.com/templecloud/glu/pkg/repl"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
