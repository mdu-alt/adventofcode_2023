// Package day01 implements Day 1 of Advent of Code, 2021: [Sonar Sweep].
// [Sonar Sweep]: https://adventofcode.com/2021/day/1
package day01

import (
	"bufio"
	"log"
	"os"
	"strings"
	"unicode"
)

func part1(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	var (
		scanner = bufio.NewScanner(file)

		first, last int
		sum         int
	)

	for scanner.Scan() {
		first, last = 0, 0
		for _, r := range scanner.Text() {
			if unicode.IsDigit(r) {
				if first == 0 {
					first = int(r - '0')
				}
				last = int(r - '0')
			}
		}
		sum += first*10 + last
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
		digits  = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
		acc     = make([]int, 0)

		line string
		sum  int
	)

	for scanner.Scan() {
		line = scanner.Text()

		for i, r := range line {
			if unicode.IsDigit(r) {
				acc = append(acc, int(r-'0'))
			} else {
				for j, digit := range digits {
					if strings.HasPrefix(line[i:], digit) {
						acc = append(acc, j+1)
					}
				}
			}
		}

		sum += acc[0]*10 + acc[len(acc)-1]
		clear(acc)
	}

	return sum
}
