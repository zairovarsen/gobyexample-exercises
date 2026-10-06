//go:build !solution

// Package epoch contains exercises for https://gobyexample.com/epoch
//
// Run the tests for this section with:
//
//	go test ./58-epoch/
package epoch

import "time"

// The Unix epoch is 1970-01-01 00:00:00 UTC. A "Unix time" is the number of
// seconds (or milliseconds, or nanoseconds) that have passed since then; it
// is the same number everywhere on earth, whatever the time zone.
//
// None of these functions calls time.Now: they get their times as
// parameters so the tests can use fixed moments. A real program would pass
// time.Now() where a "now" is asked for.

// Stamps returns t as the number of seconds, milliseconds and nanoseconds
// since the Unix epoch.
//
// Example: for 2012-10-31 16:13:58.292387 UTC it returns
// 1351700038, 1351700038292 and 1351700038292387000.
//
// time.Time has a method for each of the three.
func Stamps(t time.Time) (sec, milli, nano int64) {
	// TODO: return the three Unix times.
	return 0, 0, 0
}

// FromSeconds converts a number of seconds since the epoch back into a
// time.Time. (The tests compare instants, so the location of the result
// does not matter.)
//
// Example: FromSeconds(1351700038) is 2012-10-31 16:13:58 UTC.
//
// Use time.Unix, which takes seconds and additional nanoseconds.
func FromSeconds(sec int64) time.Time {
	// TODO: convert with time.Unix.
	return time.Time{}
}

// FromNanos converts a number of nanoseconds since the epoch into a
// time.Time.
//
// Example: FromNanos(1351700038292387000) is 2012-10-31 16:13:58.292387 UTC.
func FromNanos(nano int64) time.Time {
	// TODO: convert with time.Unix.
	return time.Time{}
}

// FromMillis converts a number of milliseconds since the epoch - the unit
// JavaScript and many web APIs use - into a time.Time.
//
// Example: FromMillis(1351700038292) is 2012-10-31 16:13:58.292 UTC.
//
// Either split ms into seconds and nanoseconds for time.Unix, or look for a
// ready-made function next to time.Unix.
func FromMillis(ms int64) time.Time {
	// TODO: convert the milliseconds.
	return time.Time{}
}

// Elapsed returns how much time has passed at the moment now since the
// Unix time startSec (in seconds). The result is negative when startSec
// lies after now.
//
// Example: with startSec = 1351700038 and now = 2012-10-31 16:15:08 UTC
// (70 seconds later) it returns 1m10s.
func Elapsed(startSec int64, now time.Time) time.Duration {
	// TODO: turn startSec into a time and subtract.
	return 0
}

// ParseTimestamp converts a Unix timestamp of unknown unit into a
// time.Time. Log files and APIs use seconds, milliseconds, microseconds or
// nanoseconds, and for present-day dates the unit can be told from the size
// of the (non-negative) number:
//
//   - below 1e11  -> seconds
//   - below 1e14  -> milliseconds
//   - below 1e17  -> microseconds
//   - otherwise   -> nanoseconds
//
// Example: 1351700038, 1351700038000, 1351700038000000 and
// 1351700038000000000 all give 2012-10-31 16:13:58 UTC.
func ParseTimestamp(ts int64) time.Time {
	// TODO: pick the unit by magnitude, then convert.
	return time.Time{}
}
