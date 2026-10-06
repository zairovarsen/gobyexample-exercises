//go:build !solution

// Package timeouts contains exercises for https://gobyexample.com/timeouts
//
// Run the tests for this section with:
//
//	go test ./34-timeouts/
package timeouts

import (
	"errors"
	"time"
)

// ErrTimeout is the error for an operation that took too long. It is complete.
var ErrTimeout = errors.New("timeout")

// RecvTimeout waits for a message on ch, but no longer than d. It returns
// the message and true, or "" and false if nothing arrived in time.
//
// Use a select with a second case that receives from time.After(d): that is
// a channel on which a value arrives once the duration d has passed.
func RecvTimeout(ch <-chan string, d time.Duration) (string, bool) {
	// TODO: select over ch and time.After(d).
	return "", false
}

// SendTimeout tries to send msg on ch, but gives up after d. It reports
// whether the message was sent.
func SendTimeout(ch chan<- string, msg string, d time.Duration) bool {
	// TODO: select over the send and time.After(d).
	return false
}

// WithTimeout runs f in a new goroutine and waits for its result, but no
// longer than d. It returns the result and a nil error, or "" and ErrTimeout
// if f did not finish in time.
//
// Example: with an f that needs 2 seconds, WithTimeout(f, time.Second)
// returns "", ErrTimeout after one second.
//
// Hint: let the goroutine deliver the result on a channel with a buffer of
// 1. Then its send succeeds even after a timeout, when nobody is listening
// any more; with an unbuffered channel the goroutine would be stuck forever
// (a goroutine leak).
func WithTimeout(f func() string, d time.Duration) (string, error) {
	// TODO: goroutine + buffered channel + select with a timeout.
	return "", nil
}

// CollectFor gathers the numbers arriving on ch for the duration d and then
// returns them, in order of arrival. It always takes d, no matter how many
// numbers arrive or how they are spaced.
//
// Hint: the deadline is for the whole call, so call time.After once, before
// the loop. Calling it inside the loop would start a new countdown with
// every pass, and a steady stream of numbers would keep the function going
// forever.
func CollectFor(ch <-chan int, d time.Duration) []int {
	// TODO: loop around a select until the deadline fires.
	return nil
}
