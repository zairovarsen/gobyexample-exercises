//go:build !solution

// Package sortingbyfunctions contains exercises for https://gobyexample.com/sorting-by-functions
//
// Run the tests for this section with:
//
//	go test ./47-sorting-by-functions/
package sortingbyfunctions

// Person is the kind of value that has no natural order: you decide what to
// sort it by.
type Person struct {
	Name string
	Age  int
}

// WordCount says how often a word occurred.
type WordCount struct {
	Word  string
	Count int
}

// SortByLength sorts words in place by their length, shortest first, e.g.
// [peach banana kiwi] becomes [kiwi peach banana]. Words of the same length
// may end up in any order.
//
// Use slices.SortFunc with a comparison function. The function returns a
// negative number when a belongs before b, a positive number when a belongs
// after b and zero when they are equal - cmp.Compare gives exactly that for
// two ints.
func SortByLength(words []string) {
	// TODO: sort with a comparison function that looks at the lengths.
}

// SortDescending sorts nums in place from largest to smallest, e.g.
// [2 9 4] becomes [9 4 2].
//
// Use slices.SortFunc; think about what swapping the arguments of
// cmp.Compare does.
func SortDescending(nums []int) {
	// TODO: sort with a comparison function that reverses the usual order.
}

// SortByAge sorts people in place from youngest to oldest.
func SortByAge(people []Person) {
	// TODO: sort with a comparison function that compares the Age fields.
}

// SortByAgeThenName sorts people in place from youngest to oldest. People of
// the same age are ordered by name (ascending), e.g.
//
//	[{Zoe 30} {Alex 25} {Adam 30}] becomes [{Alex 25} {Adam 30} {Zoe 30}]
//
// Hint: when the first comparison says "equal" (0), fall back to the second
// one. cmp.Or returns the first of its arguments that is not zero.
func SortByAgeThenName(people []Person) {
	// TODO: sort with a comparison function that uses two keys.
}

// WordFrequencies counts how often each word occurs in words and returns one
// WordCount per distinct word, the most frequent word first. Words with the
// same count are ordered alphabetically.
//
// Example: [b a b c a b] -> [{b 3} {a 2} {c 1}].
func WordFrequencies(words []string) []WordCount {
	// TODO: count with a map, turn the map into a slice, sort the slice.
	return nil
}
