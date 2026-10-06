//go:build solution

package tickers

import "time"

func Ticks(interval time.Duration, n int) []time.Time {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	times := make([]time.Time, 0, n)
	for range n {
		times = append(times, <-ticker.C)
	}
	return times
}

func Every(interval time.Duration, f func()) (stop func()) {
	ticker := time.NewTicker(interval)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				f()
			}
		}
	}()

	return func() {
		ticker.Stop()
		// Unbuffered: the send only completes once the goroutine is in its
		// select again, that is, not in the middle of a call of f.
		done <- true
	}
}

func TickFor(interval, total time.Duration) int {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	deadline := time.After(total)
	ticks := 0
	for {
		select {
		case <-ticker.C:
			ticks++
		case <-deadline:
			return ticks
		}
	}
}

func Poll(interval, timeout time.Duration, check func() bool) bool {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	timer := time.NewTimer(timeout)
	defer timer.Stop()

	for {
		select {
		case <-ticker.C:
			if check() {
				return true
			}
		case <-timer.C:
			return false
		}
	}
}
