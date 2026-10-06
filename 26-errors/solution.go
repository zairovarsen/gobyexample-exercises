//go:build solution

package errorsx

import (
	"errors"
	"fmt"
)

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

var (
	ErrOutOfTea = errors.New("no more tea available")
	ErrPower    = errors.New("can't boil water")
)

func MakeTea(arg int) error {
	if arg == 2 {
		return ErrOutOfTea
	} else if arg == 4 {
		return fmt.Errorf("making tea: %w", ErrPower)
	}
	return nil
}

func Describe(err error) string {
	switch {
	case err == nil:
		return "ok"
	case errors.Is(err, ErrOutOfTea):
		return "out of tea"
	case errors.Is(err, ErrPower):
		return "no power"
	default:
		return "unknown error: " + err.Error()
	}
}

func BrewAll(args []int) (int, error) {
	for i, arg := range args {
		if err := MakeTea(arg); err != nil {
			return i, fmt.Errorf("cup %d: %w", i, err)
		}
	}
	return len(args), nil
}
