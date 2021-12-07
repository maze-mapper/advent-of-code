// Advent of Code 2021 - Day 7
package day7

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

// parseData returns the input data as a slice of ints
func parseData(data []byte) []int {
	parts := strings.Split(
		strings.TrimSuffix(string(data), "\n"), ",",
	)
	numbers := make([]int, len(parts))
	for i, s := range parts {
		if val, err := strconv.Atoi(s); err == nil {
			numbers[i] = val
		} else {
			log.Fatal(err)
		}
	}
	return numbers
}

// difference returns the difference between two numbers
func difference(a, b int) int {
	switch {
	case a > b:
		return a - b
	case a < b:
		return b - a
	default:
		return 0
	}
}

// nop performs no operation and returns the provided int
func nop(n int) int {
	return n
}

// sumArithmeticSeries returns the sum of the first n terms of the artihmetic series for the crab fuel use
func sumArithmeticSeries(n int) int {
	return n * (1 + n) / 2
}

// calculateSpentFuel returns the fuel required to move all crabs to the given position
func calculateSpentFuel(numbers []int, pos int, f func(d int) int) int {
	fuel := 0
	for _, n := range numbers {
		diff := difference(n, pos)
		fuel += f(diff)
	}
	return fuel
}

// determineMinFuel finds the minimum amount of fuel for all crabs to reach the same position
// The function f determines how much fuel is spent to move a certain distance
func determineMinFuel(numbers []int, f func(d int) int) int {
	minFuel := int(^uint(0) >> 1) // Max value
	for i := numbers[0]; i < numbers[len(numbers)-1]; i++ {
		fuel := calculateSpentFuel(numbers, i, f)
		if fuel < minFuel {
			minFuel = fuel
		}
	}
	return minFuel
}

func part1(numbers []int) int {
	return determineMinFuel(numbers, nop)
}

func part2(numbers []int) int {
	return determineMinFuel(numbers, sumArithmeticSeries)
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	positions := parseData(data)
	// Sort positions so we can easily range over them
	sort.Ints(positions)

	p1 := part1(positions)
	fmt.Println("Part 1:", p1)

	p2 := part2(positions)
	fmt.Println("Part 2:", p2)
}
