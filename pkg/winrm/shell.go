package winrm

import (
	"bytes"
	"io"
	"sync"
	"time"

	"github.com/masterzen/winrm"
)

// TODO: Should probably be an interface.
type Shell struct {
	shell *winrm.Shell
}

type Result struct {
	Stdout   string
	Stderr   string
	ExitCode int
}

func NewShell(host Host) (*Shell, error) {
	wrappedShell := &Shell{}
	// TODO: Add support for HTTPS.
	// TODO: Add support for verification and cert auth.
	endpoint := winrm.NewEndpoint(
		host.Host,
		5985,  // HTTP
		false, // no SSL
		true,  // no-verify
		nil,   // CA certs
		nil,   // private key
		nil,   // cert
		(2 * time.Second),
	)

	client, err := winrm.NewClient(endpoint, host.User, host.Pass)
	if err != nil {
		return wrappedShell, err
	}

	shell, err := client.CreateShell()
	if err != nil {
		return wrappedShell, err
	}

	wrappedShell.shell = shell

	return wrappedShell, nil
}

func (s *Shell) Close() error {
	err := s.shell.Close()

	return err
}

func (s *Shell) Execute(command string) (*Result, error) {
	result := &Result{}
	cmd, err := s.shell.Execute(command)

	if err != nil {
		return result, err
	}

	var stdout, stderr bytes.Buffer
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		io.Copy(&stdout, cmd.Stdout)
	}()

	go func() {
		defer wg.Done()
		io.Copy(&stderr, cmd.Stderr)
	}()

	cmd.Wait()
	wg.Wait()

	result.Stdout = stdout.String()
	result.Stderr = stderr.String()
	result.ExitCode = cmd.ExitCode()

	return result, nil
}
