// Advent of Code 2019 - Day 13
package day13

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/maze-mapper/advent-of-code/2019/intcode"
)

// Coord is a Cartesian coordinate
type Coord struct {
	x, y int
}

// Tile IDs
const (
	tileEmpty  = 0
	tileWall   = 1
	tileBlock  = 2
	tilePaddle = 3
	tileBall   = 4
)

// Joystick positions
const (
	joystickNeutral = 0
	joystickLeft    = -1
	joystickRight   = 1
)

// getBallAndPaddleCoords returns the coordinates of the ball and paddle
func getBallAndPaddleCoords(tiles map[Coord]int) (Coord, Coord) {
	var ballCoord, paddleCoord Coord
	var foundBall, foundPaddle bool
	for c, id := range tiles {
		switch id {

		case tilePaddle:
			foundPaddle = true
			paddleCoord = c

		case tileBall:
			foundBall = true
			ballCoord = c

		}
	}
	if !foundPaddle {
		log.Fatalf("Did not find paddle position: paddle %v ball %v ", paddleCoord, ballCoord)
	}
	if !foundBall {
		log.Fatalf("Did not find ball position: paddle %v, ball %v", paddleCoord, ballCoord)
	}
	return ballCoord, paddleCoord
}

// chooseDirection returns the direction to move the paddle in to be as close as possible to the ball
func chooseDirection(tiles map[Coord]int) int {
	ballCoord, paddleCoord := getBallAndPaddleCoords(tiles)
	switch {

	case ballCoord.x > paddleCoord.x:
		return joystickRight

	case ballCoord.x < paddleCoord.x:
		return joystickLeft

	default:
		return joystickNeutral

	}
}

// runArcade runs the program and returns the tiles and score once finished
func runArcade(program []int) (map[Coord]int, int) {
	tiles := map[Coord]int{}

	computer := intcode.New(program)
	chanIn := make(chan int)
	chanOut := make(chan int)
	computer.SetChanIn(chanIn)
	computer.SetChanOut(chanOut)

	go computer.Run()

	var paddlePlaced, ballPlaced bool
	var x, y, score, outputCount int
	for output := range chanOut {
		// Outputs come in groups of three
		outputCount = outputCount % 3
		switch outputCount {

		case 0:
			x = output

		case 1:
			y = output

		case 2:
			if x == -1 && y == 0 {
				score = output
			} else {
				c := Coord{x: x, y: y}
				tiles[c] = output

				// If we are placing the ball then determine which direction we should move the paddle next and send to input
				// Only do this if the paddle has also been placed
				// Also do this if placing the paddle after the ball for the first time
				if (paddlePlaced && output == tileBall) || (!paddlePlaced && ballPlaced && output == tilePaddle) {
					direction := chooseDirection(tiles)
					go func() {
						chanIn <- direction
					}()
				}

				// Paddle has been placed for the first time
				if !paddlePlaced && output == tilePaddle {
					paddlePlaced = true
				}

				// Ball has been placed for the first time
				if !ballPlaced && output == tileBall {
					ballPlaced = true
				}
			}

		}
		outputCount += 1
	}

	return tiles, score
}

func part1(program []int) int {
	tiles, _ := runArcade(program)
	blockCount := 0
	for _, v := range tiles {
		if v == tileBlock {
			blockCount += 1
		}
	}
	return blockCount
}

func part2(program []int) int {
	program[0] = 2
	_, score := runArcade(program)
	return score
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	program := intcode.ReadProgram(data)

	p1 := part1(program)
	fmt.Println("Part 1:", p1)

	p2 := part2(program)
	fmt.Println("Part 2:", p2)
}
