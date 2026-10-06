//go:build solution

package channeldirections

func Ping(pings chan<- string, msg string) {
	pings <- msg
}

func Pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func Generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

func Double(in <-chan int, out chan<- int, n int) {
	for range n {
		out <- 2 * <-in
	}
}

func Pipeline(nums []int) []int {
	in := Generate(nums...)
	out := make(chan int)
	go Double(in, out, len(nums))

	doubled := make([]int, 0, len(nums))
	for range nums {
		doubled = append(doubled, <-out)
	}
	return doubled
}
