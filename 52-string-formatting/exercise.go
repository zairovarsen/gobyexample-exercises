//go:build !solution

// Package stringformatting contains exercises for https://gobyexample.com/string-formatting
//
// Run the tests for this section with:
//
//	go test ./52-string-formatting/
package stringformatting

import "io"

// Every function in this section is one call to fmt.Sprintf (or fmt.Fprintf)
// with the right format string. The "verbs" are described at
// https://pkg.go.dev/fmt and shown on the gobyexample page.

// Point is a small struct to try the general-purpose verbs on.
type Point struct {
	X, Y int
}

// Item is one line of a receipt, see Receipt.
type Item struct {
	Name  string
	Qty   int
	Price float64 // price of a single piece
}

// DescribePoint formats p with the four general verbs, labelled and
// separated by single spaces:
//
//	v=<%v> +v=<%+v> #v=<%#v> T=<%T>
//
// %v prints the plain value, %+v adds the field names, %#v prints Go syntax
// that would create the value and %T prints its type.
//
// Example: DescribePoint(Point{1, 2}) returns
//
//	v={1 2} +v={X:1 Y:2} #v=stringformatting.Point{X:1, Y:2} T=stringformatting.Point
func DescribePoint(p Point) string {
	// TODO: one Sprintf with four verbs.
	return ""
}

// DescribeInt formats n in decimal (%d), binary (%b), lower-case hexadecimal
// (%x) and as the character with that Unicode code point (%c).
//
// Example: DescribeInt(65) returns "dec=65 bin=1000001 hex=41 char=A".
func DescribeInt(n int) string {
	// TODO: one Sprintf with four verbs.
	return ""
}

// DescribeFloat formats f as a plain decimal (%f), as a decimal with exactly
// two digits after the point, and in scientific notation with a lower-case
// (%e) and an upper-case (%E) exponent.
//
// Example: DescribeFloat(78.9) returns
// "f=78.900000 2f=78.90 e=7.890000e+01 E=7.890000E+01".
func DescribeFloat(f float64) string {
	// TODO: one Sprintf with four verbs.
	return ""
}

// DescribeString formats s as it is (%s), double-quoted like in Go source
// (%q) and as the hexadecimal dump of its bytes (%x).
//
// Example: DescribeString("hi") returns `s=hi q="hi" x=6869`.
func DescribeString(s string) string {
	// TODO: one Sprintf with three verbs.
	return ""
}

// TableRow formats one row of a text table with three columns between "|"
// characters:
//
//   - name, left-aligned and padded with spaces to 10 characters
//   - qty, right-aligned in 5 characters
//   - price with 2 digits after the decimal point, right-aligned in 8
//     characters
//
// Example: TableRow("apple", 3, 1.5) returns "|apple     |    3|    1.50|".
//
// Hint: a number between % and the verb is the width, a "-" flag aligns to
// the left, and ".2" sets the precision.
func TableRow(name string, qty int, price float64) string {
	// TODO: one Sprintf with widths.
	return ""
}

// Receipt writes a receipt for items to w - use fmt.Fprintf, which formats
// like Sprintf but writes to an io.Writer (a real program would pass
// os.Stdout). There is one line per item and a final line with the total:
//
//	001 apple        3     4.50
//	002 watermelon  12    30.00
//	TOTAL                 34.50
//
// An item line consists of, separated by single spaces: the line number
// (starting at 1) padded with zeros to 3 digits, the name left-aligned in 10
// characters, the quantity right-aligned in 3 characters, and the line total
// (Qty * Price) with 2 decimals right-aligned in 8 characters.
//
// The last line is the word TOTAL left-aligned in 18 characters, a space,
// and the sum of all line totals formatted like the line totals. Every line,
// including the last one, ends with "\n". With no items only the TOTAL line
// is written.
func Receipt(w io.Writer, items []Item) {
	// TODO: Fprintf one line per item, then the total.
}
