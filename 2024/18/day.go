// Advent of Code 2024 - Day 18.
package day18

import (
	"container/heap"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/maze-mapper/advent-of-code/coordinates"
)

func parseData(data []byte) (map[coordinates.Coord]int, error) {
	lines := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
	m := map[coordinates.Coord]int{}
	for t, line := range lines {
		before, after, found := strings.Cut(line, ",")
		if !found {
			return nil, fmt.Errorf("unable to parse line %q", line)
		}
		x, err := strconv.Atoi(before)
		if err != nil {
			return nil, err
		}
		y, err := strconv.Atoi(after)
		if err != nil {
			return nil, err
		}
		m[coordinates.Coord{X: x, Y: y}] = t + 1
	}
	return m, nil
}

func part1(fallingBytes map[coordinates.Coord]int, xMax, yMax, fallenBytes int) int {
	start := coordinates.Coord{X: 0, Y: 0}
	end := coordinates.Coord{X: xMax, Y: yMax}

	pq := make(PriorityQueue, 1)
	pq[0] = &Item{
		state: gameState{
			position: start,
			steps:    0,
		},
		priority: 0,
	}
	heap.Init(&pq)

	visited := map[coordinates.Coord]int{}
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)

		if item.state.position == end {
			return item.state.steps
		}

		if s, ok := visited[item.state.position]; ok && s <= item.state.steps {
			continue
		}
		visited[item.state.position] = item.state.steps

		nextStates := []gameState{
			{
				position: coordinates.Coord{X: item.state.position.X + 1, Y: item.state.position.Y},
				steps:    item.state.steps + 1,
			},
			{
				position: coordinates.Coord{X: item.state.position.X - 1, Y: item.state.position.Y},
				steps:    item.state.steps + 1,
			},
			{
				position: coordinates.Coord{X: item.state.position.X, Y: item.state.position.Y + 1},
				steps:    item.state.steps + 1,
			},
			{
				position: coordinates.Coord{X: item.state.position.X, Y: item.state.position.Y - 1},
				steps:    item.state.steps + 1,
			},
		}

		for _, state := range nextStates {
			// Check bounds.
			if state.position.X < 0 || state.position.X > xMax || state.position.Y < 0 || state.position.Y > yMax {
				continue
			}
			// Check if the position is not within the first fallenBytes of fallign bytes.
			if t, ok := fallingBytes[state.position]; ok && t <= fallenBytes {
				continue
			}
			// Check visited.
			if s, ok := visited[state.position]; ok && state.steps >= s {
				continue
			}
			heap.Push(&pq, &Item{
				state:    state,
				priority: state.steps + coordinates.ManhattanDistance(end, state.position),
			})
		}
	}

	return 0
}

func part2(fallingBytes map[coordinates.Coord]int, xMax, yMax, fallenBytes int) string {
	for b := fallenBytes + 1; b <= len(fallingBytes); b++ {
		if steps := part1(fallingBytes, xMax, yMax, b); steps == 0 {
			for k, v := range fallingBytes {
				if v == b {
					return fmt.Sprintf("%d,%d", k.X, k.Y)
				}
			}
		}
	}
	return ""
}

type gameState struct {
	position coordinates.Coord
	steps    int
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
	fallingBytes, err := parseData(data)
	if err != nil {
		log.Fatal(err)
	}

	xMax := 70
	yMax := 70
	fallenBytes := 1024

	p1 := part1(fallingBytes, xMax, yMax, fallenBytes)
	fmt.Println("Part 1:", p1)

	p2 := part2(fallingBytes, xMax, yMax, fallenBytes)
	fmt.Println("Part 2:", p2)
}
