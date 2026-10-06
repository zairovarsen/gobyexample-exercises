//go:build !solution

// Package structembedding contains exercises for https://gobyexample.com/struct-embedding
//
// Run the tests for this section with:
//
//	go test ./23-struct-embedding/
package structembedding

// Base is a plain struct that other structs are going to embed.
type Base struct {
	Num int
}

// Describe returns "base with num=<Num>", e.g. "base with num=1".
func (b Base) Describe() string {
	// TODO: format the description.
	return ""
}

// Container embeds a Base and adds a field of its own. Thanks to the
// embedding, a Container c can be used as if it had Base's fields and
// methods itself: c.Num (short for c.Base.Num) and c.Describe().
//
// TODO: embed Base in Container. An embedding looks like a field without a
// name: just the type.
type Container struct {
	Str string
}

// NewContainer returns a Container whose embedded Base has the given num and
// whose Str is str.
//
// Hint: in a struct literal the embedded struct is initialised explicitly,
// with the type name serving as the field name.
func NewContainer(num int, str string) Container {
	// TODO: build the Container.
	return Container{}
}

// Describer is implemented by everything that can describe itself.
type Describer interface {
	Describe() string
}

// AsDescriber returns c itself as a Describer.
//
// Container never declares a Describe method: it gets one promoted from the
// embedded Base, and that is enough to implement the interface. (Returning c
// only compiles once Container embeds Base.)
func AsDescriber(c Container) Describer {
	// TODO: return c.
	return nil
}

// DescribeAll returns the description of every item, in order.
// Called without arguments it returns an empty result.
func DescribeAll(items ...Describer) []string {
	// TODO: collect the descriptions.
	return nil
}

// Point is a position on a grid.
type Point struct {
	X, Y int
}

// Move shifts the point by dx and dy. It has a pointer receiver because it
// changes the point.
func (p *Point) Move(dx, dy int) {
	// TODO: update X and Y.
}

// String formats the point like "(1,2)".
func (p Point) String() string {
	// TODO: format the point.
	return ""
}

// Circle is a Point (its centre) plus a radius. Embedding promotes the
// fields X and Y and the methods of Point, including Move, which changes
// the centre of the circle it is called on.
//
// TODO: embed Point in Circle.
type Circle struct {
	Radius int
}

// NewCircle returns a Circle with its centre at (x, y) and the given radius.
func NewCircle(x, y, radius int) Circle {
	// TODO: build the Circle.
	return Circle{}
}

// String formats the circle like "circle at (1,2) r=3".
//
// A method declared on the outer struct takes precedence over a promoted
// method with the same name. Without this method fmt.Println(circle) would
// use the String method promoted from Point and print only "(1,2)".
//
// Hint: the hidden method is still reachable as c.Point.String().
func (c Circle) String() string {
	// TODO: format the circle, reusing the String method of the embedded Point.
	return ""
}
