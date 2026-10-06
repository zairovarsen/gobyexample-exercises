package channels

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

func TestPing(t *testing.T) {
	ch := Ping("ping")
	if got := recv(t, ch); got != "ping" {
		t.Errorf("received %q, want %q", got, "ping")
	}
	if cap(ch) != 0 {
		t.Errorf("channel has capacity %d, want an unbuffered channel", cap(ch))
	}
}

func TestSumAsync(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3, 4}, 10},
		{[]int{-5, 5}, 0},
		{nil, 0},
	}
	for _, tt := range tests {
		if got := recv(t, SumAsync(tt.nums)); got != tt.want {
			t.Errorf("SumAsync(%v) delivered %d, want %d", tt.nums, got, tt.want)
		}
	}
}

func TestRelay(t *testing.T) {
	in, out := make(chan string), make(chan string)
	go Relay(in, out)

	select {
	case in <- "hi":
	case <-time.After(2 * time.Second):
		t.Fatal("Relay never received from the in channel")
	}
	if got, want := recv(t, out), "relayed: hi"; got != want {
		t.Errorf("Relay sent %q, want %q", got, want)
	}
}

func TestParallelSquares(t *testing.T) {
	nums := make([]int, 100)
	want := make([]int, 100)
	for i := range nums {
		nums[i] = i + 1
		want[i] = (i + 1) * (i + 1)
	}

	done := make(chan []int, 1)
	go func() { done <- ParallelSquares(nums) }()

	if got := recv(t, done); !slices.Equal(got, want) {
		t.Errorf("ParallelSquares(1..100) = %v, want the squares in input order", got)
	}

	if got := ParallelSquares(nil); len(got) != 0 {
		t.Errorf("ParallelSquares(nil) = %v, want an empty result", got)
	}
}
