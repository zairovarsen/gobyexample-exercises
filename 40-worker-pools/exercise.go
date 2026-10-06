//go:build !solution

// Package workerpools contains exercises for https://gobyexample.com/worker-pools
//
// Run the tests for this section with:
//
//	go test ./40-worker-pools/
package workerpools

// Worker is the body of one worker goroutine: it receives jobs until the
// jobs channel is closed and, for every job, sends the job's number times 2
// on results. It returns when jobs is closed and drained.
//
// Several workers can run this function on the same two channels at once;
// every job then goes to exactly one of them.
func Worker(jobs <-chan int, results chan<- int) {
	// TODO: range over jobs and send the results.
}

// DoubleAll doubles all the nums with a pool of workers, like the main
// function on the page:
//
//   - make a jobs and a results channel, both with room for len(nums) values;
//   - start `workers` goroutines that run Worker;
//   - send all the nums as jobs and close the jobs channel;
//   - collect len(nums) results.
//
// The results arrive in whatever order the workers finish, and they are
// returned in that order.
func DoubleAll(nums []int, workers int) []int {
	// TODO: set up the pool and collect the results.
	return nil
}

// Process is a worker pool for any kind of work: it calls work(job) for
// every job and returns the results, in whatever order they are finished.
//
// The calls of work are spread over exactly `workers` goroutines, so that
// at most that many calls are in progress at any moment. That limit is the
// point of a pool - do not start one goroutine per job.
func Process(jobs []int, workers int, work func(job int) int) []int {
	// TODO: a pool like in DoubleAll, with workers that call work.
	return nil
}

// MapOrdered applies f to every input with a pool of `workers` goroutines
// and returns the outputs in the order of the inputs: result[i] is
// f(inputs[i]), no matter which worker finished first.
//
// Hint: let the jobs and the results carry their index, e.g. as small
// structs with an index and a value, and put each result into its place in
// a slice of the right length.
func MapOrdered(inputs []string, workers int, f func(string) string) []string {
	// TODO: a pool whose results remember where they belong.
	return nil
}
