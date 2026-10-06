//go:build !solution

package testingandbenchmarking

import "testing"

// TestIntMinBasic checks one hard-coded call of IntMin.
//
// TODO: call IntMin with two different numbers and report a failure with
// t.Errorf when the answer is wrong. A good message says what was called,
// what came back and what was wanted.
func TestIntMinBasic(t *testing.T) {
}

// TestIntMinTableDriven checks many inputs with a table of test cases.
//
// TODO: declare a slice of structs with the inputs a, b and the wanted
// result, loop over it, and run every case as its own subtest with t.Run
// (give each subtest a name built from its inputs, e.g. with fmt.Sprintf).
//
// Think about which cases matter: the smaller number first, the smaller
// number second, negative numbers, zero.
func TestIntMinTableDriven(t *testing.T) {
}

// TODO: write a benchmark named BenchmarkIntMin that calls IntMin(1, 2) in a
// `for b.Loop()` loop. Benchmarks take a *testing.B instead of a *testing.T.
//
// Run it yourself with:
//
//	go test -bench . ./72-testing-and-benchmarking/
