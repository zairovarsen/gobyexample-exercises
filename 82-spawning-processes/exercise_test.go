package spawningprocesses

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

// The tests need a program to spawn that behaves the same on every operating
// system. They use this test binary itself: when the GBE_HELPER environment
// variable is set, it acts as a tiny command line tool instead of running the
// tests.
func TestMain(m *testing.M) {
	switch os.Getenv("GBE_HELPER") {
	case "echo": // print the arguments
		fmt.Println(strings.Join(os.Args[1:], " "))
	case "upper": // copy standard input to standard output in upper case
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fmt.Println(strings.ToUpper(scanner.Text()))
		}
	case "exit": // fail with the exit code given as the first argument
		code, _ := strconv.Atoi(os.Args[1])
		fmt.Fprintln(os.Stderr, "helper failing on purpose")
		os.Exit(code)
	default:
		os.Exit(m.Run())
	}
}

// helper makes spawned copies of the test binary act as the given tool and
// returns the path of the test binary.
func helper(t *testing.T, mode string) string {
	t.Helper()
	exe, err := os.Executable()
	if err != nil {
		t.Fatalf("cannot find the test binary: %v", err)
	}
	// Spawned processes inherit the environment of this process.
	t.Setenv("GBE_HELPER", mode)
	return exe
}

const noSuchProgram = "gbe-this-program-does-not-exist"

func TestOutput(t *testing.T) {
	echo := helper(t, "echo")

	got, err := Output(echo, "hello", "from", "a", "process")
	if err != nil {
		t.Fatalf("Output(echo, ...) returned an error: %v", err)
	}
	if want := "hello from a process"; got != want {
		t.Errorf("Output(echo, \"hello\", \"from\", \"a\", \"process\") = %q, want %q (trimmed)", got, want)
	}

	if got, err := Output(noSuchProgram); err == nil {
		t.Errorf("Output(%q) = %q with no error, want an error", noSuchProgram, got)
	}
}

func TestOutputOfFailingProgram(t *testing.T) {
	exit := helper(t, "exit")
	if got, err := Output(exit, "1"); err == nil {
		t.Errorf("Output of a program that exits with code 1 = %q with no error, want an error", got)
	}
}

func TestExitCode(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		echo := helper(t, "echo")
		code, err := ExitCode(echo, "fine")
		if code != 0 || err != nil {
			t.Errorf("ExitCode(echo) = %d, %v, want 0, <nil>", code, err)
		}
	})

	t.Run("failure", func(t *testing.T) {
		exit := helper(t, "exit")
		for _, want := range []int{1, 3, 42} {
			code, err := ExitCode(exit, strconv.Itoa(want))
			if code != want || err != nil {
				t.Errorf("ExitCode of a program exiting with %d = %d, %v, want %d, <nil>", want, code, err, want)
			}
		}
	})

	t.Run("cannot start", func(t *testing.T) {
		code, err := ExitCode(noSuchProgram)
		if code != -1 || err == nil {
			t.Errorf("ExitCode(%q) = %d, %v, want -1 and an error", noSuchProgram, code, err)
		}
	})
}

func TestPipeThrough(t *testing.T) {
	upper := helper(t, "upper")

	got, err := PipeThrough("hello grep\ngoodbye grep\n", upper)
	if err != nil {
		t.Fatalf("PipeThrough returned an error: %v", err)
	}
	// The helper may end lines with \r\n or \n depending on the platform.
	if lines, want := strings.Fields(got), []string{"HELLO", "GREP", "GOODBYE", "GREP"}; strings.Join(lines, " ") != strings.Join(want, " ") {
		t.Errorf("PipeThrough(\"hello grep\\ngoodbye grep\\n\", upper) = %q, want the input in upper case", got)
	}

	if got, err := PipeThrough("x", noSuchProgram); err == nil {
		t.Errorf("PipeThrough(\"x\", %q) = %q with no error, want an error", noSuchProgram, got)
	}
}

func TestPipeThroughLargeInput(t *testing.T) {
	upper := helper(t, "upper")

	type result struct {
		out string
		err error
	}
	done := make(chan result, 1)
	go func() {
		out, err := PipeThrough(strings.Repeat("spawning processes\n", 20000), upper)
		done <- result{out, err}
	}()

	select {
	case r := <-done:
		if r.err != nil {
			t.Fatalf("PipeThrough with a large input returned an error: %v", r.err)
		}
		if n := strings.Count(r.out, "SPAWNING PROCESSES"); n != 20000 {
			t.Errorf("PipeThrough with 20000 input lines returned %d upper-cased lines, want 20000", n)
		}
	case <-time.After(10 * time.Second):
		t.Fatal("PipeThrough is stuck: the program waits for someone to read its output " +
			"while PipeThrough waits to write more input - write the input from a separate goroutine")
	}
}

func TestShellCommand(t *testing.T) {
	cmd := ShellCommand("echo hello shell")
	if cmd == nil {
		t.Fatal("ShellCommand returned nil, want an *exec.Cmd")
	}
	if cmd.Process != nil {
		t.Fatal("ShellCommand already started the command, want it returned unstarted")
	}
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("running ShellCommand(\"echo hello shell\") failed: %v", err)
	}
	if got, want := strings.TrimSpace(string(out)), "hello shell"; got != want {
		t.Errorf("ShellCommand(\"echo hello shell\") printed %q, want %q", got, want)
	}
}
