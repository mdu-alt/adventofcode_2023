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

		line      string
		l, r, sum int
	)

	for scanner.Scan() {
		line = scanner.Text()
		l = strings.IndexFunc(line, unicode.IsDigit)
		r = strings.LastIndexFunc(line, unicode.IsDigit)

		sum += int(line[l]-'0')*10 + int(line[r]-'0')
	}

	return sum
}
