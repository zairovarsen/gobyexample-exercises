//go:build !solution

// Package directories contains exercises for https://gobyexample.com/directories
//
// Run the tests for this section with:
//
//	go test ./69-directories/
package directories

// MakeDir creates the single directory name inside the existing directory
// parent with os.Mkdir (permission 0755) and returns the path of the new
// directory, built with filepath.Join.
//
// os.Mkdir creates exactly one level: when parent does not exist, or when the
// directory is already there, return os.Mkdir's error (and an empty path).
func MakeDir(parent, name string) (string, error) {
	// TODO: join the path, create the directory, return the path.
	return "", nil
}

// Scaffold creates an empty file for every entry of files below root. The
// entries are slash-separated relative paths such as "parent/child/file.txt";
// directories on the way that do not exist yet are created too.
//
// Example: Scaffold(root, []string{"a.txt", "parent/child/b.txt"}) creates
// root/a.txt and root/parent/child/b.txt.
//
// Use os.MkdirAll (the "mkdir -p" of Go, permission 0755) for the directory
// of each file and os.WriteFile with empty content (permission 0644) for the
// file itself.
//
// Hint: filepath.FromSlash converts "a/b" to the separators of the current
// OS, filepath.Dir gives the directory of a path.
func Scaffold(root string, files []string) error {
	// TODO: for every file create its directory, then the empty file.
	return nil
}

// ListDir returns the names of the entries directly inside dir, as os.ReadDir
// reports them (sorted by name), with a "/" appended to the name of every
// directory.
//
// Example: a directory holding the files a.txt and c.txt and the directory b
// gives ["a.txt", "b/", "c.txt"].
func ListDir(dir string) ([]string, error) {
	// TODO: read the directory and build the list of names.
	return nil, nil
}

// FindFiles walks the whole directory tree below root with filepath.WalkDir
// and returns every file (not directory) whose extension equals ext, e.g.
// ".go". The results are paths relative to root, written with forward slashes
// on every operating system, in the order the walk visits them.
//
// Example: for a tree with root/a.go, root/notes.txt and root/sub/b.go,
// FindFiles(root, ".go") returns ["a.go", "sub/b.go"].
//
// The visit function receives an error argument: return it when it is not
// nil, so that a root that does not exist makes FindFiles return an error.
//
// Hint: filepath.Rel and filepath.ToSlash.
func FindFiles(root, ext string) ([]string, error) {
	// TODO: walk the tree and collect the matching files.
	return nil, nil
}

// WithScratchDir creates the directory "scratch" inside parent, calls fn with
// the path of that directory, and removes the directory with everything in it
// (os.RemoveAll) before returning - no matter whether fn succeeded. It
// returns the error from fn.
//
// When the directory cannot be created, the error is returned and fn is not
// called.
//
// Use defer for the clean-up, just like the page does right after creating
// its directory.
func WithScratchDir(parent string, fn func(dir string) error) error {
	// TODO: create the directory, defer its removal, call fn.
	return nil
}
