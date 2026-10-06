//go:build solution

package goroutines

import "time"

func RunAsync(f func()) {
	go f()
}

func StartEach(names []string, f func(name string)) {
	for _, name := range names {
		go func(name string) {
			f(name)
		}(name)
	}
}

func RunAndWait(wait time.Duration, fs ...func()) {
	for _, f := range fs {
		go f()
	}
	time.Sleep(wait)
}

func Tree(depth int, visit func(level int)) {
	if depth <= 0 {
		return
	}
	go func() {
		visit(depth)
		Tree(depth-1, visit)
		Tree(depth-1, visit)
	}()
}
