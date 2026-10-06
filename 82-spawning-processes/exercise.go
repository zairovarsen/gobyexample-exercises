//go:build !solution

// Package spawningprocesses contains exercises for https://gobyexample.com/spawning-processes
//
// Run the tests for this section with:
//
//	go test ./82-spawning-processes/
//
// The page spawns Unix tools such as date and grep. To work on every operating
// system, the functions here take the program to run as a parameter and the
// tests pass in a small helper program.
package spawningprocesses

import "os/exec"

// Output runs the program name with the given arguments, waits for it to
// finish and returns everything it wrote to standard output, with leading and
// trailing whitespace trimmed.
//
// If the program cannot be run or exits with a failure, the error is returned.
func Output(name string, args ...string) (string, error) {
	// TODO: build an exec.Command and collect its output.
	return "", nil
}

// ExitCode runs the program and reports how it ended:
//
//   - it ran and succeeded: 0 and a nil error
//   - it ran but exited with a failure: its exit code and a nil error
//   - it could not be started at all (e.g. the program does not exist):
//     -1 and the error
//
// Hint: a failed exit is reported as an *exec.ExitError; tell the cases apart
// with errors.As.
func ExitCode(name string, args ...string) (int, error) {
	// TODO: run the command and inspect the error.
	return 0, nil
}

// PipeThrough starts the program, writes input to its standard input, closes
// standard input so the program sees the end of the data, and returns
// everything the program wrote to standard output (untrimmed).
//
// Practise the explicit way the page shows for grep: StdinPipe, StdoutPipe,
// Start, write, close, read everything, Wait.
//
// Challenge: pipes have a small buffer. With a large input the program starts
// writing output while you are still writing input; if nobody reads that
// output, both sides wait for each other forever. The last test sends a large
// input - make the writing and the reading happen at the same time.
func PipeThrough(input string, name string, args ...string) (string, error) {
	// TODO: wire up the pipes and run the command.
	return "", nil
}

// ShellCommand returns a command (not started yet) that runs script through
// the system shell, so that a whole command line can be given as one string:
//
//   - on Windows:        cmd /C <script>
//   - everywhere else:   bash -c <script>
//
// Hint: runtime.GOOS tells you the operating system the program runs on.
func ShellCommand(script string) *exec.Cmd {
	// TODO: return the right exec.Command for this operating system.
	return nil
}
