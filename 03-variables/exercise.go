//go:build !solution

// Package variables contains exercises for https://gobyexample.com/variables
//
// Run the tests for this section with:
//
//	go test ./03-variables/
package variables

// Initial declares a variable with the var keyword that holds the string
// "initial" (let Go infer the type) and returns that variable.
func Initial() string {
	// TODO: declare the variable with var, then return it.
	var str = "initial"
	return str
}

// DeclareTwo declares two int variables in one var statement, giving them the
// values 1 and 2, and returns their sum.
func DeclareTwo() int {
	// TODO: declare both variables on a single line and add them.
	var one, two int = 1,2
	return one + two
}

// ZeroValues declares four variables WITHOUT initialising them - an int, a
// float64, a string and a bool - and reports whether every one of them holds
// its zero value: 0, 0, "" and false.
//
// In Go a variable always has a value, so the answer should be true.
func ZeroValues() bool {
	var a int
	var b float64
	var c string
	var d bool

	return a == 0 && b == 0 && c == "" && d == false
}

// Shout uses the short declaration syntax := to make a variable that holds
// word followed by "!", and returns that variable twice with a space in
// between: Shout("go") returns "go! go!".
func Shout(word string) string {
	w := word
	// TODO: declare the variable with := and build the result from it.
	return w + "!" + " " + w + "!"
}

// CoinCount returns the smallest number of coins that add up to the given
// amount of cents, using coins worth 25, 10, 5 and 1 cent.
//
// Example: 99 cents = 3 x 25 + 2 x 10 + 0 x 5 + 4 x 1, so CoinCount(99)
// returns 9.
//
// No loop or if is needed: take as many of the biggest coin as fit (integer
// division), keep what is left over in a variable (the % operator gives the
// remainder of a division) and carry on with the next coin. You will declare
// some variables and assign new values to others.
func CoinCount(cents int) int {
	// TODO: count quarters, dimes, nickels and pennies.
	var a, b, c, d int
	 a = cents / 25
	var aLeftOver = cents % 25
	 b = aLeftOver / 10
	var bLeftOver = aLeftOver % 10
	 c = bLeftOver / 5
	var cLeftOver = bLeftOver % 5
	 d = cLeftOver / 1
	return a + b + c + d
}
