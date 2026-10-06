//go:build solution

package workerpools

func Worker(jobs <-chan int, results chan<- int) {
	for j := range jobs {
		results <- j * 2
	}
}

func DoubleAll(nums []int, workers int) []int {
	jobs := make(chan int, len(nums))
	results := make(chan int, len(nums))

	for range workers {
		go Worker(jobs, results)
	}

	for _, n := range nums {
		jobs <- n
	}
	close(jobs)

	doubled := make([]int, 0, len(nums))
	for range nums {
		doubled = append(doubled, <-results)
	}
	return doubled
}

func Process(jobs []int, workers int, work func(job int) int) []int {
	jobCh := make(chan int, len(jobs))
	resultCh := make(chan int, len(jobs))

	for range workers {
		go func() {
			for j := range jobCh {
				resultCh <- work(j)
			}
		}()
	}

	for _, j := range jobs {
		jobCh <- j
	}
	close(jobCh)

	results := make([]int, 0, len(jobs))
	for range jobs {
		results = append(results, <-resultCh)
	}
	return results
}

func MapOrdered(inputs []string, workers int, f func(string) string) []string {
	type item struct {
		index int
		value string
	}
	jobs := make(chan item, len(inputs))
	results := make(chan item, len(inputs))

	for range workers {
		go func() {
			for j := range jobs {
				results <- item{j.index, f(j.value)}
			}
		}()
	}

	for i, in := range inputs {
		jobs <- item{i, in}
	}
	close(jobs)

	outputs := make([]string, len(inputs))
	for range inputs {
		r := <-results
		outputs[r.index] = r.value
	}
	return outputs
}
