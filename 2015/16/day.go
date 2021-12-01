// Advent of Code 2015 - Day 166
package day16

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var detected = map[string]int{
	"children":    3,
	"cats":        7,
	"samoyeds":    2,
	"pomeranians": 3,
	"akitas":      0,
	"vizslas":     0,
	"goldfish":    5,
	"trees":       3,
	"cars":        2,
	"perfumes":    1,
}

// part1Check validates each Aunt against the detected compounds with the part 1 conditions
func part1Check(aunt int, compounds map[string]int) {
	for compound, quanity := range compounds {
		if quanity != detected[compound] {
			return
		}
	}
	fmt.Println("Part 1: Sue", aunt)
}

// part2Check validates each Aunt against the detected compounds with the part 2 conditions
func part2Check(aunt int, compounds map[string]int) {
	for compound, quantity := range compounds {
		switch compound {
		case "cats", "trees":
			if quantity <= detected[compound] {
				return
			}
		case "pomeranians", "goldfish":
			if quantity >= detected[compound] {
				return
			}
		default:
			if quantity != detected[compound] {
				return
			}
		}
	}
	fmt.Println("Part 2: Sue", aunt)
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(
		strings.TrimSuffix(string(data), "\n"), "\n",
	)

	for _, line := range lines {
		// Separate the Aunt from the compounds
		parts := strings.SplitN(line, ": ", 2)

		// Get the Aunt number
		aunt, err := strconv.Atoi(strings.Split(parts[0], " ")[1])
		if err != nil {
			log.Fatal(err)
		}

		// Get the compounds
		compounds := make(map[string]int)
		subParts := strings.Split(parts[1], ", ")
		for _, subPart := range subParts {
			subSubParts := strings.Split(subPart, ": ")
			if quantity, err := strconv.Atoi(subSubParts[1]); err == nil {
				compounds[subSubParts[0]] = quantity
			} else {
				log.Fatal(err)
			}
		}

		// Part 1
		part1Check(aunt, compounds)

		// Part 2
		part2Check(aunt, compounds)
	}
}
