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
	t.Run("part 2", func(t *testing.T) {
		want := 61229
		data := parseInput("./testinput.txt")
		if got := data.countSum(); got != want {
			t.Errorf("part2() = %v, want %v", got, want)
		}
	})
}
