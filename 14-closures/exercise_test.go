package closures

import "testing"

func TestCounter(t *testing.T) {
	next := Counter()
	if next == nil {
		t.Fatal("Counter() = nil, want a function")
	}
	for want := 1; want <= 3; want++ {
		if got := next(); got != want {
			t.Errorf("call %d of the counter returned %d, want %d", want, got, want)
		}
	}

	other := Counter()
	if got := other(); got != 1 {
		t.Errorf("a second counter started at %d, want 1: every call of Counter needs its own variable", got)
	}
	if got := next(); got != 4 {
		t.Errorf("call 4 of the first counter returned %d, want 4: the second counter must not disturb it", got)
	}
}

func TestMultiplier(t *testing.T) {
	tests := []struct {
		factor, n int
		want      int
	}{
		{3, 5, 15},
		{2, 21, 42},
		{-1, 7, -7},
		{10, 0, 0},
	}
	for _, tt := range tests {
		mul := Multiplier(tt.factor)
		if mul == nil {
			t.Fatalf("Multiplier(%d) = nil, want a function", tt.factor)
		}
		if got := mul(tt.n); got != tt.want {
			t.Errorf("Multiplier(%d)(%d) = %d, want %d", tt.factor, tt.n, got, tt.want)
		}
	}
}

func TestAccumulator(t *testing.T) {
	acc := Accumulator(10)
	if acc == nil {
		t.Fatal("Accumulator(10) = nil, want a function")
	}
	steps := []struct {
		delta int
		want  int
	}{
		{5, 15},
		{-3, 12},
		{0, 12},
		{100, 112},
	}
	for _, s := range steps {
		if got := acc(s.delta); got != s.want {
			t.Errorf("acc(%d) = %d, want a running total of %d", s.delta, got, s.want)
		}
	}

	other := Accumulator(0)
	if got := other(1); got != 1 {
		t.Errorf("Accumulator(0)(1) = %d, want 1: accumulators must not share their total", got)
	}
}

func TestFibonacci(t *testing.T) {
	next := Fibonacci()
	if next == nil {
		t.Fatal("Fibonacci() = nil, want a function")
	}
	want := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	got := make([]int, len(want))
	for i := range got {
		got[i] = next()
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("the first %d Fibonacci numbers came out as %v, want %v", len(want), got, want)
		}
	}

	if got := Fibonacci()(); got != 0 {
		t.Errorf("a fresh generator started with %d, want 0", got)
	}
}

func TestCounterWithReset(t *testing.T) {
	next, reset := CounterWithReset()
	if next == nil || reset == nil {
		t.Fatal("CounterWithReset() returned a nil function, want two functions")
	}
	if a, b := next(), next(); a != 1 || b != 2 {
		t.Errorf("the first two calls of next returned %d, %d, want 1, 2", a, b)
	}
	reset()
	if got := next(); got != 1 {
		t.Errorf("next() after reset() = %d, want 1: both closures must use the same variable", got)
	}
	if got := next(); got != 2 {
		t.Errorf("second next() after reset() = %d, want 2", got)
	}
}

func TestMemoize(t *testing.T) {
	calls := 0
	square := func(n int) int {
		calls++
		return n * n
	}
	fast := Memoize(square)
	if fast == nil {
		t.Fatal("Memoize(f) = nil, want a function")
	}

	steps := []struct {
		n         int
		want      int
		wantCalls int
	}{
		{4, 16, 1},
		{4, 16, 1},
		{5, 25, 2},
		{4, 16, 2},
		// A result of 0 has to be cached like any other.
		{0, 0, 3},
		{0, 0, 3},
		{5, 25, 3},
	}
	for i, s := range steps {
		if got := fast(s.n); got != s.want {
			t.Errorf("step %d: memoized square(%d) = %d, want %d", i+1, s.n, got, s.want)
		}
		if calls != s.wantCalls {
			t.Fatalf("step %d: after memoized square(%d) the real function has run %d times, want %d", i+1, s.n, calls, s.wantCalls)
		}
	}

	// Two memoized functions must not share a cache.
	double := Memoize(func(n int) int { return 2 * n })
	if got := double(4); got != 8 {
		t.Errorf("a second memoized function returned %d for 4, want 8: every call of Memoize needs its own map", got)
	}
}
