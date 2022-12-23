// Advent of Code 2022 - Day 23.
package day23

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/maze-mapper/advent-of-code/coordinates"
)

func parseInput(data []byte) map[coordinates.Coord]struct{} {
	lines := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
	positions := map[coordinates.Coord]struct{}{}
	for y, line := range lines {
		for x, c := range line {
			if c == '#' {
				p := coordinates.Coord{X: x, Y: y}
				positions[p] = struct{}{}
			}
		}
	}
	return positions
}

var directions = []string{"N", "S", "W", "E"}

func printElves(elves map[coordinates.Coord]struct{}) {
	finalElves := make([]coordinates.Coord, len(elves))
	i := 0
	for e := range elves {
		finalElves[i] = e
		i += 1
	}
	min, max := coordinates.Range(finalElves)

	for y := min.Y; y <= max.Y; y++ {
		for x := min.X; x <= max.X; x++ {
			c := coordinates.Coord{X: x, Y: y}
			if _, ok := elves[c]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}

// shouldMove returns false if all eight spaces around the elf are empty.
func shouldMove(elves map[coordinates.Coord]struct{}, e coordinates.Coord) bool {
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			c := coordinates.Coord{X: e.X + x, Y: e.Y + y}
			if c == e {
				continue
			}
			if _, ok := elves[c]; ok {
				return true
			}
		}
	}
	return false
}

// canMove returns true if the elf can move in the given direction.
func canMove(elves map[coordinates.Coord]struct{}, e coordinates.Coord, dir string) (coordinates.Coord, bool) {
	var c, c1, c2 coordinates.Coord
	switch dir {
	case "N":
		c = coordinates.Coord{X: e.X, Y: e.Y - 1}
		c1 = coordinates.Coord{X: e.X - 1, Y: e.Y - 1}
		c2 = coordinates.Coord{X: e.X + 1, Y: e.Y - 1}
	case "S":
		c = coordinates.Coord{X: e.X, Y: e.Y + 1}
		c1 = coordinates.Coord{X: e.X - 1, Y: e.Y + 1}
		c2 = coordinates.Coord{X: e.X + 1, Y: e.Y + 1}
	case "W":
		c = coordinates.Coord{X: e.X - 1, Y: e.Y}
		c1 = coordinates.Coord{X: e.X - 1, Y: e.Y - 1}
		c2 = coordinates.Coord{X: e.X - 1, Y: e.Y + 1}
	case "E":
		c = coordinates.Coord{X: e.X + 1, Y: e.Y}
		c1 = coordinates.Coord{X: e.X + 1, Y: e.Y - 1}
		c2 = coordinates.Coord{X: e.X + 1, Y: e.Y + 1}
	}
	for _, p := range []coordinates.Coord{c, c1, c2} {
		if _, ok := elves[p]; ok {
			return coordinates.Coord{}, false
		}
	}
	return c, true
}

func round(elves map[coordinates.Coord]struct{}, dir int) bool {
	proposals := map[coordinates.Coord][]coordinates.Coord{}
	for e := range elves {
		if !shouldMove(elves, e) {
			continue
		}
		for i := 0; i < len(directions); i++ {
			d := directions[(dir+i)%len(directions)]
			if c, ok := canMove(elves, e, d); ok {
				proposals[c] = append(proposals[c], e)
				break
			}
		}
	}

	moved := false
	for p, el := range proposals {
		if len(el) == 1 {
			e := el[0]
			delete(elves, e)
			elves[p] = struct{}{}
			moved = true
		}
	}
	return moved
}

func part1(elves map[coordinates.Coord]struct{}) int {
	dir := 0
	for i := 0; i < 10; i++ {
		round(elves, dir)
		dir = (dir + 1) % len(directions)
	}
	finalElves := make([]coordinates.Coord, len(elves))
	i := 0
	for e := range elves {
		finalElves[i] = e
		i += 1
	}
	min, max := coordinates.Range(finalElves)
	area := (max.X - min.X + 1) * (max.Y - min.Y + 1)
	return area - len(elves)
}

func part2(elves map[coordinates.Coord]struct{}) int {
	i := 10 // Rounds from part 1
	dir := i % len(directions)
	moved := true
	for moved {
		moved = round(elves, dir)
		i += 1
		dir = (dir + 1) % len(directions)
	}
	return i
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	elves := parseInput(data)

	p1 := part1(elves)
	fmt.Println("Part 1:", p1)

	p2 := part2(elves)
	fmt.Println("Part 2:", p2)
}
