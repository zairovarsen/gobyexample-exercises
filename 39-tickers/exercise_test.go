package tickers

import (
	"sync/atomic"
	"testing"
	"time"
)

// The tests use short intervals to stay fast, and they are generous about
// the number of ticks: on a busy machine a ticker drops ticks rather than
// queueing them up.
const interval = 20 * time.Millisecond

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
	case <-time.After(3 * time.Second):
		t.Fatalf("%s did not return within 3 seconds", what)
	}
}

func TestTicks(t *testing.T) {
	const n = 3

	var got []time.Time
	start := time.Now()
	finishes(t, "Ticks", func() { got = Ticks(interval, n) })
	elapsed := time.Since(start)

	if len(got) != n {
		t.Fatalf("Ticks(%v, %d) returned %d times, want %d", interval, n, len(got), n)
	}
	if elapsed < n*interval {
		t.Errorf("Ticks(%v, %d) returned after %v, want it to wait for %d ticks (at least %v)", interval, n, elapsed, n, n*interval)
	}
	if elapsed > 10*n*interval {
		t.Errorf("Ticks(%v, %d) returned after %v, want about %v", interval, n, elapsed, n*interval)
	}
	prev := start
	for i, tick := range got {
		if !tick.After(prev) {
			t.Errorf("Ticks(%v, %d): tick %d has the time %v, want a time after %v (the times delivered by the ticker's channel, in order)", interval, n, i, tick, prev)
		}
		prev = tick
	}
}

func TestEvery(t *testing.T) {
	var calls atomic.Int32
	start := time.Now()
	stop := Every(interval, func() { calls.Add(1) })
	if stop == nil {
		t.Fatal("Every(interval, f) = nil, want a stop function")
	}
	if returned := time.Since(start); returned >= interval {
		t.Errorf("Every(%v, f) took %v to return, want it to return immediately", interval, returned)
	}

	time.Sleep(5*interval + interval/2)
	finishes(t, "the stop function returned by Every", stop)
	elapsed := time.Since(start)

	got := int(calls.Load())
	if most := int(elapsed/interval) + 1; got < 2 || got > most {
		t.Errorf("Every(%v, f) called f %d times in %v, want about one call per interval (between 2 and %d)", interval, got, elapsed, most)
	}

	time.Sleep(3 * interval)
	if after := int(calls.Load()); after != got {
		t.Errorf("Every(%v, f) called f %d more times after stop had returned, want no more calls", interval, after-got)
	}
}

func TestTickFor(t *testing.T) {
	const total = 5*interval + interval/2

	var got int
	start := time.Now()
	finishes(t, "TickFor", func() { got = TickFor(interval, total) })
	elapsed := time.Since(start)

	// 5 ticks fit into the total time. Allow for dropped ticks, and for one
	// more that shows up at the same moment as the end of the total time.
	if got < 2 || got > 6 {
		t.Errorf("TickFor(%v, %v) = %d, want about 5 (between 2 and 6)", interval, total, got)
	}
	if elapsed < total {
		t.Errorf("TickFor(%v, %v) returned after %v, want it to run for the total time", interval, total, elapsed)
	}
	if elapsed > 10*total {
		t.Errorf("TickFor(%v, %v) returned after %v, want it to stop when the total time is over", interval, total, elapsed)
	}
}

func TestPoll(t *testing.T) {
	// check succeeds on its third call.
	calls := 0
	check := func() bool {
		calls++
		return calls == 3
	}
	var got bool
	start := time.Now()
	finishes(t, "Poll", func() { got = Poll(interval, 2*time.Second, check) })
	elapsed := time.Since(start)
	if !got {
		t.Errorf("Poll(%v, 2s, check) with a check that succeeds on its third call = false, want true", interval)
	}
	if calls != 3 {
		t.Errorf("Poll(%v, 2s, check) called check %d times, want it to stop after the third call, which returns true", interval, calls)
	}
	if elapsed < 3*interval {
		t.Errorf("Poll(%v, 2s, check) returned after %v, want one check per interval (3 checks take at least %v)", interval, elapsed, 3*interval)
	}
	if elapsed > time.Second {
		t.Errorf("Poll(%v, 2s, check) returned after %v, want it to return as soon as check succeeds", interval, elapsed)
	}

	// check never succeeds.
	const timeout = 3*interval + interval/2
	calls = 0
	never := func() bool {
		calls++
		return false
	}
	start = time.Now()
	finishes(t, "Poll", func() { got = Poll(interval, timeout, never) })
	elapsed = time.Since(start)
	if got {
		t.Errorf("Poll(%v, %v, check) with a check that never succeeds = true, want false", interval, timeout)
	}
	if calls < 1 || calls > 4 {
		t.Errorf("Poll(%v, %v, check) called check %d times, want about 3 (one per interval until the timeout)", interval, timeout, calls)
	}
	if elapsed < timeout {
		t.Errorf("Poll(%v, %v, check) with a check that never succeeds returned after %v, want it to keep trying until the timeout", interval, timeout, elapsed)
	}
	if elapsed > 10*timeout {
		t.Errorf("Poll(%v, %v, check) with a check that never succeeds returned after %v, want it to give up at the timeout", interval, timeout, elapsed)
	}
}
