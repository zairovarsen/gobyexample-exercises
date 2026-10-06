package sorting

import (
	"slices"
	"testing"
)

func TestSortStrings(t *testing.T) {
	tests := []struct {
		in   []string
		want []string
	}{
		{[]string{"c", "a", "b"}, []string{"a", "b", "c"}},
		{[]string{"pear", "fig", "apple", "banana"}, []string{"apple", "banana", "fig", "pear"}},
		{[]string{"go", "Zig", "c", "Ada"}, []string{"Ada", "Zig", "c", "go"}},
		{[]string{}, []string{}},
	}
	for _, tt := range tests {
		got := slices.Clone(tt.in)
		SortStrings(got)
		if !slices.Equal(got, tt.want) {
			t.Errorf("after SortStrings(%q) the slice is %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestSortedInts(t *testing.T) {
	tests := []struct {
		in   []int
		want []int
	}{
		{[]int{7, 2, 4}, []int{2, 4, 7}},
		{[]int{3, -1, 3, 0, -8}, []int{-8, -1, 0, 3, 3}},
		{[]int{1}, []int{1}},
	}
	for _, tt := range tests {
		in := slices.Clone(tt.in)
		got := SortedInts(in)
		if !slices.Equal(got, tt.want) {
			t.Errorf("SortedInts(%v) = %v, want %v", tt.in, got, tt.want)
		}
		if !slices.Equal(in, tt.in) {
			t.Errorf("SortedInts(%v) changed its argument to %v, want it untouched", tt.in, in)
		}
	}
}

func TestIsSorted(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		{[]int{2, 4, 7}, true},
		{[]int{1, 1, 2}, true},
		{[]int{5}, true},
		{nil, true},
		{[]int{7, 2, 4}, false},
		{[]int{1, 2, 3, 2}, false},
		{[]int{3, 2, 1}, false},
	}
	for _, tt := range tests {
		if got := IsSorted(tt.nums); got != tt.want {
			t.Errorf("IsSorted(%v) = %t, want %t", tt.nums, got, tt.want)
		}
	}
}

func TestSortedKeys(t *testing.T) {
	m := map[string]int{"pear": 1, "apple": 5, "fig": 2, "banana": 0, "cherry": 9}
	want := []string{"apple", "banana", "cherry", "fig", "pear"}
	// Map iteration order changes from run to run, so try a few times.
	for range 5 {
		if got := SortedKeys(m); !slices.Equal(got, want) {
			t.Fatalf("SortedKeys(%v) = %q, want %q", m, got, want)
		}
	}
	if got := SortedKeys(nil); len(got) != 0 {
		t.Errorf("SortedKeys(nil) = %q, want an empty result", got)
	}
}

func TestTopN(t *testing.T) {
	nums := []int{3, 9, 1, 7, 5}
	intTests := []struct {
		n    int
		want []int
	}{
		{2, []int{9, 7}},
		{1, []int{9}},
		{5, []int{9, 7, 5, 3, 1}},
		{8, []int{9, 7, 5, 3, 1}},
		{0, []int{}},
		{-1, []int{}},
	}
	for _, tt := range intTests {
		in := slices.Clone(nums)
		got := TopN(in, tt.n)
		if !slices.Equal(got, tt.want) {
			t.Errorf("TopN(%v, %d) = %v, want %v", nums, tt.n, got, tt.want)
		}
		if !slices.Equal(in, nums) {
			t.Errorf("TopN(%v, %d) changed its argument to %v, want it untouched", nums, tt.n, in)
		}
	}

	words := []string{"fig", "pear", "apple"}
	if got, want := TopN(words, 2), []string{"pear", "fig"}; !slices.Equal(got, want) {
		t.Errorf("TopN(%q, 2) = %q, want %q", words, got, want)
	}
	floats := []float64{1.5, -2, 9.25}
	if got, want := TopN(floats, 1), []float64{9.25}; !slices.Equal(got, want) {
		t.Errorf("TopN(%v, 1) = %v, want %v", floats, got, want)
	}
}
