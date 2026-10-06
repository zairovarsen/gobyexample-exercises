//go:build !solution

// Package temporaryfilesanddirectories contains exercises for https://gobyexample.com/temporary-files-and-directories
//
// Run the tests for this section with:
//
//	go test ./70-temporary-files-and-directories/
package temporaryfilesanddirectories

import "os"

// WriteTemp creates a new temporary file in the directory dir with
// os.CreateTemp, writes data to it, closes it and returns its path (f.Name()).
// The file name starts with prefix; os.CreateTemp adds a random part so that
// two calls never produce the same file.
//
// A real program usually passes "" as dir, which means "the default temporary
// directory of this OS" (os.TempDir()); the tests pass their own directory.
//
// The file is NOT removed - that is the caller's job. It must be closed
// before returning, though: on Windows a file that is still open cannot be
// deleted, which the tests check.
func WriteTemp(dir, prefix string, data []byte) (string, error) {
	// TODO: create the temp file, defer the close, write, return the name.
	return "", nil
}

// CreateTempExt creates an empty temporary file in dir whose name starts with
// prefix AND ends with the extension ext (e.g. ".txt"), closes it and returns
// its path.
//
// Hint: when the pattern given to os.CreateTemp contains a "*", the random
// part replaces the star instead of being appended: "report-*.txt" becomes
// something like "report-2847593.txt".
func CreateTempExt(dir, prefix, ext string) (string, error) {
	// TODO: build the pattern, create the file, close it.
	return "", nil
}

// Stage creates a new temporary directory in dir with os.MkdirTemp (its name
// starts with prefix) and writes every entry of files into it: the map key is
// the file name, the value the content (use os.WriteFile with permission
// 0666 like the page). It returns the path of the new directory.
//
// Example: Stage(dir, "sampledir", map[string][]byte{"file1": {1, 2}})
// creates something like dir/sampledir353748116/file1 holding the bytes 1, 2.
func Stage(dir, prefix string, files map[string][]byte) (string, error) {
	// TODO: make the temp directory, then write the files into it.
	return "", nil
}

// Scratch creates a new temporary directory in dir (name starting with
// prefix) and returns its path together with a cleanup function. Calling
// cleanup removes the directory and everything inside it (os.RemoveAll) and
// returns the error of that removal.
//
// This is the function-shaped version of the page's
// "defer os.RemoveAll(dname)": the caller writes
//
//	path, cleanup, err := Scratch("", "work")
//	if err != nil { ... }
//	defer cleanup()
//
// When the directory cannot be created, return "", nil and the error.
func Scratch(dir, prefix string) (path string, cleanup func() error, err error) {
	// TODO: make the temp directory and return a closure that removes it.
	return "", nil, nil
}

// WithTempFile creates a temporary file in dir (name starting with prefix),
// passes the open file to fn, and afterwards closes AND removes the file - no
// matter whether fn succeeded. It returns the error from fn. When the file
// cannot be created, the error is returned and fn is not called.
//
// Use defer for the clean-up. Deferred calls run in reverse order, and the
// order matters here: on Windows a file cannot be removed while it is open,
// so it has to be closed first.
func WithTempFile(dir, prefix string, fn func(f *os.File) error) error {
	// TODO: create the temp file, defer remove + close, call fn.
	return nil
}
