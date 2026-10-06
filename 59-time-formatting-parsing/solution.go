//go:build solution

package timeformattingparsing

import (
	"fmt"
	"time"
)

func RFC3339(t time.Time) string {
	return t.Format(time.RFC3339)
}

func ParseRFC3339(s string) (time.Time, error) {
	return time.Parse(time.RFC3339, s)
}

func Clock(t time.Time) string {
	return t.Format("3:04PM")
}

func LongDate(t time.Time) string {
	return t.Format("Monday, January 2, 2006")
}

func ParseDate(s string) (time.Time, error) {
	return time.Parse("2006-01-02", s)
}

func Reformat(s string) (string, error) {
	t, err := time.Parse("02/01/2006 15:04", s)
	if err != nil {
		return "", err
	}
	return t.Format("Jan 2, 2006 at 3:04PM"), nil
}

func ParseAny(s string, layouts []string) (time.Time, error) {
	for _, layout := range layouts {
		if t, err := time.Parse(layout, s); err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("%q matches none of the %d layouts", s, len(layouts))
}
