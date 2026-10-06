package sortingbyfunctions

import (
	"slices"
	"testing"
)

func TestSortByLength(t *testing.T) {
	in := []string{"peach", "banana", "kiwi", "fig", "watermelon", "a"}
	want := []string{"a", "fig", "kiwi", "peach", "banana", "watermelon"}
	got := slices.Clone(in)
	SortByLength(got)
	if !slices.Equal(got, want) {
		t.Errorf("after SortByLength(%q) the slice is %q, want %q", in, got, want)
	}
}

func TestSortByLengthWithTies(t *testing.T) {
	in := []string{"ccc", "a", "dd", "bbb", "ee", "f", "gggg"}
	got := slices.Clone(in)
	SortByLength(got)

	for i := 1; i < len(got); i++ {
		if len(got[i-1]) > len(got[i]) {
			t.Fatalf("after SortByLength(%q) the slice is %q: %q comes before the shorter %q",
				in, got, got[i-1], got[i])
		}
	}
	// Still the same words?
	a, b := slices.Clone(in), slices.Clone(got)
	slices.Sort(a)
	slices.Sort(b)
	if !slices.Equal(a, b) {
		t.Errorf("after SortByLength(%q) the slice is %q, want the same words in another order", in, got)
	}
	if got[len(got)-1] != "gggg" {
		t.Errorf("after SortByLength(%q) the last word is %q, want %q", in, got[len(got)-1], "gggg")
	}
}

func TestSortDescending(t *testing.T) {
	tests := []struct {
		in   []int
		want []int
	}{
		{[]int{2, 9, 4}, []int{9, 4, 2}},
		{[]int{-3, 0, 12, 7, 7}, []int{12, 7, 7, 0, -3}},
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
	}
	for _, tt := range tests {
		got := slices.Clone(tt.in)
		SortDescending(got)
		if !slices.Equal(got, tt.want) {
			t.Errorf("after SortDescending(%v) the slice is %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestSortByAge(t *testing.T) {
	in := []Person{{"Jax", 37}, {"TJ", 25}, {"Alex", 72}, {"Kim", 4}}
	want := []Person{{"Kim", 4}, {"TJ", 25}, {"Jax", 37}, {"Alex", 72}}
	got := slices.Clone(in)
	SortByAge(got)
	if !slices.Equal(got, want) {
		t.Errorf("after SortByAge(%v) the slice is %v, want %v", in, got, want)
	}
}

func TestSortByAgeThenName(t *testing.T) {
	in := []Person{
		{"Zoe", 30}, {"Alex", 25}, {"Adam", 30}, {"Mia", 25}, {"Bo", 41}, {"Kim", 30},
	}
	want := []Person{
		{"Alex", 25}, {"Mia", 25}, {"Adam", 30}, {"Kim", 30}, {"Zoe", 30}, {"Bo", 41},
	}
	got := slices.Clone(in)
	SortByAgeThenName(got)
	if !slices.Equal(got, want) {
		t.Errorf("after SortByAgeThenName(%v) the slice is\n  %v, want\n  %v", in, got, want)
	}
}

func TestWordFrequencies(t *testing.T) {
	tests := []struct {
		words []string
		want  []WordCount
	}{
		{
			[]string{"b", "a", "b", "c", "a", "b"},
			[]WordCount{{"b", 3}, {"a", 2}, {"c", 1}},
		},
		{
			[]string{"go", "zig", "go", "ada", "zig", "c"},
			[]WordCount{{"go", 2}, {"zig", 2}, {"ada", 1}, {"c", 1}},
		},
		{
			[]string{"one"},
			[]WordCount{{"one", 1}},
		},
	}
	for _, tt := range tests {
		// Map iteration order changes from run to run, so try a few times.
		for range 5 {
			if got := WordFrequencies(tt.words); !slices.Equal(got, tt.want) {
				t.Errorf("WordFrequencies(%q) = %v, want %v", tt.words, got, tt.want)
				break
			}
		}
	}
	if got := WordFrequencies(nil); len(got) != 0 {
		t.Errorf("WordFrequencies(nil) = %v, want an empty result", got)
	}
}
