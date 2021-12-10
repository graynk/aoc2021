package main

import "testing"

func Test_day1(t *testing.T) {
	input := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	t.Run("part 1", func(t *testing.T) {
		want := 7
		if got := part1(input); got != want {
			t.Errorf("part1() = %v, want %v", got, want)
		}
	})
	t.Run("part 2", func(t *testing.T) {
		want := 5
		if got := part2(input); got != want {
			t.Errorf("part2() = %v, want %v", got, want)
		}
	})
}
