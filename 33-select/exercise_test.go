package selectx

import (
	"slices"
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
		t.Fatalf("%s did not return within 2 seconds - it is probably blocked on a channel operation", what)
	}
}

// sendOrFail sends v on ch for a function under test that runs in another
// goroutine. It fails the test when that function has already returned
// (returned is closed) or when the value is not taken in time.
func sendOrFail[T any](t *testing.T, what string, ch chan<- T, v T, returned <-chan struct{}) {
	t.Helper()
	select {
	case ch <- v:
	case <-returned:
		t.Fatalf("%s returned before it received the value %v", what, v)
	case <-time.After(2 * time.Second):
		t.Fatalf("%s did not receive the value %v within 2 seconds", what, v)
	}
}

func TestFirst(t *testing.T) {
	tests := []struct {
		ready    int // the channel that has a message
		wantMsg  string
		wantFrom int
	}{
		{1, "one", 1},
		{2, "two", 2},
	}
	for _, tt := range tests {
		c1 := make(chan string, 1)
		c2 := make(chan string, 1)
		if tt.ready == 1 {
			c1 <- "one"
		} else {
			c2 <- "two"
		}

		var msg string
		var from int
		finishes(t, "First(c1, c2)", func() { msg, from = First(c1, c2) })
		if msg != tt.wantMsg || from != tt.wantFrom {
			t.Errorf("First(c1, c2) with a message only on c%d = %q, %d, want %q, %d", tt.ready, msg, from, tt.wantMsg, tt.wantFrom)
		}
	}
}

func TestFastest(t *testing.T) {
	const fast, slow = 10 * time.Millisecond, 400 * time.Millisecond
	tests := []struct {
		d1, d2 time.Duration
		want   string
	}{
		{fast, slow, "one"},
		{slow, fast, "two"},
	}
	for _, tt := range tests {
		var got string
		start := time.Now()
		finishes(t, "Fastest", func() { got = Fastest(tt.d1, tt.d2) })
		elapsed := time.Since(start)

		if got != tt.want {
			t.Errorf("Fastest(%v, %v) = %q, want %q", tt.d1, tt.d2, got, tt.want)
		}
		if elapsed < fast {
			t.Errorf("Fastest(%v, %v) returned after %v, want it to wait %v for the faster goroutine", tt.d1, tt.d2, elapsed, fast)
		}
		if elapsed > slow/2 {
			t.Errorf("Fastest(%v, %v) returned after %v, want it to return as soon as the faster goroutine is done (about %v)", tt.d1, tt.d2, elapsed, fast)
		}
	}
}

func TestReceiveN(t *testing.T) {
	c1 := make(chan string)
	c2 := make(chan string)

	var got []string
	returned := make(chan struct{})
	go func() {
		defer close(returned)
		got = ReceiveN(c1, c2, 4)
	}()

	// The channels are unbuffered: every send completes only when ReceiveN
	// takes the value, which fixes the order of arrival.
	const what = "ReceiveN(c1, c2, 4)"
	sendOrFail(t, what, c2, "two", returned)
	sendOrFail(t, what, c1, "one", returned)
	sendOrFail(t, what, c1, "uno", returned)
	sendOrFail(t, what, c2, "dos", returned)

	select {
	case <-returned:
	case <-time.After(2 * time.Second):
		t.Fatalf("%s did not return within 2 seconds after receiving 4 messages", what)
	}
	if want := []string{"two", "one", "uno", "dos"}; !slices.Equal(got, want) {
		t.Errorf("%s = %q, want %q", what, got, want)
	}
}

func TestSendEither(t *testing.T) {
	// A send on an unbuffered channel without a receiver can never proceed;
	// a send on a channel with a free buffer slot can.
	blocked := make(chan string)
	free := make(chan string, 1)

	var got int
	finishes(t, `SendEither(blocked, free, "hi")`, func() { got = SendEither(blocked, free, "hi") })
	if got != 2 {
		t.Errorf(`SendEither(c1, c2, "hi") with only c2 ready = %d, want 2`, got)
	}
	if len(free) != 1 || <-free != "hi" {
		t.Error(`SendEither(c1, c2, "hi") with only c2 ready did not send "hi" on c2`)
	}

	finishes(t, `SendEither(free, blocked, "ho")`, func() { got = SendEither(free, blocked, "ho") })
	if got != 1 {
		t.Errorf(`SendEither(c1, c2, "ho") with only c1 ready = %d, want 1`, got)
	}
	if len(free) != 1 || <-free != "ho" {
		t.Error(`SendEither(c1, c2, "ho") with only c1 ready did not send "ho" on c1`)
	}

	// With both channels ready the message still goes out only once.
	a := make(chan string, 1)
	b := make(chan string, 1)
	finishes(t, `SendEither(a, b, "once")`, func() { got = SendEither(a, b, "once") })
	if len(a)+len(b) != 1 {
		t.Errorf(`SendEither(c1, c2, "once") with both channels ready sent the message %d times, want once`, len(a)+len(b))
	} else if (got == 1 && len(a) != 1) || (got == 2 && len(b) != 1) || (got != 1 && got != 2) {
		t.Errorf(`SendEither(c1, c2, "once") = %d, but the message is on c1: %v, on c2: %v`, got, len(a) == 1, len(b) == 1)
	}
}

func TestSumUntilQuit(t *testing.T) {
	nums := make(chan int)
	quit := make(chan bool)

	var got int
	returned := make(chan struct{})
	go func() {
		defer close(returned)
		got = SumUntilQuit(nums, quit)
	}()

	const what = "SumUntilQuit(nums, quit)"
	for _, n := range []int{1, 2, 3, 40} {
		sendOrFail(t, what, nums, n, returned)
	}
	sendOrFail(t, what, quit, true, returned)

	select {
	case <-returned:
	case <-time.After(2 * time.Second):
		t.Fatalf("%s did not return within 2 seconds after a value was sent on quit", what)
	}
	if got != 46 {
		t.Errorf("%s = %d after receiving 1, 2, 3, 40 and then quit, want 46", what, got)
	}
}
