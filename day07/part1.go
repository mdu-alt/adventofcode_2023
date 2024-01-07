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

func part1(filename string) int {
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
		hand = strings.Map(mapNoJoker, before)
		handBids = append(handBids, handBid{getTypeNoJoker(hand) + hand, bid})
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

func mapNoJoker(r rune) rune {
	switch {
	case r == 'K':
		return 'B'
	case r == 'Q':
		return 'C'
	case r == 'J':
		return 'D'
	case r == 'T':
		return 'E'
	case unicode.IsDigit(r):
		return 'F' + '9' - r
	}
	return r
}

func getTypeNoJoker(hand string) string {
	m := make(map[rune]int)
	for _, r := range hand {
		m[r]++
	}

	switch len(m) {
	case 1:
		return "A"
	case 2:
		r, _ := utf8.DecodeRuneInString(hand)
		if n := m[r]; n == 4 || n == 1 {
			return "B"
		}
		return "C"
	case 3:
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
		return "F"
	case 5:
		return "G"
	}
	panic("len(m) is unexpected")
}
