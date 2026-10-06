package signals

import (
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"testing"
	"time"
)

const timeout = 2 * time.Second

func TestNewSignalChannel(t *testing.T) {
	sigs := NewSignalChannel()
	if sigs == nil {
		t.Fatal("NewSignalChannel() returned a nil channel")
	}
	defer signal.Stop(sigs)
	if cap(sigs) != 1 {
		t.Errorf("the signal channel has capacity %d, want a buffered channel with capacity 1", cap(sigs))
	}
	select {
	case sig := <-sigs:
		t.Errorf("the signal channel already holds %v, want it to be empty until a signal arrives", sig)
	default:
	}
}

func TestAwaitSignal(t *testing.T) {
	sigs := make(chan os.Signal, 1)
	handled := make(chan os.Signal, 1)

	done := AwaitSignal(sigs, func(sig os.Signal) { handled <- sig })
	if done == nil {
		t.Fatal("AwaitSignal returned a nil done channel")
	}

	// Nothing may happen before the signal.
	select {
	case <-done:
		t.Fatal("done fired before any signal was sent")
	case sig := <-handled:
		t.Fatalf("onSignal was called with %v before any signal was sent", sig)
	case <-time.After(30 * time.Millisecond):
	}

	sigs <- syscall.SIGTERM

	select {
	case v := <-done:
		if !v {
			t.Error("received false from done, want true")
		}
	case <-time.After(timeout):
		t.Fatal("timed out waiting for done after sending a signal")
	}
	// onSignal must have run before done was signalled.
	select {
	case sig := <-handled:
		if sig != syscall.SIGTERM {
			t.Errorf("onSignal was called with %v, want %v", sig, syscall.SIGTERM)
		}
	default:
		t.Error("done fired but onSignal has not been called yet")
	}
}

func TestRunUntilSignal(t *testing.T) {
	// Unbuffered channels: every send below only completes once
	// RunUntilSignal is back in its select and receives it.
	sigs := make(chan os.Signal)
	ticks := make(chan time.Time)
	var calls atomic.Int32

	type result struct {
		count int
		sig   os.Signal
	}
	finished := make(chan result, 1)
	go func() {
		count, sig := RunUntilSignal(sigs, ticks, func() { calls.Add(1) })
		finished <- result{count, sig}
	}()

	for i := range 3 {
		select {
		case ticks <- time.Now():
		case r := <-finished:
			t.Fatalf("RunUntilSignal returned (%d, %v) before any signal was sent", r.count, r.sig)
		case <-time.After(timeout):
			t.Fatalf("RunUntilSignal did not receive tick number %d", i+1)
		}
	}

	select {
	case sigs <- os.Interrupt:
	case r := <-finished:
		t.Fatalf("RunUntilSignal returned (%d, %v) before any signal was sent", r.count, r.sig)
	case <-time.After(timeout):
		t.Fatal("RunUntilSignal did not receive the signal")
	}

	select {
	case r := <-finished:
		if r.count != 3 || r.sig != os.Interrupt {
			t.Errorf("RunUntilSignal returned (%d, %v) after 3 ticks and an interrupt, want (3, %v)", r.count, r.sig, os.Interrupt)
		}
		if got := calls.Load(); got != 3 {
			t.Errorf("work was called %d times for 3 ticks, want 3", got)
		}
	case <-time.After(timeout):
		t.Fatal("RunUntilSignal did not return after the signal")
	}
}

func TestRunUntilSignalStopsImmediately(t *testing.T) {
	sigs := make(chan os.Signal, 1)
	sigs <- syscall.SIGTERM

	type result struct {
		count int
		sig   os.Signal
	}
	finished := make(chan result, 1)
	go func() {
		count, sig := RunUntilSignal(sigs, make(chan time.Time), func() {})
		finished <- result{count, sig}
	}()

	select {
	case r := <-finished:
		if r.count != 0 || r.sig != syscall.SIGTERM {
			t.Errorf("RunUntilSignal returned (%d, %v) with a signal already waiting and no ticks, want (0, %v)", r.count, r.sig, syscall.SIGTERM)
		}
	case <-time.After(timeout):
		t.Fatal("RunUntilSignal did not return although a signal was waiting")
	}
}
