// Advent of Code 2022 - Day 10.
package day10

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type instruction struct {
	cycles int
	value  int
}

func parseInput(data []byte) []instruction {
	lines := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
	instructions := make([]instruction, len(lines))
	for i, line := range lines {
		if line == "noop" {
			instructions[i] = instruction{
				cycles: 1,
			}
		} else {
			parts := strings.Split(line, " ")
			val, err := strconv.Atoi(parts[1])
			if err != nil {
				log.Fatal(err)
			}
			instructions[i] = instruction{
				cycles: 2,
				value:  val,
			}
		}
	}
	return instructions
}

func part1(instructions []instruction) int {
	maxCycles := 220
	signalStrength := 0
	x := 1
	cycle := 0
	for _, ins := range instructions {
		for c := 0; c < ins.cycles; c++ {
			cycle += 1
			if cycle%40 == 20 && cycle <= maxCycles {
				signalStrength += cycle * x
			}
		}
		x += ins.value
	}
	return signalStrength
}

func part2(instructions []instruction) {
	rows := 6
	columns := 40
	maxCycles := rows * columns

	x := 1
	cycle := 0
	for _, ins := range instructions {
		if cycle >= maxCycles {
			break
		}
		for c := 0; c < ins.cycles; c++ {
			currentCol := cycle % columns
			cycle += 1
			if (currentCol <= x+1) && (currentCol >= x-1) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
			if currentCol == columns-1 {
				fmt.Print("\n")
			}
		}
		x += ins.value
	}
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	instructions := parseInput(data)

	p1 := part1(instructions)
	fmt.Println("Part 1:", p1)

	fmt.Println("Part 2:")
	part2(instructions)
}
