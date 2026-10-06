package channelbuffering

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
	case <-time.After(2 * time.Second):
		t.Fatalf("%s did not return within 2 seconds - it is probably blocked on a channel operation", what)
	}
}

// buffered takes the values that are waiting in the buffer of ch, without
// ever blocking.
func buffered(ch chan int) []int {
	var vals []int
	for range len(ch) {
		select {
		case v := <-ch:
			vals = append(vals, v)
		default:
		}
	}
	return vals
}

func TestNewMailbox(t *testing.T) {
	for _, size := range []int{2, 5} {
		ch := NewMailbox(size)
		if ch == nil {
			t.Fatalf("NewMailbox(%d) = nil, want a channel made with make", size)
		}
		if cap(ch) != size {
			t.Errorf("NewMailbox(%d) has capacity %d, want %d", size, cap(ch), size)
		}
	}

	// Sends up to the capacity do not need a receiver.
	ch := NewMailbox(2)
	for _, msg := range []string{"buffered", "channel"} {
		select {
		case ch <- msg:
		default:
			t.Fatalf("sending %q on NewMailbox(2) would block, want room for 2 values", msg)
		}
	}
	if got := recv(t, ch); got != "buffered" {
		t.Errorf("first value received from the mailbox = %q, want %q", got, "buffered")
	}
}

func TestPreload(t *testing.T) {
	msgs := []string{"buffered", "channel", "of", "strings"}
	ch := Preload(msgs...)
	if ch == nil {
		t.Fatalf("Preload(%q) = nil, want a channel made with make", msgs)
	}
	if cap(ch) != len(msgs) {
		t.Errorf("Preload(%q) has capacity %d, want %d", msgs, cap(ch), len(msgs))
	}
	// All values must sit in the buffer already, not in goroutines that are
	// still waiting to send.
	if len(ch) != len(msgs) {
		t.Fatalf("Preload(%q) returned a channel holding %d values, want all %d in the buffer", msgs, len(ch), len(msgs))
	}
	for i, want := range msgs {
		if got := recv(t, ch); got != want {
			t.Errorf("Preload(%q): value %d received = %q, want %q", msgs, i, got, want)
		}
	}
}

func TestFillUp(t *testing.T) {
	ch := make(chan int, 5)
	ch <- 1
	ch <- 2

	var sent int
	finishes(t, "FillUp(ch, 9)", func() { sent = FillUp(ch, 9) })
	if sent != 3 {
		t.Errorf("FillUp(ch, 9) on a channel with 2 of 5 slots taken = %d, want 3", sent)
	}
	if got, want := buffered(ch), []int{1, 2, 9, 9, 9}; !slices.Equal(got, want) {
		t.Errorf("after FillUp(ch, 9) the channel holds %v, want %v", got, want)
	}

	full := make(chan int, 2)
	full <- 7
	full <- 7
	finishes(t, "FillUp(full, 9)", func() { sent = FillUp(full, 9) })
	if sent != 0 {
		t.Errorf("FillUp(full, 9) on a full channel = %d, want 0", sent)
	}
	if got, want := buffered(full), []int{7, 7}; !slices.Equal(got, want) {
		t.Errorf("after FillUp(full, 9) the full channel holds %v, want it unchanged: %v", got, want)
	}
}

func TestKeepLatest(t *testing.T) {
	ch := make(chan int, 3)
	for v := 1; v <= 5; v++ {
		finishes(t, "KeepLatest", func() { KeepLatest(ch, v) })
	}
	if got, want := buffered(ch), []int{3, 4, 5}; !slices.Equal(got, want) {
		t.Errorf("after KeepLatest of 1, 2, 3, 4, 5 on a channel with capacity 3 it holds %v, want %v", got, want)
	}

	single := make(chan int, 1)
	for _, v := range []int{10, 20} {
		finishes(t, "KeepLatest", func() { KeepLatest(single, v) })
	}
	if got, want := buffered(single), []int{20}; !slices.Equal(got, want) {
		t.Errorf("after KeepLatest of 10, 20 on a channel with capacity 1 it holds %v, want %v", got, want)
	}
}
