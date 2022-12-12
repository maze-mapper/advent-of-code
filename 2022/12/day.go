// Advent of Code 2022 - Day 12.
package day12

import (
	"bytes"
	"container/heap"
	"fmt"
	"io/ioutil"
	"log"
	"math"

	"github.com/maze-mapper/advent-of-code/coordinates"
)

func parseInput(data []byte) ([][]uint8, coordinates.Coord, coordinates.Coord) {
	lines := bytes.Split(bytes.TrimSuffix(data, []byte("\n")), []byte("\n"))
	hill := make([][]uint8, len(lines))
	var start coordinates.Coord
	var end coordinates.Coord
	for i, line := range lines {
		hill[i] = make([]uint8, len(line))
		for j, b := range line {
			switch b {
			case byte('S'):
				start = coordinates.Coord{X: j, Y: i}
				hill[i][j] = 0
			case byte('E'):
				end = coordinates.Coord{X: j, Y: i}
				hill[i][j] = byte('z') - byte('a')
			default:
				hill[i][j] = b - byte('a')
			}
		}
	}
	return hill, start, end
}

func neighbours(hill [][]uint8, c coordinates.Coord) []coordinates.Coord {
	n := []coordinates.Coord{}
	if c.X > 0 {
		p := coordinates.Coord{X: c.X - 1, Y: c.Y}
		n = append(n, p)
	}
	if c.X < len(hill[c.Y])-1 {
		p := coordinates.Coord{X: c.X + 1, Y: c.Y}
		n = append(n, p)
	}
	if c.Y > 0 {
		p := coordinates.Coord{X: c.X, Y: c.Y - 1}
		n = append(n, p)
	}
	if c.Y < len(hill)-1 {
		p := coordinates.Coord{X: c.X, Y: c.Y + 1}
		n = append(n, p)
	}

	height := hill[c.Y][c.X]
	out := []coordinates.Coord{}
	for _, p := range n {
		pHeight := hill[p.Y][p.X]
		if pHeight <= height+1 { // Assuming climbing only.
			out = append(out, p)
		}
	}
	return out
}

type Node struct {
	c        coordinates.Coord
	step     int
	priority int
	index    int
}

// A PriorityQueue implements heap.Interface and holds Nodes.
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the lowest, not highest, priority so we use less than here.
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	node := x.(*Node)
	node.index = n
	*pq = append(*pq, node)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	old[n-1] = nil  // avoid memory leak
	node.index = -1 // for safety
	*pq = old[0 : n-1]
	return node
}

func AStar(hill [][]uint8, start coordinates.Coord, end coordinates.Coord) (int, bool) {
	// Initialise priority queue
	pq := make(PriorityQueue, 1)
	startNode := &Node{
		c:        start,
		step:     0,
		priority: 0,
		index:    0,
	}
	pq[0] = startNode
	heap.Init(&pq)

	visited := map[coordinates.Coord]int{
		start: 0,
	}

	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*Node)

		if node.c == end {
			return node.step, true
		}

		moves := neighbours(hill, node.c)
		nextStep := node.step + 1
		for _, move := range moves {
			if step, ok := visited[move]; !ok || nextStep < step {
				heap.Push(&pq, &Node{
					c:        move,
					step:     nextStep,
					priority: nextStep + coordinates.ManhattanDistance(move, end),
				})
				visited[move] = nextStep
			}

		}
	}
	return 0, false
}

func part1(hill [][]uint8, start coordinates.Coord, end coordinates.Coord) int {
	steps, _ := AStar(hill, start, end)
	return steps
}

func part2(hill [][]uint8, end coordinates.Coord) int {
	minSteps := int(math.MaxInt)
	for y, row := range hill {
		for x, p := range row {
			if p == 0 {
				if steps, ok := AStar(hill, coordinates.Coord{X: x, Y: y}, end); ok && steps < minSteps {
					minSteps = steps
				}
			}
		}
	}
	return minSteps
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	hill, start, end := parseInput(data)

	p1 := part1(hill, start, end)
	fmt.Println("Part 1:", p1)

	p2 := part2(hill, end)
	fmt.Println("Part 2:", p2)
}
