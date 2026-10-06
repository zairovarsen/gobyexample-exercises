package contextx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

// within runs f in a goroutine and returns its result, failing the test
// instead of hanging when f does not return within 2 seconds.
func within[T any](t *testing.T, what string, f func() T) T {
	t.Helper()
	done := make(chan T, 1)
	go func() { done <- f() }()
	select {
	case v := <-done:
		return v
	case <-time.After(2 * time.Second):
		t.Fatalf("%s did not return within 2 seconds", what)
	}
	panic("unreachable")
}

// long is a duration no test waits for: whatever takes this long was not
// interrupted by its context.
const long = 30 * time.Second

func TestWait(t *testing.T) {
	// Nothing interrupts: the full pause, then nil.
	start := time.Now()
	err := within(t, "Wait(ctx, 30ms)", func() error {
		return Wait(context.Background(), 30*time.Millisecond)
	})
	if err != nil {
		t.Errorf("Wait(background context, 30ms) = %v, want nil", err)
	}
	if elapsed := time.Since(start); elapsed < 20*time.Millisecond {
		t.Errorf("Wait(background context, 30ms) returned after %v, want it to pause for the whole duration", elapsed)
	}

	// The context is cancelled while waiting.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	time.AfterFunc(30*time.Millisecond, cancel)
	err = within(t, "Wait(cancelled ctx, 30s)", func() error { return Wait(ctx, long) })
	if !errors.Is(err, context.Canceled) {
		t.Errorf("Wait(ctx cancelled after 30ms, 30s) = %v, want context.Canceled", err)
	}

	// The context was already cancelled before the call.
	err = within(t, "Wait(cancelled ctx, 30s)", func() error { return Wait(ctx, long) })
	if !errors.Is(err, context.Canceled) {
		t.Errorf("Wait(already cancelled ctx, 30s) = %v, want context.Canceled", err)
	}

	// The context has a deadline.
	ctx, cancelTimeout := context.WithTimeout(context.Background(), 30*time.Millisecond)
	defer cancelTimeout()
	err = within(t, "Wait(ctx with timeout, 30s)", func() error { return Wait(ctx, long) })
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Errorf("Wait(ctx with 30ms timeout, 30s) = %v, want context.DeadlineExceeded", err)
	}
}

func TestPoll(t *testing.T) {
	// check succeeds on its third call.
	var calls atomic.Int32
	err := within(t, "Poll(ctx, 5ms, check)", func() error {
		return Poll(context.Background(), 5*time.Millisecond, func() bool {
			return calls.Add(1) == 3
		})
	})
	if err != nil {
		t.Errorf("Poll with a check that succeeds on the third call = %v, want nil", err)
	}
	if got := calls.Load(); got != 3 {
		t.Errorf("Poll called check %d times, want 3 (until it returned true)", got)
	}

	// check never succeeds; the context puts an end to it.
	calls.Store(0)
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	err = within(t, "Poll(ctx with timeout, 5ms, check)", func() error {
		return Poll(ctx, 5*time.Millisecond, func() bool {
			calls.Add(1)
			return false
		})
	})
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Errorf("Poll(ctx with 100ms timeout) with a check that always fails = %v, want context.DeadlineExceeded", err)
	}
	if got := calls.Load(); got < 2 {
		t.Errorf("Poll called check %d times in 100ms with an interval of 5ms, want it to keep polling", got)
	}

	// After Poll has returned, check must not be called any more.
	before := calls.Load()
	time.Sleep(30 * time.Millisecond)
	if after := calls.Load(); after != before {
		t.Errorf("check was called %d more times after Poll returned; stop the polling", after-before)
	}

	// A cancelled context ends a slow poll right away.
	ctx, cancelNow := context.WithCancel(context.Background())
	cancelNow()
	err = within(t, "Poll(cancelled ctx, 30s, check)", func() error {
		return Poll(ctx, long, func() bool { return false })
	})
	if !errors.Is(err, context.Canceled) {
		t.Errorf("Poll(already cancelled ctx, 30s) = %v, want context.Canceled", err)
	}
}

func TestGenerate(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := Generate(ctx)
	if ch == nil {
		t.Fatal("Generate(ctx) returned a nil channel - did you create it with make?")
	}
	if cap(ch) != 0 {
		t.Errorf("Generate(ctx) returned a channel with capacity %d, want an unbuffered channel", cap(ch))
	}
	for want := range 5 {
		select {
		case got, ok := <-ch:
			if !ok {
				t.Fatalf("the channel was closed after %d values although ctx is not done", want)
			}
			if got != want {
				t.Fatalf("received %d, want %d", got, want)
			}
		case <-time.After(2 * time.Second):
			t.Fatalf("timed out waiting for the value %d", want)
		}
	}

	// After cancel the goroutine has to stop and close the channel. A few
	// more numbers may still arrive: when both cases of its select are
	// ready, the goroutine picks one at random.
	cancel()
	deadline := time.After(2 * time.Second)
	for extra := 0; ; extra++ {
		select {
		case _, ok := <-ch:
			if !ok {
				return // closed: the goroutine is gone
			}
			if extra > 1000 {
				t.Fatal("still receiving numbers after cancel; select on ctx.Done() next to the send")
			}
		case <-deadline:
			t.Fatal("the channel was not closed within 2 seconds after cancel")
		}
	}
}

// handler returns SlowHandler(work), failing the test for a nil handler.
func handler(t *testing.T, work time.Duration) http.HandlerFunc {
	t.Helper()
	h := SlowHandler(work)
	if h == nil {
		t.Fatalf("SlowHandler(%v) returned nil", work)
	}
	return h
}

func TestSlowHandler(t *testing.T) {
	// The work finishes: 200 and "hello".
	h := handler(t, 30*time.Millisecond)
	start := time.Now()
	rec := within(t, "the handler", func() *httptest.ResponseRecorder {
		rec := httptest.NewRecorder()
		h(rec, httptest.NewRequest("GET", "/hello", nil))
		return rec
	})
	if rec.Code != http.StatusOK || rec.Body.String() != "hello\n" {
		t.Errorf("SlowHandler(30ms) responded %d %q, want 200 %q", rec.Code, rec.Body.String(), "hello\n")
	}
	if elapsed := time.Since(start); elapsed < 20*time.Millisecond {
		t.Errorf("SlowHandler(30ms) responded after %v, want it to work for the whole duration", elapsed)
	}

	// The request context is cancelled while the handler works: 500 and the
	// context's error.
	h = handler(t, long)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	time.AfterFunc(30*time.Millisecond, cancel)
	rec = within(t, "the handler (after the request context was cancelled)", func() *httptest.ResponseRecorder {
		rec := httptest.NewRecorder()
		h(rec, httptest.NewRequest("GET", "/hello", nil).WithContext(ctx))
		return rec
	})
	if rec.Code != http.StatusInternalServerError {
		t.Errorf("SlowHandler with a cancelled request responded with status %d, want 500", rec.Code)
	}
	if got, want := rec.Body.String(), "context canceled\n"; got != want {
		t.Errorf("SlowHandler with a cancelled request responded with the body %q, want %q", got, want)
	}
}

func TestSlowHandlerClientDisconnects(t *testing.T) {
	// A real server: when the client gives up and closes the connection, the
	// server cancels r.Context() and the handler must stop working.
	h := handler(t, 3*time.Second)
	finished := make(chan struct{})
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer close(finished)
		h(w, r)
	}))
	defer srv.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", srv.URL, nil)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := srv.Client().Do(req)
	if err == nil {
		resp.Body.Close()
		t.Fatal("the client got a response within 50ms although the handler should work for 3s")
	}

	select {
	case <-finished:
	case <-time.After(2 * time.Second):
		t.Fatal("the handler kept working after the client had disconnected; select on r.Context().Done()")
	}
}

func TestGet(t *testing.T) {
	release := make(chan struct{})
	mux := http.NewServeMux()
	mux.HandleFunc("/fast", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello\n")
	})
	mux.HandleFunc("/missing", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "nothing here", http.StatusNotFound)
	})
	mux.HandleFunc("/slow", func(w http.ResponseWriter, r *http.Request) {
		// Answers only when the client has given up or the test is over.
		select {
		case <-r.Context().Done():
		case <-release:
		case <-time.After(3 * time.Second):
		}
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()
	defer close(release)

	type result struct {
		body string
		err  error
	}
	get := func(ctx context.Context, path string) result {
		t.Helper()
		return within(t, "Get(ctx, server"+path+")", func() result {
			body, err := Get(ctx, srv.URL+path)
			return result{body, err}
		})
	}

	res := get(context.Background(), "/fast")
	if res.err != nil || res.body != "hello\n" {
		t.Errorf("Get(ctx, server/fast) = %q, %v, want %q, nil", res.body, res.err, "hello\n")
	}

	res = get(context.Background(), "/missing")
	if res.err == nil {
		t.Errorf("Get(ctx, server/missing) = %q, nil - want an error for the status 404", res.body)
	} else if !strings.Contains(res.err.Error(), "404") {
		t.Errorf("Get(ctx, server/missing) error = %q, want it to mention the status 404", res.err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	start := time.Now()
	res = get(ctx, "/slow")
	if !errors.Is(res.err, context.DeadlineExceeded) {
		t.Errorf("Get(ctx with 50ms timeout, server/slow) = %q, %v, want an error wrapping context.DeadlineExceeded", res.body, res.err)
	}
	if elapsed := time.Since(start); elapsed > time.Second {
		t.Errorf("Get(ctx with 50ms timeout, server/slow) took %v, want it to give up after about 50ms", elapsed)
	}

	cancelled, cancelNow := context.WithCancel(context.Background())
	cancelNow()
	res = get(cancelled, "/fast")
	if !errors.Is(res.err, context.Canceled) {
		t.Errorf("Get(already cancelled ctx, server/fast) = %q, %v, want an error wrapping context.Canceled", res.body, res.err)
	}
}
