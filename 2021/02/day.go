// Advent of Code 2021 - Day 2
package day2

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// command holds the information for a submarine command
type command struct {
	name  string
	value int
}

// parseCommands converts the input data in to a slice of command objects
func parseCommands(data []byte) []command {
	lines := strings.Split(
		strings.TrimSuffix(string(data), "\n"), "\n",
	)
	commands := make([]command, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " ")
		if value, err := strconv.Atoi(parts[1]); err == nil {
			commands[i] = command{name: parts[0], value: value}
		} else {
			log.Fatal(err)
		}
	}
	return commands
}

func part1(commands []command) int {
	var hPos, depth int
	for _, c := range commands {
		switch c.name {
		case "forward":
			hPos += c.value
		case "down":
			depth += c.value
		case "up":
			depth -= c.value
		}
	}
	return hPos * depth
}

func part2(commands []command) int {
	var hPos, depth, aim int
	for _, c := range commands {
		switch c.name {
		case "forward":
			hPos += c.value
			depth += aim * c.value
		case "down":
			aim += c.value
		case "up":
			aim -= c.value
		}
	}
	return hPos * depth
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	commands := parseCommands(data)

	p1 := part1(commands)
	fmt.Println("Part 1:", p1)

	p2 := part2(commands)
	fmt.Println("Part 2:", p2)
}
