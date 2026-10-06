//go:build solution

package timeouts

import (
	"errors"
	"time"
)

var ErrTimeout = errors.New("timeout")

func RecvTimeout(ch <-chan string, d time.Duration) (string, bool) {
	select {
	case msg := <-ch:
		return msg, true
	case <-time.After(d):
		return "", false
	}
}

func SendTimeout(ch chan<- string, msg string, d time.Duration) bool {
	select {
	case ch <- msg:
		return true
	case <-time.After(d):
		return false
	}
}

func WithTimeout(f func() string, d time.Duration) (string, error) {
	// Buffered, so that the goroutine can finish even after a timeout.
	result := make(chan string, 1)
	go func() {
		result <- f()
	}()

	select {
	case res := <-result:
		return res, nil
	case <-time.After(d):
		return "", ErrTimeout
	}
}

func CollectFor(ch <-chan int, d time.Duration) []int {
	var nums []int
	deadline := time.After(d)
	for {
		select {
		case n := <-ch:
			nums = append(nums, n)
		case <-deadline:
			return nums
		}
	}
}
