//go:build !solution

// Package randomnumbers contains exercises for https://gobyexample.com/random-numbers
//
// Run the tests for this section with:
//
//	go test ./60-random-numbers/
package randomnumbers

import "math/rand/v2"

// All exercises use the "math/rand/v2" package (note the /v2 - the older
// "math/rand" has a slightly different API). Its top-level functions such as
// rand.IntN use a generator that is seeded randomly when the program starts.
// For a reproducible sequence, create your own generator from a seeded
// source: rand.New(rand.NewPCG(seed1, seed2)) has the same methods.

// RollDie returns a random number from 1 to 6, each with the same
// probability.
//
// Use rand.IntN(n), which returns a random int with 0 <= result < n.
func RollDie() int {
	// TODO: return a random number between 1 and 6.
	return 0
}

// FloatBetween returns a random float64 f with lo <= f < hi, evenly
// distributed over that range.
//
// Use rand.Float64, which returns a number with 0.0 <= f < 1.0; stretch and
// shift it to get other ranges (the gobyexample page does it for 5.0 to 10.0).
func FloatBetween(lo, hi float64) float64 {
	// TODO: scale rand.Float64 to the range.
	return 0
}

// SeededInts returns n pseudo-random numbers, each with 0 <= number < 100,
// produced by a new generator that is seeded with seed1 and seed2. The same
// seeds must give the same numbers every time; other seeds give other
// numbers.
//
// Create the generator with rand.New(rand.NewPCG(seed1, seed2)) and call its
// IntN method.
func SeededInts(seed1, seed2 uint64, n int) []int {
	// TODO: create a seeded generator and draw n numbers from it.
	return nil
}

// Pick returns a random element of items, each with the same probability,
// using the generator r that the caller provides (instead of the package's
// global one). That way the caller decides whether the choice is
// reproducible. For an empty slice it returns "".
func Pick(r *rand.Rand, items []string) string {
	// TODO: pick a random index with r.
	return ""
}

// Shuffled returns a new slice with the elements of items in a random
// order, produced by a generator seeded with rand.NewPCG(seed, seed). The
// same seed must give the same order every time. The items slice itself
// must not be modified.
//
// Write the shuffle yourself with IntN (the Fisher-Yates shuffle): walk
// from the last index i down to 1 and swap element i with the element at a
// random index j, where 0 <= j <= i.
func Shuffled(seed uint64, items []int) []int {
	// TODO: copy items, then shuffle the copy with a seeded generator.
	return nil
}
