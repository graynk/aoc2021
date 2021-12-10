package main

import (
	"testing"
)

func Test_day7(t *testing.T) {
	t.Run("part 1", func(t *testing.T) {
		want := 37
		input := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
		if got := cheapestMove(input); got != want {
			t.Errorf("part1() = %v, want %v", got, want)
		}
	})
	t.Run("part 2", func(t *testing.T) {
		want := 168
		input := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
		if got := cheapestMoveProgressive(input); got != want {
			t.Errorf("part2() = %v, want %v", got, want)
		}
	})
}

func Test_countFuel(t *testing.T) {
	type args struct {
		from int
		to   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"whaterver",
			args{
				from: 16,
				to:   5,
			},
			66,
		},
		{
			"whaterver",
			args{
				from: 1,
				to:   5,
			},
			10,
		},
		{
			"whaterver",
			args{
				from: 2,
				to:   5,
			},
			6,
		},
		{
			"whaterver",
			args{
				from: 0,
				to:   5,
			},
			15,
		},
		{
			"whaterver",
			args{
				from: 4,
				to:   5,
			},
			1,
		},
		{
			"whaterver",
			args{
				from: 7,
				to:   5,
			},
			3,
		},
		{
			"whaterver",
			args{
				from: 1,
				to:   5,
			},
			10,
		},
		{
			"whaterver",
			args{
				from: 14,
				to:   5,
			},
			45,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countFuel(tt.args.from, tt.args.to); got != tt.want {
				t.Errorf("countFuel() = %v, want %v", got, tt.want)
			}
		})
	}
}
