//go:build solution

package methods

type Rect struct {
	Width, Height int
}

func (r Rect) Area() int {
	return r.Width * r.Height
}

func (r Rect) Perimeter() int {
	return 2*r.Width + 2*r.Height
}

func (r *Rect) Scale(factor int) {
	r.Width *= factor
	r.Height *= factor
}

func (r Rect) Scaled(factor int) Rect {
	return Rect{Width: r.Width * factor, Height: r.Height * factor}
}

type Counter struct {
	count int
}

// A pointer receiver, because Increment mutates the Counter.
func (c *Counter) Increment() {
	c.count++
}

func (c *Counter) Value() int {
	return c.count
}

type Stack struct {
	items []int
}

func (s *Stack) Push(v int) {
	s.items = append(s.items, v)
}

func (s *Stack) Pop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	last := len(s.items) - 1
	v := s.items[last]
	s.items = s.items[:last]
	return v, true
}

func (s *Stack) Len() int {
	return len(s.items)
}
