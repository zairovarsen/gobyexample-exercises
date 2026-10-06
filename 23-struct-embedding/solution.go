//go:build solution

package structembedding

import "fmt"

type Base struct {
	Num int
}

func (b Base) Describe() string {
	return fmt.Sprintf("base with num=%v", b.Num)
}

type Container struct {
	Base
	Str string
}

func NewContainer(num int, str string) Container {
	return Container{
		Base: Base{Num: num},
		Str:  str,
	}
}

type Describer interface {
	Describe() string
}

func AsDescriber(c Container) Describer {
	return c
}

func DescribeAll(items ...Describer) []string {
	descriptions := make([]string, 0, len(items))
	for _, item := range items {
		descriptions = append(descriptions, item.Describe())
	}
	return descriptions
}

type Point struct {
	X, Y int
}

func (p *Point) Move(dx, dy int) {
	p.X += dx
	p.Y += dy
}

func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}

type Circle struct {
	Point
	Radius int
}

func NewCircle(x, y, radius int) Circle {
	return Circle{
		Point:  Point{X: x, Y: y},
		Radius: radius,
	}
}

func (c Circle) String() string {
	return fmt.Sprintf("circle at %s r=%d", c.Point.String(), c.Radius)
}
