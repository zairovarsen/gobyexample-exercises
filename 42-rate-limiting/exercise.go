//go:build !solution

// Package ratelimiting contains exercises for https://gobyexample.com/rate-limiting
//
// Run the tests for this section with:
//
//	go test ./42-rate-limiting/
package ratelimiting

import "time"

// Throttle handles the requests at a limited rate: it calls handle(req) for
// every request from the requests channel, in order, but only one per
// interval - also the first request has to wait for one interval. Throttle
// returns when requests is closed and all requests are handled.
//
// The limiter is a channel that delivers a value at regular intervals, as
// time.Tick(interval) provides it (or the channel C of a time.NewTicker).
// Blocking on a receive from the limiter before each request does the
// throttling.
func Throttle(requests <-chan int, interval time.Duration, handle func(req int)) {
	// TODO: range over requests; wait for a tick before each one.
}

// NewBurstyLimiter returns a limiter that allows short bursts. The limiter
// is a buffered channel of capacity burst that works like a bucket of
// tokens: taking a value out of it gives permission for one request.
//
//   - The bucket starts out full, with burst tokens (the token values do not
//     matter - time.Now() will do), so the first burst requests need not wait.
//   - A goroutine puts one more token into the bucket per interval. When the
//     bucket is full the goroutine just waits until there is room again, so
//     unused tokens never pile up beyond burst.
//
// NewBurstyLimiter returns immediately. The stop function, which the caller
// calls once, ends the refilling for good: it makes the goroutine return and
// stops its ticker. Tokens that are in the bucket at that moment stay there.
//
// Hint: the goroutine needs a done channel in addition to its ticker (as in
// the tickers section), and its send into the bucket must not get stuck
// when stop is called - a select handles both.
func NewBurstyLimiter(burst int, interval time.Duration) (limiter <-chan time.Time, stop func()) {
	// TODO: fill the bucket, start the refilling goroutine, return both.
	return nil, nil
}

// BurstyThrottle handles the requests like Throttle does, but with a bursty
// limiter from NewBurstyLimiter: the first burst requests are handled
// without any delay, the remaining ones at one per interval. It returns when
// requests is closed and all requests are handled, and it stops its limiter
// before returning.
func BurstyThrottle(requests <-chan int, burst int, interval time.Duration, handle func(req int)) {
	// TODO: create a bursty limiter and take a token before each request.
}

// Shed handles requests under a limiter without ever making a request wait:
// for every request from the requests channel it checks whether the limiter
// has a token available right now. If so, it takes the token and calls
// handle(req); if not, the request is dropped. Shed returns the number of
// dropped requests when requests is closed and drained.
//
// (This is called load shedding: when there is too much to do, turn work
// away at once instead of building up a queue.)
//
// Hint: a non-blocking receive - a select with a default case.
func Shed(requests <-chan int, limiter <-chan time.Time, handle func(req int)) (dropped int) {
	// TODO: range over requests with a non-blocking receive from limiter.
	return 0
}
