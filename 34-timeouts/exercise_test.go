package timeouts

import (
	"errors"
	"slices"
	"testing"
	"time"
)

// The tests use short durations to stay fast. A call that is supposed to
// time out after short must take at least that long, and far less than
// 10 times as long. Where no timeout is expected the limit is long, which
// is never reached unless something is wrong.
const (
	short = 50 * time.Millisecond
	long  = 2 * time.Second
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
	case <-time.After(3 * time.Second):
		t.Fatalf("%s did not return within 3 seconds", what)
	}
}

// requireTimedOut checks the duration of a call that should have run into
// its timeout of short.
func requireTimedOut(t *testing.T, what string, elapsed time.Duration) {
	t.Helper()
	if elapsed < short {
		t.Errorf("%s returned after %v, want it to wait for the whole timeout of %v", what, elapsed, short)
	}
	if elapsed > 10*short {
		t.Errorf("%s returned after %v, want it to give up soon after the timeout of %v", what, elapsed, short)
	}
}

// requirePrompt checks the duration of a call that should have returned
// long before its timeout of long.
func requirePrompt(t *testing.T, what string, elapsed time.Duration) {
	t.Helper()
	if elapsed > long/2 {
		t.Errorf("%s returned after %v, want it to return as soon as the channel operation is possible", what, elapsed)
	}
}

func TestRecvTimeout(t *testing.T) {
	// A message is waiting already.
	ready := make(chan string, 1)
	ready <- "result 1"
	var msg string
	var ok bool
	start := time.Now()
	finishes(t, "RecvTimeout", func() { msg, ok = RecvTimeout(ready, long) })
	if msg != "result 1" || !ok {
		t.Errorf("RecvTimeout(ch, %v) with a message waiting = %q, %v, want %q, true", long, msg, ok, "result 1")
	}
	requirePrompt(t, "RecvTimeout with a message waiting", time.Since(start))

	// A message arrives a little later, but in time.
	late := make(chan string, 1)
	go func() {
		time.Sleep(20 * time.Millisecond)
		late <- "result 2"
	}()
	start = time.Now()
	finishes(t, "RecvTimeout", func() { msg, ok = RecvTimeout(late, long) })
	if msg != "result 2" || !ok {
		t.Errorf("RecvTimeout(ch, %v) with a message arriving after 20ms = %q, %v, want %q, true", long, msg, ok, "result 2")
	}
	requirePrompt(t, "RecvTimeout with a message arriving after 20ms", time.Since(start))

	// Nothing arrives.
	silent := make(chan string)
	start = time.Now()
	finishes(t, "RecvTimeout", func() { msg, ok = RecvTimeout(silent, short) })
	if msg != "" || ok {
		t.Errorf("RecvTimeout(ch, %v) without any message = %q, %v, want %q, false", short, msg, ok, "")
	}
	requireTimedOut(t, "RecvTimeout without any message", time.Since(start))
}

func TestSendTimeout(t *testing.T) {
	// There is room in the buffer.
	free := make(chan string, 1)
	var sent bool
	start := time.Now()
	finishes(t, "SendTimeout", func() { sent = SendTimeout(free, "hi", long) })
	if !sent {
		t.Errorf("SendTimeout(ch, %q, %v) on a channel with a free buffer slot = false, want true", "hi", long)
	}
	if len(free) != 1 || <-free != "hi" {
		t.Errorf("SendTimeout(ch, %q, %v) on a channel with a free buffer slot did not send %q", "hi", long, "hi")
	}
	requirePrompt(t, "SendTimeout on a channel with a free buffer slot", time.Since(start))

	// A receiver shows up a little later, but in time.
	slow := make(chan string)
	received := make(chan string, 1)
	go func() {
		time.Sleep(20 * time.Millisecond)
		select {
		case msg := <-slow:
			received <- msg
		case <-time.After(long):
			received <- ""
		}
	}()
	start = time.Now()
	finishes(t, "SendTimeout", func() { sent = SendTimeout(slow, "ho", long) })
	if !sent {
		t.Errorf("SendTimeout(ch, %q, %v) with a receiver arriving after 20ms = false, want true", "ho", long)
	} else if got := <-received; got != "ho" {
		t.Errorf("SendTimeout(ch, %q, %v): the receiver got %q, want %q", "ho", long, got, "ho")
	}
	requirePrompt(t, "SendTimeout with a receiver arriving after 20ms", time.Since(start))

	// Nobody ever receives.
	blocked := make(chan string)
	start = time.Now()
	finishes(t, "SendTimeout", func() { sent = SendTimeout(blocked, "hey", short) })
	if sent {
		t.Errorf("SendTimeout(ch, %q, %v) without a receiver = true, want false", "hey", short)
	}
	requireTimedOut(t, "SendTimeout without a receiver", time.Since(start))
}

func TestWithTimeout(t *testing.T) {
	// f is quick.
	var res string
	var err error
	start := time.Now()
	finishes(t, "WithTimeout", func() { res, err = WithTimeout(func() string { return "result 1" }, long) })
	if res != "result 1" || err != nil {
		t.Errorf("WithTimeout(f, %v) with a quick f = %q, %v, want %q, <nil>", long, res, err, "result 1")
	}
	requirePrompt(t, "WithTimeout with a quick f", time.Since(start))

	// f takes far too long: it only returns when the test is over.
	release := make(chan struct{})
	defer close(release)
	stuck := func() string {
		<-release
		return "result 2"
	}
	start = time.Now()
	finishes(t, "WithTimeout(f, timeout) with an f that does not return (run f in a goroutine)", func() { res, err = WithTimeout(stuck, short) })
	if res != "" || !errors.Is(err, ErrTimeout) {
		t.Errorf("WithTimeout(f, %v) with an f that does not return = %q, %v, want %q, ErrTimeout", short, res, err, "")
	}
	requireTimedOut(t, "WithTimeout with an f that does not return", time.Since(start))
}

func TestCollectFor(t *testing.T) {
	// Three numbers are waiting, and then nothing comes any more.
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	var got []int
	start := time.Now()
	finishes(t, "CollectFor", func() { got = CollectFor(ch, short) })
	if want := []int{1, 2, 3}; !slices.Equal(got, want) {
		t.Errorf("CollectFor(ch, %v) with 1, 2, 3 waiting = %v, want %v", short, got, want)
	}
	requireTimedOut(t, "CollectFor with 3 numbers waiting", time.Since(start))

	// A steady stream: a new number every 5ms, until the test is over.
	stream := make(chan int)
	stop := make(chan struct{})
	defer close(stop)
	go func() {
		for i := 0; ; i++ {
			select {
			case stream <- i:
				time.Sleep(5 * time.Millisecond)
			case <-stop:
				return
			}
		}
	}()
	start = time.Now()
	finishes(t, "CollectFor(ch, timeout) with a new number every 5ms (the deadline must not start over with every number)", func() {
		got = CollectFor(stream, short)
	})
	requireTimedOut(t, "CollectFor with a new number every 5ms", time.Since(start))
	if len(got) < 2 {
		t.Errorf("CollectFor(ch, %v) with a new number every 5ms = %v, want several numbers", short, got)
	}
	for i, n := range got {
		if n != i {
			t.Errorf("CollectFor(ch, %v) with a new number every 5ms = %v, want 0, 1, 2, ... without gaps", short, got)
			break
		}
	}
}
