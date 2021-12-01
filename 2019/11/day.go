// Advent of Code 2019 - Day 11
package day11

import (
	"fmt"
	"io/ioutil"
	"log"

	"adventofcode/2019/intcode"
)

// Coord is a Cartesian coordinate
type Coord struct {
	x, y int
}

// Directions for robot
const (
	up = iota
	right
	down
	left
)

// robot holds the current position and orientation of the robot
type robot struct {
	position    Coord
	orientation int
}

// TurnLeft rotates the robot 90 degrees left
func (r *robot) TurnLeft() {
	if r.orientation == up {
		r.orientation = left
	} else {
		r.orientation -= 1
	}
}

// TurnRight rotates the robot 90 degrees right
func (r *robot) TurnRight() {
	if r.orientation == left {
		r.orientation = up
	} else {
		r.orientation += 1
	}
}

// Move moves the robot one unit in the direction it is facing
func (r *robot) Move() {
	switch r.orientation {
	case up:
		r.position.y += 1
	case right:
		r.position.x += 1
	case down:
		r.position.y -= 1
	case left:
		r.position.x -= 1
	}
}

// Colours of hull
const (
	black = 0
	white = 1
)

// runRobot runs the Inctocde program on the robot given a starting colour
func runRobot(program []int, startColour int) map[Coord]int {
	r := robot{}
	painted := map[Coord]int{}

	computer := intcode.New(program)
	chanIn := make(chan int, 1) // Buffered so that we don't hang when sending final unused input
	chanOut := make(chan int)
	computer.SetChanIn(chanIn)
	computer.SetChanOut(chanOut)

	go computer.Run()
	chanIn <- startColour

	outputCount := 0
	for output := range chanOut {
		if outputCount%2 == 0 {
			// First output of pair - paint hull
			painted[r.position] = output

		} else {
			// Second output of pair - turn robot
			switch output {
			case 0:
				r.TurnLeft()
			case 1:
				r.TurnRight()
			default:
				log.Fatalf("Unrecognised output %d", output)
			}
			r.Move()
			chanIn <- painted[r.position]
		}
		outputCount += 1
	}

	return painted
}

// findCoordinateRange returns the minimum and maximum values of x and y for a set of coordinates
func findCoordinateRange(painted map[Coord]int) (int, int, int, int) {
	// Robot starts at the origin
	var minX, minY, maxX, maxY int
	for c := range painted {
		if c.x < minX {
			minX = c.x
		}
		if c.y < minY {
			minY = c.y
		}
		if c.x > maxX {
			maxX = c.x
		}
		if c.y > maxY {
			maxY = c.y
		}
	}
	return minX, minY, maxX, maxY
}

// printIdentifier prints the identifier painted by the robot
func printIdentifier(painted map[Coord]int) {
	minX, minY, maxX, maxY := findCoordinateRange(painted)

	for i := maxY; i >= minY; i-- {
		for j := minX; j <= maxX; j++ {
			c := Coord{x: j, y: i}
			colour := painted[c]
			switch colour {
			case black:
				fmt.Print(".")
			case white:
				fmt.Print("#")
			}
		}
		fmt.Print("\n")
	}
}

func part1(program []int) int {
	painted := runRobot(program, black)
	return len(painted)
}

func part2(program []int) {
	painted := runRobot(program, white)
	printIdentifier(painted)
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	program := intcode.ReadProgram(data)

	p1 := part1(program)
	fmt.Println("Part 1:", p1)

	fmt.Println("Part 2:")
	part2(program)
}
