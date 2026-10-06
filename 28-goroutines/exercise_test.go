package goroutines

import (
	"maps"
	"slices"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// The tests use channels, a mutex and atomics to observe the goroutines.
// You do not need any of them for the exercises - they are the subject of
// the following sections.

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

func TestRunAsync(t *testing.T) {
	t.Parallel()

	started := make(chan struct{}, 10)
	release := make(chan struct{})
	defer close(release)
	f := func() {
		started <- struct{}{}
		<-release // keeps f running until the test is over
	}

	finishes(t, "RunAsync(f), called with an f that keeps running (it must not wait for f),", func() { RunAsync(f) })

	select {
	case <-started:
	case <-time.After(2 * time.Second):
		t.Fatal("RunAsync(f) never called f")
	}
}

func TestStartEach(t *testing.T) {
	t.Parallel()

	names := []string{"ann", "bob", "cy", "dee"}
	started := make(chan string, 2*len(names))
	release := make(chan struct{})
	defer close(release)
	f := func(name string) {
		started <- name
		<-release // keeps every call running until the test is over
	}

	finishes(t, "StartEach(names, f), called with an f that keeps running (it must not wait for f),", func() { StartEach(names, f) })

	// No call of f returns before the test ends, so all four can only start
	// if each one has a goroutine of its own.
	var got []string
	for range names {
		select {
		case name := <-started:
			got = append(got, name)
		case <-time.After(2 * time.Second):
			t.Fatalf("StartEach(%q, f): only the calls of f for %q were started, want one running goroutine per name", names, got)
		}
	}
	slices.Sort(got)
	if !slices.Equal(got, names) {
		t.Errorf("StartEach(%q, f) called f with %q, want every name exactly once", names, got)
	}
}

func TestRunAndWait(t *testing.T) {
	const n = 5
	const wait = 50 * time.Millisecond

	// Each call of f blocks until all n calls are in progress at the same
	// time, which never happens when the functions run one after the other.
	var started, completed atomic.Int32
	all := make(chan struct{})
	f := func() {
		if started.Add(1) == n {
			close(all)
		}
		select {
		case <-all:
			completed.Add(1)
		case <-time.After(300 * time.Millisecond):
		}
	}
	fs := make([]func(), n)
	for i := range fs {
		fs[i] = f
	}

	start := time.Now()
	finishes(t, "RunAndWait", func() { RunAndWait(wait, fs...) })
	elapsed := time.Since(start)

	if got := completed.Load(); got != n {
		t.Errorf("when RunAndWait(%v, fs...) returned, %d of the %d functions had run concurrently and finished, want all of them - start each one with go", wait, got, n)
	}
	if elapsed < wait {
		t.Errorf("RunAndWait(%v, fs...) returned after %v, want it to sleep for at least %v", wait, elapsed, wait)
	}
	if elapsed > 10*wait {
		t.Errorf("RunAndWait(%v, fs...) returned after %v, want it to take only a little longer than %v", wait, elapsed, wait)
	}
}

func TestTree(t *testing.T) {
	t.Parallel()

	// Tree must return even while visit is still running.
	release := make(chan struct{})
	defer close(release)
	finishes(t, "Tree(1, visit), called with a visit that keeps running (it must not wait for visit),", func() {
		Tree(1, func(int) { <-release })
	})

	var mu sync.Mutex
	visits := map[int]int{}
	total := 0
	visit := func(level int) {
		mu.Lock()
		defer mu.Unlock()
		visits[level]++
		total++
	}
	finishes(t, "Tree(4, visit)", func() { Tree(4, visit) })

	// Wait for the 1 + 2 + 4 + 8 visits to come in.
	want := map[int]int{4: 1, 3: 2, 2: 4, 1: 8}
	deadline := time.Now().Add(2 * time.Second)
	for {
		mu.Lock()
		n := total
		mu.Unlock()
		if n >= 15 || time.Now().After(deadline) {
			break
		}
		time.Sleep(time.Millisecond)
	}
	time.Sleep(20 * time.Millisecond) // surplus visits would show up now

	mu.Lock()
	if !maps.Equal(visits, want) {
		t.Errorf("Tree(4, visit) led to these numbers of visits per level: %v, want %v", visits, want)
	}
	mu.Unlock()

	var stray atomic.Int32
	finishes(t, "Tree(0, visit)", func() { Tree(0, func(int) { stray.Add(1) }) })
	time.Sleep(20 * time.Millisecond)
	if got := stray.Load(); got != 0 {
		t.Errorf("Tree(0, visit) led to %d calls of visit, want none", got)
	}
}
