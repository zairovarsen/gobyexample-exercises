package switchcase

import (
	"testing"
	"time"
)

func TestNumberWord(t *testing.T) {
	tests := []struct {
		n    int
		want string
	}{
		{1, "one"},
		{2, "two"},
		{3, "three"},
		{4, "many"},
		{0, "many"},
		{-1, "many"},
	}
	for _, tt := range tests {
		if got := NumberWord(tt.n); got != tt.want {
			t.Errorf("NumberWord(%d) = %q, want %q", tt.n, got, tt.want)
		}
	}
}

func TestDayKind(t *testing.T) {
	tests := []struct {
		day  time.Weekday
		want string
	}{
		{time.Monday, "weekday"},
		{time.Tuesday, "weekday"},
		{time.Wednesday, "weekday"},
		{time.Thursday, "weekday"},
		{time.Friday, "weekday"},
		{time.Saturday, "weekend"},
		{time.Sunday, "weekend"},
	}
	for _, tt := range tests {
		if got := DayKind(tt.day); got != tt.want {
			t.Errorf("DayKind(time.%v) = %q, want %q", tt.day, got, tt.want)
		}
	}
}

func TestTimeOfDay(t *testing.T) {
	tests := []struct {
		hour int
		want string
	}{
		{0, "morning"},
		{9, "morning"},
		{11, "morning"},
		{12, "afternoon"},
		{17, "afternoon"},
		{18, "evening"},
		{23, "evening"},
		{24, "invalid"},
		{-1, "invalid"},
	}
	for _, tt := range tests {
		if got := TimeOfDay(tt.hour); got != tt.want {
			t.Errorf("TimeOfDay(%d) = %q, want %q", tt.hour, got, tt.want)
		}
	}
}

func TestTypeName(t *testing.T) {
	tests := []struct {
		v    any
		want string
	}{
		{true, "bool"},
		{1, "int"},
		{2.5, "float64"},
		{"hey", "string"},
		{'x', "unknown"},
		{int64(1), "unknown"},
		{[]int{1}, "unknown"},
		{nil, "unknown"},
	}
	for _, tt := range tests {
		if got := TypeName(tt.v); got != tt.want {
			t.Errorf("TypeName(%#v) = %q, want %q", tt.v, got, tt.want)
		}
	}
}

func TestMeasure(t *testing.T) {
	tests := []struct {
		v    any
		want int
	}{
		{42, 42},
		{-7, -7},
		{"hello", 5},
		{"", 0},
		{true, 1},
		{false, 0},
		{2.5, -1},
		{nil, -1},
	}
	for _, tt := range tests {
		if got := Measure(tt.v); got != tt.want {
			t.Errorf("Measure(%#v) = %d, want %d", tt.v, got, tt.want)
		}
	}
}

func TestDaysInMonth(t *testing.T) {
	tests := []struct {
		month, year int
		want        int
	}{
		{1, 2023, 31},
		{4, 2023, 30},
		{7, 2023, 31},
		{8, 2023, 31},
		{9, 2023, 30},
		{11, 2023, 30},
		{12, 2023, 31},
		{2, 2023, 28},
		{2, 2024, 29},
		{2, 1900, 28},
		{2, 2000, 29},
		{0, 2023, 0},
		{13, 2023, 0},
	}
	for _, tt := range tests {
		if got := DaysInMonth(tt.month, tt.year); got != tt.want {
			t.Errorf("DaysInMonth(%d, %d) = %d, want %d", tt.month, tt.year, got, tt.want)
		}
	}
}
