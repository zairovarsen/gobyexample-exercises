//go:build solution

package selectx

import "time"

func First(c1, c2 <-chan string) (msg string, from int) {
	select {
	case msg := <-c1:
		return msg, 1
	case msg := <-c2:
		return msg, 2
	}
}

func Fastest(d1, d2 time.Duration) string {
	// Buffered, so that the slower goroutine is not stuck in its send forever.
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	go func() {
		time.Sleep(d1)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(d2)
		c2 <- "two"
	}()

	select {
	case msg := <-c1:
		return msg
	case msg := <-c2:
		return msg
	}
}

func ReceiveN(c1, c2 <-chan string, n int) []string {
	msgs := make([]string, 0, n)
	for range n {
		select {
		case msg := <-c1:
			msgs = append(msgs, msg)
		case msg := <-c2:
			msgs = append(msgs, msg)
		}
	}
	return msgs
}

func SendEither(c1, c2 chan<- string, msg string) int {
	select {
	case c1 <- msg:
		return 1
	case c2 <- msg:
		return 2
	}
}

func SumUntilQuit(nums <-chan int, quit <-chan bool) int {
	sum := 0
	for {
		select {
		case n := <-nums:
			sum += n
		case <-quit:
			return sum
		}
	}
}
