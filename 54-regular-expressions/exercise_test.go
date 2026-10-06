package regularexpressions

import (
	"slices"
	"testing"
)

func TestIsHexColor(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"#1a2B3c", true},
		{"#000000", true},
		{"#FFFFFF", true},
		{"#abcdef", true},
		{"#12345", false},
		{"#1234567", false},
		{"123456", false},
		{"#12345g", false},
		{"x#123456", false},
		{"#123456 ", false},
		{"", false},
	}
	for _, tt := range tests {
		if got := IsHexColor(tt.s); got != tt.want {
			t.Errorf("IsHexColor(%q) = %t, want %t", tt.s, got, tt.want)
		}
	}
}

func TestFindFirst(t *testing.T) {
	tests := []struct {
		pattern, text string
		want          string
	}{
		{`p([a-z]+)ch`, "peach punch", "peach"},
		{`p([a-z]+)ch`, "a pinch of salt", "pinch"},
		{`[0-9]+`, "abc 42 and 7", "42"},
		{`[0-9]+`, "no digits", ""},
		{`^\w+`, "first word", "first"},
	}
	for _, tt := range tests {
		got, err := FindFirst(tt.pattern, tt.text)
		if err != nil {
			t.Errorf("FindFirst(%q, %q) returned error %v, want nil", tt.pattern, tt.text, err)
		}
		if got != tt.want {
			t.Errorf("FindFirst(%q, %q) = %q, want %q", tt.pattern, tt.text, got, tt.want)
		}
	}

	for _, pattern := range []string{`a(b`, `[a-`, `*oops`} {
		got, err := FindFirst(pattern, "abc")
		if err == nil {
			t.Errorf("FindFirst(%q, \"abc\") returned a nil error, want an error for the invalid pattern", pattern)
		}
		if got != "" {
			t.Errorf("FindFirst(%q, \"abc\") = %q together with an error, want an empty string", pattern, got)
		}
	}
}

func TestParseDate(t *testing.T) {
	tests := []struct {
		s                string
		year, month, day string
		ok               bool
	}{
		{"2024-03-15", "2024", "03", "15", true},
		{"1999-12-31", "1999", "12", "31", true},
		{"2024-3-15", "", "", "", false},
		{"24-03-15", "", "", "", false},
		{"2024/03/15", "", "", "", false},
		{"on 2024-03-15", "", "", "", false},
		{"2024-03-15T10:00", "", "", "", false},
		{"", "", "", "", false},
	}
	for _, tt := range tests {
		year, month, day, ok := ParseDate(tt.s)
		if year != tt.year || month != tt.month || day != tt.day || ok != tt.ok {
			t.Errorf("ParseDate(%q) = %q, %q, %q, %t; want %q, %q, %q, %t",
				tt.s, year, month, day, ok, tt.year, tt.month, tt.day, tt.ok)
		}
	}
}

func TestHashtags(t *testing.T) {
	tests := []struct {
		text string
		want []string
	}{
		{"learning #go and #regexp, day #100_days!", []string{"#go", "#regexp", "#100_days"}},
		{"#first", []string{"#first"}},
		{"#a#b # c", []string{"#a", "#b"}},
		{"issue #42 is fixed (#43 is not)", []string{"#42", "#43"}},
		{"no tags here, not even # alone", nil},
	}
	for _, tt := range tests {
		if got := Hashtags(tt.text); !slices.Equal(got, tt.want) {
			t.Errorf("Hashtags(%q) = %q, want %q", tt.text, got, tt.want)
		}
	}
}

func TestMaskDigits(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"card 1234-5678, exp 09/27", "card ####-####, exp ##/##"},
		{"007", "###"},
		{"call 555 0100 now", "call ### #### now"},
		{"nothing to hide", "nothing to hide"},
	}
	for _, tt := range tests {
		if got := MaskDigits(tt.s); got != tt.want {
			t.Errorf("MaskDigits(%q) = %q, want %q", tt.s, got, tt.want)
		}
	}
}

func TestExpandVars(t *testing.T) {
	vars := map[string]string{"user": "gopher", "lang": "Go", "empty": "", "dir_1": "/tmp"}
	tests := []struct {
		text string
		want string
	}{
		{"Hi ${user}, welcome to ${lang}! ${unknown}", "Hi gopher, welcome to Go! ${unknown}"},
		{"${user}${user}", "gophergopher"},
		{"[${empty}]", "[]"},
		{"cd ${dir_1}/x", "cd /tmp/x"},
		{"$user {user} ${} ${us er}", "$user {user} ${} ${us er}"},
		{"costs $5", "costs $5"},
	}
	for _, tt := range tests {
		if got := ExpandVars(tt.text, vars); got != tt.want {
			t.Errorf("ExpandVars(%q, %v) = %q, want %q", tt.text, vars, got, tt.want)
		}
	}
}
