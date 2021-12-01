// Advent of Code 2015 - Day 20
package day20

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func part1(minPresents int) {
	// Each elf delivers 10x presents so remove this scale factor
	scaledMinPresents := minPresents / 10
	// Upper bound is the house number equal to the target number of presents
	houses := make([]int, scaledMinPresents)
	for elf := 1; elf <= scaledMinPresents; elf++ {
		// Update any houses visted by this elf
		for house := elf; house <= scaledMinPresents; house += elf {
			idx := house - 1
			houses[idx] += elf
			if houses[idx] >= scaledMinPresents {
				fmt.Println("Part 1: House", house, "is the first to get at least", minPresents, "presents")
				return
			}
		}
	}
}

func part2(minPresents int) {
	// Each elf delivers 11x presents so remove this scale factor and account for fraction
	scaledMinPresents := 1 + minPresents/11
	fmt.Println(scaledMinPresents)
	// Upper bound is the house number equal to the target number of presents
	houses := make([]int, scaledMinPresents)
	for elf := 1; elf <= scaledMinPresents; elf++ {
		// Update any houses visted by this elf
		for house := elf; house <= scaledMinPresents && house <= elf*50; house += elf {
			idx := house - 1
			houses[idx] += elf
			if houses[idx] >= scaledMinPresents {
				fmt.Println("Part 2: House", house, "is the first to get at least", minPresents, "presents")
				return
			}
		}
	}
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	presents, err := strconv.Atoi(
		strings.TrimSuffix(string(data), "\n"),
	)
	if err != nil {
		log.Fatal(err)
	}

	part1(presents)
	part2(presents)
}
