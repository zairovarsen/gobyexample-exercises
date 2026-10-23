//go:build !solution

// Package methods contains exercises for https://gobyexample.com/methods
//
// Run the tests for this section with:
//
//	go test ./20-methods/
package methods

// Rect is a rectangle with a width and a height.
type Rect struct {
	Width, Height int
}

// Area returns the area of the rectangle: width times height.
//
// (r Rect) in front of the method name is the receiver: Area is a method
// with a value receiver, it works on a copy of the Rect it is called on.
func (r Rect) Area() int {
	return r.Width * r.Height
}

// Perimeter returns the length of the outline of the rectangle: twice the
// width plus twice the height.
func (r Rect) Perimeter() int {
	// TODO: compute the perimeter.
	return r.Height * 2 + r.Width * 2
}

// Scale multiplies both sides of the rectangle by factor, changing the Rect
// it is called on:
//
//	r := Rect{Width: 2, Height: 3}
//	r.Scale(10) // r is now {20 30}
//
// This only works because Scale has a pointer receiver: the method gets the
// address of the caller's Rect instead of a copy. (The call r.Scale(10) on a
// Rect variable is shorthand for (&r).Scale(10) - Go converts between values
// and pointers for method calls.)
func (r *Rect) Scale(factor int) {
	// TODO: change the receiver's fields.
	r.Height *= factor
	r.Width *= factor
}

// Scaled is the value-receiver counterpart of Scale: it returns a new Rect
// whose sides are factor times as long and leaves the Rect it is called on
// alone.
func (r Rect) Scaled(factor int) Rect {
	// TODO: return the bigger Rect.
	return Rect{Height: r.Height * factor, Width: r.Width * factor}
}

// Counter counts upwards, starting at zero.
//
// Its methods are NOT stubbed - declare both of them yourself, below the
// type:
//
//   - Increment, without parameters and results, adds one to the count
//   - Value, without parameters, returns the current count as an int
//
// Think about the receiver type each method needs: a method that has to
// change the Counter it is called on cannot work on a copy.
type Counter struct {
	count int
}

// TODO: declare the methods Increment and Value of Counter here.
func (s *Counter) Increment() {
	s.count += 1
}

func (s *Counter) Value() int {
	return s.count
}

// Stack is a last-in, first-out collection of ints: the number that was
// pushed last is the first one to be popped. The zero Stack is empty and
// ready to use.
type Stack struct {
	items []int
}

// Push puts v on top of the stack.
func (s *Stack) Push(v int) {
	// TODO: append to s.items.
	s.items = append(s.items, v)
}

// Pop removes the number on top of the stack and returns it together with
// true. When the stack is empty there is nothing to pop: Pop returns 0 and
// false.
func (s *Stack) Pop() (int, bool) {
	// TODO: take the last element off s.items.
	l := s.Len()

	if l == 0 {
		return 0, false
	}

	v := s.items[l - 1]
	s.items = s.items[:l - 1]	

	return v, true
}

// Len returns how many numbers are on the stack.
func (s *Stack) Len() int {
	// TODO: return the number of items.
	return len(s.items)
}
