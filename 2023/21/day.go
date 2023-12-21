// Advent of Code 2023 - Day 21.
package day21

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/maze-mapper/advent-of-code/coordinates"
)

func parseData(data []byte) (coordinates.Coord, map[coordinates.Coord]bool, int, int) {
	lines := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
	rocks := map[coordinates.Coord]bool{}
	var start coordinates.Coord
	for i, line := range lines {
		for j, r := range line {
			if r == '#' {
				rocks[coordinates.Coord{X: j, Y: i}] = true
			}
			if r == 'S' {
				start.Y = i
				start.X = j
			}
		}
	}
	yMax := len(lines) - 1
	xMax := len(lines[yMax]) - 1
	return start, rocks, yMax, xMax
}

func part1(start coordinates.Coord, rocks map[coordinates.Coord]bool, yMax int, xMax int, maxSteps int) int {
	frontier := map[coordinates.Coord]bool{start: true}
	for step := 0; step < maxSteps; step++ {
		newFrontier := map[coordinates.Coord]bool{}
		for c := range frontier {
			neighbours := []coordinates.Coord{
				{X: c.X, Y: c.Y - 1},
				{X: c.X + 1, Y: c.Y},
				{X: c.X, Y: c.Y + 1},
				{X: c.X - 1, Y: c.Y},
			}
			for _, n := range neighbours {
				if n.Y < 0 || n.X < 0 || n.Y > yMax || n.X > xMax {
					continue
				}
				if rocks[n] {
					continue
				}
				newFrontier[n] = true
			}
		}
		frontier = newFrontier
	}
	return len(frontier)
}

func part2(start coordinates.Coord, rocks map[coordinates.Coord]bool, yMax int, xMax int, maxSteps int) int {
	if yMax != xMax {
		log.Fatal("Expecting a square grid")
	}
	for c := range rocks {
		if c.X == start.X || c.Y == start.Y {
			log.Fatal("Algorithm assumes the start point is in an empty row and column")
		}
	}
	gridSize := xMax + 1
	offset := maxSteps % gridSize
	var points []int
	for i := 0; i < 3; i++ {
		points = append(points, i*gridSize+offset)
	}
	got := make([]int, len(points))

	frontier := map[coordinates.Coord]bool{start: true}
	for step := 0; step < maxSteps; step++ {
		for i, p := range points {
			if step == p {
				got[i] = len(frontier)
			}
		}
		if got[len(got)-1] != 0 {
			break
		}
		newFrontier := map[coordinates.Coord]bool{}
		for c := range frontier {
			neighbours := []coordinates.Coord{
				{X: c.X, Y: c.Y - 1},
				{X: c.X + 1, Y: c.Y},
				{X: c.X, Y: c.Y + 1},
				{X: c.X - 1, Y: c.Y},
			}
			for _, n := range neighbours {
				scaledX := n.X
				scaledY := n.Y
				for scaledX < 0 {
					scaledX += gridSize
				}
				for scaledY < 0 {
					scaledY += gridSize
				}
				scaledX = scaledX % gridSize
				scaledY = scaledY % gridSize
				scaledN := coordinates.Coord{
					X: scaledX,
					Y: scaledY,
				}
				if rocks[scaledN] {
					continue
				}
				newFrontier[n] = true
			}
		}
		frontier = newFrontier
	}

	// Calculate polynomial coefficients.
	c := got[0]
	a := (got[2] + c - 2*got[1]) / 2
	b := got[1] - c - a
	fmt.Printf("a=%d b=%d c=%d\n", a, b, c)

	n := (maxSteps - offset) / gridSize
	ans := (a * n * n) + (b * n) + c

	return ans
}

func Run(inputFile string) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	start, rocks, yMax, xMax := parseData(data)

	p1 := part1(start, rocks, yMax, xMax, 64)
	fmt.Println("Part 1:", p1)

	p2 := part2(start, rocks, yMax, xMax, 26501365)
	fmt.Println("Part 2:", p2)
}
