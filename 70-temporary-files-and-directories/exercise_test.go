package temporaryfilesanddirectories

import (
	"bytes"
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// checkTempName verifies that path names something directly inside dir whose
// name is prefix plus a random part.
func checkTempName(t *testing.T, call, path, dir, prefix string) {
	t.Helper()
	if path == "" {
		t.Fatalf("%s returned an empty path", call)
	}
	if got := filepath.Dir(path); got != dir {
		t.Errorf("%s created %q in the directory %q, want it in %q", call, path, got, dir)
	}
	name := filepath.Base(path)
	if !strings.HasPrefix(name, prefix) {
		t.Errorf("%s created %q, want a name starting with %q", call, name, prefix)
	}
	if len(name) <= len(prefix) {
		t.Errorf("%s created %q, want a random part after the prefix (use os.CreateTemp / os.MkdirTemp)", call, name)
	}
}

// checkClosed removes the file. On Windows that fails while the file is
// still open, which reveals a missing Close.
func checkClosed(t *testing.T, path string) {
	t.Helper()
	if err := os.Remove(path); err != nil {
		t.Errorf("could not remove the file afterwards - did you close it? (%v)", err)
	}
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// missingDir returns a directory that does not exist, so nothing can be
// created inside it.
func missingDir(t *testing.T) string {
	t.Helper()
	return filepath.Join(t.TempDir(), "no-such-dir")
}

func TestWriteTemp(t *testing.T) {
	dir := t.TempDir()
	data := []byte{1, 2, 3, 4}

	first, err := WriteTemp(dir, "sample", data)
	if err != nil {
		t.Fatalf(`WriteTemp(dir, "sample", [1 2 3 4]) returned error: %v`, err)
	}
	checkTempName(t, `WriteTemp(dir, "sample", ...)`, first, dir, "sample")
	got, err := os.ReadFile(first)
	if err != nil {
		t.Fatalf("the temporary file cannot be read: %v", err)
	}
	if !bytes.Equal(got, data) {
		t.Errorf("the temporary file holds %v, want %v", got, data)
	}

	second, err := WriteTemp(dir, "sample", data)
	if err != nil {
		t.Fatalf("second WriteTemp call returned error: %v", err)
	}
	if first == second {
		t.Errorf("two WriteTemp calls returned the same path %q, want two different files", first)
	}

	checkClosed(t, first)
	checkClosed(t, second)

	if path, err := WriteTemp(missingDir(t), "sample", data); err == nil {
		t.Errorf("WriteTemp(missing dir, ...) = %q, nil - want the error from os.CreateTemp", path)
	}
}

func TestCreateTempExt(t *testing.T) {
	dir := t.TempDir()
	tests := []struct {
		prefix, ext string
	}{
		{"report-", ".txt"},
		{"img", ".png"},
		{"plain", ""},
	}
	for _, tt := range tests {
		call := "CreateTempExt(dir, \"" + tt.prefix + "\", \"" + tt.ext + "\")"
		path, err := CreateTempExt(dir, tt.prefix, tt.ext)
		if err != nil {
			t.Fatalf("%s returned error: %v", call, err)
		}
		checkTempName(t, call, path, dir, tt.prefix)
		name := filepath.Base(path)
		if !strings.HasSuffix(name, tt.ext) {
			t.Errorf("%s created %q, want a name ending in %q", call, name, tt.ext)
		}
		if len(name) <= len(tt.prefix)+len(tt.ext) {
			t.Errorf("%s created %q, want a random part between prefix and extension", call, name)
		}
		info, err := os.Stat(path)
		if err != nil {
			t.Fatalf("%s: the file does not exist: %v", call, err)
		}
		if info.Size() != 0 {
			t.Errorf("%s created a file of %d bytes, want an empty file", call, info.Size())
		}
		checkClosed(t, path)
	}

	if path, err := CreateTempExt(missingDir(t), "report-", ".txt"); err == nil {
		t.Errorf("CreateTempExt(missing dir, ...) = %q, nil - want the error from os.CreateTemp", path)
	}
}

func TestStage(t *testing.T) {
	dir := t.TempDir()
	files := map[string][]byte{
		"file1":     {1, 2},
		"notes.txt": []byte("hello\n"),
		"empty":     {},
	}

	path, err := Stage(dir, "sampledir", files)
	if err != nil {
		t.Fatalf(`Stage(dir, "sampledir", 3 files) returned error: %v`, err)
	}
	checkTempName(t, `Stage(dir, "sampledir", ...)`, path, dir, "sampledir")
	if info, err := os.Stat(path); err != nil || !info.IsDir() {
		t.Fatalf("Stage returned %q, which is not a directory", path)
	}
	for name, want := range files {
		got, err := os.ReadFile(filepath.Join(path, name))
		if err != nil {
			t.Errorf("Stage did not write %s: %v", name, err)
			continue
		}
		if !bytes.Equal(got, want) {
			t.Errorf("Stage wrote %q into %s, want %q", got, name, want)
		}
	}
	entries, err := os.ReadDir(path)
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) != len(files) {
		t.Errorf("the staged directory has %d entries, want %d", len(entries), len(files))
	}

	other, err := Stage(dir, "sampledir", nil)
	if err != nil {
		t.Fatalf("Stage(dir, \"sampledir\", nil) returned error: %v", err)
	}
	if other == path {
		t.Errorf("two Stage calls returned the same directory %q, want two different ones", path)
	}

	if path, err := Stage(missingDir(t), "sampledir", files); err == nil {
		t.Errorf("Stage(missing dir, ...) = %q, nil - want the error from os.MkdirTemp", path)
	}
}

func TestScratch(t *testing.T) {
	dir := t.TempDir()

	path, cleanup, err := Scratch(dir, "work")
	if err != nil {
		t.Fatalf(`Scratch(dir, "work") returned error: %v`, err)
	}
	checkTempName(t, `Scratch(dir, "work")`, path, dir, "work")
	if info, err := os.Stat(path); err != nil || !info.IsDir() {
		t.Fatalf("Scratch returned %q, which is not a directory", path)
	}
	if cleanup == nil {
		t.Fatal("Scratch returned a nil cleanup function")
	}

	// Fill the directory: cleanup has to remove a non-empty tree.
	nested := filepath.Join(path, "a", "b")
	if err := os.MkdirAll(nested, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(nested, "f.txt"), []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}

	if err := cleanup(); err != nil {
		t.Errorf("cleanup() returned error: %v", err)
	}
	if exists(path) {
		t.Errorf("the directory %q still exists after cleanup()", filepath.Base(path))
	}
	if !exists(dir) {
		t.Error("cleanup() removed the parent directory too; it must remove only the temporary directory")
	}

	path, cleanup, err = Scratch(missingDir(t), "work")
	if err == nil {
		t.Errorf("Scratch(missing dir, ...) returned path %q and a nil error, want the error from os.MkdirTemp", path)
	}
	if err != nil && (path != "" || cleanup != nil) {
		t.Error(`Scratch(missing dir, ...) should return "" and a nil cleanup together with the error`)
	}
}

func TestWithTempFile(t *testing.T) {
	dir := t.TempDir()

	var file *os.File
	var name string
	err := WithTempFile(dir, "session-", func(f *os.File) error {
		if f == nil {
			t.Error("fn received a nil file")
			return nil
		}
		file, name = f, f.Name()
		if !exists(name) {
			t.Errorf("the temporary file %q does not exist while fn runs", name)
		}
		if _, err := f.Write([]byte("scratch data")); err != nil {
			t.Errorf("writing to the temporary file inside fn failed: %v", err)
		}
		return nil
	})
	if file == nil {
		t.Fatal("WithTempFile never called fn with a file")
	}
	if err != nil {
		t.Errorf("WithTempFile returned error %v, want the nil that fn returned", err)
	}
	checkTempName(t, `WithTempFile(dir, "session-", fn)`, name, dir, "session-")
	if _, err := file.Write([]byte("x")); !errors.Is(err, os.ErrClosed) {
		t.Errorf("the file is still open after WithTempFile returned (write error = %v, want os.ErrClosed)", err)
	}
	if exists(name) {
		t.Error("the temporary file still exists after WithTempFile returned (on Windows: close it before removing it)")
	}

	// The clean-up must also happen when fn fails, and fn's error is passed on.
	errWork := errors.New("work failed")
	name = ""
	err = WithTempFile(dir, "session-", func(f *os.File) error {
		name = f.Name()
		return errWork
	})
	if !errors.Is(err, errWork) {
		t.Errorf("WithTempFile error = %v, want the error returned by fn", err)
	}
	if name != "" && exists(name) {
		t.Error("the temporary file still exists after fn returned an error; use defer for the clean-up")
	}

	called := false
	err = WithTempFile(missingDir(t), "session-", func(*os.File) error {
		called = true
		return nil
	})
	if err == nil {
		t.Error("WithTempFile(missing dir, ...) returned a nil error, want the error from os.CreateTemp")
	}
	if called {
		t.Error("WithTempFile(missing dir, ...) called fn although no file could be created")
	}
	if _, err := os.Stat(dir); errors.Is(err, fs.ErrNotExist) {
		t.Error("the parent directory disappeared; remove only the temporary file")
	}
}
