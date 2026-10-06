package generics

import (
	"maps"
	"slices"
	"strconv"
	"testing"
)

// Named types, to check that the ~ in the constraints is honoured.
type (
	words   []string
	celsius float64
)

func TestIndex(t *testing.T) {
	strs := []string{"foo", "bar", "zoo", "bar"}
	stringTests := []struct {
		v    string
		want int
	}{
		{"zoo", 2},
		{"bar", 1},
		{"foo", 0},
		{"nope", -1},
	}
	for _, tt := range stringTests {
		if got := Index(strs, tt.v); got != tt.want {
			t.Errorf("Index(%q, %q) = %d, want %d", strs, tt.v, got, tt.want)
		}
	}

	if got := Index([]int{10, 20, 30}, 30); got != 2 {
		t.Errorf("Index([10 20 30], 30) = %d, want 2", got)
	}
	if got := Index([]float64{}, 1.5); got != -1 {
		t.Errorf("Index([], 1.5) = %d, want -1", got)
	}
	// A named slice type satisfies ~[]E; the type arguments can also be
	// spelled out explicitly.
	if got := Index[words, string](words{"a", "b", "c"}, "c"); got != 2 {
		t.Errorf(`Index(words{"a", "b", "c"}, "c") = %d, want 2`, got)
	}
}

func TestMap(t *testing.T) {
	if got, want := Map([]int{1, 2, 3}, strconv.Itoa), []string{"1", "2", "3"}; !slices.Equal(got, want) {
		t.Errorf("Map([1 2 3], strconv.Itoa) = %q, want %q", got, want)
	}

	length := func(s string) int { return len(s) }
	if got, want := Map([]string{"go", "", "gopher"}, length), []int{2, 0, 6}; !slices.Equal(got, want) {
		t.Errorf(`Map(["go" "" "gopher"], length) = %v, want %v`, got, want)
	}

	double := func(f float64) float64 { return f * 2 }
	if got := Map(nil, double); len(got) != 0 {
		t.Errorf("Map(nil, double) = %v, want an empty result", got)
	}
}

func TestSum(t *testing.T) {
	if got := Sum([]int{1, 2, 3, 4}); got != 10 {
		t.Errorf("Sum([]int{1, 2, 3, 4}) = %d, want 10", got)
	}
	if got := Sum([]int64{-5, 5, 1 << 40}); got != 1<<40 {
		t.Errorf("Sum([]int64{-5, 5, 1 << 40}) = %d, want %d", got, int64(1)<<40)
	}
	if got := Sum([]float64{0.5, 0.25, 1}); got != 1.75 {
		t.Errorf("Sum([]float64{0.5, 0.25, 1}) = %v, want 1.75", got)
	}
	if got := Sum([]celsius{20.5, 1.5}); got != 22 {
		t.Errorf("Sum([]celsius{20.5, 1.5}) = %v, want 22", got)
	}
	if got := Sum[int](nil); got != 0 {
		t.Errorf("Sum[int](nil) = %d, want 0", got)
	}
}

func TestList(t *testing.T) {
	var lst List[int]
	if got := lst.Len(); got != 0 {
		t.Errorf("empty list: Len() = %d, want 0", got)
	}
	if got := lst.AllElements(); len(got) != 0 {
		t.Errorf("empty list: AllElements() = %v, want an empty result", got)
	}

	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	if got, want := lst.AllElements(), []int{10, 13, 23}; !slices.Equal(got, want) {
		t.Errorf("after Push(10), Push(13), Push(23): AllElements() = %v, want %v", got, want)
	}
	if got := lst.Len(); got != 3 {
		t.Errorf("after three pushes: Len() = %d, want 3", got)
	}

	// The same code works for any element type.
	var names List[string]
	names.Push("ann")
	names.Push("bob")
	if got, want := names.AllElements(), []string{"ann", "bob"}; !slices.Equal(got, want) {
		t.Errorf(`after Push("ann"), Push("bob"): AllElements() = %q, want %q`, got, want)
	}
}

func TestGroupBy(t *testing.T) {
	byLen := GroupBy([]string{"go", "is", "fun", "to", "learn"}, func(s string) int { return len(s) })
	wantByLen := map[int][]string{
		2: {"go", "is", "to"},
		3: {"fun"},
		5: {"learn"},
	}
	if !maps.EqualFunc(byLen, wantByLen, slices.Equal[[]string]) {
		t.Errorf(`GroupBy(["go" "is" "fun" "to" "learn"], len) = %v, want %v`, byLen, wantByLen)
	}

	parity := func(n int) string {
		if n%2 == 0 {
			return "even"
		}
		return "odd"
	}
	byParity := GroupBy([]int{1, 2, 3, 4, 5}, parity)
	wantByParity := map[string][]int{
		"odd":  {1, 3, 5},
		"even": {2, 4},
	}
	if !maps.EqualFunc(byParity, wantByParity, slices.Equal[[]int]) {
		t.Errorf("GroupBy([1 2 3 4 5], parity) = %v, want %v", byParity, wantByParity)
	}

	if got := GroupBy(nil, parity); len(got) != 0 {
		t.Errorf("GroupBy(nil, parity) = %v, want an empty map", got)
	}
}
