package rangeoverbuiltintypes

import (
	"maps"
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{2, 3, 4}, 9},
		{[]int{-1, 1, 10}, 10},
		{[]int{42}, 42},
		{nil, 0},
	}
	for _, tt := range tests {
		if got := Sum(tt.nums); got != tt.want {
			t.Errorf("Sum(%v) = %d, want %d", tt.nums, got, tt.want)
		}
	}
}

func TestIndexOf(t *testing.T) {
	words := []string{"a", "b", "c", "b"}
	tests := []struct {
		target string
		want   int
	}{
		{"b", 1},
		{"c", 2},
		{"a", 0},
		{"z", -1},
	}
	for _, tt := range tests {
		if got := IndexOf(words, tt.target); got != tt.want {
			t.Errorf("IndexOf(%v, %q) = %d, want %d", words, tt.target, got, tt.want)
		}
	}

	if got := IndexOf(nil, "a"); got != -1 {
		t.Errorf("IndexOf(nil, %q) = %d, want -1", "a", got)
	}
}

func TestKeys(t *testing.T) {
	tests := []struct {
		m    map[string]int
		want []string
	}{
		{map[string]int{"b": 2, "a": 1, "c": 3}, []string{"a", "b", "c"}},
		{map[string]int{"pear": 0, "fig": 0, "apple": 0, "kiwi": 0}, []string{"apple", "fig", "kiwi", "pear"}},
		{map[string]int{"only": 1}, []string{"only"}},
	}
	for _, tt := range tests {
		if got := Keys(tt.m); !slices.Equal(got, tt.want) {
			t.Errorf("Keys(%v) = %v, want %v", tt.m, got, tt.want)
		}
	}

	if got := Keys(nil); len(got) != 0 {
		t.Errorf("Keys(nil) = %v, want an empty slice", got)
	}
}

func TestInvert(t *testing.T) {
	tests := []struct {
		m    map[string]string
		want map[string]string
	}{
		{
			map[string]string{"a": "x", "b": "y"},
			map[string]string{"x": "a", "y": "b"},
		},
		{
			map[string]string{"go": "gopher", "rust": "crab", "python": "snake"},
			map[string]string{"gopher": "go", "crab": "rust", "snake": "python"},
		},
	}
	for _, tt := range tests {
		if got := Invert(tt.m); !maps.Equal(got, tt.want) {
			t.Errorf("Invert(%v) = %v, want %v", tt.m, got, tt.want)
		}
	}

	if got := Invert(nil); len(got) != 0 {
		t.Errorf("Invert(nil) = %v, want an empty map", got)
	}
}

func TestByteOffsets(t *testing.T) {
	tests := []struct {
		s    string
		want []int
	}{
		{"abc", []int{0, 1, 2}},
		{"héllo", []int{0, 1, 3, 4, 5}},
		{"สวัสดี", []int{0, 3, 6, 9, 12, 15}},
		{"go 😀!", []int{0, 1, 2, 3, 7}},
	}
	for _, tt := range tests {
		if got := ByteOffsets(tt.s); !slices.Equal(got, tt.want) {
			t.Errorf("ByteOffsets(%q) = %v, want %v", tt.s, got, tt.want)
		}
	}

	if got := ByteOffsets(""); len(got) != 0 {
		t.Errorf("ByteOffsets(%q) = %v, want an empty slice", "", got)
	}
}

func TestGroupByLength(t *testing.T) {
	tests := []struct {
		words []string
		want  map[int][]string
	}{
		{
			[]string{"go", "is", "fun", "to", "learn"},
			map[int][]string{2: {"go", "is", "to"}, 3: {"fun"}, 5: {"learn"}},
		},
		{
			[]string{"b", "a", "b"},
			map[int][]string{1: {"b", "a", "b"}},
		},
	}
	for _, tt := range tests {
		got := GroupByLength(tt.words)
		if !maps.EqualFunc(got, tt.want, slices.Equal[[]string]) {
			t.Errorf("GroupByLength(%v) = %v, want %v", tt.words, got, tt.want)
		}
	}

	if got := GroupByLength(nil); len(got) != 0 {
		t.Errorf("GroupByLength(nil) = %v, want an empty map", got)
	}
}
