// Package day04 implements Day 4 of Advent of Code, 2023: [Scratchcards].
// [Scratchcards]: https://adventofcode.com/2023/day/4
package day04

import (
	"bufio"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func part1(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	var (
		scanner = bufio.NewScanner(file)

		tokens      []string
		points, sum int
	)

	for scanner.Scan() {
		tokens = strings.Fields(scanner.Text())

		points = countMatches(tokens, slices.Index(tokens, "|"))
		sum += int(math.Pow(2, float64(points-1)))
	}

	return sum
}

func part2(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	var (
		scanner      = bufio.NewScanner(file)
		scratchcards = make(map[int]int)

		tokens                  []string
		matches, cardId, i, sum int
	)

	for scanner.Scan() {
		tokens = strings.Fields(scanner.Text())

		matches = countMatches(tokens, slices.Index(tokens, "|"))

		cardId, _ = strconv.Atoi(tokens[1][:len(tokens[1])-1])

		scratchcards[cardId]++
		for i = 1; i <= matches; i++ {
			scratchcards[cardId+i] += scratchcards[cardId]
		}
		sum += scratchcards[cardId]
	}

	return sum
}

func countMatches(tokens []string, i int) int {
	var (
		m = make(map[int]struct{})

		n, matches int
	)

	for _, s := range tokens[2:i] {
		n, _ = strconv.Atoi(s)
		m[n] = struct{}{}
	}
	for _, s := range tokens[i+1:] {
		n, _ = strconv.Atoi(s)
		if _, ok := m[n]; ok {
			matches++
		}
	}

	return matches
}
