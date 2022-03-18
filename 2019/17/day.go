// Advent of Code 2019 - Day 17
package day17

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/maze-mapper/advent-of-code/2019/intcode"
)

// runProgram runs the given ASCII program and returns the output
func runProgram(program []int, input []byte) []int {
	computer := intcode.New(program)
	chanIn := make(chan int)
	chanOut := make(chan int)
	computer.SetChanIn(chanIn)
	computer.SetChanOut(chanOut)

	go computer.Run()

	go func() {
		for _, in := range input {
			chanIn <- int(in)
		}
	}()

	output := []int{}
	for out := range chanOut {
		output = append(output, out)
	}
	return output
}

// splitOutput removes trailing new lines and splits on new line characters
func splitOutput(output []byte) [][]byte {
	sep := []byte("\n")
	lines := bytes.Split(
		bytes.TrimSuffix(output, append(sep, sep...)), sep,
	)
	return lines
}

// Constants for turning the robot left or right
const (
	left  = "L"
	right = "R"
)

// Constants for the robot direction
const (
	north = iota
	south
	west
	east
)

// findRobot returns the position and direction of the robot
func findRobot(data [][]byte) (int, int, int) {
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			switch data[i][j] {
			case byte('^'):
				return i, j, north
			case byte('v'):
				return i, j, south
			case byte('<'):
				return i, j, west
			case byte('>'):
				return i, j, east
			}
		}
	}
	log.Fatal("Unable to find robot position")
	return 0, 0, 0
}

// tryMove returns the position of the robot if it moves in the given direction and whether or that is scaffold
func tryMove(data [][]byte, x, y, direction int) (int, int, bool) {
	switch direction {

	case north:
		if y == 0 {
			return x, y, false
		} else {
			y -= 1
		}

	case south:
		if y == len(data)-1 {
			return x, y, false
		} else {
			y += 1
		}

	case west:
		if x == 0 {
			return x, y, false
		} else {
			x -= 1
		}

	case east:
		if x == len(data[y])-1 {
			return x, y, false
		} else {
			x += 1
		}

	}

	return x, y, data[y][x] == byte('#')
}

// doTurnLeft turns left if there is scaffold in that direction
func doTurnLeft(data [][]byte, x, y, direction int) (int, bool) {
	var pos byte
	var dir int

	switch direction {

	case north:
		if x == 0 {
			return dir, false
		} else {
			dir = west
			pos = data[y][x-1]
		}

	case south:
		if x == len(data[y])-1 {
			return dir, false
		} else {
			dir = east
			pos = data[y][x+1]
		}

	case west:
		if y == len(data)-1 {
			return dir, false
		} else {
			dir = south
			pos = data[y+1][x]
		}

	case east:
		if y == 0 {
			return dir, false
		} else {
			dir = north
			pos = data[y-1][x]
		}

	}

	return dir, pos == byte('#')
}

// doTurnRight turns right if there is scaffold in that direction
func doTurnRight(data [][]byte, x, y, direction int) (int, bool) {
	var pos byte
	var dir int

	switch direction {

	case north:
		if x == len(data[y])-1 {
			return dir, false
		} else {
			dir = east
			pos = data[y][x+1]
		}

	case south:
		if x == 0 {
			return dir, false
		} else {
			dir = west
			pos = data[y][x-1]
		}

	case west:
		if y == 0 {
			return dir, false
		} else {
			dir = north
			pos = data[y-1][x]
		}

	case east:
		if y == len(data)-1 {
			return dir, false
		} else {
			dir = south
			pos = data[y+1][x]
		}

	}

	return dir, pos == byte('#')
}

// plotRoute returns the moves for traversing the scaffold
func plotRoute(data [][]byte) []string {
	y, x, direction := findRobot(data)
	route := []string{}

	for {
		// Try to move forward
		var steps int
		for {
			if xx, yy, ok := tryMove(data, x, y, direction); ok {
				x = xx
				y = yy
				steps += 1
			} else {
				break
			}
		}
		if steps > 0 {
			s := strconv.Itoa(steps)
			route = append(route, string(s))
		}

		// Try to turn and break if unable
		if d, ok := doTurnLeft(data, x, y, direction); ok {
			direction = d
			route = append(route, left)
		} else if d, ok := doTurnRight(data, x, y, direction); ok {
			direction = d
			route = append(route, right)
		} else {
			break
		}

	}

	return route
}

// sliceEqual compares two string slices for equality
func sliceEqual(s, ss []string) bool {
	if len(s) == len(ss) {
		for i := 0; i < len(s); i++ {
			if s[i] != ss[i] {
				return false
			}
		}
	} else {
		return false
	}
	return true
}

func part1(data [][]byte) int {
	scaff := byte('#')
	sum := 0
	for i := 1; i < len(data)-1; i++ {
		for j := 1; j < len(data[i])-1; j++ {
			// Check for scaffold intersection
			if data[i][j] == scaff && data[i-1][j] == scaff && data[i+1][j] == scaff && data[i][j-1] == scaff && data[i][j+1] == scaff {
				sum += i * j
			}
		}
	}
	return sum
}

// compress breaks the route down in to one main movement routine and three movement functions
func compress(route []string) (string, string, string, string) {
	var a, b, c []string
	// max length is 20 bytes so without commas this would be 10 single characters
	minFuncLen := 2
	maxFuncLen := 10
	// Moves are turn or move forwards so expect these to come in pairs
	stepSize := 2

	for lenA := minFuncLen; lenA <= maxFuncLen; lenA += stepSize {
		// Choose a function A
		a = route[:lenA]
		mA := []string{"A"}

		startB := lenA
		// Skip any occurrences of function A
	loop1:
		for {
			if startB+lenA <= len(route) && sliceEqual(route[startB:startB+lenA], a) {
				startB += lenA
				mA = append(mA, "A")
			} else {
				break loop1
			}
		}

		for lenB := minFuncLen; lenB <= maxFuncLen; lenB += stepSize {
			// Choose a function B
			b = route[startB : startB+lenB]
			startC := startB + lenB
			mB := []string{"B"}

			// Skip any occurrences of functions A or B
		loop2:
			for {
				switch {
				case startC+lenA <= len(route) && sliceEqual(route[startC:startC+lenA], a):
					startC += lenA
					mB = append(mB, "A")
				case startC+lenB <= len(route) && sliceEqual(route[startC:startC+lenB], b):
					startC += lenB
					mB = append(mB, "B")
				default:
					break loop2
				}
			}

			for lenC := minFuncLen; lenC <= maxFuncLen; lenC += stepSize {
				// Choose a function C
				c = route[startC : startC+lenC]
				end := startC + lenC
				mC := []string{"C"}

				// Skip any occurrences of functions A, B or C until we reach the end or can't match
			loop3:
				for end < len(route) {
					switch {
					case end+lenA <= len(route) && sliceEqual(route[end:end+lenA], a):
						end += lenA
						mC = append(mC, "A")
					case end+lenB <= len(route) && sliceEqual(route[end:end+lenB], b):
						end += lenB
						mC = append(mC, "B")
					case end+lenC <= len(route) && sliceEqual(route[end:end+lenC], c):
						end += lenC
						mC = append(mC, "C")
					default:
						break loop3
					}
				}

				// Check that all functions do not exceed 20 bytes
				if end == len(route) {
					m := append(mA, mB...)
					m = append(m, mC...)

					aOut := strings.Join(a, ",")
					bOut := strings.Join(b, ",")
					cOut := strings.Join(c, ",")
					mOut := strings.Join(m, ",")
					if len(aOut) <= 20 && len(bOut) <= 20 && len(cOut) <= 20 && len(mOut) <= 20 {
						return mOut, aOut, bOut, cOut
					}
				}
			}
		}
	}
	return "", "", "", ""
}

// makeInput converts the string movement routine and functions in to a byte slice
func makeInput(m, a, b, c string) []byte {
	var builder strings.Builder
	components := []string{m, a, b, c}
	for _, component := range components {
		builder.WriteString(component)
		builder.WriteString("\n")
	}
	// Choose "n" for continuous video feed
	builder.WriteString("n\n")
	return []byte(builder.String())
}

func part2(data [][]byte, program []int) int {
	route := plotRoute(data)

	m, a, b, c := compress(route)
	input := makeInput(m, a, b, c)

	program[0] = 2
	output := runProgram(program, input)

	return output[len(output)-1]
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	program := intcode.ReadProgram(data)

	ioutput := runProgram(program, []byte{})

	output := make([]byte, len(ioutput))
	for i, o := range ioutput {
		// Unsafe if o overflows byte
		output[i] = byte(o)
	}

	fmt.Print(string(output))

	lines := splitOutput(output)

	p1 := part1(lines)
	fmt.Println("Part 1:", p1)

	p2 := part2(lines, program)
	fmt.Println("Part 2:", p2)
}
