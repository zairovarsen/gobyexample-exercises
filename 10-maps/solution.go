//go:build solution

package mapsx

import "maps"

func Squares(n int) map[int]int {
	m := make(map[int]int)
	for i := 1; i <= n; i++ {
		m[i] = i * i
	}
	return m
}

func GetOr(m map[string]int, key string, fallback int) int {
	v, prs := m[key]
	if !prs {
		return fallback
	}
	return v
}

func WordCount(words []string) map[string]int {
	counts := make(map[string]int)
	for i := range len(words) {
		counts[words[i]]++
	}
	return counts
}

func Remove(m map[string]int, keys []string) int {
	removed := 0
	for i := range len(keys) {
		if _, prs := m[keys[i]]; prs {
			delete(m, keys[i])
			removed++
		}
	}
	return removed
}

func Reset(m map[string]int) {
	clear(m)
}

func SameWords(a, b []string) bool {
	return maps.Equal(WordCount(a), WordCount(b))
}
