package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/templecloud/glu/pkg/lexer"
)

// Prompt is the REPL prompt.
const Prompt = "glu> "

// Version is the current semantic version.
const Version = "0.0.1"

// Start begins a new REPL session.
func Start(in io.Reader, out io.Writer) {
	fmt.Printf("Glu v%s", Version)
	fmt.Println("Type 'exit' to exit.")

	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf(Prompt)
		ok := scanner.Scan()
		if !ok {
			return
		}

		input := scanner.Text()
		if input == "exit" {
			return
		}

		l := lexer.New(input)
		tokens, errors := l.ScanTokens()
		for idx, token := range tokens {
			fmt.Printf("token[%d]: %+v\n", idx, token)
		}
		for idx, err := range errors {
			fmt.Printf("err[%d]: %+v\n", idx, err)
		}
	}
}
