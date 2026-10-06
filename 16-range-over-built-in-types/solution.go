//go:build solution

package rangeoverbuiltintypes

import "slices"

func Sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func IndexOf(words []string, target string) int {
	for i, word := range words {
		if word == target {
			return i
		}
	}
	return -1
}

func Keys(m map[string]int) []string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	return keys
}

func Invert(m map[string]string) map[string]string {
	inverted := make(map[string]string)
	for k, v := range m {
		inverted[v] = k
	}
	return inverted
}

func ByteOffsets(s string) []int {
	var offsets []int
	for i := range s {
		offsets = append(offsets, i)
	}
	return offsets
}

func GroupByLength(words []string) map[int][]string {
	groups := make(map[int][]string)
	for _, word := range words {
		groups[len(word)] = append(groups[len(word)], word)
	}
	return groups
}
