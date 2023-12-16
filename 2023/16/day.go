// Advent of Code 2023 - Day 16.
package day16

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/maze-mapper/advent-of-code/coordinates"
)

func parseData(data []byte) [][]string {
	lines := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
	area := make([][]string, len(lines))
	for i, line := range lines {
		area[i] = strings.Split(line, "")
	}
	return area
}

type beamDirection int

const (
	up beamDirection = iota
	right
	down
	left
)

type beamPoint struct {
	c coordinates.Coord
	d beamDirection
}

func (bp beamPoint) nextPoint() coordinates.Coord {
	var c coordinates.Coord
	switch bp.d {
	case up:
		c = coordinates.Coord{X: bp.c.X, Y: bp.c.Y - 1}
	case right:
		c = coordinates.Coord{X: bp.c.X + 1, Y: bp.c.Y}
	case down:
		c = coordinates.Coord{X: bp.c.X, Y: bp.c.Y + 1}
	case left:
		c = coordinates.Coord{X: bp.c.X - 1, Y: bp.c.Y}
	}
	return c
}

func energised(input [][]string, beam beamPoint) int {
	beams := []beamPoint{beam}
	visited := map[coordinates.Coord]map[beamDirection]bool{}
	for len(beams) > 0 {
		var newBeams []beamPoint
		for _, beam := range beams {

			// Check if a beam at this point in the same direction has been seen before.
			if _, ok := visited[beam.c]; ok {
				if ok := visited[beam.c][beam.d]; ok {
					continue
				}
			} else {
				visited[beam.c] = map[beamDirection]bool{}
			}
			visited[beam.c][beam.d] = true

			var nextBeams []beamPoint
			switch input[beam.c.Y][beam.c.X] {
			case "/":
				switch beam.d {
				case up:
					nextBeams = append(nextBeams, beamPoint{c: coordinates.Coord{X: beam.c.X + 1, Y: beam.c.Y}, d: right})
				case right:
					nextBeams = append(nextBeams, beamPoint{c: coordinates.Coord{X: beam.c.X, Y: beam.c.Y - 1}, d: up})
				case down:
					nextBeams = append(nextBeams, beamPoint{c: coordinates.Coord{X: beam.c.X - 1, Y: beam.c.Y}, d: left})
				case left:
					nextBeams = append(nextBeams, beamPoint{c: coordinates.Coord{X: beam.c.X, Y: beam.c.Y + 1}, d: down})
				}
			case "\\":
				switch beam.d {
				case up:
					nextBeams = append(nextBeams, beamPoint{c: coordinates.Coord{X: beam.c.X - 1, Y: beam.c.Y}, d: left})
				case right:
					nextBeams = append(nextBeams, beamPoint{c: coordinates.Coord{X: beam.c.X, Y: beam.c.Y + 1}, d: down})
				case down:
					nextBeams = append(nextBeams, beamPoint{c: coordinates.Coord{X: beam.c.X + 1, Y: beam.c.Y}, d: right})
				case left:
					nextBeams = append(nextBeams, beamPoint{c: coordinates.Coord{X: beam.c.X, Y: beam.c.Y - 1}, d: up})
				}
			case "|":
				if beam.d == right || beam.d == left {
					nextBeams = append(nextBeams, beamPoint{c: coordinates.Coord{X: beam.c.X, Y: beam.c.Y - 1}, d: up})
					nextBeams = append(nextBeams, beamPoint{c: coordinates.Coord{X: beam.c.X, Y: beam.c.Y + 1}, d: down})
				} else {
					nextBeams = append(nextBeams, beamPoint{c: beam.nextPoint(), d: beam.d})
				}
			case "-":
				if beam.d == up || beam.d == down {
					nextBeams = append(nextBeams, beamPoint{c: coordinates.Coord{X: beam.c.X - 1, Y: beam.c.Y}, d: left})
					nextBeams = append(nextBeams, beamPoint{c: coordinates.Coord{X: beam.c.X + 1, Y: beam.c.Y}, d: right})
				} else {
					nextBeams = append(nextBeams, beamPoint{c: beam.nextPoint(), d: beam.d})
				}
			default:
				nextBeams = append(nextBeams, beamPoint{c: beam.nextPoint(), d: beam.d})
			}

			// Ensure the next beam positions are within bounds.
			for _, b := range nextBeams {
				if b.c.Y >= 0 && b.c.X >= 0 && b.c.Y < len(input) && b.c.X < len(input[b.c.Y]) {
					newBeams = append(newBeams, b)
				}
			}
		}
		beams = newBeams
	}
	return len(visited)
}

func part1(input [][]string) int {
	return energised(input, beamPoint{c: coordinates.Coord{}, d: right})
}

func part2(input [][]string) int {
	maxEnergised := 0
	for row := 0; row < len(input); row++ {
		if e1 := energised(input, beamPoint{c: coordinates.Coord{X: 0, Y: row}, d: right}); e1 > maxEnergised {
			maxEnergised = e1
		}
		if e2 := energised(input, beamPoint{c: coordinates.Coord{X: len(input[row]) - 1, Y: row}, d: left}); e2 > maxEnergised {
			maxEnergised = e2
		}
	}
	for col := 0; col < len(input[0]); col++ {
		if e1 := energised(input, beamPoint{c: coordinates.Coord{X: col, Y: 0}, d: down}); e1 > maxEnergised {
			maxEnergised = e1
		}
		if e2 := energised(input, beamPoint{c: coordinates.Coord{X: col, Y: len(input) - 1}, d: up}); e2 > maxEnergised {
			maxEnergised = e2
		}
	}
	return maxEnergised
}

func Run(inputFile string) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	input := parseData(data)

	p1 := part1(input)
	fmt.Println("Part 1:", p1)

	p2 := part2(input)
	fmt.Println("Part 2:", p2)
}
