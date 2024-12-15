// Advent of Code 2024 - Day 15.
package day15

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/maze-mapper/advent-of-code/coordinates"
)

func parseData(data []byte) (map[coordinates.Coord]rune, coordinates.Coord, []rune) {
	sections := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n\n")
	m := map[coordinates.Coord]rune{}
	var start coordinates.Coord
	rows := strings.Split(sections[0], "\n")
	for y, row := range rows {
		for x, cell := range row {
			switch cell {
			case '@':
				start = coordinates.Coord{X: x, Y: y}
			case '#', 'O':
				m[coordinates.Coord{X: x, Y: y}] = cell
			}
		}
	}

	directionString := strings.ReplaceAll(sections[1], "\n", "")
	var directions []rune
	for _, r := range directionString {
		directions = append(directions, r)
	}
	return m, start, directions
}

func part1(positions map[coordinates.Coord]rune, robot coordinates.Coord, directions []rune) int {
	var direction coordinates.Coord
	for _, d := range directions {
		switch d {
		case '^':
			direction = coordinates.Coord{X: 0, Y: -1}
		case '>':
			direction = coordinates.Coord{X: 1, Y: 0}
		case 'v':
			direction = coordinates.Coord{X: 0, Y: 1}
		case '<':
			direction = coordinates.Coord{X: -1, Y: 0}
		}

		nextPos := coordinates.Coord{
			X: robot.X + direction.X,
			Y: robot.Y + direction.Y,
		}
		k, ok := positions[nextPos]
		// Move robot in to empty space.
		if !ok {
			robot = nextPos
			continue
		}
		// Can't move due to wall
		if k == '#' {
			continue
		}
		// Must be a box.
		boxPos := nextPos
		for {
			// Shift box if possible.
			boxPos = coordinates.Coord{
				X: boxPos.X + direction.X,
				Y: boxPos.Y + direction.Y,
			}
			kk, ok := positions[boxPos]
			if !ok {
				robot = nextPos
				delete(positions, nextPos)
				positions[boxPos] = 'O'
				break
			}
			if kk == '#' {
				break
			}
		}
	}
	return gps(positions, 'O')
}

func part2(positions map[coordinates.Coord]rune, robot coordinates.Coord, directions []rune) int {
	var direction coordinates.Coord
	for _, d := range directions {
		switch d {
		case '^':
			direction = coordinates.Coord{X: 0, Y: -1}
		case '>':
			direction = coordinates.Coord{X: 1, Y: 0}
		case 'v':
			direction = coordinates.Coord{X: 0, Y: 1}
		case '<':
			direction = coordinates.Coord{X: -1, Y: 0}
		}

		nextPos := coordinates.Coord{
			X: robot.X + direction.X,
			Y: robot.Y + direction.Y,
		}
		k, ok := positions[nextPos]
		// Move robot in to empty space.
		if !ok {
			robot = nextPos
			continue
		}
		// Can't move due to wall
		if k == '#' {
			continue
		}
		// Must be a box.
		boxPos := nextPos
		for {
			if d == '>' || d == '<' { // Horizontal move.
				boxPos = coordinates.Coord{
					X: boxPos.X + direction.X,
					Y: boxPos.Y + direction.Y,
				}
				kk, ok := positions[boxPos]
				if !ok {
					robot = nextPos
					delete(positions, nextPos)
					step := 0
					// Move boxes right.
					if d == '>' {
						for x := nextPos.X + 1; x <= boxPos.X; x++ {
							if step%2 == 0 {
								positions[coordinates.Coord{X: x, Y: boxPos.Y}] = '['
							} else {
								positions[coordinates.Coord{X: x, Y: boxPos.Y}] = ']'
							}
							step++
						}
					}
					// Move boxes left.
					if d == '<' {
						for x := nextPos.X - 1; x >= boxPos.X; x-- {
							if step%2 == 0 {
								positions[coordinates.Coord{X: x, Y: boxPos.Y}] = ']'
							} else {
								positions[coordinates.Coord{X: x, Y: boxPos.Y}] = '['
							}
							step++
						}
					}
					break
				}
				if kk == '#' {
					break
				}
			} else { // Vertical move.
				canMove := true
				if positions[boxPos] == ']' {
					boxPos = coordinates.Coord{X: boxPos.X - 1, Y: boxPos.Y}
				}
				allBoxes := []coordinates.Coord{boxPos}
				currentBoxes := []coordinates.Coord{boxPos}
				for len(currentBoxes) > 0 {
					b := currentBoxes[0]
					currentBoxes = currentBoxes[1:]
					p1 := coordinates.Coord{
						X: b.X,
						Y: b.Y + direction.Y,
					}
					p2 := coordinates.Coord{
						X: b.X + 1,
						Y: b.Y + direction.Y,
					}
					k1, ok := positions[p1]
					if ok {
						if k1 == '#' {
							canMove = false
							break
						}
						if k1 == '[' {
							allBoxes = append(allBoxes, p1)
							currentBoxes = append(currentBoxes, p1)
							continue // No need to check p2 as that's the end of the same box.
						}
						if k1 == ']' {
							allBoxes = append(allBoxes, coordinates.Coord{X: p1.X - 1, Y: p2.Y})
							currentBoxes = append(currentBoxes, coordinates.Coord{X: p1.X - 1, Y: p2.Y})
						}
					}
					k2, ok := positions[p2]
					if ok {
						if k2 == '#' {
							canMove = false
							break
						}
						if k2 == '[' {
							allBoxes = append(allBoxes, p2)
							currentBoxes = append(currentBoxes, p2)
						}
					}
				}
				if canMove {
					for _, box := range allBoxes {
						delete(positions, box)
						delete(positions, coordinates.Coord{X: box.X + 1, Y: box.Y})
					}
					for _, box := range allBoxes {
						positions[coordinates.Coord{X: box.X, Y: box.Y + direction.Y}] = '['
						positions[coordinates.Coord{X: box.X + 1, Y: box.Y + direction.Y}] = ']'
					}
					robot = nextPos
				}
				break
			}
		}

	}
	return gps(positions, '[')
}

func scaleWidth(positions map[coordinates.Coord]rune, start coordinates.Coord) (map[coordinates.Coord]rune, coordinates.Coord) {
	newPositions := map[coordinates.Coord]rune{}
	for k, v := range positions {
		p1 := coordinates.Coord{X: 2 * k.X, Y: k.Y}
		p2 := coordinates.Coord{X: 2*k.X + 1, Y: k.Y}
		if v == 'O' {
			newPositions[p1] = '['
			newPositions[p2] = ']'
		} else {
			newPositions[p1] = v
			newPositions[p2] = v
		}
	}
	return newPositions, coordinates.Coord{X: 2 * start.X, Y: start.Y}
}

func gps(positions map[coordinates.Coord]rune, box rune) int {
	var sum int
	for k, v := range positions {
		if v == box {
			sum += 100*k.Y + k.X
		}
	}
	return sum
}

func Run(inputFile string) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	positions, start, directions := parseData(data)
	positions2, start2 := scaleWidth(positions, start)

	p1 := part1(positions, start, directions)
	fmt.Println("Part 1:", p1)

	p2 := part2(positions2, start2, directions)
	fmt.Println("Part 2:", p2)
}
