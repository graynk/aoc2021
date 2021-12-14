package main

import (
	"testing"
)

func Test_day14(t *testing.T) {
	input, pi := parseInput("./testinput.txt")
	t.Run("part 1", func(t *testing.T) {
		want := []string{"NCNBCHB", "NBCCNBBBCBHCB", "NBBBCNCCNBBNBNBBCHBHHBCHB", "NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB"}
		for i := 0; i < 4; i++ {
			if input = pi.insertPolymer(input); input != want[i] {
				t.Errorf("part1() = %v, want %v", input, want[i])
			}
		}
		for i := 4; i < 10; i++ {
			input = pi.insertPolymer(input)
		}

		var wantCount int64 = 1588
		if got := commonCounter(input); got != wantCount {
			t.Errorf("part1() = %v, want %v", got, wantCount)
		}
	})
	t.Run("part 2", func(t *testing.T) {
		var want int64 = 2188189693529
		for i := 0; i < 40; i++ {
			input = pi.insertPolymer(input)
		}
		if got := commonCounter(input); got != want {
			t.Errorf("part2() = %v, want %v", got, want)
		}
	})
}
