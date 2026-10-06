//go:build !solution

// Package writingfiles contains exercises for https://gobyexample.com/writing-files
//
// Run the tests for this section with:
//
//	go test ./66-writing-files/
package writingfiles

// WriteAll dumps content into the file at path in one call, creating the file
// when it does not exist and replacing its content when it does. Use
// permission 0644 for a new file.
//
// Hint: os.WriteFile takes a []byte.
func WriteAll(path, content string) error {
	// TODO: write the content with a single call.
	return nil
}

// WriteParts creates (or truncates) the file at path with os.Create and
// writes two pieces to it: first the bytes in data with f.Write, then text
// with f.WriteString. It calls f.Sync to flush the writes to stable storage
// and returns the total number of bytes written.
//
// Example: WriteParts(path, []byte("some\n"), "writes\n") leaves a file
// holding "some\nwrites\n" and returns 12.
//
// The file must be closed before returning (defer f.Close() right after the
// error check). On Windows a file that is still open cannot be deleted, which
// the tests check.
func WriteParts(path string, data []byte, text string) (int, error) {
	// TODO: create, defer the close, Write, WriteString, Sync.
	return 0, nil
}

// WriteLines creates (or truncates) the file at path and writes every element
// of lines followed by "\n" through a bufio.Writer. It returns the total
// number of bytes written.
//
// Example: WriteLines(path, []string{"a", "b"}) leaves a file holding
// "a\nb\n" and returns 4.
//
// A buffered writer keeps data in memory until its buffer is full: call Flush
// when you are done or the last (or all!) lines never reach the file. Close
// the file before returning.
func WriteLines(path string, lines []string) (int, error) {
	// TODO: create, wrap in a bufio.Writer, write the lines, Flush.
	return 0, nil
}

// CopyUpper reads the file src line by line and writes every line in upper
// case, followed by "\n", to the file dst (created or truncated). It returns
// the number of lines copied.
//
// Example: a src holding "hello\ngo by example\n" produces a dst holding
// "HELLO\nGO BY EXAMPLE\n" and the result 2. Lines in src may end in "\n" or
// "\r\n", and the last line may have no line ending at all.
//
// Open src first: when it cannot be opened, return the error without creating
// dst. Use a bufio.Scanner for reading and a bufio.Writer for writing, and
// close both files before returning.
func CopyUpper(src, dst string) (int, error) {
	// TODO: open src, create dst, scan + write, check scanner.Err, Flush.
	return 0, nil
}
