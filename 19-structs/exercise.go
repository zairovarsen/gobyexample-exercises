//go:build !solution

// Package structs contains exercises for https://gobyexample.com/structs
//
// Run the tests for this section with:
//
//	go test ./19-structs/
package structs

// Person describes a person by name and age.
//
// This struct is yours to declare: give it exactly two fields, Name of type
// string and Age of type int (spelled like this, with capital letters).
type Person struct {
	// TODO: declare the fields Name and Age.
	Name string
	Age int
}

// NewPerson is a constructor: it returns a pointer to a new Person with the
// given name and the age 42. Every call returns a different Person.
//
// As on the page: create the struct with a literal that names only the Name
// field, set the age with the dot notation afterwards and return the address
// of the struct.
func NewPerson(name string) *Person {
	person := Person{Name: name}
	person.Age = 42
	return &person
}

// Point is a position on a grid. (This one is declared for you.)
type Point struct {
	X, Y int
}

// Translate returns the point that lies dx to the right of and dy above p:
// Translate(Point{1, 2}, 10, 20) returns Point{11, 22}.
//
// Structs are values. p is a copy of the caller's point, so the caller's
// point does not change, whatever you do with p.
func Translate(p Point, dx, dy int) Point {
	// TODO: return the moved point.
	return Point{X: p.X + dx, Y: p.Y + dy}
}

// MoveTo changes the point that p points to so that it lies at (x, y).
//
// Structs are mutable, and with a pointer the function reaches the caller's
// struct. The dot works on struct pointers as well - the pointer is
// dereferenced automatically, so there is no need to write (*p).X.
func MoveTo(p *Point, x, y int) {
	p.X = x
	p.Y = y
}

// RexTheDog returns a value of an anonymous struct type - a struct type that
// has no name and is spelled out right where it is used. The value has the
// Name "Rex" and IsGood set to true.
//
// The stub below shows how to declare a variable of such a type. You can
// assign to its fields, or replace the declaration by a literal:
// dog := struct{ ...fields... }{ ...values... }.
func RexTheDog() struct {
	Name   string
	IsGood bool
} {
	// TODO: give the dog its name and make it a good dog.
	return struct {
		Name   string
		IsGood bool
	}{
		"Rex",
		true,
	}
	
}

// Rect is an axis-parallel rectangle, described by two opposite corners: Min
// holds the smallest X and Y, Max the largest. A struct can have fields that
// are structs themselves; they are reached with r.Min.X and so on.
type Rect struct {
	Min, Max Point
}

// BoundingBox returns the smallest Rect that contains all the given points:
// its Min has the smallest X and the smallest Y found in points, its Max the
// largest X and the largest Y. Without any points it returns the zero Rect.
//
// Example: BoundingBox([{1 5} {4 2} {3 9}]) returns Rect{Min: Point{1, 2},
// Max: Point{4, 9}}.
func BoundingBox(points []Point) Rect {
	if len(points) == 0 {
		return Rect{}
	}

	sX, lX := points[0].X, points[0].X 
	sY, lY := points[0].Y, points[0].Y

	for i := 1; i < len(points); i++ {
		v := points[i]
		if v.X < sX {
			sX = v.X
		}
		if v.X > lX {
			lX = v.X
		}
		if v.Y < sY {
			sY = v.Y
		}
		if v.Y > lY {
			lY = v.Y
		}
	}

	return Rect{Min: Point{sX, sY}, Max: Point{lX, lY}}
}
