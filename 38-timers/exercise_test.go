package timers

import (
	"sync/atomic"
	"testing"
	"time"
)

// recv receives one value from ch, failing the test instead of hanging
// forever when nothing arrives.
func recv[T any](t *testing.T, ch <-chan T) T {
	t.Helper()
	if ch == nil {
		t.Fatal("got a nil channel - did you create it with make?")
	}
	select {
	case v := <-ch:
		return v
	case <-time.After(2 * time.Second):
		t.Fatal("timed out waiting for a value on the channel")
	}
	panic("unreachable")
}

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
	case <-time.After(2 * time.Second):
		t.Fatalf("%s did not return within 2 seconds", what)
	}
}

func TestWaitTimer(t *testing.T) {
	const d = 40 * time.Millisecond

	var got time.Time
	start := time.Now()
	finishes(t, "WaitTimer", func() { got = WaitTimer(d) })
	elapsed := time.Since(start)

	if elapsed < d {
		t.Errorf("WaitTimer(%v) returned after %v, want it to block until the timer fires", d, elapsed)
	}
	if elapsed > 10*d {
		t.Errorf("WaitTimer(%v) returned after %v, want it to return soon after %v", d, elapsed, d)
	}
	if got.IsZero() {
		t.Errorf("WaitTimer(%v) returned the zero time, want the time received from the timer's channel", d)
	} else if got.Before(start) || got.After(time.Now()) {
		t.Errorf("WaitTimer(%v) returned %v, want the time at which the timer fired", d, got)
	}
}

func TestSchedule(t *testing.T) {
	const d = 40 * time.Millisecond

	// Left alone, the timer fires and f is called.
	called := make(chan time.Time, 10)
	start := time.Now()
	var timer *time.Timer
	finishes(t, "Schedule (it must not wait for the timer)", func() { timer = Schedule(d, func() { called <- time.Now() }) })
	if timer == nil {
		t.Fatal("Schedule(d, f) = nil, want the timer")
	}
	if returned := time.Since(start); returned >= d {
		t.Errorf("Schedule(%v, f) took %v to return, want it to return immediately", d, returned)
	}
	at := recv(t, called)
	if waited := at.Sub(start); waited < d {
		t.Errorf("Schedule(%v, f) called f after %v, want it to wait for the timer", d, waited)
	}

	// Stopped in time, the timer never fires and f is not called.
	var fired atomic.Bool
	finishes(t, "Schedule (it must not wait for the timer)", func() { timer = Schedule(d, func() { fired.Store(true) }) })
	if timer == nil {
		t.Fatal("Schedule(d, f) = nil, want the timer")
	}
	if !timer.Stop() {
		t.Errorf("Stop on the timer returned by Schedule(%v, f) = false right after the call, want true (a timer that is still running)", d)
	}
	time.Sleep(3 * d)
	if fired.Load() {
		t.Errorf("Schedule(%v, f) called f although the returned timer was stopped - wait on the channel of the same timer that you return", d)
	}
}

func TestTryCancel(t *testing.T) {
	running := time.NewTimer(time.Hour)
	if got, want := TryCancel(running), "stopped"; got != want {
		t.Errorf("TryCancel of a timer with an hour to go = %q, want %q", got, want)
	}
	// Whatever TryCancel did, this timer is stopped now.
	if running.Stop() {
		t.Error("TryCancel left the timer running, want it stopped")
	}

	fired := time.NewTimer(time.Millisecond)
	recv(t, fired.C) // wait until it has fired
	if got, want := TryCancel(fired), "already fired"; got != want {
		t.Errorf("TryCancel of a timer that has fired = %q, want %q", got, want)
	}
}

func TestDebounce(t *testing.T) {
	const d = 100 * time.Millisecond

	called := make(chan time.Time, 100)
	trigger := Debounce(d, func() { called <- time.Now() })
	if trigger == nil {
		t.Fatal("Debounce(d, f) = nil, want a function")
	}

	// A burst of triggers, much closer together than d.
	var last time.Time
	for range 5 {
		last = time.Now()
		trigger()
		time.Sleep(5 * time.Millisecond)
	}
	select {
	case at := <-called:
		if quiet := at.Sub(last); quiet < d {
			t.Errorf("Debounce(%v, f): f was called %v after the last trigger, want it to wait for %v without triggers", d, quiet, d)
		}
	case <-time.After(2 * time.Second):
		t.Fatalf("Debounce(%v, f): f was not called within 2 seconds after a burst of triggers", d)
	}
	select {
	case <-called:
		t.Errorf("Debounce(%v, f): f was called more than once for one burst of 5 triggers - every trigger has to cancel the countdown of the previous one", d)
	case <-time.After(d + d/2):
	}

	// After f has run, a new trigger starts all over.
	last = time.Now()
	trigger()
	select {
	case at := <-called:
		if quiet := at.Sub(last); quiet < d {
			t.Errorf("Debounce(%v, f): after a later trigger f was called again after only %v, want %v", d, quiet, d)
		}
	case <-time.After(2 * time.Second):
		t.Fatalf("Debounce(%v, f): a trigger after the first call of f did not lead to another call within 2 seconds", d)
	}
}
