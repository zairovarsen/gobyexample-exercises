package numberparsing

import (
	"errors"
	"math"
	"strconv"
	"testing"
)

func TestParsePrice(t *testing.T) {
	tests := []struct {
		s    string
		want float64
	}{
		{"1.234", 1.234},
		{"-0.5", -0.5},
		{"42", 42},
		{"1e3", 1000},
		{".25", 0.25},
	}
	for _, tt := range tests {
		got, err := ParsePrice(tt.s)
		if err != nil {
			t.Errorf("ParsePrice(%q) returned error %v, want nil", tt.s, err)
		}
		if got != tt.want {
			t.Errorf("ParsePrice(%q) = %v, want %v", tt.s, got, tt.want)
		}
	}

	for _, s := range []string{"wat", "1,5", "1.2.3", "", "12 "} {
		if _, err := ParsePrice(s); !errors.Is(err, strconv.ErrSyntax) {
			t.Errorf("ParsePrice(%q) returned error %v, want the syntax error from strconv", s, err)
		}
	}
}

func TestParseCount(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"135", 135},
		{"-7", -7},
		{"+12", 12},
		{"007", 7},
	}
	for _, tt := range tests {
		got, err := ParseCount(tt.s)
		if err != nil {
			t.Errorf("ParseCount(%q) returned error %v, want nil", tt.s, err)
		}
		if got != tt.want {
			t.Errorf("ParseCount(%q) = %d, want %d", tt.s, got, tt.want)
		}
	}

	for _, s := range []string{"wat", "1.5", "0x10", "12a", ""} {
		if _, err := ParseCount(s); !errors.Is(err, strconv.ErrSyntax) {
			t.Errorf("ParseCount(%q) returned error %v, want the syntax error from strconv", s, err)
		}
	}
}

func TestParseAnyBase(t *testing.T) {
	tests := []struct {
		s    string
		want int64
	}{
		{"123", 123},
		{"0x1c8", 456},
		{"0X1C8", 456},
		{"-0x10", -16},
		{"0b101", 5},
		{"0o17", 15},
		{"017", 15},
		{"1_000_000", 1000000},
		{"9223372036854775807", math.MaxInt64},
	}
	for _, tt := range tests {
		got, err := ParseAnyBase(tt.s)
		if err != nil {
			t.Errorf("ParseAnyBase(%q) returned error %v, want nil", tt.s, err)
		}
		if got != tt.want {
			t.Errorf("ParseAnyBase(%q) = %d, want %d", tt.s, got, tt.want)
		}
	}

	for _, s := range []string{"0xzz", "1c8", "0b102", "wat", ""} {
		if _, err := ParseAnyBase(s); !errors.Is(err, strconv.ErrSyntax) {
			t.Errorf("ParseAnyBase(%q) returned error %v, want the syntax error from strconv", s, err)
		}
	}
	if _, err := ParseAnyBase("9223372036854775808"); !errors.Is(err, strconv.ErrRange) {
		t.Errorf("ParseAnyBase(\"9223372036854775808\") returned error %v, want the range error from strconv (the number needs more than 64 bits)", err)
	}
}

func TestParseBinary(t *testing.T) {
	tests := []struct {
		s    string
		want int64
	}{
		{"1011", 11},
		{"0", 0},
		{"1", 1},
		{"11111111", 255},
		{"-100", -4},
		{"10", 2},
	}
	for _, tt := range tests {
		got, err := ParseBinary(tt.s)
		if err != nil {
			t.Errorf("ParseBinary(%q) returned error %v, want nil", tt.s, err)
		}
		if got != tt.want {
			t.Errorf("ParseBinary(%q) = %d, want %d", tt.s, got, tt.want)
		}
	}

	for _, s := range []string{"102", "0b11", "two", ""} {
		if _, err := ParseBinary(s); !errors.Is(err, strconv.ErrSyntax) {
			t.Errorf("ParseBinary(%q) returned error %v, want the syntax error from strconv", s, err)
		}
	}
}

func TestParseByte(t *testing.T) {
	tests := []struct {
		s    string
		want uint8
	}{
		{"0", 0},
		{"7", 7},
		{"200", 200},
		{"255", 255},
	}
	for _, tt := range tests {
		got, err := ParseByte(tt.s)
		if err != nil {
			t.Errorf("ParseByte(%q) returned error %v, want nil", tt.s, err)
		}
		if got != tt.want {
			t.Errorf("ParseByte(%q) = %d, want %d", tt.s, got, tt.want)
		}
	}

	for _, s := range []string{"256", "1000", "99999999999999999999999"} {
		if _, err := ParseByte(s); !errors.Is(err, strconv.ErrRange) {
			t.Errorf("ParseByte(%q) returned error %v, want the range error from strconv", s, err)
		}
	}
	for _, s := range []string{"-1", "abc", "1.5", ""} {
		if _, err := ParseByte(s); !errors.Is(err, strconv.ErrSyntax) {
			t.Errorf("ParseByte(%q) returned error %v, want the syntax error from strconv", s, err)
		}
	}
}

func TestSumFields(t *testing.T) {
	tests := []struct {
		fields []string
		want   float64
	}{
		{[]string{"1.5", "2", "-0.25"}, 3.25},
		{[]string{"10"}, 10},
		{[]string{"0.5", "0.25", "0.125", "1e2"}, 100.875},
		{nil, 0},
	}
	for _, tt := range tests {
		got, err := SumFields(tt.fields)
		if err != nil {
			t.Errorf("SumFields(%q) returned error %v, want nil", tt.fields, err)
		}
		if got != tt.want {
			t.Errorf("SumFields(%q) = %v, want %v", tt.fields, got, tt.want)
		}
	}
}

func TestSumFieldsErrors(t *testing.T) {
	tests := []struct {
		fields  []string
		wantMsg string
	}{
		{[]string{"1", "x", "3"}, `field 1: strconv.ParseFloat: parsing "x": invalid syntax`},
		{[]string{"", "2"}, `field 0: strconv.ParseFloat: parsing "": invalid syntax`},
		{[]string{"1", "2", "3", "4,5", "oops"}, `field 3: strconv.ParseFloat: parsing "4,5": invalid syntax`},
	}
	for _, tt := range tests {
		got, err := SumFields(tt.fields)
		if err == nil {
			t.Errorf("SumFields(%q) returned a nil error, want %q", tt.fields, tt.wantMsg)
			continue
		}
		if err.Error() != tt.wantMsg {
			t.Errorf("SumFields(%q) returned the error\n  %q, want\n  %q", tt.fields, err, tt.wantMsg)
		}
		// Wrapping with %w keeps the original error reachable.
		var numErr *strconv.NumError
		if !errors.As(err, &numErr) {
			t.Errorf("SumFields(%q) returned an error that does not wrap the *strconv.NumError (use %%w)", tt.fields)
		}
		if got != 0 {
			t.Errorf("SumFields(%q) = %v together with an error, want 0", tt.fields, got)
		}
	}
}
