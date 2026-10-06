package writingfiles

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// readBack returns the content of the file an exercise was supposed to write.
func readBack(t *testing.T, path string) string {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("the file was not written: %v", err)
	}
	return string(data)
}

// checkClosed removes the file. On Windows that fails while the file is
// still open, which reveals a missing Close.
func checkClosed(t *testing.T, path string) {
	t.Helper()
	if err := os.Remove(path); err != nil {
		t.Errorf("could not remove the file afterwards - did you close it? (%v)", err)
	}
}

// existingFile creates a file that already has (long) content.
func existingFile(t *testing.T) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "existing.txt")
	if err := os.WriteFile(path, []byte(strings.Repeat("old content\n", 50)), 0o644); err != nil {
		t.Fatal(err)
	}
	return path
}

// badPath returns a path inside a directory that does not exist, so creating
// a file there has to fail.
func badPath(t *testing.T) string {
	t.Helper()
	return filepath.Join(t.TempDir(), "no-such-dir", "out.txt")
}

func TestWriteAll(t *testing.T) {
	path := filepath.Join(t.TempDir(), "dat1.txt")
	if err := WriteAll(path, "hello\ngo\n"); err != nil {
		t.Fatalf(`WriteAll(new file, "hello\ngo\n") returned error: %v`, err)
	}
	if got, want := readBack(t, path), "hello\ngo\n"; got != want {
		t.Errorf("WriteAll(new file, %q) wrote %q", want, got)
	}
	checkClosed(t, path)

	path = existingFile(t)
	if err := WriteAll(path, "new"); err != nil {
		t.Fatalf(`WriteAll(existing file, "new") returned error: %v`, err)
	}
	if got := readBack(t, path); got != "new" {
		t.Errorf(`WriteAll(existing file, "new") left %q in the file, want "new"`, got)
	}

	if err := WriteAll(badPath(t), "x"); err == nil {
		t.Error("WriteAll(file in a missing directory) returned a nil error, want the error from os.WriteFile")
	}
}

func TestWriteParts(t *testing.T) {
	path := filepath.Join(t.TempDir(), "dat2.txt")
	n, err := WriteParts(path, []byte{115, 111, 109, 101, 10}, "writes\n")
	if err != nil {
		t.Fatalf("WriteParts(new file, ...) returned error: %v", err)
	}
	if got, want := readBack(t, path), "some\nwrites\n"; got != want {
		t.Errorf(`WriteParts(new file, "some\n", "writes\n") wrote %q, want %q`, got, want)
	}
	if n != 12 {
		t.Errorf(`WriteParts(new file, "some\n", "writes\n") returned %d bytes, want 12`, n)
	}
	checkClosed(t, path)

	path = existingFile(t)
	n, err = WriteParts(path, []byte("ab"), "c")
	if err != nil {
		t.Fatalf("WriteParts(existing file, ...) returned error: %v", err)
	}
	if got := readBack(t, path); got != "abc" {
		t.Errorf(`WriteParts(existing file, "ab", "c") left %q in the file, want "abc" (os.Create truncates)`, got)
	}
	if n != 3 {
		t.Errorf(`WriteParts(existing file, "ab", "c") returned %d bytes, want 3`, n)
	}
	checkClosed(t, path)

	if _, err := WriteParts(badPath(t), []byte("a"), "b"); err == nil {
		t.Error("WriteParts(file in a missing directory, ...) returned a nil error, want the error from os.Create")
	}
}

func TestWriteLines(t *testing.T) {
	// A small file fits completely into the bufio.Writer's buffer, so it
	// only shows up on disk after Flush.
	path := filepath.Join(t.TempDir(), "small.txt")
	n, err := WriteLines(path, []string{"a", "b"})
	if err != nil {
		t.Fatalf(`WriteLines(file, ["a" "b"]) returned error: %v`, err)
	}
	if got, want := readBack(t, path), "a\nb\n"; got != want {
		t.Errorf(`WriteLines(file, ["a" "b"]) wrote %q, want %q (did you Flush?)`, got, want)
	}
	if n != 4 {
		t.Errorf(`WriteLines(file, ["a" "b"]) returned %d bytes, want 4`, n)
	}
	checkClosed(t, path)

	// A large file is bigger than the buffer; its tail needs the Flush.
	lines := make([]string, 5000)
	for i := range lines {
		lines[i] = fmt.Sprintf("line %04d", i)
	}
	want := strings.Join(lines, "\n") + "\n"
	path = existingFile(t)
	n, err = WriteLines(path, lines)
	if err != nil {
		t.Fatalf("WriteLines(file, 5000 lines) returned error: %v", err)
	}
	if got := readBack(t, path); got != want {
		t.Errorf("WriteLines(file, 5000 lines) wrote %d bytes of content, want %d bytes (did you Flush?)", len(got), len(want))
	}
	if n != len(want) {
		t.Errorf("WriteLines(file, 5000 lines) returned %d bytes, want %d", n, len(want))
	}
	checkClosed(t, path)

	path = filepath.Join(t.TempDir(), "empty.txt")
	if _, err := WriteLines(path, nil); err != nil {
		t.Fatalf("WriteLines(file, nil) returned error: %v", err)
	}
	if got := readBack(t, path); got != "" {
		t.Errorf("WriteLines(file, nil) wrote %q, want an empty file", got)
	}

	if _, err := WriteLines(badPath(t), []string{"a"}); err == nil {
		t.Error("WriteLines(file in a missing directory, ...) returned a nil error, want the error from os.Create")
	}
}

func TestCopyUpper(t *testing.T) {
	tests := []struct {
		content   string
		want      string
		wantLines int
	}{
		{"hello\ngo by example\n", "HELLO\nGO BY EXAMPLE\n", 2},
		{"a\r\nb", "A\nB\n", 2},
		{"one\n\nthree\n", "ONE\n\nTHREE\n", 3},
	}
	for _, tt := range tests {
		dir := t.TempDir()
		src, dst := filepath.Join(dir, "src.txt"), filepath.Join(dir, "dst.txt")
		if err := os.WriteFile(src, []byte(tt.content), 0o644); err != nil {
			t.Fatal(err)
		}
		n, err := CopyUpper(src, dst)
		if err != nil {
			t.Fatalf("CopyUpper(src %q, dst) returned error: %v", tt.content, err)
		}
		if got := readBack(t, dst); got != tt.want {
			t.Errorf("CopyUpper(src %q, dst) wrote %q, want %q", tt.content, got, tt.want)
		}
		if n != tt.wantLines {
			t.Errorf("CopyUpper(src %q, dst) returned %d lines, want %d", tt.content, n, tt.wantLines)
		}
		checkClosed(t, src)
		checkClosed(t, dst)
	}

	dir := t.TempDir()
	dst := filepath.Join(dir, "dst.txt")
	if _, err := CopyUpper(filepath.Join(dir, "missing.txt"), dst); err == nil {
		t.Error("CopyUpper(missing src, dst) returned a nil error, want the error from os.Open")
	}
	if _, err := os.Stat(dst); !errors.Is(err, fs.ErrNotExist) {
		t.Error("CopyUpper(missing src, dst) created dst anyway; open src first and return its error")
	}
}
