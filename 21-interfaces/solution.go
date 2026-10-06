//go:build solution

package interfaces

import "math"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rect struct {
	Width, Height float64
}

func (r Rect) Area() float64 {
	return r.Width * r.Height
}

func (r Rect) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

func (s Square) Perimeter() float64 {
	return 4 * s.Side
}

func TotalArea(shapes []Shape) float64 {
	total := 0.0
	for _, s := range shapes {
		total += s.Area()
	}
	return total
}

func Radius(s Shape) (float64, bool) {
	if c, ok := s.(Circle); ok {
		return c.Radius, true
	}
	return 0, false
}

func Largest(shapes []Shape) Shape {
	var largest Shape
	for _, s := range shapes {
		if largest == nil || s.Area() > largest.Area() {
			largest = s
		}
	}
	return largest
}
