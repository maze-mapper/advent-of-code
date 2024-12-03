// Advent of Code 2024 - Day 3.
package day3

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	re  = regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	re2 = regexp.MustCompile(`(mul\([0-9]{1,3},[0-9]{1,3}\)|do\(\)|don't\(\))`)
)

func parseData(data []byte) string {
	return strings.TrimSuffix(string(data), "\n")
}

// mul performs the multiplication of a string mul instruction, e.g. mul(123,456)
func mul(s string) (int, error) {
	s = strings.TrimSuffix(strings.TrimPrefix(s, "mul("), ")")
	parts := strings.Split(s, ",")
	if len(parts) != 2 {
		return 0, fmt.Errorf("found %d parts instead of 2", len(parts))
	}
	l, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, err
	}
	r, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, err
	}
	return l * r, nil
}

func part1(s string) int {
	var total int
	matches := re.FindAllString(s, -1)
	for _, match := range matches {
		m, err := mul(match)
		if err != nil {
			log.Fatal(err)
		}
		total += m
	}
	return total
}

func part2(s string) int {
	var total int
	mulEnabled := true
	matches := re2.FindAllString(s, -1)
	for _, match := range matches {
		switch match {
		case "do()":
			mulEnabled = true
		case "don't()":
			mulEnabled = false
		default:
			if mulEnabled {
				m, err := mul(match)
				if err != nil {
					log.Fatal(err)
				}
				total += m
			}
		}
	}
	return total
}

func Run(inputFile string) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	memory := parseData(data)

	p1 := part1(memory)
	fmt.Println("Part 1:", p1)

	p2 := part2(memory)
	fmt.Println("Part 2:", p2)
}
