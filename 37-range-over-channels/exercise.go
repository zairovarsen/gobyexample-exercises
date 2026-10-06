//go:build !solution

// Package rangeoverchannels contains exercises for https://gobyexample.com/range-over-channels
//
// Run the tests for this section with:
//
//	go test ./37-range-over-channels/
package rangeoverchannels

// Collect receives values from queue until the channel is closed and
// returns them in order.
//
// Use `for elem := range queue`: the loop receives one value after the
// other and ends on its own when the channel is closed and empty.
func Collect(queue <-chan string) []string {
	// TODO: range over the channel.
	return nil
}

// Count starts a goroutine that sends the numbers 1, 2, ..., n on a new
// unbuffered channel and then closes the channel. Count itself returns the
// channel immediately. For n <= 0 the channel is closed without any value.
//
// The close matters: without it, a range loop over the channel would wait
// forever for a next value after the last one.
func Count(n int) <-chan int {
	// TODO: make a channel; in a goroutine send the numbers and close it.
	return nil
}

// Evens starts a goroutine that passes the even numbers from in on to the
// returned channel and drops the odd ones. When in is closed, the goroutine
// closes the returned channel too. Evens itself returns immediately.
//
// Example: 1, 2, 3, 4 on in gives 2, 4 on the result.
func Evens(in <-chan int) <-chan int {
	// TODO: a goroutine that ranges over in and closes its output at the end.
	return nil
}

// SumOfEvens returns the sum of the even numbers from 1 to n, e.g.
// SumOfEvens(10) = 2 + 4 + 6 + 8 + 10 = 30.
//
// Do not compute it directly: build the pipeline Count -> Evens and add up
// what comes out of it with a range loop.
func SumOfEvens(n int) int {
	// TODO: range over Evens(Count(n)).
	return 0
}

// Merge forwards everything that arrives on a and on b to the returned
// channel, in whatever order the values come in, and closes the returned
// channel once both inputs are closed and drained. Merge itself returns
// immediately.
//
// Challenge: one goroutine per input ranges over it and forwards the values.
// The output must be closed exactly once and only after BOTH goroutines are
// finished, so they have to tell someone - a done channel, as in the channel
// synchronization section, does the job.
func Merge(a, b <-chan int) <-chan int {
	// TODO: two forwarding goroutines plus one that closes the output.
	return nil
}
