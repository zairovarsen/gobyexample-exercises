//go:build !solution

// Package multiplereturnvalues contains exercises for https://gobyexample.com/multiple-return-values
//
// Run the tests for this section with:
//
//	go test ./12-multiple-return-values/
package multiplereturnvalues

// Swap returns its two arguments in opposite order: Swap("a", "b") returns
// "b", "a".
func Swap(a, b string) (string, string) {
	// TODO: return both values.
	return b, a
}

// DivMod returns the quotient and the remainder of the integer division
// a / b: DivMod(17, 5) returns 3, 2. b is never 0 in the tests.
func DivMod(a, b int) (int, int) {
	// TODO: return a / b and a % b.
	return a/b, a%b
}

// Clock splits a number of seconds into hours, minutes and seconds:
// Clock(3725) returns 1, 2, 5 because 3725 s = 1 h 2 min 5 s.
//
// A function can return more than two values. DivMod is useful here: receive
// its two results with a, b := DivMod(...).
func Clock(seconds int) (int, int, int) {
	// TODO: split off the hours, then the minutes.
	hours, seconds := DivMod(seconds, 60 * 60)
	minutes, seconds := DivMod(seconds, 60)

	return hours, minutes, seconds
}

// MinMax returns the smallest and the largest number in nums:
// MinMax([3 9 -2 4]) returns -2, 9. For an empty slice it returns 0, 0.
func MinMax(nums []int) (int, int) {
	// TODO: find both in one pass over nums.
	if len(nums) == 0 {
		return 0,0
	}
	min, max := nums[0], nums[0]
	
	for i:=1; i < len(nums); i++ {
		if (nums[i] > max) {
			max = nums[i]
		} else if (nums[i] < min) {
			min = nums[i]
		}
	}

	return min, max
}

// Largest returns the largest number in nums, or 0 for an empty slice.
//
// Do not write another loop: call MinMax and throw away the result you do
// not need with the blank identifier _.
func Largest(nums []int) int {
	// TODO: call MinMax, ignore the minimum.
	_, largest := MinMax(nums)
	return largest
}

// Winner counts the votes of an election. Every element of votes is the name
// of a candidate; Winner returns the name with the most votes and how many
// votes that candidate got. When several candidates have the same highest
// count, the one whose first vote comes earliest in the slice wins. Without
// votes there is no winner: return "" and 0.
//
// Example: Winner([ann bob ann cy bob ann]) returns "ann", 3.
//
// Hint: count with a map, then go through votes (not the map) a second time
// to find the first candidate with the highest count.
func Winner(votes []string) (string, int) {
	// TODO: count the votes and pick the winner.
	votesCount := make(map[string]int)
	for i := 0; i < len(votes); i++ {
		votesCount[votes[i]]++
	}
	highestCount, highestCountName := 0, ""

	for i := 0; i < len(votes); i++ {
		count := votesCount[votes[i]]
		if (count > highestCount) {
			highestCount = count
			highestCountName = votes[i]
		}
	}

	return highestCountName, highestCount 
}
