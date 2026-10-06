//go:build solution

package functions

func Plus(a int, b int) int {
	return a + b
}

func PlusPlus(a, b, c int) int {
	return Plus(Plus(a, b), c)
}

func Clamp(n, low, high int) int {
	if n < low {
		return low
	}
	if n > high {
		return high
	}
	return n
}

func Average(nums []float64) float64 {
	if len(nums) == 0 {
		return 0
	}
	total := 0.0
	for i := range len(nums) {
		total += nums[i]
	}
	return total / float64(len(nums))
}

func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	for d := 2; d*d <= n; d++ {
		if n%d == 0 {
			return false
		}
	}
	return true
}

func CountPrimes(limit int) int {
	count := 0
	for n := 2; n <= limit; n++ {
		if IsPrime(n) {
			count++
		}
	}
	return count
}
