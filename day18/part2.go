package day18

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func part2(filename string) int {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var (
		hex            string
		hash           int
		n              int64
		x0, y0, x1, y1 int
		boundary, area int
	)

	for scanner.Scan() {
		hex = strings.Fields(scanner.Text())[2]
		hash = strings.Index(hex, "#")
		n, _ = strconv.ParseInt("0x"+hex[hash+1:hash+6], 0, 0)

		switch hex[hash+6] {
		case '3':
			y1 += int(n)
		case '0':
			x1 += int(n)
		case '1':
			y1 -= int(n)
		case '2':
			x1 -= int(n)
		}

		// Compute the total length of the trench, alongside with the area
		// covered, using the shoelace formula (https://en.wikipedia.org/wiki/Shoelace_formula)
		boundary += int(n)
		area += x0*y1 - y0*x1

		x0, y0 = x1, y1
	}

	// Pick's theorem (https://en.wikipedia.org/wiki/Pick%27s_theorem)
	return (abs(area)/2 - boundary/2 + 1) + boundary
}
