//go:build !solution

// Package channelsynchronization contains exercises for https://gobyexample.com/channel-synchronization
//
// Run the tests for this section with:
//
//	go test ./31-channel-synchronization/
package channelsynchronization

// Worker calls work and, once work has returned, sends true on done to
// notify whoever is waiting that the job is finished.
//
// Worker is meant to be started with `go Worker(work, done)`.
func Worker(work func(), done chan bool) {
	// TODO: do the work, then signal.
}

// Start runs work in a new goroutine and returns immediately. The returned
// channel delivers true once work has finished, so the caller can wait for
// the goroutine with a blocking receive: <-Start(work).
func Start(work func()) chan bool {
	// TODO: make a done channel, start the goroutine, return the channel.
	return nil
}

// WaitAll runs every function of works in a goroutine of its own, all at the
// same time, and returns only when every one of them has finished.
//
// Use a single done channel shared by all goroutines: each one sends on it
// when it is finished, and WaitAll receives as many times as there are
// goroutines. Called without arguments WaitAll returns immediately.
func WaitAll(works ...func()) {
	// TODO: start the goroutines, then receive one signal per goroutine.
}

// RunSequence runs the steps strictly one after the other - step 0 has
// returned before step 1 begins, and so on - although every step runs in a
// goroutine of its own. It returns when the last step has finished.
//
// Challenge: start all the goroutines up front and connect each one to its
// neighbours with channels. A goroutine blocks on a receive until the
// previous step signals that it is done, runs its own step, and then signals
// the next one. Called without arguments RunSequence returns immediately.
func RunSequence(steps ...func()) {
	// TODO: build the chain of goroutines and wait for the last one.
}
