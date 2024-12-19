// Advent of Code 2024 - Day 19.
package day19

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func parseData(data []byte) ([]string, []string) {
	sections := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n\n")
	patterns := strings.Split(sections[0], ", ")
	designs := strings.Split(sections[1], "\n")
	return patterns, designs
}

func part1(patterns []string, designs []string) int {
	var valid int
	for _, design := range designs {
		if isValid(patterns, design) {
			valid++
		}
	}
	return valid
}

func isValid(patterns []string, design string) bool {
	if len(design) == 0 {
		return true
	}
	for _, p := range patterns {
		if strings.HasPrefix(design, p) {
			remainder := strings.TrimPrefix(design, p)
			if isValid(patterns, remainder) {
				return true
			}
		}
	}
	return false
}

func part2(patterns []string, designs []string) int {
	var validWays int
	for _, design := range designs {
		validWays += countValid(patterns, design)
	}
	return validWays
}

func countValid(patterns []string, design string) int {
	// Map the length of a design to the number of ways it can be made.
	designWays := map[int]int{
		0: 1, // There is one way to make a zero length design.
	}
	for i := 0; i < len(design); i++ {
		remainingDesign := design[i:]
		for _, pattern := range patterns {
			if strings.HasPrefix(remainingDesign, pattern) {
				lengthWithPattern := i + len(pattern)
				designWays[lengthWithPattern] += designWays[i]
			}
		}
	}
	return designWays[len(design)]
}

func Run(inputFile string) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	patterns, designs := parseData(data)

	p1 := part1(patterns, designs)
	fmt.Println("Part 1:", p1)

	p2 := part2(patterns, designs)
	fmt.Println("Part 2:", p2)
}
