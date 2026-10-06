//go:build !solution

// Package values contains exercises for https://gobyexample.com/values
//
// Run the tests for this section with:
//
//	go test ./02-values/
package values

import "fmt"

// FullName joins first and last with a single space between them, using the
// + operator for strings: FullName("Rob", "Pike") returns "Rob Pike".
func FullName(first, last string) string {
	// TODO: concatenate the three pieces.
	return first + " " + last
}

// Average returns the mean of two floats: Average(1, 2) returns 1.5.
func Average(a, b float64) float64 {
	// TODO: add and divide.
	return (a + b) / 2
}

// IntDivide returns a / b for integers. Dividing two ints throws the
// fractional part away: IntDivide(7, 2) returns 3, not 3.5.
//
// b is never 0 in the tests.
func IntDivide(a, b int) int {
	// TODO: divide.
	return  a / b
}

// PrintValues prints exactly these six lines to standard output:
//
//	gobyexample
//	6*7 = 42
//	1.0/4.0 = 0.25
//	false
//	true
//	false
//
// Let Go compute every value instead of typing the result: build the first
// line from "go", "by" and "example" with +, pass the arithmetic expressions
// as a second argument to fmt.Println (it puts a space between its
// arguments), and produce the booleans with true && false, true || false and
// !true.
//
// Hint: you need to import the "fmt" package.
func PrintValues() {
	// TODO: six fmt.Println calls.
	fmt.Println("gobyexample")
	fmt.Println("6*7 =", 6 * 7) 
	fmt.Println("1.0/4.0 =", 1.0/4.0) 
	fmt.Println(true && false)
	fmt.Println(false || true)
	fmt.Println(!true)
}

// Xor ("exclusive or") reports whether exactly one of a and b is true.
//
// Build the answer from the boolean operators &&, || and ! only.
func Xor(a, b bool) bool {
	// TODO: combine a and b.
	return (a && !b) || (!a && b)
}

// Majority reports whether at least two of a, b and c are true.
//
// One expression made of && and || is enough.
func Majority(a, b, c bool) bool {
	// TODO: combine a, b and c.
	return (a && b) || (b && c) || (a && c)
}
