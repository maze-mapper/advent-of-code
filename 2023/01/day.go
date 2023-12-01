// Advent of Code 2023 - Day 1.
package day1

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const alpha = "abcdefghijklmnopqrstuvwxyz"

func part1(lines []string) int {
	sum := 0
	for _, l := range lines {
		s := strings.Trim(l, alpha)
		chars := strings.Split(s, "")
		n, err := strconv.Atoi(chars[0] + chars[len(chars)-1])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Val:", n)
		sum += n
	}
	return sum
}

func part2(lines []string) int {
	textToDigits := map[string]int{
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	lengths := []int{1, 3, 4, 5}
	sum := 0

	for _, s := range lines {
		var first, last int

	loopFirst:
		for i := range s {
			for _, l := range lengths {
				if i+l <= len(s) {
					text := s[i : i+l]
					if n, ok := textToDigits[text]; ok {
						first = n
						break loopFirst
					}
				}
			}
		}

	loopLast:
		for i := len(s); i >= 0; i-- {
			for _, l := range lengths {
				if i-l >= 0 {
					text := s[i-l : i]
					if n, ok := textToDigits[text]; ok {
						last = n
						break loopLast
					}
				}
			}
		}

		sum += first*10 + last
	}

	return sum
}

func Run(inputFile string) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")

	p1 := part1(lines)
	fmt.Println("Part 1:", p1)

	p2 := part2(lines)
	fmt.Println("Part 2:", p2)
}
