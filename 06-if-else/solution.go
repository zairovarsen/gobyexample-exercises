//go:build solution

package ifelse

import "strconv"

func EvenOrOdd(n int) string {
	if n%2 == 0 {
		return "even"
	} else {
		return "odd"
	}
}

func Compare(a, b int) string {
	if diff := a - b; diff < 0 {
		return "less"
	} else if diff == 0 {
		return "equal"
	} else {
		return "greater"
	}
}

func LeapYear(year int) bool {
	return year%4 == 0 && year%100 != 0 || year%400 == 0
}

func Grade(score int) string {
	if score < 0 || score > 100 {
		return "invalid"
	} else if score >= 90 {
		return "A"
	} else if score >= 80 {
		return "B"
	} else if score >= 70 {
		return "C"
	} else if score >= 60 {
		return "D"
	} else {
		return "F"
	}
}

func FizzBuzz(n int) string {
	if n%15 == 0 {
		return "FizzBuzz"
	} else if n%3 == 0 {
		return "Fizz"
	} else if n%5 == 0 {
		return "Buzz"
	}
	return strconv.Itoa(n)
}

func FizzBuzzUpTo(n int) string {
	out := ""
	for i := 1; i <= n; i++ {
		if i > 1 {
			out += " "
		}
		out += FizzBuzz(i)
	}
	return out
}
