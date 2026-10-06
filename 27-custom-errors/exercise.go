//go:build !solution

// Package customerrors contains exercises for https://gobyexample.com/custom-errors
//
// Run the tests for this section with:
//
//	go test ./27-custom-errors/
package customerrors

import "errors"

// ArgError is a custom error type. Unlike a plain message it carries data
// that the caller can inspect: the argument that caused the problem.
type ArgError struct {
	Arg     int
	Message string
}

// Error makes *ArgError implement the built-in error interface. It returns
// "<Arg> - <Message>", e.g. "42 - can't work with it".
func (e *ArgError) Error() string {
	// TODO: format the message.
	return ""
}

// Process returns arg + 3 and a nil error - except for 42, which it refuses
// to work with: then it returns -1 and an *ArgError with Arg 42 and the
// message "can't work with it".
func Process(arg int) (int, error) {
	// TODO: return a pointer to an ArgError for 42.
	return 0, nil
}

// ArgOf looks for an *ArgError in err or in any error that err wraps. If
// there is one it returns its Arg and true, otherwise 0 and false (also for
// a nil err).
//
// Hint: errors.As is the counterpart of errors.Is for error types: it
// searches the chain for an error of the wanted type and hands it to you.
// (errors.AsType does the same with a type parameter.)
func ArgOf(err error) (int, bool) {
	// TODO: extract the *ArgError.
	return 0, false
}

// Sentinel errors saying what is wrong with a value. They are complete.
var (
	ErrEmpty    = errors.New("must not be empty")
	ErrNegative = errors.New("must not be negative")
)

// ValidationError is a custom error that says which field is invalid and
// wraps an underlying error that says why.
type ValidationError struct {
	Field string
	Err   error
}

// Error returns "<Field>: <message of Err>", e.g. "age: must not be negative".
func (e *ValidationError) Error() string {
	// TODO: format the message.
	return ""
}

// Unwrap returns the underlying error. This method is what lets errors.Is
// and errors.As look through a ValidationError, the same way they look
// through an error made with fmt.Errorf and %w.
func (e *ValidationError) Unwrap() error {
	// TODO: return the wrapped error.
	return nil
}

// Validate checks the data of a person, the name first:
//
//   - an empty name gives a *ValidationError with Field "name" wrapping ErrEmpty;
//   - a negative age gives a *ValidationError with Field "age" wrapping ErrNegative;
//   - otherwise the result is nil.
//
// Take care to return a plain nil for valid data. A nil *ValidationError
// stored in the error result would not compare equal to nil.
func Validate(name string, age int) error {
	// TODO: return the first problem found, or nil.
	return nil
}
