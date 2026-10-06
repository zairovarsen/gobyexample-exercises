//go:build !solution

// Package signals contains exercises for https://gobyexample.com/signals
//
// Run the tests for this section with:
//
//	go test ./84-signals/
//
// Real signals are awkward to produce in a test (and Windows cannot send
// them to a process at all), so apart from the first exercise the functions
// receive their signal channel as a parameter and the tests put signals into
// it by hand.
package signals

import (
	"os"
	"time"
)

// NewSignalChannel returns a channel on which the Go runtime will deliver the
// signals SIGINT (ctrl-C) and SIGTERM sent to this program.
//
// The channel must be buffered with room for one signal, because the runtime
// does not wait for a receiver: a signal that cannot be delivered right away
// is dropped.
//
// Hint: signal.Notify from os/signal, with syscall.SIGINT and syscall.SIGTERM.
func NewSignalChannel() chan os.Signal {
	// TODO: make the channel and register it.
	return nil
}

// AwaitSignal starts a goroutine that waits for one signal from sigs, calls
// onSignal with that signal, and then sends true on the returned done
// channel. AwaitSignal itself returns right away.
//
// This is the structure of the program on the page: main would block on
// <-done until the signal arrived and was handled.
func AwaitSignal(sigs <-chan os.Signal, onSignal func(os.Signal)) <-chan bool {
	// TODO: start the goroutine and return the done channel.
	return nil
}

// RunUntilSignal is the main loop of a service that shuts down gracefully.
// It calls work once for every value received from ticks, and stops as soon
// as a signal arrives on sigs. It returns how many times work was called and
// the signal that stopped it.
//
// A real program would pass time.Tick(time.Second) and NewSignalChannel().
//
// Hint: a for loop around a select.
func RunUntilSignal(sigs <-chan os.Signal, ticks <-chan time.Time, work func()) (int, os.Signal) {
	// TODO: loop until a signal arrives.
	return 0, nil
}
