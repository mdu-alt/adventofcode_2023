package day02

import (
	"bufio"
	"log"
	"os"
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
