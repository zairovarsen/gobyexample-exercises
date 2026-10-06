//go:build solution

package channels

func Ping(msg string) chan string {
	messages := make(chan string)
	go func() { messages <- msg }()
	return messages
}

func SumAsync(nums []int) chan int {
	result := make(chan int)
	go func() {
		total := 0
		for _, n := range nums {
			total += n
		}
		result <- total
	}()
	return result
}

func Relay(in, out chan string) {
	msg := <-in
	out <- "relayed: " + msg
}

func ParallelSquares(nums []int) []int {
	type result struct {
		index, square int
	}
	results := make(chan result)
	for i, n := range nums {
		go func() {
			results <- result{i, n * n}
		}()
	}
	squares := make([]int, len(nums))
	for range nums {
		r := <-results
		squares[r.index] = r.square
	}
	return squares
}
