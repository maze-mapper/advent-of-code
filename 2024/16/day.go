// Advent of Code 2024 - Day 16.
package day16

import (
	"container/heap"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/maze-mapper/advent-of-code/coordinates"
)

func parseData(data []byte) ([][]bool, coordinates.Coord, coordinates.Coord) {
	lines := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
	var start, end coordinates.Coord
	maze := make([][]bool, len(lines))
	for i, line := range lines {
		maze[i] = make([]bool, len(line))
		for j, r := range line {
			if r == 'S' {
				start = coordinates.Coord{X: j, Y: i}
			}
			if r == 'E' {
				end = coordinates.Coord{X: j, Y: i}
			}
			if r == '#' {
				maze[i][j] = true
			}
		}
	}
	return maze, start, end
}

func part1(maze [][]bool, start, end coordinates.Coord) int {
	pq := make(PriorityQueue, 1)
	pq[0] = &Item{
		state: gameState{
			position: start,
			facing:   directionEast,
			score:    0,
		},
		priority: 0,
	}
	heap.Init(&pq)

	visited := map[string]int{}

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)

		if item.state.position == end {
			return item.state.score
		}

		if s, ok := visited[item.state.key()]; ok && s <= item.state.score {
			continue
		}
		visited[item.state.key()] = item.state.score

		// Next moves.
		// Turn left or right.
		nextStates := []gameState{
			{
				position: item.state.position,
				facing:   (item.state.facing + 1) % 4,
				score:    item.state.score + 1000,
			},
			{
				position: item.state.position,
				facing:   (item.state.facing + 3) % 4,
				score:    item.state.score + 1000,
			},
		}
		// Move in the facing direction.
		c := coordinates.Coord{X: item.state.position.X, Y: item.state.position.Y}
		switch item.state.facing {
		case directionEast:
			c.Transform(coordinates.Coord{X: 1})
		case directionSouth:
			c.Transform(coordinates.Coord{Y: 1})
		case directionWest:
			c.Transform(coordinates.Coord{X: -1})
		case directionNorth:
			c.Transform(coordinates.Coord{Y: -1})
		}
		if isWall := maze[c.Y][c.X]; !isWall {
			nextStates = append(nextStates, gameState{
				position: c,
				facing:   item.state.facing,
				score:    item.state.score + 1,
			})
		}

		for _, state := range nextStates {
			if s, ok := visited[state.key()]; ok && s <= state.score {
				continue
			}
			heap.Push(&pq, &Item{
				state:    state,
				priority: state.score + coordinates.ManhattanDistance(end, state.position),
			})
		}
	}
	return 0
}

func part2(maze [][]bool, start, end coordinates.Coord) int {
	pq := make(PriorityQueue, 1)
	pq[0] = &Item{
		state: gameState{
			position: start,
			facing:   directionEast,
			score:    0,
			path:     []coordinates.Coord{start},
		},
		priority: 0,
	}
	heap.Init(&pq)

	visited := map[string]int{}
	reachedEnd := false
	var bestScore int
	bestPathTiles := map[coordinates.Coord]bool{}

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)

		if reachedEnd && item.state.score > bestScore {
			continue
		}

		if item.state.position == end {
			if !reachedEnd {
				reachedEnd = true
				bestScore = item.state.score
			}
			for _, tile := range item.state.path {
				bestPathTiles[tile] = true
			}
			continue
		}

		if s, ok := visited[item.state.key()]; ok && s < item.state.score {
			continue
		}
		visited[item.state.key()] = item.state.score

		// Next moves.
		// Turn left or right.
		nextStates := []gameState{
			{
				position: item.state.position,
				facing:   (item.state.facing + 1) % 4,
				score:    item.state.score + 1000,
				path:     item.state.path,
			},
			{
				position: item.state.position,
				facing:   (item.state.facing + 3) % 4,
				score:    item.state.score + 1000,
				path:     item.state.path,
			},
		}
		// Move in the facing direction.
		c := coordinates.Coord{X: item.state.position.X, Y: item.state.position.Y}
		switch item.state.facing {
		case directionEast:
			c.Transform(coordinates.Coord{X: 1})
		case directionSouth:
			c.Transform(coordinates.Coord{Y: 1})
		case directionWest:
			c.Transform(coordinates.Coord{X: -1})
		case directionNorth:
			c.Transform(coordinates.Coord{Y: -1})
		}
		if isWall := maze[c.Y][c.X]; !isWall {
			path := make([]coordinates.Coord, len(item.state.path))
			copy(path, item.state.path)
			path = append(path, c)
			nextStates = append(nextStates, gameState{
				position: c,
				facing:   item.state.facing,
				score:    item.state.score + 1,
				path:     path,
			})
		}

		for _, state := range nextStates {
			if s, ok := visited[state.key()]; ok && s < state.score {
				continue
			}
			heap.Push(&pq, &Item{
				state:    state,
				priority: state.score + coordinates.ManhattanDistance(end, state.position),
			})
		}
	}
	return len(bestPathTiles)
}

type direction int

const (
	directionEast = iota
	directionSouth
	directionWest
	directionNorth
)

type gameState struct {
	position coordinates.Coord
	facing   direction
	score    int
	path     []coordinates.Coord
}

func (s gameState) key() string {
	return fmt.Sprintf("%d-%d-%d", s.position.X, s.position.Y, s.facing)
}

// An Item is something we manage in a priority queue.
type Item struct {
	state    gameState
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the lowest, not highest, priority so we use less than here.
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func Run(inputFile string) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	maze, start, end := parseData(data)

	p1 := part1(maze, start, end)
	fmt.Println("Part 1:", p1)

	p2 := part2(maze, start, end)
	fmt.Println("Part 2:", p2)
}
