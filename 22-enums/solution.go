//go:build solution

package enums

import "fmt"

type ServerState int

const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

func (ss ServerState) String() string {
	if name, ok := stateName[ss]; ok {
		return name
	}
	return fmt.Sprintf("ServerState(%d)", int(ss))
}

func Transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	default:
		// StateError stays where it is, and unknown values end up there too.
		return StateError
	}
}

func ParseState(name string) (ServerState, bool) {
	for state, n := range stateName {
		if n == name {
			return state, true
		}
	}
	return 0, false
}

func Trace(start ServerState, steps int) []string {
	names := []string{start.String()}
	s := start
	for range steps {
		s = Transition(s)
		names = append(names, s.String())
	}
	return names
}
