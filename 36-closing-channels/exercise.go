//go:build !solution

// Package closingchannels contains exercises for https://gobyexample.com/closing-channels
//
// Run the tests for this section with:
//
//	go test ./36-closing-channels/
package closingchannels

// SendAndClose returns a buffered channel that holds all the nums, in
// order, and that is already closed.
//
// Closing a channel says that no more values will be sent on it. The values
// that are still in the buffer can be received as usual; only after them do
// receivers see that the channel is closed.
func SendAndClose(nums ...int) <-chan int {
	// TODO: make a buffered channel, send the numbers, close it.
	return nil
}

// ReceiveOr receives one value from ch and returns it. If the channel is
// closed and has nothing left to deliver, it returns def instead.
//
// A receive from a closed, empty channel does not block: it yields the zero
// value at once. The two-value form of the receive tells the difference:
//
//	v, ok := <-ch // ok is false when ch is closed and empty
func ReceiveOr(ch <-chan string, def string) string {
	// TODO: use the two-value receive.
	return ""
}

// SumUntilClosed receives numbers from ch until the channel is closed and
// returns their sum and how many there were.
//
// Use an endless loop with the two-value receive and leave it when the
// second value is false. (The next section shows a shorter way with range.)
// Do not stop at a 0 - that is a perfectly good number to send.
func SumUntilClosed(ch <-chan int) (sum, count int) {
	// TODO: loop until the channel is closed.
	return 0, 0
}

// StartWorker starts a goroutine that receives jobs and calls handle for
// each of them, until the jobs channel is closed and all jobs are handled.
// Then the goroutine sends true on the returned channel, to let the caller
// know that it is finished. StartWorker itself returns immediately.
//
// This is the pattern of the page: the sender closes jobs to say "no more
// work", and waits on the done channel for the worker to wrap up.
func StartWorker(jobs <-chan int, handle func(job int)) <-chan bool {
	// TODO: a goroutine with a `j, more := <-jobs` loop and a done channel.
	return nil
}

// Square starts a goroutine that receives numbers from in and sends their
// squares on the returned channel. When in is closed (and all its numbers
// are processed) the goroutine closes the returned channel in turn, so the
// "no more values" signal travels down the pipeline. Square itself returns
// immediately.
//
// Only the sending side should close a channel - here that is the goroutine,
// which is the only sender on the returned channel.
func Square(in <-chan int) <-chan int {
	// TODO: a goroutine that forwards squares and closes its output at the end.
	return nil
}
