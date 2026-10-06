package testingandbenchmarking

// This file grades the tests you write in intmin_test.go. You do not need to
// edit it, but reading it is a nice preview of the os/exec package.
//
// It compiles the tests of this package a second time, against
// intmin_mutant.go instead of intmin.go, and then runs that test binary once
// per bug to see whether your tests fail the way they should.

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime/debug"
	"strings"
	"testing"
)

// buildGrader compiles the tests of this package with the "mutant" build tag
// and returns the path of the resulting test binary.
func buildGrader(t *testing.T) string {
	t.Helper()
	tags := "mutant"
	// Keep whatever tags this test binary itself was built with.
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			if setting.Key == "-tags" && setting.Value != "" {
				tags += "," + setting.Value
			}
		}
	}
	bin := filepath.Join(t.TempDir(), "grader.test.exe")
	out, err := exec.Command("go", "test", "-c", "-tags", tags, "-o", bin, ".").CombinedOutput()
	if err != nil {
		t.Fatalf("could not compile the tests in this package: %v\n%s", err, out)
	}
	return bin
}

// runGrader runs the test binary with IntMin switched to the given bug ("" for
// the correct behaviour). It reports the output and whether the run failed.
func runGrader(t *testing.T, bin, bug string, args ...string) (output string, failed bool) {
	t.Helper()
	cmd := exec.Command(bin, args...)
	cmd.Env = append(os.Environ(), "GBE_MUTANT="+bug)
	out, err := cmd.CombinedOutput()
	var exitErr *exec.ExitError
	if err != nil && !errors.As(err, &exitErr) {
		t.Fatalf("could not run the compiled tests: %v", err)
	}
	return string(out), err != nil
}

func TestGrade(t *testing.T) {
	bin := buildGrader(t)

	if out, failed := runGrader(t, bin, "", "-test.run=^TestIntMin"); failed {
		t.Fatalf("your tests fail against the correct IntMin, fix that first:\n%s", out)
	}

	t.Run("TestIntMinBasic", func(t *testing.T) {
		if _, failed := runGrader(t, bin, "max", "-test.run=^TestIntMinBasic$"); !failed {
			t.Error("TestIntMinBasic still passes when IntMin returns the LARGER number - " +
				"does it call IntMin with two different numbers and check the result?")
		}
	})

	t.Run("TestIntMinTableDriven", func(t *testing.T) {
		const only = "-test.run=^TestIntMinTableDriven$"

		out, _ := runGrader(t, bin, "", only, "-test.v")
		if !strings.Contains(out, "=== RUN   TestIntMinTableDriven/") {
			t.Error("TestIntMinTableDriven runs no subtests - run each table entry with t.Run")
		}

		bugs := []struct {
			name string
			hint string
		}{
			{"first", "IntMin always returns its first argument - add a case where the smaller number comes second"},
			{"second", "IntMin always returns its second argument - add a case where the smaller number comes first"},
			{"max", "IntMin returns the larger number"},
			{"negative", "IntMin returns 0 whenever a negative number is involved - add cases with negative numbers"},
		}
		for _, bug := range bugs {
			if _, failed := runGrader(t, bin, bug.name, only); !failed {
				t.Errorf("TestIntMinTableDriven does not catch this bug: %s", bug.hint)
			}
		}
	})

	t.Run("BenchmarkIntMin", func(t *testing.T) {
		out, failed := runGrader(t, bin, "", "-test.run=^$", "-test.bench=^BenchmarkIntMin$", "-test.benchtime=10x")
		if failed {
			t.Fatalf("running the benchmark failed:\n%s", out)
		}
		if !regexp.MustCompile(`(?m)^BenchmarkIntMin(-\d+)?\s+\d+\s+[\d.]+ ns/op`).MatchString(out) {
			t.Errorf("found no result for BenchmarkIntMin, is the benchmark written? Output was:\n%s", out)
		}
	})
}
