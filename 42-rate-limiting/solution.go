//go:build solution

package ratelimiting

import "time"

func Throttle(requests <-chan int, interval time.Duration, handle func(req int)) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for req := range requests {
		<-ticker.C
		handle(req)
	}
}

func NewBurstyLimiter(burst int, interval time.Duration) (limiter <-chan time.Time, stop func()) {
	bucket := make(chan time.Time, burst)
	for range burst {
		bucket <- time.Now()
	}

	ticker := time.NewTicker(interval)
	done := make(chan struct{})
	go func() {
		defer ticker.Stop()
		for {
			select {
			case t := <-ticker.C:
				// Blocks while the bucket is full, but not beyond stop.
				select {
				case bucket <- t:
				case <-done:
					return
				}
			case <-done:
				return
			}
		}
	}()

	return bucket, func() { close(done) }
}

func BurstyThrottle(requests <-chan int, burst int, interval time.Duration, handle func(req int)) {
	limiter, stop := NewBurstyLimiter(burst, interval)
	defer stop()

	for req := range requests {
		<-limiter
		handle(req)
	}
}

func Shed(requests <-chan int, limiter <-chan time.Time, handle func(req int)) (dropped int) {
	for req := range requests {
		select {
		case <-limiter:
			handle(req)
		default:
			dropped++
		}
	}
	return dropped
}
