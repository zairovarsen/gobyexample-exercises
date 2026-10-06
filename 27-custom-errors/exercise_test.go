package customerrors

import (
	"errors"
	"fmt"
	"testing"
)

// Both custom types have to implement error (with pointer receivers).
var (
	_ error = (*ArgError)(nil)
	_ error = (*ValidationError)(nil)
)

func TestArgErrorError(t *testing.T) {
	tests := []struct {
		err  *ArgError
		want string
	}{
		{&ArgError{Arg: 42, Message: "can't work with it"}, "42 - can't work with it"},
		{&ArgError{Arg: -1, Message: "too small"}, "-1 - too small"},
	}
	for _, tt := range tests {
		if got := tt.err.Error(); got != tt.want {
			t.Errorf("(&ArgError{%d, %q}).Error() = %q, want %q", tt.err.Arg, tt.err.Message, got, tt.want)
		}
	}
}

func TestProcess(t *testing.T) {
	for _, arg := range []int{7, 0, -3} {
		got, err := Process(arg)
		if got != arg+3 || err != nil {
			t.Errorf("Process(%d) = %d, %v, want %d, <nil>", arg, got, err, arg+3)
		}
	}

	got, err := Process(42)
	if err == nil {
		t.Fatalf("Process(42) = %d, <nil>, want an *ArgError", got)
	}
	if got != -1 {
		t.Errorf("Process(42) returned the value %d, want -1", got)
	}
	var ae *ArgError
	if !errors.As(err, &ae) {
		t.Fatalf("Process(42) returned an error of type %T, want *ArgError", err)
	}
	if ae.Arg != 42 || ae.Message != "can't work with it" {
		t.Errorf("Process(42) returned &ArgError{%d, %q}, want &ArgError{42, %q}", ae.Arg, ae.Message, "can't work with it")
	}
}

func TestArgOf(t *testing.T) {
	tests := []struct {
		name   string
		err    error
		want   int
		wantOK bool
	}{
		{"&ArgError{5, ...}", &ArgError{Arg: 5, Message: "five"}, 5, true},
		{"wrapped &ArgError{9, ...}", fmt.Errorf("context: %w", &ArgError{Arg: 9, Message: "nine"}), 9, true},
		{
			"twice wrapped &ArgError{-2, ...}",
			fmt.Errorf("outer: %w", fmt.Errorf("inner: %w", &ArgError{Arg: -2, Message: "minus two"})),
			-2, true,
		},
		{`errors.New("plain")`, errors.New("plain"), 0, false},
		{"&ValidationError{...}", &ValidationError{Field: "age", Err: ErrNegative}, 0, false},
		{"nil", nil, 0, false},
	}
	for _, tt := range tests {
		got, ok := ArgOf(tt.err)
		if got != tt.want || ok != tt.wantOK {
			t.Errorf("ArgOf(%s) = %d, %v, want %d, %v", tt.name, got, ok, tt.want, tt.wantOK)
		}
	}
}

func TestValidationError(t *testing.T) {
	ve := &ValidationError{Field: "age", Err: ErrNegative}
	if got, want := ve.Error(), "age: must not be negative"; got != want {
		t.Errorf("(&ValidationError{%q, ErrNegative}).Error() = %q, want %q", "age", got, want)
	}
	if got := ve.Unwrap(); got != ErrNegative {
		t.Errorf("(&ValidationError{%q, ErrNegative}).Unwrap() = %v, want ErrNegative", "age", got)
	}
	// errors.Is follows Unwrap.
	if !errors.Is(ve, ErrNegative) {
		t.Error("errors.Is(&ValidationError{..., ErrNegative}, ErrNegative) = false, want true")
	}
	if errors.Is(ve, ErrEmpty) {
		t.Error("errors.Is(&ValidationError{..., ErrNegative}, ErrEmpty) = true, want false")
	}
}

func TestValidate(t *testing.T) {
	valid := []struct {
		name string
		age  int
	}{
		{"gopher", 14},
		{"baby", 0},
	}
	for _, tt := range valid {
		if err := Validate(tt.name, tt.age); err != nil {
			t.Errorf("Validate(%q, %d) = %v, want <nil>", tt.name, tt.age, err)
		}
	}

	invalid := []struct {
		name      string
		age       int
		wantField string
		wantErr   error
		wantMsg   string
	}{
		{"", 14, "name", ErrEmpty, "name: must not be empty"},
		{"gopher", -1, "age", ErrNegative, "age: must not be negative"},
		{"", -1, "name", ErrEmpty, "name: must not be empty"},
	}
	for _, tt := range invalid {
		err := Validate(tt.name, tt.age)
		if err == nil {
			t.Errorf("Validate(%q, %d) = <nil>, want a *ValidationError for the field %q", tt.name, tt.age, tt.wantField)
			continue
		}
		var ve *ValidationError
		if !errors.As(err, &ve) {
			t.Errorf("Validate(%q, %d) returned an error of type %T, want *ValidationError", tt.name, tt.age, err)
			continue
		}
		if ve.Field != tt.wantField {
			t.Errorf("Validate(%q, %d) reports the field %q, want %q", tt.name, tt.age, ve.Field, tt.wantField)
		}
		if !errors.Is(err, tt.wantErr) {
			t.Errorf("Validate(%q, %d): errors.Is(err, %v) = false, want true", tt.name, tt.age, tt.wantErr)
		}
		if got := err.Error(); got != tt.wantMsg {
			t.Errorf("Validate(%q, %d).Error() = %q, want %q", tt.name, tt.age, got, tt.wantMsg)
		}
	}
}
