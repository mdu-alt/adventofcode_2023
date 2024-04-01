package day18

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1(filename string) int {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var (
		tokens            []string
		n, x0, y0, x1, y1 int
		boundary, area    int
	)

	for scanner.Scan() {
		tokens = strings.Fields(scanner.Text())
		n, _ = strconv.Atoi(tokens[1])

		switch tokens[0][0] {
		case 'U':
			y1 += n
		case 'R':
			x1 += n
		case 'D':
			y1 -= n
		case 'L':
			x1 -= n
		}

		// Compute the total length of the trench, alongside with the area
		// covered, using the shoelace formula (https://en.wikipedia.org/wiki/Shoelace_formula)
		boundary += n
		area += x0*y1 - y0*x1

		x0, y0 = x1, y1
	}

	// Pick's theorem (https://en.wikipedia.org/wiki/Pick%27s_theorem)
	return (abs(area)/2 - boundary/2 + 1) + boundary
}
