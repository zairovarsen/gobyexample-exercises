//go:build !solution

// Package variadicfunctions contains exercises for https://gobyexample.com/variadic-functions
//
// Run the tests for this section with:
//
//	go test ./13-variadic-functions/
package variadicfunctions

// Sum adds up any number of ints: Sum(1, 2, 3) returns 6 and Sum() returns 0.
//
// Inside the function nums is an ordinary []int: you can call len on it and
// loop over it with for _, num := range nums.
func Sum(nums ...int) int {
	// TODO: add up nums.
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

// Average returns the mean of its arguments: Average(1, 2, 4) returns
// 2.3333... Without arguments it returns 0.
func Average(nums ...float64) float64 {
	if len(nums) == 0 {
		return 0
	}
	sum := 0.0
	for _, num := range nums {
		sum += num
	}
	return sum / float64(len(nums))
}

// Join glues parts together with sep between them:
// Join("-", "a", "b", "c") returns "a-b-c", Join("-", "a") returns "a" and
// Join("-") returns "".
//
// Only the last parameter of a function can be variadic; ordinary parameters
// such as sep come before it.
func Join(sep string, parts ...string) string {
	str := ""
	// TODO: concatenate the parts, with sep in between.

	for i, part := range parts {
		if i != 0 {
			str += sep
		}
		str += part
	}
	return str
}

// Max returns the largest of its arguments: Max(3, 9, 4) returns 9.
//
// The maximum of nothing makes no sense, and the signature says so: the
// ordinary parameter first forces callers to pass at least one number.
func Max(first int, rest ...int) int {
	// TODO: start with first and compare it with the rest.
	max := first
	for _, n := range rest {
		if n > max {
			max = n
		}
	}
	return max
}

// Total returns the sum of all numbers in all the given slices:
// Total([1 2], [3], [4 5 6]) returns 21.
//
// You already have Sum, but it wants individual ints, not a slice. A slice
// can be spread over a variadic parameter with a trailing ...: Sum(list...)
func Total(lists ...[]int) int {
	// TODO: call Sum for every list, spreading the slice.
	sum := 0
	for _, list := range lists {
		for i := range len(list) {
			sum += list[i]
		}
	}
	return sum
}

// Greet builds a greeting for any number of people:
//
//	Greet("Hello")                      returns "Hello!"
//	Greet("Hello", "Ann")               returns "Hello, Ann!"
//	Greet("Hello", "Ann", "Bob")        returns "Hello, Ann and Bob!"
//	Greet("Hi", "Ann", "Bob", "Cy")     returns "Hi, Ann, Bob and Cy!"
//
// With three or more names, all names but the last are separated by ", " and
// the last one is attached with " and ".
//
// Hint: names[:len(names)-1] is all names but the last - and Join accepts
// them when you spread the slice.
func Greet(greeting string, names ...string) string {
	// TODO: handle no names, one name, and several names.
	if len(names) == 0 {
		return greeting + "!"
	} else if len(names) == 1 {
		return greeting + ", " + names[0] + "!"
	} else {
		return greeting + ", " + Join(", ", names[:len(names)-1]...) + " and " + names[len(names)-1] + "!"
	}
}
