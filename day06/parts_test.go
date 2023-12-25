package day06

import "testing"

const (
	example     = "example.txt"
	puzzleInput = "puzzle_input.txt"
)

type testCase struct {
	filename string
	want     int
}

func TestPart1(t *testing.T) {
	testCases := []testCase{
		{example, 288},
		{puzzleInput, 3316275},
	}

	test(t, testCases, part1)
}

func TestPart2(t *testing.T) {
	testCases := []testCase{
		{example, 71503},
		{puzzleInput, 27102791},
	}

	test(t, testCases, part2)
}

// -----------------------------------------------------------------------------

func test(t *testing.T, testCases []testCase, part func(string) int) {
	for _, tc := range testCases {
		t.Run(tc.filename, func(t *testing.T) {
			if got := part(tc.filename); got != tc.want {
				t.Errorf("got %v; want %v", got, tc.want)
			}
		})
	}
}
