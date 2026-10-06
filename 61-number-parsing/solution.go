//go:build solution

package numberparsing

import (
	"fmt"
	"strconv"
)

func ParsePrice(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

func ParseCount(s string) (int, error) {
	return strconv.Atoi(s)
}

func ParseAnyBase(s string) (int64, error) {
	return strconv.ParseInt(s, 0, 64)
}

func ParseBinary(s string) (int64, error) {
	return strconv.ParseInt(s, 2, 64)
}

func ParseByte(s string) (uint8, error) {
	n, err := strconv.ParseUint(s, 10, 8)
	if err != nil {
		return 0, err
	}
	return uint8(n), nil
}

func SumFields(fields []string) (float64, error) {
	sum := 0.0
	for i, field := range fields {
		f, err := strconv.ParseFloat(field, 64)
		if err != nil {
			return 0, fmt.Errorf("field %d: %w", i, err)
		}
		sum += f
	}
	return sum, nil
}
