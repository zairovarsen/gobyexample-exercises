//go:build solution

package nonblockingchanneloperations

import "fmt"

func TryReceive(messages <-chan string) (string, bool) {
	select {
	case msg := <-messages:
		return msg, true
	default:
		return "", false
	}
}

func TrySend(messages chan<- string, msg string) bool {
	select {
	case messages <- msg:
		return true
	default:
		return false
	}
}

func Poll(messages <-chan string, signals <-chan bool) string {
	select {
	case msg := <-messages:
		return "received message " + msg
	case sig := <-signals:
		return fmt.Sprint("received signal ", sig)
	default:
		return "no activity"
	}
}

func DrainAvailable(ch <-chan int) []int {
	var vals []int
	for {
		select {
		case v := <-ch:
			vals = append(vals, v)
		default:
			return vals
		}
	}
}

func Broadcast(subscribers []chan string, msg string) int {
	delivered := 0
	for _, sub := range subscribers {
		select {
		case sub <- msg:
			delivered++
		default:
		}
	}
	return delivered
}
