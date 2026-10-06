package interfaces

import (
	"math"
	"reflect"
	"testing"
)

// This file never calls a method on a variable of type Shape directly and
// looks at Square's methods at run time only. That way it compiles while the
// interface is still empty and Square has no methods yet.

func near(a, b float64) bool {
	return math.Abs(a-b) < 1e-9
}

// square returns a Square as a Shape, failing the test when Square does not
// implement the interface (yet).
func square(t *testing.T, side float64) Shape {
	t.Helper()
	s, ok := any(Square{Side: side}).(Shape)
	if !ok {
		t.Fatal("Square does not implement Shape - declare its methods Area() float64 and Perimeter() float64")
	}
	return s
}

// triangle is a Shape that the exercise code has never heard of.
type triangle struct {
	base, height float64
}

func (tr triangle) Area() float64      { return tr.base * tr.height / 2 }
func (tr triangle) Perimeter() float64 { return 0 }

func TestShapeInterface(t *testing.T) {
	typ := reflect.TypeFor[Shape]()
	for _, name := range []string{"Area", "Perimeter"} {
		m, ok := typ.MethodByName(name)
		if !ok {
			t.Errorf("interface Shape has no method %s - add it to the interface declaration", name)
			continue
		}
		if m.Type.NumIn() != 0 || m.Type.NumOut() != 1 || m.Type.Out(0).Kind() != reflect.Float64 {
			t.Errorf("method %s of Shape has the signature %v, want func() float64", name, m.Type)
		}
	}
	if n := typ.NumMethod(); n != 2 {
		t.Errorf("interface Shape has %d methods, want exactly 2 (Area and Perimeter)", n)
	}
}

func TestRect(t *testing.T) {
	tests := []struct {
		r                   Rect
		wantArea, wantPerim float64
	}{
		{Rect{Width: 3, Height: 4}, 12, 14},
		{Rect{Width: 2.5, Height: 2}, 5, 9},
		{Rect{Width: 1, Height: 1}, 1, 4},
	}
	for _, tt := range tests {
		if got := tt.r.Area(); !near(got, tt.wantArea) {
			t.Errorf("%+v.Area() = %v, want %v", tt.r, got, tt.wantArea)
		}
		if got := tt.r.Perimeter(); !near(got, tt.wantPerim) {
			t.Errorf("%+v.Perimeter() = %v, want %v", tt.r, got, tt.wantPerim)
		}
	}
}

func TestCircle(t *testing.T) {
	tests := []struct {
		c                   Circle
		wantArea, wantPerim float64
	}{
		{Circle{Radius: 1}, math.Pi, 2 * math.Pi},
		{Circle{Radius: 5}, 78.53981633974483, 31.41592653589793},
		{Circle{Radius: 0.5}, 0.7853981633974483, math.Pi},
	}
	for _, tt := range tests {
		if got := tt.c.Area(); !near(got, tt.wantArea) {
			t.Errorf("%+v.Area() = %v, want %v", tt.c, got, tt.wantArea)
		}
		if got := tt.c.Perimeter(); !near(got, tt.wantPerim) {
			t.Errorf("%+v.Perimeter() = %v, want %v", tt.c, got, tt.wantPerim)
		}
	}
}

func TestSquare(t *testing.T) {
	sq, ok := any(Square{Side: 3}).(interface {
		Area() float64
		Perimeter() float64
	})
	if !ok {
		t.Fatal("Square needs the methods Area() float64 and Perimeter() float64 - declare both of them")
	}
	if got := sq.Area(); !near(got, 9) {
		t.Errorf("Square{Side: 3}.Area() = %v, want 9", got)
	}
	if got := sq.Perimeter(); !near(got, 12) {
		t.Errorf("Square{Side: 3}.Perimeter() = %v, want 12", got)
	}
}

func TestTotalArea(t *testing.T) {
	shapes := []Shape{Rect{Width: 3, Height: 4}, Circle{Radius: 1}, triangle{base: 4, height: 5}}
	if got, want := TotalArea(shapes), 12+math.Pi+10; !near(got, want) {
		t.Errorf("TotalArea(%+v) = %v, want %v", shapes, got, want)
	}

	shapes = []Shape{square(t, 2), Rect{Width: 1, Height: 6}}
	if got, want := TotalArea(shapes), 10.0; !near(got, want) {
		t.Errorf("TotalArea(%+v) = %v, want %v", shapes, got, want)
	}

	if got := TotalArea(nil); got != 0 {
		t.Errorf("TotalArea(nil) = %v, want 0", got)
	}
}

func TestRadius(t *testing.T) {
	tests := []struct {
		s          Shape
		wantRadius float64
		wantOK     bool
	}{
		{Circle{Radius: 2.5}, 2.5, true},
		{Circle{Radius: 7}, 7, true},
		{Circle{}, 0, true},
		{Rect{Width: 3, Height: 4}, 0, false},
		{triangle{base: 1, height: 1}, 0, false},
	}
	for _, tt := range tests {
		if r, ok := Radius(tt.s); r != tt.wantRadius || ok != tt.wantOK {
			t.Errorf("Radius(%#v) = %v, %v, want %v, %v", tt.s, r, ok, tt.wantRadius, tt.wantOK)
		}
	}
}

func TestLargest(t *testing.T) {
	small := Rect{Width: 1, Height: 2}
	medium := Circle{Radius: 1}
	big := triangle{base: 10, height: 10}

	tests := []struct {
		shapes []Shape
		want   Shape
	}{
		{[]Shape{small, medium, big}, big},
		{[]Shape{big, small, medium}, big},
		{[]Shape{small, medium}, medium},
		{[]Shape{small}, small},
		// Equal areas: the first one wins.
		{[]Shape{Rect{Width: 2, Height: 8}, small, Rect{Width: 4, Height: 4}}, Rect{Width: 2, Height: 8}},
		{nil, nil},
	}
	for _, tt := range tests {
		// Two interface values are equal when they hold equal values of the
		// same concrete type.
		if got := Largest(tt.shapes); got != tt.want {
			t.Errorf("Largest(%+v) = %#v, want %#v", tt.shapes, got, tt.want)
		}
	}
}
