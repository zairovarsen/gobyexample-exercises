//go:build solution

package signals

import (
	"os"
	"os/signal"
	"syscall"
	"time"
)

func NewSignalChannel() chan os.Signal {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	return sigs
}

func AwaitSignal(sigs <-chan os.Signal, onSignal func(os.Signal)) <-chan bool {
	done := make(chan bool, 1)
	go func() {
		sig := <-sigs
		onSignal(sig)
		done <- true
	}()
	return done
}

func RunUntilSignal(sigs <-chan os.Signal, ticks <-chan time.Time, work func()) (int, os.Signal) {
	count := 0
	for {
		select {
		case <-ticks:
			work()
			count++
		case sig := <-sigs:
			return count, sig
		}
	}
}
