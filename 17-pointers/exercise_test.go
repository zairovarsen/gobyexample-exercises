package pointers

import "testing"

func TestZeroPtr(t *testing.T) {
	for _, start := range []int{1, 5, -42} {
		i := start
		ZeroPtr(&i)
		if i != 0 {
			t.Errorf("i := %d; ZeroPtr(&i) left i = %d, want 0", start, i)
		}
	}
}

func TestSwap(t *testing.T) {
	tests := []struct {
		x, y int
	}{
		{1, 2},
		{-7, 30},
		{0, 9},
	}
	for _, tt := range tests {
		x, y := tt.x, tt.y
		Swap(&x, &y)
		if x != tt.y || y != tt.x {
			t.Errorf("x, y := %d, %d; Swap(&x, &y) left x, y = %d, %d, want %d, %d", tt.x, tt.y, x, y, tt.y, tt.x)
		}
	}
}

func TestNewInt(t *testing.T) {
	p, q := NewInt(5), NewInt(5)
	if p == nil || q == nil {
		t.Fatal("NewInt(5) = nil, want a pointer to an int")
	}
	if *p != 5 {
		t.Errorf("*NewInt(5) = %d, want 5", *p)
	}
	if p == q {
		t.Error("two calls of NewInt returned the same pointer, want pointers to two different variables")
	}
	*p = 6
	if *q != 5 {
		t.Errorf("changing the int behind one pointer changed the other one to %d", *q)
	}

	if r := NewInt(-3); r == nil || *r != -3 {
		t.Errorf("NewInt(-3) does not point to -3")
	}
}

func TestValueOr(t *testing.T) {
	seven, zero := 7, 0
	tests := []struct {
		name     string
		p        *int
		fallback int
		want     int
	}{
		{"pointer to 7", &seven, 1, 7},
		{"pointer to 0", &zero, 1, 0},
		{"nil", nil, 1, 1},
		{"nil", nil, -9, -9},
	}
	for _, tt := range tests {
		if got := ValueOr(tt.p, tt.fallback); got != tt.want {
			t.Errorf("ValueOr(%s, %d) = %d, want %d", tt.name, tt.fallback, got, tt.want)
		}
	}
}

func TestFillArray(t *testing.T) {
	a := [3]int{1, 2, 3}
	FillArray(&a, 7)
	if want := [3]int{7, 7, 7}; a != want {
		t.Errorf("a := [1 2 3]; FillArray(&a, 7) left a = %v, want %v", a, want)
	}
	FillArray(&a, 0)
	if want := [3]int{}; a != want {
		t.Errorf("FillArray(&a, 0) left a = %v, want %v", a, want)
	}
}

func TestLargest(t *testing.T) {
	tests := []struct {
		nums      []int
		wantIndex int
	}{
		{[]int{3, 9, 4}, 1},
		{[]int{8, 1, 2}, 0},
		{[]int{-5, -9, -1}, 2},
		{[]int{4, 7, 7, 1}, 1},
		{[]int{6}, 0},
	}
	for _, tt := range tests {
		p := Largest(tt.nums)
		if p == nil {
			t.Fatalf("Largest(%v) = nil, want a pointer to element %d", tt.nums, tt.wantIndex)
		}
		if want := tt.nums[tt.wantIndex]; *p != want {
			t.Errorf("Largest(%v) points to %d, want %d", tt.nums, *p, want)
		}
		if p != &tt.nums[tt.wantIndex] {
			t.Errorf("Largest(%v) does not point to element %d of the slice itself (is it the address of a copy?)", tt.nums, tt.wantIndex)
		}
	}

	nums := []int{3, 9, 4}
	if p := Largest(nums); p != nil {
		*p = 0
	}
	if want := [3]int{3, 0, 4}; [3]int(nums) != want {
		t.Errorf("nums := [3 9 4]; *Largest(nums) = 0 left nums = %v, want %v", nums, want)
	}

	if p := Largest(nil); p != nil {
		t.Errorf("Largest(nil) = pointer to %d, want nil", *p)
	}
}
