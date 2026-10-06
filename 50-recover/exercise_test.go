package recovery

import (
	"errors"
	"slices"
	"testing"
)

// guard runs f and turns a panic that escapes from it into a test failure
// (instead of a crash of the whole test binary).
func guard(t *testing.T, call string, f func()) {
	t.Helper()
	defer func() {
		t.Helper()
		if r := recover(); r != nil {
			t.Fatalf("%s let the panic \"%v\" escape, want it recovered", call, r)
		}
	}()
	f()
}

func TestDidPanic(t *testing.T) {
	var nums []int
	tests := []struct {
		name string
		f    func()
		want bool
	}{
		{"a function that panics with a string", func() { panic("a problem") }, true},
		{"a function that panics with an error", func() { panic(errors.New("bad")) }, true},
		{"a function that indexes out of range", func() { _ = nums[3] }, true},
		{"a function that returns normally", func() {}, false},
	}
	for _, tt := range tests {
		guard(t, "DidPanic", func() {
			if got := DidPanic(tt.f); got != tt.want {
				t.Errorf("DidPanic(%s) = %t, want %t", tt.name, got, tt.want)
			}
		})
	}
}

func TestDidPanicCallsF(t *testing.T) {
	called := false
	DidPanic(func() { called = true })
	if !called {
		t.Errorf("DidPanic(f) never called f")
	}
}

func TestPanicValue(t *testing.T) {
	errBad := errors.New("bad")
	tests := []struct {
		name string
		f    func()
		want any
	}{
		{`func() { panic("a problem") }`, func() { panic("a problem") }, "a problem"},
		{`func() { panic(42) }`, func() { panic(42) }, 42},
		{`func() { panic(errBad) }`, func() { panic(errBad) }, errBad},
	}
	for _, tt := range tests {
		guard(t, "PanicValue", func() {
			if got := PanicValue(tt.f); got != tt.want {
				t.Errorf("PanicValue(%s) = %v, want %v", tt.name, got, tt.want)
			}
		})
	}

	called := false
	if got := PanicValue(func() { called = true }); got != nil {
		t.Errorf("PanicValue(a function that does not panic) = %v, want nil", got)
	}
	if !called {
		t.Errorf("PanicValue(f) never called f")
	}
}

func TestSafeDivide(t *testing.T) {
	tests := []struct {
		a, b int
		want int
	}{
		{7, 2, 3},
		{-9, 3, -3},
		{0, 5, 0},
		{12, 4, 3},
	}
	for _, tt := range tests {
		got, err := SafeDivide(tt.a, tt.b)
		if err != nil {
			t.Errorf("SafeDivide(%d, %d) returned error %v, want nil", tt.a, tt.b, err)
		}
		if got != tt.want {
			t.Errorf("SafeDivide(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}

	guard(t, "SafeDivide(1, 0)", func() {
		got, err := SafeDivide(1, 0)
		if err == nil {
			t.Fatalf("SafeDivide(1, 0) returned a nil error, want the recovered panic as an error")
		}
		if want := "recovered: runtime error: integer divide by zero"; err.Error() != want {
			t.Errorf("SafeDivide(1, 0) returned the error %q, want %q", err, want)
		}
		if got != 0 {
			t.Errorf("SafeDivide(1, 0) = %d, want 0", got)
		}
	})
}

func TestTry(t *testing.T) {
	errFailed := errors.New("failed")

	called := false
	if err := Try(func() error { called = true; return nil }); err != nil {
		t.Errorf("Try(f returning nil) = %v, want nil", err)
	}
	if !called {
		t.Errorf("Try(f) never called f")
	}

	if err := Try(func() error { return errFailed }); !errors.Is(err, errFailed) {
		t.Errorf("Try(f returning errFailed) = %v, want errFailed", err)
	}

	guard(t, "Try", func() {
		err := Try(func() error { panic("boom") })
		if err == nil {
			t.Fatalf(`Try(f panicking with "boom") = nil, want an error`)
		}
		if want := "panic: boom"; err.Error() != want {
			t.Errorf(`Try(f panicking with "boom") = %q, want %q`, err, want)
		}
	})

	guard(t, "Try", func() {
		err := Try(func() error { panic(errFailed) })
		if err == nil {
			t.Fatalf("Try(f panicking with errFailed) = nil, want an error")
		}
		if !errors.Is(err, errFailed) {
			t.Errorf("Try(f panicking with errFailed) = %q, want an error that wraps errFailed (use %%w)", err)
		}
		if want := "panic: failed"; err.Error() != want {
			t.Errorf("Try(f panicking with errFailed) = %q, want %q", err, want)
		}
	})

	guard(t, "Try", func() {
		var m map[string]int
		err := Try(func() error { m["x"] = 1; return nil })
		if err == nil {
			t.Fatalf("Try(f writing to a nil map) = nil, want an error")
		}
		if want := "panic: assignment to entry in nil map"; err.Error() != want {
			t.Errorf("Try(f writing to a nil map) = %q, want %q", err, want)
		}
	})
}

func TestRunAll(t *testing.T) {
	var ran []string
	tasks := []func(){
		func() { ran = append(ran, "first") },
		func() { ran = append(ran, "second"); panic("boom") },
		func() { ran = append(ran, "third") },
		func() { ran = append(ran, "fourth"); panic(errors.New("out of coffee")) },
		func() { ran = append(ran, "fifth") },
	}
	wantMsgs := []string{"", "task 1 panicked: boom", "", "task 3 panicked: out of coffee", ""}

	guard(t, "RunAll", func() {
		errs := RunAll(tasks)
		if want := []string{"first", "second", "third", "fourth", "fifth"}; !slices.Equal(ran, want) {
			t.Errorf("RunAll ran the tasks %q, want all of them in order: %q", ran, want)
		}
		if len(errs) != len(tasks) {
			t.Fatalf("RunAll returned %d results for %d tasks, want one per task", len(errs), len(tasks))
		}
		for i, want := range wantMsgs {
			switch {
			case want == "" && errs[i] != nil:
				t.Errorf("errs[%d] = %q, want nil (task %d did not panic)", i, errs[i], i)
			case want != "" && errs[i] == nil:
				t.Errorf("errs[%d] = nil, want the error %q", i, want)
			case want != "" && errs[i].Error() != want:
				t.Errorf("errs[%d] = %q, want %q", i, errs[i], want)
			}
		}
	})

	if errs := RunAll(nil); len(errs) != 0 {
		t.Errorf("RunAll(nil) returned %d results, want none", len(errs))
	}
}
