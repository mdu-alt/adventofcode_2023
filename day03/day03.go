// Package day03 implements Day 3 of Advent of Code, 2023: [Gear Ratios].
//
// [Gear Ratios]: https://adventofcode.com/2023/day/3
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
								first = toInt(schematic[ii][l:r])
								jj = r
							} else if second == 0 {
								second = toInt(schematic[ii][l:r])
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

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func toInt(line []byte) int {
	n := 0
	for _, b := range line {
		n = n*10 + int(b-'0')
	}
	return n
}

func findNumber(i int, line []byte) (int, int) {
	l, r := i, i
	for l >= 0 && isDigit(line[l]) {
		l--
	}
	for r < len(line) && isDigit(line[r]) {
		r++
	}
	return l + 1, r
}
