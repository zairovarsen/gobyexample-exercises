//go:build !solution

// Package contextx contains exercises for https://gobyexample.com/context
//
// Run the tests for this section with:
//
//	go test ./81-context/
package contextx

import (
	"context"
	"net/http"
	"time"
)

// A context.Context carries a cancellation signal across function calls and
// goroutines. Code that does slow work watches the channel ctx.Done(): once
// it is closed the work is no longer wanted (the caller gave up, a deadline
// passed, the HTTP client disconnected ...) and ctx.Err() tells why:
// context.Canceled or context.DeadlineExceeded.
//
// By convention ctx is the first parameter of a function.

// Wait pauses for the duration d, unless ctx is done first. It returns nil
// after a full pause and ctx.Err() when it was interrupted.
//
// This is the core of the page's handler: a select with two cases, one
// receiving from time.After(d) and one receiving from ctx.Done().
func Wait(ctx context.Context, d time.Duration) error {
	// TODO: select on the timer and on ctx.Done().
	return nil
}

// Poll calls check once per interval until check returns true; then Poll
// returns nil. When ctx is done before that, Poll stops calling check and
// returns ctx.Err().
//
// Example: waiting for a server to come up - check tries to connect, the
// context limits how long we are prepared to wait in total.
//
// Hint: a time.Ticker (stop it with defer) and a for loop around a select.
func Poll(ctx context.Context, interval time.Duration, check func() bool) error {
	// TODO: loop: wait for the next tick or for ctx.Done().
	return nil
}

// Generate returns an unbuffered channel and starts a goroutine that sends
// the numbers 0, 1, 2, ... on it, for as long as somebody receives them.
// When ctx is done the goroutine stops and closes the channel.
//
// Without the context the goroutine would stay blocked in its send forever
// once the receiver loses interest - a goroutine leak. The fix is a select
// with two cases: the send, and ctx.Done().
func Generate(ctx context.Context) <-chan int {
	// TODO: make the channel, start the goroutine, return the channel.
	return nil
}

// SlowHandler returns the HTTP handler from the page, with the duration of
// the simulated work as a parameter (the page waits 10 seconds).
//
// The handler waits for the duration work and then responds with "hello\n".
// While waiting it keeps an eye on the context of the request, r.Context(),
// which the server cancels when the client disconnects. In that case the
// handler stops waiting at once and responds with the status
// 500 Internal Server Error and the text of ctx.Err() as the body:
//
//	http.Error(w, err.Error(), http.StatusInternalServerError)
func SlowHandler(work time.Duration) http.HandlerFunc {
	// TODO: return a handler function that selects on the work and on the
	// request context.
	return nil
}

// Get is the other side of SlowHandler: an HTTP client that can be told to
// give up. It sends a GET request to url that is bound to ctx and returns
// the response body. A status other than 200 OK is an error.
//
// When ctx is done before the response has arrived, the request is aborted
// and the returned error wraps ctx.Err(), so that for instance
// errors.Is(err, context.DeadlineExceeded) is true - the client does that
// for you.
//
// http.Get cannot take a context. Build the request yourself with
// http.NewRequestWithContext(ctx, http.MethodGet, url, nil) and send it with
// http.DefaultClient.Do(req). Close the response body.
func Get(ctx context.Context, url string) (string, error) {
	// TODO: build the request with ctx, send it, check the status, read the
	// body.
	return "", nil
}
