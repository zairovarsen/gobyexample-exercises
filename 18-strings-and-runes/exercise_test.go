package stringsandrunes

import (
	"testing"
	"unicode/utf8"
)

func TestLengths(t *testing.T) {
	tests := []struct {
		s                    string
		wantBytes, wantRunes int
	}{
		{"hello", 5, 5},
		{"héllo", 6, 5},
		{"สวัสดี", 18, 6},
		{"😀", 4, 1},
		{"", 0, 0},
	}
	for _, tt := range tests {
		if b, r := Lengths(tt.s); b != tt.wantBytes || r != tt.wantRunes {
			t.Errorf("Lengths(%q) = %d, %d, want %d, %d", tt.s, b, r, tt.wantBytes, tt.wantRunes)
		}
	}
}

func TestHexBytes(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"hé", "68 c3 a9"},
		{"go", "67 6f"},
		{"ส", "e0 b8 aa"},
		{"a\tb", "61 09 62"},
		{"", ""},
	}
	for _, tt := range tests {
		if got := HexBytes(tt.s); got != tt.want {
			t.Errorf("HexBytes(%q) = %q, want %q", tt.s, got, tt.want)
		}
	}
}

func TestCountRune(t *testing.T) {
	tests := []struct {
		s      string
		target rune
		want   int
	}{
		{"สวัสดี", 'ส', 2},
		{"hello", 'l', 2},
		{"héllo é", 'é', 2},
		{"hello", 'z', 0},
		{"", 'a', 0},
	}
	for _, tt := range tests {
		if got := CountRune(tt.s, tt.target); got != tt.want {
			t.Errorf("CountRune(%q, %q) = %d, want %d", tt.s, tt.target, got, tt.want)
		}
	}
}

func TestFirstRune(t *testing.T) {
	tests := []struct {
		s         string
		wantRune  rune
		wantWidth int
	}{
		{"สวัสดี", 'ส', 3},
		{"hi", 'h', 1},
		{"éa", 'é', 2},
		{"😀!", '😀', 4},
		{"", utf8.RuneError, 0},
	}
	for _, tt := range tests {
		if r, width := FirstRune(tt.s); r != tt.wantRune || width != tt.wantWidth {
			t.Errorf("FirstRune(%q) = %q, %d, want %q, %d", tt.s, r, width, tt.wantRune, tt.wantWidth)
		}
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"hello", "olleh"},
		{"héllo", "olléh"},
		{"Go 语言", "言语 oG"},
		{"a", "a"},
		{"", ""},
	}
	for _, tt := range tests {
		if got := Reverse(tt.s); got != tt.want {
			t.Errorf("Reverse(%q) = %q, want %q", tt.s, got, tt.want)
		}
	}
}

func TestTruncate(t *testing.T) {
	tests := []struct {
		s    string
		n    int
		want string
	}{
		{"hello", 3, "hel"},
		{"héllo", 2, "hé"},
		{"สวัสดี", 2, "สว"},
		{"日本語", 1, "日"},
		{"hi", 2, "hi"},
		{"hi", 5, "hi"},
		{"hi", 0, ""},
		{"hi", -1, ""},
		{"", 3, ""},
	}
	for _, tt := range tests {
		if got := Truncate(tt.s, tt.n); got != tt.want {
			t.Errorf("Truncate(%q, %d) = %q, want %q", tt.s, tt.n, got, tt.want)
		}
	}
}
