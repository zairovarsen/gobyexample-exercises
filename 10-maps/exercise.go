//go:build !solution

// Package mapsx contains exercises for https://gobyexample.com/maps
//
// (The package is not called "maps" because that is the name of the standard
// library package you will import here.)
//
// Run the tests for this section with:
//
//	go test ./10-maps/
package mapsx

// Squares returns a map from every number 1, 2, ..., n to its square:
// Squares(3) returns map[1:1 2:4 3:9]. For n < 1 the map is empty (but not
// nil).
//
// Create the map with make and set the pairs with m[key] = value.
func Squares(n int) map[int]int {
	// TODO: make the map and fill it in a loop.
	return nil
}

// GetOr returns the value stored under key in m, or fallback when m has no
// such key.
//
// A plain m[key] is not enough, because it returns 0 both for a missing key
// and for a key that really holds 0: GetOr(map[a:0], "a", 5) must return 0
// while GetOr(map[a:0], "b", 5) returns 5. Use the two-value form of the map
// lookup, value, present := m[key].
func GetOr(m map[string]int, key string, fallback int) int {
	// TODO: look the key up with the two-value form.
	return 0
}

// WordCount returns how often each word occurs in words:
// WordCount([a b a]) returns map[a:2 b:1].
//
// Reading a missing key gives the zero value, which makes counting very
// short.
func WordCount(words []string) map[string]int {
	// TODO: count every word.
	return nil
}

// Remove deletes the given keys from m and returns how many keys really were
// removed: a key that is not in the map does not count.
//
// Example: with m = map[a:1 b:2 c:3], Remove(m, [a c x]) returns 2 and
// leaves map[b:2].
//
// The function has no need to return the map: the caller sees the change,
// because a map value refers to the same underlying data when it is passed
// around. Use the builtin delete.
func Remove(m map[string]int, keys []string) int {
	// TODO: check that each key is present, delete it and count.
	return 0
}

// Reset removes every key/value pair from m with the builtin clear, so that
// the caller's map is empty afterwards.
func Reset(m map[string]int) {
	// TODO: clear the map.
}

// SameWords reports whether a and b contain the same words the same number
// of times, in any order: [x y x] and [x x y] are the same words, [x y] and
// [x y y] are not.
//
// Hint: you already have WordCount, and the "maps" package has maps.Equal.
func SameWords(a, b []string) bool {
	// TODO: compare the word counts of a and b.
	return false
}
