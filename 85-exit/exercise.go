//go:build !solution

// Package exit contains exercises for https://gobyexample.com/exit
//
// Run the tests for this section with:
//
//	go test ./85-exit/
//
// os.Exit ends the program on the spot, so the tests run the functions that
// call it in a separate process and look at that process's exit status.
package exit

import (
	"errors"
	"io"
)

// ErrUsage marks errors caused by calling a program the wrong way.
var ErrUsage = errors.New("usage error")

// StatusFor turns the error a program ended with into its exit status,
// following the usual convention:
//
//   - nil: 0 (success)
//   - ErrUsage, or any error wrapping it: 2
//   - every other error: 1
func StatusFor(err error) int {
	// TODO: pick the status.
	return 0
}

// Greet is the "real main" of a small greeting program. Because os.Exit does
// not run deferred functions, a common pattern is to keep main tiny:
//
//	func main() { os.Exit(Greet(os.Args[1:], os.Stdout, os.Stderr)) }
//
// and do all the work (including any defers) in a function that returns the
// exit status.
//
// With exactly one argument, Greet writes "hello <arg>\n" to stdout and
// returns 0. With any other number of arguments it writes
// "usage: greet <name>\n" to stderr, writes nothing to stdout and returns 2.
func Greet(args []string, stdout, stderr io.Writer) int {
	// TODO: check the arguments, write the output, return the status.
	return 0
}

// ExitAfter ends the program with the given exit status, but runs cleanup
// first. A `defer cleanup()` would NOT work here - deferred functions are
// skipped by os.Exit - so cleanup has to be called explicitly.
func ExitAfter(status int, cleanup func()) {
	// TODO: clean up, then exit.
}

// MustSucceed is a helper for programs that cannot continue after an error.
// When err is nil it does nothing. Otherwise it writes "error: <err>\n" to
// stderr and exits with the status that StatusFor gives for err.
func MustSucceed(err error, stderr io.Writer) {
	// TODO: report the error and exit, if there is one.
}
