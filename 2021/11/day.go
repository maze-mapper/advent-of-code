// Advent of Code 2021 - Day 11
package day11

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// parseData returns the data as a 10x10 grid of ints
func parseData(data []byte) [10][10]int {
	lines := strings.Split(
		strings.TrimSuffix(string(data), "\n"), "\n",
	)
	octopuses := [10][10]int{}
	for i, line := range lines {
		octopuses[i] = [10]int{}
		for j, r := range line {
			if val, err := strconv.Atoi(string(r)); err == nil {
				octopuses[i][j] = val
			} else {
				log.Fatal(err)
			}
		}
	}
	return octopuses
}

var maxEnergyLevel int = 9

func doStep(octopuses *[10][10]int) int {
	// Increment energy level
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			octopuses[i][j] += 1
		}
	}

	// Count flashes
	totalFlashes := 0
	noFlashes := false
	for noFlashes == false {
		flashes := 0
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				if octopuses[i][j] > maxEnergyLevel {
					flashes += 1
					for ii := i - 1; ii < i+2; ii++ {
						if ii < 0 || ii >= 10 {
							continue
						}
						for jj := j - 1; jj < j+2; jj++ {
							if jj < 0 || jj >= 10 {
								continue
							}
							if octopuses[ii][jj] != 0 {
								octopuses[ii][jj] += 1
							}
						}
					}
					octopuses[i][j] = 0
				}
			}
		}
		if flashes == 0 {
			noFlashes = true
		} else {
			totalFlashes += flashes
		}
	}

	return totalFlashes
}

func part1(octopuses *[10][10]int, steps int) int {
	flashes := 0
	for i := 0; i < steps; i++ {
		f := doStep(octopuses)
		flashes += f
	}
	return flashes
}

func part2(octopuses *[10][10]int, step int) int {
	flashes := 0
	for flashes != 100 {
		step += 1
		flashes = doStep(octopuses)
	}
	return step
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	octopuses := parseData(data)
	steps := 100

	p1 := part1(&octopuses, steps)
	fmt.Println("Part 1:", p1)

	p2 := part2(&octopuses, steps)
	fmt.Println("Part 2:", p2)
}
