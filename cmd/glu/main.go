package main

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/templecloud/glu/pkg/repl"
)

const (
	// Repl identifier.
	Repl = "repl"
	// File switch
	File = "-f"
)

func main() {
	if len(os.Args) == 2 && os.Args[1] == Repl {
		repl.New().Start(os.Stdin, os.Stdout)
	} else if len(os.Args) == 3 && os.Args[1] == File {
		data, err := ioutil.ReadFile(os.Args[2])
		if err != nil {
			panic(err)
		}
		repl.NewCmd().Exec(string(data))
	} else {
		input := strings.Join(os.Args[1:], " ")
		repl.NewCmd().Exec(input)
	}
}
