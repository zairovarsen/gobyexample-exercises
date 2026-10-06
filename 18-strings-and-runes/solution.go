//go:build solution

package stringsandrunes

import (
	"fmt"
	"unicode/utf8"
)

func Lengths(s string) (int, int) {
	return len(s), utf8.RuneCountInString(s)
}

func HexBytes(s string) string {
	out := ""
	for i := 0; i < len(s); i++ {
		if i > 0 {
			out += " "
		}
		out += fmt.Sprintf("%02x", s[i])
	}
	return out
}

func CountRune(s string, target rune) int {
	count := 0
	for _, r := range s {
		if r == target {
			count++
		}
	}
	return count
}

func FirstRune(s string) (rune, int) {
	return utf8.DecodeRuneInString(s)
}

func Reverse(s string) string {
	out := ""
	for _, r := range s {
		out = string(r) + out
	}
	return out
}

func Truncate(s string, n int) string {
	if n <= 0 {
		return ""
	}
	seen := 0
	for i := range s {
		// i is the byte offset of rune number `seen`.
		if seen == n {
			return s[:i]
		}
		seen++
	}
	return s
}
