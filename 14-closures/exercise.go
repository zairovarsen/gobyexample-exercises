//go:build !solution

// Package closures contains exercises for https://gobyexample.com/closures
//
// Run the tests for this section with:
//
//	go test ./14-closures/
package closures

// Counter returns a function that returns 1 the first time it is called, 2
// the second time, 3 the third time and so on.
//
// The count lives in a variable of Counter that the returned anonymous
// function "closes over". Every call of Counter creates a new, independent
// count.
func Counter() func() int {
	// TODO: declare the variable and return a func() int that updates it.
	return nil
}

// Multiplier returns a function that multiplies its argument by factor:
// triple := Multiplier(3) makes triple(5) return 15.
//
// A closure can capture a parameter just like any other variable.
func Multiplier(factor int) func(int) int {
	// TODO: return a func(int) int that uses factor.
	return nil
}

// Accumulator returns a function that keeps a running total. The total
// begins at start; every call adds its argument to the total and returns the
// new total.
//
// Example: acc := Accumulator(10); acc(5) returns 15, then acc(-3) returns
// 12.
func Accumulator(start int) func(delta int) int {
	// TODO: return a closure that adds to a captured total.
	return nil
}

// Fibonacci returns a generator for the Fibonacci numbers: successive calls
// of the returned function return 0, 1, 1, 2, 3, 5, 8, 13, ... (every number
// is the sum of the two before it).
//
// Hint: the closure has to remember two numbers between calls.
func Fibonacci() func() int {
	// TODO: return a closure over two variables.
	return nil
}

// CounterWithReset returns two functions that share one count: next works
// like the function returned by Counter (1, 2, 3, ...), and reset sets the
// count back so that the following call of next returns 1 again.
//
// Two closures created in the same function call see the same variable.
func CounterWithReset() (next func() int, reset func()) {
	// TODO: declare one variable and return two closures that use it.
	return nil, nil
}

// Memoize wraps a (slow) function f in a cache. The returned function gives
// the same results as f, but it calls f at most once for every distinct
// argument: the first time an argument is seen the result is computed with
// f and remembered, afterwards it is served from memory.
//
// Hint: close over a map. A cached result may well be 0, so check whether the
// key is present instead of comparing the value with 0.
func Memoize(f func(int) int) func(int) int {
	// TODO: return a closure that consults a map before calling f.
	return nil
}
