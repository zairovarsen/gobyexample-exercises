//go:build solution

package variadicfunctions

func Sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func Average(nums ...float64) float64 {
	if len(nums) == 0 {
		return 0
	}
	total := 0.0
	for _, num := range nums {
		total += num
	}
	return total / float64(len(nums))
}

func Join(sep string, parts ...string) string {
	out := ""
	for i, part := range parts {
		if i > 0 {
			out += sep
		}
		out += part
	}
	return out
}

func Max(first int, rest ...int) int {
	largest := first
	for _, num := range rest {
		if num > largest {
			largest = num
		}
	}
	return largest
}

func Total(lists ...[]int) int {
	total := 0
	for _, list := range lists {
		total += Sum(list...)
	}
	return total
}

func Greet(greeting string, names ...string) string {
	switch len(names) {
	case 0:
		return greeting + "!"
	case 1:
		return greeting + ", " + names[0] + "!"
	default:
		last := len(names) - 1
		return greeting + ", " + Join(", ", names[:last]...) + " and " + names[last] + "!"
	}
}
