package main

import (
	"testing"
)

func Test_day10(t *testing.T) {
	input := parseInput("./testinput.txt")
	t.Run("part 1", func(t *testing.T) {
		want := 26397
		if got := findCorrupted(input); got != want {
			t.Errorf("part1() = %v, want %v", got, want)
		}
	})
	t.Run("part 2", func(t *testing.T) {
		want := 288957
		if got := completeIncompletes(input); got != want {
			t.Errorf("part1() = %v, want %v", got, want)
		}
	})
}
