// Advent of Code 2021 - Day 1
package day1

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// parseDepths converts the input data in to an integer slice
func parseDepths(data []byte) []int {
	lines := strings.Split(
		strings.TrimSuffix(string(data), "\n"), "\n",
	)
	depths := make([]int, len(lines))
	for i, line := range lines {
		if depth, err := strconv.Atoi(line); err == nil {
			depths[i] = depth
		} else {
			log.Fatal(err)
		}
	}
	return depths
}

func part1(depths []int) int {
	currentDepth := depths[0]
	increases := 0
	for i := 1; i < len(depths); i++ {
		if depths[i] > currentDepth {
			increases += 1
		}
		currentDepth = depths[i]
	}
	return increases
}

// sum adds up the integers in the given slice
func sum(values []int) int {
	s := 0
	for _, v := range values {
		s += v
	}
	return s
}

func part2(depths []int) int {
	windowSize := 3
	// Determine sum of measurements in first window
	currentSum := sum(depths[0:windowSize])

	increases := 0
	for i := windowSize + 1; i <= len(depths); i++ {
		newSum := sum(depths[i-windowSize : i])
		if newSum > currentSum {
			increases += 1
		}
		currentSum = newSum
	}
	return increases
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	depths := parseDepths(data)

	p1 := part1(depths)
	fmt.Println("Part 1:", p1)

	p2 := part2(depths)
	fmt.Println("Part 2:", p2)
}
