package main

import (
	"testing"
)

func Test_day20(t *testing.T) {
	t.Run("part 1", func(t *testing.T) {
		want := 35
		enhancement, image := FromFile("./testinput1.txt")

		if got := iterate(enhancement, image, 2); got != want {
			t.Errorf("part1() = %v, want %v", got, want)
		}
	})

	t.Run("part 2", func(t *testing.T) {
		want := 3351
		enhancement, image := FromFile("./testinput1.txt")

		if got := iterate(enhancement, image, 50); got != want {
			t.Errorf("part2() = %v, want %v", got, want)
		}
	})
}
