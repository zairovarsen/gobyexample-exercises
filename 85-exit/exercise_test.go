package exit

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
)

// Functions that call os.Exit would end the test run, so they are tested in a
// separate process: when GBE_HELPER is set, this test binary runs one of the
// scenarios below instead of the tests.
func TestMain(m *testing.M) {
	switch os.Getenv("GBE_HELPER") {
	case "exit-after":
		ExitAfter(3, func() { fmt.Println("cleanup done") })
		fmt.Println("still running")
	case "must-succeed-nil":
		MustSucceed(nil, os.Stderr)
		fmt.Println("still running")
	case "must-succeed-error":
		MustSucceed(errors.New("disk full"), os.Stderr)
		fmt.Println("still running")
	case "must-succeed-usage":
		MustSucceed(fmt.Errorf("missing file name: %w", ErrUsage), os.Stderr)
		fmt.Println("still running")
	default:
		os.Exit(m.Run())
	}
}

// runScenario runs the test binary as a separate process in the given helper
// mode and returns what it printed and its exit status.
func runScenario(t *testing.T, mode string) (stdout, stderr string, status int) {
	t.Helper()
	exe, err := os.Executable()
	if err != nil {
		t.Fatalf("cannot find the test binary: %v", err)
	}
	cmd := exec.Command(exe)
	cmd.Env = append(os.Environ(), "GBE_HELPER="+mode)
	var outBuf, errBuf bytes.Buffer
	cmd.Stdout, cmd.Stderr = &outBuf, &errBuf

	err = cmd.Run()
	var exitErr *exec.ExitError
	switch {
	case err == nil:
		status = 0
	case errors.As(err, &exitErr):
		status = exitErr.ExitCode()
	default:
		t.Fatalf("cannot run the test binary: %v", err)
	}
	clean := func(s string) string { return strings.ReplaceAll(s, "\r\n", "\n") }
	return clean(outBuf.String()), clean(errBuf.String()), status
}

func TestStatusFor(t *testing.T) {
	tests := []struct {
		err  error
		want int
	}{
		{nil, 0},
		{errors.New("boom"), 1},
		{ErrUsage, 2},
		{fmt.Errorf("unknown flag -x: %w", ErrUsage), 2},
		{fmt.Errorf("outer: %w", fmt.Errorf("inner: %w", ErrUsage)), 2},
		{errors.New("usage error"), 1}, // same text, but not ErrUsage
	}
	for _, tt := range tests {
		if got := StatusFor(tt.err); got != tt.want {
			t.Errorf("StatusFor(%v) = %d, want %d", tt.err, got, tt.want)
		}
	}
}

func TestGreet(t *testing.T) {
	tests := []struct {
		name       string
		args       []string
		wantStatus int
		wantOut    string
		wantErr    string
	}{
		{"one argument", []string{"gopher"}, 0, "hello gopher\n", ""},
		{"no arguments", nil, 2, "", "usage: greet <name>\n"},
		{"too many arguments", []string{"a", "b"}, 2, "", "usage: greet <name>\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stdout, stderr bytes.Buffer
			status := Greet(tt.args, &stdout, &stderr)
			if status != tt.wantStatus {
				t.Errorf("Greet(%q) returned status %d, want %d", tt.args, status, tt.wantStatus)
			}
			if got := stdout.String(); got != tt.wantOut {
				t.Errorf("Greet(%q) wrote %q to stdout, want %q", tt.args, got, tt.wantOut)
			}
			if got := stderr.String(); got != tt.wantErr {
				t.Errorf("Greet(%q) wrote %q to stderr, want %q", tt.args, got, tt.wantErr)
			}
		})
	}
}

func TestExitAfter(t *testing.T) {
	stdout, _, status := runScenario(t, "exit-after")
	if status != 3 {
		t.Errorf("a program calling ExitAfter(3, cleanup) exited with status %d, want 3", status)
	}
	if !strings.Contains(stdout, "cleanup done") {
		t.Errorf("cleanup did not run before the exit (program printed %q)", stdout)
	}
	if strings.Contains(stdout, "still running") {
		t.Error("the program kept running after ExitAfter, want it to exit")
	}
}

func TestMustSucceed(t *testing.T) {
	t.Run("nil error", func(t *testing.T) {
		stdout, stderr, status := runScenario(t, "must-succeed-nil")
		if status != 0 || !strings.Contains(stdout, "still running") || stderr != "" {
			t.Errorf("after MustSucceed(nil) the program had status %d, stdout %q, stderr %q; want it to carry on normally and print nothing to stderr", status, stdout, stderr)
		}
	})

	failures := []struct {
		mode       string
		wantStatus int
		wantStderr string
	}{
		{"must-succeed-error", 1, "error: disk full\n"},
		{"must-succeed-usage", 2, "error: missing file name: usage error\n"},
	}
	for _, tt := range failures {
		t.Run(tt.mode, func(t *testing.T) {
			stdout, stderr, status := runScenario(t, tt.mode)
			if status != tt.wantStatus {
				t.Errorf("exit status = %d, want %d", status, tt.wantStatus)
			}
			if stderr != tt.wantStderr {
				t.Errorf("stderr = %q, want %q", stderr, tt.wantStderr)
			}
			if strings.Contains(stdout, "still running") {
				t.Error("the program kept running after MustSucceed got an error, want it to exit")
			}
		})
	}
}
