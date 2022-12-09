// Advent of Code 2022 - Day 9.
package day9

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/maze-mapper/advent-of-code/coordinates"
)

type move struct {
	direction string
	distance  int
}

func parseInput(data []byte) []move {
	lines := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
	moves := make([]move, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " ")
		dist, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}
		moves[i] = move{direction: parts[0], distance: dist}
	}
	return moves
}

func moveRope(moves []move, knotCount int) map[coordinates.Coord]struct{} {
	knots := make([]*coordinates.Coord, knotCount)
	for i := range knots {
		knots[i] = &coordinates.Coord{}
	}
	visited := map[coordinates.Coord]struct{}{
		coordinates.Coord{}: struct{}{},
	}
	for _, m := range moves {
	//	fmt.Println("Move", m)
		for d := 0; d < m.distance; d++ {
			// Move head.
			v := coordinates.Coord{}
			switch m.direction {
			case "U":
				v = coordinates.Coord{Y: 1}
			case "D":
				v = coordinates.Coord{Y: -1}
			case "L":
				v = coordinates.Coord{X: -1}
			case "R":
				v = coordinates.Coord{X: 1}
			}
			knots[0].Transform(v)

			// Move other knots.
			loop:
			for i := 1; i < knotCount; i++ {
				previous := knots[i-1]
				current := knots[i]
				sep := coordinates.Coord{
					X: previous.X - current.X,
					Y: previous.Y - current.Y,
				}
//				fmt.Println(i, previous, current, sep)
				switch {
				case sep.X == 0 && sep.Y > 1:
					current.Y += 1
				case sep.X == 0 && sep.Y < -1:
					current.Y -= 1
				case sep.Y == 0 && sep.X > 1:
					current.X += 1
				case sep.Y == 0 && sep.X < -1:
					current.X -= 1
				case (sep.X >= 1 && sep.Y > 1) || (sep.Y >= 1 && sep.X > 1):
					current.X += 1
					current.Y += 1
				case (sep.X <= -1 && sep.Y > 1) || (sep.Y >= 1 && sep.X < -1):
					current.X -= 1
					current.Y += 1
				case (sep.X >= 1 && sep.Y < -1) || (sep.Y <= -1 && sep.X > 1):
					current.X += 1
					current.Y -= 1
				case (sep.X <= -1 && sep.Y < -1) || (sep.Y <= -1 && sep.X < -1):
					current.X -= 1
					current.Y -= 1
				default:
					// Current knot did not move so skip others.
					break loop

				}
//				fmt.Println(i, previous, current, sep)
				pos := make([]coordinates.Coord, knotCount)
                		for i, k := range knots {
		                        pos[i] = *k
                		}
//				fmt.Println(pos)


			}
			tail := *knots[knotCount-1]
//			fmt.Println(tail)
			visited[tail] = struct{}{}
		}

/*		pos := make([]coordinates.Coord, knotCount)
		for i, k := range knots {
			pos[i] = *k
		}
		fmt.Println(pos)*/
	}
	return visited
}

func part1(moves []move) int {
	visited := moveRope(moves, 2)
	return len(visited)
}

func part2(moves []move) int {
	visited := moveRope(moves, 10)
	return len(visited)
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	moves := parseInput(data)

	p1 := part1(moves)
	fmt.Println("Part 1:", p1)

	p2 := part2(moves)
	fmt.Println("Part 2:", p2)
}
