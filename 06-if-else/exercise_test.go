package ifelse

import "testing"

func TestEvenOrOdd(t *testing.T) {
	tests := []struct {
		n    int
		want string
	}{
		{7, "odd"},
		{8, "even"},
		{0, "even"},
		{-3, "odd"},
		{-4, "even"},
	}
	for _, tt := range tests {
		if got := EvenOrOdd(tt.n); got != tt.want {
			t.Errorf("EvenOrOdd(%d) = %q, want %q", tt.n, got, tt.want)
		}
	}
}

func TestCompare(t *testing.T) {
	tests := []struct {
		a, b int
		want string
	}{
		{1, 2, "less"},
		{2, 2, "equal"},
		{3, 2, "greater"},
		{-5, -1, "less"},
		{0, -1, "greater"},
	}
	for _, tt := range tests {
		if got := Compare(tt.a, tt.b); got != tt.want {
			t.Errorf("Compare(%d, %d) = %q, want %q", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestLeapYear(t *testing.T) {
	tests := []struct {
		year int
		want bool
	}{
		{2024, true},
		{2023, false},
		{1900, false},
		{2000, true},
		{2100, false},
		{1996, true},
	}
	for _, tt := range tests {
		if got := LeapYear(tt.year); got != tt.want {
			t.Errorf("LeapYear(%d) = %v, want %v", tt.year, got, tt.want)
		}
	}
}

func TestGrade(t *testing.T) {
	tests := []struct {
		score int
		want  string
	}{
		{100, "A"},
		{90, "A"},
		{89, "B"},
		{80, "B"},
		{75, "C"},
		{60, "D"},
		{59, "F"},
		{0, "F"},
		{-1, "invalid"},
		{101, "invalid"},
	}
	for _, tt := range tests {
		if got := Grade(tt.score); got != tt.want {
			t.Errorf("Grade(%d) = %q, want %q", tt.score, got, tt.want)
		}
	}
}

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		n    int
		want string
	}{
		{1, "1"},
		{3, "Fizz"},
		{5, "Buzz"},
		{7, "7"},
		{9, "Fizz"},
		{15, "FizzBuzz"},
		{30, "FizzBuzz"},
		{98, "98"},
		{100, "Buzz"},
	}
	for _, tt := range tests {
		if got := FizzBuzz(tt.n); got != tt.want {
			t.Errorf("FizzBuzz(%d) = %q, want %q", tt.n, got, tt.want)
		}
	}
}

func TestFizzBuzzUpTo(t *testing.T) {
	tests := []struct {
		n    int
		want string
	}{
		{1, "1"},
		{5, "1 2 Fizz 4 Buzz"},
		{15, "1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz 11 Fizz 13 14 FizzBuzz"},
		{0, ""},
	}
	for _, tt := range tests {
		if got := FizzBuzzUpTo(tt.n); got != tt.want {
			t.Errorf("FizzBuzzUpTo(%d) = %q, want %q", tt.n, got, tt.want)
		}
	}
}
