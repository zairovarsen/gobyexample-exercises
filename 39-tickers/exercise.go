//go:build !solution

// Package tickers contains exercises for https://gobyexample.com/tickers
//
// Run the tests for this section with:
//
//	go test ./39-tickers/
package tickers

import "time"

// Ticks creates a ticker with the given interval, receives n ticks from it
// and returns the times that they delivered, in order. So it returns after
// about n intervals.
//
// A ticker made with time.NewTicker sends the current time on its channel C
// at regular intervals, again and again. Stop the ticker when you are done
// with it (defer ticker.Stop() is the usual way).
func Ticks(interval time.Duration, n int) []time.Time {
	// TODO: time.NewTicker, then receive n times from its channel C.
	return nil
}

// Every calls f once per interval, in a goroutine that it starts, and
// returns immediately. It hands back a function that the caller calls once
// to end this: stop halts the ticker, makes the goroutine return, and
// guarantees that f is not called any more after stop has returned.
//
// This is the pattern of the page: the goroutine loops around a select with
// one case for the ticker's channel and one for a done channel, and stop
// sends on (or closes) the done channel.
func Every(interval time.Duration, f func()) (stop func()) {
	// TODO: ticker + done channel + goroutine; return a func that stops them.
	return nil
}

// TickFor runs a ticker with the given interval for the duration total and
// returns how many ticks it got in that time. It returns after total.
//
// Example: TickFor(500*time.Millisecond, 1600*time.Millisecond) returns 3,
// like the program on the page that prints three ticks.
//
// Hint: select over the ticker's channel and the channel of time.After(total).
func TickFor(interval, total time.Duration) int {
	// TODO: count ticks until the total time is over.
	return 0
}

// Poll calls check once per interval - the first time after one interval -
// until check returns true or the timeout is over. It reports whether check
// succeeded: true as soon as a call of check returns true, false if the
// timeout passes first. check is only ever called from Poll's own goroutine.
//
// This is how a program waits for something that it cannot be notified
// about, e.g. for a server to come up.
func Poll(interval, timeout time.Duration, check func() bool) bool {
	// TODO: a ticker for the checks, a timer for the timeout, and a select.
	return false
}
