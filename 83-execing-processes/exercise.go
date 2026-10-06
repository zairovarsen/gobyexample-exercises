//go:build !solution

// Package execingprocesses contains exercises for https://gobyexample.com/execing-processes
//
// Run the tests for this section with:
//
//	go test ./83-execing-processes/
//
// Exec replaces the running Go process with another program. That is a Unix
// feature: on Windows syscall.Exec exists but always returns an error
// ("not supported by windows"), and the usual substitute is to spawn the
// program, wait for it, and exit with its exit code. The preparation steps -
// finding the binary, building the argument list and the environment - are
// the same everywhere, and that is what most of these exercises practise.
package execingprocesses

// FindBinary returns the absolute path of the program name, found by looking
// through the directories of the PATH environment variable, e.g. "go" ->
// "/usr/local/go/bin/go" or `C:\Program Files\Go\bin\go.exe`.
//
// When the program is not found, the error from the lookup must be returned
// wrapped with the program name, in the form "find <name>: <lookup error>".
//
// Hint: exec.LookPath, and fmt.Errorf with %w.
func FindBinary(name string) (string, error) {
	// TODO: look the program up.
	return "", nil
}

// Argv builds the argument list for Exec. Exec wants the arguments in "argv"
// form, where the first element is the program name itself:
//
//	Argv("ls", "-a", "-l") returns ["ls", "-a", "-l"]
//
// The result must be a new slice; the caller's args must stay untouched.
func Argv(name string, args ...string) []string {
	// TODO: build the slice.
	return nil
}

// EnvWith returns a copy of env in which the variable key is set to value.
// env is in the form that os.Environ returns: a slice of "KEY=value" strings.
//
// An existing entry for key is replaced in place (same position); otherwise
// "key=value" is appended. env itself must not be modified.
//
//	EnvWith(["A=1", "B=2"], "B", "x") returns ["A=1", "B=x"]
//	EnvWith(["A=1", "B=2"], "C", "x") returns ["A=1", "B=2", "C=x"]
func EnvWith(env []string, key, value string) []string {
	// TODO: copy env and set the variable.
	return nil
}

// Replace replaces the current process with the program name, run with the
// given arguments (without the program name; Replace adds it) and environment.
//
// It finds the binary with FindBinary and returns that error if the program
// does not exist. Then it calls syscall.Exec. On success Exec never returns,
// because this process no longer exists; if Exec fails (as it always does on
// Windows) Replace returns its error.
//
// A real program would call it like: Replace("ls", []string{"-a"}, os.Environ())
func Replace(name string, args []string, env []string) error {
	// TODO: find the binary, build argv, exec.
	return nil
}
