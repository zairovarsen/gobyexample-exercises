//go:build !solution

// Package channelbuffering contains exercises for https://gobyexample.com/channel-buffering
//
// Run the tests for this section with:
//
//	go test ./30-channel-buffering/
package channelbuffering

// NewMailbox returns a channel of strings that buffers up to size values:
// that many sends succeed even when nobody is receiving yet.
func NewMailbox(size int) chan string {
	// TODO: make takes the buffer size as its second argument.
	return nil
}

// Preload returns a buffered channel that already holds all the msgs, in
// order, and whose capacity is exactly len(msgs).
//
// No goroutine is needed: because the channel is buffered, the sends
// complete without a concurrent receive.
func Preload(msgs ...string) chan string {
	// TODO: make a channel that is just big enough and send every message.
	return nil
}

// FillUp sends the value v on ch again and again until the buffer of ch is
// full, and returns how many values it sent. It never blocks: when the
// buffer is already full it sends nothing and returns 0.
//
// Assume that no other goroutine uses ch at the same time.
//
// Hint: len(ch) is the number of values waiting in the buffer and cap(ch)
// is the size of the buffer.
func FillUp(ch chan int, v int) int {
	// TODO: send while there is room.
	return 0
}

// KeepLatest sends v on the buffered channel ch without ever blocking: when
// the buffer is full it first makes room by receiving - and so discarding -
// the oldest value. Used like this, a buffered channel keeps the most recent
// cap(ch) values.
//
// Example: on a channel with capacity 3 that holds 1, 2, 3, KeepLatest(ch, 4)
// leaves 2, 3, 4 in the buffer.
//
// Assume that no other goroutine uses ch at the same time.
func KeepLatest(ch chan int, v int) {
	// TODO: drop the oldest value if necessary, then send.
}
