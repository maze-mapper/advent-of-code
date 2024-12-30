// Advent of Code 2024 - Day 25.
package day25

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const height int = 5

func parseData(data []byte) ([][]int, [][]int) {
	var locks, keys [][]int
	sections := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n\n")
	for _, section := range sections {
		lines := strings.Split(section, "\n")
		isKey := false
		for _, r := range lines[0] {
			if r != '#' {
				isKey = true
				break
			}
		}
		profile := make([]int, len(lines[0]))
		for i := 1; i < len(lines)-1; i++ {
			for j, r := range lines[i] {
				if r == '#' {
					profile[j]++
				}
			}
		}
		if isKey {
			keys = append(keys, profile)
		} else {
			locks = append(locks, profile)
		}
	}
	return locks, keys
}

func part1(locks, keys [][]int) int {
	var count int
	for _, lock := range locks {
		for _, key := range keys {
			noOverlap := true
			for i := 0; i < len(lock); i++ {
				if lock[i]+key[i] > height {
					noOverlap = false
				}
			}
			if noOverlap {
				count++
			}
		}
	}
	return count
}

func Run(inputFile string) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	locks, keys := parseData(data)

	p1 := part1(locks, keys)
	fmt.Println("Part 1:", p1)
}
