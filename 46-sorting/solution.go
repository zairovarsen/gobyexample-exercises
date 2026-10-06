//go:build solution

package sorting

import (
	"cmp"
	"slices"
)

func SortStrings(strs []string) {
	slices.Sort(strs)
}

func SortedInts(nums []int) []int {
	sorted := slices.Clone(nums)
	slices.Sort(sorted)
	return sorted
}

func IsSorted(nums []int) bool {
	return slices.IsSorted(nums)
}

func SortedKeys(m map[string]int) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	return keys
}

func TopN[T cmp.Ordered](items []T, n int) []T {
	if n <= 0 {
		return nil
	}
	sorted := slices.Clone(items)
	slices.Sort(sorted)
	slices.Reverse(sorted)
	if n > len(sorted) {
		n = len(sorted)
	}
	return sorted[:n]
}
