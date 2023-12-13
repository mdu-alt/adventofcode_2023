package day03

import (
	"bytes"
	"log"
	"os"
)

func part1(filename string) int {
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}

	var (
		schematic = bytes.Fields(file)

		n, size, sum int
	)

	for i, line := range schematic {
		for j, b := range line {
			if isDigit(b) {
				n = n*10 + int(b-'0')
				size++

				if j == len(line)-1 || !isDigit(line[j+1]) {
				out:
					for ii := i - 1; ii <= i+1; ii++ {
						if ii < 0 || ii > len(schematic)-1 {
							continue
						}
						for jj := j - size; jj <= j+1; jj++ {
							if jj < 0 || jj > len(schematic[ii])-1 {
								continue
							}
							if surround := schematic[ii][jj]; !isDigit(surround) && surround != '.' {
								sum += n
								break out
							}
						}
					}

					n, size = 0, 0
				}
			}
		}
	}

	return sum
}
