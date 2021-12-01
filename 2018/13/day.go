// Advent of Code 2018 - Day 13
package day13

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

// Constants for cart direction, order is important
const (
	Up = iota
	Right
	Down
	Left
)

// Constants for cart turning
const (
	TurnLeft = iota
	Straight
	TurnRight
)

// coord is a coordiante (x, y)
type coord [2]int

// Cart holds information on a mine cart
type Cart struct {
	direction, nextTurn int
}

// doTurnLeft makes the cart turn 90 degrees left
func (cart *Cart) doTurnLeft() {
	cart.direction = (cart.direction + 3) % 4
}

// doTurnRight makes the cart turn 90 degrees right
func (cart *Cart) doTurnRight() {
	cart.direction = (cart.direction + 1) % 4
}

// doIntersectionTurn turns the cart at an intersection
func (cart *Cart) doIntersectionTurn() {
	switch cart.nextTurn {
	case TurnLeft:
		cart.doTurnLeft()
	case TurnRight:
		cart.doTurnRight()
	}
	cart.nextTurn = (cart.nextTurn + 1) % 3
}

// parseData reads the input text file and returns data structures
func parseData(file string) ([][]rune, map[coord]Cart) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(
		strings.TrimSuffix(string(data), "\n"), "\n",
	)
	track := make([][]rune, len(lines))
	carts := map[coord]Cart{}
	for i, line := range lines {
		track[i] = make([]rune, len(line))
		for j := 0; j < len(line); j++ {
			c := coord{j, i}
			// Carts are known to be on straight track initially
			switch line[j] {
			case '^':
				track[i][j] = '|'
				carts[c] = Cart{
					direction: Up,
					nextTurn:  TurnLeft,
				}
			case '>':
				track[i][j] = '-'
				carts[c] = Cart{
					direction: Right,
					nextTurn:  TurnLeft,
				}
			case 'v':
				track[i][j] = '|'
				carts[c] = Cart{
					direction: Down,
					nextTurn:  TurnLeft,
				}
			case '<':
				track[i][j] = '-'
				carts[c] = Cart{
					direction: Left,
					nextTurn:  TurnLeft,
				}
			default:
				track[i][j] = rune(line[j])
			}
		}
	}

	return track, carts
}

// getMoveOrder returns the coordinates of the carts in the order they will move
func getMoveOrder(carts map[coord]Cart) []coord {
	// Get the cart coordinates
	coords := make([]coord, len(carts))
	i := 0
	for c := range carts {
		coords[i] = c
		i++
	}
	// Sort the cart coordinates
	sort.Slice(
		coords,
		func(i, j int) bool {
			if coords[i][1] == coords[j][1] {
				// Sort by x if y is the same
				return coords[i][0] < coords[j][0]
			} else {
				// Sort by y
				return coords[i][1] < coords[j][1]
			}
		},
	)
	return coords
}

// getNextCoordinate returns the next location of the cart
func getNextCoordinate(c coord, direction int) coord {
	switch direction {
	case Up:
		c[1] -= 1
	case Right:
		c[0] += 1
	case Down:
		c[1] += 1
	case Left:
		c[0] -= 1
	}
	return c
}

// tick moves all carts one by one
func tick(track [][]rune, carts map[coord]Cart) []coord {
	coords := getMoveOrder(carts)
	collisions := []coord{}

	for _, c := range coords {
		cart, ok := carts[c]
		// Cart may have been removed if it was hit by another cart in the same tick
		if !ok {
			continue
		}

		delete(carts, c)

		nextCoord := getNextCoordinate(c, cart.direction)

		// Check for collision
		if _, ok := carts[nextCoord]; ok {
			// Collision!
			collisions = append(collisions, nextCoord)
			delete(carts, nextCoord)
			// TODO remove coord from list
		} else {
			// Handle next track segment
			switch track[nextCoord[1]][nextCoord[0]] {
			case '+':
				cart.doIntersectionTurn()
			case '/':
				switch cart.direction {
				case Up, Down:
					cart.doTurnRight()
				case Right, Left:
					cart.doTurnLeft()
				}
			case '\\':
				switch cart.direction {
				case Up, Down:
					cart.doTurnLeft()
				case Right, Left:
					cart.doTurnRight()
				}
			}
			carts[nextCoord] = cart
		}
	}
	return collisions
}

// printState prints the track with the current location of the carts
func printState(track [][]rune, carts map[coord]Cart) {
	for i, row := range track {
		for j, val := range row {
			c := coord{j, i}
			if cart, ok := carts[c]; ok {
				switch cart.direction {
				case Up:
					fmt.Print("^")
				case Right:
					fmt.Print(">")
				case Down:
					fmt.Print("v")
				case Left:
					fmt.Print("<")
				}
			} else {
				fmt.Print(string(val))
			}
		}
		fmt.Print("\n")
	}
}

func Run(inputFile string) {
	track, carts := parseData(inputFile)
	collisions := []coord{}

	// Part 1
	for len(collisions) == 0 {
		collisions = tick(track, carts)
	}
	fmt.Println("Part 1: Collision at", collisions[0])

	// Part 2
	// Exit loop if there are zero or one carts
	for len(carts) > 1 {
		collisions = tick(track, carts)
	}
	for k := range carts {
		fmt.Println("Part 2: Last cart at", k)
	}
}
