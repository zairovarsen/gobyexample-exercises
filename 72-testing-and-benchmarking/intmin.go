//go:build !mutant

// Package testingandbenchmarking contains exercises for https://gobyexample.com/testing-and-benchmarking
//
// Run the tests for this section with:
//
//	go test ./72-testing-and-benchmarking/
//
// This section is the other way around: the code below is finished and YOU
// write the tests, in intmin_test.go. Your tests are then graded by
// check_test.go, which runs them against a few deliberately broken versions
// of IntMin and expects them to notice. That takes a few seconds.
package testingandbenchmarking

// IntMin returns the smaller of a and b.
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
