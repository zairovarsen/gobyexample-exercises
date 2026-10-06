//go:build !solution

// Package atomiccounters contains exercises for https://gobyexample.com/atomic-counters
//
// Run the tests for this section with:
//
//	go test ./43-atomic-counters/
package atomiccounters

// CountOps is the program from the page as a function: it starts the given
// number of goroutines, each of which increments one shared counter
// perGoroutine times. When all goroutines are done (use a WaitGroup) it
// returns the value of the counter.
//
// The counter has to be an atomic.Uint64 from the package sync/atomic, with
// its methods Add and Load. With a plain integer and ops++ the goroutines
// would get in each other's way, and increments would be lost.
//
// Example: CountOps(50, 1000) returns 50000.
func CountOps(goroutines, perGoroutine int) uint64 {
	// TODO: atomic counter + WaitGroup.
	return 0
}

// Counter is a counter that any number of goroutines may use at the same
// time. Its zero value is a counter at 0, ready for use.
type Counter struct {
	// TODO: add a field of type atomic.Int64 for the count. (The tests check
	// that the struct keeps its state in a type of sync/atomic.)
}

// Inc adds 1 to the counter.
func (c *Counter) Inc() {
	// TODO
}

// Add adds delta, which may be negative, to the counter.
func (c *Counter) Add(delta int64) {
	// TODO
}

// Value returns the current count. Reading while other goroutines might be
// writing has to be atomic as well: use Load.
func (c *Counter) Value() int64 {
	// TODO
	return 0
}

// Flag is a switch that can be turned on once, for example to mark that a
// shutdown has begun. Any number of goroutines may use it at the same time.
// Its zero value is a flag that is not set.
type Flag struct {
	// TODO: add a field of type atomic.Bool.
}

// TrySet sets the flag. It returns true if this call was the one that
// changed the flag from unset to set, and false if the flag was set before.
// However many goroutines call TrySet at the same moment, exactly one of
// them gets true.
//
// Hint: checking the flag and then storing true are two steps, and another
// goroutine can get in between. CompareAndSwap(false, true) does both in a
// single atomic step and reports whether it made the swap.
func (f *Flag) TrySet() bool {
	// TODO
	return false
}

// IsSet reports whether the flag has been set.
func (f *Flag) IsSet() bool {
	// TODO
	return false
}

// HighWater remembers the largest value it has seen, for example the peak
// number of connections. Any number of goroutines may use it at the same
// time. Its zero value is ready for use and has the maximum 0.
type HighWater struct {
	// TODO: add a field of type atomic.Int64.
}

// Observe records v: if v is larger than the current maximum it becomes the
// new maximum, otherwise nothing changes.
//
// Challenge: "load, compare, store" is a race, because another goroutine
// may store a larger value between your Load and your Store, and you would
// overwrite it. Use a loop: load the current maximum; if v is not larger,
// return; otherwise try CompareAndSwap(current, v) and start over if it
// reports that the value has changed in the meantime.
func (h *HighWater) Observe(v int64) {
	// TODO
}

// Max returns the largest value observed so far (0 if there was none above 0).
func (h *HighWater) Max() int64 {
	// TODO
	return 0
}
