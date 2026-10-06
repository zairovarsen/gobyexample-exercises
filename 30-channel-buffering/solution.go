//go:build solution

package channelbuffering

func NewMailbox(size int) chan string {
	return make(chan string, size)
}

func Preload(msgs ...string) chan string {
	messages := make(chan string, len(msgs))
	for _, msg := range msgs {
		messages <- msg
	}
	return messages
}

func FillUp(ch chan int, v int) int {
	sent := 0
	for len(ch) < cap(ch) {
		ch <- v
		sent++
	}
	return sent
}

func KeepLatest(ch chan int, v int) {
	if len(ch) == cap(ch) {
		<-ch
	}
	ch <- v
}
