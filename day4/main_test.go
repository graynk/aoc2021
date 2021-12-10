package main

import "testing"

func Test_day4(t *testing.T) {
	numbers, boards := parseInput("./testinput.txt")

	t.Run("part 1", func(t *testing.T) {
		want := 4512
		if got := part1(numbers, boards); got != want {
			t.Errorf("part1() = %v, want %v", got, want)
		}
	})
	t.Run("part 2", func(t *testing.T) {
		want := 1924
		if got := part2(numbers, boards); got != want {
			t.Errorf("part2() = %v, want %v", got, want)
		}
	})
}
