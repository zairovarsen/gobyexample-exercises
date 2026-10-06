//go:build !solution

// Package regularexpressions contains exercises for https://gobyexample.com/regular-expressions
//
// Run the tests for this section with:
//
//	go test ./54-regular-expressions/
package regularexpressions

// All exercises use the standard "regexp" package. Write patterns in
// backquotes (`\d+`) so that the backslashes need no escaping. The syntax is
// described at https://pkg.go.dev/regexp/syntax.

// IsHexColor reports whether s is a CSS colour of the form "#" followed by
// exactly six hexadecimal digits (0-9, a-f, A-F) and nothing else.
//
// Examples: "#1a2B3c" -> true, "#12345" -> false, "x#123456" -> false.
//
// Hint: regexp.MatchString tests a pattern directly. A pattern matches
// anywhere in the string unless you anchor it with ^ and $.
func IsHexColor(s string) bool {
	// TODO: match s against a pattern.
	return false
}

// FindFirst returns the first (leftmost) match of pattern in text, or the
// empty string when there is no match. The pattern comes from the caller and
// may be invalid; in that case FindFirst returns the error from the regexp
// package.
//
// Examples:
//
//	FindFirst(`p([a-z]+)ch`, "peach punch") -> "peach", nil
//	FindFirst(`[0-9]+`, "no digits")        -> "", nil
//	FindFirst(`a(b`, "abc")                 -> "", an error
//
// Use regexp.Compile (not MustCompile, which panics on a bad pattern) and
// the FindString method.
func FindFirst(pattern, text string) (string, error) {
	// TODO: compile the pattern, then search.
	return "", nil
}

// ParseDate splits a date written as YYYY-MM-DD (4, 2 and 2 digits) into
// its three parts, which are returned as strings just as they are written.
// The whole string must be a date, otherwise ok is false and the parts are
// empty.
//
// Examples:
//
//	ParseDate("2024-03-15")    -> "2024", "03", "15", true
//	ParseDate("2024-3-15")     -> "", "", "", false
//	ParseDate("on 2024-03-15") -> "", "", "", false
//
// Use a pattern with three (parenthesised) groups and FindStringSubmatch,
// which returns the whole match followed by what each group matched, or nil.
// A pattern that never changes is best compiled once, in a package-level
// variable, with regexp.MustCompile.
func ParseDate(s string) (year, month, day string, ok bool) {
	// TODO: match with groups and return the submatches.
	return "", "", "", false
}

// Hashtags returns all hashtags in text, in order of appearance. A hashtag
// is a "#" directly followed by one or more word characters (letters, digits
// or underscore, \w in a pattern). Without any hashtag the result is empty.
//
// Example: Hashtags("learning #go and #regexp, day #100_days!") returns
// [#go #regexp #100_days].
//
// Use FindAllString; -1 as the second argument means "all matches".
func Hashtags(text string) []string {
	// TODO: find all matches.
	return nil
}

// MaskDigits returns s with every digit replaced by "#".
//
// Example: MaskDigits("card 1234-5678, exp 09/27") returns
// "card ####-####, exp ##/##".
//
// Use ReplaceAllString.
func MaskDigits(s string) string {
	// TODO: replace all matches.
	return ""
}

// ExpandVars replaces every placeholder of the form ${name} in text with
// vars["name"]. A name consists of one or more word characters (\w).
// Placeholders whose name is not a key of vars stay in the text unchanged.
//
// Example: with vars = {"user": "gopher", "lang": "Go"}
//
//	ExpandVars("Hi ${user}, welcome to ${lang}! ${unknown}", vars)
//
// returns "Hi gopher, welcome to Go! ${unknown}".
//
// Use ReplaceAllStringFunc, which calls a function with every match and
// replaces the match with what the function returns. Remember that $, { and
// } have a special meaning in patterns and need a backslash.
func ExpandVars(text string, vars map[string]string) string {
	// TODO: replace every placeholder using a function.
	return ""
}
