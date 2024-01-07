package day08

import (
	"bufio"
	"log"
	"os"
	"strings"
	"unicode"
)

func part2(filename string) int {
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
		starts  = make([]string, 0, 64)
		nodes   []string
	)
	for scanner.Scan() {
		nodes = strings.FieldsFunc(scanner.Text(), func(r rune) bool {
			return !unicode.IsLetter(r) && !unicode.IsDigit(r)
		})
		network[nodes[0]] = pair{nodes[1], nodes[2]}
		if strings.HasSuffix(nodes[0], "A") {
			starts = append(starts, nodes[0])
		}
	}

	steps, totalSteps := 0, 1
	for _, node := range starts {
		for !strings.HasSuffix(node, "Z") {
			if instructions[steps%len(instructions)] == 'L' {
				node = network[node].l
			} else {
				node = network[node].r
			}
			steps++
		}
		totalSteps = lcm(totalSteps, steps)
		steps = 0
	}

	return totalSteps
}

// -----------------------------------------------------------------------------

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}
