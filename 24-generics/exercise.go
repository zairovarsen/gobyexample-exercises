//go:build !solution

// Package generics contains exercises for https://gobyexample.com/generics
//
// Run the tests for this section with:
//
//	go test ./24-generics/
package generics

// Index returns the index of the first occurrence of v in s, or -1 if v is
// not present, e.g. Index([]string{"foo", "bar", "zoo"}, "zoo") returns 2.
//
// It has two type parameters: S, any slice type (the ~ also admits named
// types such as `type Words []string`), and E, the element type, which must
// be comparable so that == can be used on it.
func Index[S ~[]E, E comparable](s S, v E) int {
	// TODO: search s for v.
	for i, val := range s {
		if val == v {
			return i
		}
	}
	return -1
}

// Map returns a new slice holding f(x) for every element x of s, in order.
// The result has a different element type whenever f changes the type:
// Map([]int{1, 2}, strconv.Itoa) returns []string{"1", "2"}.
func Map[T, U any](s []T, f func(T) U) []U {
	// TODO: apply f to every element.
	result := make([]U, 0, len(s))

	for _, v := range s {
		value := f(v)
		result = append(result, value)
	}
	return result
}

// Number is a constraint: an interface that lists the types a type parameter
// may be instantiated with. It permits int, int64 and float64 and, because of
// the ~, every named type built on one of them.
type Number interface {
	~int | ~int64 | ~float64
}

// Sum adds up nums; the sum of no numbers is 0. Thanks to the constraint the
// + operator is available for T.
func Sum[T Number](nums []T) T {
	// TODO: add up the numbers.
	var zero T

	for _, num := range nums {
		zero += num
	}
	return zero
}

// List is a singly-linked list with elements of any type T. Its zero value
// is an empty list that is ready to use: `var lst List[int]`.
type List[T any] struct {
	// TODO: declare a generic element type (a value of type T plus a pointer
	// to the next element) and give List pointers to its first and its last
	// element.
	head, tail *Element[T]
}

type Element[T any] struct {
	value T
	next  *Element[T]
}

// Push appends v to the end of the list.
//
// Note that the receiver keeps the type parameter: List[T], not List.
func (lst *List[T]) Push(v T) {
	// TODO: link a new element behind the current last one.
	if lst.tail == nil {
		lst.head = &Element[T]{value: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &Element[T]{value: v}
		lst.tail = lst.tail.next
	}
}

// Len returns the number of elements in the list.
func (lst *List[T]) Len() int {
	// TODO: walk the list and count.
	n := 0
	tmp := lst.head

	for tmp != nil {
		n++
		tmp = tmp.next
	}
	return n
}

// AllElements returns the elements of the list as a slice, in the order they
// were pushed. For an empty list the result is empty.
func (lst *List[T]) AllElements() []T {
	// TODO: walk the list and collect the values.
	values := make([]T, 0)

	for tmp := lst.head; tmp != nil; tmp = tmp.next {
		values = append(values, tmp.value)
	}
	return values
}

// GroupBy sorts the elements of s into groups: the result maps every key
// that the key function produced to the elements it produced that key for,
// in their original order.
//
// Example: grouping []string{"go", "is", "fun"} by len gives
// map[int][]string{2: {"go", "is"}, 3: {"fun"}}.
func GroupBy[T any, K comparable](s []T, key func(T) K) map[K][]T {
	// TODO: build the map of groups.
	groups := make(map[K][]T, 0)

	for _, val := range s {
		k := key(val)
		if _, ok := groups[k]; !ok {
			groups[k] = make([]T, 0)
		}
		groups[k] = append(groups[k], val)
	}
	return groups
}
