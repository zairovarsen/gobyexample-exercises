//go:build solution

package exit

import (
	"errors"
	"fmt"
	"io"
	"os"
)

var ErrUsage = errors.New("usage error")

func StatusFor(err error) int {
	switch {
	case err == nil:
		return 0
	case errors.Is(err, ErrUsage):
		return 2
	default:
		return 1
	}
}

func Greet(args []string, stdout, stderr io.Writer) int {
	if len(args) != 1 {
		fmt.Fprintln(stderr, "usage: greet <name>")
		return 2
	}
	fmt.Fprintln(stdout, "hello", args[0])
	return 0
}

func ExitAfter(status int, cleanup func()) {
	cleanup()
	os.Exit(status)
}

func MustSucceed(err error, stderr io.Writer) {
	if err == nil {
		return
	}
	fmt.Fprintf(stderr, "error: %v\n", err)
	os.Exit(StatusFor(err))
}
