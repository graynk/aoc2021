package main

import "testing"

func Test_day2(t *testing.T) {
	input := []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}

	t.Run("part 1", func(t *testing.T) {
		want := 150
		if got := part1(input); got != want {
			t.Errorf("part2() = %v, want %v", got, want)
		}
	})
	t.Run("part 2", func(t *testing.T) {
		want := 900
		if got := part2(input); got != want {
			t.Errorf("part2() = %v, want %v", got, want)
		}
	})
}
