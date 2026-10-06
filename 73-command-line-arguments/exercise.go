//go:build !solution

// Package commandlinearguments contains exercises for https://gobyexample.com/command-line-arguments
//
// Run the tests for this section with:
//
//	go test ./73-command-line-arguments/
package commandlinearguments

// In all exercises of this section args has the shape of os.Args: the first
// element is the name (path) of the program, the real arguments follow. A
// real program would call these functions with os.Args; taking the slice as a
// parameter is what makes them testable.

// ProgramName returns the first element of args - the program name - or ""
// when args is empty.
//
// Example: ProgramName([]string{"./prog", "a", "b"}) is "./prog".
func ProgramName(args []string) string {
	// TODO: return the first element, if there is one.
	return ""
}

// Arguments returns the arguments without the program name, like the page's
// os.Args[1:]. When there are none (or args is empty) the result is empty.
//
// Example: Arguments([]string{"./prog", "a", "b"}) is ["a", "b"].
func Arguments(args []string) []string {
	// TODO: slice off the program name; mind the empty slice.
	return nil
}

// ArgAt returns the n-th argument, counted from 1 and not counting the
// program name, and true. When there is no such argument (n is smaller than 1
// or larger than the number of arguments) it returns "" and false instead of
// panicking like a plain os.Args[n] would.
//
// Example: ArgAt([]string{"./prog", "a", "b", "c", "d"}, 3) is "c", true.
func ArgAt(args []string, n int) (string, bool) {
	// TODO: check the bounds, then index.
	return "", false
}

// Sum parses every argument (not the program name) as a whole number and
// returns the total. Without arguments the sum is 0. When an argument is not
// a number, Sum returns 0 and an error.
//
// Example: Sum([]string{"./sum", "1", "2", "39"}) is 42.
//
// Hint: command-line arguments are always strings; strconv.Atoi converts.
func Sum(args []string) (int, error) {
	// TODO: loop over the arguments, convert and add them up.
	return 0, nil
}

// Repeat implements a tiny command-line program with the usage
//
//	repeat <count> <word>...
//
// It returns the words joined by single spaces as one line ending in "\n",
// repeated count times.
//
// Example: Repeat([]string{"./repeat", "2", "hello", "world"}) returns
// "hello world\nhello world\n".
//
// It returns "" and an error when fewer than two arguments are given (the
// count and at least one word) and when count is not a number or is negative.
// A count of 0 is valid and gives "".
func Repeat(args []string) (string, error) {
	// TODO: validate the arguments, then build the output.
	return "", nil
}
