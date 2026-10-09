//go:build !solution

// Package slicesx contains exercises for https://gobyexample.com/slices
//
// (The package is not called "slices" because that is the name of the
// standard library package you will import here.)
//
// Run the tests for this section with:
//
//	go test ./09-slices/
package slicesx

import "slices"

// Countdown returns the numbers from n down to 1: Countdown(3) returns
// [3 2 1]. For n < 1 the result is empty.
//
// Start with a nil slice (var s []int) and grow it with append. Remember
// that append returns the new slice.
func Countdown(n int) []int {
	var s []int
	for i := range n {
		s = append(s, n - i)
	}

	return s
}

// Middle returns s without its first and its last element, using the slice
// operator s[low:high]: Middle([1 2 3 4 5]) returns [2 3 4]. A slice with
// fewer than two elements has no middle: return an empty slice.
//
// The result must be a view of s, not a copy - slicing a slice does not copy
// the elements.
func Middle(s []int) []int {
	length := len(s)
	if length < 3 {
		return []int{}
	}
	return s[1:length - 1]
}

// Clone returns a new slice with the same elements as s that does NOT share
// memory with s: changing an element of the clone must leave s alone.
//
// Create a slice of the right length with make and fill it with the builtin
// copy.
func Clone(s []int) []int {
	s2 := make([]int, len(s))
	copy(s2, s)
	return s2
}

// HasPrefix reports whether s starts with the elements of prefix, in the same
// order. Every slice starts with the empty prefix, and a prefix longer than s
// can never match.
//
// Example: HasPrefix([a b c], [a b]) is true, HasPrefix([a b c], [b]) is
// false.
//
// Hint: cut the beginning off s with the slice operator and compare it with
// slices.Equal from the "slices" package.
func HasPrefix(s, prefix []string) bool {
	// TODO: compare the first len(prefix) elements of s with prefix.
	if len(prefix) > len(s) {
		return false
	}

	start := s[0:len(prefix)]
	return slices.Equal(start, prefix)
}

// Triangle returns a two-dimensional slice with n rows whose rows have
// different lengths: row i has i+1 elements, and element [i][j] is i + j.
//
// Example: Triangle(3) returns [[0] [1 2] [2 3 4]].
//
// Create the outer slice with make, then make every inner slice separately.
func Triangle(n int) [][]int {
	// TODO: make the rows and fill them.
	var triangle [][]int = make([][]int, n)

	for i := range n {
		innerLen := i + 1
		triangle[i] = make([]int, innerLen)
		for j := range innerLen {
			triangle[i][j] = i + j
		}	
	}
	return triangle
}

// Remove returns a NEW slice that holds the elements of s except the one at
// index i: Remove([10 20 30 40], 1) returns [10 30 40]. The index is always
// valid in the tests.
//
// The slice s that was passed in must still hold its old elements
// afterwards. That rules out the well-known shortcut of appending the tail
// to s[:i]: the result would share its memory with s and overwrite it.
// Make a fresh slice and copy the two parts into it.
func Remove(s []int, i int) []int {
	s2 := make([]int, len(s) - 1)
	copy(s2, s[:i])
	copy(s2[i:], s[i+1:])
	return s2
}
