// Advent of Code 2015 - Day 100
package day10

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

// lookAndSay determines the next term in the look-and-say sequence
func lookAndSay(digits []int) []int {
	var output []int
	var currentDigit, count int

	for _, d := range digits {
		if d == currentDigit {
			count++
		} else {
			if count != 0 {
				output = append(output, count, currentDigit)
			}
			currentDigit = d
			count = 1
		}
	}
	if count != 0 {
		output = append(output, count, currentDigit)
	}
	return output
}

// doLookAndSay repeats the lookAndSay function a certain number of times
func doLookAndSay(digits []int, iterations int) []int {
	for i := 0; i < iterations; i++ {
		digits = lookAndSay(digits)
	}
	return digits
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	numbers := []int{}
	for _, d := range data {
		if i, err := strconv.Atoi(string(d)); err == nil {
			numbers = append(numbers, i)
		}
	}

	part1 := doLookAndSay(numbers, 40)
	fmt.Println("Part 1:", len(part1))

	part2 := doLookAndSay(part1, 10) // Total of 50 iterations
	fmt.Println("Part 2:", len(part2))
}
