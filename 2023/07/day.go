// Advent of Code 2023 - Day 7.
package day7

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type hand struct {
	cards    [5]int
	bid      int
	handType int
}

func (h *hand) initHandType() {
	m := map[int]int{}
	for _, c := range h.cards {
		m[c] += 1
	}
	var counts []int
	for _, v := range m {
		counts = append(counts, v)
	}
	switch {
	case len(counts) == 1:
		h.handType = fiveOfAKind

	case len(counts) == 2 && (counts[0] == 4 || counts[1] == 4):
		h.handType = fourOfAKind

	case len(counts) == 2 && (counts[0] == 3 || counts[1] == 3):
		h.handType = fullHouse

	case len(counts) == 3 && (counts[0] == 3 || counts[1] == 3 || counts[2] == 3):
		h.handType = threeOfAKind

	case len(counts) == 3:
		h.handType = twoPair

	case len(counts) == 4:
		h.handType = onePair

	case len(counts) == 5:
		h.handType = highCard
	}
}

func (h *hand) jackToJoker() {
	for i, c := range h.cards {
		if c == 11 {
			h.cards[i] = 1
		}
	}
	h.setJokerHandType()
}

func (h *hand) setJokerHandType() {
	m := map[int]int{}
	for _, c := range h.cards {
		m[c] += 1
	}
	var counts []int
	for _, v := range m {
		counts = append(counts, v)
	}
	jokers := m[1]

	switch {
	case len(counts) == 1:
		h.handType = fiveOfAKind

	case len(counts) == 2 && (counts[0] == 4 || counts[1] == 4):
		switch jokers {
		case 1, 4:
			h.handType = fiveOfAKind
		default:
			h.handType = fourOfAKind
		}

	case len(counts) == 2 && (counts[0] == 3 || counts[1] == 3):
		switch jokers {
		case 2, 3:
			h.handType = fiveOfAKind
		default:
			h.handType = fullHouse
		}

	case len(counts) == 3 && (counts[0] == 3 || counts[1] == 3 || counts[2] == 3):
		switch jokers {
		case 1, 3:
			h.handType = fourOfAKind
		default:
			h.handType = threeOfAKind
		}

	case len(counts) == 3:
		switch jokers {
		case 2:
			h.handType = fourOfAKind
		case 1:
			h.handType = fullHouse
		default:
			h.handType = twoPair
		}

	case len(counts) == 4:
		switch jokers {
		case 1, 2:
			h.handType = threeOfAKind
		default:
			h.handType = onePair
		}

	case len(counts) == 5:
		switch jokers {
		case 1:
			h.handType = onePair
		default:
			h.handType = highCard
		}
	}
}

// Hand types in ascending strength.
const (
	highCard = iota
	onePair
	twoPair
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

type handSorter []hand

func (hs handSorter) Len() int {
	return len(hs)
}

func (hs handSorter) Swap(i, j int) {
	hs[i], hs[j] = hs[j], hs[i]
}

func (hs handSorter) Less(i, j int) bool {
	if hs[i].handType == hs[j].handType {
		for k := 0; k < 5; k++ {
			if hs[i].cards[k] != hs[j].cards[k] {
				return hs[i].cards[k] < hs[j].cards[k]
			}
		}
		return false
	}
	return hs[i].handType < hs[j].handType
}

func cardToInt(r rune) (int, error) {
	switch r {
	case 'A':
		return 14, nil
	case 'K':
		return 13, nil
	case 'Q':
		return 12, nil
	case 'J':
		return 11, nil
	case 'T':
		return 10, nil
	default:
		n, err := strconv.Atoi(string(r))
		if err != nil {
			return 0, err
		}
		return n, nil
	}
}

func parseData(data []byte) []hand {
	lines := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
	hands := make([]hand, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " ")
		cards := [5]int{}
		for i, r := range parts[0] {
			n, err := cardToInt(r)
			if err != nil {
				log.Fatal(err)
			}
			cards[i] = n
		}
		bid, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}
		h := hand{
			cards: cards,
			bid:   bid,
		}
		h.initHandType()
		hands[i] = h
	}
	return hands
}

func part1(hands handSorter) int {
	sort.Sort(hands)
	total := 0
	for i, h := range hands {
		total += (i + 1) * h.bid
	}
	return total
}

func part2(hands handSorter) int {
	for i, h := range hands {
		h.jackToJoker()
		hands[i] = h
	}
	sort.Sort(hands)
	total := 0
	for i, h := range hands {
		total += (i + 1) * h.bid
	}
	return total
}

func Run(inputFile string) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	hands := parseData(data)

	p1 := part1(hands)
	fmt.Println("Part 1:", p1)

	p2 := part2(hands)
	fmt.Println("Part 2:", p2)
}
