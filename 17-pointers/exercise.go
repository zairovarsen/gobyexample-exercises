//go:build !solution

// Package pointers contains exercises for https://gobyexample.com/pointers
//
// Run the tests for this section with:
//
//	go test ./17-pointers/
package pointers

// ZeroPtr sets the int that p points to to 0. Unlike a function that takes a
// plain int (and therefore works on a copy), it changes the caller's
// variable:
//
//	i := 5
//	ZeroPtr(&i) // i is now 0
//
// *p "dereferences" the pointer: it stands for the variable p points to.
func ZeroPtr(p *int) {
	// TODO: assign through the pointer.
}

// Swap exchanges the values of the two variables that a and b point to:
//
//	x, y := 1, 2
//	Swap(&x, &y) // x is now 2, y is now 1
func Swap(a, b *int) {
	// TODO: exchange *a and *b.
}

// NewInt returns a pointer to a new int variable that holds v. Every call
// returns a pointer to a different variable.
//
// The & operator gives the address of a variable - and in Go it is fine to
// return the address of a local variable.
func NewInt(v int) *int {
	// TODO: return the address of a variable holding v.
	return nil
}

// ValueOr returns the int that p points to. A pointer can also point to
// nothing at all: the zero value of a pointer is nil, and dereferencing nil
// crashes the program. For a nil pointer ValueOr returns fallback instead.
func ValueOr(p *int, fallback int) int {
	// TODO: check for nil before dereferencing.
	return 0
}

// FillArray sets all three elements of the array that a points to to v.
//
// Arrays are values, so a function with a parameter of type [3]int would
// only fill its own copy. A pointer to the array reaches the caller's array.
// (You can write a[i] instead of (*a)[i]: Go dereferences array pointers
// automatically.)
func FillArray(a *[3]int, v int) {
	// TODO: set every element through the pointer.
}

// Largest returns a pointer to the largest element of nums - to the element
// inside the slice, not to a copy of its value - so that the caller can
// change that element through the pointer:
//
//	nums := []int{3, 9, 4}
//	*Largest(nums) = 0 // nums is now [3 0 4]
//
// When the largest number occurs more than once, return a pointer to its
// first occurrence. For an empty slice return nil.
//
// Careful: the value variable of a range loop is a copy of the element, so
// its address is useless here. &nums[i] is the address of the element.
func Largest(nums []int) *int {
	// TODO: find the index of the largest element and return its address.
	return nil
}
