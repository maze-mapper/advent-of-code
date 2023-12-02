// Advent of Code 2023 - Day 2.
package day2

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type cubes struct {
	r, g, b int
}

func (c cubes) isSubsetOf(other cubes) bool {
	return c.r <= other.r && c.g <= other.g && c.b <= other.b
}

func (c cubes) power() int {
	return c.r * c.g * c.b
}

func parseData(data []byte) [][]cubes {
	lines := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
	games := make([][]cubes, len(lines))
	for i, line := range lines {
		draws := strings.Split(line, ": ")[1]
		var results []cubes
		for _, draw := range strings.Split(draws, "; ") {
			var c cubes
			for _, s := range strings.Split(draw, ", ") {
				parts := strings.Split(s, " ")
				n, err := strconv.Atoi(parts[0])
				if err != nil {
					log.Fatal(err)
				}
				switch parts[1] {
				case "red":
					c.r = n
				case "green":
					c.g = n
				case "blue":
					c.b = n
				}
			}
			results = append(results, c)
		}
		games[i] = results
	}
	return games
}

func part1(games [][]cubes) int {
	target := cubes{r: 12, g: 13, b: 14}
	count := 0
	for i, game := range games {
		possible := true
		for _, draw := range game {
			if !draw.isSubsetOf(target) {
				possible = false
			}
		}
		if possible {
			count += i + 1
		}
	}
	return count
}

func part2(games [][]cubes) int {
	sum := 0
	for _, game := range games {
		var minCubes cubes
		for _, draw := range game {
			minCubes.r = max(minCubes.r, draw.r)
			minCubes.g = max(minCubes.g, draw.g)
			minCubes.b = max(minCubes.b, draw.b)
		}
		sum += minCubes.power()
	}
	return sum
}

func Run(inputFile string) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	games := parseData(data)

	p1 := part1(games)
	fmt.Println("Part 1:", p1)

	p2 := part2(games)
	fmt.Println("Part 2:", p2)
}
