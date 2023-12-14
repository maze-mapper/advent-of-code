// Advent of Code 2023 - Day 14.
package day14

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/maze-mapper/advent-of-code/coordinates"
)

type rockArrangement struct {
	fixed, rolling map[coordinates.Coord]bool
	xLen, yLen     int
}

func (ra rockArrangement) String() string {
	var b strings.Builder
	for y := 0; y < ra.yLen; y++ {
		for x := 0; x < ra.xLen; x++ {
			c := coordinates.Coord{X: x, Y: y}
			if ra.fixed[c] {
				b.WriteString("#")
			} else if ra.rolling[c] {
				b.WriteString("O")
			} else {
				b.WriteString(".")
			}
		}
		b.WriteString("\n")
	}
	return b.String()
}

func (ra rockArrangement) totalLoad() int {
	total := 0
	for r := range ra.rolling {
		total += ra.yLen - r.Y
	}
	return total
}

func (ra *rockArrangement) spinCycle() {
	ra.tiltNorth()
	ra.tiltWest()
	ra.tiltSouth()
	ra.tiltEast()
}

func (ra *rockArrangement) tiltNorth() {
	for i := 0; i < ra.yLen; i++ {
		for r := range ra.rolling {
			if r.Y != i {
				continue
			}
			c := coordinates.Coord{X: r.X, Y: i - 1}
			for c.Y >= 0 {
				if ra.rolling[c] || ra.fixed[c] {
					break
				}
				c.Y -= 1
			}
			c.Y += 1
			if c.Y != r.Y {
				delete(ra.rolling, r)
				ra.rolling[c] = true
			}
		}
	}
}

func (ra *rockArrangement) tiltSouth() {
	for i := ra.yLen - 1; i >= 0; i-- {
		for r := range ra.rolling {
			if r.Y != i {
				continue
			}
			c := coordinates.Coord{X: r.X, Y: i + 1}
			for c.Y < ra.yLen {
				if ra.rolling[c] || ra.fixed[c] {
					break
				}
				c.Y += 1
			}
			c.Y -= 1
			if c.Y != r.Y {
				delete(ra.rolling, r)
				ra.rolling[c] = true
			}
		}
	}
}

func (ra *rockArrangement) tiltWest() {
	for i := 0; i < ra.xLen; i++ {
		for r := range ra.rolling {
			if r.X != i {
				continue
			}
			c := coordinates.Coord{X: i - 1, Y: r.Y}
			for c.X >= 0 {
				if ra.rolling[c] || ra.fixed[c] {
					break
				}
				c.X -= 1
			}
			c.X += 1
			if c.X != r.X {
				delete(ra.rolling, r)
				ra.rolling[c] = true
			}
		}
	}
}

func (ra *rockArrangement) tiltEast() {
	for i := ra.xLen - 1; i >= 0; i-- {
		for r := range ra.rolling {
			if r.X != i {
				continue
			}
			c := coordinates.Coord{X: i + 1, Y: r.Y}
			for c.X < ra.xLen {
				if ra.rolling[c] || ra.fixed[c] {
					break
				}
				c.X += 1
			}
			c.X -= 1
			if c.X != r.X {
				delete(ra.rolling, r)
				ra.rolling[c] = true
			}
		}
	}
}

func parseData(data []byte) rockArrangement {
	lines := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
	rocks := rockArrangement{
		fixed:   map[coordinates.Coord]bool{},
		rolling: map[coordinates.Coord]bool{},
		xLen:    len(lines[0]),
		yLen:    len(lines),
	}
	for i, line := range lines {
		for j, r := range line {
			if r == '#' {
				rocks.fixed[coordinates.Coord{X: j, Y: i}] = true
			}
			if r == 'O' {
				rocks.rolling[coordinates.Coord{X: j, Y: i}] = true
			}
		}
	}
	return rocks
}

func part1(ra rockArrangement) int {
	ra.tiltNorth()
	return ra.totalLoad()
}

func part2(ra rockArrangement) int {
	cycles := 1000000000
	arrangements := map[string]int{}
	for cycle := 0; cycle < cycles; cycle++ {
		ra.spinCycle()
		s := ra.String()
		if v, ok := arrangements[s]; ok {
			period := cycle - v
			jumps := (cycles - cycle) / period
			cycle += jumps * period
		} else {
			arrangements[s] = cycle
		}
	}
	return ra.totalLoad()
}

func Run(inputFile string) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	rocks := parseData(data)

	p1 := part1(rocks)
	fmt.Println("Part 1:", p1)

	p2 := part2(rocks)
	fmt.Println("Part 2:", p2)
}
