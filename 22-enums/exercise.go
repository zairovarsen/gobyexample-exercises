//go:build !solution

// Package enums contains exercises for https://gobyexample.com/enums
//
// Run the tests for this section with:
//
//	go test ./22-enums/
package enums

// ServerState is an enumerated type: a named int whose possible values are
// the State constants below.
type ServerState int

// The possible states of a server. They must be numbered 0, 1, 2, 3 in this
// order.
//
// TODO: the zeros are placeholders that only make the package compile.
// Declare StateIdle as `ServerState = iota` and let the other constants
// repeat that expression implicitly, so that each one gets the next number.
const (
	StateIdle      ServerState = 0
	StateConnected ServerState = 0
	StateError     ServerState = 0
	StateRetrying  ServerState = 0
)

// String returns the name of the state: "idle", "connected", "error" or
// "retrying". Any other value is formatted like "ServerState(7)".
//
// A String method makes ServerState implement fmt.Stringer, so
// fmt.Println(StateIdle) prints "idle" instead of 0.
//
// Hint: a map[ServerState]string or a switch both work for the names, and
// fmt.Sprintf builds the fallback.
func (ss ServerState) String() string {
	// TODO: return the name of ss.
	return ""
}

// Transition returns the state that follows s:
//
//	idle      -> connected
//	connected -> idle
//	retrying  -> idle
//	error     -> error
//
// A value that is not one of the four constants leads to StateError.
//
// Use a switch on s; a single case can list several states.
func Transition(s ServerState) ServerState {
	// TODO: switch on s.
	return 0
}

// ParseState is the inverse of String: it turns a name back into its state,
// e.g. ParseState("connected") returns StateConnected, true.
//
// The match is case-sensitive. For an unknown name it returns 0, false.
func ParseState(name string) (ServerState, bool) {
	// TODO: look the name up.
	return 0, false
}

// Trace follows Transition for the given number of steps and returns the
// names of all states it passes through, beginning with start itself.
//
// Example: Trace(StateRetrying, 3) returns
// ["retrying" "idle" "connected" "idle"] - always steps+1 names.
func Trace(start ServerState, steps int) []string {
	// TODO: loop, collecting s.String() and moving on with Transition.
	return nil
}
