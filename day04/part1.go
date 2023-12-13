package day04

import (
	"bufio"
	"log"
	"math"
	"os"
	"slices"
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

		tokens              []string
		index, matches, sum int
	)

	for scanner.Scan() {
		tokens = strings.Fields(scanner.Text())

		index = slices.Index(tokens, "|")
		matches = countMatches(tokens[2:index], tokens[index+1:])

		sum += int(math.Pow(2, float64(matches-1)))
	}

	return sum
}
