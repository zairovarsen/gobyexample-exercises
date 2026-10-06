package channelsynchronization

import (
	"runtime"
	"slices"
	"strings"
	"sync"
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

// goroutineID returns the number of the calling goroutine, taken from the
// first line of its stack trace ("goroutine 18 [running]:"). Only tests
// should ever do this.
func goroutineID() string {
	buf := make([]byte, 64)
	buf = buf[:runtime.Stack(buf, false)]
	fields := strings.Fields(string(buf))
	if len(fields) < 2 {
		return "unknown"
	}
	return fields[1]
}

func TestWorker(t *testing.T) {
	done := make(chan bool, 1)
	ran := false
	signalledEarly := false
	work := func() {
		signalledEarly = len(done) > 0
		ran = true
	}

	finishes(t, "Worker(work, done)", func() { Worker(work, done) })

	if !ran {
		t.Error("Worker(work, done) did not call work")
	}
	if signalledEarly {
		t.Error("Worker(work, done) sent on done before calling work, want the signal after the work")
	}
	select {
	case v := <-done:
		if !v {
			t.Errorf("Worker(work, done) sent %v on done, want true", v)
		}
	default:
		t.Error("Worker(work, done) returned without sending on done")
	}
}

func TestStart(t *testing.T) {
	release := make(chan struct{})
	released := false
	defer func() {
		if !released {
			close(release)
		}
	}()
	var finished atomic.Bool
	work := func() {
		<-release
		finished.Store(true)
	}

	var done chan bool
	finishes(t, "Start(work), called with a work function that keeps running (it must not wait for work),", func() { done = Start(work) })
	if done == nil {
		t.Fatal("Start(work) = nil, want a channel made with make")
	}

	select {
	case <-done:
		t.Fatal("Start(work): the channel delivered a value while work was still running")
	case <-time.After(20 * time.Millisecond):
	}

	close(release)
	released = true
	if v := recv(t, done); !v {
		t.Errorf("Start(work): the channel delivered %v, want true", v)
	}
	if !finished.Load() {
		t.Error("Start(work): the channel delivered a value before work had finished")
	}
}

func TestWaitAll(t *testing.T) {
	const n = 4

	// Each work function blocks until all n are in progress at the same
	// time, which never happens when they run one after the other.
	var started, finished atomic.Int32
	all := make(chan struct{})
	work := func() {
		if started.Add(1) == n {
			close(all)
		}
		select {
		case <-all:
		case <-time.After(300 * time.Millisecond):
			return
		}
		time.Sleep(20 * time.Millisecond)
		finished.Add(1)
	}
	works := make([]func(), n)
	for i := range works {
		works[i] = work
	}

	finishes(t, "WaitAll(works...)", func() { WaitAll(works...) })
	if got := finished.Load(); got != n {
		t.Errorf("when WaitAll(works...) returned, %d of the %d functions had run concurrently and finished, want all of them", got, n)
	}

	finishes(t, "WaitAll()", func() { WaitAll() })
}

func TestRunSequence(t *testing.T) {
	const n = 5

	var mu sync.Mutex
	var order []int
	ids := map[string]bool{}
	steps := make([]func(), n)
	for i := range steps {
		steps[i] = func() {
			// Later steps are quicker, so they would overtake the earlier
			// ones if the steps did not wait for each other.
			time.Sleep(time.Duration(n-i) * 4 * time.Millisecond)
			mu.Lock()
			defer mu.Unlock()
			order = append(order, i)
			ids[goroutineID()] = true
		}
	}

	var callerID string
	finishes(t, "RunSequence(steps...)", func() {
		callerID = goroutineID()
		RunSequence(steps...)
	})

	mu.Lock()
	defer mu.Unlock()
	if want := []int{0, 1, 2, 3, 4}; !slices.Equal(order, want) {
		t.Errorf("when RunSequence returned, the steps had finished in the order %v, want %v", order, want)
	}
	if ids[callerID] {
		t.Error("RunSequence ran a step in the goroutine of its caller, want every step in a goroutine of its own")
	}
	if len(ids) != len(order) {
		t.Errorf("RunSequence ran its %d steps in %d different goroutines, want one goroutine per step", len(order), len(ids))
	}

	finishes(t, "RunSequence()", func() { RunSequence() })
}
