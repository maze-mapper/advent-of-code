// Advent of Code 2022 - Day 24.
package day24

import (
	"container/heap"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/maze-mapper/advent-of-code/coordinates"
)

func parseInput(data []byte) (coordinates.Coord, coordinates.Coord, map[coordinates.Coord][]rune, map[coordinates.Coord]struct{}, int, int) {
	lines := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
	maxY := len(lines) - 1
	maxX := len(lines[0]) - 1

	var start, end coordinates.Coord
	blizzards := map[coordinates.Coord][]rune{}
	walls := map[coordinates.Coord]struct{}{}

	for y, line := range lines {
		for x, r := range line {
			c := coordinates.Coord{X: x, Y: y}
			switch r {
			case '^', 'v', '<', '>':
				blizzards[c] = append(blizzards[c], r)
			case '#':
				walls[c] = struct{}{}
			case '.':
				if y == 0 {
					start = c
				} else if y == maxY {
					end = c
				}
			}
		}
	}

	return start, end, blizzards, walls, maxX, maxY
}

func (node *Node) state(maxX, maxY int, walls map[coordinates.Coord]struct{}) string {
	var b strings.Builder
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			c := coordinates.Coord{X: x, Y: y}
			if _, ok := walls[c]; ok {
				b.WriteString("#")
			} else if bl, ok := node.blizzards[c]; ok {
				if len(bl) > 1 {
					l := strconv.Itoa(len(bl))
					b.WriteString(l)
				} else if len(bl) > 1 {
					b.WriteString("2")
				} else {
					b.WriteRune(bl[0])
				}
			} else if c == node.elves {
				b.WriteString("E")
			} else {
				b.WriteString(".")
			}
		}
		b.WriteString("\n")
	}
	return b.String()
}

type Node struct {
	minute    int
	elves     coordinates.Coord
	blizzards map[coordinates.Coord][]rune
	priority  int
	index     int
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

func moveBlizzards(blizzards map[coordinates.Coord][]rune, maxX, maxY int) map[coordinates.Coord][]rune {
	newBlizzards := map[coordinates.Coord][]rune{}

	for c, directions := range blizzards {
		for _, d := range directions {
			var newC coordinates.Coord
			switch d {
			case '^':
				newC = coordinates.Coord{X: c.X, Y: c.Y - 1}
				if newC.Y == 0 {
					newC.Y = maxY - 1
				}
			case 'v':
				newC = coordinates.Coord{X: c.X, Y: c.Y + 1}
				if newC.Y == maxY {
					newC.Y = 1
				}
			case '<':
				newC = coordinates.Coord{X: c.X - 1, Y: c.Y}
				if newC.X == 0 {
					newC.X = maxX - 1
				}
			case '>':
				newC = coordinates.Coord{X: c.X + 1, Y: c.Y}
				if newC.X == maxX {
					newC.X = 1
				}
			}
			newBlizzards[newC] = append(newBlizzards[newC], d)
		}
	}
	return newBlizzards
}

func elfMoves(elves coordinates.Coord, blizzards map[coordinates.Coord][]rune, walls map[coordinates.Coord]struct{}) []coordinates.Coord {
	moves := []coordinates.Coord{}
	for x := -1; x <= 1; x++ {
	loop:
		for y := -1; y <= 1; y++ {
			// Only move in cardinal directions.
			if x != 0 && y != 0 {
				continue
			}
			c := coordinates.Coord{X: elves.X + x, Y: elves.Y + y}
			// Prevent moving in to wall or above the start.
			if _, ok := walls[c]; ok || c.Y < 0 {
				continue
			}
			// Prevent moving in to a blizzard.
			for b := range blizzards {
				if c == b {
					continue loop
				}
			}
			moves = append(moves, c)
		}
	}
	return moves
}

func aStar(start, end coordinates.Coord, blizzards map[coordinates.Coord][]rune, walls map[coordinates.Coord]struct{}, maxX, maxY int) *Node {
	pq := make(PriorityQueue, 1)
	startNode := &Node{
		minute:    0,
		elves:     start,
		blizzards: blizzards,
	}
	pq[0] = startNode
	heap.Init(&pq)

	visited := map[string]int{
		startNode.state(maxX, maxY, walls): 0,
	}

	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*Node)

		if node.elves == end {
			return node
		}

		newBlizzards := moveBlizzards(node.blizzards, maxX, maxY)
		moves := elfMoves(node.elves, newBlizzards, walls)
		newMinute := node.minute + 1

		for _, move := range moves {
			newNode := &Node{
				minute:    newMinute,
				elves:     move,
				blizzards: newBlizzards,
				priority:  newMinute + coordinates.ManhattanDistance(move, end),
			}
			s := newNode.state(maxX, maxY, walls)
			if m, ok := visited[s]; !ok || newMinute < m {
				visited[s] = newMinute
				heap.Push(&pq, newNode)
			}
		}
	}

	return &Node{}
}

func part1(start, end coordinates.Coord, blizzards map[coordinates.Coord][]rune, walls map[coordinates.Coord]struct{}, maxX, maxY int) int {
	node := aStar(start, end, blizzards, walls, maxX, maxY)
	return node.minute
}

func part2(start, end coordinates.Coord, blizzards map[coordinates.Coord][]rune, walls map[coordinates.Coord]struct{}, maxX, maxY int) int {
	minutes := 0
	node := aStar(start, end, blizzards, walls, maxX, maxY)
	minutes += node.minute
	node = aStar(end, start, node.blizzards, walls, maxX, maxY)
	minutes += node.minute
	node = aStar(start, end, node.blizzards, walls, maxX, maxY)
	minutes += node.minute
	return minutes
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	start, end, blizzards, walls, maxX, maxY := parseInput(data)

	p1 := part1(start, end, blizzards, walls, maxX, maxY)
	fmt.Println("Part 1:", p1)

	p2 := part2(start, end, blizzards, walls, maxX, maxY)
	fmt.Println("Part 2:", p2)
}
