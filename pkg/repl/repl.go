// Simple Read-Eval-Print-Loop built on bufio from the Go stdlib.
package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/Sharpie/winrm-repl/pkg/winrm"
)

type Repl struct {
	scanner *bufio.Scanner
	output  io.Writer
	shell   *winrm.Shell
}

const (
	exitKwd = "exit"
)

func NewRepl(input io.Reader, output io.Writer, shell *winrm.Shell) *Repl {
	return &Repl{
		scanner: bufio.NewScanner(input),
		output:  output,
		shell:   shell,
	}
}

func (r *Repl) Run() {
	for {
		r.prompt()

		if r.read() {
			r.eval()
		} else {
			r.shell.Close()
			break
		}
	}
}

func (r *Repl) prompt() {
	fmt.Fprint(r.output, "\nCMD > ")
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
	// FIXME: Check error value.
	result, _ := r.shell.Execute(r.scanner.Text())

	// FIXME: Maybe just pass these right through to the Execute function and
	//        have it write the results as they stream in.
	fmt.Fprintln(r.output, result.Stderr)
	fmt.Fprintln(r.output, result.Stdout)
}
