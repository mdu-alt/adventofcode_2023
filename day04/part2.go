package day04

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func part2(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	var (
		scanner      = bufio.NewScanner(file)
		scratchcards = make(map[int]int)

		tokens                      []string
		index, matches, cardId, sum int
	)

	for scanner.Scan() {
		tokens = strings.Fields(scanner.Text())

		index = slices.Index(tokens, "|")
		matches = countMatches(tokens[2:index], tokens[index+1:])
		cardId, _ = strconv.Atoi(tokens[1][:len(tokens[1])-1])

		scratchcards[cardId]++
		for i := 1; i <= matches; i++ {
			scratchcards[cardId+i] += scratchcards[cardId]
		}
		sum += scratchcards[cardId]
	}

	return sum
}
