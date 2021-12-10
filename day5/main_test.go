package main

import "testing"

func Test_day5(t *testing.T) {

	t.Run("part 1", func(t *testing.T) {
		want := 5
		board := parseInput("./testinput.txt", 10, false)
		if got := part1(board); got != want {
			t.Errorf("part1() = %v, want %v", got, want)
		}
	})
	t.Run("part 2", func(t *testing.T) {
		want := 12
		board := parseInput("./testinput.txt", 10, true)
		if got := part2(board); got != want {
			t.Errorf("part2() = %v, want %v", got, want)
		}
	})
}
