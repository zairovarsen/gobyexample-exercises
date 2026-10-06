//go:build solution

package structs

type Person struct {
	Name string
	Age  int
}

func NewPerson(name string) *Person {
	p := Person{Name: name}
	p.Age = 42
	return &p
}

type Point struct {
	X, Y int
}

func Translate(p Point, dx, dy int) Point {
	return Point{X: p.X + dx, Y: p.Y + dy}
}

func MoveTo(p *Point, x, y int) {
	p.X = x
	p.Y = y
}

func RexTheDog() struct {
	Name   string
	IsGood bool
} {
	dog := struct {
		Name   string
		IsGood bool
	}{
		"Rex",
		true,
	}
	return dog
}

type Rect struct {
	Min, Max Point
}

func BoundingBox(points []Point) Rect {
	if len(points) == 0 {
		return Rect{}
	}
	box := Rect{Min: points[0], Max: points[0]}
	for _, p := range points[1:] {
		if p.X < box.Min.X {
			box.Min.X = p.X
		}
		if p.Y < box.Min.Y {
			box.Min.Y = p.Y
		}
		if p.X > box.Max.X {
			box.Max.X = p.X
		}
		if p.Y > box.Max.Y {
			box.Max.Y = p.Y
		}
	}
	return box
}
