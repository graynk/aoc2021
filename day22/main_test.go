package main

import (
	"testing"
)

func Test_day22(t *testing.T) {
	t.Run("part 1", func(t *testing.T) {
		want := 39
		r := reboot("./testinput.txt", 100)

		if got := r.countOn(); got != want {
			t.Errorf("part1() = %v, want %v", got, want)
		}
	})
	t.Run("part 1 2", func(t *testing.T) {
		want := 590784
		r := reboot("./testinput2.txt", 100)

		if got := r.countOn(); got != want {
			t.Errorf("part1() = %v, want %v", got, want)
		}
	})
	t.Run("part 2", func(t *testing.T) {
		want := 2758514936282235
		r := reboot("./testinput3.txt", 105000)

		if got := r.countOn(); got != want {
			t.Errorf("part2() = %v, want %v", got, want)
		}
	})
}
