package day06

import (
	"bufio"
	"log"
	"math"
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

	var (
		scanner = bufio.NewScanner(f)
		line    string
	)

	scanner.Scan()
	line = strings.Join(strings.Fields(scanner.Text())[1:], "")
	time, _ := strconv.ParseFloat(line, 64)

	scanner.Scan()
	line = strings.Join(strings.Fields(scanner.Text())[1:], "")
	dist, _ := strconv.ParseFloat(line, 64)

	delta := math.Pow(time, 2) - 4*dist
	r1 := (-time - math.Sqrt(delta)) / -2
	r2 := (-time + math.Sqrt(delta)) / -2

	return int(math.Ceil(r1) - math.Floor(r2) - 1)
}
