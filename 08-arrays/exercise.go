//go:build !solution

// Package arrays contains exercises for https://gobyexample.com/arrays
//
// Run the tests for this section with:
//
//	go test ./08-arrays/
package arrays

// Squares returns an array of the first five square numbers: element i holds
// i*i, so the result is [0 1 4 9 16].
//
// Declare the array with var (all elements start as zero) and set the
// elements in a loop with a[i] = ...
func Squares() [5]int {
	return [5]int{0,1,4,9,16}
}

// Sum adds up all elements of nums: Sum([5]int{1, 2, 3, 4, 5}) returns 15.
//
// Use the builtin len for the loop bound instead of writing 5.
func Sum(nums [5]int) int {
	length := len(nums)
	sum := 0

	for i := 0; i < length; i++ {
		sum += nums[i]		
	}
	return sum
}

// Sparse returns the array [10 0 0 40 50] - written as ONE array literal.
//
// Let the compiler count the elements with [...] and jump to an index with
// the index: value form, as in [...]int{100, 3: 400, 500}.
func Sparse() [5]int {
	return [...]int{10, 3: 40, 50}
}

// Reversed returns the elements of a in opposite order:
// Reversed([5]int{1, 2, 3, 4, 5}) returns [5 4 3 2 1].
//
// Arrays are values: the function works on its own copy of a, so whatever
// you do here, the caller's array stays as it was.
func Reversed(a [5]int) [5]int {
	var b [5]int
	
	for i := range 5 {
		b[i] = a[5 - i - 1]
	}
	return b
}

// Table returns a 3x3 multiplication table as a two-dimensional array:
// element [i][j] holds (i+1) * (j+1).
//
//	[[1 2 3] [2 4 6] [3 6 9]]
func Table() [3][3]int {
	var table [3][3]int
	
	for i := range 3 {
		for j := range 3 {
			table[i][j] = (i + 1) * (j + 1)
		}
	}
	// TODO: fill the table with two nested loops.
	return table
}

// Transpose flips a matrix with 2 rows and 3 columns over its diagonal: the
// result has 3 rows and 2 columns, and its element [j][i] is m[i][j].
//
//	[[1 2 3] [4 5 6]]  ->  [[1 4] [2 5] [3 6]]
//
// Note that [2][3]int and [3][2]int are different types: the lengths are
// part of an array's type.
func Transpose(m [2][3]int) [3][2]int {
	// TODO: copy every element to its new position.
	return [3][2]int{}
}
