// Advent of Code 2023 - Day 23.
package day23

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/maze-mapper/advent-of-code/coordinates"
)

func parseData(data []byte) (coordinates.Coord, coordinates.Coord, [][]rune) {
	var start, end coordinates.Coord
	lines := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
	area := make([][]rune, len(lines))
	for i, line := range lines {
		area[i] = make([]rune, len(line))
		for j, r := range line {
			area[i][j] = r
			if i == 0 && r == '.' {
				start.X = j
			}
			if i == len(area)-1 && r == '.' {
				end.X = j
				end.Y = len(area) - 1
			}
		}
	}
	return start, end, area
}

// graph is an adjacency list which maps nodes to adjacent nodes with the edge weight.
type graph map[coordinates.Coord]map[coordinates.Coord]int

func solve(start, end coordinates.Coord, area [][]rune, moveFunc func(coordinates.Coord, [][]rune) []coordinates.Coord) int {
	maxSteps := 0
	junctions := findNodes(start, end, area)
	g := graphFromNodes(junctions, area, moveFunc)
	dfs2(start, end, map[coordinates.Coord]bool{start: true}, 0, g, &maxSteps)
	return maxSteps
}

func dfs2(current, end coordinates.Coord, visited map[coordinates.Coord]bool, pathLength int, g graph, maxPath *int) {
	if current == end {
		if pathLength > *maxPath {
			*maxPath = pathLength
		}
	}
	neighbours, _ := g[current]
	for neighbour, weight := range neighbours {
		if visited[neighbour] {
			continue
		}
		newVisited := map[coordinates.Coord]bool{}
		for k := range visited {
			newVisited[k] = true
		}
		newVisited[neighbour] = true
		dfs2(neighbour, end, newVisited, pathLength+weight, g, maxPath)
	}
}

func findNodes(start, end coordinates.Coord, area [][]rune) map[coordinates.Coord]bool {
	nodes := map[coordinates.Coord]bool{
		start: true,
		end:   true,
	}
	for y := 0; y < len(area); y++ {
		for x := 0; x < len(area[y]); x++ {
			if area[y][x] == '#' {
				continue
			}
			c := coordinates.Coord{X: x, Y: y}
			neighbours := possibleDryMoves(c, area)
			if len(neighbours) > 2 {
				nodes[c] = true
			}
		}
	}
	return nodes
}

func bfs(start, current coordinates.Coord, targets, visited map[coordinates.Coord]bool, area [][]rune, moveFunc func(coordinates.Coord, [][]rune) []coordinates.Coord, steps int, g graph) {
	if targets[current] {
		if _, ok := g[start]; !ok {
			g[start] = map[coordinates.Coord]int{}
		}
		g[start][current] = steps
		return
	}
	for _, neighbour := range moveFunc(current, area) {
		if visited[neighbour] {
			continue
		}
		newVisited := map[coordinates.Coord]bool{}
		for k := range visited {
			newVisited[k] = true
		}
		newVisited[neighbour] = true
		bfs(start, neighbour, targets, newVisited, area, moveFunc, steps+1, g)
	}
}

func graphFromNodes(nodes map[coordinates.Coord]bool, area [][]rune, moveFunc func(coordinates.Coord, [][]rune) []coordinates.Coord) graph {
	g := graph{}
	for node := range nodes {
		for _, neighbour := range moveFunc(node, area) {
			bfs(node, neighbour, nodes, map[coordinates.Coord]bool{node: true, neighbour: true}, area, moveFunc, 1, g)
		}
	}
	return g
}

func connectNodes(nodes map[coordinates.Coord]bool, area [][]rune, moveFunc func(coordinates.Coord, [][]rune) []coordinates.Coord) graph {
	g := graph{}
	for node := range nodes {
	loop:
		for _, neighbour := range moveFunc(node, area) {
			fmt.Println("node", node, "has neighbour", neighbour)
			n := neighbour
			visited := map[coordinates.Coord]bool{node: true, n: true}
			steps := 1
			for !nodes[n] {
				fmt.Println("###", n)
				neighbours := moveFunc(n, area)
				newN := n
				for _, nn := range neighbours {
					if !visited[nn] {
						newN = nn
						visited[nn] = true
						fmt.Println("next step is", n)
						break
					}
				}
				if n == newN {
					continue loop
				}
				steps += 1

			}

			if _, ok := g[node]; !ok {
				g[node] = map[coordinates.Coord]int{}
			}
			if _, ok := g[n]; !ok {
				g[n] = map[coordinates.Coord]int{}
			}
			g[node][n] = steps
			g[n][node] = steps
		}
	}
	return g
}

func areaToGraph(area [][]rune, moveFunc func(coordinates.Coord, [][]rune) []coordinates.Coord) graph {
	g := graph{}
	for y := 0; y < len(area); y++ {
		for x := 0; x < len(area[y]); x++ {
			c := coordinates.Coord{X: x, Y: y}
			if area[y][x] == '#' {
				continue
			}
			neighbours := moveFunc(c, area)
			for _, n := range neighbours {
				if _, ok := g[n]; !ok {
					g[n] = map[coordinates.Coord]int{}
				}
				g[c][n] = 1
				g[n][c] = 1
			}
		}
	}
	return g
}

// func areaToContractedGraph(start, end coordinates.Coord, area [][]rune, moveFunc func(coordinates.Coord, [][]rune) []coordinates.Coord) graph {
// 	g := graph{}
// 	a := start
// 	neighbours := moveFunc(a, area)
// 	if len(neighbours) >

// 	return g
// }

// func (g graph) contractEdges() graph {
// 	newG := graph{}
// 	for a, neighbours := range g {
// 		// Check for a single path
// 		if len(neighbours) == 1 {
// 			var keys []coordinates.Coord
// 			for k := range neighbours {
// 				keys = append(keys, k)
// 			}
// 			key := keys[0]
// 		}
// 	}
// 	return newG
// }

// type Item struct {
// 	c       coordinates.Coord
// 	steps   int
// 	visited map[coordinates.Coord]bool
// 	// The priority of the item in the queue.
// 	priority int
// 	// The index of the item in the heap.
// 	// It is maintained by the heap.Interface methods.
// 	index int
// }

// // A PriorityQueue implements heap.Interface and holds Items.
// type PriorityQueue []*Item

// func (pq PriorityQueue) Len() int {
// 	return len(pq)
// }

// func (pq PriorityQueue) Less(i, j int) bool {
// 	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
// 	return pq[i].priority > pq[j].priority
// }

// func (pq PriorityQueue) Swap(i, j int) {
// 	pq[i], pq[j] = pq[j], pq[i]
// 	pq[i].index = i
// 	pq[j].index = j
// }

// func (pq *PriorityQueue) Push(x any) {
// 	n := len(*pq)
// 	item := x.(*Item)
// 	item.index = n
// 	*pq = append(*pq, item)
// }

// func (pq *PriorityQueue) Pop() any {
// 	old := *pq
// 	n := len(old)
// 	item := old[n-1]
// 	old[n-1] = nil  // avoid memory leak
// 	item.index = -1 // for safety
// 	*pq = old[0 : n-1]
// 	return item
// }

func possibleIcyMoves(c coordinates.Coord, area [][]rune) []coordinates.Coord {
	var moves []coordinates.Coord
	// Up.
	if c.Y > 0 {
		p := area[c.Y-1][c.X]
		if p == '.' || p == '^' {
			moves = append(moves, coordinates.Coord{X: c.X, Y: c.Y - 1})
		}
	}
	// Down.
	if c.Y < len(area)-1 {
		p := area[c.Y+1][c.X]
		if p == '.' || p == 'v' {
			moves = append(moves, coordinates.Coord{X: c.X, Y: c.Y + 1})
		}
	}
	// Left.
	if c.X > 0 {
		p := area[c.Y][c.X-1]
		if p == '.' || p == '<' {
			moves = append(moves, coordinates.Coord{X: c.X - 1, Y: c.Y})
		}
	}
	// Right.
	if c.X < len(area[c.Y])-1 {
		p := area[c.Y][c.X+1]
		if p == '.' || p == '>' {
			moves = append(moves, coordinates.Coord{X: c.X + 1, Y: c.Y})
		}
	}
	return moves
}

func possibleDryMoves(c coordinates.Coord, area [][]rune) []coordinates.Coord {
	var moves []coordinates.Coord
	for _, move := range []coordinates.Coord{
		{X: c.X, Y: c.Y - 1},
		{X: c.X, Y: c.Y + 1},
		{X: c.X - 1, Y: c.Y},
		{X: c.X + 1, Y: c.Y},
	} {
		if move.Y < 0 || move.Y >= len(area) || move.X < 0 || move.X >= len(area[c.Y]) {
			continue
		}
		if area[move.Y][move.X] != '#' {
			moves = append(moves, move)
		}
	}
	return moves
}

func dfs(current, end coordinates.Coord, area [][]rune, visited map[coordinates.Coord]bool, moveFunc func(coordinates.Coord, [][]rune) []coordinates.Coord, maxSteps *int) {
	if current == end {
		// Path length minus starting position.
		l := len(visited) - 1
		if l > *maxSteps {
			fmt.Println("Reached end in new max:", l, "steps")
			*maxSteps = l
		}
	}
	for _, neighbour := range moveFunc(current, area) {
		if visited[neighbour] {
			continue
		}
		newVisited := map[coordinates.Coord]bool{}
		for k := range visited {
			newVisited[k] = true
		}
		newVisited[neighbour] = true
		dfs(neighbour, end, area, newVisited, moveFunc, maxSteps)
	}

}

// func solve(start, end coordinates.Coord, area [][]rune, moveFunc func(coordinates.Coord, [][]rune) []coordinates.Coord) int {
// 	pq := make(PriorityQueue, 1)
// 	pq[0] = &Item{
// 		c:        start,
// 		steps:    0,
// 		visited:  map[coordinates.Coord]bool{start: true},
// 		priority: 0,
// 		index:    0,
// 	}
// 	heap.Init(&pq)

// 	maxSteps := 0
// 	for pq.Len() > 0 {
// 		item := heap.Pop(&pq).(*Item)
// 		if item.c == end {
// 			fmt.Println("Reached end in", item.steps, "steps")
// 		}
// 		if item.c == end && item.steps > maxSteps {
// 			maxSteps = item.steps
// 		}

// 		newSteps := item.steps + 1
// 		for _, move := range moveFunc(item.c, area) {
// 			if item.visited[move] {
// 				continue
// 			}
// 			newVisited := map[coordinates.Coord]bool{}
// 			for k := range item.visited {
// 				newVisited[k] = true
// 			}
// 			newVisited[move] = true
// 			it := &Item{
// 				c:        move,
// 				steps:    newSteps,
// 				visited:  newVisited,
// 				priority: newSteps + coordinates.ManhattanDistance(end, move),
// 			}
// 			heap.Push(&pq, it)
// 		}
// 	}
// 	return maxSteps
// }

func part1(start, end coordinates.Coord, area [][]rune) int {
	return solve(start, end, area, possibleIcyMoves)
	// maxSteps := 0
	// dfs(start, end, area, map[coordinates.Coord]bool{start: true}, possibleIcyMoves, &maxSteps)
	// return maxSteps
	// return solve(start, end, area, possibleIcyMoves)
}

func part2(start, end coordinates.Coord, area [][]rune) int {
	return solve(start, end, area, possibleDryMoves)
	// maxSteps := 0
	// dfs(start, end, area, map[coordinates.Coord]bool{start: true}, possibleDryMoves, &maxSteps)
	// return maxSteps
	// return solve(start, end, area, possibleDryMoves)
}

func Run(inputFile string) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	start, end, area := parseData(data)

	p1 := part1(start, end, area)
	fmt.Println("Part 1:", p1)

	p2 := part2(start, end, area)
	fmt.Println("Part 2:", p2)
}
