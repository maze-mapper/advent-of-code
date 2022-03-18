// Advent of Code 2019 - Day 15
package day15

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/maze-mapper/advent-of-code/2019/intcode"
	"github.com/maze-mapper/advent-of-code/coordinates"
)

// Directions of movement
const (
	north = 1
	south = 2
	west  = 3
	east  = 4
)

// Repair droid status codes
const (
	hitWall       = 0
	movedStep     = 1
	reachedOxygen = 2
)

// Droid holds state related to the repair droid
type Droid struct {
	position coordinates.Coord
	explored map[coordinates.Coord]int
}

// NewDroid returns a droid positioned at the origin
func NewDroid() *Droid {
	var pos coordinates.Coord
	d := Droid{
		position: pos,
		explored: map[coordinates.Coord]int{pos: movedStep},
	}
	return &d
}

// getAvailableMoves returns the adjacent coordinates that have not already been explored
func (d *Droid) getAvailableMoves() map[int]coordinates.Coord {
	moves := map[int]coordinates.Coord{}
	directions := []int{north, south, west, east}
	for _, direction := range directions {
		neighbour := movePosition(d.position, direction)
		if _, ok := d.explored[neighbour]; !ok {
			moves[direction] = neighbour
		}
	}
	return moves
}

// reverseDirection returns the opposite direction
func reverseDirection(direction int) int {
	switch direction {
	case north:
		return south
	case south:
		return north
	case west:
		return east
	case east:
		return west
	}
	log.Fatalf("Unknown direction %d", direction)
	return 0
}

// Move sends a movement command to the droid
func (d *Droid) Move(direction int, expectedPosition coordinates.Coord, chanIn, chanOut *chan int) {
	*chanIn <- direction
	status := <-*chanOut

	// Mark status code of position
	d.explored[expectedPosition] = status

	// Update position if moved
	if status != hitWall {
		d.position = expectedPosition
	}
}

// DFS performs a depth first search
func (d *Droid) DFS(chanIn, chanOut *chan int) {
	// Add location to visited
	moves := d.getAvailableMoves()
	for direction, newPosition := range moves {
		oldPosition := d.position

		d.Move(direction, newPosition, chanIn, chanOut)
		d.DFS(chanIn, chanOut)

		// Backtrack if the droid has moved
		if oldPosition != d.position {
			reverse := reverseDirection(direction)
			d.Move(reverse, oldPosition, chanIn, chanOut)
		}
	}
}

// PrintExplored prints the region explored by the repair droid
func (d *Droid) PrintExplored() {
	coords := make([]coordinates.Coord, len(d.explored))
	i := 0
	for k := range d.explored {
		coords[i] = k
		i += 1
	}
	minCoord, maxCoord := coordinates.Range(coords)

	for i := maxCoord.Y; i >= minCoord.Y; i-- {
		for j := minCoord.X; j <= maxCoord.X; j++ {
			c := coordinates.Coord{X: j, Y: i}
			if c == d.position {
				fmt.Print("D")
			} else {
				if value, ok := d.explored[c]; ok {
					switch value {
					case hitWall:
						fmt.Print("#")
					case movedStep:
						fmt.Print(".")
					case reachedOxygen:
						fmt.Print("O")
					}
				} else {
					fmt.Print(" ")
				}
			}
		}
		fmt.Print("\n")
	}
}

// movePosition returns the new coordinate if moving in the given direction
func movePosition(c coordinates.Coord, direction int) coordinates.Coord {
	switch direction {
	case north:
		c.Y += 1
	case south:
		c.Y -= 1
	case west:
		c.X -= 1
	case east:
		c.X += 1
	}
	return c
}

// exploreArea runs the repair droid program to fully explore an area
func exploreArea(program []int) map[coordinates.Coord]int {
	computer := intcode.New(program)
	chanIn := make(chan int)
	chanOut := make(chan int)
	computer.SetChanIn(chanIn)
	computer.SetChanOut(chanOut)

	// Bad practice: this will not terminate
	go computer.Run()

	droid := NewDroid()
	droid.DFS(&chanIn, &chanOut)

	//	droid.PrintExplored()
	return droid.explored
}

// getNeighbours returns the adjacent coordinates
func getNeighbours(c coordinates.Coord) []coordinates.Coord {
	directions := []int{north, south, west, east}
	neighbours := make([]coordinates.Coord, len(directions))
	for i, direction := range directions {
		neighbours[i] = movePosition(c, direction)
	}
	return neighbours
}

// Node holds a coordinate and its distance from the start
type Node struct {
	c          coordinates.Coord
	pathLength int
}

// BFS performs a Breadth First Search on a maze
func BFS(start coordinates.Coord, maze map[coordinates.Coord]int) map[coordinates.Coord]int {
	// Mark start position as visited with zero path length
	visited := map[coordinates.Coord]int{start: 0}

	// Initialise queue with start position
	queue := list.New()
	queue.PushBack(Node{c: start, pathLength: 0})

	for queue.Len() > 0 {
		elem := queue.Front()

		n := elem.Value.(Node)
		for _, neighbour := range getNeighbours(n.c) {
			// Check neighbour is traversable
			if val := maze[neighbour]; val != hitWall {
				// Check neighbour hasn't already been visited
				if _, ok := visited[neighbour]; !ok {
					visited[neighbour] = n.pathLength + 1
					queue.PushBack(Node{c: neighbour, pathLength: n.pathLength + 1})
				}
			}
		}
		queue.Remove(elem)
	}
	return visited
}

// findOxygenCoordinates returns the coordinates of the oxygen in the area
func findOxygenCoordinates(area map[coordinates.Coord]int) coordinates.Coord {
	var oxygenCoordinates coordinates.Coord
	var foundOxygen bool
	for k, v := range area {
		if v == reachedOxygen {
			oxygenCoordinates = k
			foundOxygen = true
			break
		}
	}
	if !foundOxygen {
		log.Fatal("Did not find oxgen source in area")
	}
	return oxygenCoordinates
}

// part1 returns the distance between the oxygen and origin
func part1(distances map[coordinates.Coord]int) int {
	return distances[coordinates.Coord{}]
}

// part2 returns the distance between the oxygen and furthest point in the area
func part2(distances map[coordinates.Coord]int) int {
	var maxDistance int
	for _, v := range distances {
		if v > maxDistance {
			maxDistance = v
		}
	}
	return maxDistance
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	program := intcode.ReadProgram(data)

	area := exploreArea(program)
	oxygen := findOxygenCoordinates(area)
	distances := BFS(oxygen, area)

	p1 := part1(distances)
	fmt.Println("Part 1:", p1)

	p2 := part2(distances)
	fmt.Println("Part 2:", p2)
}
