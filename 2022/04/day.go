// Advent of Code 2022 - Day 4.
package day4

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// interval holds the lower and upper bounds for assigned areas.
type interval struct {
	lower, upper int
}

// contains returns true if the interval j is contained by the interval i.
func (i interval) contains(j interval) bool {
	return j.lower >= i.lower && j.upper <= i.upper
}

// overlaps returns true if interval j overlaps with interval i.
func (i interval) overlaps(j interval) bool {
	return i.upper >= j.lower && i.lower <= j.upper
}

func parseInput(data []byte) [][2]interval {
	lines := strings.Split(
		strings.TrimSuffix(string(data), "\n"), "\n",
	)
	output := make([][2]interval, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, ",")
		intervals := [2]interval{}
		for j, part := range parts {
			nums := strings.Split(part, "-")
			lo, err := strconv.Atoi(nums[0])
			if err != nil {
				log.Fatal(err)
			}
			hi, err := strconv.Atoi(nums[1])
			if err != nil {
				log.Fatal(err)
			}
			intervals[j] = interval{
				lower: lo,
				upper: hi,
			}
		}
		output[i] = intervals
	}
	return output
}

func part1(assignments [][2]interval) int {
	count := 0
	for _, a := range assignments {
		if a[0].contains(a[1]) || a[1].contains(a[0]) {
			count += 1
		}
	}
	return count
}

func part2(assignments [][2]interval) int {
	count := 0
	for _, a := range assignments {
		if a[0].overlaps(a[1]) {
			count += 1
		}
	}
	return count
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	assignments := parseInput(data)

	p1 := part1(assignments)
	fmt.Println("Part 1:", p1)

	p2 := part2(assignments)
	fmt.Println("Part 2:", p2)
}
