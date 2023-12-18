// Advent of Code 2023 - Day 18.
package day18

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/maze-mapper/advent-of-code/coordinates"
)

type digDirection int

const (
	up digDirection = iota
	right
	down
	left
)

type digInstruction struct {
	direction digDirection
	distance  int
	colour    string
}

func (di digInstruction) fromColour() digInstruction {
	var direction digDirection
	switch string(di.colour[5]) {
	case "0":
		direction = right
	case "1":
		direction = down
	case "2":
		direction = left
	case "3":
		direction = up
	default:
		log.Fatalf("Invalid direction %s", string(di.colour[5]))
	}

	distance, err := strconv.ParseInt(di.colour[:5], 16, 64)
	if err != nil {
		log.Fatal(err)
	}

	return digInstruction{
		direction: direction,
		distance:  int(distance),
	}
}

func parseData(data []byte) []digInstruction {
	lines := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
	instructions := make([]digInstruction, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " ")
		instruction := digInstruction{}

		switch parts[0] {
		case "U":
			instruction.direction = up
		case "R":
			instruction.direction = right
		case "D":
			instruction.direction = down
		case "L":
			instruction.direction = left
		default:
			log.Fatalf("Unknown direction %s", parts[0])
		}

		n, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}
		instruction.distance = n

		instruction.colour = strings.TrimSuffix(strings.TrimPrefix(parts[2], "(#"), ")")

		instructions[i] = instruction
	}
	return instructions
}

func part1(instructions []digInstruction) int {
	c := coordinates.Coord{}
	trench := []coordinates.Coord{c}
	for _, instruction := range instructions {
		switch instruction.direction {
		case up:
			for d := 0; d < instruction.distance; d++ {
				c.Y -= 1
				trench = append(trench, c)
			}

		case right:
			for d := 0; d < instruction.distance; d++ {
				c.X += 1
				trench = append(trench, c)
			}
		case down:
			for d := 0; d < instruction.distance; d++ {
				c.Y += 1
				trench = append(trench, c)
			}
		case left:
			for d := 0; d < instruction.distance; d++ {
				c.X -= 1
				trench = append(trench, c)
			}
		}
	}
	topLeft, bottomRight := coordinates.Range(trench)

	trenchPoints := map[coordinates.Coord]bool{}
	for _, p := range trench {
		trenchPoints[p] = true
	}

	emptyPoints := map[coordinates.Coord]bool{}
	for x := topLeft.X; x <= bottomRight.X; x++ {
		pTop := coordinates.Coord{X: x, Y: topLeft.Y}
		if !trenchPoints[pTop] {
			emptyPoints[pTop] = true
		}
		pBottom := coordinates.Coord{X: x, Y: bottomRight.Y}
		if !trenchPoints[pBottom] {
			emptyPoints[pBottom] = true
		}
	}
	for y := topLeft.Y + 1; y <= bottomRight.Y-1; y++ {
		pLeft := coordinates.Coord{X: topLeft.X, Y: y}
		if !trenchPoints[pLeft] {
			emptyPoints[pLeft] = true
		}
		pRight := coordinates.Coord{X: bottomRight.X, Y: y}
		if !trenchPoints[pRight] {
			emptyPoints[pRight] = true
		}
	}

	var frontier []coordinates.Coord
	for k := range emptyPoints {
		frontier = append(frontier, k)
	}
	found := map[coordinates.Coord]bool{}
	floodFill(frontier, found, trenchPoints, topLeft.X, bottomRight.X, topLeft.Y, bottomRight.Y)

	area := (1 + bottomRight.X - topLeft.X) * (1 + bottomRight.Y - topLeft.Y)

	return area - len(found)

}

func floodFill(frontier []coordinates.Coord, found, trenchPoints map[coordinates.Coord]bool, xMin, xMax, yMin, yMax int) {
	for _, p := range frontier {
		if found[p] || trenchPoints[p] {
			continue
		}
		if p.X < xMin || p.Y < yMin || p.X > xMax || p.Y > yMax {
			continue
		}
		found[p] = true
		neighbours := []coordinates.Coord{
			{X: p.X, Y: p.Y - 1},
			{X: p.X + 1, Y: p.Y},
			{X: p.X, Y: p.Y + 1},
			{X: p.X - 1, Y: p.Y},
		}
		floodFill(neighbours, found, trenchPoints, xMin, xMax, yMin, yMax)
	}
}

func part2(instructions []digInstruction) int {
	var path []coordinates.Coord
	c := coordinates.Coord{}
	for _, instruction := range instructions {
		instruction = instruction.fromColour()
		switch instruction.direction {
		case up:
			c.Y -= instruction.distance
		case right:
			c.X += instruction.distance
		case down:
			c.Y += instruction.distance
		case left:
			c.X -= instruction.distance
		}
		path = append(path, c)
	}

	return shoelaceWithPerimeter(path)
}

func shoelaceWithPerimeter(points []coordinates.Coord) int {
	area := 0
	perimeter := 0
	length := len(points)
	for i := 0; i < length; i++ {
		j := (i + 1) % length
		p1 := points[i]
		p2 := points[j]
		area += (p1.X * p2.Y) - (p2.X * p1.Y)

		x := p2.X - p1.X
		if x < 0 {
			x = -x
		}
		y := p2.Y - p1.Y
		if y < 0 {
			y = -y
		}
		perimeter += x + y
	}
	area /= 2
	area += 1 + perimeter/2
	return area
}

func Run(inputFile string) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	instructions := parseData(data)

	p1 := part1(instructions)
	fmt.Println("Part 1:", p1)

	p2 := part2(instructions)
	fmt.Println("Part 2:", p2)
}
