//go:build !solution

// Package recovery contains exercises for https://gobyexample.com/recover
//
// Run the tests for this section with:
//
//	go test ./50-recover/
package recovery

// DidPanic calls f and reports whether f panicked. The panic must not escape
// from DidPanic.
//
// recover only does something when it is called from a deferred function:
// there it stops the panic and returns the panic value (nil when there is no
// panic). Hint: a deferred function literal can set the named result.
func DidPanic(f func()) (panicked bool) {
	// TODO: defer a function that calls recover, then call f.
	return false
}

// PanicValue calls f and returns the value f panicked with, or nil when f
// returned normally.
//
// Example: PanicValue(func() { panic("a problem") }) returns "a problem".
func PanicValue(f func()) (value any) {
	// TODO: recover the panic value in a deferred function.
	return nil
}

// SafeDivide returns a / b. Dividing an int by zero makes the Go runtime
// panic; SafeDivide recovers from that panic and turns it into an ordinary
// error built like fmt.Errorf("recovered: %v", r), where r is the recovered
// value. The result is 0 in that case.
//
// Example: SafeDivide(1, 0) returns 0 and an error with the message
// "recovered: runtime error: integer divide by zero".
//
// Do not check b before dividing - let the panic happen and recover from it.
func SafeDivide(a, b int) (result int, err error) {
	// TODO: divide, and convert a panic into an error with defer + recover.
	return 0, nil
}

// Try is a "safe call" wrapper: it calls f and returns f's error. When f
// panics, Try recovers and returns an error instead of crashing:
//
//   - the panic value is an error e: Try returns an error that wraps e
//     (so errors.Is still finds it), built like fmt.Errorf("panic: %w", e)
//   - any other panic value r: the error is fmt.Errorf("panic: %v", r)
//
// Example: Try(func() error { panic("boom") }) returns an error with the
// message "panic: boom".
func Try(f func() error) (err error) {
	// TODO: call f; recover in a deferred function and set err.
	return nil
}

// RunAll runs every task, one after another in the given order, and returns
// one result per task: errs[i] is nil when tasks[i] returned normally, and
// an error with the message fmt.Sprintf("task %d panicked: %v", i, r) when
// it panicked with the value r.
//
// A panicking task must not stop the tasks after it - this is what a server
// does so that one failing request cannot take down the whole process.
//
// Hint: recover ends the function whose deferred call it runs in, so every
// task needs a function call of its own to recover in.
func RunAll(tasks []func()) []error {
	// TODO: run each task in a helper that recovers.
	return nil
}
