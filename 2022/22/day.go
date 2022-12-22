// Advent of Code 2022 - Day 22.
package day22

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// Facing directions.
const (
	facingRight = 0
	facingDown  = 1
	facingLeft  = 2
	facingUp    = 3
	facingMax   = 4
)

// Turn directions.
const (
	turnLeft = iota
	turnRight
)

func parseInput(data []byte) ([][]string, []int) {
	parts := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n\n")

	lines := strings.Split(parts[0], "\n")
	board := make([][]string, len(lines))
	maxWidth := 0
	for _, line := range lines {
		w := len(line)
		if w > maxWidth {
			maxWidth = w
		}
	}
	for i, line := range lines {
		w := len(line)
		if w < maxWidth {
			suffix := strings.Repeat(" ", maxWidth-w)
			line += suffix
		}
		board[i] = strings.Split(line, "")
	}

	pathInstructions := strings.Split(strings.Replace(strings.Replace(parts[1], "L", " L ", -1), "R", " R ", -1), " ")
	path := make([]int, len(pathInstructions))
	for i, p := range pathInstructions {
		switch p {
		case "L":
			path[i] = turnLeft

		case "R":
			path[i] = turnRight

		default:
			v, err := strconv.Atoi(p)
			if err != nil {
				log.Fatal(err)
			}
			path[i] = v
		}
	}

	return board, path
}

func changeDirection(facing int, turn int) int {
	if turn == turnLeft {
		facing -= 1
		if facing < 0 {
			facing += facingMax
		}
	} else {
		facing += 1
		if facing >= facingMax {
			facing = facing % facingMax
		}
	}
	return facing
}

func moveDirection(facing int) (int, int) {
	xDir := 0
	yDir := 0
	switch facing {
	case facingRight:
		xDir = 1
	case facingDown:
		yDir = 1
	case facingLeft:
		xDir = -1
	case facingUp:
		yDir = -1
	}
	return xDir, yDir
}

func startingPosition(board [][]string) (int, int) {
	for x, b := range board[0] {
		if b == "." {
			return x, 0
		}
	}
	return 0, 0
}

func wrap(x, y, width, height int) (int, int) {
	if x < 0 {
		x += width
	}
	if x >= width {
		x -= width
	}
	if y < 0 {
		y += height
	}
	if y >= height {
		y -= height
	}
	return x, y
}

func customWrap(x, y, facing int) (int, int, int) {
	xDir, yDir := moveDirection(facing)
	x += xDir
	y += yDir

	xOffset := x % 50
	yOffset := y % 50

	switch {
	// blue
	case y == -1 && x >= 50 && x < 100 && facing == facingUp:
		facing = facingRight
		x = 0
		y = 150 + xOffset

		//black
	case y == -1 && x >= 100 && x < 150 && facing == facingUp:
		facing = facingUp
		y = 199
		x = 0 + xOffset

	//red
	case y == 50 && x >= 100 && x < 150 && facing == facingDown:
		facing = facingLeft
		x = 99
		y = 50 + xOffset

		// red
	case y == 99 && x >= 0 && x < 50 && facing == facingUp:
		facing = facingRight
		x = 50
		y = 50 + xOffset

		//red
	case y == 150 && x >= 50 && x < 100 && facing == facingDown:
		facing = facingLeft
		x = 49
		y = 150 + xOffset

		//black
	case y == 200 && x >= 0 && x < 50 && facing == facingDown:
		facing = facingDown
		y = 0
		x = 100 + xOffset

		// blue
	case x == -1 && y >= 150 && y < 200 && facing == facingLeft:
		facing = facingDown
		y = 0
		x = 50 + yOffset

		//red
	case x == 50 && y >= 150 && y < 200 && facing == facingRight:
		facing = facingUp
		y = 149
		x = 50 + yOffset

		//green dots
	case x == -1 && y >= 100 && y < 150 && facing == facingLeft:
		facing = facingRight
		x = 50
		y = 49 - yOffset

		//green
	case x == 100 && y >= 100 && y < 150 && facing == facingRight:
		facing = facingLeft
		x = 149
		y = 49 - yOffset

		//red
	case x == 49 && y >= 50 && y < 100 && facing == facingLeft:
		facing = facingDown
		y = 100
		x = 0 + yOffset

		//red
	case x == 100 && y >= 50 && y < 100 && facing == facingRight:
		facing = facingUp
		y = 49
		x = 100 + yOffset

	//green dots
	case x == 49 && y >= 0 && y < 50 && facing == facingLeft:
		facing = facingRight
		x = 0
		y = 149 - yOffset

		//green
	case x == 150 && y >= 0 && y < 50 && facing == facingRight:
		facing = facingLeft
		x = 99
		y = 149 - yOffset

	}

	return x, y, facing
}

func cubeWrap(board [][]string, x, y, facing, size int) (int, int, int) {
	xDir, yDir := moveDirection(facing)
	nextXPos := x + xDir
	nextYPos := y + yDir

	// Determine if the new position is on another cube face.
	currentXTile := x / size
	currentYTile := y / size
	nextXTile := nextXPos / size
	nextYTile := nextYPos / size

	if currentXTile == nextXTile && currentYTile == nextYTile {
		return nextXPos, nextYPos, facing
	}

	// Moving in the x direction.
	if nextXTile != currentXTile {
		fmt.Println("Change tile x")
		// TODO x > len

		// Next space is on an adjacent face.
		if board[nextYPos][nextXPos] != " " {
			return nextXPos, nextYPos, facing
		}

		yOffset := nextYPos % size

		// Check if there is a cube face downwards one tile.
		yDown := nextYPos - yOffset - 1
		xDown := x + (yOffset+1)*xDir
		if board[yDown][xDown] != " " {
			nextXPos = xDown
			nextYPos = yDown
			facing = changeDirection(facing, turnRight)
			return nextXPos, nextYPos, facing
		}

		// Check if there is a cube face upwards one tile.
		yUp := nextYPos + size - yOffset
		xUp := x + (size-yOffset)*xDir
		if board[yUp][xUp] != " " {
			nextXPos = xUp
			nextYPos = yUp
			facing = changeDirection(facing, turnLeft)
			return nextXPos, nextYPos, facing
		}
	}

	// Moving in the y direction.
	if nextYTile != currentYTile {
		fmt.Println("Change tile y")
		// TODO check bounds

		// Next space is on an adjacent face.
		if board[nextYPos][nextXPos] != " " {
			return nextXPos, nextYPos, facing
		}

		xOffset := nextXPos % size

		// Check if there is a cube face left one tile.
		yLeft := y + (xOffset+1)*yDir
		xLeft := nextXPos - xOffset - 1
		if board[yLeft][xLeft] != " " {
			nextXPos = xLeft
			nextYPos = yLeft
			facing = changeDirection(facing, turnLeft)
			return nextXPos, nextYPos, facing
		}

		// Check if there is a cube face right one tile.
		yRight := y + (size-xOffset)*yDir
		xRight := nextXPos - xOffset - 1
		if board[yRight][xRight] != " " {
			nextXPos = xRight
			nextYPos = yRight
			facing = changeDirection(facing, turnRight)
			return nextXPos, nextYPos, facing
		}
	}

	return nextXPos, nextYPos, facing
}

func score(row, col, facing int) int {
	return 1000*(row+1) + 4*(col+1) + facing
}

func part1(board [][]string, path []int) int {
	xPos, yPos := startingPosition(board)
	facing := facingRight

	// First instruction is a move.
	isMove := true

	for _, p := range path {
		if isMove {
			xDir, yDir := moveDirection(facing)
			for step := 0; step < p; step++ {
				nextXPos := xPos + xDir
				nextYPos := yPos + yDir
				width := len(board[yPos])
				height := len(board)
				nextXPos, nextYPos = wrap(nextXPos, nextYPos, width, height)

				for board[nextYPos][nextXPos] == " " {
					nextXPos += xDir
					nextYPos += yDir
					nextXPos, nextYPos = wrap(nextXPos, nextYPos, width, height)
				}
				if board[nextYPos][nextXPos] == "#" {
					break
				} else {
					xPos = nextXPos
					yPos = nextYPos
				}
			}
		} else {
			facing = changeDirection(facing, p)
		}
		isMove = !isMove
	}

	return score(yPos, xPos, facing)
}

func part2(board [][]string, path []int) int {
	//	cubeSize := 50
	xPos, yPos := startingPosition(board)
	facing := facingRight

	// First instruction is a move.
	isMove := true

	for _, p := range path {
		if isMove {
			fmt.Printf("Move %d, facing %d from (%d,%d)\n", p, facing, xPos, yPos)
			for step := 0; step < p; step++ {
				//nextXPos, nextYPos, nextFacing := cubeWrap(board, xPos, yPos, facing, 50)
				nextXPos, nextYPos, nextFacing := customWrap(xPos, yPos, facing)

				if board[nextYPos][nextXPos] == "#" {
					break
				} else if board[nextYPos][nextXPos] == " " {
					log.Fatal("not a valid space")
				} else {
					xPos = nextXPos
					yPos = nextYPos
					facing = nextFacing
				}
				fmt.Println(xPos, yPos, facing)
			}
		} else {
			facing = changeDirection(facing, p)
			fmt.Printf("Turn %d, now facing %d\n", p, facing)
		}
		isMove = !isMove
	}

	return score(yPos, xPos, facing)
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	board, path := parseInput(data)
	fmt.Println(path)
	for _, row := range board {
		fmt.Println(row)
	}

	p1 := part1(board, path)
	fmt.Println("Part 1:", p1)

	p2 := part2(board, path)
	fmt.Println("Part 2:", p2)
}
