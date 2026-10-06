package execingprocesses

import (
	"errors"
	"os"
	"syscall"
	"testing"
)

// On Windows a process cannot be replaced, so the best the test can do is to
// check that Replace got as far as calling syscall.Exec and passed its error on.
func TestReplace(t *testing.T) {
	err := Replace("go", []string{"version"}, os.Environ())
	if err == nil {
		t.Fatal("Replace(\"go\", [\"version\"], env) returned no error, want the error of syscall.Exec (which Windows does not support)")
	}
	if !errors.Is(err, syscall.EWINDOWS) {
		t.Errorf("Replace(\"go\", [\"version\"], env) error = %q, want the \"not supported by windows\" error returned by syscall.Exec", err)
	}
}
