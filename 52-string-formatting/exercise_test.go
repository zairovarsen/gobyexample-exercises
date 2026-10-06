package stringformatting

import (
	"bytes"
	"testing"
)

func TestDescribePoint(t *testing.T) {
	tests := []struct {
		p    Point
		want string
	}{
		{Point{1, 2}, "v={1 2} +v={X:1 Y:2} #v=stringformatting.Point{X:1, Y:2} T=stringformatting.Point"},
		{Point{-7, 0}, "v={-7 0} +v={X:-7 Y:0} #v=stringformatting.Point{X:-7, Y:0} T=stringformatting.Point"},
	}
	for _, tt := range tests {
		if got := DescribePoint(tt.p); got != tt.want {
			t.Errorf("DescribePoint(%#v) =\n  %q, want\n  %q", tt.p, got, tt.want)
		}
	}
}

func TestDescribeInt(t *testing.T) {
	tests := []struct {
		n    int
		want string
	}{
		{65, "dec=65 bin=1000001 hex=41 char=A"},
		{97, "dec=97 bin=1100001 hex=61 char=a"},
		{955, "dec=955 bin=1110111011 hex=3bb char=λ"},
		{456, "dec=456 bin=111001000 hex=1c8 char=ǈ"},
	}
	for _, tt := range tests {
		if got := DescribeInt(tt.n); got != tt.want {
			t.Errorf("DescribeInt(%d) = %q, want %q", tt.n, got, tt.want)
		}
	}
}

func TestDescribeFloat(t *testing.T) {
	tests := []struct {
		f    float64
		want string
	}{
		{78.9, "f=78.900000 2f=78.90 e=7.890000e+01 E=7.890000E+01"},
		{123400000.0, "f=123400000.000000 2f=123400000.00 e=1.234000e+08 E=1.234000E+08"},
		{0.000123, "f=0.000123 2f=0.00 e=1.230000e-04 E=1.230000E-04"},
	}
	for _, tt := range tests {
		if got := DescribeFloat(tt.f); got != tt.want {
			t.Errorf("DescribeFloat(%v) = %q, want %q", tt.f, got, tt.want)
		}
	}
}

func TestDescribeString(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"hi", `s=hi q="hi" x=6869`},
		{"hex this", `s=hex this q="hex this" x=6865782074686973`},
		{`say "go"`, `s=say "go" q="say \"go\"" x=7361792022676f22`},
	}
	for _, tt := range tests {
		if got := DescribeString(tt.s); got != tt.want {
			t.Errorf("DescribeString(%q) = %q, want %q", tt.s, got, tt.want)
		}
	}
}

func TestTableRow(t *testing.T) {
	tests := []struct {
		name  string
		qty   int
		price float64
		want  string
	}{
		{"apple", 3, 1.5, "|apple     |    3|    1.50|"},
		{"watermelon", 12345, 1234.567, "|watermelon|12345| 1234.57|"},
		{"fig", -7, 0, "|fig       |   -7|    0.00|"},
	}
	for _, tt := range tests {
		if got := TableRow(tt.name, tt.qty, tt.price); got != tt.want {
			t.Errorf("TableRow(%q, %d, %v) =\n  %q, want\n  %q", tt.name, tt.qty, tt.price, got, tt.want)
		}
	}
}

func TestReceipt(t *testing.T) {
	items := []Item{
		{"apple", 3, 1.5},
		{"watermelon", 12, 2.5},
		{"fig", 100, 0.25},
	}
	want := "" +
		"001 apple        3     4.50\n" +
		"002 watermelon  12    30.00\n" +
		"003 fig        100    25.00\n" +
		"TOTAL                 59.50\n"

	var buf bytes.Buffer
	Receipt(&buf, items)
	if got := buf.String(); got != want {
		t.Errorf("Receipt(w, %v) wrote\n%s\nwant\n%s", items, got, want)
	}

	buf.Reset()
	Receipt(&buf, nil)
	if got, want := buf.String(), "TOTAL                  0.00\n"; got != want {
		t.Errorf("Receipt(w, nil) wrote %q, want %q", got, want)
	}
}
