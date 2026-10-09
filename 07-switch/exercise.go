//go:build !solution

// Package switchcase contains exercises for https://gobyexample.com/switch
//
// Run the tests for this section with:
//
//	go test ./07-switch/
package switchcase

import "time"

// NumberWord returns "one", "two" or "three" for the numbers 1, 2 and 3 and
// "many" for every other number.
//
// Use a basic switch on n with a default case.
func NumberWord(n int) string {
	// TODO: switch on n.
	switch {
		case n == 1:
			return "one"	
		case n == 2:
			return "two"
		case n == 3:
			return "three"
		default:
			return "many"
	}
}

// DayKind returns "weekend" for Saturday and Sunday and "weekday" for the
// other five days. A real program would pass time.Now().Weekday().
//
// List several values in one case, separated by commas: the constants are
// called time.Saturday, time.Sunday and so on.
func DayKind(day time.Weekday) string {
	// TODO: one case for the weekend, default for the rest.
	switch day {
	case time.Sunday, time.Saturday:
		return "weekend"
	default:
		return "weekday"
	}
}

// TimeOfDay names the part of the day an hour (0-23) belongs to: "morning"
// before 12, "afternoon" from 12 up to (not including) 18 and "evening" from
// 18 on. For an hour outside 0-23 it returns "invalid".
//
// Use a switch without an expression - it works like an if/else chain, and
// every case holds a condition such as hour < 12.
func TimeOfDay(hour int) string {
	switch {
	case hour < 0, hour > 23:
		return "invalid"
	case hour < 12:
		return "morning"
	case hour >= 18:
		return "evening"
	default:
		return "afternoon"
	}
}

// TypeName returns the name of the type of the value stored in v: "bool",
// "int", "float64" or "string". For a value of any other type (or nil) it
// returns "unknown".
//
// The parameter type any is another name for interface{}: it accepts a value
// of every type. Use a type switch: switch v.(type) { case bool: ... }.
func TypeName(v any) string {
	// TODO: type switch.
	switch v.(type) {
	case bool:
		return "bool"
	case int:
		return "int"
	case float64:
		return "float64"
	case string: 
		return "string"
	default: 
		return "unknown"
	}
}

// Measure gives a value a size, depending on its type:
//
//   - an int is its own size: Measure(42) returns 42
//   - a string is as big as its length: Measure("hello") returns 5
//   - a bool is 1 when true and 0 when false
//   - everything else returns -1
//
// This time you need the value itself inside the cases, with its real type:
// switch t := v.(type) makes t an int in the int case, a string in the
// string case, and so on.
func Measure(v any) int {
	// TODO: type switch that binds the value to a variable.
	switch t := v.(type) {
	case bool: 
		if t {
			return 1
		}
		return 0
	case int:
		return t
	case string:
		return len(t)
	default: 
		return -1
	}
}

// DaysInMonth returns the number of days of a month (1 = January, ...,
// 12 = December) in the given year. February has 29 days in a leap year and
// 28 otherwise; a year is a leap year when it is divisible by 4 but not by
// 100, or divisible by 400. For a month outside 1-12 it returns 0.
//
// Example: DaysInMonth(2, 2024) returns 29, DaysInMonth(4, 2023) returns 30.
func DaysInMonth(month, year int) int {
	// TODO: group the months with comma-separated cases; February needs an if.
	return 0
}
