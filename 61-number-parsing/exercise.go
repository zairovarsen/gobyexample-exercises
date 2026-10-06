//go:build !solution

// Package numberparsing contains exercises for https://gobyexample.com/number-parsing
//
// Run the tests for this section with:
//
//	go test ./61-number-parsing/
package numberparsing

// All exercises use the standard "strconv" package. Its parse functions
// return an error for input that is not a number, or that does not fit into
// the requested size. Unless an exercise says otherwise, hand that error on
// to the caller as it is.

// ParsePrice parses a decimal number such as "1.234" or "-0.5" into a
// float64.
//
// Use strconv.ParseFloat; its second argument is the precision in bits (64).
func ParsePrice(s string) (float64, error) {
	// TODO: parse s as a 64-bit float.
	return 0, nil
}

// ParseCount parses a plain base-10 integer such as "135" or "-7" into an
// int.
//
// Use strconv.Atoi, the convenience function for exactly this.
func ParseCount(s string) (int, error) {
	// TODO: parse s as an int.
	return 0, nil
}

// ParseAnyBase parses an integer whose base is given by its prefix, the way
// Go source code writes numbers: "0x1c8" is hexadecimal (456), "0b101"
// binary (5), "0o17" octal (15), and without a prefix the number is decimal
// ("123"). A leading minus sign is allowed.
//
// Use strconv.ParseInt(s, base, bitSize) with base 0, which means "infer
// the base from the prefix", and bitSize 64.
func ParseAnyBase(s string) (int64, error) {
	// TODO: parse s with an inferred base.
	return 0, nil
}

// ParseBinary parses a string of binary digits without any prefix, such as
// "1011" (11), into an int64.
//
// This time give strconv.ParseInt the base explicitly.
func ParseBinary(s string) (int64, error) {
	// TODO: parse s in base 2.
	return 0, nil
}

// ParseByte parses a decimal number from 0 to 255 into a uint8. Negative
// numbers and numbers above 255 are errors - and you do not have to check
// the range yourself.
//
// Use strconv.ParseUint with base 10 and a bitSize of 8: it then reports
// values that do not fit into 8 bits with an error (one that wraps
// strconv.ErrRange). The result is still a uint64 and needs a conversion.
func ParseByte(s string) (uint8, error) {
	// TODO: parse s as an unsigned 8-bit number.
	return 0, nil
}

// SumFields parses every field as a floating point number and returns the
// sum, e.g. SumFields([]string{"1.5", "2", "-0.25"}) returns 3.25. Without
// fields the sum is 0.
//
// When a field cannot be parsed, SumFields returns 0 and an error that says
// which field was bad and wraps the error from strconv, built like
// fmt.Errorf("field %d: %w", i, err) where i is the index of the field.
//
// Example: SumFields([]string{"1", "x", "3"}) returns the error
// `field 1: strconv.ParseFloat: parsing "x": invalid syntax`.
func SumFields(fields []string) (float64, error) {
	// TODO: parse and add the fields; wrap the first error with its index.
	return 0, nil
}
