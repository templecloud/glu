package repl

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/templecloud/glu/pkg/ast"
	"github.com/templecloud/glu/pkg/interpreter"
	"github.com/templecloud/glu/pkg/lexer"
	"github.com/templecloud/glu/pkg/parser"
)

const (
	// Prompt is the REPL prompt.
	Prompt = "glu> "
	// version is the current semantic version.
	version = "0.0.1"
	// exit is a repl command to exist the repl.
	exit = "exit"
	// debugOn is a repl command to turn on full debugging.
	debugOn = "debug on"
	// debugOff is a repl command to turn off debugging.
	debugOff = "debug off"
)

// Repl ===================================================================
//

// Repl is a 'Read Evaluate Print loop' for Glu statements.
type Repl struct {
	config
	evaluator *interpreter.Interpreter
}

// New creates a new default Repl.
func New() *Repl {
	return &Repl{
		config:    defaultConfig(),
		evaluator: interpreter.New(),
	}
}

// NewCmd creates a new command Repl.
func NewCmd() *Repl {
	return &Repl{
		config:    cmdConfig(),
		evaluator: interpreter.New(),
	}
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

	// Parse
	p := parser.New(tokens)
	stmts := p.Parse()
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
		for idx, stmt := range stmts {
			// Print
			printer := ast.Printer{}
			exprStr := printer.Print(stmt)
			if r.config.exprHeader {
				fmt.Printf("expr   :")
			}
			if r.config.expr {
				fmt.Printf("%s\n", exprStr)
			}

			// Evaluate
			i := r.evaluator
			result, evalErr := i.Eval(stmt)
			if evalErr != nil {
				if r.config.evalErrHeader {
					fmt.Printf("runtime error: ")
				}
				if r.config.evalErr {
					fmt.Printf("%v", evalErr)
				}
			} else {
				// Result
				if r.config.resultHeader {
					fmt.Printf("result : ")
				}
				if r.config.result && result != nil && idx == len(stmts)-1 {
					fmt.Printf("%v\n", result)
				} else if strings.HasPrefix(exprStr, "(#ls") {
					// NB: Small hack to recognise REPL 'log'
					//     statements.
					fmt.Println()
				}
			}
		}
	}
}
