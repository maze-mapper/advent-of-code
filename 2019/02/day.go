// Advent of Code 2019 - Day 2
package day2

import (
	"fmt"
	"io/ioutil"
	"log"

	"adventofcode/2019/intcode"
)

func part1(program []int) int {
	return runGravityAssist(program, 12, 2)
}

// runGravityAssist will run the given program with the input values substituted and return the first address element
func runGravityAssist(program []int, noun, verb int) int {
	program[1] = noun
	program[2] = verb

	computer := intcode.New(program)
	computer.Run()
	return computer.Program()[0]
}

func part2(program []int) int {
	maxVal := 99
	goal := 19690720
	for i := 0; i <= maxVal; i++ {
		for j := 0; j <= maxVal; j++ {
			// TODO: can goroutines be used here but still not execute all possible permutations?
			// If so we will need to create local copies of noun and verb within each iteration like below
			noun, verb := i, j
			if output := runGravityAssist(program, noun, verb); output == goal {
				return 100*noun + verb
			}
		}
	}
	return 0
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	program := intcode.ReadProgram(data)

	p1 := part1(program)
	fmt.Println("Part 1:", p1)

	p2 := part2(program)
	fmt.Println("Part 2:", p2)
}
