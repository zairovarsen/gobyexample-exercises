//go:build solution

package panics

import (
	"errors"
	"fmt"
)

var ErrEmptyStack = errors.New("stack is empty")

func Fail() {
	panic("a problem")
}

func MustPositive(n int) int {
	if n <= 0 {
		panic(fmt.Sprintf("not positive: %d", n))
	}
	return n
}

func Element(nums []int, i int) int {
	return nums[i]
}

func Must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

type Stack struct {
	items []int
}

func (s *Stack) Push(n int) {
	s.items = append(s.items, n)
}

func (s *Stack) Len() int {
	return len(s.items)
}

func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		panic(ErrEmptyStack)
	}
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top
}
