// Advent of Code 2024 - Day 12.
package day12

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func parseData(data []byte) [][]string {
	lines := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
	region := make([][]string, len(lines))
	for i, line := range lines {
		region[i] = strings.Split(line, "")
	}
	return region
}

func part1(region [][]string) int {
	visited := make([][]bool, len(region))
	for i := range region {
		visited[i] = make([]bool, len(region[i]))
	}

	var totalPrice int
	for i, row := range region {
		for j, plant := range row {
			var area int
			var perimeter int
			plotSpaces := [][2]int{{i, j}}
			for len(plotSpaces) > 0 {
				space := plotSpaces[0]
				si := space[0]
				sj := space[1]
				plotSpaces = plotSpaces[1:]
				if visited[si][sj] {
					continue
				}
				visited[si][sj] = true
				area += 1
				for _, neighbour := range [][2]int{
					{si - 1, sj}, {si + 1, sj}, {si, sj - 1}, {si, sj + 1},
				} {
					ii := neighbour[0]
					jj := neighbour[1]
					if ii < 0 || ii >= len(region) || jj < 0 || jj >= len(region[ii]) || region[ii][jj] != plant {
						perimeter += 1
					} else {
						plotSpaces = append(plotSpaces, [2]int{ii, jj})
					}
				}
			}
			totalPrice += area * perimeter
		}
	}
	return totalPrice
}

func part2(region [][]string) int {
	visited := make([][]bool, len(region))
	for i := range region {
		visited[i] = make([]bool, len(region[i]))
	}

	var totalPrice int
	for i, row := range region {
		for j, plant := range row {
			var area int
			edgeSpaces := map[[2]int]bool{}
			outerSpaces := map[[2]int]bool{}
			plotSpaces := [][2]int{{i, j}}
			for len(plotSpaces) > 0 {
				space := plotSpaces[0]
				si := space[0]
				sj := space[1]
				plotSpaces = plotSpaces[1:]
				if visited[si][sj] {
					continue
				}
				visited[si][sj] = true
				area += 1
				for _, neighbour := range [][2]int{
					{si - 1, sj}, {si + 1, sj}, {si, sj - 1}, {si, sj + 1},
				} {
					ii := neighbour[0]
					jj := neighbour[1]
					if ii < 0 || ii >= len(region) || jj < 0 || jj >= len(region[ii]) || region[ii][jj] != plant {
						edgeSpaces[space] = true
						outerSpaces[neighbour] = true
					} else {
						plotSpaces = append(plotSpaces, [2]int{ii, jj})
					}
				}
			}
			edges := countEdges(edgeSpaces, outerSpaces)
			totalPrice += area * edges
		}
	}
	return totalPrice
}

type direction int

const (
	directionRight = iota
	directionUp
	directionLeft
	directionDown
)

func turnLeft(d direction) direction {
	return (d + 1) % 4
}

func turnRight(d direction) direction {
	return (d + 3) % 4
}

func nextSpace(current [2]int, d direction) [2]int {
	switch d {
	case directionUp:
		current[0]--
	case directionRight:
		current[1]++
	case directionLeft:
		current[1]--
	case directionDown:
		current[0]++
	}
	return current
}

func countEdges(edgeSpaces, outerSpaces map[[2]int]bool) int {
	visitedOuterSpaces := map[[2]int]bool{}
	var edgeCount int

	for len(visitedOuterSpaces) != len(outerSpaces) {

		// Find any unvisited outer edge space.
		var currentOuterSpace [2]int
		for k := range outerSpaces {
			if !visitedOuterSpaces[k] {
				currentOuterSpace = k
				break
			}
		}

		// Find the direction of the inside and get that point. Turn left from there.
		var currentInnerSpace [2]int
		var currentDirection direction
		for _, d := range []direction{directionRight, directionUp, directionLeft, directionDown} {
			sp := nextSpace(currentOuterSpace, d)
			if edgeSpaces[sp] {
				currentInnerSpace = sp
				currentDirection = turnLeft(d)
				break
			}
		}

		// Walk the edge.
		endSpace := [2]int{currentInnerSpace[0], currentInnerSpace[1]}
		endDirection := currentDirection
		for {
			visitedOuterSpaces[currentOuterSpace] = true

			// Possible turns, keeping the plot on your right.
			// The left turn is a diagonal move.
			i := currentInnerSpace[0]
			j := currentInnerSpace[1]
			var leftTurn [2]int
			var noTurn [2]int
			switch currentDirection {
			case directionRight:
				leftTurn = [2]int{i - 1, j + 1}
				noTurn = [2]int{i, j + 1}
			case directionUp:
				leftTurn = [2]int{i - 1, j - 1}
				noTurn = [2]int{i - 1, j}
			case directionLeft:
				leftTurn = [2]int{i + 1, j - 1}
				noTurn = [2]int{i, j - 1}
			case directionDown:
				leftTurn = [2]int{i + 1, j + 1}
				noTurn = [2]int{i + 1, j}
			}

			switch {
			case edgeSpaces[leftTurn]:
				edgeCount += 1
				currentInnerSpace = leftTurn
				currentDirection = turnLeft(currentDirection)
			case edgeSpaces[noTurn]:
				currentInnerSpace = noTurn
			default:
				edgeCount += 1
				currentDirection = turnRight(currentDirection)
			}
			// Outer space to the left of the current edge.
			currentOuterSpace = nextSpace(currentInnerSpace, turnLeft(currentDirection))

			// Check if we are back where we started.
			if currentInnerSpace == endSpace && currentDirection == endDirection {
				break
			}
		}
	}

	return edgeCount
}

func Run(inputFile string) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	region := parseData(data)

	p1 := part1(region)
	fmt.Println("Part 1:", p1)

	p2 := part2(region)
	fmt.Println("Part 2:", p2)
}
