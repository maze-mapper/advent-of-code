// Advent of Code 2021 - Day 15
package day15

import (
	"container/heap"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"adventofcode/coordinates"
)

type Node struct {
	c        coordinates.Coord
	priority int // The priority of the item in the queue.
	index    int // The index of the item in the heap. The index is needed by update and is maintained by the heap.Interface methods.
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

func parseData(data []byte) [][]int {
	lines := strings.Split(
		strings.TrimSuffix(string(data), "\n"), "\n",
	)

	maze := make([][]int, len(lines))
	for i, line := range lines {
		maze[i] = make([]int, len(line))
		for j, r := range line {
			if val, err := strconv.Atoi(string(r)); err == nil {
				maze[i][j] = val
			} else {
				log.Fatal(err)
			}
		}
	}

	return maze
}

func getNeighbours(maze [][]int, c coordinates.Coord) []coordinates.Coord {
	neighbours := []coordinates.Coord{}
	if c.X > 0 {
		neighbours = append(neighbours, coordinates.Coord{X: c.X - 1, Y: c.Y})
	}
	if c.X < len(maze[c.Y])-1 {
		neighbours = append(neighbours, coordinates.Coord{X: c.X + 1, Y: c.Y})
	}
	if c.Y > 0 {
		neighbours = append(neighbours, coordinates.Coord{X: c.X, Y: c.Y - 1})
	}
	if c.Y < len(maze)-1 {
		neighbours = append(neighbours, coordinates.Coord{X: c.X, Y: c.Y + 1})
	}
	return neighbours
}

func traverse(maze [][]int) int {
	// Initialise priority queue
	pq := make(PriorityQueue, 1)
	start := coordinates.Coord{}
	end := coordinates.Coord{X: len(maze[len(maze)-1]) - 1, Y: len(maze) - 1}
	pq[0] = &Node{c: start}
	heap.Init(&pq)

	explored := map[coordinates.Coord]int{}

	for pq.Len() > 0 {
		// Take next item off priority queue
		node := heap.Pop(&pq).(*Node)

		if node.c == end {
			return node.priority
		}

		// Find all possible next moves and add any unexplored ones to the priority queue
		moves := getNeighbours(maze, node.c)
		for _, move := range moves {
			cost := maze[move.Y][move.X]
			newCost := node.priority + cost
			if oldCost, ok := explored[move]; !ok || newCost < oldCost {
				n := Node{
					c:        move,
					priority: newCost,
					index:    -1,
				}
				heap.Push(&pq, &n)
				explored[move] = newCost
			}
		}
	}
	log.Fatal("Did not find path")
	return 0
}

func extendMaze(maze [][]int, factor int) [][]int {
	newMaze := make([][]int, len(maze)*factor)
	for i := 0; i < len(maze); i++ {
		for f := 0; f < factor; f++ {
			newMaze[i+len(maze)*f] = make([]int, len(maze[i])*factor)
		}
	}

	for i := 0; i < len(maze); i++ {
		for j := 0; j < len(maze[i]); j++ {
			for fx := 0; fx < factor; fx++ {
				for fy := 0; fy < factor; fy++ {
					risk := maze[i][j] + fx + fy
					if risk > 9 {
						risk = risk % 9
					}
					newMaze[i+len(maze)*fy][j+len(maze[i])*fx] = risk
				}
			}
		}
	}

	return newMaze
}

func part1(maze [][]int) int {
	return traverse(maze)
}

func part2(maze [][]int) int {
	maze = extendMaze(maze, 5)
	return traverse(maze)
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	maze := parseData(data)

	p1 := part1(maze)
	fmt.Println("Part 1:", p1)

	p2 := part2(maze)
	fmt.Println("Part 2:", p2)
}
