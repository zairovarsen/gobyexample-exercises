package panics

import (
	"errors"
	"runtime"
	"strconv"
	"testing"
)

// catch calls f and reports whether it panicked and with which value. It uses
// recover, which is the topic of the next section - no need to understand it
// yet.
func catch(f func()) (value any, panicked bool) {
	defer func() {
		if r := recover(); r != nil {
			value, panicked = r, true
		}
	}()
	f()
	return nil, false
}

// mustPanic calls f and returns the value it panicked with. The test fails
// when f returns normally.
func mustPanic(t *testing.T, call string, f func()) any {
	t.Helper()
	value, panicked := catch(f)
	if !panicked {
		t.Fatalf("%s returned normally, want a panic", call)
	}
	return value
}

// mustNotPanic calls f and fails the test when it panics.
func mustNotPanic(t *testing.T, call string, f func()) {
	t.Helper()
	if value, panicked := catch(f); panicked {
		t.Fatalf("%s panicked with %v, want a normal return", call, value)
	}
}

func TestFail(t *testing.T) {
	value := mustPanic(t, "Fail()", Fail)
	msg, ok := value.(string)
	if !ok {
		t.Fatalf("Fail() panicked with a %T (%v), want a string", value, value)
	}
	if want := "a problem"; msg != want {
		t.Errorf("Fail() panicked with %q, want %q", msg, want)
	}
}

func TestMustPositive(t *testing.T) {
	for _, n := range []int{1, 7, 1000} {
		var got int
		mustNotPanic(t, "MustPositive("+strconv.Itoa(n)+")", func() { got = MustPositive(n) })
		if got != n {
			t.Errorf("MustPositive(%d) = %d, want %d", n, got, n)
		}
	}

	panics := []struct {
		n    int
		want string
	}{
		{0, "not positive: 0"},
		{-3, "not positive: -3"},
	}
	for _, tt := range panics {
		value := mustPanic(t, "MustPositive("+strconv.Itoa(tt.n)+")", func() { MustPositive(tt.n) })
		msg, ok := value.(string)
		if !ok {
			t.Errorf("MustPositive(%d) panicked with a %T (%v), want a string", tt.n, value, value)
			continue
		}
		if msg != tt.want {
			t.Errorf("MustPositive(%d) panicked with %q, want %q", tt.n, msg, tt.want)
		}
	}
}

func TestElement(t *testing.T) {
	nums := []int{10, 20, 30}
	for i, want := range nums {
		var got int
		mustNotPanic(t, "Element(nums, "+strconv.Itoa(i)+")", func() { got = Element(nums, i) })
		if got != want {
			t.Errorf("Element(%v, %d) = %d, want %d", nums, i, got, want)
		}
	}

	for _, i := range []int{3, -1, 100} {
		value := mustPanic(t, "Element(nums, "+strconv.Itoa(i)+")", func() { Element(nums, i) })
		if _, ok := value.(runtime.Error); !ok {
			t.Errorf("Element(%v, %d) panicked with a %T (%v), want the runtime's own index-out-of-range error",
				nums, i, value, value)
		}
	}
}

func TestMust(t *testing.T) {
	var n int
	mustNotPanic(t, "Must(42, nil)", func() { n = Must(42, nil) })
	if n != 42 {
		t.Errorf("Must(42, nil) = %d, want 42", n)
	}

	var s string
	mustNotPanic(t, `Must("ok", nil)`, func() { s = Must("ok", nil) })
	if s != "ok" {
		t.Errorf(`Must("ok", nil) = %q, want "ok"`, s)
	}

	errBroken := errors.New("broken")
	value := mustPanic(t, "Must(0, errBroken)", func() { Must(0, errBroken) })
	err, ok := value.(error)
	if !ok {
		t.Fatalf("Must(0, errBroken) panicked with a %T (%v), want the error itself", value, value)
	}
	if !errors.Is(err, errBroken) {
		t.Errorf("Must(0, errBroken) panicked with the error %q, want errBroken", err)
	}

	// The typical use: wrapping a call that returns (value, error).
	value = mustPanic(t, `Must(strconv.Atoi("x"))`, func() { Must(strconv.Atoi("x")) })
	if err, ok := value.(error); !ok || !errors.Is(err, strconv.ErrSyntax) {
		t.Errorf(`Must(strconv.Atoi("x")) panicked with %v (%T), want the error returned by Atoi`, value, value)
	}
}

func TestStack(t *testing.T) {
	var s Stack
	s.Push(1)
	s.Push(2)
	s.Push(3)
	if got := s.Len(); got != 3 {
		t.Fatalf("after 3 pushes Len() = %d, want 3", got)
	}
	for _, want := range []int{3, 2} {
		var got int
		mustNotPanic(t, "Pop()", func() { got = s.Pop() })
		if got != want {
			t.Errorf("Pop() = %d, want %d", got, want)
		}
	}
	s.Push(9)
	for _, want := range []int{9, 1} {
		var got int
		mustNotPanic(t, "Pop()", func() { got = s.Pop() })
		if got != want {
			t.Errorf("Pop() = %d, want %d", got, want)
		}
	}
	if got := s.Len(); got != 0 {
		t.Errorf("after popping everything Len() = %d, want 0", got)
	}
}

func TestStackPopEmpty(t *testing.T) {
	var s Stack
	value := mustPanic(t, "Pop() on an empty stack", func() { s.Pop() })
	err, ok := value.(error)
	if !ok {
		t.Fatalf("Pop() on an empty stack panicked with a %T (%v), want ErrEmptyStack", value, value)
	}
	if !errors.Is(err, ErrEmptyStack) {
		t.Errorf("Pop() on an empty stack panicked with the error %q, want ErrEmptyStack", err)
	}

	// The same after the stack was used and emptied again.
	s.Push(5)
	mustNotPanic(t, "Pop() after Push(5)", func() { s.Pop() })
	mustPanic(t, "Pop() on an emptied stack", func() { s.Pop() })
}
