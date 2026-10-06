//go:build solution

package pointers

func ZeroPtr(p *int) {
	*p = 0
}

func Swap(a, b *int) {
	*a, *b = *b, *a
}

func NewInt(v int) *int {
	// v is this call's own copy of the argument, so its address is unique.
	return &v
}

func ValueOr(p *int, fallback int) int {
	if p == nil {
		return fallback
	}
	return *p
}

func FillArray(a *[3]int, v int) {
	for i := range len(a) {
		a[i] = v
	}
}

func Largest(nums []int) *int {
	if len(nums) == 0 {
		return nil
	}
	best := 0
	for i, num := range nums {
		if num > nums[best] {
			best = i
		}
	}
	return &nums[best]
}
