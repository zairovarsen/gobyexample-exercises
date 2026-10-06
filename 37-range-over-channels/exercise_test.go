package rangeoverchannels

import (
	"slices"
	"testing"
	"time"
)

// drain receives from ch until it is closed and returns what it got. It
// fails the test instead of hanging forever when the channel is nil or is
// never closed.
func drain[T any](t *testing.T, what string, ch <-chan T) []T {
	t.Helper()
	if ch == nil {
		t.Fatalf("%s returned a nil channel - did you create it with make?", what)
	}
	var got []T
	timeout := time.After(2 * time.Second)
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				return got
			}
			got = append(got, v)
		case <-timeout:
			t.Fatalf("the channel returned by %s was not closed within 2 seconds (values received so far: %v)", what, got)
		}
	}
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

// closed returns a closed channel that still holds vals.
func closed[T any](vals ...T) <-chan T {
	ch := make(chan T, len(vals))
	for _, v := range vals {
		ch <- v
	}
	close(ch)
	return ch
}

func TestCollect(t *testing.T) {
	var got []string
	finishes(t, "Collect", func() { got = Collect(closed("one", "two", "three")) })
	if want := []string{"one", "two", "three"}; !slices.Equal(got, want) {
		t.Errorf(`Collect of a closed channel holding "one", "two", "three" = %q, want %q`, got, want)
	}

	finishes(t, "Collect", func() { got = Collect(closed[string]()) })
	if len(got) != 0 {
		t.Errorf("Collect of a closed, empty channel = %q, want an empty result", got)
	}

	// The loop has to keep going until the channel is closed, not just
	// until it is empty for a moment.
	slow := make(chan string)
	go func() {
		for _, s := range []string{"a", "b"} {
			time.Sleep(10 * time.Millisecond)
			slow <- s
		}
		close(slow)
	}()
	finishes(t, "Collect", func() { got = Collect(slow) })
	if want := []string{"a", "b"}; !slices.Equal(got, want) {
		t.Errorf(`Collect of a channel that slowly delivers "a", "b" and is then closed = %q, want %q`, got, want)
	}
}

func TestCount(t *testing.T) {
	tests := []struct {
		n    int
		want []int
	}{
		{5, []int{1, 2, 3, 4, 5}},
		{1, []int{1}},
		{0, nil},
		{-3, nil},
	}
	for _, tt := range tests {
		var ch <-chan int
		finishes(t, "Count (it must not wait for a receiver)", func() { ch = Count(tt.n) })
		if got := drain(t, "Count", ch); !slices.Equal(got, tt.want) {
			t.Errorf("Count(%d) delivered %v before closing, want %v", tt.n, got, tt.want)
		}
	}
}

func TestEvens(t *testing.T) {
	tests := []struct {
		in   []int
		want []int
	}{
		{[]int{1, 2, 3, 4}, []int{2, 4}},
		{[]int{10, 0, -2, 7, 8}, []int{10, 0, -2, 8}},
		{[]int{1, 3, 5}, nil},
		{nil, nil},
	}
	for _, tt := range tests {
		var ch <-chan int
		finishes(t, "Evens (it must not wait for a receiver)", func() { ch = Evens(closed(tt.in...)) })
		if got := drain(t, "Evens", ch); !slices.Equal(got, tt.want) {
			t.Errorf("Evens of a channel delivering %v delivered %v before closing, want %v", tt.in, got, tt.want)
		}
	}
}

func TestSumOfEvens(t *testing.T) {
	tests := []struct {
		n, want int
	}{
		{10, 30},
		{7, 12},
		{2, 2},
		{1, 0},
		{0, 0},
	}
	for _, tt := range tests {
		var got int
		finishes(t, "SumOfEvens", func() { got = SumOfEvens(tt.n) })
		if got != tt.want {
			t.Errorf("SumOfEvens(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}

func TestMerge(t *testing.T) {
	tests := []struct {
		a, b []int
		want []int // sorted
	}{
		{[]int{1, 3, 5}, []int{2, 4, 6}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{7, 7}, nil, []int{7, 7}},
		{nil, []int{9}, []int{9}},
		{nil, nil, nil},
	}
	for _, tt := range tests {
		var ch <-chan int
		finishes(t, "Merge (it must not wait for a receiver)", func() { ch = Merge(closed(tt.a...), closed(tt.b...)) })
		got := drain(t, "Merge", ch)
		slices.Sort(got)
		if !slices.Equal(got, tt.want) {
			t.Errorf("Merge of channels delivering %v and %v delivered (sorted) %v before closing, want %v", tt.a, tt.b, got, tt.want)
		}
	}

	// One input stays open longer than the other: the output may only be
	// closed when both are done.
	a := make(chan int)
	b := closed(1)
	var ch <-chan int
	finishes(t, "Merge (it must not wait for a receiver)", func() { ch = Merge(a, b) })
	go func() {
		time.Sleep(30 * time.Millisecond)
		a <- 2
		close(a)
	}()
	got := drain(t, "Merge", ch)
	slices.Sort(got)
	if want := []int{1, 2}; !slices.Equal(got, want) {
		t.Errorf("Merge with one input that delivers late delivered (sorted) %v before closing, want %v", got, want)
	}
}
