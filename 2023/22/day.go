// Advent of Code 2023 - Day 22.
package day22

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/maze-mapper/advent-of-code/coordinates"
)

type brick struct {
	c1, c2 coordinates.Coord
}

func (b brick) overlap(other brick) bool {
	return b.c1.X <= other.c2.X && b.c2.X >= other.c1.X &&
		b.c1.Y <= other.c2.Y && b.c2.Y >= other.c1.Y &&
		b.c1.Z <= other.c2.Z && b.c2.Z >= other.c1.Z
}

func parseData(data []byte) []brick {
	lines := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
	bricks := make([]brick, len(lines))
	for i, line := range lines {
		line = strings.ReplaceAll(strings.ReplaceAll(line, ",", " "), "~", " ")
		var c1, c2 coordinates.Coord
		n, err := fmt.Sscanf(line, "%d %d %d %d %d %d", &c1.X, &c1.Y, &c1.Z, &c2.X, &c2.Y, &c2.Z)
		if err != nil {
			log.Fatal(err)
		}
		if n != 6 {
			log.Fatal("Not enough matches")
		}
		bricks[i] = brick{c1: c1, c2: c2}
	}
	return bricks
}

func dropBricks(bricks []brick) ([]brick, int) {
	fallen := make([]bool, len(bricks))
	for !isStable(bricks) {
		for i, b := range bricks {
			drop := 0
		loop:
			for tryZ := b.c1.Z - 1; tryZ > 0; tryZ-- {
				surface := brick{
					c1: coordinates.Coord{X: b.c1.X, Y: b.c1.Y, Z: tryZ},
					c2: coordinates.Coord{X: b.c2.X, Y: b.c2.Y, Z: tryZ},
				}
				for ii, bb := range bricks {
					if ii == i {
						continue
					}
					if surface.overlap(bb) {
						break loop
					}
				}
				drop += 1
			}
			if drop > 0 {
				bricks[i] = brick{
					c1: coordinates.Coord{X: b.c1.X, Y: b.c1.Y, Z: b.c1.Z - drop},
					c2: coordinates.Coord{X: b.c2.X, Y: b.c2.Y, Z: b.c2.Z - drop},
				}
				fallen[i] = true
			}
		}
	}
	totalFallen := 0
	for _, fell := range fallen {
		if fell {
			totalFallen += 1
		}
	}
	return bricks, totalFallen
}

func isStable(bricks []brick) bool {
loop:
	for i, b := range bricks {
		tryZ := b.c1.Z - 1
		if tryZ == 0 {
			// Brick is already on the ground.
			continue
		}
		surface := brick{
			c1: coordinates.Coord{X: b.c1.X, Y: b.c1.Y, Z: tryZ},
			c2: coordinates.Coord{X: b.c2.X, Y: b.c2.Y, Z: tryZ},
		}
		for ii, bb := range bricks {
			if ii == i {
				continue
			}
			if surface.overlap(bb) {
				continue loop
			}
		}
		return false
	}
	return true
}

func part1(bricks []brick) int {
	bricks, _ = dropBricks(bricks)
	count := 0
	for i := range bricks {
		var remaining []brick
		remaining = append(remaining, bricks[:i]...)
		remaining = append(remaining, bricks[i+1:]...)
		if isStable(remaining) {
			count += 1
		}
	}
	return count
}

func part2(bricks []brick) int {
	bricks, _ = dropBricks(bricks)
	total := 0
	for i := range bricks {
		var remaining []brick
		remaining = append(remaining, bricks[:i]...)
		remaining = append(remaining, bricks[i+1:]...)
		_, fell := dropBricks(remaining)
		total += fell
	}
	return total
}

func Run(inputFile string) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	bricks := parseData(data)

	p1 := part1(bricks)
	fmt.Println("Part 1:", p1)

	p2 := part2(bricks)
	fmt.Println("Part 2:", p2)
}
