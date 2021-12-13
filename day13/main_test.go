package main

import (
	"testing"
)

func Test_day13(t *testing.T) {
	input, instructions := parseInput("./testinput.txt")
	t.Run("part 1", func(t *testing.T) {
		want := 17
		newPf := input.foldInstruction(instructions[0])
		if got := newPf.countDots(); got != want {
			t.Errorf("part1() = %v, want %v", got, want)
		}
		wantField := `
#####
#...#
#...#
#...#
#####
.....
.....
`
		if got := newPf.foldInstruction(instructions[1]).String(); got != wantField[1:] {
			t.Errorf("part1() = %v, want %v", got, wantField[1:])
		}
	})
}
