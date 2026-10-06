//go:build solution

package slicesx

import "slices"

func Countdown(n int) []int {
	var s []int
	for i := n; i >= 1; i-- {
		s = append(s, i)
	}
	return s
}

func Middle(s []int) []int {
	if len(s) < 2 {
		return []int{}
	}
	return s[1 : len(s)-1]
}

func Clone(s []int) []int {
	c := make([]int, len(s))
	copy(c, s)
	return c
}

func HasPrefix(s, prefix []string) bool {
	if len(prefix) > len(s) {
		return false
	}
	return slices.Equal(s[:len(prefix)], prefix)
}

func Triangle(n int) [][]int {
	rows := make([][]int, n)
	for i := range n {
		innerLen := i + 1
		rows[i] = make([]int, innerLen)
		for j := range innerLen {
			rows[i][j] = i + j
		}
	}
	return rows
}

func Remove(s []int, i int) []int {
	out := make([]int, len(s)-1)
	copy(out, s[:i])
	// out[i:] is the still-empty rest of out.
	copy(out[i:], s[i+1:])
	return out
}
