//go:build !solution

// Package timeformattingparsing contains exercises for https://gobyexample.com/time-formatting-parsing
//
// Run the tests for this section with:
//
//	go test ./59-time-formatting-parsing/
package timeformattingparsing

import "time"

// Go describes a time format with an example instead of codes like
// %Y-%m-%d: a layout shows how the reference time
//
//	Mon Jan 2 15:04:05 MST 2006
//
// would be written. "2006-01-02" therefore means year-month-day and
// "3:04PM" a 12-hour clock. The reference time is easy to remember as
// 01/02 03:04:05PM '06 -0700 (1 2 3 4 5 6 7). t.Format(layout) formats a
// time and time.Parse(layout, s) parses a string with it.

// RFC3339 formats t in the RFC 3339 format that most APIs use, e.g.
// "2012-11-01T22:08:41Z" for a UTC time or "2012-11-01T22:08:41+02:00" for a
// time in a zone two hours ahead of UTC.
//
// The time package has constants for common layouts: use time.RFC3339.
func RFC3339(t time.Time) string {
	// TODO: format t.
	return ""
}

// ParseRFC3339 parses a time written in the RFC 3339 format, e.g.
// "2012-11-01T22:08:41+00:00". It returns the error from time.Parse when s
// does not have that format.
func ParseRFC3339(s string) (time.Time, error) {
	// TODO: parse s.
	return time.Time{}, nil
}

// Clock formats the time of day of t on a 12-hour clock with AM/PM, the
// hour without a leading zero, e.g. "6:00AM", "12:30PM" or "10:08PM".
//
// Write a custom layout: how would the reference time (15:04) look in this
// format?
func Clock(t time.Time) string {
	// TODO: format t with a custom layout.
	return ""
}

// LongDate formats the date of t like "Thursday, November 1, 2012": the
// full weekday name, the full month name, the day without a leading zero
// and the four-digit year.
func LongDate(t time.Time) string {
	// TODO: format t with a custom layout.
	return ""
}

// ParseDate parses a calendar date written as year-month-day with four, two
// and two digits, e.g. "2012-11-01", and returns midnight UTC on that day.
// It returns an error when s is not written like that or is no real date
// ("2012-13-01", "2012-02-30").
//
// time.Parse does all of that when given the right layout.
func ParseDate(s string) (time.Time, error) {
	// TODO: parse s with a custom layout.
	return time.Time{}, nil
}

// Reformat converts a timestamp from the European form "01/11/2012 22:08"
// (day/month/year, 24-hour clock) into the form "Nov 1, 2012 at 10:08PM".
// It returns an empty string and the parse error when s is not in the first
// form.
//
// Parse with one layout, format with another.
func Reformat(s string) (string, error) {
	// TODO: parse s, then format the result.
	return "", nil
}

// ParseAny tries to parse s with each of the layouts, in the given order,
// and returns the time from the first layout that fits. When no layout
// fits (or there are no layouts) it returns the zero time.Time and an
// error.
//
// Example: with layouts = [time.RFC3339, "2006-01-02", "Jan 2, 2006"] the
// strings "2012-11-01T22:08:41Z", "2012-11-01" and "Nov 1, 2012" all parse.
func ParseAny(s string, layouts []string) (time.Time, error) {
	// TODO: loop over the layouts and return the first successful parse.
	return time.Time{}, nil
}
