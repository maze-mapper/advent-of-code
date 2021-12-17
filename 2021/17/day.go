// Advent of Code 2021 - Day 17
package day17

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// parseData returns the x and y coordinates for the target area range
func parseData(data []byte) (int, int, int, int) {
	line := strings.TrimSuffix(string(data), "\n")
	line = strings.TrimPrefix(line, "target area: x=")
	line = strings.Replace(line, ", y=", " ", 1)
	line = strings.Replace(line, "..", " ", 2)
	var x1, x2, y1, y2 int
	fmt.Sscanf(line, "%d %d %d %d", &x1, &x2, &y1, &y2)
	return x1, x2, y1, y2
}

// triangularNumber returns the n-th triangular number
func triangularNumber(n int) int {
	return n * (n + 1) / 2
}

// getMinVelocity returns the smallest velocity required to reach a certain distance before the velocity reaches zero
func getMinVelocity(target int) int {
	n := 1
	for {
		if tn := triangularNumber(n); tn >= target {
			return n
		}
		n += 1
	}
	return -1
}

func part1(yMin int) int {
	// Any positive y speed will become the same speed in the negative direction at y=0
	// For the fastest y speed we want it to just be within the target area in one step
	// Once it reaches y=0 it will next move one space further than original y speed
	ySpeed := -yMin - 1
	return triangularNumber(ySpeed)
}

func part2(xMin, xMax, yMin, yMax int) int {
	// Min x velocity will only just reach or exceed the left edge
	xVelocityMin := getMinVelocity(xMin)
	// Max x velocity will take us to the far right edge in one step
	xVelocityMax := xMax
	// Min y velocity (negative) will take us to the bottom edge in one step
	yVelocityMin := yMin
	// Max y velocity (positive) is one less than the distance to the bottom edge
	yVelocityMax := -yMin - 1

	count := 0
	for x := xVelocityMin; x <= xVelocityMax; x++ {
		for y := yVelocityMin; y <= yVelocityMax; y++ {
			xVel := x
			yVel := y
			xPos := 0
			yPos := 0
			for n := 0; xPos <= xMax && yPos >= yMin; n++ {
				// Update x postion and speed
				if xVel > 0 {
					xPos += xVel
					xVel -= 1
				}

				// Update y position and speed
				yPos += yVel
				yVel -= 1

				// Check if position is within the target area
				if xPos >= xMin && xPos <= xMax && yPos >= yMin && yPos <= yMax {
					count += 1
					break
				}
			}
		}
	}
	return count
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	x1, x2, y1, y2 := parseData(data)

	p1 := part1(y1)
	fmt.Println("Part 1:", p1)

	p2 := part2(x1, x2, y1, y2)
	fmt.Println("Part 2:", p2)
}
