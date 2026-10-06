//go:build solution

package waitgroups

import "sync"

func RunAll(tasks ...func()) {
	var wg sync.WaitGroup
	for _, task := range tasks {
		wg.Go(task)
	}
	wg.Wait()
}

func Worker(id int, wg *sync.WaitGroup, report func(id int)) {
	defer wg.Done()
	report(id)
}

func MapConcurrently(nums []int, f func(n int) int) []int {
	results := make([]int, len(nums))
	var wg sync.WaitGroup
	for i, n := range nums {
		wg.Go(func() {
			results[i] = f(n)
		})
	}
	wg.Wait()
	return results
}

func Gather(producers ...func(out chan<- int)) []int {
	out := make(chan int)
	var wg sync.WaitGroup
	for _, produce := range producers {
		wg.Go(func() {
			produce(out)
		})
	}

	// Closing must wait for the producers, but the loop below must already
	// be receiving while they send: so the waiting gets its own goroutine.
	go func() {
		wg.Wait()
		close(out)
	}()

	var nums []int
	for n := range out {
		nums = append(nums, n)
	}
	return nums
}
