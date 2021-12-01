// Advent of Code 2018 - Day 23
package day23

import (
	"container/heap"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type coord struct {
	x, y, z int
}

type nanobot struct {
	pos coord
	r   int
}

// manhattanDistance returns the Manhattan distance between two coordinates
func manhattanDistance(a, b coord) int {
	x := a.x - b.x
	if x < 0 {
		x = -x
	}

	y := a.y - b.y
	if y < 0 {
		y = -y
	}

	z := a.z - b.z
	if z < 0 {
		z = -z
	}

	return x + y + z
}

// isInRange determines if a coordinate is in range of a nanobot
func (bot *nanobot) isInRange(c coord) bool {
	md := manhattanDistance(bot.pos, c)
	if md <= bot.r {
		return true
	}
	return false
}

// parseData reads the input text file and returns the nanobots
func parseData(file string) []nanobot {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(
		strings.TrimSuffix(string(data), "\n"), "\n",
	)

	bots := make([]nanobot, len(lines))
	for i, line := range lines {
		var x, y, z, r int
		fmt.Sscanf(line, "pos=<%d,%d,%d>, r=%d", &x, &y, &z, &r)
		bots[i] = nanobot{
			pos: coord{x, y, z},
			r:   r,
		}
	}

	return bots
}

// findStrongestNanobot finds the nanobot with the highest signal radius
func findStrongestNanobot(bots []nanobot) *nanobot {
	maxR := 0
	var strongest *nanobot
	for i := 0; i < len(bots); i++ {
		if bots[i].r > maxR {
			maxR = bots[i].r
			strongest = &bots[i]
		}
	}
	return strongest
}

func part1(bots []nanobot) {
	strongest := findStrongestNanobot(bots)
	inRange := 0
	for _, bot := range bots {
		if strongest.isInRange(bot.pos) {
			inRange += 1
		}
	}
	fmt.Println("Part 1: Strongest nanobot has", inRange, "nanobots in range")
}

type searchCube struct {
	pos      coord // Coordinate of the vertex with the lowest values
	size     int   // Size of the seach cube region, a power of 2
	bots     int   // Number of nanobots that are in range of this search cube
	distance int   // Distance to origin

	index int // The index of the item in the heap. The index is needed by update and is maintained by the heap.Interface methods.
}

// distToRange returns the distance of a point to a range of values
func distToRange(p, min, max int) int {
	switch {
	case p < min:
		return min - p
	case p > max:
		return p - max
	default:
		return 0
	}
}

// countNanobots returns the number of nanobots that are in range of the search cube
func (cube *searchCube) countNanobots(bots []nanobot) {
	count := 0
	for i := 0; i < len(bots); i++ {
		xDist := distToRange(bots[i].pos.x, cube.pos.x, cube.pos.x+cube.size-1)
		yDist := distToRange(bots[i].pos.y, cube.pos.y, cube.pos.y+cube.size-1)
		zDist := distToRange(bots[i].pos.z, cube.pos.z, cube.pos.z+cube.size-1)
		if xDist+yDist+zDist <= bots[i].r {
			count += 1
		}
	}
	cube.bots = count
}

// makeChildren returns the eight sub-cubes of the parent
func (cube *searchCube) makeChildren(bots []nanobot) []*searchCube {
	if cube.size == 1 {
		log.Fatal("Cannot create search cube children for parent of size 1")
	}
	newSize := cube.size / 2
	children := make([]*searchCube, 8)

	i := 0
	for x := cube.pos.x; x <= cube.pos.x+newSize; x += newSize {
		for y := cube.pos.y; y <= cube.pos.y+newSize; y += newSize {
			for z := cube.pos.z; z <= cube.pos.z+newSize; z += newSize {
				c := coord{x, y, z}
				sc := &searchCube{
					pos:      c,
					size:     newSize,
					distance: manhattanDistance(c, coord{0, 0, 0}),
					index:    -1,
				}
				sc.countNanobots(bots)
				children[i] = sc
				i += 1
			}
		}
	}

	return children
}

type priorityQueue []*searchCube

func (pq priorityQueue) Len() int {
	return len(pq)
}

func (pq priorityQueue) Less(i, j int) bool {
	// Prioritise by number of nanobots, distance to origin then size
	if pq[i].bots == pq[j].bots {
		if pq[i].distance == pq[j].distance {
			return pq[i].size < pq[j].size
		} else {
			return pq[i].distance < pq[j].distance
		}
	} else {
		return pq[i].bots > pq[j].bots
	}
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	sc := x.(*searchCube)
	sc.index = n
	*pq = append(*pq, sc)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	sc := old[n-1]
	old[n-1] = nil // avoid memory leak
	sc.index = -1  // for safety
	*pq = old[0 : n-1]
	return sc
}

func part2(bots []nanobot) {
	// Create a search cube with a large volume to contain all nanobots
	size := 1 << 62
	loc := -1 << 31
	pos := coord{loc, loc, loc}
	sc := &searchCube{
		pos:      pos,
		size:     size,
		distance: manhattanDistance(coord{0, 0, 0}, pos),
		index:    0,
	}
	sc.countNanobots(bots)
	if sc.bots != len(bots) {
		log.Fatal("Initial search cube does not hold all nanobots")
	}

	// Initialise priority queue
	pq := make(priorityQueue, 1)
	pq[0] = sc
	heap.Init(&pq)

	for pq.Len() > 0 {
		cube := heap.Pop(&pq).(*searchCube)
		if cube.size == 1 {
			fmt.Println("Part 2: Coordinate", cube.pos, "with Manhattan distance", cube.distance)
			break
		} else {
			children := cube.makeChildren(bots)
			for i := 0; i < len(children); i++ {
				if children[i].bots != 0 {
					heap.Push(&pq, children[i])
				}
			}
		}
	}
}

func Run(inputFile string) {
	bots := parseData(inputFile)
	part1(bots)
	part2(bots)
}
