// A simple Read-Eval-Print-Loop for Windows Remote Management.
package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"syscall"

	"golang.org/x/crypto/ssh/terminal" // For reading passwords securely.

	"github.com/Sharpie/winrm-repl/pkg/repl"
	"github.com/Sharpie/winrm-repl/pkg/winrm"
)

const (
	Version = "0.0.1"
)

var (
	action = "runRepl"
	host   = winrm.Host{}
	// FIXME: The range of characters allowed in Windows usernames is more
	//        complex than \w+.
	hostParser = regexp.MustCompile(`^(?P<username>\w+)@(?P<hostname>.+)`)
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

	flag.Usage = func() {
		fmt.Printf("Usage: %s user@hostname\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	if *printVersion {
		action = "printVersion"
	} else if flag.NArg() != 1 {
		action = "printHelp"
	} else {
		hostInfo := flag.Arg(0)

		if !parseHost(hostInfo) {
			fmt.Printf("Error: Host info \"%s\" did not match: user@hostname\n", hostInfo)
			os.Exit(1)
		}

		host.Pass = readPassword()
	}
}

func parseHost(hostStr string) bool {
	match := hostParser.FindStringSubmatch(hostStr)

	if match != nil {
		host.User = match[1]
		host.Host = match[2]

		return true
	} else {
		return false
	}
}

func readPassword() string {
	fmt.Printf("Enter password for %s: ", host.User)
	bytes, err := terminal.ReadPassword(int(syscall.Stdin))

	if err != nil {
		fmt.Println("Error: Could not read password from stdin.")
		os.Exit(1)
	}

	return string(bytes)
}
