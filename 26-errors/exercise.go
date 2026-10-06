//go:build !solution

// Package errorsx contains exercises for https://gobyexample.com/errors
//
// Run the tests for this section with:
//
//	go test ./26-errors/
package errorsx

// Divide returns a / b (integer division).
//
// Dividing by zero is not possible: in that case Divide returns 0 and an
// error with the message "division by zero". By convention the error is the
// last return value, and nil means that everything went fine.
//
// Hint: errors.New builds an error from a message.
func Divide(a, b int) (int, error) {
	// TODO: return an error when b is 0, the quotient and nil otherwise.
	return 0, nil
}

// Sentinel errors: predeclared error values that stand for one specific
// condition, so that callers can test for exactly that condition.
//
// TODO: these are nil placeholders. Create both with errors.New:
//
//	ErrOutOfTea: "no more tea available"
//	ErrPower:    "can't boil water"
var (
	ErrOutOfTea error
	ErrPower    error
)

// MakeTea tries to make cup number arg:
//
//   - arg 2 fails with ErrOutOfTea, returned as it is;
//   - arg 4 fails with ErrPower wrapped in an error that adds context, so
//     that its message is "making tea: can't boil water";
//   - every other arg succeeds: the result is nil.
//
// Hint: fmt.Errorf with the %w verb wraps an error. The wrapped error stays
// reachable for errors.Is, which a plain %v would not allow.
func MakeTea(arg int) error {
	// TODO: return the right error, or nil.
	return nil
}

// Describe turns an error into a short text:
//
//	nil                          -> "ok"
//	ErrOutOfTea, however wrapped -> "out of tea"
//	ErrPower, however wrapped    -> "no power"
//	anything else                -> "unknown error: " + the error's message
//
// Compare with errors.Is rather than ==, so that wrapped errors are
// recognised too.
func Describe(err error) string {
	// TODO: classify err.
	return ""
}

// BrewAll calls MakeTea for every arg, in order, and stops at the first
// failure. It returns how many cups were made, together with the error of
// the failed cup wrapped as "cup <index>: <error>", where index is the
// position in args. If every cup succeeds the error is nil.
//
// Example: BrewAll([]int{1, 3, 4, 5}) returns 2 and an error with the message
// "cup 2: making tea: can't boil water" for which errors.Is(err, ErrPower)
// still holds - errors can be wrapped any number of times.
func BrewAll(args []int) (int, error) {
	// TODO: loop over args, wrapping the first error with %w.
	return 0, nil
}
