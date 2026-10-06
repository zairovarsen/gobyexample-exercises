package multiplereturnvalues

import "testing"

func TestSwap(t *testing.T) {
	tests := []struct {
		a, b string
	}{
		{"a", "b"},
		{"hello", "world"},
		{"", "x"},
	}
	for _, tt := range tests {
		if x, y := Swap(tt.a, tt.b); x != tt.b || y != tt.a {
			t.Errorf("Swap(%q, %q) = %q, %q, want %q, %q", tt.a, tt.b, x, y, tt.b, tt.a)
		}
	}
}

func TestDivMod(t *testing.T) {
	tests := []struct {
		a, b         int
		wantQ, wantR int
	}{
		{17, 5, 3, 2},
		{20, 4, 5, 0},
		{3, 7, 0, 3},
		{100, 9, 11, 1},
	}
	for _, tt := range tests {
		if q, r := DivMod(tt.a, tt.b); q != tt.wantQ || r != tt.wantR {
			t.Errorf("DivMod(%d, %d) = %d, %d, want %d, %d", tt.a, tt.b, q, r, tt.wantQ, tt.wantR)
		}
	}
}

func TestClock(t *testing.T) {
	tests := []struct {
		seconds             int
		wantH, wantM, wantS int
	}{
		{3725, 1, 2, 5},
		{59, 0, 0, 59},
		{60, 0, 1, 0},
		{3600, 1, 0, 0},
		{86399, 23, 59, 59},
		{0, 0, 0, 0},
	}
	for _, tt := range tests {
		if h, m, s := Clock(tt.seconds); h != tt.wantH || m != tt.wantM || s != tt.wantS {
			t.Errorf("Clock(%d) = %d, %d, %d, want %d, %d, %d", tt.seconds, h, m, s, tt.wantH, tt.wantM, tt.wantS)
		}
	}
}

func TestMinMax(t *testing.T) {
	tests := []struct {
		nums             []int
		wantMin, wantMax int
	}{
		{[]int{3, 9, -2, 4}, -2, 9},
		{[]int{7}, 7, 7},
		{[]int{-5, -9, -1}, -9, -1},
		{[]int{1, 2, 3}, 1, 3},
		{[]int{3, 2, 1}, 1, 3},
		{nil, 0, 0},
	}
	for _, tt := range tests {
		if lo, hi := MinMax(tt.nums); lo != tt.wantMin || hi != tt.wantMax {
			t.Errorf("MinMax(%v) = %d, %d, want %d, %d", tt.nums, lo, hi, tt.wantMin, tt.wantMax)
		}
	}
}

func TestLargest(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{3, 9, -2, 4}, 9},
		{[]int{-5, -9, -1}, -1},
		{[]int{7}, 7},
		{nil, 0},
	}
	for _, tt := range tests {
		if got := Largest(tt.nums); got != tt.want {
			t.Errorf("Largest(%v) = %d, want %d", tt.nums, got, tt.want)
		}
	}
}

func TestWinner(t *testing.T) {
	tests := []struct {
		votes     []string
		wantName  string
		wantCount int
	}{
		{[]string{"ann", "bob", "ann", "cy", "bob", "ann"}, "ann", 3},
		{[]string{"bob"}, "bob", 1},
		{[]string{"cy", "bob", "bob"}, "bob", 2},
		// Ties: the candidate who appears first wins.
		{[]string{"bob", "ann", "ann", "bob"}, "bob", 2},
		{[]string{"cy", "ann", "bob"}, "cy", 1},
		{nil, "", 0},
	}
	for _, tt := range tests {
		if name, count := Winner(tt.votes); name != tt.wantName || count != tt.wantCount {
			t.Errorf("Winner(%v) = %q, %d, want %q, %d", tt.votes, name, count, tt.wantName, tt.wantCount)
		}
	}
}
