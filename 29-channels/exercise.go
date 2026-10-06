//go:build !solution

// Package channels contains exercises for https://gobyexample.com/channels
//
// Run the tests for this section with:
//
//	go test ./29-channels/
package channels

// Ping creates an unbuffered channel, starts a goroutine that sends msg on
// it, and returns the channel right away (without waiting for the send).
func Ping(msg string) chan string {
	// TODO: make the channel, send from a new goroutine, return the channel.
	return nil
}

// SumAsync starts a goroutine that adds up nums and delivers the total on the
// returned channel. SumAsync itself must return immediately.
func SumAsync(nums []int) chan int {
	// TODO: compute the sum in a goroutine and send it on the channel.
	return nil
}

// Relay receives exactly one message from in and sends it on out with the
// prefix "relayed: ". It blocks until both the receive and the send happened.
//
// Example: receiving "hi" from in sends "relayed: hi" on out.
func Relay(in, out chan string) {
	// TODO: receive from in, then send on out.
}

// ParallelSquares squares every number in its own goroutine and returns the
// squares in the same order as the input, e.g. [1 2 3] -> [1 4 9].
//
// Use a channel to carry the results back to the calling goroutine - do not
// share a slice between the goroutines. Since goroutines finish in any order,
// each result needs to carry the index it belongs to.
func ParallelSquares(nums []int) []int {
	// TODO: fan the work out to goroutines and collect the results in order.
	return nil
}
