//go:build solution

package arrays

func Squares() [5]int {
	var a [5]int
	for i := range len(a) {
		a[i] = i * i
	}
	return a
}

func Sum(nums [5]int) int {
	total := 0
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}
	return total
}

func Sparse() [5]int {
	return [...]int{10, 3: 40, 50}
}

func Reversed(a [5]int) [5]int {
	var r [5]int
	for i := range len(a) {
		r[len(a)-1-i] = a[i]
	}
	return r
}

func Table() [3][3]int {
	var table [3][3]int
	for i := range 3 {
		for j := range 3 {
			table[i][j] = (i + 1) * (j + 1)
		}
	}
	return table
}

func Transpose(m [2][3]int) [3][2]int {
	var out [3][2]int
	for i := range 2 {
		for j := range 3 {
			out[j][i] = m[i][j]
		}
	}
	return out
}
