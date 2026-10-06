//go:build !solution

// Package goroutines contains exercises for https://gobyexample.com/goroutines
//
// Run the tests for this section with:
//
//	go test ./28-goroutines/
package goroutines

import "time"

// RunAsync starts f in a new goroutine and returns right away, without
// waiting for f to finish (or even to start).
func RunAsync(f func()) {
	// TODO: use the go keyword.
}

// StartEach calls f once for every name, each call in a goroutine of its
// own, so that all the calls run concurrently. It returns right away without
// waiting for them.
//
// Practise starting a goroutine from an anonymous function that receives
// the name as its argument: go func(name string) { ... }(name)
func StartEach(names []string, f func(name string)) {
	// TODO: start one goroutine per name.
}

// RunAndWait starts every function of fs in a goroutine of its own and then
// sleeps for the duration wait before it returns, to give the goroutines
// time to finish.
//
// Sleeping is the crude way of waiting that the page uses. It only works if
// you know how long the goroutines take; channels and WaitGroups, coming up
// in the next sections, do this properly.
func RunAndWait(wait time.Duration, fs ...func()) {
	// TODO: start the goroutines, then time.Sleep.
}

// Tree starts a tree of goroutines and returns right away.
//
// For depth <= 0 it does nothing. Otherwise it starts ONE new goroutine,
// and that goroutine calls visit(depth) and then calls Tree(depth-1, visit)
// twice. So Tree(3, visit) leads to one call of visit(3), two calls of
// visit(2) and four calls of visit(1), spread over seven goroutines.
//
// The visit function is called from many goroutines at once; it is the
// caller's job to make that safe.
func Tree(depth int, visit func(level int)) {
	// TODO: a goroutine that visits and then recurses.
}
