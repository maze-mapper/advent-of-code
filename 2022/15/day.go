// Advent of Code 2022 - Day 15.
package day15

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/maze-mapper/advent-of-code/coordinates"
)

func parseInput(data []byte) [][2]coordinates.Coord {
	lines := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
	locations := make([][2]coordinates.Coord, len(lines))
	for i, line := range lines {
		s := coordinates.Coord{}
		b := coordinates.Coord{}
		n, err := fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &(s.X), &(s.Y), &(b.X), &(b.Y))
		if err != nil {
			log.Fatal(err)
		}
		if n != 4 {
			log.Fatal("Expecting 4 items.")
		}
		locations[i] = [2]coordinates.Coord{s, b}
	}
	return locations
}

func part1(locations [][2]coordinates.Coord, row int) int {
	coveredArea := map[coordinates.Coord]struct{}{}
	for _, pair := range locations {
		s := pair[0]
		b := pair[1]
		md := coordinates.ManhattanDistance(s, b)

		t := coordinates.Coord{X: s.X, Y: row}
		md2 := coordinates.ManhattanDistance(s, t)

		d := md - md2
		if d >= 0 {
			for i := 0; i <= d; i++ {
				a := coordinates.Coord{X: s.X + i, Y: row}
				b := coordinates.Coord{X: s.X - i, Y: row}
				coveredArea[a] = struct{}{}
				coveredArea[b] = struct{}{}
			}
		}
		delete(coveredArea, b)
	}
	return len(coveredArea)
}

func part2(locations [][2]coordinates.Coord, size int) int {
	for y := 0; y < size; y++ {
		x := 0
		for x < size {
			p := coordinates.Coord{X: x, Y: y}
			covered := false
			for _, pair := range locations {
				s := pair[0]
				b := pair[1]
				coverDist := coordinates.ManhattanDistance(s, b)
				md := coordinates.ManhattanDistance(s, p)
				if md <= coverDist {
					covered = true

					// Skip x coordinate ahead to end of covered area.
					yDist := p.Y - s.Y
					if yDist < 0 {
						yDist = -yDist
					}
					xDist := coverDist - yDist
					x = s.X + xDist + 1

					break
				}
			}
			if !covered {
				return x*4000000 + y
			}
		}
	}
	return 0
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	locations := parseInput(data)

	p1 := part1(locations, 2000000)
	fmt.Println("Part 1:", p1)

	p2 := part2(locations, 4000000)
	fmt.Println("Part 2:", p2)
}
