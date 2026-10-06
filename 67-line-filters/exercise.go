//go:build !solution

// Package linefilters contains exercises for https://gobyexample.com/line-filters
//
// Run the tests for this section with:
//
//	go test ./67-line-filters/
package linefilters

import "io"

// Upper is the line filter from the page: it reads r line by line and writes
// every line in upper case, followed by "\n", to w.
//
// Example: the input "hello\nfilter\n" produces "HELLO\nFILTER\n".
//
// A real program would call Upper(os.Stdin, os.Stdout). Use a bufio.Scanner
// for reading: it accepts lines ending in "\n" or "\r\n" and a last line
// without any line ending. After the loop return scanner.Err(), so a failing
// reader is reported to the caller.
func Upper(r io.Reader, w io.Writer) error {
	// TODO: scan the lines, upper-case them, print them to w.
	return nil
}

// NumberLines writes every input line prefixed with its line number, a colon
// and a space; the first line has number 1.
//
// Example: the input "foo\nbar\n" produces "1: foo\n2: bar\n".
//
// Return scanner.Err() at the end, as in Upper.
func NumberLines(r io.Reader, w io.Writer) error {
	// TODO: scan the lines and write them numbered.
	return nil
}

// Grep writes only the lines that contain substr (unchanged, followed by
// "\n") and returns how many lines it wrote.
//
// Example: with substr "an", the input "apple\nbanana\nmango\n" produces
// "banana\nmango\n" and the result 2.
func Grep(r io.Reader, w io.Writer, substr string) (int, error) {
	// TODO: scan the lines, write and count the matching ones.
	return 0, nil
}

// Uniq drops repeated lines that directly follow each other, like the Unix
// tool uniq: of every run of identical adjacent lines only the first is
// written.
//
// Example: the input "a\na\nb\na\n" produces "a\nb\na\n".
//
// Careful with empty lines: an empty first line is a line like any other and
// must be written.
func Uniq(r io.Reader, w io.Writer) error {
	// TODO: scan the lines, remember the previous one, skip repeats.
	return nil
}

// Filter is a build-your-own line filter. Every input line is passed through
// the steps in order. A step returns the (possibly changed) line and whether
// to keep it: the changed line is handed to the next step, and as soon as one
// step reports false the line is dropped and the remaining steps are not
// called for it. Lines that survive all steps are written to w followed by
// "\n". Without any steps Filter copies the lines unchanged.
//
// Example: with the steps strings.TrimSpace-and-keep followed by
// drop-if-empty, the input "  a  \n   \nb\n" produces "a\nb\n".
func Filter(r io.Reader, w io.Writer, steps ...func(line string) (string, bool)) error {
	// TODO: scan the lines, run each through the steps, write the survivors.
	return nil
}
