//go:build solution

package recursion

func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

func SumDigits(n int) int {
	if n < 10 {
		return n
	}
	return n%10 + SumDigits(n/10)
}

func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

func Sum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return nums[0] + Sum(nums[1:])
}

func Fib(n int) int {
	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	return fib(n)
}

func FibFast(n int) int {
	known := make(map[int]int)

	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		if v, ok := known[n]; ok {
			return v
		}
		v := fib(n-1) + fib(n-2)
		known[n] = v
		return v
	}
	return fib(n)
}
