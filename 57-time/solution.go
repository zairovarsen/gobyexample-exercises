//go:build solution

package timex

import (
	"fmt"
	"time"
)

func Noon(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 12, 0, 0, 0, time.UTC)
}

func Describe(t time.Time) string {
	return fmt.Sprintf("%s %d %s %d %02d:%02d",
		t.Weekday(), t.Day(), t.Month(), t.Year(), t.Hour(), t.Minute())
}

func Order(a, b time.Time) string {
	switch {
	case a.Before(b):
		return "before"
	case a.After(b):
		return "after"
	}
	return "equal"
}

func HoursBetween(start, end time.Time) float64 {
	return end.Sub(start).Hours()
}

func Shift(t time.Time, d time.Duration) (later, earlier time.Time) {
	return t.Add(d), t.Add(-d)
}

func Age(born, now time.Time) int {
	age := now.Year() - born.Year()
	// Not there yet this year?
	if now.Month() < born.Month() || (now.Month() == born.Month() && now.Day() < born.Day()) {
		age--
	}
	return age
}
