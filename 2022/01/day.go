// Advent of Code 2022 - Day 1.
package day1

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

// Parse calories returns the total number of calories carried by each elf.
func parseCalories(data []byte) []int {
	groups := strings.Split(
                strings.TrimSuffix(string(data), "\n"), "\n\n",
        )
	elves := make([]int, len(groups))
	for i, group := range groups {
		total := 0
		for _, line := range strings.Split(strings.TrimSuffix(group, "\n"), "\n") {
			if calories, err := strconv.Atoi(strings.TrimSuffix(string(line), "\n")); err == nil {
				total += calories
			} else {
				log.Fatal(err)
			}
		}
		elves[i] = total
	}
	return elves
}

func part1(calories []int) int {
	return calories[len(calories) - 1]
}

func part2(calories []int) int {
	total := 0
	for i := len(calories) - 3; i < len(calories); i++ {
		total += calories[i]
	}
	return total
}

func Run(inputFile string) {
        data, err := ioutil.ReadFile(inputFile)
        if err != nil {
                log.Fatal(err)
        }
        calories := parseCalories(data)
	sort.Ints(calories)

        p1 := part1(calories)
        fmt.Println("Part 1:", p1)

        p2 := part2(calories)
        fmt.Println("Part 2:", p2)
}

