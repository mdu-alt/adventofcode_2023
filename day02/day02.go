// Package day02 implements Day 2 of Advent of Code, 2023: [Cube Conundrum].
//
// [Cube Conundrum]: https://adventofcode.com/2023/day/2
package day02

import (
	"bufio"
	"log"
	"os"
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
		scanner    = bufio.NewScanner(file)
		isPossible = true

		tokens      []string
		i, n        int
		gameId, sum int
	)

	for scanner.Scan() {
		tokens = strings.Fields(scanner.Text())

		gameId, _ = strconv.Atoi(strings.Split(tokens[1], ":")[0])

		for _, token := range tokens[2:] {
			if i, err = strconv.Atoi(token); err != nil {
				if (strings.HasPrefix(token, "red") && n > 12) || (strings.HasPrefix(token, "green") && n > 13) || n > 14 {
					isPossible = false
					break
				}
			} else {
				n = i
			}
		}

		if isPossible {
			sum += gameId
		}
		isPossible = true
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
		scanner = bufio.NewScanner(file)

		tokens                []string
		i, n                  int
		maxR, maxG, maxB, sum int
	)

	for scanner.Scan() {
		tokens = strings.Fields(scanner.Text())

		for _, token := range tokens[2:] {
			if i, err = strconv.Atoi(token); err != nil {
				switch {
				case strings.HasPrefix(token, "red"):
					maxR = max(maxR, n)
				case strings.HasPrefix(token, "green"):
					maxG = max(maxG, n)
				default:
					maxB = max(maxB, n)
				}
			} else {
				n = i
			}
		}

		sum += maxR * maxG * maxB
		maxR, maxG, maxB = 0, 0, 0
	}

	return sum
}
