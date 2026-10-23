//go:build !solution

// Package rangeoverbuiltintypes contains exercises for https://gobyexample.com/range-over-built-in-types
//
// Run the tests for this section with:
//
//	go test ./16-range-over-built-in-types/
package rangeoverbuiltintypes

import "slices"

// Sum adds up the numbers in nums.
//
// range over a slice gives the index and the value of each element. The
// index is of no use here: ignore it with the blank identifier _.
func Sum(nums []int) int {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	return sum
}

// IndexOf returns the index of the first element of words that equals
// target, or -1 when there is none: IndexOf([a b c b], "b") returns 1.
//
// This time you need the index that range provides.
func IndexOf(words []string, target string) int {

	for i, word := range words {
		if word == target {
			return i
		}
	}
	
	return -1
}

// Keys returns the keys of m in ascending order: Keys(map[b:2 a:1 c:3])
// returns [a b c].
//
// range over a map with a single variable gives just the keys - in no
// particular order, which is why they must be sorted afterwards.
//
// Hint: slices.Sort from the "slices" package sorts a slice in place.
func Keys(m map[string]int) []string {
	keys := make([]string, 0, len(m))

	for k := range m {
		keys = append(keys, k)	
	}

	slices.Sort(keys)
	return keys
}

// Invert returns a new map with keys and values swapped:
// Invert(map[a:x b:y]) returns map[x:a y:b]. The tests only use maps whose
// values are all different.
//
// range over a map with two variables gives every key and its value.
func Invert(m map[string]string) map[string]string {
	mi := make(map[string]string)

	for k, v := range m {
		mi[v] = k
	}
	// TODO: for k, v := range m ...
	return mi
}

// ByteOffsets returns, for every character (rune) of s, the index of the
// byte at which it starts. In "abc" every character takes one byte, so the
// result is [0 1 2]. In "héllo" the é takes two bytes, so the result is
// [0 1 3 4 5].
//
// That is exactly the first value that range produces when it iterates over
// a string: range walks through a string rune by rune, not byte by byte.
func ByteOffsets(s string) []int {
	runes := make([]int, 0)

	for v := range s {
		runes = append(runes, v)
	}
	
	return runes
}

// GroupByLength sorts words into groups of equal length: the result maps a
// length to all the words with that many bytes (len(word)), in the order in
// which they appear in words.
//
// Example: GroupByLength([go is fun to learn]) returns
// map[2:[go is to] 3:[fun] 5:[learn]].
//
// Hint: append works on a slice that is the zero value of a missing map key.
func GroupByLength(words []string) map[int][]string {
	m := make(map[int][]string)

	for _, word := range words {
		l := len(word)
		if m[l] == nil {
			m[l] = make([]string, 0)
		}

		m[l] = append(m[l], word)
	}

	return m
}
