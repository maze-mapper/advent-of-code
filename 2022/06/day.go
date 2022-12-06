// Advent of Code 2022 - Day 6.
package day6

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// findNDistinctCharacters returns the first index at which n characters in the string are distinct.
func findNDistinctCharacters(s string, n int) int {
	chars := map[byte]int{}
	i := 0
	for ; i < n; i++ {
		chars[s[i]] += 1
	}
	for len(chars) < n && i < len(s) {
		oldChar := s[i-n]
		chars[oldChar] -= 1
		if chars[oldChar] == 0 {
			delete(chars, oldChar)
		}
		chars[s[i]] += 1
		i += 1
	}
	return i
}

func part1(datastream string) int {
	return findNDistinctCharacters(datastream, 4)
}

func part2(datastream string) int {
	return findNDistinctCharacters(datastream, 14)
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	datastream := strings.TrimSuffix(string(data), "\n")

	p1 := part1(datastream)
	fmt.Println("Part 1:", p1)

	p2 := part2(datastream)
	fmt.Println("Part 2:", p2)
}
