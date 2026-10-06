package arrays

import "testing"

func TestSquares(t *testing.T) {
	// Arrays of the same type can be compared with == and !=.
	if got, want := Squares(), [5]int{0, 1, 4, 9, 16}; got != want {
		t.Errorf("Squares() = %v, want %v", got, want)
	}
}

func TestSum(t *testing.T) {
	tests := []struct {
		nums [5]int
		want int
	}{
		{[5]int{1, 2, 3, 4, 5}, 15},
		{[5]int{10, -10, 7, 0, 0}, 7},
		{[5]int{0, 0, 0, 0, 100}, 100},
		{[5]int{}, 0},
	}
	for _, tt := range tests {
		if got := Sum(tt.nums); got != tt.want {
			t.Errorf("Sum(%v) = %d, want %d", tt.nums, got, tt.want)
		}
	}
}

func TestSparse(t *testing.T) {
	if got, want := Sparse(), [5]int{10, 0, 0, 40, 50}; got != want {
		t.Errorf("Sparse() = %v, want %v", got, want)
	}
}

func TestReversed(t *testing.T) {
	tests := []struct {
		a    [5]int
		want [5]int
	}{
		{[5]int{1, 2, 3, 4, 5}, [5]int{5, 4, 3, 2, 1}},
		{[5]int{7, 0, 0, 0, 9}, [5]int{9, 0, 0, 0, 7}},
		{[5]int{1, 1, 2, 1, 1}, [5]int{1, 1, 2, 1, 1}},
	}
	for _, tt := range tests {
		if got := Reversed(tt.a); got != tt.want {
			t.Errorf("Reversed(%v) = %v, want %v", tt.a, got, tt.want)
		}
	}
}

func TestTable(t *testing.T) {
	want := [3][3]int{
		{1, 2, 3},
		{2, 4, 6},
		{3, 6, 9},
	}
	if got := Table(); got != want {
		t.Errorf("Table() = %v, want %v", got, want)
	}
}

func TestTranspose(t *testing.T) {
	tests := []struct {
		m    [2][3]int
		want [3][2]int
	}{
		{
			[2][3]int{{1, 2, 3}, {4, 5, 6}},
			[3][2]int{{1, 4}, {2, 5}, {3, 6}},
		},
		{
			[2][3]int{{0, 0, 7}, {0, 0, 0}},
			[3][2]int{{0, 0}, {0, 0}, {7, 0}},
		},
	}
	for _, tt := range tests {
		if got := Transpose(tt.m); got != tt.want {
			t.Errorf("Transpose(%v) = %v, want %v", tt.m, got, tt.want)
		}
	}
}
