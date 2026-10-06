package timex

import (
	"math"
	"testing"
	"time"
)

// utc builds a fixed moment in UTC, so that no test depends on the clock or
// the time zone of the machine it runs on.
func utc(year int, month time.Month, day, hour, minute, sec int) time.Time {
	return time.Date(year, month, day, hour, minute, sec, 0, time.UTC)
}

func TestNoon(t *testing.T) {
	tests := []struct {
		year  int
		month time.Month
		day   int
	}{
		{2009, time.November, 17},
		{2024, time.February, 29},
		{1999, time.December, 31},
	}
	for _, tt := range tests {
		got := Noon(tt.year, tt.month, tt.day)
		want := utc(tt.year, tt.month, tt.day, 12, 0, 0)
		if !got.Equal(want) {
			t.Errorf("Noon(%d, %v, %d) = %v, want %v", tt.year, tt.month, tt.day, got, want)
		}
		if got.Location() != time.UTC {
			t.Errorf("Noon(%d, %v, %d) is in location %v, want UTC", tt.year, tt.month, tt.day, got.Location())
		}
	}
}

func TestDescribe(t *testing.T) {
	tests := []struct {
		t    time.Time
		want string
	}{
		{time.Date(2009, time.November, 17, 20, 34, 58, 651387237, time.UTC), "Tuesday 17 November 2009 20:34"},
		{utc(2024, time.February, 29, 9, 5, 0), "Thursday 29 February 2024 09:05"},
		{utc(2000, time.January, 1, 0, 0, 0), "Saturday 1 January 2000 00:00"},
		{utc(2023, time.July, 9, 23, 59, 59), "Sunday 9 July 2023 23:59"},
	}
	for _, tt := range tests {
		if got := Describe(tt.t); got != tt.want {
			t.Errorf("Describe(%v) = %q, want %q", tt.t, got, tt.want)
		}
	}
}

func TestOrder(t *testing.T) {
	early := utc(2009, time.November, 17, 20, 34, 58)
	late := utc(2009, time.November, 17, 20, 34, 59)
	// The same instant as early, written down in a zone two hours ahead.
	earlyElsewhere := early.In(time.FixedZone("UTC+2", 2*60*60))

	tests := []struct {
		name string
		a, b time.Time
		want string
	}{
		{"early, late", early, late, "before"},
		{"late, early", late, early, "after"},
		{"early, early", early, early, "equal"},
		{"early, the same instant in another zone", early, earlyElsewhere, "equal"},
		{"the same instant in another zone, late", earlyElsewhere, late, "before"},
		{"a day in 2030, a day in 1970", utc(2030, 1, 1, 0, 0, 0), utc(1970, 1, 1, 0, 0, 0), "after"},
	}
	for _, tt := range tests {
		if got := Order(tt.a, tt.b); got != tt.want {
			t.Errorf("Order(%s) = %q, want %q\n  a = %v\n  b = %v", tt.name, got, tt.want, tt.a, tt.b)
		}
	}
}

func TestHoursBetween(t *testing.T) {
	tests := []struct {
		start, end time.Time
		want       float64
	}{
		{utc(2024, 3, 1, 9, 0, 0), utc(2024, 3, 1, 17, 30, 0), 8.5},
		{utc(2024, 3, 1, 17, 30, 0), utc(2024, 3, 1, 9, 0, 0), -8.5},
		{utc(2024, 2, 28, 12, 0, 0), utc(2024, 3, 1, 12, 0, 0), 48},
		{utc(2023, 2, 28, 12, 0, 0), utc(2023, 3, 1, 12, 0, 0), 24},
		{utc(2024, 3, 1, 9, 0, 0), utc(2024, 3, 1, 9, 0, 36), 0.01},
	}
	for _, tt := range tests {
		if got := HoursBetween(tt.start, tt.end); math.Abs(got-tt.want) > 1e-9 {
			t.Errorf("HoursBetween(%v, %v) = %v, want %v", tt.start, tt.end, got, tt.want)
		}
	}
}

func TestShift(t *testing.T) {
	tests := []struct {
		t           time.Time
		d           time.Duration
		wantLater   time.Time
		wantEarlier time.Time
	}{
		{
			utc(2024, 3, 1, 12, 0, 0), 90 * time.Minute,
			utc(2024, 3, 1, 13, 30, 0), utc(2024, 3, 1, 10, 30, 0),
		},
		{
			utc(2024, 3, 1, 0, 0, 0), 24 * time.Hour,
			utc(2024, 3, 2, 0, 0, 0), utc(2024, 2, 29, 0, 0, 0),
		},
		{
			utc(2023, 12, 31, 23, 59, 59), 2 * time.Second,
			utc(2024, 1, 1, 0, 0, 1), utc(2023, 12, 31, 23, 59, 57),
		},
	}
	for _, tt := range tests {
		later, earlier := Shift(tt.t, tt.d)
		if !later.Equal(tt.wantLater) {
			t.Errorf("Shift(%v, %v): later = %v, want %v", tt.t, tt.d, later, tt.wantLater)
		}
		if !earlier.Equal(tt.wantEarlier) {
			t.Errorf("Shift(%v, %v): earlier = %v, want %v", tt.t, tt.d, earlier, tt.wantEarlier)
		}
	}
}

func TestAge(t *testing.T) {
	born := utc(1990, time.June, 15, 8, 30, 0)
	leapling := utc(2000, time.February, 29, 0, 0, 0)
	tests := []struct {
		born, now time.Time
		want      int
	}{
		{born, utc(2020, time.June, 14, 23, 59, 59), 29},
		{born, utc(2020, time.June, 15, 0, 0, 0), 30},
		{born, utc(2020, time.June, 16, 12, 0, 0), 30},
		{born, utc(2021, time.January, 1, 0, 0, 0), 30},
		{born, utc(2020, time.May, 31, 0, 0, 0), 29},
		{born, utc(2020, time.July, 1, 0, 0, 0), 30},
		{born, utc(1991, time.June, 15, 0, 0, 0), 1},
		{leapling, utc(2023, time.February, 28, 12, 0, 0), 22},
		{leapling, utc(2023, time.March, 1, 12, 0, 0), 23},
		{leapling, utc(2024, time.February, 29, 12, 0, 0), 24},
	}
	for _, tt := range tests {
		if got := Age(tt.born, tt.now); got != tt.want {
			t.Errorf("Age(born %s, now %s) = %d, want %d",
				tt.born.Format("2006-01-02"), tt.now.Format("2006-01-02"), got, tt.want)
		}
	}
}
