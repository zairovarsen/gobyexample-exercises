package commandlinearguments

import (
	"slices"
	"testing"
)

func TestProgramName(t *testing.T) {
	tests := []struct {
		args []string
		want string
	}{
		{[]string{"./prog", "a", "b"}, "./prog"},
		{[]string{"command-line-arguments"}, "command-line-arguments"},
		{[]string{}, ""},
		{nil, ""},
	}
	for _, tt := range tests {
		if got := ProgramName(tt.args); got != tt.want {
			t.Errorf("ProgramName(%q) = %q, want %q", tt.args, got, tt.want)
		}
	}
}

func TestArguments(t *testing.T) {
	tests := []struct {
		args []string
		want []string
	}{
		{[]string{"./prog", "a", "b", "c", "d"}, []string{"a", "b", "c", "d"}},
		{[]string{"./prog", "only"}, []string{"only"}},
		{[]string{"./prog"}, nil},
		{[]string{}, nil},
		{nil, nil},
	}
	for _, tt := range tests {
		if got := Arguments(tt.args); !slices.Equal(got, tt.want) {
			t.Errorf("Arguments(%q) = %q, want %q", tt.args, got, tt.want)
		}
	}
}

func TestArgAt(t *testing.T) {
	args := []string{"./prog", "a", "b", "c", "d"}
	tests := []struct {
		args   []string
		n      int
		want   string
		wantOK bool
	}{
		{args, 3, "c", true},
		{args, 1, "a", true},
		{args, 4, "d", true},
		{args, 5, "", false},
		{args, 0, "", false},
		{args, -1, "", false},
		{[]string{"./prog"}, 1, "", false},
		{nil, 1, "", false},
	}
	for _, tt := range tests {
		got, ok := ArgAt(tt.args, tt.n)
		if got != tt.want || ok != tt.wantOK {
			t.Errorf("ArgAt(%q, %d) = %q, %v, want %q, %v", tt.args, tt.n, got, ok, tt.want, tt.wantOK)
		}
	}
}

func TestSum(t *testing.T) {
	tests := []struct {
		args []string
		want int
	}{
		{[]string{"./sum", "1", "2", "39"}, 42},
		{[]string{"./sum", "-5", "15"}, 10},
		{[]string{"./sum", "7"}, 7},
		{[]string{"./sum"}, 0},
	}
	for _, tt := range tests {
		got, err := Sum(tt.args)
		if err != nil {
			t.Errorf("Sum(%q) returned error: %v", tt.args, err)
			continue
		}
		if got != tt.want {
			t.Errorf("Sum(%q) = %d, want %d", tt.args, got, tt.want)
		}
	}

	// The program name is never part of the sum, even if it looks like a number.
	if got, err := Sum([]string{"100", "1", "2"}); err != nil || got != 3 {
		t.Errorf(`Sum(["100" "1" "2"]) = %d, %v, want 3, nil (the first element is the program name)`, got, err)
	}

	for _, args := range [][]string{
		{"./sum", "1", "two", "3"},
		{"./sum", "1.5"},
		{"./sum", ""},
	} {
		if got, err := Sum(args); err == nil {
			t.Errorf("Sum(%q) = %d, nil - want an error", args, got)
		}
	}
}

func TestRepeat(t *testing.T) {
	tests := []struct {
		args []string
		want string
	}{
		{[]string{"./repeat", "2", "hello", "world"}, "hello world\nhello world\n"},
		{[]string{"./repeat", "1", "go"}, "go\n"},
		{[]string{"./repeat", "3", "a", "b", "c"}, "a b c\na b c\na b c\n"},
		{[]string{"./repeat", "2", "7"}, "7\n7\n"},
	}
	for _, tt := range tests {
		got, err := Repeat(tt.args)
		if err != nil {
			t.Errorf("Repeat(%q) returned error: %v", tt.args, err)
			continue
		}
		if got != tt.want {
			t.Errorf("Repeat(%q) = %q, want %q", tt.args, got, tt.want)
		}
	}

	if got, err := Repeat([]string{"./repeat", "0", "quiet"}); err != nil || got != "" {
		t.Errorf(`Repeat(["./repeat" "0" "quiet"]) = %q, %v, want "", nil`, got, err)
	}

	for _, args := range [][]string{
		nil,
		{"./repeat"},
		{"./repeat", "3"},
		{"./repeat", "hello", "world"},
		{"./repeat", "-1", "hello"},
	} {
		if got, err := Repeat(args); err == nil {
			t.Errorf("Repeat(%q) = %q, nil - want an error", args, got)
		}
	}
}
