//go:build solution

package forloop

func SumTo(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func Repeat(s string, n int) string {
	out := ""
	for range n {
		out += s
	}
	return out
}

func CountDigits(n int) int {
	digits := 1
	for n >= 10 {
		n = n / 10
		digits++
	}
	return digits
}

func SumOdd(n int) int {
	sum := 0
	for i := range n {
		if i%2 == 0 {
			continue
		}
		sum += i
	}
	return sum
}

func SmallestDivisor(n int) int {
	d := 2
	for {
		if n%d == 0 {
			break
		}
		d++
	}
	return d
}

func CollatzSteps(n int) int {
	steps := 0
	for {
		if n == 1 {
			break
		}
		steps++
		if n%2 == 0 {
			n = n / 2
			continue
		}
		n = 3*n + 1
	}
	return steps
}
