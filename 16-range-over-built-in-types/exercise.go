//go:build !solution

// Package rangeoverbuiltintypes contains exercises for https://gobyexample.com/range-over-built-in-types
//
// Run the tests for this section with:
//
//	go test ./16-range-over-built-in-types/
package rangeoverbuiltintypes

// Sum adds up the numbers in nums.
//
// range over a slice gives the index and the value of each element. The
// index is of no use here: ignore it with the blank identifier _.
func Sum(nums []int) int {
	// TODO: for _, num := range nums ...
	return 0
}

// IndexOf returns the index of the first element of words that equals
// target, or -1 when there is none: IndexOf([a b c b], "b") returns 1.
//
// This time you need the index that range provides.
func IndexOf(words []string, target string) int {
	// TODO: range with index and value, return at the first match.
	return 0
}

// Keys returns the keys of m in ascending order: Keys(map[b:2 a:1 c:3])
// returns [a b c].
//
// range over a map with a single variable gives just the keys - in no
// particular order, which is why they must be sorted afterwards.
//
// Hint: slices.Sort from the "slices" package sorts a slice in place.
func Keys(m map[string]int) []string {
	// TODO: collect the keys with range, then sort them.
	return nil
}

// Invert returns a new map with keys and values swapped:
// Invert(map[a:x b:y]) returns map[x:a y:b]. The tests only use maps whose
// values are all different.
//
// range over a map with two variables gives every key and its value.
func Invert(m map[string]string) map[string]string {
	// TODO: for k, v := range m ...
	return nil
}

// ByteOffsets returns, for every character (rune) of s, the index of the
// byte at which it starts. In "abc" every character takes one byte, so the
// result is [0 1 2]. In "héllo" the é takes two bytes, so the result is
// [0 1 3 4 5].
//
// That is exactly the first value that range produces when it iterates over
// a string: range walks through a string rune by rune, not byte by byte.
func ByteOffsets(s string) []int {
	// TODO: range over the string and collect the index of every rune.
	return nil
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
	// TODO: range over words and append every word to its group.
	return nil
}
