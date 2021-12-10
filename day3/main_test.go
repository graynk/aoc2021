package main

import "testing"

func Test_day3(t *testing.T) {
	input := []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}

	t.Run("part 1", func(t *testing.T) {
		want := int64(198)
		if got := part1(input); got != want {
			t.Errorf("part1() = %v, want %v", got, want)
		}
	})
	t.Run("part 2", func(t *testing.T) {
		want := int64(230)
		if got := part2(input); got != want {
			t.Errorf("part2() = %v, want %v", got, want)
		}
	})
}
