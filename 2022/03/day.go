// Advent of Code 2022 - Day 3.
package day3

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
)

func parseInput(data []byte) [][2][]byte {
	lines := bytes.Split(bytes.TrimSuffix(data, []byte("\n")), []byte("\n"))
	rucksacks := make([][2][]byte, len(lines))
	for i, line := range lines {
		if len(line)%2 != 0 {
			log.Fatalf("Line number %d does not contain an even number of characters: %v", i, line)
		}
		m := len(line) / 2
		rucksacks[i] = [2][]byte{line[:m], line[m:]}
	}
	return rucksacks
}

// itemScore returns the score for an item.
func itemScore(item byte) byte {
	if item >= 97 {
		// Lowercase.
		return item - 96
	}
	// Uppercase.
	return item - 38
}

// findCommonItem finds the item that appears in both compartments of a rucksack.
func findCommonItem(items [2][]byte) byte {
	allocations := map[byte]int{}
	numItems := len(items[0])
	for i := 0; i < numItems; i++ {
		for c := 0; c < 2; c++ {
			obj := items[c][i]
			if n, ok := allocations[obj]; ok {
				if n != c {
					return obj
				}
			} else {
				allocations[obj] = c
			}
		}
	}
	return byte(0)
}

// findBadge finds the item that appears in all provided rucksacks.
func findBadge(rucksacks [][2][]byte) byte {
	allocations := map[byte]uint8{}
	for i, rucksack := range rucksacks {
		all := append(rucksack[0], rucksack[1]...)
		flag := uint8(1) << i
		for _, item := range all {
			allocations[item] |= flag
		}
	}
	numRucksacks := len(rucksacks)
	goal := uint8(0)
	for i := 0; i < numRucksacks; i++ {
		goal |= (uint8(1) << i)
	}
	for k, v := range allocations {
		if v == goal {
			return k
		}
	}
	return byte(0)
}

func part1(rucksacks [][2][]byte) int {
	total := 0
	for _, rucksack := range rucksacks {
		common := findCommonItem(rucksack)
		total += int(itemScore(common))
	}
	return total
}

func part2(rucksacks [][2][]byte) int {
	if len(rucksacks)%3 != 0 {
		log.Fatal("Cannot separate %d rucksacks in to groups of three", len(rucksacks))
	}
	total := 0
	for i := 0; i < len(rucksacks); i += 3 {
		badge := findBadge(rucksacks[i : i+3])
		total += int(itemScore(badge))
	}
	return total
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	rucksacks := parseInput(data)

	p1 := part1(rucksacks)
	fmt.Println("Part 1:", p1)

	p2 := part2(rucksacks)
	fmt.Println("Part 2:", p2)
}
