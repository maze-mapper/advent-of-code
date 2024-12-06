// Advent of Code 2024 - Day 6.
package day6

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/maze-mapper/advent-of-code/coordinates"
)

func parseData(data []byte) ([][]string, coordinates.Coord) {
	lines := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
	var start coordinates.Coord
	grid := make([][]string, len(lines))
	for i, line := range lines {
		grid[i] = make([]string, len(line))
		for j, r := range line {
			if r == '^' {
				start = coordinates.Coord{X: j, Y: i}
				grid[i][j] = "."
			} else {
				grid[i][j] = string(r)
			}
		}
	}
	return grid, start
}

type direction int

const (
	directionUp = iota
	directionRight
	directionDown
	directionLeft
)

func step(grid [][]string, c coordinates.Coord, d direction) (coordinates.Coord, direction, bool) {
	next := coordinates.Coord{X: c.X, Y: c.Y}
	switch d {
	case directionUp:
		next.Y--
	case directionRight:
		next.X++
	case directionDown:
		next.Y++
	case directionLeft:
		next.X--
	}
	// Check if the guard has left the area.
	if next.Y < 0 || next.X < 0 || next.Y >= len(grid) || next.X >= len(grid[next.Y]) {
		return next, d, true
	}
	if grid[next.Y][next.X] == "#" {
		// Turn right at an obstacle.
		return c, (d + 1) % 4, false
	}
	// Continue moving the same direction.
	return next, d, false
}

func patrol(grid [][]string, c coordinates.Coord) (int, bool) {
	visited := map[coordinates.Coord]map[direction]bool{}
	var d direction = directionUp
	var leftArea bool
	for {
		// Check if the guard is looping.
		if visited[c][d] {
			return len(visited), true
		}
		// Mark location and direction as visited.
		if _, ok := visited[c]; !ok {
			visited[c] = map[direction]bool{}
		}
		visited[c][d] = true
		// Step forward.
		c, d, leftArea = step(grid, c, d)
		if leftArea {
			return len(visited), false
		}
	}
}

func part1(grid [][]string, c coordinates.Coord) int {
	v, _ := patrol(grid, c)
	return v
}

func part2(grid [][]string, c coordinates.Coord) int {
	var count int
	var wg sync.WaitGroup
	var mu sync.Mutex
	for i, row := range grid {
		for j, el := range row {
			if i == c.Y && j == c.X {
				// Cannot place an obstacle at the guard's starting position.
				continue
			}
			if el == "#" {
				// There is already an obstacle here.
				continue
			}
			newGrid := make([][]string, len(grid))
			for ii, arr := range grid {
				newGrid[ii] = make([]string, len(arr))
				copy(newGrid[ii], arr)
			}
			newGrid[i][j] = "#"
			wg.Add(1)
			go func() {
				defer wg.Done()
				_, loop := patrol(newGrid, coordinates.Coord{X: c.X, Y: c.Y})
				if loop {
					mu.Lock()
					count++
					mu.Unlock()
				}
			}()
		}
	}
	wg.Wait()
	return count
}

func Run(inputFile string) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	grid, start := parseData(data)

	p1 := part1(grid, start)
	fmt.Println("Part 1:", p1)

	p2 := part2(grid, start)
	fmt.Println("Part 2:", p2)
}
