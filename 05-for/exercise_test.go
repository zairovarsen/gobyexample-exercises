package forloop

import "testing"

func TestSumTo(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{1, 1},
		{4, 10},
		{100, 5050},
		{0, 0},
		{-3, 0},
	}
	for _, tt := range tests {
		if got := SumTo(tt.n); got != tt.want {
			t.Errorf("SumTo(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}

func TestRepeat(t *testing.T) {
	tests := []struct {
		s    string
		n    int
		want string
	}{
		{"ab", 3, "ababab"},
		{"go", 1, "go"},
		{"-", 5, "-----"},
		{"ab", 0, ""},
	}
	for _, tt := range tests {
		if got := Repeat(tt.s, tt.n); got != tt.want {
			t.Errorf("Repeat(%q, %d) = %q, want %q", tt.s, tt.n, got, tt.want)
		}
	}
}

func TestCountDigits(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{0, 1},
		{7, 1},
		{10, 2},
		{99, 2},
		{1000, 4},
		{1234567, 7},
	}
	for _, tt := range tests {
		if got := CountDigits(tt.n); got != tt.want {
			t.Errorf("CountDigits(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}

func TestSumOdd(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{10, 25},
		{6, 9},
		{2, 1},
		{100, 2500},
		{0, 0},
	}
	for _, tt := range tests {
		if got := SumOdd(tt.n); got != tt.want {
			t.Errorf("SumOdd(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}

func TestSmallestDivisor(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{2, 2},
		{9, 3},
		{15, 3},
		{49, 7},
		{100, 2},
		{13, 13},
		{97, 97},
	}
	for _, tt := range tests {
		if got := SmallestDivisor(tt.n); got != tt.want {
			t.Errorf("SmallestDivisor(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}

func TestCollatzSteps(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{6, 8},
		{2, 1},
		{7, 16},
		{27, 111},
		{1, 0},
	}
	for _, tt := range tests {
		if got := CollatzSteps(tt.n); got != tt.want {
			t.Errorf("CollatzSteps(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}
