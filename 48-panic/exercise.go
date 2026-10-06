//go:build !solution

// Package panics contains exercises for https://gobyexample.com/panic
//
// Run the tests for this section with:
//
//	go test ./48-panic/
package panics

import "errors"

// A panic means something went unexpectedly wrong: a bug in the program or a
// situation it is not prepared to handle. Ordinary, expected failures are
// reported with error return values instead. The functions below are the
// kind of places where a panic is appropriate.

// ErrEmptyStack is the value Stack.Pop panics with when the stack is empty.
var ErrEmptyStack = errors.New("stack is empty")

// Fail always panics, with the string "a problem" as the panic value.
func Fail() {
	// TODO: panic with the message.
}

// MustPositive returns n when n is greater than zero. For any other n it
// panics with a string that names the offending value, built like
// fmt.Sprintf("not positive: %d", n) - e.g. "not positive: -3".
func MustPositive(n int) int {
	// TODO: panic for n <= 0, otherwise return n.
	return 0
}

// Element returns nums[i]. It does no bounds checking of its own: indexing
// outside the slice is a mistake of the caller, and the Go runtime answers
// that with a panic all by itself. (The tests check that the panic comes
// from the runtime.)
func Element(nums []int, i int) int {
	// TODO: return the element.
	return 0
}

// Must is for errors that "cannot happen": it returns value when err is nil
// and panics when err is not nil. The panic value must be err itself (not
// its message), so whoever sees the panic gets the original error.
//
// Example: n := Must(strconv.Atoi("42")) gives 42, while
// Must(strconv.Atoi("x")) panics with the error from Atoi.
func Must[T any](value T, err error) T {
	// TODO: panic with err when it is not nil, otherwise return value.
	var zero T
	return zero
}

// Stack is a last-in-first-out stack of ints. Its zero value is an empty
// stack that is ready to use.
type Stack struct {
	items []int
}

// Push puts n on top of the stack.
func (s *Stack) Push(n int) {
	// TODO: append n to the items.
}

// Len returns the number of values on the stack.
func (s *Stack) Len() int {
	// TODO: return the number of items.
	return 0
}

// Pop removes the value on top of the stack and returns it. Calling Pop on
// an empty stack is a bug in the calling code: Pop panics with ErrEmptyStack
// as the panic value.
func (s *Stack) Pop() int {
	// TODO: panic when the stack is empty, otherwise remove and return the top.
	return 0
}
