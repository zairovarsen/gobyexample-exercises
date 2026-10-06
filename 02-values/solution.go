//go:build solution

package values

import "fmt"

func FullName(first, last string) string {
	return first + " " + last
}

func Average(a, b float64) float64 {
	return (a + b) / 2.0
}

func IntDivide(a, b int) int {
	return a / b
}

func PrintValues() {
	fmt.Println("go" + "by" + "example")
	fmt.Println("6*7 =", 6*7)
	fmt.Println("1.0/4.0 =", 1.0/4.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}

func Xor(a, b bool) bool {
	return (a || b) && !(a && b)
}

func Majority(a, b, c bool) bool {
	return a && b || a && c || b && c
}
