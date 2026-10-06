//go:build solution

package spawningprocesses

import (
	"errors"
	"io"
	"os/exec"
	"runtime"
	"strings"
)

func Output(name string, args ...string) (string, error) {
	out, err := exec.Command(name, args...).Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

func ExitCode(name string, args ...string) (int, error) {
	err := exec.Command(name, args...).Run()
	if err == nil {
		return 0, nil
	}
	var exitErr *exec.ExitError
	if errors.As(err, &exitErr) {
		return exitErr.ExitCode(), nil
	}
	return -1, err
}

func PipeThrough(input string, name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return "", err
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}
	if err := cmd.Start(); err != nil {
		return "", err
	}

	// Write from a separate goroutine: with a large input the program may
	// start writing output (and block on it) before it has read everything.
	go func() {
		io.WriteString(stdin, input)
		stdin.Close()
	}()

	out, readErr := io.ReadAll(stdout)
	// Wait must come after all reads from the pipe are done.
	if err := cmd.Wait(); err != nil {
		return "", err
	}
	if readErr != nil {
		return "", readErr
	}
	return string(out), nil
}

func ShellCommand(script string) *exec.Cmd {
	if runtime.GOOS == "windows" {
		return exec.Command("cmd", "/C", script)
	}
	return exec.Command("bash", "-c", script)
}
