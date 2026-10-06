package rangeoveriterators

import (
	"iter"
	"slices"
	"testing"
)

// collect gathers all values of seq. It fails the test instead of crashing
// on a nil iterator, and gives up on an iterator that does not seem to end.
func collect[T any](t *testing.T, what string, seq iter.Seq[T]) []T {
	t.Helper()
	if seq == nil {
		t.Fatalf("%s returned a nil iterator", what)
	}
	const limit = 10_000
	var got []T
	for v := range seq {
		got = append(got, v)
		if len(got) > limit {
			t.Fatalf("%s yielded more than %d values - is a stop condition missing?", what, limit)
		}
	}
	return got
}

// firstN gathers the first n values of seq and then breaks out of the loop.
func firstN[T any](t *testing.T, what string, seq iter.Seq[T], n int) []T {
	t.Helper()
	if seq == nil {
		t.Fatalf("%s returned a nil iterator", what)
	}
	var got []T
	for v := range seq {
		got = append(got, v)
		if len(got) == n {
			break
		}
	}
	return got
}

// requireStopsEarly calls the iterator by hand with a yield function that
// asks it to stop right away, and checks that it does.
func requireStopsEarly[T any](t *testing.T, what string, seq iter.Seq[T]) {
	t.Helper()
	if seq == nil {
		t.Fatalf("%s returned a nil iterator", what)
	}
	calls := 0
	func() {
		defer func() { recover() }() // only for the runaway panic below
		seq(func(T) bool {
			calls++
			if calls > 100 {
				panic("runaway iterator")
			}
			return false
		})
	}()
	if calls != 1 {
		t.Fatalf("%s called yield %d times although the first call returned false - check the result of yield and return when it is false", what, calls)
	}
}

func TestCountdown(t *testing.T) {
	tests := []struct {
		n    int
		want []int
	}{
		{3, []int{3, 2, 1}},
		{1, []int{1}},
		{0, nil},
		{-2, nil},
	}
	for _, tt := range tests {
		if got := collect(t, "Countdown", Countdown(tt.n)); !slices.Equal(got, tt.want) {
			t.Errorf("Countdown(%d) yielded %v, want %v", tt.n, got, tt.want)
		}
	}
	requireStopsEarly(t, "Countdown(5)", Countdown(5))
}

func TestListAll(t *testing.T) {
	var lst List[int]
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	if got, want := collect(t, "List.All", lst.All()), []int{10, 13, 23}; !slices.Equal(got, want) {
		t.Errorf("list of 10, 13, 23: All() yielded %v, want %v", got, want)
	}
	// An iterator can be consumed again; it starts from the beginning.
	if got, want := collect(t, "List.All", lst.All()), []int{10, 13, 23}; !slices.Equal(got, want) {
		t.Errorf("second pass over the list: All() yielded %v, want %v", got, want)
	}
	requireStopsEarly(t, "List.All", lst.All())

	var names List[string]
	if got := collect(t, "List.All", names.All()); len(got) != 0 {
		t.Errorf("empty list: All() yielded %q, want nothing", got)
	}
	names.Push("ann")
	if got, want := collect(t, "List.All", names.All()), []string{"ann"}; !slices.Equal(got, want) {
		t.Errorf(`list of "ann": All() yielded %q, want %q`, got, want)
	}
}

func TestFib(t *testing.T) {
	// Checked first: breaking out of a range loop over an iterator that
	// ignores the result of yield makes the program panic.
	requireStopsEarly(t, "Fib()", Fib())

	want := []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	if got := firstN(t, "Fib", Fib(), len(want)); !slices.Equal(got, want) {
		t.Errorf("the first %d values of Fib() are %v, want %v", len(want), got, want)
	}
}

func TestTake(t *testing.T) {
	requireStopsEarly(t, "Fib()", Fib()) // Take(Fib(), n) only ends if Fib stops when asked to

	if got, want := collect(t, "Take", Take(Fib(), 8)), []int{1, 1, 2, 3, 5, 8, 13, 21}; !slices.Equal(got, want) {
		t.Errorf("Take(Fib(), 8) yielded %v, want %v", got, want)
	}
	if got, want := collect(t, "Take", Take(Countdown(3), 10)), []int{3, 2, 1}; !slices.Equal(got, want) {
		t.Errorf("Take(Countdown(3), 10) yielded %v, want %v", got, want)
	}
	if got := collect(t, "Take", Take(Fib(), 0)); len(got) != 0 {
		t.Errorf("Take(Fib(), 0) yielded %v, want nothing", got)
	}
	requireStopsEarly(t, "Take(Fib(), 5)", Take(Fib(), 5))

	// Take must not pull more values out of seq than it passes on.
	pulled := 0
	counting := func(yield func(int) bool) {
		for i := 1; i <= 10; i++ {
			pulled++
			if !yield(i) {
				return
			}
		}
	}
	collect(t, "Take", Take(counting, 3))
	if pulled != 3 {
		t.Errorf("Take(seq, 3) pulled %d values out of seq, want 3 - stop right after the third value", pulled)
	}
}

func TestEnumerate(t *testing.T) {
	var lst List[string]
	lst.Push("a")
	lst.Push("b")
	lst.Push("c")

	seq := Enumerate(lst.All())
	if seq == nil {
		t.Fatal("Enumerate returned a nil iterator")
	}
	type pair struct {
		i int
		v string
	}
	var got []pair
	for i, v := range seq {
		got = append(got, pair{i, v})
	}
	want := []pair{{0, "a"}, {1, "b"}, {2, "c"}}
	if !slices.Equal(got, want) {
		t.Errorf(`Enumerate over "a", "b", "c" yielded %v, want %v`, got, want)
	}

	calls := 0
	seq(func(int, string) bool {
		calls++
		return false
	})
	if calls != 1 {
		t.Errorf("Enumerate called yield %d times although the first call returned false", calls)
	}
}

func TestEvenFibs(t *testing.T) {
	tests := []struct {
		n    int
		want []int
	}{
		{1, []int{2}},
		{4, []int{2, 8, 34, 144}},
		{7, []int{2, 8, 34, 144, 610, 2584, 10946}},
		{0, nil},
	}
	for _, tt := range tests {
		if got := EvenFibs(tt.n); !slices.Equal(got, tt.want) {
			t.Errorf("EvenFibs(%d) = %v, want %v", tt.n, got, tt.want)
		}
	}
}
