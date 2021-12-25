package main

import (
	"fmt"
	"testing"
)

func Test_day25(t *testing.T) {
	t.Run("day 25", func(t *testing.T) {
		want := 58
		field := parseInput("./testinput.txt")
		fmt.Println(field.String())
		if got := field.StepTillWeStop(); got != want {
			t.Errorf("day 25 = %v, want %v", got, want)
		}
	})
}
