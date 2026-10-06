//go:build solution

package generics

func Index[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

func Map[T, U any](s []T, f func(T) U) []U {
	mapped := make([]U, 0, len(s))
	for _, x := range s {
		mapped = append(mapped, f(x))
	}
	return mapped
}

type Number interface {
	~int | ~int64 | ~float64
}

func Sum[T Number](nums []T) T {
	var total T
	for _, n := range nums {
		total += n
	}
	return total
}

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) Len() int {
	n := 0
	for e := lst.head; e != nil; e = e.next {
		n++
	}
	return n
}

func (lst *List[T]) AllElements() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func GroupBy[T any, K comparable](s []T, key func(T) K) map[K][]T {
	groups := make(map[K][]T)
	for _, x := range s {
		k := key(x)
		groups[k] = append(groups[k], x)
	}
	return groups
}
