package values

import "testing"

func TestFullName(t *testing.T) {
	tests := []struct {
		first, last string
		want        string
	}{
		{"Rob", "Pike", "Rob Pike"},
		{"Ada", "Lovelace", "Ada Lovelace"},
		{"go", "lang", "go lang"},
	}
	for _, tt := range tests {
		if got := FullName(tt.first, tt.last); got != tt.want {
			t.Errorf("FullName(%q, %q) = %q, want %q", tt.first, tt.last, got, tt.want)
		}
	}
}

func TestAverage(t *testing.T) {
	tests := []struct {
		a, b float64
		want float64
	}{
		{1, 2, 1.5},
		{7, 3, 5},
		{0.5, 0.25, 0.375},
		{-4, 1, -1.5},
	}
	for _, tt := range tests {
		if got := Average(tt.a, tt.b); got != tt.want {
			t.Errorf("Average(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestIntDivide(t *testing.T) {
	tests := []struct {
		a, b int
		want int
	}{
		{7, 2, 3},
		{9, 3, 3},
		{100, 7, 14},
		{1, 2, 0},
		{-7, 2, -3},
	}
	for _, tt := range tests {
		if got := IntDivide(tt.a, tt.b); got != tt.want {
			t.Errorf("IntDivide(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

// An Example function is a test too: `go test` runs it and compares what it
// printed to standard output with the "Output:" comment at the end.
func ExamplePrintValues() {
	PrintValues()
	// Output:
	// gobyexample
	// 6*7 = 42
	// 1.0/4.0 = 0.25
	// false
	// true
	// false
}

func TestXor(t *testing.T) {
	tests := []struct {
		a, b bool
		want bool
	}{
		{false, false, false},
		{true, false, true},
		{false, true, true},
		{true, true, false},
	}
	for _, tt := range tests {
		if got := Xor(tt.a, tt.b); got != tt.want {
			t.Errorf("Xor(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestMajority(t *testing.T) {
	tests := []struct {
		a, b, c bool
		want    bool
	}{
		{false, false, false, false},
		{true, false, false, false},
		{false, true, false, false},
		{false, false, true, false},
		{true, true, false, true},
		{true, false, true, true},
		{false, true, true, true},
		{true, true, true, true},
	}
	for _, tt := range tests {
		if got := Majority(tt.a, tt.b, tt.c); got != tt.want {
			t.Errorf("Majority(%v, %v, %v) = %v, want %v", tt.a, tt.b, tt.c, got, tt.want)
		}
	}
}
