package day01

import (
	"bufio"
	"log"
	"os"
	"strings"
	"unicode"
)

func part2(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	var (
		scanner = bufio.NewScanner(file)
		digits  = map[string]string{
			"one":   "o1e",
			"two":   "t2o",
			"three": "t3e",
			"four":  "f4r",
			"five":  "f5e",
			"six":   "s6x",
			"seven": "s7n",
			"eight": "e8t",
			"nine":  "n9e",
		}

		line      string
		l, r, sum int
	)

	for scanner.Scan() {
		line = scanner.Text()

		for k, v := range digits {
			line = strings.ReplaceAll(line, k, v)
		}

		l = strings.IndexFunc(line, unicode.IsDigit)
		r = strings.LastIndexFunc(line, unicode.IsDigit)

		sum += int(line[l]-'0')*10 + int(line[r]-'0')
	}

	return sum
}
