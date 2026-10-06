//go:build solution

package rangeoverchannels

func Collect(queue <-chan string) []string {
	var elems []string
	for elem := range queue {
		elems = append(elems, elem)
	}
	return elems
}

func Count(n int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 1; i <= n; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func Evens(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			if n%2 == 0 {
				out <- n
			}
		}
		close(out)
	}()
	return out
}

func SumOfEvens(n int) int {
	sum := 0
	for even := range Evens(Count(n)) {
		sum += even
	}
	return sum
}

func Merge(a, b <-chan int) <-chan int {
	out := make(chan int)
	done := make(chan bool)

	forward := func(in <-chan int) {
		for n := range in {
			out <- n
		}
		done <- true
	}
	go forward(a)
	go forward(b)

	go func() {
		<-done
		<-done
		close(out)
	}()
	return out
}
