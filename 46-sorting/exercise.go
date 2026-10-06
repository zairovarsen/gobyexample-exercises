//go:build !solution

// Package sorting contains exercises for https://gobyexample.com/sorting
//
// Run the tests for this section with:
//
//	go test ./46-sorting/
package sorting

import "cmp"

// SortStrings sorts strs in place in ascending order, e.g. [c a b] becomes
// [a b c]. Strings are compared byte by byte, so upper-case letters come
// before lower-case ones.
//
// Hint: the "slices" package sorts a slice of any ordered built-in type.
func SortStrings(strs []string) {
	// TODO: sort the slice in place.
}

// SortedInts returns a new slice holding the numbers of nums in ascending
// order, e.g. [7 2 4] -> [2 4 7]. The nums slice itself must stay untouched.
func SortedInts(nums []int) []int {
	// TODO: copy nums, sort the copy and return it.
	return nil
}

// IsSorted reports whether nums is already in ascending order. An empty slice
// and a slice with one element count as sorted.
//
// Hint: the "slices" package can answer this without sorting anything.
func IsSorted(nums []int) bool {
	// TODO: check the order.
	return false
}

// SortedKeys returns the keys of m in ascending order. Map iteration order is
// random, so sorting the keys is the usual way to get a stable result.
//
// Example: {"pear": 1, "apple": 5, "fig": 2} -> [apple fig pear].
func SortedKeys(m map[string]int) []string {
	// TODO: collect the keys, then sort them.
	return nil
}

// TopN returns the n largest items, largest first, e.g.
// TopN([]int{3, 9, 1, 7}, 2) returns [9 7]. It works for every ordered type.
//
// When n is larger than len(items) all items are returned (largest first);
// when n is zero or negative the result is empty. The items slice must not be
// modified.
func TopN[T cmp.Ordered](items []T, n int) []T {
	// TODO: sort a copy, then pick the n largest in descending order.
	return nil
}
