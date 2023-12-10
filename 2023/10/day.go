// Advent of Code 2023 - Day 10.
package day10

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/maze-mapper/advent-of-code/coordinates"
)

func parseData(data []byte) (coordinates.Coord, [][]rune) {
	lines := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
	pipes := make([][]rune, len(lines))
	var start coordinates.Coord
	for i, line := range lines {
		pipes[i] = make([]rune, len(line))
		for j, r := range line {
			pipes[i][j] = r
			if r == 'S' {
				start = coordinates.Coord{X: j, Y: i}
			}
		}
	}
	return start, pipes
}

type cardinalMove struct {
	nextPosition coordinates.Coord
	direction    int
}

func cardinalNeighbours(c coordinates.Coord) []cardinalMove {
	moves := []cardinalMove{
		{
			nextPosition: coordinates.Coord{Y: -1},
			direction:    north,
		},
		{
			nextPosition: coordinates.Coord{X: 1},
			direction:    east,
		},
		{
			nextPosition: coordinates.Coord{Y: 1},
			direction:    south,
		},
		{
			nextPosition: coordinates.Coord{X: -1},
			direction:    west,
		},
	}
	for i := range moves {
		moves[i].nextPosition.Transform(c)
	}
	return moves
}

const (
	north = iota
	east
	south
	west
)

func validMove(direction int, pipe rune) (int, bool) {
	switch pipe {
	case '|':
		if direction == north || direction == south {
			return direction, true
		}
	case '-':
		if direction == east || direction == west {
			return direction, true
		}
	case 'L':
		if direction == south {
			return east, true
		} else if direction == west {
			return north, true
		}
	case 'J':
		if direction == south {
			return west, true
		} else if direction == east {
			return north, true
		}
	case '7':
		if direction == north {
			return west, true
		} else if direction == east {
			return south, true
		}
	case 'F':
		if direction == north {
			return east, true
		} else if direction == west {
			return south, true
		}
	}
	return direction, false
}

func leftAndRight(direction int, pipe rune) ([]coordinates.Coord, []coordinates.Coord) {
	switch pipe {
	case '|':
		left, right := []coordinates.Coord{{X: -1}}, []coordinates.Coord{{X: 1}}
		if direction == north {
			return left, right
		} else if direction == south {
			return right, left
		}
	case '-':
		top, bottom := []coordinates.Coord{{Y: -1}}, []coordinates.Coord{{Y: 1}}
		if direction == east {
			return top, bottom
		} else if direction == west {
			return bottom, top
		}
	case 'L':
		outer := []coordinates.Coord{{X: -1}, {Y: 1}}
		if direction == north {
			return outer, nil
		} else if direction == east {
			return nil, outer
		}
	case 'J':
		outer := []coordinates.Coord{{X: 1}, {Y: 1}}
		if direction == north {
			return nil, outer
		} else if direction == west {
			return outer, nil
		}
	case '7':
		outer := []coordinates.Coord{{X: 1}, {Y: -1}}
		if direction == south {
			return outer, nil
		} else if direction == west {
			return nil, outer
		}
	case 'F':
		outer := []coordinates.Coord{{X: -1}, {Y: -1}}
		if direction == east {
			return outer, nil
		} else if direction == south {
			return nil, outer
		}
	}
	return nil, nil
}

func startingPipeAndDirection(start coordinates.Coord, pipes [][]rune) (rune, int) {
	var allowedDirections []int
	for _, move := range cardinalNeighbours(start) {
		if move.nextPosition.X < 0 || move.nextPosition.Y < 0 || move.nextPosition.X >= len(pipes[0]) || move.nextPosition.Y >= len(pipes) {
			continue
		}
		if _, ok := validMove(move.direction, pipes[move.nextPosition.Y][move.nextPosition.X]); ok {
			allowedDirections = append(allowedDirections, move.direction)
		}
	}
	if len(allowedDirections) != 2 {
		log.Fatal("Not enough pipe directions")
	}

	pipe := 'S'
	switch {
	case allowedDirections[0] == north && allowedDirections[1] == south:
		pipe = '|'
	case allowedDirections[0] == east && allowedDirections[1] == west:
		pipe = '-'
	case allowedDirections[0] == north && allowedDirections[1] == east:
		pipe = 'L'
	case allowedDirections[0] == north && allowedDirections[1] == west:
		pipe = 'J'
	case allowedDirections[0] == south && allowedDirections[1] == west:
		pipe = '7'
	case allowedDirections[0] == east && allowedDirections[1] == south:
		pipe = 'F'
	}
	return pipe, allowedDirections[0]
}

func part1(c coordinates.Coord, pipes [][]rune, visted map[coordinates.Coord]bool, direction int) int {
	for !visted[c] {
		visted[c] = true
		switch direction {
		case north:
			c.Transform(coordinates.Coord{Y: -1})
		case east:
			c.Transform(coordinates.Coord{X: 1})
		case south:
			c.Transform(coordinates.Coord{Y: 1})
		case west:
			c.Transform(coordinates.Coord{X: -1})
		}
		d, ok := validMove(direction, pipes[c.Y][c.X])
		if !ok {
			log.Fatal("Not a valid move")
		}
		direction = d
	}

	return len(visted) / 2
}

func part2(c coordinates.Coord, pipes [][]rune, mainLoop map[coordinates.Coord]bool, direction int) int {
	left, right, visted := map[coordinates.Coord]bool{}, map[coordinates.Coord]bool{}, map[coordinates.Coord]bool{}
	for !visted[c] {
		visted[c] = true

		leftPoints, rightPoints := leftAndRight(direction, pipes[c.Y][c.X])
		for _, p := range leftPoints {
			p.Transform(c)
			if p.X < 0 || p.Y < 0 || p.X >= len(pipes[0]) || p.Y >= len(pipes) {
				continue
			}
			if !mainLoop[p] {
				left[p] = true
			}
		}
		for _, p := range rightPoints {
			p.Transform(c)
			if p.X < 0 || p.Y < 0 || p.X >= len(pipes[0]) || p.Y >= len(pipes) {
				continue
			}
			if !mainLoop[p] {

				right[p] = true
			}
		}

		switch direction {
		case north:
			c.Transform(coordinates.Coord{Y: -1})
		case east:
			c.Transform(coordinates.Coord{X: 1})
		case south:
			c.Transform(coordinates.Coord{Y: 1})
		case west:
			c.Transform(coordinates.Coord{X: -1})
		}
		d, ok := validMove(direction, pipes[c.Y][c.X])
		if !ok {
			log.Fatal("Not a valid move")
		}
		direction = d
	}

	// testPrint(left, right, mainLoop, pipes)

	floodFill(left, mainLoop, len(pipes[0])-1, len(pipes)-1)
	floodFill(right, mainLoop, len(pipes[0])-1, len(pipes)-1)

	// testPrint(left, right, mainLoop, pipes)

	for k := range left {
		if k.X == 0 || k.X == len(pipes[0])-1 || k.Y == 0 || k.Y == len(pipes)-1 {
			return len(right)
		}
	}
	for k := range right {
		if k.X == 0 || k.X == len(pipes[0])-1 || k.Y == 0 || k.Y == len(pipes)-1 {
			return len(left)
		}
	}
	fmt.Printf("Unable to determine outside. Left points: %d, right points: %d\n", len(left), len(right))

	return 0
}

func floodFill(points, mainLoop map[coordinates.Coord]bool, xMax, yMax int) {
	for p := range points {
		for _, neighbour := range cardinalNeighbours(p) {
			if neighbour.nextPosition.X < 0 || neighbour.nextPosition.Y < 0 || neighbour.nextPosition.X > xMax || neighbour.nextPosition.Y > yMax {
				continue
			}
			if mainLoop[neighbour.nextPosition] {
				continue
			}
			if points[neighbour.nextPosition] {
				continue
			}
			points[neighbour.nextPosition] = true
			floodFill(points, mainLoop, xMax, yMax)
		}
	}
}

func testPrint(left, right, mainLoop map[coordinates.Coord]bool, pipes [][]rune) {
	for y, line := range pipes {
		for x, r := range line {
			c := coordinates.Coord{X: x, Y: y}
			if left[c] {
				fmt.Print("A")
			} else if right[c] {
				fmt.Print("B")
			} else if mainLoop[c] {
				fmt.Print(string(r))
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}

func Run(inputFile string) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	start, pipes := parseData(data)

	pipe, direction := startingPipeAndDirection(start, pipes)
	if pipe == 'S' {
		log.Fatal("Could not determine starting pipe")
	}
	pipes[start.Y][start.X] = pipe

	mainLoop := map[coordinates.Coord]bool{}

	p1 := part1(start, pipes, mainLoop, direction)
	fmt.Println("Part 1:", p1)

	p2 := part2(start, pipes, mainLoop, direction)
	fmt.Println("Part 2:", p2)
}
