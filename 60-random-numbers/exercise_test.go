package randomnumbers

import (
	"math"
	"math/rand/v2"
	"slices"
	"testing"
)

// Random results cannot be compared with fixed expected values. The tests
// check properties instead: the range, a roughly even distribution over many
// draws, and that the same seed reproduces the same sequence. The tolerances
// are so wide that a correct solution fails by chance practically never.

func TestRollDie(t *testing.T) {
	const rolls = 6000
	var counts [7]int
	for range rolls {
		n := RollDie()
		if n < 1 || n > 6 {
			t.Fatalf("RollDie() = %d, want a number from 1 to 6", n)
		}
		counts[n]++
	}
	// Each face is expected about 1000 times.
	for face := 1; face <= 6; face++ {
		if counts[face] < 700 || counts[face] > 1300 {
			t.Errorf("in %d rolls the face %d came up %d times, want roughly %d - is every face equally likely?",
				rolls, face, counts[face], rolls/6)
		}
	}
}

func TestFloatBetween(t *testing.T) {
	tests := []struct{ min, max float64 }{
		{5, 10},
		{0, 1},
		{-3, 3},
		{100, 100.5},
	}
	for _, tt := range tests {
		const draws = 10000
		sum := 0.0
		lowest, highest := math.Inf(1), math.Inf(-1)
		for range draws {
			f := FloatBetween(tt.min, tt.max)
			if f < tt.min || f >= tt.max {
				t.Fatalf("FloatBetween(%v, %v) = %v, want %v <= f < %v", tt.min, tt.max, f, tt.min, tt.max)
			}
			sum += f
			lowest, highest = math.Min(lowest, f), math.Max(highest, f)
		}
		width := tt.max - tt.min
		mean, wantMean := sum/draws, (tt.min+tt.max)/2
		if math.Abs(mean-wantMean) > 0.05*width {
			t.Errorf("the mean of %d draws of FloatBetween(%v, %v) is %v, want about %v",
				draws, tt.min, tt.max, mean, wantMean)
		}
		// The draws should cover the whole range, not just a part of it.
		if lowest > tt.min+0.05*width || highest < tt.max-0.05*width {
			t.Errorf("%d draws of FloatBetween(%v, %v) only cover %v to %v, want (nearly) the whole range",
				draws, tt.min, tt.max, lowest, highest)
		}
	}
}

func TestSeededInts(t *testing.T) {
	const n = 50
	a := SeededInts(42, 1024, n)
	if len(a) != n {
		t.Fatalf("SeededInts(42, 1024, %d) returned %d numbers, want %d", n, len(a), n)
	}
	for _, v := range a {
		if v < 0 || v >= 100 {
			t.Fatalf("SeededInts(42, 1024, %d) contains %d, want only numbers with 0 <= number < 100", n, v)
		}
	}

	// Same seeds, same sequence.
	if b := SeededInts(42, 1024, n); !slices.Equal(a, b) {
		t.Errorf("two calls of SeededInts(42, 1024, %d) differ:\n  %v\n  %v\nwant the same sequence for the same seeds", n, a, b)
	}
	// A shorter request is the beginning of the same sequence.
	if short := SeededInts(42, 1024, 10); !slices.Equal(short, a[:10]) {
		t.Errorf("SeededInts(42, 1024, 10) = %v, want the first 10 numbers of the longer sequence %v", short, a[:10])
	}
	// Other seeds, other sequence. Both seed values have to matter.
	for _, seeds := range [][2]uint64{{5, 1024}, {42, 1025}, {1024, 42}} {
		if other := SeededInts(seeds[0], seeds[1], n); slices.Equal(a, other) {
			t.Errorf("SeededInts(%d, %d, %d) equals SeededInts(42, 1024, %d), want another sequence for other seeds",
				seeds[0], seeds[1], n, n)
		}
	}
	// It should look random: many different values, not one repeated.
	distinct := slices.Clone(a)
	slices.Sort(distinct)
	distinct = slices.Compact(distinct)
	if len(distinct) < 20 {
		t.Errorf("SeededInts(42, 1024, %d) has only %d distinct values (%v), want a random-looking sequence", n, len(distinct), a)
	}

	if got := SeededInts(1, 2, 0); len(got) != 0 {
		t.Errorf("SeededInts(1, 2, 0) = %v, want no numbers", got)
	}
}

func TestPick(t *testing.T) {
	items := []string{"rock", "paper", "scissors", "lizard"}
	const picks = 4000

	r := rand.New(rand.NewPCG(7, 11))
	counts := make(map[string]int)
	var sequence []string
	for range picks {
		got := Pick(r, items)
		if !slices.Contains(items, got) {
			t.Fatalf("Pick(r, %q) = %q, want one of the items", items, got)
		}
		counts[got]++
		sequence = append(sequence, got)
	}
	// Each item is expected about 1000 times.
	for _, item := range items {
		if counts[item] < 700 || counts[item] > 1300 {
			t.Errorf("in %d picks %q was chosen %d times, want roughly %d - is every item equally likely?",
				picks, item, counts[item], picks/len(items))
		}
	}

	// A generator with the same seed must lead to the same picks: Pick has
	// to use r, not the global generator.
	r = rand.New(rand.NewPCG(7, 11))
	for i, want := range sequence[:100] {
		if got := Pick(r, items); got != want {
			t.Fatalf("with an identically seeded generator pick number %d is %q, but it was %q the first time - use r for the random choice",
				i+1, got, want)
		}
	}

	if got := Pick(r, []string{"only"}); got != "only" {
		t.Errorf("Pick(r, [only]) = %q, want %q", got, "only")
	}
	if got := Pick(r, nil); got != "" {
		t.Errorf("Pick(r, nil) = %q, want an empty string", got)
	}
}

func TestShuffled(t *testing.T) {
	items := make([]int, 30)
	for i := range items {
		items[i] = i + 1
	}
	in := slices.Clone(items)

	a := Shuffled(1, in)
	if !slices.Equal(in, items) {
		t.Fatalf("Shuffled changed its argument to %v, want it untouched", in)
	}
	// Still the same elements?
	sorted := slices.Clone(a)
	slices.Sort(sorted)
	if !slices.Equal(sorted, items) {
		t.Fatalf("Shuffled(1, 1..30) = %v, want the numbers 1..30 in another order", a)
	}
	if slices.Equal(a, items) {
		t.Errorf("Shuffled(1, 1..30) returned the numbers in their original order, want them shuffled")
	}

	// Same seed, same order.
	if b := Shuffled(1, in); !slices.Equal(a, b) {
		t.Errorf("two calls of Shuffled(1, 1..30) differ:\n  %v\n  %v\nwant the same order for the same seed", a, b)
	}
	// Another seed, another order.
	if c := Shuffled(2, in); slices.Equal(a, c) {
		t.Errorf("Shuffled(2, 1..30) equals Shuffled(1, 1..30), want another order for another seed")
	}

	// Over many seeds every number should reach the first position about
	// equally often (300 times each here).
	const seeds = 9000
	var first [31]int
	for seed := range uint64(seeds) {
		s := Shuffled(seed, in)
		if len(s) != len(items) || s[0] < 1 || s[0] > 30 {
			t.Fatalf("Shuffled(%d, 1..30) = %v, want the numbers 1..30 in another order", seed, s)
		}
		first[s[0]]++
	}
	for n := 1; n <= 30; n++ {
		if first[n] < 180 || first[n] > 420 {
			t.Errorf("over %d seeds the number %d came first %d times, want roughly %d - the shuffle is not uniform",
				seeds, n, first[n], seeds/30)
		}
	}

	if got := Shuffled(1, nil); len(got) != 0 {
		t.Errorf("Shuffled(1, nil) = %v, want an empty result", got)
	}
}
