package day01

import (
	"testing"
)

const (
	Example1    = "example_1.txt"
	Example2    = "example_2.txt"
	PuzzleInput = "puzzle_input.txt"
)

func TestPart1(t *testing.T) {
	testCases := []struct {
		filename string
		want     int
	}{
		{Example1, 142},
		{PuzzleInput, 55538},
	}

	for _, tc := range testCases {
		t.Run(tc.filename, func(t *testing.T) {
			if got := part1(tc.filename); got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	testCases := []struct {
		filename string
		want     int
	}{
		{Example2, 281},
		{PuzzleInput, 54875},
	}

	for _, tc := range testCases {
		t.Run(tc.filename, func(t *testing.T) {
			if got := part2(tc.filename); got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
