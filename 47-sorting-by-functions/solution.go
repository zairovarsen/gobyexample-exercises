//go:build solution

package sortingbyfunctions

import (
	"cmp"
	"slices"
)

type Person struct {
	Name string
	Age  int
}

type WordCount struct {
	Word  string
	Count int
}

func SortByLength(words []string) {
	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}
	slices.SortFunc(words, lenCmp)
}

func SortDescending(nums []int) {
	slices.SortFunc(nums, func(a, b int) int {
		return cmp.Compare(b, a)
	})
}

func SortByAge(people []Person) {
	slices.SortFunc(people, func(a, b Person) int {
		return cmp.Compare(a.Age, b.Age)
	})
}

func SortByAgeThenName(people []Person) {
	slices.SortFunc(people, func(a, b Person) int {
		return cmp.Or(
			cmp.Compare(a.Age, b.Age),
			cmp.Compare(a.Name, b.Name),
		)
	})
}

func WordFrequencies(words []string) []WordCount {
	counts := make(map[string]int)
	for _, w := range words {
		counts[w]++
	}
	result := make([]WordCount, 0, len(counts))
	for w, n := range counts {
		result = append(result, WordCount{Word: w, Count: n})
	}
	slices.SortFunc(result, func(a, b WordCount) int {
		return cmp.Or(
			cmp.Compare(b.Count, a.Count),
			cmp.Compare(a.Word, b.Word),
		)
	})
	return result
}
