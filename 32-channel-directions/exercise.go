//go:build !solution

// Package channeldirections contains exercises for https://gobyexample.com/channel-directions
//
// Run the tests for this section with:
//
//	go test ./32-channel-directions/
package channeldirections

// Ping sends msg on pings.
//
// The parameter type chan<- string is a send-only channel: inside Ping,
// trying to receive from pings is a compile-time error. (Try it!)
func Ping(pings chan<- string, msg string) {
	// TODO: send the message.
}

// Pong receives one message from pings and sends it on to pongs.
//
// pings is receive-only (<-chan) and pongs is send-only (chan<-), so the
// signature alone tells which way the data flows.
func Pong(pings <-chan string, pongs chan<- string) {
	// TODO: receive, then send.
}

// Generate starts a goroutine that sends nums, in order, on a new channel,
// and returns that channel immediately.
//
// The result type is receive-only: the caller can only read from the
// channel, sending stays the business of the goroutine. A bidirectional
// channel converts to a directional one automatically, so the function can
// simply return the channel it made.
func Generate(nums ...int) <-chan int {
	// TODO: make a channel, send on it from a goroutine, return it.
	return nil
}

// Double receives n numbers from in and sends twice each number on out, in
// the same order. It returns after the n-th number is sent.
func Double(in <-chan int, out chan<- int, n int) {
	// TODO: loop n times: receive, double, send.
}

// Pipeline doubles every number of nums by wiring the functions above
// together: Generate produces the numbers, a goroutine running Double
// transforms them, and Pipeline itself collects the len(nums) results from
// Double's output channel, in order.
//
// Example: Pipeline([]int{1, 2, 3}) returns [2 4 6].
func Pipeline(nums []int) []int {
	// TODO: connect Generate and Double with channels and collect the output.
	return nil
}
