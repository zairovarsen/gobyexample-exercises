//go:build solution

package randomnumbers

import (
	"math/rand/v2"
	"slices"
)

func RollDie() int {
	return rand.IntN(6) + 1
}

func FloatBetween(lo, hi float64) float64 {
	return lo + rand.Float64()*(hi-lo)
}

func SeededInts(seed1, seed2 uint64, n int) []int {
	r := rand.New(rand.NewPCG(seed1, seed2))
	nums := make([]int, n)
	for i := range nums {
		nums[i] = r.IntN(100)
	}
	return nums
}

func Pick(r *rand.Rand, items []string) string {
	if len(items) == 0 {
		return ""
	}
	return items[r.IntN(len(items))]
}

func Shuffled(seed uint64, items []int) []int {
	r := rand.New(rand.NewPCG(seed, seed))
	shuffled := slices.Clone(items)
	for i := len(shuffled) - 1; i > 0; i-- {
		j := r.IntN(i + 1)
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	}
	return shuffled
}
