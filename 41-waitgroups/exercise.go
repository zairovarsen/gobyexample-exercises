//go:build !solution

// Package waitgroups contains exercises for https://gobyexample.com/waitgroups
//
// Run the tests for this section with:
//
//	go test ./41-waitgroups/
package waitgroups

import "sync"

// RunAll runs every task in a goroutine of its own, all at the same time,
// and returns when all of them have finished. Called without tasks it
// returns immediately.
//
// Use a sync.WaitGroup: wg.Go(f) starts f in a new goroutine that the group
// keeps track of, and wg.Wait() blocks until all those goroutines are done.
func RunAll(tasks ...func()) {
	// TODO: wg.Go for every task, then wg.Wait.
}

// Worker is written in the classic style from before wg.Go existed, which
// you will meet in a lot of code: the caller does wg.Add(1) and then starts
// `go Worker(id, wg, report)`, and it is the worker's job to call wg.Done()
// when it is finished.
//
// Worker calls report(id) - that is its "work" - and tells the WaitGroup
// that it is done. Use defer for the Done call, so that it cannot be
// forgotten on any path out of the function.
//
// Note that the WaitGroup is passed as a pointer. A WaitGroup must not be
// copied: Done on a copy would never reach the Wait of the original.
func Worker(id int, wg *sync.WaitGroup, report func(id int)) {
	// TODO: defer wg.Done(), then do the work.
}

// MapConcurrently returns a slice holding f(n) for every n of nums, at the
// same positions, with all the calls of f running concurrently: one
// goroutine per number, awaited with a WaitGroup.
//
// The goroutines can store their results directly in a slice that was made
// with the right length beforehand: each goroutine writes only the element
// with its own index, and writes to different elements do not conflict.
func MapConcurrently(nums []int, f func(n int) int) []int {
	// TODO: one goroutine per element, each filling in its own slot.
	return nil
}

// Gather runs every producer in a goroutine of its own. Each producer gets
// the same channel to send on and sends as many numbers as it likes before
// it returns. Gather returns all the numbers that were sent, in whatever
// order they arrive, once all producers have returned.
//
// Challenge: nobody knows in advance how many numbers there will be, so
// Gather has to range over the channel - and somebody has to close the
// channel when, and only when, every producer has returned. Use a WaitGroup
// for the producers and one more goroutine that waits for the group and
// then closes the channel.
func Gather(producers ...func(out chan<- int)) []int {
	// TODO: producers in a WaitGroup, a closer goroutine, and a range loop.
	return nil
}
