// A simple Read-Eval-Print-Loop for Windows Remote Management.
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Sharpie/winrm-repl/pkg/repl"
)

const (
	Version = "0.0.1"
)

var (
	action = "runRepl"
)

func main() {
	setup()

	switch action {
	case "printVersion":
		fmt.Println(Version)
	case "printHelp":
		flag.Usage()
		defer os.Exit(1)
	default:
		repl := repl.NewRepl(os.Stdin, os.Stdout)

		// TODO: Right now, just a simple "echo" REPL. Need to pass a real WinRM
		// connection.
		repl.Run()
	}
}

func setup() {
	printVersion := flag.Bool("version", false, "Print version number")
	// TODO: Flags for seting WinRM auth parameters.

	flag.Parse()

	if *printVersion {
		action = "printVersion"
	} else if flag.NArg() != 1 {
		action = "printHelp"
	}
}
