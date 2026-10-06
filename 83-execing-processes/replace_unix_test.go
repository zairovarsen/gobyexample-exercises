//go:build !windows

package execingprocesses

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
)

// When GBE_HELPER is set, the test binary calls Replace instead of running
// the tests, so that a spawned copy of it can be replaced safely.
func TestMain(m *testing.M) {
	if os.Getenv("GBE_HELPER") == "replace" {
		os.Unsetenv("GBE_HELPER")
		err := Replace("go", []string{"version"}, os.Environ())
		// Only reached when the exec did not happen.
		fmt.Println("still the old process:", err)
		os.Exit(1)
	}
	os.Exit(m.Run())
}

func TestReplace(t *testing.T) {
	exe, err := os.Executable()
	if err != nil {
		t.Fatalf("cannot find the test binary: %v", err)
	}
	cmd := exec.Command(exe)
	cmd.Env = append(os.Environ(), "GBE_HELPER=replace")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("a process calling Replace(\"go\", [\"version\"], env) failed: %v\n%s", err, out)
	}
	if !strings.HasPrefix(string(out), "go version") {
		t.Errorf("a process calling Replace(\"go\", [\"version\"], env) printed %q, want the output of `go version`", out)
	}
}
