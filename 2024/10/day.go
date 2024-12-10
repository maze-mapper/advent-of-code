// Advent of Code 2024 - Day 10.
package day10

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseData(data []byte) ([][]int, error) {
	lines := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
	area := make([][]int, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, "")
		area[i] = make([]int, len(parts))
		for j, p := range parts {
			n, err := strconv.Atoi(p)
			if err != nil {
				return nil, err
			}
			area[i][j] = n
		}
	}
	return area, nil
}

func part1(region [][]int) int {
	var sum int
	for i, line := range region {
		for j, p := range line {
			if p == 0 {
				score := len(findHikingTrails(region, j, i))
				sum += score
			}
		}
	}
	return sum
}

func part2(region [][]int) int {
	var sum int
	for i, line := range region {
		for j, p := range line {
			if p == 0 {
				routes := findHikingTrails(region, j, i)
				var rating int
				for _, count := range routes {
					rating += count
				}
				sum += rating
			}
		}
	}
	return sum
}

func findHikingTrails(region [][]int, x, y int) map[[2]int]int {
	trails := map[[2]int]int{}
	currentHeight := region[y][x]
	frontier := [][2]int{{x, y}}
	for len(frontier) > 0 {
		var newFrontier [][2]int
		for _, p := range frontier {
			x = p[0]
			y = p[1]
			for _, neighbour := range getNeighbours(region, x, y) {
				newX := neighbour[0]
				newY := neighbour[1]
				if region[newY][newX] == currentHeight+1 {
					newFrontier = append(newFrontier, [2]int{newX, newY})
				}
			}
		}
		frontier = newFrontier
		currentHeight++

		if currentHeight == 9 {
			for _, p := range frontier {
				trails[p]++
			}
			break
		}
	}
	return trails
}

func getNeighbours(region [][]int, x, y int) [][2]int {
	var neighbours [][2]int
	candidates := [][2]int{
		{x - 1, y},
		{x + 1, y},
		{x, y - 1},
		{x, y + 1},
	}
	for _, c := range candidates {
		newX := c[0]
		newY := c[1]
		if newY < 0 || newY >= len(region) || newX < 0 || newX >= len(region[newY]) {
			continue
		}
		neighbours = append(neighbours, c)
	}
	return neighbours
}

func Run(inputFile string) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	region, err := parseData(data)
	if err != nil {
		log.Fatal(err)
	}

	p1 := part1(region)
	fmt.Println("Part 1:", p1)

	p2 := part2(region)
	fmt.Println("Part 2:", p2)
}
