package mapsx

import (
	"maps"
	"testing"
)

func TestSquares(t *testing.T) {
	tests := []struct {
		n    int
		want map[int]int
	}{
		{1, map[int]int{1: 1}},
		{3, map[int]int{1: 1, 2: 4, 3: 9}},
		{5, map[int]int{1: 1, 2: 4, 3: 9, 4: 16, 5: 25}},
	}
	for _, tt := range tests {
		if got := Squares(tt.n); !maps.Equal(got, tt.want) {
			t.Errorf("Squares(%d) = %v, want %v", tt.n, got, tt.want)
		}
	}

	got := Squares(0)
	if got == nil {
		t.Errorf("Squares(0) = nil, want an empty map created with make")
	}
	if len(got) != 0 {
		t.Errorf("Squares(0) = %v, want an empty map", got)
	}
}

func TestGetOr(t *testing.T) {
	m := map[string]int{"a": 1, "zero": 0}
	tests := []struct {
		key      string
		fallback int
		want     int
	}{
		{"a", 5, 1},
		{"missing", 5, 5},
		{"missing", -1, -1},
		{"zero", 5, 0},
	}
	for _, tt := range tests {
		if got := GetOr(m, tt.key, tt.fallback); got != tt.want {
			t.Errorf("GetOr(%v, %q, %d) = %d, want %d", m, tt.key, tt.fallback, got, tt.want)
		}
	}

	if got := GetOr(nil, "a", 7); got != 7 {
		t.Errorf("GetOr(nil, %q, 7) = %d, want 7", "a", got)
	}
}

func TestWordCount(t *testing.T) {
	tests := []struct {
		words []string
		want  map[string]int
	}{
		{[]string{"a", "b", "a"}, map[string]int{"a": 2, "b": 1}},
		{[]string{"go", "go", "go"}, map[string]int{"go": 3}},
		{[]string{"x"}, map[string]int{"x": 1}},
	}
	for _, tt := range tests {
		if got := WordCount(tt.words); !maps.Equal(got, tt.want) {
			t.Errorf("WordCount(%v) = %v, want %v", tt.words, got, tt.want)
		}
	}

	if got := WordCount(nil); len(got) != 0 {
		t.Errorf("WordCount(nil) = %v, want an empty map", got)
	}
}

func TestRemove(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	keys := []string{"a", "c", "x"}
	if got := Remove(m, keys); got != 2 {
		t.Errorf("Remove(map[a:1 b:2 c:3], %v) = %d, want 2", keys, got)
	}
	if want := map[string]int{"b": 2}; !maps.Equal(m, want) {
		t.Errorf("after Remove(map[a:1 b:2 c:3], %v) the map is %v, want %v", keys, m, want)
	}

	// The same key twice can only be removed once, and a key that holds the
	// zero value is a key like any other.
	m = map[string]int{"zero": 0, "one": 1}
	keys = []string{"zero", "zero"}
	if got := Remove(m, keys); got != 1 {
		t.Errorf("Remove(map[one:1 zero:0], %v) = %d, want 1", keys, got)
	}
	if want := map[string]int{"one": 1}; !maps.Equal(m, want) {
		t.Errorf("after Remove(map[one:1 zero:0], %v) the map is %v, want %v", keys, m, want)
	}
}

func TestReset(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	Reset(m)
	if len(m) != 0 {
		t.Errorf("after Reset(m) the map is %v, want an empty map", m)
	}
}

func TestSameWords(t *testing.T) {
	tests := []struct {
		a, b []string
		want bool
	}{
		{[]string{"x", "y", "x"}, []string{"x", "x", "y"}, true},
		{[]string{"a"}, []string{"a"}, true},
		{nil, nil, true},
		{[]string{"x", "y"}, []string{"x", "y", "y"}, false},
		{[]string{"x", "x", "y"}, []string{"x", "y", "y"}, false},
		{[]string{"a"}, []string{"b"}, false},
		{[]string{"a"}, nil, false},
	}
	for _, tt := range tests {
		if got := SameWords(tt.a, tt.b); got != tt.want {
			t.Errorf("SameWords(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
		}
	}
}
