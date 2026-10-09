//go:build !solution

// Package functions contains exercises for https://gobyexample.com/functions
//
// Run the tests for this section with:
//
//	go test ./11-functions/
package functions

// Plus returns the sum of a and b.
//
// Go requires an explicit return: it never returns the value of the last
// expression on its own.
func Plus(a int, b int) int {
	// TODO: return the sum.
	return a + b
}

// PlusPlus returns the sum of a, b and c. Its parameters share one type
// declaration: (a, b, c int).
//
// Do not use + more than once here: call Plus instead.
func PlusPlus(a, b, c int) int {
	// TODO: build the result from calls to Plus.
	return Plus(Plus(a,b),c)
}

// Clamp forces n into the range from low to high: it returns low when n is
// smaller than low, high when n is bigger than high, and n itself otherwise.
//
// Example: Clamp(15, 0, 10) returns 10 and Clamp(5, 0, 10) returns 5.
func Clamp(n, low, high int) int {
	if n < low {
		return low
	} else if n > high {
		return high
	} else {
		return n
	}
}

// Average returns the mean of nums: Average([1 2 4]) returns 2.3333...
// For an empty slice it returns 0.
//
// Hint: float64(len(nums)) converts the length to a float.
func Average(nums []float64) float64 {
	if len(nums) == 0 {
		return 0
	}
	// TODO: sum the numbers and divide by how many there are.
	sum := 0.0
	for i := range len(nums) {
		sum += nums[i]
	}
	return sum / float64(len(nums))
}

// IsPrime reports whether n is a prime number: a number greater than 1 that
// is divisible only by 1 and by itself. 2, 3, 5, 7 and 11 are prime; 0, 1,
// 4 and 9 are not, and neither is any negative number.
func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	divisor := 2
	
	for {
		if divisor == n {
			break
		}
		
		if n % divisor == 0 {
			return false
		}

		divisor++
	}
	// TODO: look for a divisor between 2 and n-1.
	return true
}

// CountPrimes returns how many prime numbers there are from 1 up to and
// including limit: CountPrimes(10) returns 4 (2, 3, 5 and 7).
//
// This is what functions are for: let IsPrime do the hard part.
func CountPrimes(limit int) int {
	// TODO: call IsPrime in a loop and count.
	count := 0
	for i := range limit {
		if (IsPrime(i)) {
			count++	
		}
	}

	if (IsPrime(limit)) {
		count++
	}

	return count
}
