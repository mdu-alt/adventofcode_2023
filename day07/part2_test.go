package day07

import "testing"

func TestPart2(t *testing.T) {
	testCases := []struct {
		filename string
		want     int
	}{
		{"example.txt", 5905},
		{"puzzle_input.txt", 249620106},
	}
	for _, tc := range testCases {
		t.Run(tc.filename, func(t *testing.T) {
			if got := part2(tc.filename); got != tc.want {
				t.Errorf("got %v; want %v", got, tc.want)
			}
		})
	}
}
