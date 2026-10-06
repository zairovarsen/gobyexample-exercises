package readingfiles

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"testing"
)

// sample is 27 bytes long: "hello go\n" (0-8), "second line\n" (9-20),
// "third\n" (21-26).
const sample = "hello go\nsecond line\nthird\n"

// writeFile creates a file with the given content in a fresh temporary
// directory and returns its path.
func writeFile(t *testing.T, content string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "data.txt")
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	return path
}

// missingFile returns a path at which no file exists.
func missingFile(t *testing.T) string {
	t.Helper()
	return filepath.Join(t.TempDir(), "no-such-file.txt")
}

// checkClosed removes the file. On Windows that fails while the file is
// still open, which reveals a missing Close.
func checkClosed(t *testing.T, path string) {
	t.Helper()
	if err := os.Remove(path); err != nil {
		t.Errorf("could not remove the file afterwards - did you close it? (%v)", err)
	}
}

func TestReadAll(t *testing.T) {
	path := writeFile(t, sample)
	got, err := ReadAll(path)
	if err != nil {
		t.Fatalf("ReadAll(file) returned error: %v", err)
	}
	if got != sample {
		t.Errorf("ReadAll(file) = %q, want %q", got, sample)
	}
	checkClosed(t, path)

	if _, err := ReadAll(missingFile(t)); !errors.Is(err, fs.ErrNotExist) {
		t.Errorf("ReadAll(missing file) error = %v, want an error matching fs.ErrNotExist", err)
	}
}

func TestReadFirst(t *testing.T) {
	path := writeFile(t, sample)
	tests := []struct {
		n    int
		want string
	}{
		{5, "hello"},
		{8, "hello go"},
		{1000, sample},
	}
	for _, tt := range tests {
		got, err := ReadFirst(path, tt.n)
		if err != nil {
			t.Errorf("ReadFirst(file, %d) returned error: %v", tt.n, err)
			continue
		}
		if string(got) != tt.want {
			t.Errorf("ReadFirst(file, %d) = %q, want %q", tt.n, got, tt.want)
		}
	}
	checkClosed(t, path)

	empty := writeFile(t, "")
	if got, err := ReadFirst(empty, 5); err != nil || len(got) != 0 {
		t.Errorf("ReadFirst(empty file, 5) = %q, %v, want an empty result and a nil error", got, err)
	}
	checkClosed(t, empty)

	if _, err := ReadFirst(missingFile(t), 5); err == nil {
		t.Error("ReadFirst(missing file, 5) returned a nil error, want the error from os.Open")
	}
}

func TestReadSection(t *testing.T) {
	path := writeFile(t, sample)
	tests := []struct {
		offset int64
		n      int
		want   string
	}{
		{6, 2, "go"},
		{9, 6, "second"},
		{0, 5, "hello"},
		{0, len(sample), sample},
	}
	for _, tt := range tests {
		got, err := ReadSection(path, tt.offset, tt.n)
		if err != nil {
			t.Errorf("ReadSection(file, %d, %d) returned error: %v", tt.offset, tt.n, err)
			continue
		}
		if string(got) != tt.want {
			t.Errorf("ReadSection(file, %d, %d) = %q, want %q", tt.offset, tt.n, got, tt.want)
		}
	}

	if got, err := ReadSection(path, 20, 100); err == nil {
		t.Errorf("ReadSection(file, 20, 100) = %q, nil - want an error because only 7 bytes are left", got)
	}
	checkClosed(t, path)

	if _, err := ReadSection(missingFile(t), 0, 1); err == nil {
		t.Error("ReadSection(missing file, 0, 1) returned a nil error, want the error from os.Open")
	}
}

func TestReadTail(t *testing.T) {
	path := writeFile(t, sample)
	tests := []struct {
		n    int64
		want string
	}{
		{6, "third\n"},
		{1, "\n"},
		{int64(len(sample)), sample},
		{1000, sample},
	}
	for _, tt := range tests {
		got, err := ReadTail(path, tt.n)
		if err != nil {
			t.Errorf("ReadTail(file, %d) returned error: %v", tt.n, err)
			continue
		}
		if string(got) != tt.want {
			t.Errorf("ReadTail(file, %d) = %q, want %q", tt.n, got, tt.want)
		}
	}
	if got, err := ReadTail(path, 0); err != nil || len(got) != 0 {
		t.Errorf("ReadTail(file, 0) = %q, %v, want an empty result and a nil error", got, err)
	}
	checkClosed(t, path)

	if _, err := ReadTail(missingFile(t), 3); err == nil {
		t.Error("ReadTail(missing file, 3) returned a nil error, want the error from os.Open")
	}
}

func TestPeekAndReadLine(t *testing.T) {
	tests := []struct {
		content    string
		n          int
		wantPeeked string
		wantLine   string
	}{
		{sample, 5, "hello", "hello go"},
		{sample, 11, "hello go\nse", "hello go"},
		{"first\r\nsecond\r\n", 3, "fir", "first"},
		{"only", 2, "on", "only"},
	}
	for _, tt := range tests {
		path := writeFile(t, tt.content)
		peeked, line, err := PeekAndReadLine(path, tt.n)
		if err != nil {
			t.Errorf("PeekAndReadLine(file %q, %d) returned error: %v", tt.content, tt.n, err)
			continue
		}
		if peeked != tt.wantPeeked || line != tt.wantLine {
			t.Errorf("PeekAndReadLine(file %q, %d) = %q, %q, want %q, %q",
				tt.content, tt.n, peeked, line, tt.wantPeeked, tt.wantLine)
		}
		checkClosed(t, path)
	}

	short := writeFile(t, "hi")
	if _, _, err := PeekAndReadLine(short, 10); err == nil {
		t.Error(`PeekAndReadLine(file "hi", 10) returned a nil error, want the error from Peek`)
	}
	checkClosed(t, short)

	if _, _, err := PeekAndReadLine(missingFile(t), 1); err == nil {
		t.Error("PeekAndReadLine(missing file, 1) returned a nil error, want the error from os.Open")
	}
}

func TestFindLines(t *testing.T) {
	tests := []struct {
		content string
		substr  string
		want    []string
	}{
		{"apple pie\nbanana\npineapple\ncherry\n", "apple", []string{"1: apple pie", "3: pineapple"}},
		{"apple pie\r\nbanana\r\npineapple\r\n", "apple", []string{"1: apple pie", "3: pineapple"}},
		{"one\ntwo\nthree", "t", []string{"2: two", "3: three"}},
		{"one\ntwo\n", "", []string{"1: one", "2: two"}},
		{"one\ntwo\n", "zzz", nil},
	}
	for _, tt := range tests {
		path := writeFile(t, tt.content)
		got, err := FindLines(path, tt.substr)
		if err != nil {
			t.Errorf("FindLines(file %q, %q) returned error: %v", tt.content, tt.substr, err)
			continue
		}
		if !slices.Equal(got, tt.want) {
			t.Errorf("FindLines(file %q, %q) = %q, want %q", tt.content, tt.substr, got, tt.want)
		}
		checkClosed(t, path)
	}

	if _, err := FindLines(missingFile(t), "x"); err == nil {
		t.Error(`FindLines(missing file, "x") returned a nil error, want the error from os.Open`)
	}
}
