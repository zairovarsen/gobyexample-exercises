//go:build solution

package closingchannels

func SendAndClose(nums ...int) <-chan int {
	ch := make(chan int, len(nums))
	for _, n := range nums {
		ch <- n
	}
	close(ch)
	return ch
}

func ReceiveOr(ch <-chan string, def string) string {
	v, ok := <-ch
	if !ok {
		return def
	}
	return v
}

func SumUntilClosed(ch <-chan int) (sum, count int) {
	for {
		n, more := <-ch
		if !more {
			return sum, count
		}
		sum += n
		count++
	}
}

func StartWorker(jobs <-chan int, handle func(job int)) <-chan bool {
	done := make(chan bool, 1)
	go func() {
		for {
			j, more := <-jobs
			if !more {
				done <- true
				return
			}
			handle(j)
		}
	}()
	return done
}

func Square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			n, more := <-in
			if !more {
				return
			}
			out <- n * n
		}
	}()
	return out
}
