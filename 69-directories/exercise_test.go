package directories

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"testing"
)

// createTree creates the given slash-separated files (with some content)
// below root, including their directories. A name ending in "/" creates an
// empty directory.
func createTree(t *testing.T, root string, names ...string) {
	t.Helper()
	for _, name := range names {
		path := filepath.Join(root, filepath.FromSlash(name))
		if name[len(name)-1] == '/' {
			if err := os.MkdirAll(path, 0o755); err != nil {
				t.Fatal(err)
			}
			continue
		}
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(path, []byte("content"), 0o644); err != nil {
			t.Fatal(err)
		}
	}
}

func isDir(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func TestMakeDir(t *testing.T) {
	parent := t.TempDir()
	want := filepath.Join(parent, "subdir")

	got, err := MakeDir(parent, "subdir")
	if err != nil {
		t.Fatalf(`MakeDir(parent, "subdir") returned error: %v`, err)
	}
	if got != want {
		t.Errorf(`MakeDir(parent, "subdir") = %q, want %q`, got, want)
	}
	if !isDir(want) {
		t.Errorf(`MakeDir(parent, "subdir") did not create the directory %q`, want)
	}

	if _, err := MakeDir(parent, "subdir"); !errors.Is(err, fs.ErrExist) {
		t.Errorf("MakeDir for a directory that already exists: error = %v, want an error matching fs.ErrExist", err)
	}

	missing := filepath.Join(parent, "missing")
	if got, err := MakeDir(missing, "child"); err == nil {
		t.Errorf("MakeDir(missing parent, \"child\") = %q, nil - want an error (use os.Mkdir, not os.MkdirAll)", got)
	}
	if exists(missing) {
		t.Error("MakeDir created the missing parent directory; os.Mkdir must create one level only")
	}
}

func TestScaffold(t *testing.T) {
	root := t.TempDir()
	files := []string{"a.txt", "parent/b.txt", "parent/child/c.txt", "parent/child/d.txt", "other/deep/down/e.txt"}

	if err := Scaffold(root, files); err != nil {
		t.Fatalf("Scaffold(root, %q) returned error: %v", files, err)
	}
	for _, file := range files {
		path := filepath.Join(root, filepath.FromSlash(file))
		info, err := os.Stat(path)
		if err != nil {
			t.Errorf("Scaffold did not create %s: %v", file, err)
			continue
		}
		if info.IsDir() {
			t.Errorf("Scaffold created %s as a directory, want a file", file)
		} else if info.Size() != 0 {
			t.Errorf("Scaffold created %s with %d bytes, want an empty file", file, info.Size())
		}
	}
	if !isDir(filepath.Join(root, "parent", "child")) {
		t.Error("Scaffold did not create the directory parent/child")
	}

	// Running it again over the existing tree is fine for MkdirAll.
	if err := Scaffold(root, []string{"parent/child/again.txt"}); err != nil {
		t.Errorf("Scaffold into directories that already exist returned error: %v", err)
	}
	if !exists(filepath.Join(root, "parent", "child", "again.txt")) {
		t.Error("Scaffold did not create parent/child/again.txt in the existing directories")
	}
}

func TestListDir(t *testing.T) {
	dir := t.TempDir()
	createTree(t, dir, "gamma.txt", "alpha.txt", "beta/inner.txt", "delta/")

	got, err := ListDir(dir)
	if err != nil {
		t.Fatalf("ListDir(dir) returned error: %v", err)
	}
	want := []string{"alpha.txt", "beta/", "delta/", "gamma.txt"}
	if !slices.Equal(got, want) {
		t.Errorf("ListDir(dir) = %q, want %q", got, want)
	}

	got, err = ListDir(filepath.Join(dir, "delta"))
	if err != nil || len(got) != 0 {
		t.Errorf("ListDir(empty dir) = %q, %v, want no entries and a nil error", got, err)
	}

	if _, err := ListDir(filepath.Join(dir, "missing")); err == nil {
		t.Error("ListDir(missing dir) returned a nil error, want the error from os.ReadDir")
	}
}

func TestFindFiles(t *testing.T) {
	root := t.TempDir()
	createTree(t, root,
		"z.go", "a.go", "notes.txt", "go/",
		"sub/b.go", "sub/readme.md", "sub/deep/c.go", "sub/deep/c.txt",
		"dir.go/inside.txt",
	)

	tests := []struct {
		ext  string
		want []string
	}{
		{".go", []string{"a.go", "sub/b.go", "sub/deep/c.go", "z.go"}},
		{".txt", []string{"dir.go/inside.txt", "notes.txt", "sub/deep/c.txt"}},
		{".md", []string{"sub/readme.md"}},
		{".rs", nil},
	}
	for _, tt := range tests {
		got, err := FindFiles(root, tt.ext)
		if err != nil {
			t.Errorf("FindFiles(root, %q) returned error: %v", tt.ext, err)
			continue
		}
		if !slices.Equal(got, tt.want) {
			t.Errorf("FindFiles(root, %q) = %q, want %q", tt.ext, got, tt.want)
		}
	}

	if _, err := FindFiles(filepath.Join(root, "missing"), ".go"); err == nil {
		t.Error("FindFiles(missing root) returned a nil error; return the err argument of the visit function")
	}
}

func TestWithScratchDir(t *testing.T) {
	parent := t.TempDir()
	want := filepath.Join(parent, "scratch")

	called := false
	err := WithScratchDir(parent, func(dir string) error {
		called = true
		if dir != want {
			t.Errorf("fn received the directory %q, want %q", dir, want)
		}
		if !isDir(dir) {
			t.Errorf("the directory %q does not exist while fn runs", dir)
			return nil
		}
		// Leave something behind that the clean-up has to remove.
		createTree(t, dir, "tmp.txt", "nested/more/tmp.txt")
		return nil
	})
	if !called {
		t.Fatal("WithScratchDir never called fn")
	}
	if err != nil {
		t.Errorf("WithScratchDir returned error %v, want the nil that fn returned", err)
	}
	if exists(want) {
		t.Error("the scratch directory still exists after WithScratchDir returned")
	}

	// The clean-up must also happen when fn fails, and fn's error is passed on.
	errWork := errors.New("work failed")
	err = WithScratchDir(parent, func(dir string) error {
		createTree(t, dir, "partial.txt")
		return errWork
	})
	if !errors.Is(err, errWork) {
		t.Errorf("WithScratchDir error = %v, want the error returned by fn", err)
	}
	if exists(want) {
		t.Error("the scratch directory still exists after fn returned an error; use defer for the clean-up")
	}

	// When the directory cannot be created fn must not run.
	called = false
	err = WithScratchDir(filepath.Join(parent, "missing"), func(string) error {
		called = true
		return nil
	})
	if err == nil {
		t.Error("WithScratchDir(missing parent) returned a nil error, want the error from os.Mkdir")
	}
	if called {
		t.Error("WithScratchDir(missing parent) called fn although the directory could not be created")
	}
}
