// Advent of Code 2023 - Day 4.
package day4

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type scratchcard struct {
	winningNumbers, haveNumbers map[int]bool
}

func (sc scratchcard) matches() int {
	total := 0
	for n := range sc.haveNumbers {
		if sc.winningNumbers[n] {
			total += 1
		}
	}
	return total
}

func (sc scratchcard) points() int {
	matches := sc.matches()
	switch matches {
	case 0:
		return 0
	default:
		return int(math.Pow(2, float64(matches-1)))
	}
}

func parseData(data []byte) []scratchcard {
	lines := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
	cards := make([]scratchcard, len(lines))
	for i, line := range lines {
		card := scratchcard{winningNumbers: map[int]bool{}, haveNumbers: map[int]bool{}}
		line = strings.ReplaceAll(line, "  ", " ")
		parts := strings.Split(strings.Split(line, ": ")[1], " | ")

		for _, s := range strings.Split(parts[0], " ") {
			n, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}
			card.winningNumbers[n] = true
		}
		for _, s := range strings.Split(parts[1], " ") {
			n, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}
			card.haveNumbers[n] = true
		}
		cards[i] = card
	}
	return cards
}

func part1(cards []scratchcard) int {
	total := 0
	for _, card := range cards {
		total += card.points()
	}
	return total
}

func part2(cards []scratchcard) int {
	counts := make([]int, len(cards))
	for i, card := range cards {
		counts[i] += 1
		n := card.matches()
		for j := 1; i+j < len(counts) && j <= n; j++ {
			counts[i+j] += counts[i]
		}
		fmt.Println(counts)
	}
	total := 0
	for _, c := range counts {
		total += c
	}
	return total
}

func Run(inputFile string) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	cards := parseData(data)

	p1 := part1(cards)
	fmt.Println("Part 1:", p1)

	p2 := part2(cards)
	fmt.Println("Part 2:", p2)
}
