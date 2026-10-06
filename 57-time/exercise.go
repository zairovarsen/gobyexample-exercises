//go:build !solution

// Package timex contains exercises for https://gobyexample.com/time
//
// Run the tests for this section with:
//
//	go test ./57-time/
package timex

import "time"

// A note on testability: none of these functions calls time.Now. They get
// the times to work with as parameters, so the tests can use fixed moments.
// A real program would pass time.Now() where a "now" is asked for.

// Noon returns the moment 12:00:00 (and zero nanoseconds) on the given day
// in the UTC time zone.
//
// Example: Noon(2009, time.November, 17) is 2009-11-17 12:00:00 UTC.
//
// Use time.Date; its last argument is the location, time.UTC here.
func Noon(year int, month time.Month, day int) time.Time {
	// TODO: build the time.
	return time.Time{}
}

// Describe returns a description of t in the form
// "<weekday> <day> <month> <year> <hh>:<mm>", built from the components of
// t (the Weekday, Day, Month, Year, Hour and Minute methods). Hour and
// minute always have two digits.
//
// Example: for 2009-11-17 20:34:58 it returns "Tuesday 17 November 2009 20:34".
//
// Hint: time.Weekday and time.Month print as English names with %s or %v.
// (The next sections show t.Format, which is the usual tool for this job;
// here the point is to practise the component methods.)
func Describe(t time.Time) string {
	// TODO: extract the components and format them with fmt.Sprintf.
	return ""
}

// Order compares two moments and returns "before" when a happens before b,
// "after" when a happens after b and "equal" when they are the same instant.
//
// Use the Before, After and Equal methods. Do not compare with ==: two
// time.Time values can describe the same instant in different time zones.
func Order(a, b time.Time) string {
	// TODO: compare a and b.
	return ""
}

// HoursBetween returns the time from start to end in hours, as a floating
// point number. It is negative when end is before start.
//
// Example: from 09:00 to 17:30 on the same day it returns 8.5.
//
// Use the Sub method, which returns a time.Duration, and one of the
// Duration methods.
func HoursBetween(start, end time.Time) float64 {
	// TODO: subtract and convert the duration to hours.
	return 0
}

// Shift moves t by d in both directions: later is t advanced by d, earlier
// is t moved back by d.
//
// Example: for 12:00 and d = 90 minutes it returns 13:30 and 10:30.
//
// Use the Add method; a negative duration moves back in time.
func Shift(t time.Time, d time.Duration) (later, earlier time.Time) {
	// TODO: add d and -d.
	return time.Time{}, time.Time{}
}

// Age returns the age in completed years, on the day of now, of somebody
// born on the day of born. Only the dates (year, month, day) matter; the
// birthday itself already counts as a completed year.
//
// Examples for somebody born on 1990-06-15:
//
//	now = 2020-06-14 -> 29
//	now = 2020-06-15 -> 30
//	now = 2021-01-01 -> 30
//
// Hint: start with the difference of the years, then check with Month and
// Day whether this year's birthday is still ahead.
func Age(born, now time.Time) int {
	// TODO: compute the completed years from the date components.
	return 0
}
