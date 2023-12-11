// Advent of Code 2023 - Day 11.
package day11

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/maze-mapper/advent-of-code/coordinates"
)

func parseData(data []byte) ([]coordinates.Coord, []bool, []bool) {
	var galaxies []coordinates.Coord
	lines := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
	rowsWithGalaxies := make([]bool, len(lines))
	columnsWithGalaxies := make([]bool, len(lines[0]))
	for i, line := range lines {
		for j, r := range line {
			if r == '#' {
				rowsWithGalaxies[i] = true
				columnsWithGalaxies[j] = true
				galaxies = append(galaxies, coordinates.Coord{X: j, Y: i})
			}
		}
	}
	return galaxies, rowsWithGalaxies, columnsWithGalaxies
}

func expand(galaxies []coordinates.Coord, rowsWithGalaxies, columnsWithGalaxies []bool, scaleFactor int) []coordinates.Coord {
	columnExpansion := make([]int, len(columnsWithGalaxies))
	emptyColumnCount := 0
	for i, b := range columnsWithGalaxies {
		if !b {
			emptyColumnCount += 1
		}
		columnExpansion[i] = emptyColumnCount
	}

	rowExpansion := make([]int, len(rowsWithGalaxies))
	emptyRowCount := 0
	for i, b := range rowsWithGalaxies {
		if !b {
			emptyRowCount += 1
		}
		rowExpansion[i] = emptyRowCount
	}

	expandedGalaxies := make([]coordinates.Coord, len(galaxies))
	for i, g := range galaxies {
		expandedGalaxies[i] = coordinates.Coord{
			X: g.X + (scaleFactor-1)*columnExpansion[g.X],
			Y: g.Y + (scaleFactor-1)*rowExpansion[g.Y],
		}
	}
	return expandedGalaxies
}

func sumOfManhattanDistances(galaxies []coordinates.Coord) int {
	sum := 0
	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			sum += coordinates.ManhattanDistance(galaxies[i], galaxies[j])
		}
	}
	return sum
}

func part1(galaxies []coordinates.Coord, rowsWithGalaxies, columnsWithGalaxies []bool) int {
	expandedGalaxies := expand(galaxies, rowsWithGalaxies, columnsWithGalaxies, 2)
	return sumOfManhattanDistances(expandedGalaxies)
}

func part2(galaxies []coordinates.Coord, rowsWithGalaxies, columnsWithGalaxies []bool) int {
	expandedGalaxies := expand(galaxies, rowsWithGalaxies, columnsWithGalaxies, 1000000)
	return sumOfManhattanDistances(expandedGalaxies)
}

func Run(inputFile string) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	galaxies, rowsWithGalaxies, columnsWithGalaxies := parseData(data)

	p1 := part1(galaxies, rowsWithGalaxies, columnsWithGalaxies)
	fmt.Println("Part 1:", p1)

	p2 := part2(galaxies, rowsWithGalaxies, columnsWithGalaxies)
	fmt.Println("Part 2:", p2)
}
