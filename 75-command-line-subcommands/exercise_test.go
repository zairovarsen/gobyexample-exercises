package commandlinesubcommands

import (
	"errors"
	"slices"
	"strings"
	"testing"
)

func TestParseFoo(t *testing.T) {
	tests := []struct {
		args []string
		want FooOptions
	}{
		{[]string{"-enable", "-name=joe", "a1", "a2"}, FooOptions{Enable: true, Name: "joe", Tail: []string{"a1", "a2"}}},
		{[]string{"-name", "ann"}, FooOptions{Enable: false, Name: "ann"}},
		{[]string{"-enable"}, FooOptions{Enable: true, Name: ""}},
		{[]string{"a1", "-enable"}, FooOptions{Enable: false, Name: "", Tail: []string{"a1", "-enable"}}},
	}
	for _, tt := range tests {
		got, err := ParseFoo(tt.args)
		if err != nil {
			t.Errorf("ParseFoo(%q) returned error: %v", tt.args, err)
			continue
		}
		if got.Enable != tt.want.Enable || got.Name != tt.want.Name || !slices.Equal(got.Tail, tt.want.Tail) {
			t.Errorf("ParseFoo(%q) = %+v, want %+v", tt.args, got, tt.want)
		}
	}

	for _, args := range [][]string{
		{"-level", "8"}, // a flag of bar, not of foo
		{"-name"},       // the value is missing
	} {
		if got, err := ParseFoo(args); err == nil {
			t.Errorf("ParseFoo(%q) = %+v, nil - want an error", args, got)
		}
	}
}

func TestParseBar(t *testing.T) {
	tests := []struct {
		args []string
		want BarOptions
	}{
		{[]string{"-level", "8", "a1"}, BarOptions{Level: 8, Tail: []string{"a1"}}},
		{[]string{"-level=3"}, BarOptions{Level: 3}},
		{[]string{"x", "y"}, BarOptions{Level: 0, Tail: []string{"x", "y"}}},
	}
	for _, tt := range tests {
		got, err := ParseBar(tt.args)
		if err != nil {
			t.Errorf("ParseBar(%q) returned error: %v", tt.args, err)
			continue
		}
		if got.Level != tt.want.Level || !slices.Equal(got.Tail, tt.want.Tail) {
			t.Errorf("ParseBar(%q) = %+v, want %+v", tt.args, got, tt.want)
		}
	}

	for _, args := range [][]string{
		{"-enable", "a1"}, // a flag of foo, not of bar
		{"-level=high"},
	} {
		if got, err := ParseBar(args); err == nil {
			t.Errorf("ParseBar(%q) = %+v, nil - want an error", args, got)
		}
	}
}

func TestRun(t *testing.T) {
	tests := []struct {
		args []string
		want string
	}{
		{[]string{"foo", "-enable", "-name=joe", "a1", "a2"}, "foo: enable=true name=joe tail=[a1 a2]"},
		{[]string{"foo", "-name=ann"}, "foo: enable=false name=ann tail=[]"},
		{[]string{"foo"}, "foo: enable=false name= tail=[]"},
		{[]string{"bar", "-level", "8", "a1"}, "bar: level=8 tail=[a1]"},
		{[]string{"bar"}, "bar: level=0 tail=[]"},
		{[]string{"bar", "foo"}, "bar: level=0 tail=[foo]"},
	}
	for _, tt := range tests {
		got, err := Run(tt.args)
		if err != nil {
			t.Errorf("Run(%q) returned error: %v", tt.args, err)
			continue
		}
		if got != tt.want {
			t.Errorf("Run(%q) = %q, want %q", tt.args, got, tt.want)
		}
	}

	for _, args := range [][]string{nil, {}} {
		if got, err := Run(args); !errors.Is(err, ErrNoSubcommand) {
			t.Errorf("Run(%q) = %q, %v, want the error ErrNoSubcommand", args, got, err)
		}
	}

	got, err := Run([]string{"baz", "-enable"})
	if !errors.Is(err, ErrUnknownSubcommand) {
		t.Errorf(`Run(["baz" "-enable"]) = %q, %v, want an error wrapping ErrUnknownSubcommand`, got, err)
	} else if !strings.Contains(err.Error(), "baz") {
		t.Errorf(`Run(["baz" "-enable"]) error = %q, want it to mention the unknown subcommand "baz"`, err)
	}

	// The flags belong to their subcommand: bar does not know -enable.
	got, err = Run([]string{"bar", "-enable", "a1"})
	if err == nil {
		t.Errorf(`Run(["bar" "-enable" "a1"]) = %q, nil - want the flag parsing error`, got)
	} else if errors.Is(err, ErrNoSubcommand) || errors.Is(err, ErrUnknownSubcommand) {
		t.Errorf(`Run(["bar" "-enable" "a1"]) error = %v, want the error from parsing bar's flags`, err)
	}
}

func TestDispatch(t *testing.T) {
	var gotArgs []string
	errFailed := errors.New("command failed")
	commands := map[string]func(args []string) (string, error){
		"greet": func(args []string) (string, error) {
			gotArgs = args
			return "hello " + strings.Join(args, " and "), nil
		},
		"fail": func(args []string) (string, error) {
			return "partial", errFailed
		},
	}

	got, err := Dispatch([]string{"greet", "ann", "bob"}, commands)
	if err != nil {
		t.Fatalf(`Dispatch(["greet" "ann" "bob"]) returned error: %v`, err)
	}
	if want := "hello ann and bob"; got != want {
		t.Errorf(`Dispatch(["greet" "ann" "bob"]) = %q, want %q`, got, want)
	}
	if !slices.Equal(gotArgs, []string{"ann", "bob"}) {
		t.Errorf("the greet command received the arguments %q, want [\"ann\" \"bob\"] (everything after the subcommand)", gotArgs)
	}

	gotArgs = []string{"stale"}
	if got, err := Dispatch([]string{"greet"}, commands); err != nil || got != "hello " {
		t.Errorf(`Dispatch(["greet"]) = %q, %v, want "hello ", nil`, got, err)
	}
	if len(gotArgs) != 0 {
		t.Errorf("the greet command received the arguments %q, want none", gotArgs)
	}

	got, err = Dispatch([]string{"fail", "x"}, commands)
	if !errors.Is(err, errFailed) || got != "partial" {
		t.Errorf(`Dispatch(["fail" "x"]) = %q, %v, want exactly what the command returned: "partial", %v`, got, err, errFailed)
	}

	for _, args := range [][]string{nil, {}} {
		if got, err := Dispatch(args, commands); !errors.Is(err, ErrNoSubcommand) {
			t.Errorf("Dispatch(%q) = %q, %v, want the error ErrNoSubcommand", args, got, err)
		}
	}

	got, err = Dispatch([]string{"dance", "greet"}, commands)
	if !errors.Is(err, ErrUnknownSubcommand) {
		t.Errorf(`Dispatch(["dance" "greet"]) = %q, %v, want an error wrapping ErrUnknownSubcommand`, got, err)
	} else if !strings.Contains(err.Error(), "dance") {
		t.Errorf(`Dispatch(["dance" "greet"]) error = %q, want it to mention the unknown subcommand "dance"`, err)
	}
}
