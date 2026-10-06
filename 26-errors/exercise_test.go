package errorsx

import (
	"errors"
	"fmt"
	"testing"
)

// requireSentinels stops a test early while the sentinel errors are still
// nil: errors.Is cannot tell anything apart before they exist.
func requireSentinels(t *testing.T) {
	t.Helper()
	if ErrOutOfTea == nil || ErrPower == nil {
		t.Fatal("ErrOutOfTea and ErrPower are still nil - create them with errors.New first (see TestSentinelErrors)")
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		a, b int
		want int
	}{
		{7, 2, 3},
		{-9, 3, -3},
		{0, 5, 0},
	}
	for _, tt := range tests {
		got, err := Divide(tt.a, tt.b)
		if got != tt.want || err != nil {
			t.Errorf("Divide(%d, %d) = %d, %v, want %d, <nil>", tt.a, tt.b, got, err, tt.want)
		}
	}

	got, err := Divide(1, 0)
	if err == nil {
		t.Fatalf("Divide(1, 0) = %d, <nil>, want an error", got)
	}
	if got != 0 || err.Error() != "division by zero" {
		t.Errorf("Divide(1, 0) = %d, %q, want 0 and the error %q", got, err, "division by zero")
	}
}

func TestSentinelErrors(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want string
	}{
		{"ErrOutOfTea", ErrOutOfTea, "no more tea available"},
		{"ErrPower", ErrPower, "can't boil water"},
	}
	for _, tt := range tests {
		if tt.err == nil {
			t.Errorf("%s is nil, want an error with the message %q", tt.name, tt.want)
			continue
		}
		if got := tt.err.Error(); got != tt.want {
			t.Errorf("%s.Error() = %q, want %q", tt.name, got, tt.want)
		}
	}
}

func TestMakeTea(t *testing.T) {
	requireSentinels(t)

	for _, arg := range []int{0, 1, 3, 5} {
		if err := MakeTea(arg); err != nil {
			t.Errorf("MakeTea(%d) = %q, want <nil>", arg, err)
		}
	}

	if err := MakeTea(2); err != ErrOutOfTea {
		t.Errorf("MakeTea(2) = %v, want ErrOutOfTea itself", err)
	}

	err := MakeTea(4)
	if err == nil {
		t.Fatal("MakeTea(4) = <nil>, want an error wrapping ErrPower")
	}
	if err == ErrPower {
		t.Error("MakeTea(4) returned ErrPower itself, want it wrapped with the context \"making tea\"")
	}
	if !errors.Is(err, ErrPower) {
		t.Errorf("errors.Is(MakeTea(4), ErrPower) = false, want true - wrap with %%w (got the error %q)", err)
	}
	if errors.Is(err, ErrOutOfTea) {
		t.Error("errors.Is(MakeTea(4), ErrOutOfTea) = true, want false")
	}
	if got, want := err.Error(), "making tea: can't boil water"; got != want {
		t.Errorf("MakeTea(4).Error() = %q, want %q", got, want)
	}
}

func TestDescribe(t *testing.T) {
	requireSentinels(t)

	tests := []struct {
		name string
		err  error
		want string
	}{
		{"nil", nil, "ok"},
		{"ErrOutOfTea", ErrOutOfTea, "out of tea"},
		{"ErrPower", ErrPower, "no power"},
		{"wrapped ErrOutOfTea", fmt.Errorf("kitchen: %w", ErrOutOfTea), "out of tea"},
		{"twice wrapped ErrPower", fmt.Errorf("house: %w", fmt.Errorf("kitchen: %w", ErrPower)), "no power"},
		{`errors.New("boom")`, errors.New("boom"), "unknown error: boom"},
		// Same message as the sentinel, but a different error value.
		{`errors.New("can't boil water")`, errors.New("can't boil water"), "unknown error: can't boil water"},
	}
	for _, tt := range tests {
		if got := Describe(tt.err); got != tt.want {
			t.Errorf("Describe(%s) = %q, want %q", tt.name, got, tt.want)
		}
	}
}

func TestBrewAll(t *testing.T) {
	requireSentinels(t)

	n, err := BrewAll([]int{1, 3, 5})
	if n != 3 || err != nil {
		t.Errorf("BrewAll([1 3 5]) = %d, %v, want 3, <nil>", n, err)
	}
	n, err = BrewAll(nil)
	if n != 0 || err != nil {
		t.Errorf("BrewAll(nil) = %d, %v, want 0, <nil>", n, err)
	}

	n, err = BrewAll([]int{1, 3, 4, 5})
	if err == nil {
		t.Fatalf("BrewAll([1 3 4 5]) = %d, <nil>, want an error for cup 2", n)
	}
	if n != 2 {
		t.Errorf("BrewAll([1 3 4 5]) made %d cups, want 2", n)
	}
	if got, want := err.Error(), "cup 2: making tea: can't boil water"; got != want {
		t.Errorf("BrewAll([1 3 4 5]) error = %q, want %q", got, want)
	}
	if !errors.Is(err, ErrPower) {
		t.Error("errors.Is(err, ErrPower) = false for the error of BrewAll([1 3 4 5]), want true - wrap with %w")
	}

	n, err = BrewAll([]int{2, 4})
	if n != 0 || !errors.Is(err, ErrOutOfTea) {
		t.Errorf("BrewAll([2 4]) = %d, %v, want 0 and an error wrapping ErrOutOfTea", n, err)
	} else if got, want := err.Error(), "cup 0: no more tea available"; got != want {
		t.Errorf("BrewAll([2 4]) error = %q, want %q", got, want)
	}
}
