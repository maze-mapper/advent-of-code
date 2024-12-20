// Advent of Code 2024 - Day 20.
package day20

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
	region := make([][]bool, len(lines))
	var start, end coordinates.Coord
	for i, line := range lines {
		region[i] = make([]bool, len(line))
		for j, r := range line {
			switch r {
			case 'S':
				start = coordinates.Coord{X: j, Y: i}
			case 'E':
				end = coordinates.Coord{X: j, Y: i}
			case '#':
				region[i][j] = true
			}
		}
	}
	return region, start, end
}

func aStarWithCheats(region [][]bool, start, end coordinates.Coord, cheatMaxSteps, stepsToSave int) int {
	// Find the distances of every open space to the end.
	nonCheatingStepsToEnd := bfs(region, end)
	nonCheatingMinSteps := nonCheatingStepsToEnd[start]

	pq := make(PriorityQueue, 1)
	pq[0] = &Item{
		state: gameState{
			position: start,
			steps:    0,
		},
		priority: 0,
	}
	heap.Init(&pq)

	var cheatsToReachEnd int
	visited := map[coordinates.Coord]int{}
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)

		if item.state.position == end {
			// A-star search ensures that no shorter paths will be found after this.
			break
		}

		if s, ok := visited[item.state.position]; ok && s <= item.state.steps {
			continue
		}
		visited[item.state.position] = item.state.steps

		// Do a cheat move. This uses the data from the BFS to short circuit the remaining A-star after the cheat.
		for x := -cheatMaxSteps; x <= cheatMaxSteps; x++ {
			for y := -cheatMaxSteps; y <= cheatMaxSteps; y++ {
				absX := x
				if absX < 0 {
					absX *= -1
				}
				absY := y
				if absY < 0 {
					absY *= -1
				}
				totalCheatSteps := absX + absY
				if totalCheatSteps > cheatMaxSteps {
					continue
				}
				newPosition := coordinates.Coord{X: item.state.position.X + x, Y: item.state.position.Y + y}
				distToEnd, ok := nonCheatingStepsToEnd[newPosition]
				if !ok {
					// Not a valid position.
					continue
				}
				stepsToEnd := item.state.steps + totalCheatSteps + distToEnd
				if nonCheatingMinSteps-stepsToEnd >= stepsToSave {
					cheatsToReachEnd++
				}

			}
		}

		// Non-cheating moves.
		for _, state := range []gameState{
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
		} {
			// Check bounds.
			if state.position.Y < 0 || state.position.Y >= len(region) || state.position.X < 0 || state.position.X >= len(region[state.position.Y]) {
				continue
			}
			// Check that the move does not hit a wall.
			if region[state.position.Y][state.position.X] {
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

	return cheatsToReachEnd
}

// bfs does a Bredth First Search of the region to find the distance of every
// open space to the given point without any cheating.
func bfs(region [][]bool, point coordinates.Coord) map[coordinates.Coord]int {
	queue := []coordinates.Coord{point}
	distances := map[coordinates.Coord]int{
		point: 0,
	}
	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]
		for _, n := range []coordinates.Coord{
			{X: c.X - 1, Y: c.Y},
			{X: c.X + 1, Y: c.Y},
			{X: c.X, Y: c.Y - 1},
			{X: c.X, Y: c.Y + 1},
		} {
			if region[n.Y][n.X] {
				continue
			}
			if _, ok := distances[n]; !ok {
				distances[n] = distances[c] + 1
				queue = append(queue, n)
			}
		}
	}
	return distances
}

func part1(region [][]bool, start, end coordinates.Coord, stepsToSave int) int {
	return aStarWithCheats(region, start, end, 2, stepsToSave)
}

func part2(region [][]bool, start, end coordinates.Coord, stepsToSave int) int {
	return aStarWithCheats(region, start, end, 20, stepsToSave)
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
	region, start, end := parseData(data)

	stepsToSave := 100

	p1 := part1(region, start, end, stepsToSave)
	fmt.Println("Part 1:", p1)

	p2 := part2(region, start, end, stepsToSave)
	fmt.Println("Part 2:", p2)
}
