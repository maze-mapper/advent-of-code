// Advent of Code 2024 - Day 8.
package day8

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/maze-mapper/advent-of-code/coordinates"
)

func parseData(data []byte) (map[string][]coordinates.Coord, coordinates.Coord, coordinates.Coord) {
	lines := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
	m := map[string][]coordinates.Coord{}
	for i, line := range lines {
		for j, r := range line {
			if r == '.' {
				continue
			}
			s := string(r)
			m[s] = append(m[s], coordinates.Coord{X: j, Y: i})
		}
	}
	minC := coordinates.Coord{}
	maxC := coordinates.Coord{X: len(lines[0]) - 1, Y: len(lines) - 1}
	return m, minC, maxC
}

func inBounds(c, minC, maxC coordinates.Coord) bool {
	return c.X >= minC.X && c.X <= maxC.X && c.Y >= minC.Y && c.Y <= maxC.Y
}

func part1(antennas map[string][]coordinates.Coord, minC, maxC coordinates.Coord) int {
	return solve(antennas, minC, maxC, findAntinodes)
}

func part2(antennas map[string][]coordinates.Coord, minC, maxC coordinates.Coord) int {
	return solve(antennas, minC, maxC, findResonantAntinodes)
}

type antinodeFunc func([]coordinates.Coord, coordinates.Coord, coordinates.Coord) []coordinates.Coord

func solve(antennas map[string][]coordinates.Coord, minC, maxC coordinates.Coord, f antinodeFunc) int {
	m := map[coordinates.Coord]bool{}
	for _, v := range antennas {
		positions := f(v, minC, maxC)
		for _, p := range positions {
			m[p] = true
		}
	}
	return len(m)
}

func findAntinodes(locations []coordinates.Coord, minC, maxC coordinates.Coord) []coordinates.Coord {
	var antinodes []coordinates.Coord
	for i := 0; i < len(locations)-1; i++ {
		for j := i + 1; j < len(locations); j++ {
			a := locations[i]
			b := locations[j]
			aToB := coordinates.Coord{
				X: b.X - a.X,
				Y: b.Y - a.Y,
			}
			p1 := coordinates.Coord{
				X: a.X - aToB.X,
				Y: a.Y - aToB.Y,
			}
			p2 := coordinates.Coord{
				X: b.X + aToB.X,
				Y: b.Y + aToB.Y,
			}
			for _, p := range []coordinates.Coord{p1, p2} {
				if inBounds(p, minC, maxC) {
					antinodes = append(antinodes, p)
				}
			}
		}
	}
	return antinodes
}

func findResonantAntinodes(locations []coordinates.Coord, minC, maxC coordinates.Coord) []coordinates.Coord {
	var antinodes []coordinates.Coord
	for i := 0; i < len(locations)-1; i++ {
		for j := i + 1; j < len(locations); j++ {
			a := locations[i]
			b := locations[j]
			aToB := coordinates.Coord{
				X: b.X - a.X,
				Y: b.Y - a.Y,
			}
			p1 := coordinates.Coord{X: a.X, Y: a.Y}
			for {
				if !inBounds(p1, minC, maxC) {
					break
				}
				antinodes = append(antinodes, coordinates.Coord{X: p1.X, Y: p1.Y})
				p1.X -= aToB.X
				p1.Y -= aToB.Y
			}
			p2 := coordinates.Coord{X: b.X, Y: b.Y}
			for {
				if !inBounds(p2, minC, maxC) {
					break
				}
				antinodes = append(antinodes, coordinates.Coord{X: p2.X, Y: p2.Y})
				p2.X += aToB.X
				p2.Y += aToB.Y
			}
		}
	}
	return antinodes
}

func Run(inputFile string) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	antennas, minC, maxC := parseData(data)

	p1 := part1(antennas, minC, maxC)
	fmt.Println("Part 1:", p1)

	p2 := part2(antennas, minC, maxC)
	fmt.Println("Part 2:", p2)
}
