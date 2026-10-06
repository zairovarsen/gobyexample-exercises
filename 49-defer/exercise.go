//go:build !solution

// Package deferx contains exercises for https://gobyexample.com/defer
//
// Run the tests for this section with:
//
//	go test ./49-defer/
package deferx

import (
	"errors"
	"io"
	"sync"
)

// Errors returned by Account.Withdraw.
var (
	ErrInvalidAmount     = errors.New("amount must be positive")
	ErrInsufficientFunds = errors.New("insufficient funds")
)

// Bracket writes the line "enter\n" to w, calls body, and writes the line
// "exit\n" to w when it is done.
//
// Write the exit line with defer, right after the enter line. A deferred
// call runs when the surrounding function returns - even when it "returns"
// because body panicked. The tests check that too.
func Bracket(w io.Writer, body func()) {
	// TODO: write "enter", defer writing "exit", call body.
}

// ReverseOrder shows in which order deferred calls run. For every name, in
// order, it defers the call record(name). After that loop it calls
// record("body") directly.
//
// Deferred calls run last-in-first-out, so record sees "body" first and then
// the names from last to first: ReverseOrder([a b c], record) results in the
// calls record("body"), record("c"), record("b"), record("a").
func ReverseOrder(names []string, record func(string)) {
	// TODO: defer inside a loop, then call record("body").
}

// Account is a bank account that is safe to use from several goroutines.
type Account struct {
	mu      sync.Mutex
	balance int
}

// Withdraw takes amount out of the account.
//
//   - amount is zero or negative: it returns ErrInvalidAmount
//   - amount is larger than the balance: it returns ErrInsufficientFunds
//   - otherwise it reduces the balance and returns nil
//
// The balance must only be touched while a.mu is locked, and the mutex must
// be unlocked again on every one of the three ways out. Lock it first and
// defer the Unlock right away, so no return path can forget it.
func (a *Account) Withdraw(amount int) error {
	// TODO: lock, defer the unlock, then check and update the balance.
	return nil
}

// WriteAll writes every line, followed by a newline, to wc and closes wc.
//
//   - wc.Close must be called exactly once, whatever happens - use defer.
//   - When a write fails, WriteAll stops writing and returns that error.
//   - When all writes succeed but Close fails, WriteAll returns the error
//     from Close: data may be lost when a close fails, so the error must not
//     be dropped.
//
// Hint: a deferred function literal can assign to a named result (err here)
// after the return statement has run.
func WriteAll(wc io.WriteCloser, lines []string) (err error) {
	// TODO: defer the close (reporting its error), then write the lines.
	return nil
}

// WriteFile creates (or truncates) the file at path, writes every line
// followed by a newline ("\n") to it and closes it. It returns the first
// error that occurs while creating, writing or closing the file.
//
// This is the example from the gobyexample page turned into a function:
// os.Create, a deferred Close whose error is checked, and fmt.Fprintln to
// write. The file must be closed on every return path.
func WriteFile(path string, lines []string) (err error) {
	// TODO: create the file, defer closing it, write the lines.
	return nil
}
