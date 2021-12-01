// Advent of Code 2018 - Day 22
package day22

import (
	"container/heap"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// parseData reads the input text file and returns the cave depth and target coordinates
func parseData(file string) *Cave {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(
		strings.TrimSuffix(string(data), "\n"), "\n",
	)

	// Extract depth
	depth, err := strconv.Atoi(
		strings.TrimPrefix(lines[0], "depth: "),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Extract target
	targetParts := strings.Split(
		strings.TrimPrefix(lines[1], "target: "), ",",
	)
	x, errX := strconv.Atoi(targetParts[0])
	y, errY := strconv.Atoi(targetParts[1])
	if errX != nil || errY != nil {
		log.Fatal(errX, errY)
	}

	cave := Cave{
		depth:             depth,
		target:            coord{x, y},
		erosionLevelCache: map[coord]int{},
	}
	return &cave
}

// Define constants to represent regions of the cave
const (
	rocky  = '.'
	wet    = '='
	narrow = '|'
	mouth  = 'M'
	target = 'T'
)

// Define enums for tools
const (
	noTools = iota
	torch
	climbingGear

	allTools // Total number of tool options
)

// printRoute prints the route chosen through the cave
func (cave *Cave) printRoute(path []nodeState) {
	ANSIGreen := "\033[30m\033[42m"
	ANSIYellow := "\033[30m\033[43m"
	ANSIBlue := "\033[30m\033[44m"
	ANSIClear := "\033[0m"

	maxX := 0
	maxY := 0
	route := map[coord]int{}
	for _, v := range path {
		if v.c[0] > maxX {
			maxX = v.c[0]
		}
		if v.c[1] > maxY {
			maxY = v.c[1]
		}
		route[v.c] = v.tool
	}

	for i := 0; i <= maxY; i++ {
		for j := 0; j <= maxX; j++ {
			c := coord{j, i}
			if t, ok := route[c]; ok {
				switch t {
				case noTools:
					fmt.Print(ANSIGreen)
				case torch:
					fmt.Print(ANSIYellow)
				case climbingGear:
					fmt.Print(ANSIBlue)
				}
				fmt.Print(string(cave.regionType(c)))
				fmt.Print(ANSIClear)
			} else {
				fmt.Print(string(cave.regionType(c)))
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

// coord holds a coordinate (x, y)
type coord [2]int

// Cave holds basic information about the cave
type Cave struct {
	depth             int
	target            coord
	erosionLevelCache map[coord]int
}

// geologicIndex returns the geologic index for a coordinate
func (cave *Cave) geologicIndex(c coord) int {
	switch {

	// The region at (0, 0) (the mouth of the cave) has a geologic index of 0.
	// The region at the coordinates of the target has a geologic index of 0.
	case c == coord{0, 0} || c == cave.target:
		return 0

	// If the region's Y coordinate is 0, the geologic index is its X coordinate times 16807.
	case c[1] == 0:
		return c[0] * 16807

	// If the region's X coordinate is 0, the geologic index is its Y coordinate times 48271.
	case c[0] == 0:
		return c[1] * 48271

	// Otherwise, the region's geologic index is the result of multiplying the erosion levels of the regions at (X-1, Y) and (X, Y-1).
	default:
		el1 := cave.erosionLevel(coord{c[0] - 1, c[1]})
		el2 := cave.erosionLevel(coord{c[0], c[1] - 1})
		return el1 * el2

	}
}

// erosionLevel returns the erosion level for a coordinate
func (cave *Cave) erosionLevel(c coord) int {
	// Check if we have already calculated the erosion level
	if el, ok := cave.erosionLevelCache[c]; ok {
		return el
	}

	gi := cave.geologicIndex(c)
	el := (gi + cave.depth) % 20183
	cave.erosionLevelCache[c] = el
	return el
}

// regionType determines the region type at a coordinate
func (cave *Cave) regionType(c coord) rune {
	var rt rune
	el := cave.erosionLevel(c)
	r := el % 3
	switch r {
	case 0:
		rt = rocky
	case 1:
		rt = wet
	case 2:
		rt = narrow
	}

	return rt
}

// riskLevel returns the risk level of the regions between the cave mouth and target coordinates
func (cave *Cave) riskLevel() int {
	risk := 0
	for i := 0; i <= cave.target[1]; i++ {
		for j := 0; j <= cave.target[0]; j++ {
			r := cave.regionType(coord{j, i})
			switch r {
			case wet:
				risk += 1
			case narrow:
				risk += 2
			}
		}
	}
	return risk
}

// Node hold the state information for a region in the cave
type Node struct {
	c        coord
	tool     int // The current tool in use.
	time     int // The time taken to travel from the cave mouth to this coordinate.
	priority int // The priority of the item in the queue.
	index    int // The index of the item in the heap. The index is needed by update and is maintained by the heap.Interface methods.
}

type nodeState struct {
	c    coord
	tool int
}

// state returns a string representation of the node
func (node *Node) state() string {
	s := strconv.Itoa(node.c[0]) + " " + strconv.Itoa(node.c[1]) + " " + strconv.Itoa(node.tool)
	return s
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

// manhattanDistance returns the Manhattan distance from a coordinate to the target
func (cave *Cave) manhattanDistance(c coord) int {
	x := cave.target[0] - c[0]
	y := cave.target[1] - c[1]

	// Find absolute values
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}

	return x + y
}

// getNextMoves returns the available moves at the current position
func (cave *Cave) getNextMoves(node *Node) []*Node {
	nextMoves := []*Node{}
	// Find spaces that can be moved to
	coords := []coord{}
	for y := node.c[1] - 1; y <= node.c[1]+1; y++ {
		if y >= 0 {
			coords = append(coords, coord{node.c[0], y})
		}
	}
	for x := node.c[0] - 1; x <= node.c[0]+1; x++ {
		if x >= 0 {
			coords = append(coords, coord{x, node.c[1]})
		}
	}

	for _, nextC := range coords {
		// Check if we can enter this terrain type
		rt := cave.regionType(nextC)
		switch rt {

		case rocky:
			if node.tool == noTools {
				continue
			}

		case wet:
			if node.tool == torch {
				continue
			}

		case narrow:
			if node.tool == climbingGear {
				continue
			}
		}

		// Append next valid node to list of possible moves
		nextNode := &Node{
			c:        nextC,
			tool:     node.tool,
			time:     node.time + 1,
			priority: node.time + 1 + cave.manhattanDistance(nextC),
			index:    -1,
		}
		nextMoves = append(nextMoves, nextNode)

	}

	// Change tool
	rt := cave.regionType(node.c)
	for t := 0; t < allTools; t++ {
		if t != node.tool {
			// Check if we equip tool in current terrain type
			switch rt {

			case rocky:
				if t == noTools {
					continue
				}

			case wet:
				if t == torch {
					continue
				}

			case narrow:
				if t == climbingGear {
					continue
				}
			}
			nextNode := &Node{
				c:        node.c,
				tool:     t,
				time:     node.time + 7,
				priority: node.priority + 7, // Manhattan distance remains unchanged
				index:    -1,
			}
			nextMoves = append(nextMoves, nextNode)
		}
	}

	return nextMoves
}

// traverse finds the shortest path from the cave mouth to the target using an A* approach
func (cave *Cave) traverse() {
	// Initialise priority queue
	pq := make(PriorityQueue, 1)
	mouthCoord := coord{0, 0}
	mouthNode := &Node{
		c:        mouthCoord,
		tool:     torch,
		time:     0,
		priority: cave.manhattanDistance(mouthCoord),
		index:    0,
	}
	pq[0] = mouthNode
	heap.Init(&pq)

	// Track each coordinate that has been visited and with what tool equipped
	explored := map[string]int{}
	explored[mouthNode.state()] = mouthNode.time

	path := map[nodeState]nodeState{
		nodeState{c: mouthCoord, tool: torch}: nodeState{c: mouthCoord, tool: torch},
	}

	for pq.Len() > 0 {
		// Take next item off priority queue
		node := heap.Pop(&pq).(*Node)

		// Check if we have reached the target and have the torch equipped
		if node.c == cave.target && node.tool == torch {
			p := nodeState{c: node.c, tool: node.tool}
			route := []nodeState{}
			for {
				route = append(route, p)
				q := path[p]
				if p == q {
					break
				}
				p = q
			}
			cave.printRoute(route)

			fmt.Println("Part 2: Time taken to reach target at", cave.target, "is", node.time, "minutes")
			break
		}

		// Find all possible next moves and add any unexplored ones to the priority queue
		moves := cave.getNextMoves(node)
		for _, move := range moves {
			st := move.state()
			if t, ok := explored[st]; !ok || move.time < t {
				heap.Push(&pq, move)
				explored[st] = move.time
				path[nodeState{c: move.c, tool: move.tool}] = nodeState{c: node.c, tool: node.tool}
			}
		}
	}
}

func Run(inputFile string) {
	cave := parseData(inputFile)
	risk := cave.riskLevel()
	fmt.Println("Part 1: The risk level is", risk)
	cave.traverse()
}
