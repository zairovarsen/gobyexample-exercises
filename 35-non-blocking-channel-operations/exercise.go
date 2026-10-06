//go:build !solution

// Package nonblockingchanneloperations contains exercises for https://gobyexample.com/non-blocking-channel-operations
//
// Run the tests for this section with:
//
//	go test ./35-non-blocking-channel-operations/
package nonblockingchanneloperations

// TryReceive receives from messages only if a value is available right now.
// It returns the value and true, or "" and false when it would have to wait.
// It never blocks.
//
// Use a select with a default case: default is chosen when no other case is
// ready.
func TryReceive(messages <-chan string) (string, bool) {
	// TODO: select with a receive case and a default case.
	return "", false
}

// TrySend sends msg on messages only if that is possible right now - a
// receiver is waiting or the buffer has room - and reports whether the
// message was sent. It never blocks.
func TrySend(messages chan<- string, msg string) bool {
	// TODO: select with a send case and a default case.
	return false
}

// Poll looks at two channels at once without blocking and reports what it
// found:
//
//	a message is available -> "received message <msg>"
//	a signal is available  -> "received signal <true or false>"
//	neither                -> "no activity"
//
// If both are available it may report either one.
//
// Use a multi-way select: several cases plus a default.
func Poll(messages <-chan string, signals <-chan bool) string {
	// TODO: select with two receive cases and a default case.
	return ""
}

// DrainAvailable receives from ch for as long as values are available
// right now and returns them in order. As soon as it would have to wait it
// stops; it never blocks. Without any available value the result is empty.
func DrainAvailable(ch <-chan int) []int {
	// TODO: a loop around a non-blocking receive.
	return nil
}

// Broadcast offers msg to every subscriber with a non-blocking send, so
// that a subscriber who is not ready (nobody is receiving and the buffer, if
// there is one, is full) is skipped instead of holding up everybody else.
// It returns the number of subscribers that got the message.
func Broadcast(subscribers []chan string, msg string) int {
	// TODO: a non-blocking send to each subscriber, counting the successes.
	return 0
}
