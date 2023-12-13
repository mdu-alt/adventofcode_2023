package day03

import (
	"bytes"
	"log"
	"os"
	"strconv"
)

func part2(filename string) int {
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}

	var (
		schematic = bytes.Fields(file)

		first, second, sum int
	)

	for i, line := range schematic {
		for j, b := range line {
			if b == '*' {
			out:
				for ii := i - 1; ii <= i+1; ii++ {
					if ii < 0 || ii > len(schematic)-1 {
						continue
					}
					for jj := j - 1; jj <= j+1; jj++ {
						if jj < 0 || jj > len(schematic[ii])-1 {
							continue
						}
						if isDigit(schematic[ii][jj]) {
							l, r := findNumber(jj, schematic[ii])
							if first == 0 {
								first, _ = strconv.Atoi(string(schematic[ii][l:r]))
								jj = r
							} else if second == 0 {
								second, _ = strconv.Atoi(string(schematic[ii][l:r]))
								sum += first * second
								break out
							}
						}
					}
				}
			}
			first, second = 0, 0
		}
	}

	return sum
}
