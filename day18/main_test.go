package main

import (
	"testing"
)

func Test_day18(t *testing.T) {
	t.Run("part 1 test sum simplest", func(t *testing.T) {
		want := "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"
		numbers := parseInput("./testinput1.txt")

		if got := sumNumbers(numbers).String(); got != want {
			t.Errorf("part1() = %v, want %v", got, want)
		}
	})
	t.Run("part 1 test sum", func(t *testing.T) {
		want := "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]"
		numbers := parseInput("./testinput2.txt")

		if got := sumNumbers(numbers).String(); got != want {
			t.Errorf("part1() = %v, want %v", got, want)
		}
	})
	t.Run("part 1 test another sum", func(t *testing.T) {
		want := "[[[[6,6],[7,6]],[[7,7],[7,0]]],[[[7,7],[7,7]],[[7,8],[9,9]]]]"
		numbers := parseInput("./testinput3.txt")

		if got := sumNumbers(numbers).String(); got != want {
			t.Errorf("part1() = %v, want %v", got, want)
		}
	})
	t.Run("part 1 test magnitude", func(t *testing.T) {
		want := 4140
		numbers := parseInput("./testinput3.txt")

		if got := sumNumbers(numbers).magnitude(); got != want {
			t.Errorf("part1() = %v, want %v", got, want)
		}
	})
	t.Run("part 2 max magnitude", func(t *testing.T) {
		want := 3993
		numbers := parseInput("./testinput3.txt")

		if got := maxSumMagnitude(numbers); got != want {
			t.Errorf("part2() = %v, want %v", got, want)
		}
	})
}
