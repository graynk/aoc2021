package main

import (
	"testing"
)

func Test_day6(t *testing.T) {
	input := []int{3, 4, 3, 1, 2}

	t.Run("part 1", func(t *testing.T) {
		want := uint64(5934)
		if got := countFishesNoMemory(input, 80); got != want {
			t.Errorf("part1() = %v, want %v", got, want)
		}
	})
	t.Run("part 2", func(t *testing.T) {
		want := uint64(26984457539)
		if got := countFishesNoMemory(input, 256); got != want {
			t.Errorf("part2() = %v, want %v", got, want)
		}
	})
}
