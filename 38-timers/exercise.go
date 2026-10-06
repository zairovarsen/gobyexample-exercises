//go:build !solution

// Package timers contains exercises for https://gobyexample.com/timers
//
// Run the tests for this section with:
//
//	go test ./38-timers/
package timers

import "time"

// WaitTimer creates a timer that fires after d, blocks until it does, and
// returns the value that the timer delivered.
//
// A timer made with time.NewTimer has a channel C. When the time is up, the
// current time is sent on C, so a receive from C blocks until the timer
// fires.
func WaitTimer(d time.Duration) time.Time {
	// TODO: time.NewTimer, then receive from its channel C.
	return time.Time{}
}

// Schedule arranges for f to be called once, after the duration d: it
// creates a timer and starts a goroutine that waits for the timer to fire
// and then calls f. Schedule returns the timer immediately, which gives the
// caller the chance to cancel the call with the timer's Stop method.
//
// If the timer is stopped before it fires, f is never called. (The goroutine
// then stays blocked on the timer's channel for good. That is acceptable for
// this exercise; real code would use time.AfterFunc, which does all of this
// without a waiting goroutine.)
func Schedule(d time.Duration, f func()) *time.Timer {
	// TODO: create the timer, wait for it in a goroutine, return it.
	return nil
}

// TryCancel stops the timer and says how that went: "stopped" if the timer
// was stopped before it fired, "already fired" if it was too late.
//
// Hint: Stop returns exactly this information as a bool.
func TryCancel(timer *time.Timer) string {
	// TODO: call Stop and look at its result.
	return ""
}

// Debounce returns a function trigger that delays calls of f until things
// have calmed down: f is called once the duration d has passed since the
// LAST call of trigger. Every call of trigger cancels the countdown that is
// in progress and starts a new one, so a quick burst of triggers leads to a
// single call of f. A trigger that comes after f has run starts a new
// countdown and leads to another call.
//
// (Editors use this to save a file only when the typing pauses.)
//
// trigger will only be called from one goroutine; f may be called from a
// different goroutine.
//
// Hint: keep a *time.Timer in a variable that the returned closure captures.
// Stop the old timer, if there is one, and start a new one - time.AfterFunc
// is the most convenient way.
func Debounce(d time.Duration, f func()) (trigger func()) {
	// TODO: return a closure that restarts a timer.
	return nil
}
