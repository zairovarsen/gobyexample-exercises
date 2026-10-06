package ratelimiting

import (
	"slices"
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
	case <-time.After(3 * time.Second):
		t.Fatalf("%s did not return within 3 seconds", what)
	}
}

// queue returns a closed channel holding the requests 1 to n.
func queue(n int) <-chan int {
	requests := make(chan int, n)
	for i := 1; i <= n; i++ {
		requests <- i
	}
	close(requests)
	return requests
}

// handled is one call of a handle function: the request and when it came.
type handled struct {
	req   int
	after time.Duration // since the start of the measurement
}

// recorder returns a handle function and the slice it logs its calls to.
// The handle function is only called from one goroutine at a time.
func recorder(start time.Time) (handle func(int), log *[]handled) {
	log = new([]handled)
	return func(req int) { *log = append(*log, handled{req, time.Since(start)}) }, log
}

func reqs(log []handled) []int {
	var rs []int
	for _, h := range log {
		rs = append(rs, h.req)
	}
	return rs
}

func TestThrottle(t *testing.T) {
	const interval = 25 * time.Millisecond
	const n = 4

	start := time.Now()
	handle, log := recorder(start)
	finishes(t, "Throttle", func() { Throttle(queue(n), interval, handle) })

	if got, want := reqs(*log), []int{1, 2, 3, 4}; !slices.Equal(got, want) {
		t.Fatalf("Throttle handled the requests %v, want %v", got, want)
	}
	// Request number i has to wait for tick number i.
	for i, h := range *log {
		if earliest := time.Duration(i+1) * interval; h.after < earliest {
			t.Errorf("Throttle(requests, %v, handle) handled request %d after %v, want one request per interval (not before %v)", interval, h.req, h.after, earliest)
		}
	}
	if total := (*log)[n-1].after; total > 10*n*interval {
		t.Errorf("Throttle(requests, %v, handle) took %v for %d requests, want about %v", interval, total, n, n*interval)
	}
}

func TestNewBurstyLimiter(t *testing.T) {
	const interval = 50 * time.Millisecond
	const burst = 3

	start := time.Now()
	limiter, stop := NewBurstyLimiter(burst, interval)
	if limiter == nil || stop == nil {
		t.Fatal("NewBurstyLimiter returned a nil limiter or a nil stop function")
	}
	stopped := false
	defer func() {
		if !stopped {
			stop()
		}
	}()

	if cap(limiter) != burst {
		t.Errorf("NewBurstyLimiter(%d, %v): the limiter has capacity %d, want %d", burst, interval, cap(limiter), burst)
	}
	// The bucket starts out full.
	for i := range burst {
		select {
		case <-limiter:
		default:
			t.Fatalf("NewBurstyLimiter(%d, %v): only %d tokens were available right away, want %d", burst, interval, i, burst)
		}
	}
	// The next token has to be refilled.
	select {
	case <-limiter:
		t.Fatalf("NewBurstyLimiter(%d, %v): a token number %d was available right away, want the bucket to be empty after %d", burst, interval, burst+1, burst)
	default:
	}
	recv(t, limiter)
	if waited := time.Since(start); waited < interval {
		t.Errorf("NewBurstyLimiter(%d, %v): the first refilled token came after %v, want one token per interval", burst, interval, waited)
	} else if waited > 10*interval {
		t.Errorf("NewBurstyLimiter(%d, %v): the first refilled token came after %v, want it after about %v", burst, interval, waited, interval)
	}

	// After stop the bucket is not refilled any more.
	finishes(t, "the stop function returned by NewBurstyLimiter", stop)
	stopped = true
	time.Sleep(2 * interval)
	for range burst + 1 { // empty the bucket
		select {
		case <-limiter:
		default:
		}
	}
	time.Sleep(3 * interval)
	select {
	case <-limiter:
		t.Errorf("NewBurstyLimiter(%d, %v): a token was put into the bucket after stop had been called", burst, interval)
	default:
	}
}

func TestBurstyThrottle(t *testing.T) {
	const interval = 100 * time.Millisecond
	const burst = 3
	const n = 5

	start := time.Now()
	handle, log := recorder(start)
	finishes(t, "BurstyThrottle", func() { BurstyThrottle(queue(n), burst, interval, handle) })

	if got, want := reqs(*log), []int{1, 2, 3, 4, 5}; !slices.Equal(got, want) {
		t.Fatalf("BurstyThrottle handled the requests %v, want %v", got, want)
	}
	for i, h := range *log {
		if i < burst {
			if h.after >= interval {
				t.Errorf("BurstyThrottle(requests, %d, %v, handle) handled request %d after %v, want the first %d requests without delay", burst, interval, h.req, h.after, burst)
			}
			continue
		}
		// Request number burst+k has to wait for tick number k.
		if earliest := time.Duration(i-burst+1) * interval; h.after < earliest {
			t.Errorf("BurstyThrottle(requests, %d, %v, handle) handled request %d after %v, want one request per interval after the burst (not before %v)", burst, interval, h.req, h.after, earliest)
		}
	}
	if total := (*log)[n-1].after; total > 10*(n-burst)*interval {
		t.Errorf("BurstyThrottle(requests, %d, %v, handle) took %v for %d requests, want about %v", burst, interval, total, n, (n-burst)*interval)
	}
}

func TestShed(t *testing.T) {
	tokens := func(n int) chan time.Time {
		limiter := make(chan time.Time, 5)
		for range n {
			limiter <- time.Now()
		}
		return limiter
	}
	tests := []struct {
		tokens, requests int
		wantHandled      []int
		wantDropped      int
		wantLeft         int // tokens left in the limiter
	}{
		{2, 5, []int{1, 2}, 3, 0},
		{5, 3, []int{1, 2, 3}, 0, 2},
		{0, 4, nil, 4, 0},
		{3, 0, nil, 0, 3},
	}
	for _, tt := range tests {
		limiter := tokens(tt.tokens)
		handle, log := recorder(time.Now())
		var dropped int
		finishes(t, "Shed", func() { dropped = Shed(queue(tt.requests), limiter, handle) })

		if got := reqs(*log); !slices.Equal(got, tt.wantHandled) || dropped != tt.wantDropped {
			t.Errorf("Shed with %d requests and %d tokens handled %v and dropped %d, want it to handle %v and drop %d",
				tt.requests, tt.tokens, got, dropped, tt.wantHandled, tt.wantDropped)
		}
		if len(limiter) != tt.wantLeft {
			t.Errorf("Shed with %d requests and %d tokens left %d tokens in the limiter, want %d", tt.requests, tt.tokens, len(limiter), tt.wantLeft)
		}
	}
}
