// Advent of Code 2021 - Day 5
package day5

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/maze-mapper/advent-of-code/coordinates"
)

// Line holds the two coordinates that define the line
type Line struct {
	A, B coordinates.Coord
}

// parsePoint convers a string "x,y" in to a coordinate
func parsePoint(s string) coordinates.Coord {
	parts := strings.Split(s, ",")

	x, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Fatal(err)
	}

	y, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatal(err)
	}

	return coordinates.Coord{X: x, Y: y}

}

// parseData converts the input data in to a slice of Line objects
func parseData(data []byte) []Line {
	inputLines := strings.Split(
		strings.TrimSuffix(string(data), "\n"), "\n",
	)
	lines := make([]Line, len(inputLines))
	for i, l := range inputLines {
		points := strings.Split(l, " -> ")
		a := parsePoint(points[0])
		b := parsePoint(points[1])
		lines[i] = Line{A: a, B: b}
	}
	return lines
}

// removeDiagonal filters out any diagonal lines
func removeDiagonal(lines []Line) []Line {
	filteredLines := []Line{}
	for _, line := range lines {
		if line.A.X == line.B.X || line.A.Y == line.B.Y {
			filteredLines = append(filteredLines, line)
		}
	}
	return filteredLines
}

// generatePoints returns the list of points making up a line
func generatePoints(line Line) []coordinates.Coord {
	points := []coordinates.Coord{line.A}
	c := coordinates.Coord{X: line.A.X, Y: line.A.Y}
	for c != line.B {
		// Update x coordinate
		switch {
		case line.B.X > c.X:
			c.X += 1
		case line.B.X < c.X:
			c.X -= 1
		}

		// Update y coordinate
		switch {
		case line.B.Y > c.Y:
			c.Y += 1
		case line.B.Y < c.Y:
			c.Y -= 1
		}

		points = append(points, coordinates.Coord{X: c.X, Y: c.Y})
	}
	return points
}

// placeLines marks the positions of lines and the number of lines passing through each point
func placeLines(lines []Line) map[coordinates.Coord]int {
	placed := map[coordinates.Coord]int{}
	for _, line := range lines {
		points := generatePoints(line)
		for _, p := range points {
			placed[p] += 1
		}
	}
	return placed
}

// countOverlaps returns the number of coordinates where lines overlap
func countOverlaps(placed map[coordinates.Coord]int) int {
	count := 0
	for _, v := range placed {
		if v > 1 {
			count += 1
		}
	}
	return count
}

func part1(lines []Line) int {
	lines = removeDiagonal(lines)
	placed := placeLines(lines)
	return countOverlaps(placed)
}

func part2(lines []Line) int {
	placed := placeLines(lines)
	return countOverlaps(placed)
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	lines := parseData(data)

	p1 := part1(lines)
	fmt.Println("Part 1:", p1)

	p2 := part2(lines)
	fmt.Println("Part 2:", p2)
}
