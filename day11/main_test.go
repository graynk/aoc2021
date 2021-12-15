package main

import (
	"testing"
)

func Test_day11(t *testing.T) {
	t.Run("part 1", func(t *testing.T) {
		want := 1656
		field := parseInput("./testinput.txt")
		if got := field.iWouldStepOneHundredSteps(); got != want {
			t.Errorf("part1() = %v, want %v", got, want)
		}
	})
	t.Run("part 1", func(t *testing.T) {
		want := 195
		field := parseInput("./testinput.txt")
		if got := field.andIWouldStepOneHundredMore(); got != want {
			t.Errorf("part1() = %v, want %v", got, want)
		}
	})
}
