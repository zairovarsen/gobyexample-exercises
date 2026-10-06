//go:build solution

package closures

func Counter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func Multiplier(factor int) func(int) int {
	return func(n int) int {
		return n * factor
	}
}

func Accumulator(start int) func(delta int) int {
	total := start
	return func(delta int) int {
		total += delta
		return total
	}
}

func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		current := a
		a, b = b, a+b
		return current
	}
}

func CounterWithReset() (next func() int, reset func()) {
	i := 0
	next = func() int {
		i++
		return i
	}
	reset = func() {
		i = 0
	}
	return next, reset
}

func Memoize(f func(int) int) func(int) int {
	cache := make(map[int]int)
	return func(n int) int {
		if v, ok := cache[n]; ok {
			return v
		}
		v := f(n)
		cache[n] = v
		return v
	}
}
