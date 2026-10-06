package waitgroups

import (
	"slices"
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
	case <-time.After(2 * time.Second):
		t.Fatalf("%s did not return within 2 seconds", what)
	}
}

// meeting builds a function for n concurrent callers: every call of join
// blocks until n calls are in progress at the same time, then takes a
// little longer and returns true. That never happens when the calls are
// made one after the other; then join gives up after a while and returns
// false. arrived reports how many calls returned true so far.
func meeting(n int) (join func() bool, arrived func() int) {
	var started, finished atomic.Int32
	all := make(chan struct{})
	join = func() bool {
		if int(started.Add(1)) == n {
			close(all)
		}
		select {
		case <-all:
		case <-time.After(300 * time.Millisecond):
			return false
		}
		time.Sleep(20 * time.Millisecond)
		finished.Add(1)
		return true
	}
	return join, func() int { return int(finished.Load()) }
}

func TestRunAll(t *testing.T) {
	const n = 5
	join, arrived := meeting(n)
	tasks := make([]func(), n)
	for i := range tasks {
		tasks[i] = func() { join() }
	}

	finishes(t, "RunAll(tasks...)", func() { RunAll(tasks...) })
	if got := arrived(); got != n {
		t.Errorf("when RunAll(tasks...) returned, %d of the %d tasks had run concurrently and finished, want all of them", got, n)
	}

	finishes(t, "RunAll()", func() { RunAll() })
}

func TestWorker(t *testing.T) {
	// A single worker, called directly.
	var wg sync.WaitGroup
	var reported []int
	wg.Add(1)
	finishes(t, "Worker(7, &wg, report)", func() { Worker(7, &wg, func(id int) { reported = append(reported, id) }) })
	if want := []int{7}; !slices.Equal(reported, want) {
		t.Fatalf("Worker(7, &wg, report) called report with %v, want %v", reported, want)
	}
	finishes(t, "wg.Wait() after Worker returned (Worker has to call wg.Done())", wg.Wait)

	// Several workers as goroutines, the way Worker is meant to be used.
	const n = 5
	join, arrived := meeting(n)
	var sum atomic.Int32
	for id := 1; id <= n; id++ {
		wg.Add(1)
		go Worker(id, &wg, func(id int) {
			join()
			sum.Add(int32(id))
		})
	}
	finishes(t, "wg.Wait() for 5 workers (every Worker has to call wg.Done() once)", wg.Wait)
	if got := arrived(); got != n {
		t.Errorf("when wg.Wait() returned, %d of the %d workers had finished their work, want all of them - call wg.Done() after the work, not before", got, n)
	}
	if got := sum.Load(); got != 1+2+3+4+5 {
		t.Errorf("the ids that the 5 workers reported add up to %d, want 15 (each worker reports its own id)", got)
	}
}

func TestMapConcurrently(t *testing.T) {
	nums := []int{3, 1, 4, 1, 5, 9}
	want := []int{9, 1, 16, 1, 25, 81}
	join, arrived := meeting(len(nums))
	square := func(n int) int {
		join()
		return n * n
	}

	var got []int
	finishes(t, "MapConcurrently", func() { got = MapConcurrently(nums, square) })
	if !slices.Equal(got, want) {
		t.Errorf("MapConcurrently(%v, square) = %v, want %v", nums, got, want)
	}
	if n := arrived(); n != len(nums) {
		t.Errorf("MapConcurrently(%v, square): %d of the %d calls of square ran concurrently, want all of them", nums, n, len(nums))
	}

	finishes(t, "MapConcurrently without numbers", func() { got = MapConcurrently(nil, square) })
	if len(got) != 0 {
		t.Errorf("MapConcurrently(nil, square) = %v, want an empty result", got)
	}
}

func TestGather(t *testing.T) {
	send := func(nums ...int) func(chan<- int) {
		return func(out chan<- int) {
			for _, n := range nums {
				time.Sleep(time.Millisecond)
				out <- n
			}
		}
	}

	var got []int
	finishes(t, "Gather", func() { got = Gather(send(1, 2, 3), send(), send(10), send(20, 30, 40, 50)) })
	slices.Sort(got)
	if want := []int{1, 2, 3, 10, 20, 30, 40, 50}; !slices.Equal(got, want) {
		t.Errorf("Gather of producers sending (1 2 3), (), (10), (20 30 40 50) = (sorted) %v, want %v", got, want)
	}

	// The producers must run concurrently.
	const n = 4
	join, arrived := meeting(n)
	producers := make([]func(chan<- int), n)
	for i := range producers {
		producers[i] = func(out chan<- int) {
			if join() {
				out <- i
			}
		}
	}
	finishes(t, "Gather", func() { got = Gather(producers...) })
	slices.Sort(got)
	if want := []int{0, 1, 2, 3}; !slices.Equal(got, want) || arrived() != n {
		t.Errorf("Gather of 4 producers that only send when all of them run at the same time = (sorted) %v, want %v", got, want)
	}

	finishes(t, "Gather()", func() { got = Gather() })
	if len(got) != 0 {
		t.Errorf("Gather() = %v, want an empty result", got)
	}
}
