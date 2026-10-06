package constants

import (
	"math"
	"testing"
)

func TestLanguage(t *testing.T) {
	if Language != "Go" {
		t.Errorf("Language = %q, want %q", Language, "Go")
	}
}

func TestTimeout(t *testing.T) {
	if Timeout != 30 {
		t.Errorf("Timeout = %v, want 30", Timeout)
	}
	// An untyped integer constant becomes an int when it is stored in a
	// variable; only a constant declared with a type arrives as int64.
	if _, ok := any(Timeout).(int64); !ok {
		t.Errorf("Timeout has type %T, want int64 - declare it as `const Timeout int64 = ...`", Timeout)
	}
}

func TestSecondsPer(t *testing.T) {
	if SecondsPerHour != 3600 {
		t.Errorf("SecondsPerHour = %v, want 3600", SecondsPerHour)
	}
	if SecondsPerDay != 86400 {
		t.Errorf("SecondsPerDay = %v, want 86400", SecondsPerDay)
	}
}

func TestZebisPerYobi(t *testing.T) {
	if ZebisPerYobi != 1024 {
		// float64 because a wrong answer might not fit into an int.
		t.Errorf("ZebisPerYobi = %v, want 1024", float64(ZebisPerYobi))
	}
}

func TestCircleArea(t *testing.T) {
	tests := []struct {
		radius float64
		want   float64
	}{
		{1, 3.141592653589793},
		{2, 12.566370614359172},
		{0.5, 0.7853981633974483},
		{0, 0},
	}
	for _, tt := range tests {
		if got := CircleArea(tt.radius); math.Abs(got-tt.want) > 1e-9 {
			t.Errorf("CircleArea(%v) = %v, want %v", tt.radius, got, tt.want)
		}
	}
}

func TestKilogramsAndGrams(t *testing.T) {
	kgTests := []struct {
		grams float64
		want  float64
	}{
		{2500, 2.5},
		{1000, 1},
		{1, 0.001},
	}
	for _, tt := range kgTests {
		if got := Kilograms(tt.grams); math.Abs(got-tt.want) > 1e-9 {
			t.Errorf("Kilograms(%v) = %v, want %v", tt.grams, got, tt.want)
		}
	}

	gramTests := []struct {
		kilograms int
		want      int
	}{
		{3, 3000},
		{1, 1000},
		{0, 0},
	}
	for _, tt := range gramTests {
		if got := Grams(tt.kilograms); got != tt.want {
			t.Errorf("Grams(%d) = %d, want %d", tt.kilograms, got, tt.want)
		}
	}
}

func TestTotalTimeout(t *testing.T) {
	tests := []struct {
		retries int
		want    int
	}{
		{3, 90},
		{1, 30},
		{10, 300},
		{0, 0},
	}
	for _, tt := range tests {
		if got := TotalTimeout(tt.retries); got != tt.want {
			t.Errorf("TotalTimeout(%d) = %d, want %d", tt.retries, got, tt.want)
		}
	}
}
