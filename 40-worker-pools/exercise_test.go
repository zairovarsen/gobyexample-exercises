package workerpools

import (
	"fmt"
	"slices"
	"strings"
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

// gauge measures how many calls are in progress at the same time. The work
// functions of the tests call enter when they begin and leave when they end.
type gauge struct {
	active, peak atomic.Int32
}

func (g *gauge) enter() {
	n := g.active.Add(1)
	for {
		p := g.peak.Load()
		if n <= p || g.peak.CompareAndSwap(p, n) {
			return
		}
	}
}

func (g *gauge) leave() { g.active.Add(-1) }

// requireBounded checks the peak that g saw against the size of the pool.
func requireBounded(t *testing.T, what string, g *gauge, workers int) {
	t.Helper()
	peak := int(g.peak.Load())
	if peak > workers {
		t.Errorf("%s had %d jobs in progress at the same time, want at most %d (the number of workers)", what, peak, workers)
	}
	if peak < 2 {
		t.Errorf("%s never had more than %d job in progress at a time, want the %d workers to work concurrently", what, peak, workers)
	}
}

func TestWorker(t *testing.T) {
	jobs := make(chan int, 5)
	results := make(chan int, 5)
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	finishes(t, "Worker(jobs, results) with jobs closed", func() { Worker(jobs, results) })

	if len(results) != 5 {
		t.Fatalf("Worker(jobs, results) returned after sending %d results for 5 jobs, want 5", len(results))
	}
	got := []int{<-results, <-results, <-results, <-results, <-results}
	if want := []int{2, 4, 6, 8, 10}; !slices.Equal(got, want) {
		t.Errorf("Worker(jobs, results) turned the jobs 1, 2, 3, 4, 5 into the results %v, want %v", got, want)
	}
}

func TestDoubleAll(t *testing.T) {
	tests := []struct {
		nums    []int
		workers int
		want    []int // sorted
	}{
		{[]int{1, 2, 3, 4, 5}, 3, []int{2, 4, 6, 8, 10}},
		{[]int{7, -1}, 5, []int{-2, 14}},
		{[]int{4, 4, 4}, 1, []int{8, 8, 8}},
		{nil, 3, nil},
	}
	for _, tt := range tests {
		var got []int
		finishes(t, "DoubleAll", func() { got = DoubleAll(tt.nums, tt.workers) })
		slices.Sort(got)
		if !slices.Equal(got, tt.want) {
			t.Errorf("DoubleAll(%v, %d) = (sorted) %v, want %v", tt.nums, tt.workers, got, tt.want)
		}
	}
}

func TestProcess(t *testing.T) {
	const workers = 4
	jobs := make([]int, 24)
	want := make([]int, 24)
	for i := range jobs {
		jobs[i] = i + 1
		want[i] = (i + 1) * (i + 1)
	}

	var g gauge
	square := func(j int) int {
		g.enter()
		defer g.leave()
		time.Sleep(5 * time.Millisecond)
		return j * j
	}

	var got []int
	finishes(t, "Process", func() { got = Process(jobs, workers, square) })
	slices.Sort(got)
	if !slices.Equal(got, want) {
		t.Fatalf("Process(1..24, %d, square) = (sorted) %v, want the squares of 1 to 24", workers, got)
	}
	requireBounded(t, fmt.Sprintf("Process(1..24, %d, square)", workers), &g, workers)

	finishes(t, "Process without jobs", func() { got = Process(nil, workers, square) })
	if len(got) != 0 {
		t.Errorf("Process(nil, %d, square) = %v, want an empty result", workers, got)
	}
}

func TestMapOrdered(t *testing.T) {
	const workers = 3
	var inputs, want []string
	for i := range 20 {
		s := fmt.Sprintf("word%02d", i)
		inputs = append(inputs, s)
		want = append(want, strings.ToUpper(s))
	}

	var g gauge
	var calls atomic.Int32
	upper := func(s string) string {
		g.enter()
		defer g.leave()
		// Uneven durations, so that the results finish out of order.
		time.Sleep(time.Duration(calls.Add(1)%4) * 2 * time.Millisecond)
		return strings.ToUpper(s)
	}

	var got []string
	finishes(t, "MapOrdered", func() { got = MapOrdered(inputs, workers, upper) })
	if !slices.Equal(got, want) {
		t.Fatalf("MapOrdered(word00..word19, %d, upper) = %q, want %q", workers, got, want)
	}
	requireBounded(t, fmt.Sprintf("MapOrdered(word00..word19, %d, upper)", workers), &g, workers)

	finishes(t, "MapOrdered without inputs", func() { got = MapOrdered(nil, workers, upper) })
	if len(got) != 0 {
		t.Errorf("MapOrdered(nil, %d, upper) = %q, want an empty result", workers, got)
	}
}
