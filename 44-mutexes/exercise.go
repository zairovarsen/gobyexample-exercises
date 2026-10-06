//go:build !solution

// Package mutexes contains exercises for https://gobyexample.com/mutexes
//
// Run the tests for this section with:
//
//	go test ./44-mutexes/
package mutexes

// Container holds a set of named counters that any number of goroutines
// may update at the same time. A map is not safe for concurrent use on its
// own, so every access to it has to be guarded by a mutex.
type Container struct {
	// TODO: add a sync.Mutex and a map[string]int.
}

// NewContainer returns an empty Container that is ready for use (a map has
// to be created before values can be stored in it).
//
// It returns a pointer because a struct with a mutex must not be copied:
// the copy would get a mutex of its own, which guards nothing.
func NewContainer() *Container {
	// TODO: create the Container with its map.
	return nil
}

// Inc adds 1 to the counter with the given name. A counter that does not
// exist yet starts at 0.
//
// Lock the mutex before touching the map, and unlock it with defer.
func (c *Container) Inc(name string) {
	// TODO
}

// Get returns the value of the counter with the given name, or 0 if there
// is none. Reading needs the lock too.
func (c *Container) Get(name string) int {
	// TODO
	return 0
}

// Snapshot returns a copy of all counters. The caller may do with the copy
// whatever it likes, without a lock, and changes to the copy do not affect
// the Container (returning the internal map itself would let the caller
// bypass the mutex).
func (c *Container) Snapshot() map[string]int {
	// TODO
	return nil
}

// IncMany starts the given number of goroutines, each of which calls
// c.Inc(name) n times, and returns when all of them are finished (use a
// WaitGroup). Afterwards the counter is goroutines*n higher than before.
func IncMany(c *Container, name string, goroutines, n int) {
	// TODO
}

// Account is a bank account that any number of goroutines may use at the
// same time.
type Account struct {
	// TODO: add a sync.Mutex and the balance.
}

// NewAccount returns an account with the given starting balance.
func NewAccount(balance int) *Account {
	// TODO
	return nil
}

// Deposit adds amount to the balance.
func (a *Account) Deposit(amount int) {
	// TODO
}

// Withdraw takes amount from the account if the balance covers it, and
// reports whether it did. The balance never becomes negative.
//
// The check and the update have to happen under one and the same lock.
// Otherwise two goroutines could both see enough money and then both
// withdraw it.
func (a *Account) Withdraw(amount int) bool {
	// TODO
	return false
}

// Balance returns the current balance.
func (a *Account) Balance() int {
	// TODO
	return 0
}

// Cache remembers the results of expensive computations by key. Any number
// of goroutines may use it at the same time.
type Cache struct {
	// TODO: add a sync.Mutex and a map[string]int.
}

// NewCache returns an empty Cache that is ready for use.
func NewCache() *Cache {
	// TODO
	return nil
}

// GetOrCompute returns the value stored for key. If there is none yet, it
// calls compute, stores the result for key and returns it.
//
// Challenge: compute is called at most once per key, even when many
// goroutines ask for the same missing key at the same moment - the others
// wait and then get the stored value. The simplest correct way is to hold
// the lock for the whole method, including the call of compute.
func (c *Cache) GetOrCompute(key string, compute func() int) int {
	// TODO
	return 0
}
