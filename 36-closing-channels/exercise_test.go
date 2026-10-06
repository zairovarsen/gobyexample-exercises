package closingchannels

import (
	"slices"
	"testing"
	"time"
)

// recv does a two-value receive from ch: it returns the value and whether
// the channel was still open. It fails the test instead of hanging forever
// when nothing arrives and the channel is not closed either.
func recv[T any](t *testing.T, ch <-chan T) (T, bool) {
	t.Helper()
	if ch == nil {
		t.Fatal("got a nil channel - did you create it with make?")
	}
	select {
	case v, ok := <-ch:
		return v, ok
	case <-time.After(2 * time.Second):
		t.Fatal("timed out on the channel: no value arrived and it was not closed either")
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

func TestSendAndClose(t *testing.T) {
	nums := []int{7, 0, 9}
	ch := SendAndClose(nums...)
	if ch == nil {
		t.Fatal("SendAndClose(7, 0, 9) = nil, want a channel made with make")
	}
	if len(ch) != len(nums) {
		t.Fatalf("SendAndClose(7, 0, 9) returned a channel holding %d values, want all 3 in the buffer", len(ch))
	}
	for i, want := range nums {
		got, ok := recv(t, ch)
		if got != want || !ok {
			t.Errorf("SendAndClose(7, 0, 9): receive %d = %d, %v, want %d, true", i+1, got, ok, want)
		}
	}
	// With the buffer empty, the receivers get to see that it is closed.
	if got, ok := recv(t, ch); got != 0 || ok {
		t.Errorf("SendAndClose(7, 0, 9): receive after the last value = %d, %v, want 0, false (a closed channel)", got, ok)
	}

	if _, ok := recv(t, SendAndClose()); ok {
		t.Error("SendAndClose(): the receive reports an open channel, want a closed one")
	}
}

func TestReceiveOr(t *testing.T) {
	ch := make(chan string, 2)
	ch <- "one"
	ch <- "" // an empty string is a value like any other
	close(ch)

	for i, want := range []string{"one", "", "default", "default"} {
		var got string
		finishes(t, "ReceiveOr", func() { got = ReceiveOr(ch, "default") })
		if got != want {
			t.Errorf(`call %d of ReceiveOr(ch, "default") on a closed channel that held "one", "" = %q, want %q`, i+1, got, want)
		}
	}
}

func TestSumUntilClosed(t *testing.T) {
	tests := []struct {
		nums      []int
		wantSum   int
		wantCount int
	}{
		{[]int{1, 2, 3}, 6, 3},
		{[]int{5, 0, 0, 7}, 12, 4},
		{[]int{0}, 0, 1},
		{nil, 0, 0},
	}
	for _, tt := range tests {
		ch := make(chan int, len(tt.nums))
		for _, n := range tt.nums {
			ch <- n
		}
		close(ch)

		var sum, count int
		finishes(t, "SumUntilClosed", func() { sum, count = SumUntilClosed(ch) })
		if sum != tt.wantSum || count != tt.wantCount {
			t.Errorf("SumUntilClosed of a closed channel holding %v = %d, %d, want %d, %d", tt.nums, sum, count, tt.wantSum, tt.wantCount)
		}
	}
}

func TestStartWorker(t *testing.T) {
	jobs := make(chan int, 3)
	handled := make(chan int, 10)
	var done <-chan bool
	finishes(t, "StartWorker(jobs, handle)", func() { done = StartWorker(jobs, func(j int) { handled <- j }) })
	if done == nil {
		t.Fatal("StartWorker(jobs, handle) = nil, want a channel made with make")
	}

	var got []int
	for j := 1; j <= 3; j++ {
		jobs <- j
		v, _ := recv(t, handled)
		got = append(got, v)
	}
	if want := []int{1, 2, 3}; !slices.Equal(got, want) {
		t.Errorf("StartWorker: handle was called with the jobs %v, want %v", got, want)
	}

	select {
	case <-done:
		t.Fatal("StartWorker: the done channel fired although jobs is not closed yet")
	case <-time.After(20 * time.Millisecond):
	}

	close(jobs)
	if v, ok := recv(t, done); !v || !ok {
		t.Errorf("StartWorker: after closing jobs the done channel delivered %v (open: %v), want the value true", v, ok)
	}
}

func TestSquare(t *testing.T) {
	// Input that is complete and closed from the start.
	in := make(chan int, 4)
	for _, n := range []int{1, 2, 0, -3} {
		in <- n
	}
	close(in)

	var out <-chan int
	finishes(t, "Square(in)", func() { out = Square(in) })
	var got []int
	for {
		v, ok := recv(t, out)
		if !ok {
			break
		}
		got = append(got, v)
		if len(got) > 10 {
			t.Fatalf("Square delivered more than 10 values for 4 numbers: %v", got)
		}
	}
	if want := []int{1, 4, 0, 9}; !slices.Equal(got, want) {
		t.Errorf("Square of a closed channel holding 1, 2, 0, -3 delivered %v before closing, want %v", got, want)
	}

	// Input that stays open for a while.
	live := make(chan int)
	finishes(t, "Square(in)", func() { out = Square(live) })
	select {
	case live <- 5:
	case <-time.After(2 * time.Second):
		t.Fatal("Square did not receive the number sent on its input within 2 seconds")
	}
	if v, ok := recv(t, out); v != 25 || !ok {
		t.Errorf("Square: after sending 5 the output delivered %d (open: %v), want 25", v, ok)
	}
	select {
	case v, ok := <-out:
		t.Fatalf("Square: the output delivered %d (open: %v) although the input is idle and still open", v, ok)
	case <-time.After(20 * time.Millisecond):
	}
	close(live)
	if v, ok := recv(t, out); ok {
		t.Errorf("Square: after closing the input the output delivered %d, want it to be closed", v)
	}
}
