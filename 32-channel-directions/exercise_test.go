package channeldirections

import (
	"slices"
	"testing"
	"time"
)

// The channel directions in the signatures are part of the exercise. These
// assignments only compile as long as nobody loosens them.
var (
	_ func(chan<- string, string)        = Ping
	_ func(<-chan string, chan<- string) = Pong
	_ func(...int) <-chan int            = Generate
	_ func(<-chan int, chan<- int, int)  = Double
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

func TestPing(t *testing.T) {
	pings := make(chan string, 1)
	finishes(t, `Ping(pings, "passed message")`, func() { Ping(pings, "passed message") })

	select {
	case got := <-pings:
		if got != "passed message" {
			t.Errorf(`Ping(pings, "passed message") sent %q, want %q`, got, "passed message")
		}
	default:
		t.Error(`Ping(pings, "passed message") returned without sending on pings`)
	}
}

func TestPong(t *testing.T) {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	pings <- "passed message"
	finishes(t, "Pong(pings, pongs)", func() { Pong(pings, pongs) })

	if len(pings) != 0 {
		t.Error("Pong(pings, pongs) returned without receiving the message waiting on pings")
	}
	select {
	case got := <-pongs:
		if got != "passed message" {
			t.Errorf("Pong(pings, pongs) sent %q on pongs, want %q", got, "passed message")
		}
	default:
		t.Error("Pong(pings, pongs) returned without sending on pongs")
	}
}

func TestGenerate(t *testing.T) {
	nums := []int{3, 1, 4, 1, 5}
	var ch <-chan int
	finishes(t, "Generate(3, 1, 4, 1, 5) (it must not wait for a receiver)", func() { ch = Generate(nums...) })
	if ch == nil {
		t.Fatal("Generate(3, 1, 4, 1, 5) = nil, want a channel made with make")
	}
	var got []int
	for range nums {
		got = append(got, recv(t, ch))
	}
	if !slices.Equal(got, nums) {
		t.Errorf("Generate(3, 1, 4, 1, 5) delivered %v, want %v", got, nums)
	}
}

func TestDouble(t *testing.T) {
	in := make(chan int, 3)
	out := make(chan int, 3)
	in <- 1
	in <- 20
	in <- -3
	finishes(t, "Double(in, out, 3)", func() { Double(in, out, 3) })

	if len(out) != 3 {
		t.Fatalf("Double(in, out, 3) returned after sending %d values on out, want 3", len(out))
	}
	got := []int{<-out, <-out, <-out}
	if want := []int{2, 40, -6}; !slices.Equal(got, want) {
		t.Errorf("Double(in, out, 3) turned 1, 20, -3 into %v, want %v", got, want)
	}

	// With n = 0 there is nothing to do.
	finishes(t, "Double(in, out, 0)", func() { Double(in, out, 0) })
}

func TestPipeline(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 3}, []int{2, 4, 6}},
		{[]int{-7, 0, 50, 2, 2}, []int{-14, 0, 100, 4, 4}},
		{[]int{21}, []int{42}},
	}
	for _, tt := range tests {
		var got []int
		finishes(t, "Pipeline", func() { got = Pipeline(tt.nums) })
		if !slices.Equal(got, tt.want) {
			t.Errorf("Pipeline(%v) = %v, want %v", tt.nums, got, tt.want)
		}
	}

	var got []int
	finishes(t, "Pipeline(nil)", func() { got = Pipeline(nil) })
	if len(got) != 0 {
		t.Errorf("Pipeline(nil) = %v, want an empty result", got)
	}
}
