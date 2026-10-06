package timeformattingparsing

import (
	"testing"
	"time"
)

// utc builds a fixed moment in UTC, so that no test depends on the clock or
// the time zone of the machine it runs on.
func utc(year int, month time.Month, day, hour, minute, sec int) time.Time {
	return time.Date(year, month, day, hour, minute, sec, 0, time.UTC)
}

func TestRFC3339(t *testing.T) {
	plus2 := time.FixedZone("", 2*60*60)
	minus530 := time.FixedZone("", -(5*60*60 + 30*60))
	tests := []struct {
		t    time.Time
		want string
	}{
		{utc(2012, time.November, 1, 22, 8, 41), "2012-11-01T22:08:41Z"},
		{time.Date(2012, time.November, 1, 22, 8, 41, 0, plus2), "2012-11-01T22:08:41+02:00"},
		{time.Date(2024, time.February, 29, 6, 0, 5, 0, minus530), "2024-02-29T06:00:05-05:30"},
		{utc(1999, time.December, 31, 23, 59, 59), "1999-12-31T23:59:59Z"},
	}
	for _, tt := range tests {
		if got := RFC3339(tt.t); got != tt.want {
			t.Errorf("RFC3339(%v) = %q, want %q", tt.t, got, tt.want)
		}
	}
}

func TestParseRFC3339(t *testing.T) {
	tests := []struct {
		s    string
		want time.Time
	}{
		{"2012-11-01T22:08:41+00:00", utc(2012, time.November, 1, 22, 8, 41)},
		{"2012-11-01T22:08:41Z", utc(2012, time.November, 1, 22, 8, 41)},
		{"2012-11-01T22:08:41+02:00", utc(2012, time.November, 1, 20, 8, 41)},
		{"2024-02-29T00:30:00-05:30", utc(2024, time.February, 29, 6, 0, 0)},
	}
	for _, tt := range tests {
		got, err := ParseRFC3339(tt.s)
		if err != nil {
			t.Errorf("ParseRFC3339(%q) returned error %v, want nil", tt.s, err)
		}
		if !got.Equal(tt.want) {
			t.Errorf("ParseRFC3339(%q) = %v, want the instant %v", tt.s, got, tt.want)
		}
	}

	for _, s := range []string{"2012-11-01", "8:41PM", "2012-11-01 22:08:41", "tomorrow", ""} {
		if _, err := ParseRFC3339(s); err == nil {
			t.Errorf("ParseRFC3339(%q) returned a nil error, want an error", s)
		}
	}
}

func TestClock(t *testing.T) {
	tests := []struct {
		t    time.Time
		want string
	}{
		{utc(2012, time.November, 1, 22, 8, 41), "10:08PM"},
		{utc(2012, time.November, 1, 6, 0, 0), "6:00AM"},
		{utc(2012, time.November, 1, 12, 30, 0), "12:30PM"},
		{utc(2012, time.November, 1, 0, 5, 0), "12:05AM"},
		{utc(2012, time.November, 1, 13, 59, 59), "1:59PM"},
	}
	for _, tt := range tests {
		if got := Clock(tt.t); got != tt.want {
			t.Errorf("Clock(%v) = %q, want %q", tt.t, got, tt.want)
		}
	}
}

func TestLongDate(t *testing.T) {
	tests := []struct {
		t    time.Time
		want string
	}{
		{utc(2012, time.November, 1, 22, 8, 41), "Thursday, November 1, 2012"},
		{utc(2024, time.February, 29, 0, 0, 0), "Thursday, February 29, 2024"},
		{utc(2000, time.January, 1, 12, 0, 0), "Saturday, January 1, 2000"},
		{utc(2023, time.July, 9, 12, 0, 0), "Sunday, July 9, 2023"},
	}
	for _, tt := range tests {
		if got := LongDate(tt.t); got != tt.want {
			t.Errorf("LongDate(%v) = %q, want %q", tt.t, got, tt.want)
		}
	}
}

func TestParseDate(t *testing.T) {
	tests := []struct {
		s    string
		want time.Time
	}{
		{"2012-11-01", utc(2012, time.November, 1, 0, 0, 0)},
		{"2024-02-29", utc(2024, time.February, 29, 0, 0, 0)},
		{"1999-12-31", utc(1999, time.December, 31, 0, 0, 0)},
	}
	for _, tt := range tests {
		got, err := ParseDate(tt.s)
		if err != nil {
			t.Errorf("ParseDate(%q) returned error %v, want nil", tt.s, err)
		}
		if !got.Equal(tt.want) {
			t.Errorf("ParseDate(%q) = %v, want %v", tt.s, got, tt.want)
		}
	}

	for _, s := range []string{"01/11/2012", "2012-13-01", "2012-02-30", "2023-02-29", "2012-11-01T22:08:41Z", "12-11-01", ""} {
		if _, err := ParseDate(s); err == nil {
			t.Errorf("ParseDate(%q) returned a nil error, want an error", s)
		}
	}
}

func TestReformat(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"01/11/2012 22:08", "Nov 1, 2012 at 10:08PM"},
		{"25/12/2024 07:30", "Dec 25, 2024 at 7:30AM"},
		{"09/07/2023 00:00", "Jul 9, 2023 at 12:00AM"},
		{"12/01/2000 12:00", "Jan 12, 2000 at 12:00PM"},
	}
	for _, tt := range tests {
		got, err := Reformat(tt.s)
		if err != nil {
			t.Errorf("Reformat(%q) returned error %v, want nil", tt.s, err)
		}
		if got != tt.want {
			t.Errorf("Reformat(%q) = %q, want %q", tt.s, got, tt.want)
		}
	}

	for _, s := range []string{"25/13/2024 07:30", "2012-11-01 22:08", "01/11/2012", "01/11/2012 25:00", ""} {
		got, err := Reformat(s)
		if err == nil {
			t.Errorf("Reformat(%q) returned a nil error, want an error", s)
		}
		if got != "" {
			t.Errorf("Reformat(%q) = %q together with an error, want an empty string", s, got)
		}
	}
}

func TestParseAny(t *testing.T) {
	layouts := []string{time.RFC3339, "2006-01-02", "Jan 2, 2006", "02/01/2006 15:04"}
	tests := []struct {
		s    string
		want time.Time
	}{
		{"2012-11-01T22:08:41Z", utc(2012, time.November, 1, 22, 8, 41)},
		{"2012-11-01", utc(2012, time.November, 1, 0, 0, 0)},
		{"Nov 1, 2012", utc(2012, time.November, 1, 0, 0, 0)},
		{"01/11/2012 22:08", utc(2012, time.November, 1, 22, 8, 0)},
	}
	for _, tt := range tests {
		got, err := ParseAny(tt.s, layouts)
		if err != nil {
			t.Errorf("ParseAny(%q, %q) returned error %v, want nil", tt.s, layouts, err)
		}
		if !got.Equal(tt.want) {
			t.Errorf("ParseAny(%q, %q) = %v, want %v", tt.s, layouts, got, tt.want)
		}
	}

	// The first layout that fits wins: 03/04 is the 3rd of April with the
	// day-first layout in front and the 4th of March the other way round.
	dayFirst := []string{"02/01/2006", "01/02/2006"}
	monthFirst := []string{"01/02/2006", "02/01/2006"}
	if got, _ := ParseAny("03/04/2020", dayFirst); !got.Equal(utc(2020, time.April, 3, 0, 0, 0)) {
		t.Errorf("ParseAny(\"03/04/2020\", %q) = %v, want April 3rd", dayFirst, got)
	}
	if got, _ := ParseAny("03/04/2020", monthFirst); !got.Equal(utc(2020, time.March, 4, 0, 0, 0)) {
		t.Errorf("ParseAny(\"03/04/2020\", %q) = %v, want March 4th", monthFirst, got)
	}
	// 25 cannot be a month, so the second layout has to step in.
	if got, _ := ParseAny("25/12/2020", monthFirst); !got.Equal(utc(2020, time.December, 25, 0, 0, 0)) {
		t.Errorf("ParseAny(\"25/12/2020\", %q) = %v, want December 25th", monthFirst, got)
	}
}

func TestParseAnyErrors(t *testing.T) {
	layouts := []string{time.RFC3339, "2006-01-02"}
	for _, s := range []string{"8:41PM", "next week", ""} {
		got, err := ParseAny(s, layouts)
		if err == nil {
			t.Errorf("ParseAny(%q, %q) returned a nil error, want an error", s, layouts)
		}
		if !got.IsZero() {
			t.Errorf("ParseAny(%q, %q) = %v together with an error, want the zero time", s, layouts, got)
		}
	}
	if _, err := ParseAny("2012-11-01", nil); err == nil {
		t.Errorf("ParseAny(\"2012-11-01\", no layouts) returned a nil error, want an error")
	}
}
