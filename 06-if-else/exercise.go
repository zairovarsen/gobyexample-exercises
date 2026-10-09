//go:build !solution

// Package ifelse contains exercises for https://gobyexample.com/if-else
//
// Run the tests for this section with:
//
//	go test ./06-if-else/
package ifelse

// EvenOrOdd returns "even" when n is divisible by 2 and "odd" otherwise.
func EvenOrOdd(n int) string {
	// TODO: decide with if / else.
	if n % 2 == 0 {
		return "even"
	}
	return "odd"
}

// Compare returns "less" when a < b, "equal" when a == b and "greater" when
// a > b.
//
// Practise the statement that may precede the condition: compute the
// difference inside the if (if diff := a - b; diff < 0 { ... }) and reuse
// diff in the else if branch.
func Compare(a, b int) string {
	// TODO: use an if / else if / else chain with a short statement.

	if diff := a -b; diff < 0 {
		return "less"
	} else if diff > 0  {
		return "greater"
	} else {
		return "equal"
	}
}

// LeapYear reports whether year is a leap year: a year divisible by 4 is a
// leap year, except that years divisible by 100 are not, except that years
// divisible by 400 are after all. So 2024 and 2000 are leap years, 2023 and
// 1900 are not.
//
// Combine the conditions with && and ||.
func LeapYear(year int) bool {
	return year%4 == 0 && year%100 != 0 || year%400 == 0
}

// Grade turns a score from 0 to 100 into a letter: 90 and above is "A", 80 to
// 89 is "B", 70 to 79 is "C", 60 to 69 is "D" and everything below is "F".
// A score below 0 or above 100 is "invalid".
func Grade(score int) string {
	if score >= 90 {
		return "A"
	} else if score >= 80 && score < 90 {
		return "B"
	} else if score >= 70 && score < 80 {
		return "C"
	} else if score >= 60 && score < 70 {
		return "D"
	} else {
		return "F"
	}
}

// FizzBuzz returns "Fizz" when n is divisible by 3, "Buzz" when n is
// divisible by 5, "FizzBuzz" when it is divisible by both, and otherwise the
// number itself as a string: FizzBuzz(7) returns "7".
//
// Hint: strconv.Itoa(n) from the "strconv" package turns an int into its
// decimal string.
func FizzBuzz(n int) string {
	// TODO: pick the right answer - the order of the checks matters.
	return ""
}

// FizzBuzzUpTo plays the whole game: it returns the FizzBuzz answers for 1,
// 2, ..., n in one string, separated by single spaces (no space at the start
// or at the end). For n < 1 it returns "".
//
// Example: FizzBuzzUpTo(5) returns "1 2 Fizz 4 Buzz".
func FizzBuzzUpTo(n int) string {
	// TODO: loop from 1 to n, call FizzBuzz and join the answers with spaces.
	return ""
}
