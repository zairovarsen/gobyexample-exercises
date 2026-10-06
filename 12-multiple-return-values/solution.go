//go:build solution

package multiplereturnvalues

func Swap(a, b string) (string, string) {
	return b, a
}

func DivMod(a, b int) (int, int) {
	return a / b, a % b
}

func Clock(seconds int) (int, int, int) {
	hours, rest := DivMod(seconds, 3600)
	minutes, secs := DivMod(rest, 60)
	return hours, minutes, secs
}

func MinMax(nums []int) (int, int) {
	if len(nums) == 0 {
		return 0, 0
	}
	lo, hi := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < lo {
			lo = nums[i]
		}
		if nums[i] > hi {
			hi = nums[i]
		}
	}
	return lo, hi
}

func Largest(nums []int) int {
	_, hi := MinMax(nums)
	return hi
}

func Winner(votes []string) (string, int) {
	counts := make(map[string]int)
	for i := range len(votes) {
		counts[votes[i]]++
	}

	// Walking the votes in order (and replacing the leader only on a strictly
	// higher count) makes the earliest candidate win a tie.
	winner, best := "", 0
	for i := range len(votes) {
		if c := counts[votes[i]]; c > best {
			winner, best = votes[i], c
		}
	}
	return winner, best
}
