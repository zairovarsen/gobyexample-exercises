//go:build !solution

// Package rangeoveriterators contains exercises for https://gobyexample.com/range-over-iterators
//
// Run the tests for this section with:
//
//	go test ./25-range-over-iterators/
package rangeoveriterators

import "iter"

// Countdown returns an iterator over n, n-1, ..., 1. For n <= 0 it yields
// nothing.
//
// An iter.Seq[int] is a function that takes a callback, by convention called
// yield. The iterator calls yield once for every value, and it has to stop
// as soon as yield returns false - that is how a `break` in the caller's
// range loop reaches the iterator.
//
// Example: `for v := range Countdown(3)` sees 3, 2, 1.
func Countdown(n int) iter.Seq[int] {
	// TODO: return a func(yield func(int) bool) that yields the numbers.
	return nil
}

// List is the generic linked list from the generics section. The type and
// its Push method are complete; only All is left to do.
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

// Push appends v to the end of the list.
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

// All returns an iterator over the elements of the list, in the order they
// were pushed, so that a list can be used directly in a range loop:
//
//	for v := range lst.All() { ... }
//
// All is the conventional name for such a method.
func (lst *List[T]) All() iter.Seq[T] {
	// TODO: return an iterator that walks the list.
	return nil
}

// Fib returns an iterator over the Fibonacci numbers 1, 1, 2, 3, 5, 8, ...
//
// The sequence never ends: the iterator keeps yielding for as long as yield
// returns true, and it is up to the consumer to break out of the loop.
func Fib() iter.Seq[int] {
	// TODO: yield Fibonacci numbers until yield returns false.
	return nil
}

// Take returns an iterator over the first n values of seq (or over all of
// them if seq has fewer). For n <= 0 it yields nothing. It stops right after
// the n-th value, without pulling a value out of seq that it will not pass on.
//
// Iterators compose: Take receives one and returns a new one. It makes an
// endless sequence finite, e.g. Take(Fib(), 5) yields 1, 1, 2, 3, 5.
func Take[T any](seq iter.Seq[T], n int) iter.Seq[T] {
	// TODO: range over seq inside the new iterator and stop after n values.
	return nil
}

// Enumerate pairs every value of seq with its position, counting from 0.
//
// The result is an iter.Seq2, whose yield takes two values, for range loops
// with two variables:
//
//	for i, v := range Enumerate(lst.All()) { ... }
func Enumerate[T any](seq iter.Seq[T]) iter.Seq2[int, T] {
	// TODO: yield (0, first value), (1, second value), ...
	return nil
}

// EvenFibs returns the first n even Fibonacci numbers as a slice, e.g.
// EvenFibs(4) returns [2 8 34 144]. For n <= 0 the result is empty.
//
// Hint: either range over Fib() and break once you have enough, or build a
// small pipeline of iterators and turn it into a slice with slices.Collect.
func EvenFibs(n int) []int {
	// TODO: collect even Fibonacci numbers.
	return nil
}
