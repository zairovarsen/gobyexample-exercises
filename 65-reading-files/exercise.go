//go:build !solution

// Package readingfiles contains exercises for https://gobyexample.com/reading-files
//
// Run the tests for this section with:
//
//	go test ./65-reading-files/
package readingfiles

// ReadAll returns the entire content of the file at path as a string - the
// most basic way of reading a file, slurping it into memory in one call.
//
// Hint: os.ReadFile opens, reads and closes the file for you.
func ReadAll(path string) (string, error) {
	// TODO: read the whole file and convert the bytes to a string.
	return "", nil
}

// ReadFirst opens the file at path with os.Open and reads up to n bytes from
// its start with a single f.Read call. It returns only the bytes that were
// really read: for a file holding "hello go\n", ReadFirst(path, 5) returns
// "hello" and ReadFirst(path, 1000) returns all 9 bytes.
//
// Reading from an empty file makes Read report io.EOF; that is not a failure
// here - return an empty result and a nil error.
//
// The file must be closed before returning (defer f.Close()). On Windows a
// file that is still open cannot be deleted, which the tests check.
func ReadFirst(path string, n int) ([]byte, error) {
	// TODO: open, defer the close, read into a buffer of n bytes.
	return nil, nil
}

// ReadSection returns exactly n bytes starting at byte offset (counted from
// the start of the file). For a file holding "hello go\n",
// ReadSection(path, 6, 2) returns "go".
//
// Use f.Seek with io.SeekStart to jump to the offset and io.ReadAtLeast to
// fill the buffer. When the file has fewer than n bytes after offset, return
// the error from io.ReadAtLeast. Close the file before returning.
func ReadSection(path string, offset int64, n int) ([]byte, error) {
	// TODO: open, seek, read at least n bytes.
	return nil, nil
}

// ReadTail returns the last n bytes of the file. For a file holding
// "hello go\n", ReadTail(path, 3) returns "go\n". When the file is shorter
// than n bytes the whole file is returned.
//
// Use f.Seek relative to the end of the file (io.SeekEnd) - no reading of the
// part you do not need. Close the file before returning.
//
// Hint: Seek returns the new offset, so f.Seek(0, io.SeekEnd) tells you the
// size of the file.
func ReadTail(path string, n int64) ([]byte, error) {
	// TODO: open, find the size, seek to n bytes before the end, read.
	return nil, nil
}

// PeekAndReadLine wraps the file in a bufio.Reader, peeks at the first n
// bytes WITHOUT consuming them, and then reads the first line with the same
// reader. Because peeking does not advance the reader, the line still starts
// at the beginning of the file.
//
// For a file holding "hello go\nsecond line\n", PeekAndReadLine(path, 5)
// returns peeked "hello" and line "hello go".
//
// The line is returned without its line ending ("\n" or "\r\n"). A file
// without any line ending is one single line. When the file has fewer than n
// bytes, return the error from Peek. Close the file before returning.
func PeekAndReadLine(path string, n int) (peeked, line string, err error) {
	// TODO: open, bufio.NewReader, Peek, then read up to the first '\n'.
	return "", "", nil
}

// FindLines returns every line of the file that contains substr, prefixed
// with its line number (the first line is number 1) in the form "N: line".
//
// For a file holding "apple pie\nbanana\npineapple\n",
// FindLines(path, "apple") returns ["1: apple pie", "3: pineapple"].
//
// Read the file line by line with a bufio.Scanner (it strips "\n" as well as
// "\r\n") and remember to check scanner.Err() after the loop. Close the file
// before returning.
func FindLines(path, substr string) ([]string, error) {
	// TODO: open, scan line by line, collect the matching lines.
	return nil, nil
}
