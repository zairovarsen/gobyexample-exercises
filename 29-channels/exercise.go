//go:build !solution

// Package channels contains exercises for https://gobyexample.com/channels
//
// Run the tests for this section with:
//
//	go test ./29-channels/
package channels

import "fmt"

// Ping creates an unbuffered channel, starts a goroutine that sends msg on
// it, and returns the channel right away (without waiting for the send).
func Ping(msg string) chan string {
	// TODO: make the channel, send from a new goroutine, return the channel.
	messages := make(chan string)

	go func(msg string) {
		messages <- msg
	}(msg)

	return messages
}

// SumAsync starts a goroutine that adds up nums and delivers the total on the
// returned channel. SumAsync itself must return immediately.
func SumAsync(nums []int) chan int {
	// TODO: compute the sum in a goroutine and send it on the channel.
	numbers := make(chan int)

	go func() {
		sum := 0
		for _, num := range nums {
			sum += num
		}
		numbers <- sum
	}()

	return numbers
}

// Relay receives exactly one message from in and sends it on out with the
// prefix "relayed: ". It blocks until both the receive and the send happened.
//
// Example: receiving "hi" from in sends "relayed: hi" on out.
func Relay(in, out chan string) {
	// TODO: receive from in, then send on out.
	msg := <-in
	out <- fmt.Sprintf("relayed: %s", msg)
}

// ParallelSquares squares every number in its own goroutine and returns the
// squares in the same order as the input, e.g. [1 2 3] -> [1 4 9].
//
// Use a channel to carry the results back to the calling goroutine - do not
// share a slice between the goroutines. Since goroutines finish in any order,
// each result needs to carry the index it belongs to.
func ParallelSquares(nums []int) []int {
	type result struct {
		index, square int
	}
	results := make(chan result)

	go func() {
		for i, num := range nums {
			results <- result{index: i, square: num * num}
		}
	}()

	squares := make([]int, len(nums))
	for range nums {
		result := <-results
		squares[result.index] = result.square
	}

	return squares
}
