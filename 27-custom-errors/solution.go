//go:build solution

package customerrors

import (
	"errors"
	"fmt"
)

type ArgError struct {
	Arg     int
	Message string
}

func (e *ArgError) Error() string {
	return fmt.Sprintf("%d - %s", e.Arg, e.Message)
}

func Process(arg int) (int, error) {
	if arg == 42 {
		return -1, &ArgError{Arg: arg, Message: "can't work with it"}
	}
	return arg + 3, nil
}

func ArgOf(err error) (int, bool) {
	if ae, ok := errors.AsType[*ArgError](err); ok {
		return ae.Arg, true
	}
	return 0, false
}

var (
	ErrEmpty    = errors.New("must not be empty")
	ErrNegative = errors.New("must not be negative")
)

type ValidationError struct {
	Field string
	Err   error
}

func (e *ValidationError) Error() string {
	return e.Field + ": " + e.Err.Error()
}

func (e *ValidationError) Unwrap() error {
	return e.Err
}

func Validate(name string, age int) error {
	if name == "" {
		return &ValidationError{Field: "name", Err: ErrEmpty}
	}
	if age < 0 {
		return &ValidationError{Field: "age", Err: ErrNegative}
	}
	return nil
}
