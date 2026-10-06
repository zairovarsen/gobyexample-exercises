package execingprocesses

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"
	"testing"
)

const noSuchProgram = "gbe-this-program-does-not-exist"

func TestFindBinary(t *testing.T) {
	// The go command must be on the PATH, since it is running these tests.
	got, err := FindBinary("go")
	if err != nil {
		t.Fatalf("FindBinary(\"go\") returned an error: %v", err)
	}
	if !filepath.IsAbs(got) {
		t.Errorf("FindBinary(\"go\") = %q, want an absolute path", got)
	}
	if name := strings.TrimSuffix(strings.ToLower(filepath.Base(got)), ".exe"); name != "go" {
		t.Errorf("FindBinary(\"go\") = %q, want the path of the go binary", got)
	}
	if _, err := os.Stat(got); err != nil {
		t.Errorf("FindBinary(\"go\") = %q, but that file cannot be found: %v", got, err)
	}
}

func TestFindBinaryNotFound(t *testing.T) {
	got, err := FindBinary(noSuchProgram)
	if err == nil {
		t.Fatalf("FindBinary(%q) = %q with no error, want an error", noSuchProgram, got)
	}
	if !errors.Is(err, exec.ErrNotFound) {
		t.Errorf("FindBinary(%q) error = %q, want it to wrap the lookup error (%%w) so that errors.Is(err, exec.ErrNotFound) holds", noSuchProgram, err)
	}
	if want := "find " + noSuchProgram + ": "; !strings.HasPrefix(err.Error(), want) {
		t.Errorf("FindBinary(%q) error = %q, want it to start with %q", noSuchProgram, err, want)
	}
}

func TestArgv(t *testing.T) {
	args := []string{"-a", "-l", "-h"}
	got := Argv("ls", args...)
	if want := []string{"ls", "-a", "-l", "-h"}; !slices.Equal(got, want) {
		t.Errorf("Argv(\"ls\", \"-a\", \"-l\", \"-h\") = %q, want %q", got, want)
	}
	if want := []string{"-a", "-l", "-h"}; !slices.Equal(args, want) {
		t.Errorf("Argv changed the caller's args to %q", args)
	}

	if got, want := Argv("date"), []string{"date"}; !slices.Equal(got, want) {
		t.Errorf("Argv(\"date\") = %q, want %q", got, want)
	}

	// Appending to a slice with spare capacity must not leak into the result.
	spare := make([]string, 1, 10)
	spare[0] = "-x"
	got = Argv("tool", spare...)
	_ = append(spare, "-y")
	if want := []string{"tool", "-x"}; !slices.Equal(got, want) {
		t.Errorf("Argv(\"tool\", \"-x\") = %q, want %q", got, want)
	}
}

func TestEnvWith(t *testing.T) {
	tests := []struct {
		name       string
		env        []string
		key, value string
		want       []string
	}{
		{"replace", []string{"A=1", "B=2", "C=3"}, "B", "x", []string{"A=1", "B=x", "C=3"}},
		{"append", []string{"A=1", "B=2"}, "C", "x", []string{"A=1", "B=2", "C=x"}},
		{"empty env", nil, "A", "1", []string{"A=1"}},
		{"key is a prefix of another key", []string{"PATHEXT=.exe", "PATH=/bin"}, "PATH", "/usr/bin", []string{"PATHEXT=.exe", "PATH=/usr/bin"}},
		{"value contains equals sign", []string{"OPTS=a=b"}, "OPTS", "c=d", []string{"OPTS=c=d"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := slices.Clone(tt.env)
			got := EnvWith(tt.env, tt.key, tt.value)
			if !slices.Equal(got, tt.want) {
				t.Errorf("EnvWith(%q, %q, %q) = %q, want %q", original, tt.key, tt.value, got, tt.want)
			}
			if !slices.Equal(tt.env, original) {
				t.Errorf("EnvWith modified its env argument: %q became %q", original, tt.env)
			}
		})
	}
}

func TestReplaceNotFound(t *testing.T) {
	err := Replace(noSuchProgram, nil, os.Environ())
	if err == nil {
		t.Fatalf("Replace(%q, ...) returned no error, want the error from FindBinary", noSuchProgram)
	}
	if !errors.Is(err, exec.ErrNotFound) {
		t.Errorf("Replace(%q, ...) error = %q, want the error from FindBinary", noSuchProgram, err)
	}
}
