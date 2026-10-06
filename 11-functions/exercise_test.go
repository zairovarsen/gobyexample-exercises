package functions

import (
	"math"
	"testing"
)

func TestPlus(t *testing.T) {
	tests := []struct {
		a, b int
		want int
	}{
		{1, 2, 3},
		{-5, 2, -3},
		{40, 2, 42},
	}
	for _, tt := range tests {
		if got := Plus(tt.a, tt.b); got != tt.want {
			t.Errorf("Plus(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestPlusPlus(t *testing.T) {
	tests := []struct {
		a, b, c int
		want    int
	}{
		{1, 2, 3, 6},
		{10, -20, 5, -5},
		{0, 0, 7, 7},
	}
	for _, tt := range tests {
		if got := PlusPlus(tt.a, tt.b, tt.c); got != tt.want {
			t.Errorf("PlusPlus(%d, %d, %d) = %d, want %d", tt.a, tt.b, tt.c, got, tt.want)
		}
	}
}

func TestClamp(t *testing.T) {
	tests := []struct {
		n, low, high int
		want         int
	}{
		{5, 0, 10, 5},
		{15, 0, 10, 10},
		{-3, 0, 10, 0},
		{10, 0, 10, 10},
		{-7, -5, -1, -5},
		{3, 3, 3, 3},
	}
	for _, tt := range tests {
		if got := Clamp(tt.n, tt.low, tt.high); got != tt.want {
			t.Errorf("Clamp(%d, %d, %d) = %d, want %d", tt.n, tt.low, tt.high, got, tt.want)
		}
	}
}

func TestAverage(t *testing.T) {
	tests := []struct {
		nums []float64
		want float64
	}{
		{[]float64{1, 2, 4}, 7.0 / 3.0},
		{[]float64{2.5}, 2.5},
		{[]float64{1, 2, 3, 4}, 2.5},
		{[]float64{-1, 1}, 0},
		{nil, 0},
	}
	for _, tt := range tests {
		if got := Average(tt.nums); math.IsNaN(got) || math.Abs(got-tt.want) > 1e-9 {
			t.Errorf("Average(%v) = %v, want %v", tt.nums, got, tt.want)
		}
	}
}

func TestIsPrime(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{2, true},
		{3, true},
		{5, true},
		{11, true},
		{97, true},
		{7919, true},
		{0, false},
		{1, false},
		{4, false},
		{9, false},
		{91, false},
		{-7, false},
	}
	for _, tt := range tests {
		if got := IsPrime(tt.n); got != tt.want {
			t.Errorf("IsPrime(%d) = %v, want %v", tt.n, got, tt.want)
		}
	}
}

func TestCountPrimes(t *testing.T) {
	tests := []struct {
		limit int
		want  int
	}{
		{10, 4},
		{2, 1},
		{100, 25},
		{1000, 168},
		{1, 0},
	}
	for _, tt := range tests {
		if got := CountPrimes(tt.limit); got != tt.want {
			t.Errorf("CountPrimes(%d) = %d, want %d", tt.limit, got, tt.want)
		}
	}
}
