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
		scanner = bufio.NewScanner(file)

		tokens, draws  []string
		color          string
		n, gameId, sum int
	)

	for scanner.Scan() {
		tokens = strings.FieldsFunc(scanner.Text(), isNotAlphanumeric)

		gameId, _ = strconv.Atoi(tokens[1])
		draws = tokens[2:]
		for i := 0; i < len(draws); i = i + 2 {
			n, _ = strconv.Atoi(draws[i])
			color = draws[i+1]

			if (color == "red" && n > 12) || (color == "green" && n > 13) || n > 14 {
				gameId = 0
			}
		}
		sum += gameId
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

		tokens, draws            []string
		color                    string
		n, maxR, maxG, maxB, sum int
	)

	for scanner.Scan() {
		tokens = strings.FieldsFunc(scanner.Text(), isNotAlphanumeric)

		draws = tokens[2:]
		for i := 0; i < len(draws); i = i + 2 {
			n, _ = strconv.Atoi(draws[i])
			color = draws[i+1]

			switch color {
			case "red":
				maxR = max(maxR, n)
			case "green":
				maxG = max(maxG, n)
			default:
				maxB = max(maxB, n)
			}
		}
		sum += maxR * maxG * maxB
		maxR, maxG, maxB = 0, 0, 0
	}

	return sum
}

func isNotAlphanumeric(r rune) bool {
	return !('0' <= r && r <= '9') && !('A' <= r && r <= 'Z') && !('a' <= r && r <= 'z')
}
