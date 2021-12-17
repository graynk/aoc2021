package main

import (
	"testing"
)

func Test_day17(t *testing.T) {
	t.Run("parse test input", func(t *testing.T) {
		wantMinX, wantMaxX, wantMinY, wantMaxY := 20, 30, -10, -5
		minX, maxX, minY, maxY := parseInput("./testinput.txt")

		if minX != wantMinX {
			t.Errorf("part1() = %v, want %v", minX, wantMinX)
		}
		if maxX != wantMaxX {
			t.Errorf("part1() = %v, want %v", maxX, wantMaxX)
		}
		if minY != wantMinY {
			t.Errorf("part1() = %v, want %v", minY, wantMinY)
		}
		if maxY != wantMaxY {
			t.Errorf("part1() = %v, want %v", maxY, wantMaxY)
		}
	})
	t.Run("part 1", func(t *testing.T) {
		want := 45
		minX, maxX, minY, maxY := parseInput("./testinput.txt")

		if got, _ := probeLauncher(minX, maxX, minY, maxY); got != want {
			t.Errorf("part1() = %v, want %v", got, want)
		}
	})
	t.Run("part 2", func(t *testing.T) {
		want := 112
		minX, maxX, minY, maxY := parseInput("./testinput.txt")

		if _, got := probeLauncher(minX, maxX, minY, maxY); got != want {
			t.Errorf("part1() = %v, want %v", got, want)
		}
	})
}
