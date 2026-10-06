package commandlineflags

import (
	"bytes"
	"errors"
	"flag"
	"slices"
	"strings"
	"testing"
)

func TestParseWord(t *testing.T) {
	tests := []struct {
		args []string
		want string
	}{
		{[]string{"-word=opt"}, "opt"},
		{[]string{"-word", "opt"}, "opt"},
		{[]string{"--word=two words"}, "two words"},
		{[]string{"-word="}, ""},
		{[]string{}, "foo"},
		{nil, "foo"},
		{[]string{"positional"}, "foo"},
	}
	for _, tt := range tests {
		got, err := ParseWord(tt.args)
		if err != nil {
			t.Errorf("ParseWord(%q) returned error: %v", tt.args, err)
			continue
		}
		if got != tt.want {
			t.Errorf("ParseWord(%q) = %q, want %q", tt.args, got, tt.want)
		}
	}

	for _, args := range [][]string{
		{"-wat"},
		{"-word"}, // the value is missing
	} {
		if got, err := ParseWord(args); err == nil {
			t.Errorf("ParseWord(%q) = %q, nil - want an error", args, got)
		}
	}
}

func TestParseNumbFork(t *testing.T) {
	tests := []struct {
		args     []string
		wantNumb int
		wantFork bool
	}{
		{[]string{"-numb=7", "-fork"}, 7, true},
		{[]string{"-fork", "-numb", "7"}, 7, true},
		{[]string{"-numb=-3"}, -3, false},
		{[]string{"-fork=false"}, 42, false},
		{[]string{"-fork"}, 42, true},
		{nil, 42, false},
	}
	for _, tt := range tests {
		numb, fork, err := ParseNumbFork(tt.args)
		if err != nil {
			t.Errorf("ParseNumbFork(%q) returned error: %v", tt.args, err)
			continue
		}
		if numb != tt.wantNumb || fork != tt.wantFork {
			t.Errorf("ParseNumbFork(%q) = %d, %v, want %d, %v", tt.args, numb, fork, tt.wantNumb, tt.wantFork)
		}
	}

	for _, args := range [][]string{
		{"-numb=seven"},
		{"-numb=1.5"},
		{"-word=opt"}, // not declared in this flag set
	} {
		if numb, fork, err := ParseNumbFork(args); err == nil {
			t.Errorf("ParseNumbFork(%q) = %d, %v, nil - want an error", args, numb, fork)
		}
	}
}

func TestParse(t *testing.T) {
	tests := []struct {
		args []string
		want Config
	}{
		{
			[]string{"-word=opt", "-numb=7", "-fork", "-svar=flag"},
			Config{Word: "opt", Numb: 7, Fork: true, SVar: "flag"},
		},
		{
			[]string{"-word=opt"},
			Config{Word: "opt", Numb: 42, Fork: false, SVar: "bar"},
		},
		{
			[]string{"-word=opt", "a1", "a2", "a3"},
			Config{Word: "opt", Numb: 42, Fork: false, SVar: "bar", Tail: []string{"a1", "a2", "a3"}},
		},
		{
			// Flags after the first positional argument are not parsed.
			[]string{"-word=opt", "a1", "a2", "a3", "-numb=7"},
			Config{Word: "opt", Numb: 42, Fork: false, SVar: "bar", Tail: []string{"a1", "a2", "a3", "-numb=7"}},
		},
		{
			[]string{"--svar", "s", "--numb", "0", "tail"},
			Config{Word: "foo", Numb: 0, Fork: false, SVar: "s", Tail: []string{"tail"}},
		},
		{
			nil,
			Config{Word: "foo", Numb: 42, Fork: false, SVar: "bar"},
		},
	}
	for _, tt := range tests {
		var output bytes.Buffer
		got, err := Parse(tt.args, &output)
		if err != nil {
			t.Errorf("Parse(%q) returned error: %v", tt.args, err)
			continue
		}
		if got.Word != tt.want.Word || got.Numb != tt.want.Numb || got.Fork != tt.want.Fork ||
			got.SVar != tt.want.SVar || !slices.Equal(got.Tail, tt.want.Tail) {
			t.Errorf("Parse(%q) = %+v, want %+v", tt.args, got, tt.want)
		}
		if output.Len() != 0 {
			t.Errorf("Parse(%q) wrote %q to output, want nothing for valid arguments", tt.args, output.String())
		}
	}
}

func TestParseErrors(t *testing.T) {
	flagNames := []string{"-word", "-numb", "-fork", "-svar"}

	var output bytes.Buffer
	got, err := Parse([]string{"-wat"}, &output)
	if err == nil {
		t.Fatalf(`Parse(["-wat"]) = %+v, nil - want an error for the unknown flag`, got)
	}
	if !strings.Contains(output.String(), "-wat") {
		t.Errorf(`Parse(["-wat"]) wrote %q to output, want the flag set's error message mentioning -wat (fs.SetOutput(output))`, output.String())
	}
	for _, name := range flagNames {
		if !strings.Contains(output.String(), name) {
			t.Errorf(`Parse(["-wat"]): the usage text in output does not list the flag %s`, name)
		}
	}

	output.Reset()
	if _, err := Parse([]string{"-numb=seven"}, &output); err == nil {
		t.Error(`Parse(["-numb=seven"]) returned a nil error, want an error for the invalid int`)
	}

	// -h is defined automatically; it reports flag.ErrHelp and prints the usage.
	output.Reset()
	_, err = Parse([]string{"-h"}, &output)
	if !errors.Is(err, flag.ErrHelp) {
		t.Errorf(`Parse(["-h"]) error = %v, want flag.ErrHelp`, err)
	}
	for _, name := range flagNames {
		if !strings.Contains(output.String(), name) {
			t.Errorf(`Parse(["-h"]): the usage text in output does not list the flag %s`, name)
		}
	}
}

func TestParseServer(t *testing.T) {
	tests := []struct {
		args []string
		want string
	}{
		{nil, "http://localhost:8080"},
		{[]string{"-port=9000", "-tls"}, "https://localhost:9000"},
		{[]string{"-host", "example.com"}, "http://example.com:8080"},
		{[]string{"-tls", "-host=0.0.0.0", "-port=443"}, "https://0.0.0.0:443"},
		{[]string{"-port=1"}, "http://localhost:1"},
		{[]string{"-port=65535"}, "http://localhost:65535"},
	}
	for _, tt := range tests {
		got, err := ParseServer(tt.args)
		if err != nil {
			t.Errorf("ParseServer(%q) returned error: %v", tt.args, err)
			continue
		}
		if got != tt.want {
			t.Errorf("ParseServer(%q) = %q, want %q", tt.args, got, tt.want)
		}
	}

	for _, args := range [][]string{
		{"-port=0"},
		{"-port=65536"},
		{"-port=-1"},
		{"-port=http"},
		{"-verbose"},
		{"-port=9000", "extra"},
		{"start"},
	} {
		if got, err := ParseServer(args); err == nil {
			t.Errorf("ParseServer(%q) = %q, nil - want an error", args, got)
		}
	}
}
