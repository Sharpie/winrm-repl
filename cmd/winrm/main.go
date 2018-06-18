// A simple Read-Eval-Print-Loop for Windows Remote Management.
package main

import (
	"flag"
	"fmt"
	"os"
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
		// TODO: Launch REPL.
		defer os.Exit(0)
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
