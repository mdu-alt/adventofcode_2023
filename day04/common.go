package day04

import "strconv"

func countMatches(winning []string, draw []string) int {
	var (
		set = make(map[int]struct{})

		n, matches int
	)

	for _, t := range winning {
		n, _ = strconv.Atoi(t)
		set[n] = struct{}{}
	}
	for _, t := range draw {
		n, _ = strconv.Atoi(t)
		if _, ok := set[n]; ok {
			matches++
		}
	}

	return matches
}
