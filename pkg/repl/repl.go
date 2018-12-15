package repl

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
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
	// run is a replc command for running a file.
	run = "run"
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
		fmt.Printf("\n%s", Prompt)
		ok := scanner.Scan()
		if !ok {
			return
		}
		input := scanner.Text()
		if input == exit {
			return
		} else if input == debugOn {
			r.config.debug = fullDebug()
			continue
		} else if input == debugOff {
			r.config.debug = defaultDebug()
			continue
		} else if strings.HasPrefix(input, run) {
			fp := strings.Trim(strings.Replace(input, run, "", 1), " ")
			if fp != "" {
				data, err := ioutil.ReadFile(fp)
				if err != nil {
					fmt.Println("Failed to open file: ", err)
				}
				input = string(data)
			} else {
				fmt.Printf("'%s' requires a valid file.\n", run)
			}
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
				fmt.Printf("%+v", parserErr)
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
					fmt.Printf("Runtime error: ")
				}
				if r.config.evalErr {
					fmt.Printf("%+v", evalErr)
				}
			} else {
				// Result
				if r.config.resultHeader {
					fmt.Printf("result : ")
				}
				if r.config.result && result != nil && idx == len(stmts)-1 {
					fmt.Printf("%v", result)
				}
			}
		}
	}
}
