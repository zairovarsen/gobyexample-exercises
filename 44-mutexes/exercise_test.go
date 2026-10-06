package mutexes

import (
	"maps"
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
		t.Fatalf("%s did not return within 5 seconds - is a mutex locked and never unlocked?", what)
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

func TestContainer(t *testing.T) {
	c := NewContainer()
	if c == nil {
		t.Fatal("NewContainer() = nil, want a Container")
	}

	if got := c.Get("a"); got != 0 {
		t.Errorf(`new Container: Get("a") = %d, want 0`, got)
	}
	c.Inc("a")
	c.Inc("a")
	c.Inc("b")
	if a, b := c.Get("a"), c.Get("b"); a != 2 || b != 1 {
		t.Errorf(`after Inc("a"), Inc("a"), Inc("b"): Get("a"), Get("b") = %d, %d, want 2, 1`, a, b)
	}

	// Many goroutines at once, spread over three counters.
	// (Without a mutex this is where Go reports "concurrent map writes".)
	shared := NewContainer()
	names := []string{"x", "y", "z"}
	const goroutines, perGoroutine = 30, 1000
	hammer(t, goroutines, func(g int) {
		for range perGoroutine {
			shared.Inc(names[g%len(names)])
			shared.Get(names[(g+1)%len(names)])
		}
	})
	want := map[string]int{"x": 10000, "y": 10000, "z": 10000}
	snapshot := shared.Snapshot()
	if !maps.Equal(snapshot, want) {
		t.Fatalf("after %d goroutines did %d Incs each, spread over 3 counters: Snapshot() = %v, want %v", goroutines, perGoroutine, snapshot, want)
	}
	for _, name := range names {
		if got := shared.Get(name); got != 10000 {
			t.Errorf("after the concurrent Incs: Get(%q) = %d, want 10000", name, got)
		}
	}

	// The snapshot is a copy.
	snapshot["x"] = -1
	snapshot["new"] = 1
	if x, n := shared.Get("x"), shared.Get("new"); x != 10000 || n != 0 {
		t.Errorf(`changing the map returned by Snapshot() changed the Container: Get("x"), Get("new") = %d, %d, want 10000, 0`, x, n)
	}
}

func TestIncMany(t *testing.T) {
	c := NewContainer()
	if c == nil {
		t.Fatal("NewContainer() = nil, want a Container")
	}
	c.Inc("a")

	finishes(t, "IncMany", func() { IncMany(c, "a", 50, 200) })
	if got := c.Get("a"); got != 10001 {
		t.Errorf(`IncMany(c, "a", 50, 200) on a counter at 1: Get("a") = %d when it returned, want 10001`, got)
	}
	if got := c.Get("b"); got != 0 {
		t.Errorf(`IncMany(c, "a", 50, 200): Get("b") = %d, want 0`, got)
	}
}

func TestAccount(t *testing.T) {
	a := NewAccount(100)
	if a == nil {
		t.Fatal("NewAccount(100) = nil, want an Account")
	}
	if got := a.Balance(); got != 100 {
		t.Errorf("NewAccount(100): Balance() = %d, want 100", got)
	}
	a.Deposit(50)
	if got := a.Balance(); got != 150 {
		t.Errorf("after Deposit(50): Balance() = %d, want 150", got)
	}
	if a.Withdraw(151) {
		t.Error("Withdraw(151) with a balance of 150 = true, want false")
	}
	if got := a.Balance(); got != 150 {
		t.Errorf("after the refused Withdraw(151): Balance() = %d, want 150", got)
	}
	if !a.Withdraw(150) {
		t.Error("Withdraw(150) with a balance of 150 = false, want true")
	}
	if got := a.Balance(); got != 0 {
		t.Errorf("after Withdraw(150): Balance() = %d, want 0", got)
	}

	// 200 goroutines try to withdraw 10 each from 1000: only 100 can succeed.
	for round := range 20 {
		shared := NewAccount(1000)
		var succeeded atomic.Int32
		hammer(t, 200, func(int) {
			if shared.Withdraw(10) {
				succeeded.Add(1)
			}
		})
		if got, balance := succeeded.Load(), shared.Balance(); got != 100 || balance != 0 {
			t.Fatalf("round %d: of 200 concurrent Withdraw(10) calls on a balance of 1000, %d succeeded and the balance is %d, want 100 and 0", round, got, balance)
		}

		hammer(t, 100, func(int) { shared.Deposit(5) })
		if got := shared.Balance(); got != 500 {
			t.Fatalf("round %d: after 100 concurrent Deposit(5) calls on an empty account: Balance() = %d, want 500", round, got)
		}
	}
}

func TestCache(t *testing.T) {
	c := NewCache()
	if c == nil {
		t.Fatal("NewCache() = nil, want a Cache")
	}

	if got := c.GetOrCompute("seven", func() int { return 7 }); got != 7 {
		t.Errorf(`GetOrCompute("seven", compute) with a compute that returns 7 = %d, want 7`, got)
	}
	if got := c.GetOrCompute("seven", func() int { return 8 }); got != 7 {
		t.Errorf(`second GetOrCompute("seven", ...) = %d, want the stored value 7`, got)
	}
	// A stored zero is a stored value, too.
	c.GetOrCompute("zero", func() int { return 0 })
	if got := c.GetOrCompute("zero", func() int { return 1 }); got != 0 {
		t.Errorf(`GetOrCompute("zero", ...) after 0 was stored for "zero" = %d, want the stored value 0`, got)
	}

	// Many goroutines want the same missing key at the same moment.
	var computed, wrong atomic.Int32
	slow := func() int {
		computed.Add(1)
		time.Sleep(10 * time.Millisecond)
		return 42
	}
	hammer(t, 50, func(int) {
		if c.GetOrCompute("answer", slow) != 42 {
			wrong.Add(1)
		}
	})
	if got := computed.Load(); got != 1 {
		t.Errorf(`50 concurrent GetOrCompute("answer", compute) calls called compute %d times, want once`, got)
	}
	if got := wrong.Load(); got != 0 {
		t.Errorf(`%d of 50 concurrent GetOrCompute("answer", compute) calls returned something other than 42`, got)
	}
}
