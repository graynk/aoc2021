package main

import (
	"testing"
)

func Test_day15(t *testing.T) {
	t.Run("part 1", func(t *testing.T) {
		want := 40
		rt := parseInput("./testinput.txt", 10, 10)

		if got := rt.traverse(); got != want {
			t.Errorf("part1() = %v, want %v", got, want)
		}
	})
	//t.Run("part 2", func(t *testing.T) {
	//	var want int64 = 2188189693529
	//	pi := parseInput("./testinput.txt")
	//	for i := 0; i < 40; i++ {
	//		pi.insertPolymers()
	//	}
	//	if got := pi.commonCounter(); got != want {
	//		t.Errorf("part2() = %v, want %v", got, want)
	//	}
	//})
}
