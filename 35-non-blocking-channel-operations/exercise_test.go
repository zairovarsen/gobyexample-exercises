package nonblockingchanneloperations

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
		t.Fatalf("%s blocked (it did not return within 2 seconds) - use a select with a default case", what)
	}
}

func TestTryReceive(t *testing.T) {
	var msg string
	var ok bool

	ready := make(chan string, 1)
	ready <- "hi"
	finishes(t, "TryReceive", func() { msg, ok = TryReceive(ready) })
	if msg != "hi" || !ok {
		t.Errorf("TryReceive(ch) with %q waiting = %q, %v, want %q, true", "hi", msg, ok, "hi")
	}
	if len(ready) != 0 {
		t.Error("TryReceive(ch) left the waiting message in the channel")
	}

	finishes(t, "TryReceive on an empty buffered channel", func() { msg, ok = TryReceive(ready) })
	if msg != "" || ok {
		t.Errorf("TryReceive(ch) on an empty buffered channel = %q, %v, want %q, false", msg, ok, "")
	}

	finishes(t, "TryReceive on an unbuffered channel without a sender", func() { msg, ok = TryReceive(make(chan string)) })
	if msg != "" || ok {
		t.Errorf("TryReceive(ch) on an unbuffered channel without a sender = %q, %v, want %q, false", msg, ok, "")
	}
}

func TestTrySend(t *testing.T) {
	var sent bool

	buffered := make(chan string, 1)
	finishes(t, "TrySend", func() { sent = TrySend(buffered, "hi") })
	if !sent {
		t.Errorf("TrySend(ch, %q) on a channel with a free buffer slot = false, want true", "hi")
	}
	if len(buffered) != 1 {
		t.Fatalf("TrySend(ch, %q) on a channel with a free buffer slot did not send the message", "hi")
	}

	finishes(t, "TrySend on a full channel", func() { sent = TrySend(buffered, "ho") })
	if sent {
		t.Errorf("TrySend(ch, %q) on a full channel = true, want false", "ho")
	}
	if got := <-buffered; got != "hi" {
		t.Errorf("after the two TrySend calls the channel holds %q, want %q", got, "hi")
	}

	finishes(t, "TrySend on an unbuffered channel without a receiver", func() { sent = TrySend(make(chan string), "hey") })
	if sent {
		t.Errorf("TrySend(ch, %q) on an unbuffered channel without a receiver = true, want false", "hey")
	}
}

func TestPoll(t *testing.T) {
	tests := []struct {
		name    string
		message string // sent on messages unless empty
		signal  *bool  // sent on signals unless nil
		want    string
	}{
		{"a message waiting", "hi", nil, "received message hi"},
		{"the signal true waiting", "", new(true), "received signal true"},
		{"the signal false waiting", "", new(false), "received signal false"},
		{"nothing waiting", "", nil, "no activity"},
	}
	for _, tt := range tests {
		messages := make(chan string, 1)
		signals := make(chan bool, 1)
		if tt.message != "" {
			messages <- tt.message
		}
		if tt.signal != nil {
			signals <- *tt.signal
		}

		var got string
		finishes(t, "Poll with "+tt.name, func() { got = Poll(messages, signals) })
		if got != tt.want {
			t.Errorf("Poll(messages, signals) with %s = %q, want %q", tt.name, got, tt.want)
		}
	}
}

func TestDrainAvailable(t *testing.T) {
	var got []int

	ch := make(chan int, 5)
	ch <- 3
	ch <- 1
	ch <- 2
	finishes(t, "DrainAvailable", func() { got = DrainAvailable(ch) })
	if want := []int{3, 1, 2}; !slices.Equal(got, want) {
		t.Errorf("DrainAvailable(ch) with 3, 1, 2 waiting = %v, want %v", got, want)
	}
	if len(ch) != 0 {
		t.Errorf("DrainAvailable(ch) left %d values in the channel, want it empty", len(ch))
	}

	finishes(t, "DrainAvailable on an empty channel", func() { got = DrainAvailable(ch) })
	if len(got) != 0 {
		t.Errorf("DrainAvailable(ch) on an empty channel = %v, want an empty result", got)
	}
}

func TestBroadcast(t *testing.T) {
	ready1 := make(chan string, 1)
	full := make(chan string, 1)
	full <- "old news"
	nobodyListening := make(chan string)
	ready2 := make(chan string, 2)
	subscribers := []chan string{ready1, full, nobodyListening, ready2}

	var got int
	finishes(t, "Broadcast", func() { got = Broadcast(subscribers, "news") })
	if got != 2 {
		t.Errorf("Broadcast(subscribers, %q) with 2 of 4 subscribers ready = %d, want 2", "news", got)
	}
	for i, ch := range []chan string{ready1, ready2} {
		if len(ch) != 1 {
			t.Errorf("Broadcast: ready subscriber %d has %d messages waiting, want 1", i+1, len(ch))
		} else if msg := <-ch; msg != "news" {
			t.Errorf("Broadcast: ready subscriber %d got %q, want %q", i+1, msg, "news")
		}
	}
	if msg := <-full; msg != "old news" {
		t.Errorf("Broadcast: the subscriber with the full buffer now holds %q, want its old message %q untouched", msg, "old news")
	}

	finishes(t, "Broadcast without subscribers", func() { got = Broadcast(nil, "news") })
	if got != 0 {
		t.Errorf("Broadcast(nil, %q) = %d, want 0", "news", got)
	}
}
