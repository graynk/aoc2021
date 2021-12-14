package main

import (
	"testing"
)

func Test_day14(t *testing.T) {
	t.Run("part 1", func(t *testing.T) {
		var want int64 = 1588
		pi := parseInput("./testinput.txt")
		for i := 0; i < 10; i++ {
			pi.insertPolymers()
		}

		if got := pi.commonCounter(); got != want {
			t.Errorf("part1() = %v, want %v", got, want)
		}
	})
	t.Run("part 2", func(t *testing.T) {
		var want int64 = 2188189693529
		pi := parseInput("./testinput.txt")
		for i := 0; i < 40; i++ {
			pi.insertPolymers()
		}
		if got := pi.commonCounter(); got != want {
			t.Errorf("part2() = %v, want %v", got, want)
		}
	})
}
