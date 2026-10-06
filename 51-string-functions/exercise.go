//go:build !solution

// Package stringfunctions contains exercises for https://gobyexample.com/string-functions
//
// Run the tests for this section with:
//
//	go test ./51-string-functions/
package stringfunctions

// Everything in this section is a job for the standard "strings" package.
// Each doc comment names the functions that fit; look them up with
// `go doc strings.HasPrefix` and friends.

// FileKind classifies a file name. The first rule that matches wins:
//
//   - the name starts with "."           -> "hidden"
//   - the name contains "_gen."          -> "generated"
//   - the name ends with "_test.go"      -> "test"
//   - the name ends with ".go"           -> "source"
//   - anything else                      -> "other"
//
// Examples: ".gitignore" -> "hidden", "types_gen.go" -> "generated",
// "main_test.go" -> "test", "main.go" -> "source", "README.md" -> "other".
//
// Use strings.HasPrefix, strings.Contains and strings.HasSuffix.
func FileKind(name string) string {
	// TODO: check the rules in order.
	return ""
}

// Slug turns a title into the lower-case, dash-separated form used in URLs:
// every space becomes "-" and all letters become lower case.
//
// Example: Slug("Go by Example") returns "go-by-example".
//
// Use strings.ToLower and strings.ReplaceAll (or strings.Replace with -1).
func Slug(title string) string {
	// TODO: lower-case the title and replace the spaces.
	return ""
}

// CountWord returns how many times word occurs in text, ignoring upper and
// lower case. Occurrences inside longer words count as well.
//
// Example: CountWord("Go go GO gopher", "go") returns 4.
//
// Use strings.Count; bring both strings to the same case first.
func CountWord(text, word string) int {
	// TODO: count the case-insensitive occurrences.
	return 0
}

// Underline returns the title followed by a newline and a line of "="
// characters that is exactly as long as the title (in bytes).
//
// Example: Underline("Hello") returns "Hello\n=====".
//
// Use strings.Repeat.
func Underline(title string) string {
	// TODO: build the two lines.
	return ""
}

// Domain returns the part of an e-mail address after the first "@", e.g.
// Domain("gopher@example.com") returns "example.com". When address contains
// no "@" the result is the empty string.
//
// Use strings.Index, which returns -1 when there is no match, and slice the
// string.
func Domain(address string) string {
	// TODO: find the "@" and return what follows it.
	return ""
}

// CleanList tidies up a comma-separated list typed by a human: it splits
// list at the commas, removes the white space around every item, converts
// the items to lower case, drops the items that are empty and joins the rest
// with a single "," again.
//
// Example: CleanList("  Go, rust ,, ZIG ") returns "go,rust,zig".
//
// Use strings.Split, strings.TrimSpace, strings.ToLower and strings.Join.
func CleanList(list string) string {
	// TODO: split, clean every item, join.
	return ""
}
