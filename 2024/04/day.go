// Advent of Code 2024 - Day 4.
package day4

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func parseData(data []byte) [][]string {
	lines := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
	ws := make([][]string, len(lines))
	for i, line := range lines {
		ws[i] = make([]string, len(line))
		for j, r := range line {
			ws[i][j] = string(r)
		}
	}
	return ws
}

type direction int

const (
	directionUp = iota
	directionDown
	directionLeft
	directionRight
	directionUpLeft
	directionDownLeft
	directionUpRight
	directionDownRight
)

var allDirections = []direction{
	directionUp,
	directionDown,
	directionLeft,
	directionRight,
	directionUpLeft,
	directionDownLeft,
	directionUpRight,
	directionDownRight,
}

func shift(x, y int, d direction) (int, int) {
	switch d {
	case directionUp:
		y--
	case directionDown:
		y++
	case directionLeft:
		x--
	case directionRight:
		x++
	case directionUpLeft:
		y--
		x--
	case directionDownLeft:
		y++
		x--
	case directionUpRight:
		y--
		x++
	case directionDownRight:
		y++
		x++
	}
	return x, y
}

func part1(ws [][]string) int {
	var count int
	for i := range ws {
		for j := range ws[i] {
			if ws[i][j] != "X" {
				continue
			}
			for _, d := range allDirections {
				y := i
				x := j
				for n, l := range []string{"M", "A", "S"} {
					x, y = shift(x, y, d)
					if x < 0 || y < 0 || x >= len(ws[i]) || y >= len(ws) {
						break
					}
					if ws[y][x] != l {
						break
					}
					if n == 2 {
						count++
					}
				}
			}
		}
	}
	return count
}

func part2(ws [][]string) int {
	var count int
	// Omit outer cells as these can never form the centre of an X shape.
	for i := 1; i < len(ws)-1; i++ {
		for j := 1; j < len(ws[i])-1; j++ {
			if ws[i][j] != "A" {
				continue
			}
			// Check diagonal: \
			if (ws[i-1][j-1] == "M" && ws[i+1][j+1] == "S") || (ws[i-1][j-1] == "S" && ws[i+1][j+1] == "M") {
				// Check diagonal: /
				if (ws[i+1][j-1] == "M" && ws[i-1][j+1] == "S") || (ws[i+1][j-1] == "S" && ws[i-1][j+1] == "M") {
					count++
				}
			}
		}
	}
	return count
}

func Run(inputFile string) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	ws := parseData(data)

	p1 := part1(ws)
	fmt.Println("Part 1:", p1)

	p2 := part2(ws)
	fmt.Println("Part 2:", p2)
}
