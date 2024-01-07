package day07

import (
	"bufio"
	"cmp"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

func part2(filename string) int {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	var (
		scanner             = bufio.NewScanner(f)
		handBids            = make([]handBid, 0, 128)
		before, after, hand string
		bid                 int
	)

	for scanner.Scan() {
		before, after, _ = strings.Cut(scanner.Text(), " ")

		bid, _ = strconv.Atoi(after)
		hand = strings.Map(mapJoker, before)
		handBids = append(handBids, handBid{getTypeJoker(hand) + hand, bid})
	}

	slices.SortFunc(handBids, func(a, b handBid) int {
		return -cmp.Compare(a.typeAndHand, b.typeAndHand)
	})

	var total int
	for i, handBid := range handBids {
		total += (i + 1) * handBid.bid
	}

	return total
}

// -----------------------------------------------------------------------------

func mapJoker(r rune) rune {
	switch {
	case r == 'K':
		return 'B'
	case r == 'Q':
		return 'C'
	case r == 'T':
		return 'D'
	case unicode.IsDigit(r):
		return 'E' + '9' - r
	case r == 'J':
		return 'M'
	}
	return r
}

func getTypeJoker(hand string) string {
	m := make(map[rune]int)
	for _, r := range hand {
		m[r]++
	}

	switch len(m) {
	case 1:
		return "A"
	case 2:
		if _, ok := m['M']; ok {
			return "A"
		}
		r, _ := utf8.DecodeRuneInString(hand)
		if n := m[r]; n == 4 || n == 1 {
			return "B"
		}
		return "C"
	case 3:
		if n, ok := m['M']; ok {
			if n == 2 || n == 3 {
				return "B"
			}
			for _, r := range hand {
				if r == 'M' {
					continue
				}
				if x := m[r]; x == 1 || x == 3 {
					return "B"
				}
				return "C"
			}
		}
		for _, r := range hand {
			if m[r] == 1 {
				continue
			}
			if m[r] == 3 {
				return "D"
			}
			return "E"
		}
	case 4:
		if _, ok := m['M']; ok {
			return "D"
		}
		return "F"
	case 5:
		if _, ok := m['M']; ok {
			return "F"
		}
		return "G"
	}
	panic("len(m) is unexpected")
}
