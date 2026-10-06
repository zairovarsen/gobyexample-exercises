package linefilters

import (
	"bytes"
	"errors"
	"io"
	"strings"
	"testing"
	"testing/iotest"
)

var errBroken = errors.New("broken reader")

// brokenReader returns a reader that fails with errBroken on the first Read,
// like a pipe that went away.
func brokenReader() io.Reader {
	return iotest.ErrReader(errBroken)
}

func TestUpper(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"hello\nfilter\n", "HELLO\nFILTER\n"},
		{"hello\r\nfilter\r\n", "HELLO\nFILTER\n"},
		{"no newline at the end", "NO NEWLINE AT THE END\n"},
		{"a\n\nb\n", "A\n\nB\n"},
		{"", ""},
	}
	for _, tt := range tests {
		var out bytes.Buffer
		if err := Upper(strings.NewReader(tt.input), &out); err != nil {
			t.Errorf("Upper(%q) returned error: %v", tt.input, err)
			continue
		}
		if got := out.String(); got != tt.want {
			t.Errorf("Upper(%q) wrote %q, want %q", tt.input, got, tt.want)
		}
	}

	if err := Upper(brokenReader(), io.Discard); !errors.Is(err, errBroken) {
		t.Errorf("Upper(broken reader) error = %v, want the reader's error (return scanner.Err())", err)
	}
}

func TestNumberLines(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"foo\nbar\n", "1: foo\n2: bar\n"},
		{"foo\r\nbar\r\nbaz", "1: foo\n2: bar\n3: baz\n"},
		{"\n\n", "1: \n2: \n"},
		{"", ""},
	}
	for _, tt := range tests {
		var out bytes.Buffer
		if err := NumberLines(strings.NewReader(tt.input), &out); err != nil {
			t.Errorf("NumberLines(%q) returned error: %v", tt.input, err)
			continue
		}
		if got := out.String(); got != tt.want {
			t.Errorf("NumberLines(%q) wrote %q, want %q", tt.input, got, tt.want)
		}
	}

	if err := NumberLines(brokenReader(), io.Discard); !errors.Is(err, errBroken) {
		t.Errorf("NumberLines(broken reader) error = %v, want the reader's error (return scanner.Err())", err)
	}
}

func TestGrep(t *testing.T) {
	tests := []struct {
		input  string
		substr string
		want   string
		wantN  int
	}{
		{"apple\nbanana\nmango\n", "an", "banana\nmango\n", 2},
		{"apple\r\nbanana\r\nmango\r\n", "an", "banana\nmango\n", 2},
		{"apple\nbanana\n", "", "apple\nbanana\n", 2},
		{"apple\nbanana\n", "kiwi", "", 0},
		{"Go\ngo\n", "go", "go\n", 1},
	}
	for _, tt := range tests {
		var out bytes.Buffer
		n, err := Grep(strings.NewReader(tt.input), &out, tt.substr)
		if err != nil {
			t.Errorf("Grep(%q, %q) returned error: %v", tt.input, tt.substr, err)
			continue
		}
		if got := out.String(); got != tt.want || n != tt.wantN {
			t.Errorf("Grep(%q, %q) wrote %q and returned %d, want %q and %d",
				tt.input, tt.substr, got, n, tt.want, tt.wantN)
		}
	}

	if _, err := Grep(brokenReader(), io.Discard, "x"); !errors.Is(err, errBroken) {
		t.Errorf("Grep(broken reader) error = %v, want the reader's error (return scanner.Err())", err)
	}
}

func TestUniq(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"a\na\nb\na\n", "a\nb\na\n"},
		{"a\r\na\r\nb\r\n", "a\nb\n"},
		{"x\nx\nx\nx", "x\n"},
		{"a\nb\nc\n", "a\nb\nc\n"},
		{"\n\na\n\n\nb\n", "\na\n\nb\n"},
		{"", ""},
	}
	for _, tt := range tests {
		var out bytes.Buffer
		if err := Uniq(strings.NewReader(tt.input), &out); err != nil {
			t.Errorf("Uniq(%q) returned error: %v", tt.input, err)
			continue
		}
		if got := out.String(); got != tt.want {
			t.Errorf("Uniq(%q) wrote %q, want %q", tt.input, got, tt.want)
		}
	}

	if err := Uniq(brokenReader(), io.Discard); !errors.Is(err, errBroken) {
		t.Errorf("Uniq(broken reader) error = %v, want the reader's error (return scanner.Err())", err)
	}
}

func TestFilter(t *testing.T) {
	trim := func(line string) (string, bool) { return strings.TrimSpace(line), true }
	dropEmpty := func(line string) (string, bool) { return line, line != "" }
	shout := func(line string) (string, bool) { return strings.ToUpper(line) + "!", true }

	var out bytes.Buffer
	if err := Filter(strings.NewReader("  a  \n   \nb\n"), &out, trim, dropEmpty); err != nil {
		t.Fatalf("Filter(trim, dropEmpty) returned error: %v", err)
	}
	if got, want := out.String(), "a\nb\n"; got != want {
		t.Errorf(`Filter("  a  \n   \nb\n", trim, dropEmpty) wrote %q, want %q`, got, want)
	}

	// The steps run in the order given: each one sees the previous result.
	out.Reset()
	if err := Filter(strings.NewReader(" go \r\n\r\nby example"), &out, trim, dropEmpty, shout); err != nil {
		t.Fatalf("Filter(trim, dropEmpty, shout) returned error: %v", err)
	}
	if got, want := out.String(), "GO!\nBY EXAMPLE!\n"; got != want {
		t.Errorf(`Filter(" go \r\n\r\nby example", trim, dropEmpty, shout) wrote %q, want %q`, got, want)
	}

	// A dropped line must not reach the steps that follow.
	var seen []string
	record := func(line string) (string, bool) {
		seen = append(seen, line)
		return line, true
	}
	out.Reset()
	if err := Filter(strings.NewReader("keep\n\nalso\n"), &out, dropEmpty, record); err != nil {
		t.Fatalf("Filter(dropEmpty, record) returned error: %v", err)
	}
	if len(seen) != 2 || seen[0] != "keep" || seen[1] != "also" {
		t.Errorf("after dropEmpty the next step was called for %q, want only [\"keep\" \"also\"]", seen)
	}

	out.Reset()
	if err := Filter(strings.NewReader("one\ntwo\n"), &out); err != nil {
		t.Fatalf("Filter() without steps returned error: %v", err)
	}
	if got, want := out.String(), "one\ntwo\n"; got != want {
		t.Errorf(`Filter("one\ntwo\n") without steps wrote %q, want %q`, got, want)
	}

	if err := Filter(brokenReader(), io.Discard, trim); !errors.Is(err, errBroken) {
		t.Errorf("Filter(broken reader) error = %v, want the reader's error (return scanner.Err())", err)
	}
}
