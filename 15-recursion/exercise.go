//go:build !solution

// Package recursion contains exercises for https://gobyexample.com/recursion
//
// Run the tests for this section with:
//
//	go test ./15-recursion/
package recursion

// Factorial returns n! = n * (n-1) * ... * 2 * 1, with Factorial(0) = 1.
// n is never negative in the tests.
//
// No loop: the function calls itself with a smaller n until it reaches the
// base case.
func Factorial(n int) int {
	// TODO: base case, then n times the factorial of n-1.
	if n < 0 {
		return 0
	}

	if n == 0 || n == 1 {
		return 1
	}


	return n * Factorial(n-1)
}

// SumDigits returns the sum of the decimal digits of the non-negative number
// n: SumDigits(123) returns 6. Solve it recursively.
//
// Hint: n%10 is the last digit and n/10 is the number without it.
func SumDigits(n int) int {

	if n == 0 {
		return n
	}

	last := n % 10
	rest := n / 10
	
	return last + SumDigits(rest)
}

// GCD returns the greatest common divisor of two non-negative numbers, i.e.
// the largest number that divides both: GCD(12, 18) returns 6.
//
// Use Euclid's algorithm, which is recursive by nature: the GCD of a and 0
// is a, and otherwise the GCD of a and b equals the GCD of b and a%b.
func GCD(a, b int) int {
	// TODO: translate Euclid's rule.
	return 0
}

// Sum adds up the numbers in nums recursively, without a loop: the sum of an
// empty slice is 0, and the sum of any other slice is its first element plus
// the sum of the rest (nums[1:]).
func Sum(nums []int) int {
	// TODO: first element plus Sum of the rest.
	return 0
}

// Fib returns the n-th Fibonacci number: Fib(0) = 0, Fib(1) = 1 and beyond
// that Fib(n) = Fib(n-1) + Fib(n-2). The tests ask for small n only.
//
// Practise the second technique of the page: write the recursion as a
// closure inside Fib. Because the closure calls itself, its variable must be
// declared before the function is assigned to it:
//
//	var fib func(n int) int
//	fib = func(n int) int { ... fib(n-1) ... }
func Fib(n int) int {
	// TODO: declare the recursive closure and call it.
	return 0
}

// FibFast returns the n-th Fibonacci number as well, but it has to cope with
// n = 90. The simple recursion is hopeless there: it would compute the same
// numbers over and over again, billions of billions of times.
//
// Keep the recursive definition, but remember every result in a map that the
// recursive closure closes over, and look there first before computing
// anything (this is called memoization).
func FibFast(n int) int {
	// TODO: recursive closure + map of results that are already known.
	return 0
}
