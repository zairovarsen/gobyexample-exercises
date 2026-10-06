package enums

import (
	"fmt"
	"slices"
	"testing"
)

// ServerState has to implement fmt.Stringer.
var _ fmt.Stringer = StateIdle

// requireDistinctStates stops a test early while the constants still share
// their placeholder value: nothing else can be checked before that is fixed.
func requireDistinctStates(t *testing.T) {
	t.Helper()
	seen := map[ServerState]bool{}
	for _, s := range []ServerState{StateIdle, StateConnected, StateError, StateRetrying} {
		if seen[s] {
			t.Fatal("the State constants do not have distinct values yet - give them values with iota first (see TestStateConstants)")
		}
		seen[s] = true
	}
}

func TestStateConstants(t *testing.T) {
	tests := []struct {
		name string
		got  ServerState
		want int
	}{
		{"StateIdle", StateIdle, 0},
		{"StateConnected", StateConnected, 1},
		{"StateError", StateError, 2},
		{"StateRetrying", StateRetrying, 3},
	}
	for _, tt := range tests {
		if int(tt.got) != tt.want {
			t.Errorf("%s = %d, want %d", tt.name, int(tt.got), tt.want)
		}
	}
}

func TestString(t *testing.T) {
	requireDistinctStates(t)

	tests := []struct {
		state ServerState
		want  string
	}{
		{StateIdle, "idle"},
		{StateConnected, "connected"},
		{StateError, "error"},
		{StateRetrying, "retrying"},
		{ServerState(7), "ServerState(7)"},
		{ServerState(-1), "ServerState(-1)"},
	}
	for _, tt := range tests {
		if got := tt.state.String(); got != tt.want {
			t.Errorf("ServerState(%d).String() = %q, want %q", int(tt.state), got, tt.want)
		}
	}

	// fmt uses the String method automatically.
	if got, want := fmt.Sprint(StateConnected), "connected"; got != want {
		t.Errorf("fmt.Sprint(StateConnected) = %q, want %q", got, want)
	}
	if got, want := fmt.Sprintf("%v then %s", StateError, StateRetrying), "error then retrying"; got != want {
		t.Errorf(`fmt.Sprintf("%%v then %%s", StateError, StateRetrying) = %q, want %q`, got, want)
	}
}

func TestTransition(t *testing.T) {
	requireDistinctStates(t)

	tests := []struct {
		fromName string
		from     ServerState
		wantName string
		want     ServerState
	}{
		{"StateIdle", StateIdle, "StateConnected", StateConnected},
		{"StateConnected", StateConnected, "StateIdle", StateIdle},
		{"StateRetrying", StateRetrying, "StateIdle", StateIdle},
		{"StateError", StateError, "StateError", StateError},
		{"ServerState(42)", ServerState(42), "StateError", StateError},
		{"ServerState(-1)", ServerState(-1), "StateError", StateError},
	}
	for _, tt := range tests {
		if got := Transition(tt.from); got != tt.want {
			t.Errorf("Transition(%s) = %d, want %d (%s)", tt.fromName, int(got), int(tt.want), tt.wantName)
		}
	}
}

func TestParseState(t *testing.T) {
	requireDistinctStates(t)

	tests := []struct {
		name   string
		want   ServerState
		wantOK bool
	}{
		{"idle", StateIdle, true},
		{"connected", StateConnected, true},
		{"error", StateError, true},
		{"retrying", StateRetrying, true},
		{"Idle", 0, false},
		{"bogus", 0, false},
		{"", 0, false},
	}
	for _, tt := range tests {
		got, ok := ParseState(tt.name)
		if got != tt.want || ok != tt.wantOK {
			t.Errorf("ParseState(%q) = %d, %v, want %d, %v", tt.name, int(got), ok, int(tt.want), tt.wantOK)
		}
	}
}

func TestTrace(t *testing.T) {
	requireDistinctStates(t)

	tests := []struct {
		startName string
		start     ServerState
		steps     int
		want      []string
	}{
		{"StateIdle", StateIdle, 0, []string{"idle"}},
		{"StateIdle", StateIdle, 3, []string{"idle", "connected", "idle", "connected"}},
		{"StateRetrying", StateRetrying, 3, []string{"retrying", "idle", "connected", "idle"}},
		{"StateError", StateError, 2, []string{"error", "error", "error"}},
		{"ServerState(9)", ServerState(9), 1, []string{"ServerState(9)", "error"}},
	}
	for _, tt := range tests {
		if got := Trace(tt.start, tt.steps); !slices.Equal(got, tt.want) {
			t.Errorf("Trace(%s, %d) = %q, want %q", tt.startName, tt.steps, got, tt.want)
		}
	}
}
