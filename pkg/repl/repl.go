package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/templecloud/glu/pkg/ast"
	"github.com/templecloud/glu/pkg/interpreter"
	"github.com/templecloud/glu/pkg/lexer"
	"github.com/templecloud/glu/pkg/parser"
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
		// Read
		fmt.Printf(Prompt)
		ok := scanner.Scan()
		if !ok {
			return
		}
		input := scanner.Text()
		if input == "exit" {
			return
		}

		// Lexer
		l := lexer.New(input)
		tokens, errors := l.ScanTokens()
		for idx, token := range tokens {
			fmt.Printf("token[%d]: %+v\n", idx, token)
		}
		for idx, err := range errors {
			fmt.Printf("err[%d]: %+v\n", idx, err)
		}

		// Parser
		p := parser.New(tokens)
		expr := p.Parse()
		printer := ast.Printer{}
		exprStr := printer.Print(expr)
		fmt.Printf("expr   : %s\n", exprStr)

		// Evaluate
		i := interpreter.Interpreter{}
		result := i.Evaluate(expr)
		fmt.Printf("result : %v\n", result)
	}
}
