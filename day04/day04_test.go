package day04

import (
	"testing"
)

const (
	Example     = "example.txt"
	PuzzleInput = "puzzle_input.txt"
)

func TestPart1(t *testing.T) {
	testCases := []struct {
		filename string
		want     int
	}{
		{Example, 13},
		{PuzzleInput, 24706},
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
		{Example, 30},
		{PuzzleInput, 13114317},
	}

	for _, tc := range testCases {
		t.Run(tc.filename, func(t *testing.T) {
			if got := part2(tc.filename); got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
