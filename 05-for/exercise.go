//go:build !solution

// Package forloop contains exercises for https://gobyexample.com/for
//
// Run the tests for this section with:
//
//	go test ./05-for/
package forloop

// SumTo returns 1 + 2 + ... + n, e.g. SumTo(4) returns 10. For n < 1 the
// result is 0.
//
// Use the classic three-part loop: for i := 1; i <= n; i++.
func SumTo(n int) int {
	if n < 1 {
		return 0
	}

	result := 0
	i := 1

	for i <= n {
		result += i
		i += 1
	}
	return result
}

// Repeat returns s glued together n times: Repeat("ab", 3) returns "ababab"
// and Repeat("ab", 0) returns "".
//
// Use the "do this n times" form: for range n (or for i := range n).
func Repeat(s string, n int) string {
	// TODO: build the result in a loop over range n.
	result := ""
	i := 0

	for ; i < n; i++ {
		result += s
	}
	
	return result
}

// CountDigits returns how many decimal digits the non-negative number n has:
// CountDigits(7) is 1, CountDigits(1234) is 4 and CountDigits(0) is 1.
//
// Use the most basic form of for, with a single condition, and divide n by
// 10 until nothing is left to chop off.
func CountDigits(n int) int {
	// TODO: count with a condition-only loop.
	if n == 0 {
		return 1
	}
	count := 0	
	for n != 0 {
		n = n / 10
		count++;
	}

	return count
}

// SumOdd returns the sum of all odd numbers i with 0 <= i < n, e.g.
// SumOdd(10) returns 1 + 3 + 5 + 7 + 9 = 25.
//
// Loop with for i := range n and use continue to skip the even numbers
// (i%2 == 0 is true for even i).
func SumOdd(n int) int {
	// TODO: range over n, continue on even numbers, add the rest.
	sum := 0

	for i := range n {
		if i % 2 == 0 {
			continue
		}

		sum += i		
	}
	return sum
}

// SmallestDivisor returns the smallest number greater than 1 that divides n
// without a remainder: SmallestDivisor(15) is 3, SmallestDivisor(49) is 7 and
// for a prime number such as 13 it is the number itself. n is at least 2.
//
// Use a for without any condition, try 2, 3, 4, ... and leave the loop with
// break (or return) as soon as you found a divisor.
func SmallestDivisor(n int) int {
	// TODO: loop "forever" and break out at the first divisor.
	divisor := 2
	for n % divisor != 0 {
		divisor++
	}
	return divisor
}

// CollatzSteps counts the steps of the Collatz sequence from n down to 1.
// One step turns n into n/2 when n is even and into 3*n+1 when n is odd.
//
// Example: 6 -> 3 -> 10 -> 5 -> 16 -> 8 -> 4 -> 2 -> 1 takes 8 steps, so
// CollatzSteps(6) returns 8. CollatzSteps(1) returns 0. n is at least 1.
//
// Everything on the page is useful here: a loop, a counter, an if, and
// break or continue.
func CollatzSteps(n int) int {
	// TODO: loop until n is 1, counting the steps.
	steps := 0

	for n != 1 {
		if n % 2 == 0{
			n = n/2
		} else {
			n = 3*n+1
		}
		steps++
	}
	return steps
}
