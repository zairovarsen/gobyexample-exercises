//go:build solution

package variables

func Initial() string {
	var a = "initial"
	return a
}

func DeclareTwo() int {
	var b, c int = 1, 2
	return b + c
}

func ZeroValues() bool {
	var i int
	var f float64
	var s string
	var b bool
	return i == 0 && f == 0 && s == "" && !b
}

func Shout(word string) string {
	loud := word + "!"
	return loud + " " + loud
}

func CoinCount(cents int) int {
	var coins int

	quarters := cents / 25
	cents = cents % 25
	coins = coins + quarters

	dimes := cents / 10
	cents = cents % 10
	coins = coins + dimes

	nickels := cents / 5
	cents = cents % 5
	coins = coins + nickels

	// Whatever is left is paid in pennies.
	return coins + cents
}
