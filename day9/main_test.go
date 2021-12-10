package main

import (
	"testing"
)

func Test_day9(t *testing.T) {
	input := parseInput("./testinput.txt")
	t.Run("part 1", func(t *testing.T) {
		want := 15
		if got := findRisk(input); got != want {
			t.Errorf("part1() = %v, want %v", got, want)
		}
	})
	t.Run("part 2", func(t *testing.T) {
		want := 1134
		if got := basinCounter(input); got != want {
			t.Errorf("part2() = %v, want %v", got, want)
		}
	})
}
