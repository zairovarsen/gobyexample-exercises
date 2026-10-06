//go:build solution

package timers

import "time"

func WaitTimer(d time.Duration) time.Time {
	timer := time.NewTimer(d)
	return <-timer.C
}

func Schedule(d time.Duration, f func()) *time.Timer {
	timer := time.NewTimer(d)
	go func() {
		<-timer.C
		f()
	}()
	return timer
}

func TryCancel(timer *time.Timer) string {
	if timer.Stop() {
		return "stopped"
	}
	return "already fired"
}

func Debounce(d time.Duration, f func()) (trigger func()) {
	var timer *time.Timer
	return func() {
		if timer != nil {
			timer.Stop()
		}
		timer = time.AfterFunc(d, f)
	}
}
