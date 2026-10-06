//go:build !solution

// Package interfaces contains exercises for https://gobyexample.com/interfaces
//
// Run the tests for this section with:
//
//	go test ./21-interfaces/
package interfaces

// Shape is the interface for geometric shapes. An interface is a named
// collection of method signatures.
//
// This interface is yours to declare: give it exactly two methods, Area and
// Perimeter, both without parameters and both returning a float64.
type Shape interface {
	// TODO: list the two method signatures.
}

// Rect is a rectangle. It implements Shape simply by having the methods the
// interface asks for - there is no "implements" declaration in Go.
type Rect struct {
	Width, Height float64
}

// Area returns width times height.
func (r Rect) Area() float64 {
	// TODO: compute the area.
	return 0
}

// Perimeter returns twice the width plus twice the height.
func (r Rect) Perimeter() float64 {
	// TODO: compute the perimeter.
	return 0
}

// Circle is a circle with the given radius.
type Circle struct {
	Radius float64
}

// Area returns pi times the radius squared. Use math.Pi.
func (c Circle) Area() float64 {
	// TODO: compute the area.
	return 0
}

// Perimeter returns the circumference: two times pi times the radius.
func (c Circle) Perimeter() float64 {
	// TODO: compute the circumference.
	return 0
}

// Square is a square with the given side length.
//
// Its methods are NOT stubbed: make Square a Shape by declaring Area (side
// times side) and Perimeter (four times the side) yourself, below the type.
type Square struct {
	Side float64
}

// TODO: declare the methods Area and Perimeter of Square here.

// TotalArea returns the sum of the areas of all the given shapes.
//
// The function knows nothing about rectangles or circles: it works with
// every type that implements Shape, including types written long after it.
func TotalArea(shapes []Shape) float64 {
	// TODO: call Area on every shape and add up the results.
	return 0
}

// Radius returns the radius of s and true when the Shape holds a Circle. For
// every other shape it returns 0 and false.
//
// An interface value hides the concrete type inside it. A type assertion
// gets it back: c, ok := s.(Circle) sets ok to true when s holds a Circle,
// and then c is that Circle with all its fields.
func Radius(s Shape) (float64, bool) {
	// TODO: use a type assertion with the two-value form.
	return 0, false
}

// Largest returns the shape with the biggest area. When several shapes share
// the biggest area it returns the first of them, and for an empty slice it
// returns nil (the zero value of an interface type).
func Largest(shapes []Shape) Shape {
	// TODO: compare the areas and remember the best shape so far.
	return nil
}
