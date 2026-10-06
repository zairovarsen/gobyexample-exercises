package variables

import "testing"

func TestInitial(t *testing.T) {
	if got, want := Initial(), "initial"; got != want {
		t.Errorf("Initial() = %q, want %q", got, want)
	}
}

func TestDeclareTwo(t *testing.T) {
	if got, want := DeclareTwo(), 3; got != want {
		t.Errorf("DeclareTwo() = %d, want %d", got, want)
	}
}

func TestZeroValues(t *testing.T) {
	if !ZeroValues() {
		t.Error("ZeroValues() = false, want true: uninitialised variables hold 0, 0, \"\" and false")
	}
}

func TestShout(t *testing.T) {
	tests := []struct {
		word string
		want string
	}{
		{"go", "go! go!"},
		{"apple", "apple! apple!"},
		{"", "! !"},
	}
	for _, tt := range tests {
		if got := Shout(tt.word); got != tt.want {
			t.Errorf("Shout(%q) = %q, want %q", tt.word, got, tt.want)
		}
	}
}

func TestCoinCount(t *testing.T) {
	tests := []struct {
		cents int
		want  int
	}{
		{99, 9},
		{25, 1},
		{30, 2},
		{41, 4},
		{4, 4},
		{100, 4},
		{0, 0},
	}
	for _, tt := range tests {
		if got := CoinCount(tt.cents); got != tt.want {
			t.Errorf("CoinCount(%d) = %d, want %d", tt.cents, got, tt.want)
		}
	}
}
