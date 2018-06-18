// Simple Read-Eval-Print-Loop built on bufio from the Go stdlib.
package repl

import (
	"bufio"
	"fmt"
	"io"
)

type Repl struct {
	scanner *bufio.Scanner
	output  io.Writer
}

const (
	exitKwd = "exit"
)

func NewRepl(input io.Reader, output io.Writer) *Repl {
	return &Repl{
		scanner: bufio.NewScanner(input),
		output:  output,
	}
}

func (r *Repl) Run() {
	for {
		r.prompt()

		if r.read() {
			r.eval()
		} else {
			// TODO: Also shut down WinRM connection.
			break
		}
	}
}

func (r *Repl) prompt() {
	fmt.Fprint(r.output, "PS > ")
}

func (r *Repl) read() bool {
	if r.scanner.Scan() {
		// TODO: Also handle signals such as Ctrl-C.
		return r.scanner.Text() != exitKwd
	} else {
		return false
	}
}

func (r *Repl) eval() {
	// TODO: Just a simple "echo" action. Need to fit a real WinRM connection in.
	fmt.Fprintln(r.output, r.scanner.Text())
}
