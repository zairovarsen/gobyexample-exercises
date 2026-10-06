//go:build !solution

// Package statefulgoroutines contains exercises for https://gobyexample.com/stateful-goroutines
//
// Run the tests for this section with:
//
//	go test ./45-stateful-goroutines/
package statefulgoroutines

// ReadOp is a request to read the value stored under Key. The goroutine
// that owns the state answers by sending the value on Resp.
type ReadOp struct {
	Key  int
	Resp chan int
}

// WriteOp is a request to store Val under Key. The goroutine that owns the
// state confirms by sending true on Resp.
type WriteOp struct {
	Key  int
	Val  int
	Resp chan bool
}

// StartStore starts a goroutine that owns a map[int]int - no other
// goroutine ever touches that map - and returns the two channels through
// which everybody else asks it to read and to write. Because only one
// goroutine accesses the state, no mutex is needed: the access is
// synchronised by the channels.
//
// The goroutine loops forever around a select:
//
//   - for a ReadOp from reads it sends the value stored under the key (0 if
//     there is none) on the Resp channel of the request;
//   - for a WriteOp from writes it stores the value and sends true on the
//     Resp channel of the request.
func StartStore() (reads chan<- ReadOp, writes chan<- WriteOp) {
	// TODO: make both channels, start the state-owning goroutine, return them.
	return nil, nil
}

// Read asks the store behind the reads channel for the value under key and
// returns the answer: it builds a ReadOp with a fresh Resp channel, sends
// it, and waits for the response.
func Read(reads chan<- ReadOp, key int) int {
	// TODO: send a request, receive the response.
	return 0
}

// Write asks the store behind the writes channel to store val under key
// and returns once the store has confirmed it.
func Write(writes chan<- WriteOp, key, val int) {
	// TODO: send a request, wait for the confirmation.
}

// CountOps is the program from the page as a function, without the
// sleeping: it starts a store with StartStore, then `readers` goroutines
// that each do opsEach reads with Read, and `writers` goroutines that each
// do opsEach writes with Write (use any keys and values you like). Every
// completed operation is counted in one of two atomic counters. When all
// goroutines are finished (use a WaitGroup) it returns the two totals.
//
// Example: CountOps(10, 3, 20) returns 200, 60.
func CountOps(readers, writers, opsEach int) (readOps, writeOps uint64) {
	// TODO: combine the store with atomic counters and a WaitGroup.
	return 0, 0
}

// Counter is a counter that any number of goroutines may use at the same
// time - without a mutex and without sync/atomic. Its count is a local
// variable of a goroutine that NewCounter starts, and the methods talk to
// that goroutine through channels.
type Counter struct {
	// TODO: add the channels that the methods need to reach the goroutine,
	// directly as fields of this struct. (The tests check that Counter has
	// channel fields and nothing from the packages sync and sync/atomic.)
}

// NewCounter returns a Counter at 0 and starts the goroutine that owns its
// count.
func NewCounter() *Counter {
	// TODO: make the channels and start the goroutine.
	return nil
}

// Inc adds 1 to the counter. When Inc returns the increment has reached
// the owning goroutine, so a later call of Value includes it.
func (c *Counter) Inc() {
	// TODO
}

// Value returns the current count. As the count lives in another goroutine,
// Value has to ask for it and wait for the answer (remember ReadOp).
func (c *Counter) Value() int {
	// TODO
	return 0
}

// Close ends the goroutine of the counter. It is called once, when the
// counter is not needed any more; no other method is called after it.
func (c *Counter) Close() {
	// TODO: tell the goroutine to return, e.g. by closing a channel.
}
