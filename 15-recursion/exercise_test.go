package recursion

import (
	"testing"
	"time"
)

func TestFactorial(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{0, 1},
		{1, 1},
		{5, 120},
		{7, 5040},
		{10, 3628800},
		{20, 2432902008176640000},
	}
	for _, tt := range tests {
		if got := Factorial(tt.n); got != tt.want {
			t.Errorf("Factorial(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}

func TestSumDigits(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{7, 7},
		{10, 1},
		{123, 6},
		{9999, 36},
		{1000000, 1},
		{0, 0},
	}
	for _, tt := range tests {
		if got := SumDigits(tt.n); got != tt.want {
			t.Errorf("SumDigits(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}

func TestGCD(t *testing.T) {
	tests := []struct {
		a, b int
		want int
	}{
		{12, 18, 6},
		{18, 12, 6},
		{48, 36, 12},
		{17, 5, 1},
		{10, 0, 10},
		{0, 10, 10},
		{7, 7, 7},
	}
	for _, tt := range tests {
		if got := GCD(tt.a, tt.b); got != tt.want {
			t.Errorf("GCD(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestSum(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3, 4}, 10},
		{[]int{5}, 5},
		{[]int{-2, 2, 7}, 7},
		{nil, 0},
	}
	for _, tt := range tests {
		if got := Sum(tt.nums); got != tt.want {
			t.Errorf("Sum(%v) = %d, want %d", tt.nums, got, tt.want)
		}
	}
}

var fibTests = []struct {
	n    int
	want int
}{
	{0, 0},
	{1, 1},
	{2, 1},
	{3, 2},
	{7, 13},
	{10, 55},
	{20, 6765},
}

func TestFib(t *testing.T) {
	for _, tt := range fibTests {
		if got := Fib(tt.n); got != tt.want {
			t.Errorf("Fib(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}

func TestFibFast(t *testing.T) {
	for _, tt := range fibTests {
		if got := FibFast(tt.n); got != tt.want {
			t.Errorf("FibFast(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}

	// Without memoization this call would run for many years, so it gets a
	// deadline instead of the chance to hang the test.
	const n, want = 90, 2880067194370816120
	done := make(chan int, 1)
	go func() { done <- FibFast(n) }()
	select {
	case got := <-done:
		if got != want {
			t.Errorf("FibFast(%d) = %d, want %d", n, got, want)
		}
	case <-time.After(2 * time.Second):
		t.Fatalf("FibFast(%d) did not finish within 2 seconds: remember the results in a map", n)
	}
}
