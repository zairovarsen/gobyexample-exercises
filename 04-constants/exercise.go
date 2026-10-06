//go:build !solution

// Package constants contains exercises for https://gobyexample.com/constants
//
// Run the tests for this section with:
//
//	go test ./04-constants/
package constants

// Language is a string constant with the value "Go".
//
// TODO: give it the right value.
const Language = ""

// Timeout is a TYPED constant: its type is int64 and its value is 30.
// (A constant declared without a type, like the one below, is "untyped" and
// only gets a type when it is used.)
//
// TODO: add the type and the value.
const Timeout = 0

// These three are given.
const (
	SecondsPerMinute = 60
	MinutesPerHour   = 60
	HoursPerDay      = 24
)

// SecondsPerHour and SecondsPerDay are constant expressions: constants
// computed (at compile time) from the three constants above.
//
// TODO: replace both zeros with an expression - do not type 3600 or 86400.
const (
	SecondsPerHour = 0
	SecondsPerDay  = 0
)

// Zebi and Yobi are given. Both are far too big for any integer type (int64
// ends just below 1 << 63), yet they are perfectly fine as untyped constants:
// constant arithmetic has arbitrary precision.
const (
	Zebi = 1 << 70
	Yobi = 1 << 80
)

// ZebisPerYobi is the number of times Zebi fits into Yobi.
//
// TODO: replace the zero with a constant expression that divides the two
// huge constants. The result is small enough to be used as a normal int.
const ZebisPerYobi = 0

// CircleArea returns the area of a circle with the given radius: pi times
// the radius squared.
//
// Hint: the "math" package has the untyped constant math.Pi.
func CircleArea(radius float64) float64 {
	// TODO: compute the area with math.Pi.
	return 0
}

// Kilo is given. It is untyped, so it takes on whatever numeric type the
// expression around it needs.
const Kilo = 1000

// Kilograms converts grams to kilograms: Kilograms(2500) returns 2.5.
// Use Kilo.
func Kilograms(grams float64) float64 {
	// TODO: divide by Kilo.
	return 0
}

// Grams converts whole kilograms to grams: Grams(3) returns 3000.
// Use the very same Kilo - this time it acts as an int.
func Grams(kilograms int) int {
	// TODO: multiply by Kilo.
	return 0
}

// TotalTimeout returns the time spent waiting when an operation is tried the
// given number of times and each try waits for Timeout:
// TotalTimeout(3) returns 90.
//
// Once Timeout is a typed int64 constant, `retries * Timeout` no longer
// compiles, because Go never mixes int and int64 silently. Give the constant
// the type you need with an explicit conversion, the way the page does it
// with int64(d).
func TotalTimeout(retries int) int {
	// TODO: multiply retries by Timeout, converting Timeout to int.
	return 0
}
