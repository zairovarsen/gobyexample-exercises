//go:build mutant

package testingandbenchmarking

import "os"

// This file is only compiled by check_test.go. It replaces IntMin with a
// version that can be switched to one of several buggy behaviours, to find
// out whether your tests catch the bugs. You do not need to edit it.

func IntMin(a, b int) int {
	switch os.Getenv("GBE_MUTANT") {
	case "first":
		return a
	case "second":
		return b
	case "max":
		if a > b {
			return a
		}
		return b
	case "negative":
		// Correct for non-negative numbers only.
		if a < 0 || b < 0 {
			return 0
		}
	}
	if a < b {
		return a
	}
	return b
}
