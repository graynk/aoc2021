package main

import (
	"fmt"
	"testing"
)

func Test_day12(t *testing.T) {
	t.Run("part 1", func(t *testing.T) {
		want := 10
		caveMap := parseInput("./testinput.txt")
		paths := make([]Path, 0, 1)
		paths = caveMap["start"].Explore(make(Path, 0, 1), paths)
		for _, path := range paths {
			fmt.Println(path)
		}

		if got := len(paths); got != want {
			t.Errorf("part1() = %v, want %v", got, want)
		}
	})
}
