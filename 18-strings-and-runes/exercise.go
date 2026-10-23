//go:build !solution

// Package stringsandrunes contains exercises for https://gobyexample.com/strings-and-runes
//
// Run the tests for this section with:
//
//	go test ./18-strings-and-runes/
package stringsandrunes

import (
	"fmt"
	"unicode/utf8"
)

// Lengths returns the two lengths of s: first the number of bytes, then the
// number of runes (characters). For "hello" both are 5, for "héllo" they are
// 6 and 5, and the Thai word "สวัสดี" has 18 bytes but 6 runes.
//
// Hint: len counts bytes; utf8.RuneCountInString from the "unicode/utf8"
// package counts runes.
func Lengths(s string) (int, int) {
	// TODO: return the byte count and the rune count.
	return len(s), utf8.RuneCountInString(s)
}

// HexBytes shows the raw bytes that make up s: every byte as two lower-case
// hex digits, separated by single spaces.
//
// Example: HexBytes("hé") returns "68 c3 a9" - one byte for h, two for é.
// HexBytes("") returns "".
//
// Indexing a string (s[i]) gives a byte, not a character.
//
// Hint: fmt.Sprintf("%02x", b) formats one byte as two hex digits and
// returns the string instead of printing it.
func HexBytes(s string) string {
	// TODO: loop over the byte indexes 0 .. len(s)-1 and format every byte.
	str := ""

	for i := range len(s) {
		if i != 0 {
			str += " "
		}
		str += fmt.Sprintf("%02x", s[i])
	}
	return str
}

// CountRune returns how often the rune target occurs in s:
// CountRune("สวัสดี", 'ส') returns 2.
//
// A value in single quotes is a rune literal. range over a string yields
// runes that can be compared with it using ==.
func CountRune(s string, target rune) int {
	// TODO: range over s and compare every rune with target.
	count := 0
	for _, runeValue := range s {
		if runeValue == target {
			count += 1
		}	
    }
	return count
}

// FirstRune returns the first rune of s and its width, that is the number of
// bytes it occupies: FirstRune("สวัสดี") returns 'ส', 3 and FirstRune("hi")
// returns 'h', 1.
//
// Hint: this is what utf8.DecodeRuneInString does. (For an empty string it
// returns utf8.RuneError and 0, and so should FirstRune.)
func FirstRune(s string) (rune, int) {
	// TODO: decode the first rune.
	return utf8.DecodeRuneInString(s[0:])
}

// Reverse returns s with its characters in opposite order:
// Reverse("héllo") returns "olléh". It must reverse the runes - reversing
// the bytes would tear multi-byte characters such as é apart.
//
// Hint: string(r) turns a rune into a string that can be concatenated.
func Reverse(s string) string {
	// TODO: range over the runes and build the result back to front.
	str := ""
	for _, v := range s {
		str = string(v) + str
	}
	return str
}

// Truncate shortens s to at most n characters (runes, not bytes) and never
// cuts a multi-byte character in half: Truncate("héllo", 2) returns "hé".
// A string with n runes or fewer is returned unchanged, and for n <= 0 the
// result is "".
//
// Hint: the index that range gives for a rune is the byte position at which
// it starts - exactly what the slice expression s[:i] needs.
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
