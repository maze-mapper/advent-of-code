// Advent of Code 2024 - Day 14.
package day14

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/maze-mapper/advent-of-code/coordinates"
)

func parseData(data []byte) ([]robot, error) {
	lines := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
	robots := make([]robot, len(lines))
	for i, line := range lines {
		line = strings.ReplaceAll(line, ",", " ")
		line = strings.ReplaceAll(line, "=", " ")
		var p, v coordinates.Coord
		n, err := fmt.Sscanf(line, "p %d %d v %d %d", &p.X, &p.Y, &v.X, &v.Y)
		if err != nil {
			return nil, err
		}
		if n != 4 {
			return nil, fmt.Errorf("expected to parse 4 items, instead parsed %d", n)
		}
		robots[i] = robot{position: p, velocity: v}
	}
	return robots, nil
}

type robot struct {
	position, velocity coordinates.Coord
}

func part1(robots []robot, xSize, ySize int) int {
	positions := simulatePositions(robots, xSize, ySize, 100)
	return safetyFactor(positions, xSize, ySize)
}

func simulatePositions(robots []robot, xSize, ySize, seconds int) []coordinates.Coord {
	positions := make([]coordinates.Coord, len(robots))
	for i, r := range robots {
		nonWrappedPosition := coordinates.Coord{
			X: seconds*r.velocity.X + r.position.X,
			Y: seconds*r.velocity.Y + r.position.Y,
		}
		positions[i] = wrap(nonWrappedPosition, xSize, ySize)
	}
	return positions
}

func part2(robots []robot, xSize, ySize int) int {
	step := 1
	for {
		positions := map[coordinates.Coord]int{}
		newRobots := make([]robot, len(robots))
		for i, r := range robots {
			r.position.Transform(r.velocity)
			r.position = wrap(r.position, xSize, ySize)
			positions[r.position] += 1
			newRobots[i] = r
		}
		if hasSquare(positions) {
			printPositions(positions, xSize, ySize)
			break
		}
		robots = newRobots
		step++
	}
	return step
}

// hasSquare returns true if any position is fully surrounded by other robots.
// This is a good indicator of whether the robots are arranged into an image.
func hasSquare(positions map[coordinates.Coord]int) bool {
	for k := range positions {
		surrounded := true
	loop:
		for x := k.X - 1; x <= k.X+1; x++ {
			for y := k.Y - 1; y <= k.Y+1; y++ {
				if _, ok := positions[coordinates.Coord{X: x, Y: y}]; !ok {
					surrounded = false
				}
				if !surrounded {
					break loop
				}
			}
		}
		if surrounded {
			return true
		}
	}
	return false
}

func printPositions(positions map[coordinates.Coord]int, xSize, ySize int) {
	var b strings.Builder
	for y := 0; y < ySize; y++ {
		for x := 0; x < xSize; x++ {
			if _, ok := positions[coordinates.Coord{X: x, Y: y}]; ok {
				b.WriteString("#")
			} else {
				b.WriteString(".")
			}
		}
		b.WriteString("\n")
	}
	fmt.Println(b.String())
}

func wrap(position coordinates.Coord, xSize, ySize int) coordinates.Coord {
	wrappedPosition := coordinates.Coord{
		X: position.X % xSize,
		Y: position.Y % ySize,
	}
	if wrappedPosition.X < 0 {
		wrappedPosition.X += xSize
	}
	if wrappedPosition.Y < 0 {
		wrappedPosition.Y += ySize
	}
	return wrappedPosition
}

func safetyFactor(positions []coordinates.Coord, xSize, ySize int) int {
	// Determine boundaries of quadrants.
	x1 := xSize / 2
	x2 := (xSize + 1) / 2
	y1 := ySize / 2
	y2 := (ySize + 1) / 2

	var q1, q2, q3, q4 int
	for _, p := range positions {
		switch {
		case p.X < x1 && p.Y < y1:
			q1++
		case p.X >= x2 && p.Y < y1:
			q2++
		case p.X < x1 && p.Y >= y2:
			q3++
		case p.X >= x2 && p.Y >= y2:
			q4++
		}
	}
	return q1 * q2 * q3 * q4
}

func Run(inputFile string) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	robots, err := parseData(data)
	if err != nil {
		log.Fatal(err)
	}

	xSize := 101
	ySize := 103

	p1 := part1(robots, xSize, ySize)
	fmt.Println("Part 1:", p1)

	p2 := part2(robots, xSize, ySize)
	fmt.Println("Part 2:", p2)
}
