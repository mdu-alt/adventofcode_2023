package day08

import (
	"bufio"
	"log"
	"os"
	"strings"
	"unicode"
)

func part1(filename string) int {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	scanner.Scan()
	instructions := scanner.Text()
	scanner.Scan()

	var (
		network = make(map[string]pair)
		nodes   []string
	)
	for scanner.Scan() {
		nodes = strings.FieldsFunc(scanner.Text(), func(r rune) bool {
			return !unicode.IsLetter(r)
		})
		network[nodes[0]] = pair{nodes[1], nodes[2]}
	}

	var (
		node  = "AAA"
		steps int
	)
	for node != "ZZZ" {
		if instructions[steps%len(instructions)] == 'L' {
			node = network[node].l
		} else {
			node = network[node].r
		}
		steps++
	}

	return steps
}
