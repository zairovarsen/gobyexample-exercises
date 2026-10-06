package deferx

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"
)

// catch calls f and reports whether it panicked. It uses recover, which is
// the topic of the next section.
func catch(f func()) (panicked bool) {
	defer func() {
		if r := recover(); r != nil {
			panicked = true
		}
	}()
	f()
	return false
}

func TestBracket(t *testing.T) {
	var buf bytes.Buffer
	Bracket(&buf, func() { buf.WriteString("body\n") })
	if got, want := buf.String(), "enter\nbody\nexit\n"; got != want {
		t.Errorf("Bracket wrote %q, want %q", got, want)
	}
}

func TestBracketPanickingBody(t *testing.T) {
	var buf bytes.Buffer
	panicked := catch(func() {
		Bracket(&buf, func() { panic("boom") })
	})
	if got, want := buf.String(), "enter\nexit\n"; got != want {
		t.Errorf("with a panicking body Bracket wrote %q, want %q (the deferred call still runs)", got, want)
	}
	if !panicked {
		t.Errorf("the panic of the body did not come out of Bracket; defer must not swallow it")
	}
}

func TestReverseOrder(t *testing.T) {
	tests := []struct {
		names []string
		want  []string
	}{
		{[]string{"a", "b", "c"}, []string{"body", "c", "b", "a"}},
		{[]string{"open", "lock"}, []string{"body", "lock", "open"}},
		{nil, []string{"body"}},
	}
	for _, tt := range tests {
		var got []string
		ReverseOrder(tt.names, func(s string) { got = append(got, s) })
		if !slices.Equal(got, tt.want) {
			t.Errorf("ReverseOrder(%q, record) made the calls %q, want %q", tt.names, got, tt.want)
		}
	}
}

func TestReverseOrderUsesDefer(t *testing.T) {
	// When the direct call panics, only calls that really were deferred
	// still happen.
	var got []string
	catch(func() {
		ReverseOrder([]string{"a", "b", "c"}, func(s string) {
			if s == "body" {
				panic("boom")
			}
			got = append(got, s)
		})
	})
	if want := []string{"c", "b", "a"}; !slices.Equal(got, want) {
		t.Errorf("with record(\"body\") panicking, the remaining calls were %q, want %q - defer the calls", got, want)
	}
}

func TestWithdraw(t *testing.T) {
	a := &Account{balance: 100}
	steps := []struct {
		amount      int
		wantErr     error
		wantBalance int
	}{
		{30, nil, 70},
		{0, ErrInvalidAmount, 70},
		{-5, ErrInvalidAmount, 70},
		{71, ErrInsufficientFunds, 70},
		{70, nil, 0},
		{1, ErrInsufficientFunds, 0},
	}
	for _, step := range steps {
		err := a.Withdraw(step.amount)
		if !errors.Is(err, step.wantErr) {
			t.Errorf("Withdraw(%d) returned error %v, want %v", step.amount, err, step.wantErr)
		}
		// Every return path must leave the mutex unlocked.
		if !a.mu.TryLock() {
			t.Fatalf("after Withdraw(%d) returned %v the mutex is still locked", step.amount, err)
		}
		balance := a.balance
		a.mu.Unlock()
		if balance != step.wantBalance {
			t.Errorf("after Withdraw(%d) the balance is %d, want %d", step.amount, balance, step.wantBalance)
		}
	}
}

// fakeFile is an io.WriteCloser that remembers what happened to it and can
// be told to fail.
type fakeFile struct {
	data     bytes.Buffer
	failOn   string // a Write containing this text fails with errWrite
	closeErr error  // returned by Close
	closed   int    // number of Close calls
}

var (
	errWrite = errors.New("disk full")
	errClose = errors.New("close failed")
)

func (f *fakeFile) Write(p []byte) (int, error) {
	if f.closed > 0 {
		return 0, errors.New("write after close")
	}
	if f.failOn != "" && bytes.Contains(p, []byte(f.failOn)) {
		return 0, errWrite
	}
	return f.data.Write(p)
}

func (f *fakeFile) Close() error {
	f.closed++
	return f.closeErr
}

func TestWriteAll(t *testing.T) {
	f := &fakeFile{}
	err := WriteAll(f, []string{"one", "two", "three"})
	if err != nil {
		t.Errorf("WriteAll returned error %v, want nil", err)
	}
	if got, want := f.data.String(), "one\ntwo\nthree\n"; got != want {
		t.Errorf("WriteAll wrote %q, want %q", got, want)
	}
	if f.closed != 1 {
		t.Errorf("WriteAll called Close %d times, want exactly once", f.closed)
	}
}

func TestWriteAllWriteFails(t *testing.T) {
	f := &fakeFile{failOn: "bad"}
	err := WriteAll(f, []string{"ok", "bad", "never"})
	if !errors.Is(err, errWrite) {
		t.Errorf("WriteAll returned error %v, want the write error %q", err, errWrite)
	}
	if got, want := f.data.String(), "ok\n"; got != want {
		t.Errorf("WriteAll wrote %q, want %q (stop at the first failing write)", got, want)
	}
	if f.closed != 1 {
		t.Errorf("after a failed write WriteAll called Close %d times, want exactly once", f.closed)
	}
}

func TestWriteAllCloseFails(t *testing.T) {
	f := &fakeFile{closeErr: errClose}
	err := WriteAll(f, []string{"one"})
	if !errors.Is(err, errClose) {
		t.Errorf("WriteAll returned error %v, want the close error %q", err, errClose)
	}
	if f.closed != 1 {
		t.Errorf("WriteAll called Close %d times, want exactly once", f.closed)
	}

	// When both fail, the write error wins: it happened first.
	f = &fakeFile{failOn: "bad", closeErr: errClose}
	err = WriteAll(f, []string{"bad"})
	if !errors.Is(err, errWrite) {
		t.Errorf("with a failing write and a failing close WriteAll returned %v, want the write error %q", err, errWrite)
	}
	if f.closed != 1 {
		t.Errorf("WriteAll called Close %d times, want exactly once", f.closed)
	}
}

func TestWriteFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "defer.txt")
	if err := WriteFile(path, []string{"data", "more data"}); err != nil {
		t.Fatalf("WriteFile(%q, ...) returned error %v, want nil", path, err)
	}
	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("WriteFile did not create the file: %v", err)
	}
	got := strings.ReplaceAll(string(content), "\r\n", "\n")
	if want := "data\nmore data\n"; got != want {
		t.Errorf("the file contains %q, want %q", got, want)
	}

	// Writing again replaces the old content.
	if err := WriteFile(path, []string{"new"}); err != nil {
		t.Fatalf("second WriteFile(%q, ...) returned error %v, want nil", path, err)
	}
	content, err = os.ReadFile(path)
	if err != nil {
		t.Fatalf("reading the file back: %v", err)
	}
	got = strings.ReplaceAll(string(content), "\r\n", "\n")
	if want := "new\n"; got != want {
		t.Errorf("after the second WriteFile the file contains %q, want %q", got, want)
	}

	// On Windows a file that is still open cannot be removed.
	if err := os.Remove(path); err != nil {
		t.Errorf("removing the file failed (%v) - was it closed?", err)
	}
}

func TestWriteFileError(t *testing.T) {
	path := filepath.Join(t.TempDir(), "no-such-dir", "defer.txt")
	if err := WriteFile(path, []string{"data"}); err == nil {
		t.Errorf("WriteFile(%q, ...) returned nil, want an error because the directory does not exist", path)
	}
}
