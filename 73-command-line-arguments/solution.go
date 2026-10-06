//go:build solution

package commandlinearguments

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func ProgramName(args []string) string {
	if len(args) == 0 {
		return ""
	}
	return args[0]
}

func Arguments(args []string) []string {
	if len(args) == 0 {
		return nil
	}
	return args[1:]
}

func ArgAt(args []string, n int) (string, bool) {
	if n < 1 || n >= len(args) {
		return "", false
	}
	return args[n], true
}

func Sum(args []string) (int, error) {
	total := 0
	for _, arg := range Arguments(args) {
		n, err := strconv.Atoi(arg)
		if err != nil {
			return 0, err
		}
		total += n
	}
	return total, nil
}

func Repeat(args []string) (string, error) {
	argsWithoutProg := Arguments(args)
	if len(argsWithoutProg) < 2 {
		return "", errors.New("usage: repeat <count> <word>...")
	}
	count, err := strconv.Atoi(argsWithoutProg[0])
	if err != nil {
		return "", fmt.Errorf("count: %w", err)
	}
	if count < 0 {
		return "", fmt.Errorf("count must not be negative, got %d", count)
	}
	line := strings.Join(argsWithoutProg[1:], " ") + "\n"
	return strings.Repeat(line, count), nil
}
