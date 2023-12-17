// Advent of Code 2023 - Day 17.
package day17

import (
	"container/heap"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/maze-mapper/advent-of-code/coordinates"
)

func parseData(data []byte) [][]int {
	lines := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
	area := make([][]int, len(lines))
	for i, line := range lines {
		area[i] = make([]int, len(line))
		for j, r := range line {
			n, err := strconv.Atoi(string(r))
			if err != nil {
				log.Fatal(err)
			}
			area[i][j] = n
		}
	}
	return area
}

type cruicibleDirection int

const (
	up cruicibleDirection = iota
	right
	down
	left
)

func nextPosStraight(c coordinates.Coord, d cruicibleDirection) coordinates.Coord {
	var cc coordinates.Coord
	switch d {
	case up:
		cc = coordinates.Coord{X: c.X, Y: c.Y - 1}
	case right:
		cc = coordinates.Coord{X: c.X + 1, Y: c.Y}
	case down:
		cc = coordinates.Coord{X: c.X, Y: c.Y + 1}
	case left:
		cc = coordinates.Coord{X: c.X - 1, Y: c.Y}
	}
	return cc
}

func nextPosTurn(c coordinates.Coord, d cruicibleDirection) (coordinates.Coord, cruicibleDirection, coordinates.Coord, cruicibleDirection) {
	var a coordinates.Coord
	var b coordinates.Coord
	var d1, d2 cruicibleDirection
	switch d {
	case up, down:
		a = coordinates.Coord{X: c.X - 1, Y: c.Y}
		d1 = left
		b = coordinates.Coord{X: c.X + 1, Y: c.Y}
		d2 = right
	case right, left:
		a = coordinates.Coord{X: c.X, Y: c.Y - 1}
		d1 = up
		b = coordinates.Coord{X: c.X, Y: c.Y + 1}
		d2 = down
	}
	return a, d1, b, d2
}

// An Item is something we manage in a priority queue.
type Item struct {
	pos                  coordinates.Coord
	priority             int // The priority of the item in the queue.
	direction            cruicibleDirection
	stepsInSameDirection int
	heatLoss             int
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

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
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func part1(area [][]int) int {
	return solve(area, 0, 3)
}

func part2(area [][]int) int {
	return solve(area, 4, 10)
}

func solve(area [][]int, minStraightSteps, maxStraightSteps int) int {
	goal := coordinates.Coord{X: len(area[len(area)-1]) - 1, Y: len(area) - 1}
	pq := make(PriorityQueue, 2)
	pq[0] = &Item{
		pos:                  coordinates.Coord{X: 1, Y: 0},
		heatLoss:             area[0][1],
		direction:            right,
		stepsInSameDirection: 1,
		index:                0,
	}
	pq[1] = &Item{
		pos:                  coordinates.Coord{X: 0, Y: 1},
		heatLoss:             area[1][0],
		direction:            down,
		stepsInSameDirection: 1,
		index:                1,
	}
	heap.Init(&pq)

	visited := map[string]int{
		fmt.Sprintf("%d %d %d %d", 1, 0, right, 1): area[0][1],
		fmt.Sprintf("%d %d %d %d", 0, 1, down, 1):  area[1][0],
	}

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		if item.pos == goal && item.stepsInSameDirection >= minStraightSteps {
			return item.heatLoss
		}

		if item.stepsInSameDirection < maxStraightSteps {
			p := nextPosStraight(item.pos, item.direction)
			if p.X >= 0 && p.Y >= 0 && p.Y < len(area) && p.X < len(area[p.Y]) {
				heatLoss := item.heatLoss + area[p.Y][p.X]
				it := &Item{
					pos:                  p,
					priority:             heatLoss + coordinates.ManhattanDistance(p, goal),
					direction:            item.direction,
					stepsInSameDirection: item.stepsInSameDirection + 1,
					heatLoss:             heatLoss,
				}
				key := fmt.Sprintf("%d %d %d %d", it.pos.X, it.pos.Y, it.direction, it.stepsInSameDirection)
				if v, ok := visited[key]; !ok || it.heatLoss < v {
					heap.Push(&pq, it)
					visited[key] = it.heatLoss
				}
			}
		}

		if item.stepsInSameDirection >= minStraightSteps {
			a, d1, b, d2 := nextPosTurn(item.pos, item.direction)
			if a.X >= 0 && a.Y >= 0 && a.Y < len(area) && a.X < len(area[a.Y]) {
				heatLoss := item.heatLoss + area[a.Y][a.X]
				it := &Item{
					pos:                  a,
					priority:             heatLoss + coordinates.ManhattanDistance(a, goal),
					direction:            d1,
					stepsInSameDirection: 1,
					heatLoss:             heatLoss,
				}
				key := fmt.Sprintf("%d %d %d %d", it.pos.X, it.pos.Y, it.direction, it.stepsInSameDirection)
				if v, ok := visited[key]; !ok || it.heatLoss < v {
					heap.Push(&pq, it)
					visited[key] = it.heatLoss
				}
			}
			if b.X >= 0 && b.Y >= 0 && b.Y < len(area) && b.X < len(area[b.Y]) {
				heatLoss := item.heatLoss + area[b.Y][b.X]
				it := &Item{
					pos:                  b,
					priority:             heatLoss + coordinates.ManhattanDistance(b, goal),
					direction:            d2,
					stepsInSameDirection: 1,
					heatLoss:             heatLoss,
				}
				key := fmt.Sprintf("%d %d %d %d", it.pos.X, it.pos.Y, it.direction, it.stepsInSameDirection)
				if v, ok := visited[key]; !ok || it.heatLoss < v {
					heap.Push(&pq, it)
					visited[key] = it.heatLoss
				}
			}
		}
	}
	return 0
}

func Run(inputFile string) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	area := parseData(data)

	p1 := part1(area)
	fmt.Println("Part 1:", p1)

	p2 := part2(area)
	fmt.Println("Part 2:", p2)
}
