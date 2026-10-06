package methods

import "testing"

func TestArea(t *testing.T) {
	tests := []struct {
		r    Rect
		want int
	}{
		{Rect{Width: 10, Height: 5}, 50},
		{Rect{Width: 3, Height: 3}, 9},
		{Rect{Width: 7, Height: 1}, 7},
		{Rect{Width: 0, Height: 4}, 0},
	}
	for _, tt := range tests {
		if got := tt.r.Area(); got != tt.want {
			t.Errorf("%+v.Area() = %d, want %d", tt.r, got, tt.want)
		}
	}

	// Methods with a value receiver can be called through a pointer too.
	rp := &Rect{Width: 2, Height: 4}
	if got := rp.Area(); got != 8 {
		t.Errorf("rp := &Rect{2, 4}; rp.Area() = %d, want 8", got)
	}
}

func TestPerimeter(t *testing.T) {
	tests := []struct {
		r    Rect
		want int
	}{
		{Rect{Width: 10, Height: 5}, 30},
		{Rect{Width: 3, Height: 3}, 12},
		{Rect{Width: 7, Height: 1}, 16},
	}
	for _, tt := range tests {
		if got := tt.r.Perimeter(); got != tt.want {
			t.Errorf("%+v.Perimeter() = %d, want %d", tt.r, got, tt.want)
		}
	}
}

func TestScale(t *testing.T) {
	r := Rect{Width: 2, Height: 3}
	r.Scale(10)
	if want := (Rect{Width: 20, Height: 30}); r != want {
		t.Errorf("r := Rect{2, 3}; r.Scale(10) left r = %+v, want %+v", r, want)
	}

	rp := &Rect{Width: 1, Height: 5}
	rp.Scale(3)
	if want := (Rect{Width: 3, Height: 15}); *rp != want {
		t.Errorf("rp := &Rect{1, 5}; rp.Scale(3) left *rp = %+v, want %+v", *rp, want)
	}
}

func TestScaled(t *testing.T) {
	r := Rect{Width: 2, Height: 3}
	got := r.Scaled(10)
	if want := (Rect{Width: 20, Height: 30}); got != want {
		t.Errorf("Rect{2, 3}.Scaled(10) = %+v, want %+v", got, want)
	}
	if want := (Rect{Width: 2, Height: 3}); r != want {
		t.Errorf("r.Scaled(10) changed r to %+v, want it to stay %+v", r, want)
	}
}

func TestCounter(t *testing.T) {
	// The test asks at run time whether the methods exist (with type
	// assertions, a topic of the next section), so that this file compiles
	// before you have declared them.
	c := &Counter{}
	incrementer, ok := any(c).(interface{ Increment() })
	if !ok {
		t.Fatal("Counter has no method Increment() - declare it on Counter (which receiver type does it need?)")
	}
	valuer, ok := any(c).(interface{ Value() int })
	if !ok {
		t.Fatal("Counter has no method Value() int - declare it on Counter")
	}

	if got := valuer.Value(); got != 0 {
		t.Errorf("Value() of a new Counter = %d, want 0", got)
	}
	for range 3 {
		incrementer.Increment()
	}
	if got := valuer.Value(); got != 3 {
		t.Errorf("Value() after 3 calls of Increment() = %d, want 3 (with a value receiver Increment only changes a copy)", got)
	}

	// Counters are independent of each other.
	other := &Counter{}
	if got := any(other).(interface{ Value() int }).Value(); got != 0 {
		t.Errorf("Value() of a second, new Counter = %d, want 0", got)
	}
}

func TestStack(t *testing.T) {
	var s Stack
	if got := s.Len(); got != 0 {
		t.Errorf("Len() of a new Stack = %d, want 0", got)
	}

	for _, v := range []int{1, 2, 3} {
		s.Push(v)
	}
	if got := s.Len(); got != 3 {
		t.Fatalf("Len() after pushing 1, 2, 3 = %d, want 3", got)
	}

	for i, want := range []int{3, 2} {
		if got, ok := s.Pop(); got != want || !ok {
			t.Errorf("pop number %d returned %d, %v, want %d, true", i+1, got, ok, want)
		}
	}
	if got := s.Len(); got != 1 {
		t.Errorf("Len() after pushing three numbers and popping two = %d, want 1", got)
	}

	s.Push(0)
	if got, ok := s.Pop(); got != 0 || !ok {
		t.Errorf("Pop() after Push(0) = %d, %v, want 0, true", got, ok)
	}
	if got, ok := s.Pop(); got != 1 || !ok {
		t.Errorf("Pop() of the last number = %d, %v, want 1, true", got, ok)
	}

	if got, ok := s.Pop(); got != 0 || ok {
		t.Errorf("Pop() on an empty Stack = %d, %v, want 0, false", got, ok)
	}
	if got := s.Len(); got != 0 {
		t.Errorf("Len() of the emptied Stack = %d, want 0", got)
	}
}
