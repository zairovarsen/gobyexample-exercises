//go:build solution

package switchcase

import "time"

func NumberWord(n int) string {
	switch n {
	case 1:
		return "one"
	case 2:
		return "two"
	case 3:
		return "three"
	default:
		return "many"
	}
}

func DayKind(day time.Weekday) string {
	switch day {
	case time.Saturday, time.Sunday:
		return "weekend"
	default:
		return "weekday"
	}
}

func TimeOfDay(hour int) string {
	switch {
	case hour < 0 || hour > 23:
		return "invalid"
	case hour < 12:
		return "morning"
	case hour < 18:
		return "afternoon"
	default:
		return "evening"
	}
}

func TypeName(v any) string {
	switch v.(type) {
	case bool:
		return "bool"
	case int:
		return "int"
	case float64:
		return "float64"
	case string:
		return "string"
	default:
		return "unknown"
	}
}

func Measure(v any) int {
	switch t := v.(type) {
	case int:
		return t
	case string:
		return len(t)
	case bool:
		if t {
			return 1
		}
		return 0
	default:
		return -1
	}
}

func DaysInMonth(month, year int) int {
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	case 4, 6, 9, 11:
		return 30
	case 2:
		if year%4 == 0 && year%100 != 0 || year%400 == 0 {
			return 29
		}
		return 28
	default:
		return 0
	}
}
