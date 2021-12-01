// Advent of Code 2018 - Day 2
package day2

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func parseInput(file string) []string {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(
		strings.TrimSuffix(string(data), "\n"), "\n",
	)

	return lines
}

// countChar counts the number of each character in a string
func countChar(s string) map[rune]int {
	counts := map[rune]int{}
	for _, r := range s {
		counts[r] += 1
	}
	return counts
}

// containsExactlyNSameCharacters reports if a map has a value of n for any of its keys
func containsExactlyNSameCharacters(m map[rune]int, n int) bool {
	for _, v := range m {
		if v == n {
			return true
		}
	}
	return false
}

func calculateChecksum(m map[int]int) int {
	checksum := 1
	for _, v := range m {
		checksum *= v
	}
	return checksum
}

func part1(ids []string) int {
	// Checksum requires counts for pairs and triplets
	components := map[int]int{
		2: 0,
		3: 0,
	}

	for _, id := range ids {
		counts := countChar(id)
		for k := range components {
			if containsExactlyNSameCharacters(counts, k) {
				components[k] += 1
			}
		}
	}

	return calculateChecksum(components)
}

// compare determines if two strings differ by only one character and returns the index of the differing character
func compare(s, ss string) (bool, int) {
	index := -1
	if len(s) != len(ss) {
		return false, index
	}
	differences := 0
	for i := 0; i < len(s); i++ {
		if s[i] != ss[i] {
			differences += 1
			index = i
		}
		if differences > 1 {
			return false, index
		}
	}
	return true, index
}

func part2(ids []string) string {
	length := len(ids)
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if found, index := compare(ids[i], ids[j]); found {
				return ids[i][:index] + ids[i][index+1:]
			}
		}
	}
	return ""
}

func Run(inputFile string) {
	ids := parseInput(inputFile)

	p1 := part1(ids)
	fmt.Println("Part 1:", p1)

	p2 := part2(ids)
	fmt.Println("Part 2:", p2)
}
