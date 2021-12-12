package main

import (
	"testing"
)

func Test_day8(t *testing.T) {
	t.Run("part 1", func(t *testing.T) {
		want := 26
		data := parseInput("./testinput.txt")
		if got := data.countEasyNumbers(); got != want {
			t.Errorf("part1() = %v, want %v", got, want)
		}
	})
	//t.Run("part 2", func(t *testing.T) {
	//	want := 168
	//	input := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
	//	if got := cheapestMoveProgressive(input); got != want {
	//		t.Errorf("part2() = %v, want %v", got, want)
	//	}
	//})
}
