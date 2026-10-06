package atomiccounters

import (
	"reflect"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// finishes runs f and fails the test when f does not return in time,
// instead of hanging forever.
func finishes(t *testing.T, what string, f func()) {
	t.Helper()
	done := make(chan struct{})
	go func() {
		defer close(done)
		f()
	}()
	select {
	case <-done:
	case <-time.After(5 * time.Second):
		t.Fatalf("%s did not return within 5 seconds", what)
	}
}

// hammer calls f from the given number of goroutines, passing each one its
// number. The goroutines are all released at the same moment, to make them
// collide as much as possible, and hammer returns when all are finished.
func hammer(t *testing.T, goroutines int, f func(g int)) {
	t.Helper()
	start := make(chan struct{})
	var wg sync.WaitGroup
	for g := range goroutines {
		wg.Go(func() {
			<-start
			f(g)
		})
	}
	close(start)
	finishes(t, "the concurrent calls", wg.Wait)
}

// requireAtomicField fails the test unless the struct type has a field
// whose type comes from the package sync/atomic.
func requireAtomicField(t *testing.T, typ reflect.Type) {
	t.Helper()
	for i := range typ.NumField() {
		if typ.Field(i).Type.PkgPath() == "sync/atomic" {
			return
		}
	}
	t.Fatalf("%v has no field of a sync/atomic type such as atomic.Int64 - keep its state in one", typ)
}

func TestCountOps(t *testing.T) {
	tests := []struct {
		goroutines, perGoroutine int
		want                     uint64
	}{
		{50, 1000, 50000},
		{1, 7, 7},
		{8, 20000, 160000},
		{3, 0, 0},
		{0, 100, 0},
	}
	for _, tt := range tests {
		var got uint64
		finishes(t, "CountOps", func() { got = CountOps(tt.goroutines, tt.perGoroutine) })
		if got != tt.want {
			t.Errorf("CountOps(%d, %d) = %d, want %d", tt.goroutines, tt.perGoroutine, got, tt.want)
		}
	}
}

func TestCounter(t *testing.T) {
	requireAtomicField(t, reflect.TypeFor[Counter]())

	var c Counter
	if got := c.Value(); got != 0 {
		t.Errorf("new Counter: Value() = %d, want 0", got)
	}
	c.Inc()
	c.Add(10)
	c.Add(-4)
	if got := c.Value(); got != 7 {
		t.Errorf("Counter after Inc(), Add(10), Add(-4): Value() = %d, want 7", got)
	}

	const goroutines, perGoroutine = 50, 2000
	var shared Counter
	hammer(t, goroutines, func(int) {
		for range perGoroutine {
			shared.Inc()
			shared.Add(2)
		}
	})
	if got, want := shared.Value(), int64(goroutines*perGoroutine*3); got != want {
		t.Errorf("Counter after %d goroutines did Inc() and Add(2) %d times each: Value() = %d, want %d", goroutines, perGoroutine, got, want)
	}
}

func TestFlag(t *testing.T) {
	requireAtomicField(t, reflect.TypeFor[Flag]())

	var f Flag
	if f.IsSet() {
		t.Error("new Flag: IsSet() = true, want false")
	}
	if !f.TrySet() {
		t.Error("first TrySet() on a new Flag = false, want true")
	}
	if !f.IsSet() {
		t.Error("IsSet() after TrySet() = false, want true")
	}
	if f.TrySet() {
		t.Error("second TrySet() = true, want false")
	}

	// A race for the flag has exactly one winner, every time.
	const goroutines = 64
	for round := range 200 {
		var contested Flag
		var winners atomic.Int32
		hammer(t, goroutines, func(int) {
			if contested.TrySet() {
				winners.Add(1)
			}
		})
		if got := winners.Load(); got != 1 {
			t.Fatalf("round %d: of %d goroutines calling TrySet() at the same time, %d got true, want exactly 1", round, goroutines, got)
		}
		if !contested.IsSet() {
			t.Fatalf("round %d: IsSet() = false after %d goroutines called TrySet(), want true", round, goroutines)
		}
	}
}

func TestHighWater(t *testing.T) {
	requireAtomicField(t, reflect.TypeFor[HighWater]())

	var h HighWater
	if got := h.Max(); got != 0 {
		t.Errorf("new HighWater: Max() = %d, want 0", got)
	}
	steps := []struct {
		observe, want int64
	}{
		{5, 5},
		{3, 5},
		{-1, 5},
		{12, 12},
		{12, 12},
		{11, 12},
	}
	for _, s := range steps {
		h.Observe(s.observe)
		if got := h.Max(); got != s.want {
			t.Errorf("HighWater: Max() after Observe(%d) = %d, want %d", s.observe, got, s.want)
		}
	}

	// Each goroutine observes a different mix of numbers; the largest
	// number overall is goroutines*perGoroutine.
	const goroutines, perGoroutine = 50, 2000
	for round := range 20 {
		var shared HighWater
		hammer(t, goroutines, func(g int) {
			for i := 1; i <= perGoroutine; i++ {
				shared.Observe(int64(i*goroutines - g))
			}
		})
		if got, want := shared.Max(), int64(goroutines*perGoroutine); got != want {
			t.Fatalf("round %d: HighWater after %d goroutines observed numbers up to %d: Max() = %d, want %d", round, goroutines, want, got, want)
		}
	}
}
