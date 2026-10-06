//go:build solution

package constants

import "math"

const Language = "Go"

const Timeout int64 = 30

const (
	SecondsPerMinute = 60
	MinutesPerHour   = 60
	HoursPerDay      = 24
)

const (
	SecondsPerHour = SecondsPerMinute * MinutesPerHour
	SecondsPerDay  = SecondsPerHour * HoursPerDay
)

const (
	Zebi = 1 << 70
	Yobi = 1 << 80
)

const ZebisPerYobi = Yobi / Zebi

func CircleArea(radius float64) float64 {
	return math.Pi * radius * radius
}

const Kilo = 1000

func Kilograms(grams float64) float64 {
	return grams / Kilo
}

func Grams(kilograms int) int {
	return kilograms * Kilo
}

func TotalTimeout(retries int) int {
	return retries * int(Timeout)
}
