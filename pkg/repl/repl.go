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

const (
	// Prompt is the REPL prompt.
	Prompt = "glu> "
	// Version is the current semantic version.
	version  = "0.0.1"
	exit     = "exit"
	debugOn  = "debug.on"
	debugOff = "debug.off"
)

// Repl ===================================================================
//

// Repl is a 'Read Evaluate Print loop' for Glu statements.
type Repl struct {
	config
}

// New creates a new default Repl.
func New() *Repl {
	return &Repl{config: defaultConfig()}
}

// Start begins a new REPL session.
func (r *Repl) Start(in io.Reader, out io.Writer) {
	fmt.Printf("Glu %s\n", version)
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
		if input == exit {
			return
		}
		if input == debugOn {
			r.config.debug = fullDebug()
			continue
		}
		if input == debugOff {
			r.config.debug = defaultDebug()
			continue
		}
		r.Exec(input)
	}
}

// Exec tokenizes, parses, and, executes the specified input string.
func (r *Repl) Exec(input string) {
	// Lexer
	l := lexer.New(input)
	tokens, errors := l.ScanTokens()
	for idx, token := range tokens {
		if r.config.tokenHeader {
			fmt.Printf("token[%d]: ", idx)
		}
		if r.config.token {
			fmt.Printf("%+v\n", token)
		}
	}
	for idx, tokenErr := range errors {
		if r.config.tokenErrHeader {
			fmt.Printf("t_err[%d]: ", idx)
		}
		if r.config.tokenErr {
			fmt.Printf("%+v\n", tokenErr)
		}
	}

	// Parser
	p := parser.New(tokens)
	stmts := p.Parse()

	for _, stmt := range stmts {
		if len(p.Errors) > 0 {
			for idx, parserErr := range p.Errors {
				if r.config.parseErrHeader {
					fmt.Printf("p_err[%d]: ", idx)
				}
				if r.config.parseErr {
					fmt.Printf("%+v\n", parserErr)
				}
			}
		} else {
			printer := ast.Printer{}
			exprStr := printer.Print(stmt)

			if r.config.exprHeader {
				fmt.Printf("expr   :")
			}
			if r.config.expr {
				fmt.Printf("%s\n", exprStr)
			}

			// Evaluate
			i := interpreter.Interpreter{}
			result := i.Evaluate(stmt)
			if len(i.Errors) > 0 {
				for idx, evalErr := range i.Errors {
					if r.config.evalErrHeader {
						fmt.Printf("i_err[%d] :", idx)
					}
					if r.config.evalErr {
						fmt.Printf("%v\n", evalErr)
					}
				}
			} else {
				if r.config.resultHeader {
					fmt.Printf("result : ")
				}
				if r.config.result {
					fmt.Printf("%v\n", result)
				}
			}
		}
	}
}
