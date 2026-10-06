//go:build !solution

// Package selectx contains exercises for https://gobyexample.com/select
//
// Run the tests for this section with:
//
//	go test ./33-select/
package selectx

import "time"

// First waits on both channels at once and returns the first message that
// arrives, together with the number of the channel it came from: 1 for c1,
// 2 for c2.
//
// Use a select statement with one case per channel.
func First(c1, c2 <-chan string) (msg string, from int) {
	// TODO: select over the two receives.
	return "", 0
}

// Fastest simulates two concurrent operations, like the page does: it
// starts two goroutines, the first of which sleeps for d1 and then sends
// "one" on a channel of its own, while the second sleeps for d2 and then
// sends "two" on another channel. Fastest returns the message that arrives
// first, without waiting for the slower goroutine.
//
// Example: Fastest(10*time.Millisecond, time.Second) returns "one" after
// about 10ms.
//
// Hint: give the channels a buffer of 1, so that the slower goroutine can
// still complete its send (and end) when nobody receives its message.
func Fastest(d1, d2 time.Duration) string {
	// TODO: two channels, two goroutines, one select.
	return ""
}

// ReceiveN receives n messages in total from c1 and c2, taking each message
// from whichever channel has one ready, and returns them in the order they
// arrived.
//
// Put the select statement into a loop.
func ReceiveN(c1, c2 <-chan string, n int) []string {
	// TODO: loop n times around a select.
	return nil
}

// SendEither sends msg on c1 or on c2, whichever is ready to take it first,
// and returns the number of the channel it used: 1 or 2. The message is sent
// exactly once.
//
// The cases of a select can be sends just as well as receives.
func SendEither(c1, c2 chan<- string, msg string) int {
	// TODO: select over the two sends.
	return 0
}

// SumUntilQuit adds up the numbers arriving on nums and keeps doing so until
// a value arrives on quit. Then it returns the sum.
//
// Example: after 1, 2 and 3 on nums and then a value on quit, it returns 6.
func SumUntilQuit(nums <-chan int, quit <-chan bool) int {
	// TODO: an endless loop around a select; return from the quit case.
	return 0
}
