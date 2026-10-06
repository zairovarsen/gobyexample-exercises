package variadicfunctions

import (
	"math"
	"testing"
)

func TestSum(t *testing.T) {
	if got := Sum(1, 2); got != 3 {
		t.Errorf("Sum(1, 2) = %d, want 3", got)
	}
	if got := Sum(1, 2, 3); got != 6 {
		t.Errorf("Sum(1, 2, 3) = %d, want 6", got)
	}
	if got := Sum(-4); got != -4 {
		t.Errorf("Sum(-4) = %d, want -4", got)
	}
	if got := Sum(); got != 0 {
		t.Errorf("Sum() = %d, want 0", got)
	}

	// A slice can be passed to a variadic function with a trailing "...".
	nums := []int{1, 2, 3, 4}
	if got := Sum(nums...); got != 10 {
		t.Errorf("Sum(%v...) = %d, want 10", nums, got)
	}
}

func TestAverage(t *testing.T) {
	check := func(call string, got, want float64) {
		t.Helper()
		if math.IsNaN(got) || math.Abs(got-want) > 1e-9 {
			t.Errorf("%s = %v, want %v", call, got, want)
		}
	}
	check("Average(1, 2, 4)", Average(1, 2, 4), 7.0/3.0)
	check("Average(2.5)", Average(2.5), 2.5)
	check("Average(1, 2, 3, 4)", Average(1, 2, 3, 4), 2.5)
	check("Average()", Average(), 0)
}

func TestJoin(t *testing.T) {
	tests := []struct {
		sep   string
		parts []string
		want  string
	}{
		{"-", []string{"a", "b", "c"}, "a-b-c"},
		{", ", []string{"x", "y"}, "x, y"},
		{"-", []string{"a"}, "a"},
		{"", []string{"go", "lang"}, "golang"},
		{"-", nil, ""},
	}
	for _, tt := range tests {
		if got := Join(tt.sep, tt.parts...); got != tt.want {
			t.Errorf("Join(%q, %q...) = %q, want %q", tt.sep, tt.parts, got, tt.want)
		}
	}
}

func TestMax(t *testing.T) {
	if got := Max(3, 9, 4); got != 9 {
		t.Errorf("Max(3, 9, 4) = %d, want 9", got)
	}
	if got := Max(7); got != 7 {
		t.Errorf("Max(7) = %d, want 7", got)
	}
	if got := Max(9, 3, 4); got != 9 {
		t.Errorf("Max(9, 3, 4) = %d, want 9", got)
	}
	if got := Max(-5, -9, -1); got != -1 {
		t.Errorf("Max(-5, -9, -1) = %d, want -1", got)
	}
	if got := Max(-5, -9); got != -5 {
		t.Errorf("Max(-5, -9) = %d, want -5", got)
	}
}

func TestTotal(t *testing.T) {
	if got := Total([]int{1, 2}, []int{3}, []int{4, 5, 6}); got != 21 {
		t.Errorf("Total([1 2], [3], [4 5 6]) = %d, want 21", got)
	}
	if got := Total([]int{5}, nil); got != 5 {
		t.Errorf("Total([5], nil) = %d, want 5", got)
	}
	if got := Total(); got != 0 {
		t.Errorf("Total() = %d, want 0", got)
	}
}

func TestGreet(t *testing.T) {
	tests := []struct {
		greeting string
		names    []string
		want     string
	}{
		{"Hello", nil, "Hello!"},
		{"Hello", []string{"Ann"}, "Hello, Ann!"},
		{"Hello", []string{"Ann", "Bob"}, "Hello, Ann and Bob!"},
		{"Hi", []string{"Ann", "Bob", "Cy"}, "Hi, Ann, Bob and Cy!"},
		{"Bye", []string{"Ann", "Bob", "Cy", "Di"}, "Bye, Ann, Bob, Cy and Di!"},
	}
	for _, tt := range tests {
		if got := Greet(tt.greeting, tt.names...); got != tt.want {
			t.Errorf("Greet(%q, %q...) = %q, want %q", tt.greeting, tt.names, got, tt.want)
		}
	}
}
