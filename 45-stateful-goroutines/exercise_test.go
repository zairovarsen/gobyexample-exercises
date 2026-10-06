package statefulgoroutines

import (
	"reflect"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

const timeout = 2 * time.Second

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
	case <-time.After(5 * time.Second):
		t.Fatalf("%s did not return within 5 seconds", what)
	}
}

// ask sends the request op on ch and waits for the answer on resp, the way
// a client of the store does. It reports false instead of hanging forever
// when the store does not take the request or does not answer.
func ask[Op, R any](ch chan<- Op, op Op, resp <-chan R) (R, bool) {
	var zero R
	select {
	case ch <- op:
	case <-time.After(timeout):
		return zero, false
	}
	select {
	case r := <-resp:
		return r, true
	case <-time.After(timeout):
		return zero, false
	}
}

func storeRead(reads chan<- ReadOp, key int) (int, bool) {
	resp := make(chan int)
	return ask(reads, ReadOp{Key: key, Resp: resp}, resp)
}

func storeWrite(writes chan<- WriteOp, key, val int) bool {
	resp := make(chan bool)
	confirmed, ok := ask(writes, WriteOp{Key: key, Val: val, Resp: resp}, resp)
	return ok && confirmed
}

func TestStartStore(t *testing.T) {
	var reads chan<- ReadOp
	var writes chan<- WriteOp
	finishes(t, "StartStore() (the loop belongs in a goroutine)", func() { reads, writes = StartStore() })
	if reads == nil || writes == nil {
		t.Fatal("StartStore() returned a nil channel, want two channels made with make")
	}

	mustRead := func(key, want int) {
		t.Helper()
		got, ok := storeRead(reads, key)
		if !ok {
			t.Fatalf("a ReadOp for the key %d was not answered within %v", key, timeout)
		}
		if got != want {
			t.Errorf("a ReadOp for the key %d was answered with %d, want %d", key, got, want)
		}
	}
	mustWrite := func(key, val int) {
		t.Helper()
		if !storeWrite(writes, key, val) {
			t.Fatalf("a WriteOp for the key %d was not confirmed with true within %v", key, timeout)
		}
	}

	mustRead(1, 0)
	mustWrite(1, 42)
	mustRead(1, 42)
	mustRead(2, 0)
	mustWrite(2, -5)
	mustWrite(1, 7)
	mustRead(1, 7)
	mustRead(2, -5)

	// Many clients at once, each with a key of its own.
	var failed atomic.Int32
	var wg sync.WaitGroup
	for g := range 20 {
		wg.Go(func() {
			key := 100 + g
			for i := range 20 {
				if !storeWrite(writes, key, g*1000+i) {
					failed.Add(1)
					return
				}
				if got, ok := storeRead(reads, key); !ok || got != g*1000+i {
					failed.Add(1)
					return
				}
			}
		})
	}
	finishes(t, "the concurrent clients of the store", wg.Wait)
	if got := failed.Load(); got != 0 {
		t.Errorf("%d of 20 concurrent clients did not read back the value they had just written under their own key", got)
	}
}

func TestRead(t *testing.T) {
	// A stand-in for the store that answers every read with key * 10.
	reads := make(chan ReadOp)
	go func() {
		for op := range reads {
			if op.Resp != nil {
				op.Resp <- op.Key * 10
			}
		}
	}()

	for _, key := range []int{7, 0, -3} {
		var got int
		finishes(t, "Read (it needs a Resp channel made with make, and has to wait for the answer on it)", func() { got = Read(reads, key) })
		if got != key*10 {
			t.Errorf("Read(reads, %d) = %d, want %d (the store of this test answers with key * 10)", key, got, key*10)
		}
	}
}

func TestWrite(t *testing.T) {
	// A stand-in for the store that takes its time before it confirms.
	writes := make(chan WriteOp)
	got := make(chan WriteOp, 10)
	var confirming atomic.Bool
	go func() {
		for op := range writes {
			got <- op
			time.Sleep(20 * time.Millisecond)
			confirming.Store(true)
			if op.Resp != nil {
				op.Resp <- true
			}
		}
	}()

	finishes(t, "Write (it needs a Resp channel made with make, and has to wait for the confirmation on it)", func() { Write(writes, 3, 99) })

	select {
	case op := <-got:
		if op.Key != 3 || op.Val != 99 {
			t.Errorf("Write(writes, 3, 99) sent a WriteOp with Key %d and Val %d, want 3 and 99", op.Key, op.Val)
		}
		if op.Resp == nil {
			t.Error("Write(writes, 3, 99) sent a WriteOp with a nil Resp channel, want one made with make")
		}
	default:
		t.Fatal("Write(writes, 3, 99) returned without sending a WriteOp")
	}
	if !confirming.Load() {
		t.Error("Write(writes, 3, 99) returned before the store confirmed the write, want it to wait for the value on Resp")
	}
}

func TestCountOps(t *testing.T) {
	tests := []struct {
		readers, writers, opsEach int
		wantReads, wantWrites     uint64
	}{
		{10, 3, 20, 200, 60},
		{1, 1, 1, 1, 1},
		{4, 0, 5, 20, 0},
		{0, 2, 50, 0, 100},
	}
	for _, tt := range tests {
		var reads, writes uint64
		finishes(t, "CountOps", func() { reads, writes = CountOps(tt.readers, tt.writers, tt.opsEach) })
		if reads != tt.wantReads || writes != tt.wantWrites {
			t.Errorf("CountOps(%d, %d, %d) = %d, %d, want %d, %d", tt.readers, tt.writers, tt.opsEach, reads, writes, tt.wantReads, tt.wantWrites)
		}
	}
}

func TestCounter(t *testing.T) {
	// The state must live in a goroutine, reached through channels.
	typ := reflect.TypeFor[Counter]()
	channels := 0
	for i := range typ.NumField() {
		ft := typ.Field(i).Type
		if ft.Kind() == reflect.Chan {
			channels++
		}
		if pkg := ft.PkgPath(); pkg == "sync" || pkg == "sync/atomic" {
			t.Errorf("Counter has a field of type %v, want it to work without the packages sync and sync/atomic", ft)
		}
	}
	if channels == 0 {
		t.Fatal("Counter has no channel fields, want it to talk to a state-owning goroutine through channels")
	}

	c := NewCounter()
	if c == nil {
		t.Fatal("NewCounter() = nil, want a Counter")
	}

	var got int
	finishes(t, "Value", func() { got = c.Value() })
	if got != 0 {
		t.Errorf("new Counter: Value() = %d, want 0", got)
	}
	finishes(t, "Inc", func() {
		c.Inc()
		c.Inc()
		c.Inc()
	})
	finishes(t, "Value", func() { got = c.Value() })
	if got != 3 {
		t.Errorf("Counter after 3 calls of Inc(): Value() = %d, want 3", got)
	}

	const goroutines, perGoroutine = 50, 100
	var wg sync.WaitGroup
	for range goroutines {
		wg.Go(func() {
			for range perGoroutine {
				c.Inc()
				c.Value()
			}
		})
	}
	finishes(t, "the concurrent calls of Inc and Value", wg.Wait)
	finishes(t, "Value", func() { got = c.Value() })
	if want := 3 + goroutines*perGoroutine; got != want {
		t.Errorf("Counter after %d goroutines did %d more calls of Inc() each: Value() = %d, want %d", goroutines, perGoroutine, got, want)
	}

	finishes(t, "Close", c.Close)
}
