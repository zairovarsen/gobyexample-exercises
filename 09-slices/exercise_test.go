package slicesx

import (
	"slices"
	"testing"
)

func TestCountdown(t *testing.T) {
	tests := []struct {
		n    int
		want []int
	}{
		{3, []int{3, 2, 1}},
		{1, []int{1}},
		{5, []int{5, 4, 3, 2, 1}},
	}
	for _, tt := range tests {
		if got := Countdown(tt.n); !slices.Equal(got, tt.want) {
			t.Errorf("Countdown(%d) = %v, want %v", tt.n, got, tt.want)
		}
	}
	if got := Countdown(0); len(got) != 0 {
		t.Errorf("Countdown(0) = %v, want an empty slice", got)
	}
}

func TestMiddle(t *testing.T) {
	tests := []struct {
		s    []int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{2, 3, 4}},
		{[]int{1, 2, 3}, []int{2}},
		{[]int{1, 2}, []int{}},
		{[]int{1}, []int{}},
		{nil, []int{}},
	}
	for _, tt := range tests {
		if got := Middle(tt.s); !slices.Equal(got, tt.want) {
			t.Errorf("Middle(%v) = %v, want %v", tt.s, got, tt.want)
		}
	}

	s := []int{1, 2, 3, 4, 5}
	m := Middle(s)
	if len(m) != 3 {
		t.Fatalf("Middle(%v) has %d elements, want 3", s, len(m))
	}
	m[0] = 99
	if s[1] != 99 {
		t.Errorf("writing to Middle(s)[0] did not change s[1]: the result must be a view of s (s[low:high]), not a copy")
	}
}

func TestClone(t *testing.T) {
	s := []int{1, 2, 3}
	c := Clone(s)
	if !slices.Equal(c, []int{1, 2, 3}) {
		t.Fatalf("Clone(%v) = %v, want %v", s, c, s)
	}
	c[0] = 99
	if s[0] != 1 {
		t.Errorf("changing the clone changed the original to %v: the clone must not share memory with s", s)
	}

	if got := Clone(nil); len(got) != 0 {
		t.Errorf("Clone(nil) = %v, want an empty slice", got)
	}
}

func TestHasPrefix(t *testing.T) {
	abc := []string{"a", "b", "c"}
	tests := []struct {
		s, prefix []string
		want      bool
	}{
		{abc, []string{"a", "b"}, true},
		{abc, []string{"a"}, true},
		{abc, []string{"a", "b", "c"}, true},
		{abc, nil, true},
		{nil, nil, true},
		{abc, []string{"b"}, false},
		{abc, []string{"a", "c"}, false},
		{abc, []string{"a", "b", "c", "d"}, false},
		{nil, []string{"a"}, false},
	}
	for _, tt := range tests {
		if got := HasPrefix(tt.s, tt.prefix); got != tt.want {
			t.Errorf("HasPrefix(%v, %v) = %v, want %v", tt.s, tt.prefix, got, tt.want)
		}
	}
}

func TestTriangle(t *testing.T) {
	want := [][]int{
		{0},
		{1, 2},
		{2, 3, 4},
		{3, 4, 5, 6},
	}
	got := Triangle(4)
	if len(got) != len(want) {
		t.Fatalf("Triangle(4) = %v: it has %d rows, want %d", got, len(got), len(want))
	}
	for i := range want {
		if !slices.Equal(got[i], want[i]) {
			t.Errorf("Triangle(4)[%d] = %v, want %v", i, got[i], want[i])
		}
	}

	if got := Triangle(0); len(got) != 0 {
		t.Errorf("Triangle(0) = %v, want no rows", got)
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		i    int
		want []int
	}{
		{0, []int{20, 30, 40}},
		{1, []int{10, 30, 40}},
		{2, []int{10, 20, 40}},
		{3, []int{10, 20, 30}},
	}
	for _, tt := range tests {
		s := []int{10, 20, 30, 40}
		got := Remove(s, tt.i)
		if !slices.Equal(got, tt.want) {
			t.Errorf("Remove([10 20 30 40], %d) = %v, want %v", tt.i, got, tt.want)
		}
		if !slices.Equal(s, []int{10, 20, 30, 40}) {
			t.Errorf("Remove([10 20 30 40], %d) changed the slice that was passed in to %v", tt.i, s)
		}
	}

	if got := Remove([]int{7}, 0); len(got) != 0 {
		t.Errorf("Remove([7], 0) = %v, want an empty slice", got)
	}
}
