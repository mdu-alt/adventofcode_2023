package day06

import (
	"bufio"
	"log"
	"math"
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

	scanner.Scan()
	times := strings.Fields(scanner.Text())[1:]

	scanner.Scan()
	dists := strings.Fields(scanner.Text())[1:]

	var (
		t, d  float64
		races []race
	)
	for i := range times {
		t, _ = strconv.ParseFloat(times[i], 64)
		d, _ = strconv.ParseFloat(dists[i], 64)
		races = append(races, race{t, d})
	}

	var (
		delta, r1, r2 float64
		ways          = 1
	)
	for _, race := range races {
		delta = math.Pow(race.time, 2) - 4*race.dist
		r1 = (-race.time - math.Sqrt(delta)) / -2
		r2 = (-race.time + math.Sqrt(delta)) / -2

		ways *= int(math.Ceil(r1) - math.Floor(r2) - 1)
	}

	return ways
}
