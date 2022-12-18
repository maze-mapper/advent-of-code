// Advent of Code 2022 - Day 17.
package day17

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strings"
)

var shapeOrder = [][][]bool{
	// ####
	[][]bool{
		[]bool{true, true, true, true},
	},
	// .#.
	// ###
	// .#.
	[][]bool{
		[]bool{false, true, false},
		[]bool{true, true, true},
		[]bool{false, true, false},
	},
	// ..# ###
	// ..# ..#
	// ### ..#
	[][]bool{
		[]bool{true, true, true},
		[]bool{false, false, true},
		[]bool{false, false, true},
	},
	// #
	// #
	// #
	// #
	[][]bool{
		[]bool{true},
		[]bool{true},
		[]bool{true},
		[]bool{true},
	},
	// ##
	// ##
	[][]bool{
		[]bool{true, true},
		[]bool{true, true},
	},
}

func printChamber(chamber, shape [][]bool, posX, posY, currentHeight int) {
	shapeCoords := map[[2]int]bool{}
	for y := len(shape) - 1; y >= 0; y-- {
		for x, b := range shape[y] {
			if b {
				shapeCoords[[2]int{posY + y, posX + x}] = true
			}
		}
	}

	for y := currentHeight ; y >= 0; y-- {
		for x, b := range chamber[y] {
			c := [2]int{y, x}
			if _, ok := shapeCoords[c]; ok {
				fmt.Print("@")
			} else if b {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}

		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func canMove(chamber, shape [][]bool, posX, posY, dirX, dirY int) bool {
	// Check if floor and wall block movement.
	if dirX < 0 && posX == 0 || dirX > 0 && posX+len(shape[0]) == len(chamber[0]) || dirY < 0 && posY == 0 {
		return false
	}
	for y := len(shape) - 1; y >= 0; y-- {
		for x, b := range shape[y] {
//			fmt.Println(posY+y+dirY, posX+x+dirX)
			if b && chamber[posY+y+dirY][posX+x+dirX] {
				return false
			}
		}

	}
	return true
}

type info struct {
	lastShapeNum int
	currentHeight int
}

func tetris(jets string, numShapes int, maxHeight int) int {
	width := 7
//	maxHeight := 4 * numShapes
	chamber := make([][]bool, maxHeight)
	for i := range chamber {
		chamber[i] = make([]bool, width)
	}
	frontier := make([]int, width)
	pastFrontier := false
	// jet index mapped to shape mapped to height
//	states := map[int]map[int]info{}
	states := make([]map[int]info, len(jets))

	currentHeight := -1
	bonusHeight := 0
//	truncatedHeight := 0

	// Bottom left corner of shape.
	newPosX := 2
	newPosY := 4

	var posX, posY int
	var shape [][]bool
	j := 0
	for s := 0; s < numShapes; s++ {
		// Truncate chamber.
/*		minFrontier := int(math.MaxInt)
		for _, f := range frontier {
			if f < minFrontier {
					minFrontier = f
			}
		}*/
//		fmt.Println(frontier, minFrontier, len(chamber), truncatedHeight, chamber[0])
/*		if minFrontier > 0 {
//			printChamber(chamber, shape, posX, posY, currentHeight + newPosY + 4)
			newRows := make([][]bool, minFrontier)
			for i := range newRows {
				newRows[i] = make([]bool, width)
			}
			chamber = append(chamber[minFrontier:], newRows...)
			for i := range frontier {
				frontier[i] -= minFrontier
			}
			truncatedHeight += minFrontier
			currentHeight -= minFrontier
//			printChamber(chamber, shape, posX, posY, currentHeight + newPosY + 4)
		}*/


		shapeIndex := s%len(shapeOrder)
		shape = shapeOrder[shapeIndex]
		posX = newPosX
		posY = currentHeight + newPosY
		falling := true

//		fmt.Printf("Shape number %d has index %d has left corner at x=%d, y=%d\n", s, s%len(shapeOrder), posX, posY)

//		fmt.Println(frontier)
		for falling {
//			printChamber(chamber, shape, posX, posY, currentHeight + newPosY + 4)
			// Movement from jet.
			j = j % len(jets)
			var dir int
			switch string(jets[j]) {
			case "<":
				dir = -1
			case ">":
				dir = 1
			}
			j += 1
			if canMove(chamber, shape, posX, posY, dir, 0) {
				posX += dir
			}
		//	fmt.Printf("Jet movement, now x=%d\n", posX)

			// Movement down.
			falling = canMove(chamber, shape, posX, posY, 0, -1)
			if falling {
				posY -= 1
			}
		//	fmt.Printf("falling=%t at x=%d, y=%d\n", falling, posX, posY)
		}

		//fmt.Printf("Update pos at x=%d, y=%d\n", posX, posY)
		for y := range shape {
			for x, b := range shape[y] {
				if b {
					chamber[posY+y][posX+x] = true

					if posY+y + 1 > frontier[posX+x] {
						frontier[posX+x] = posY+y + 1
					}
				}
				if posY+y > currentHeight {
					currentHeight = posY + y
				}
			}
		}

		if !pastFrontier {
                        minFrontier := int(math.MaxInt)
                        for _, f := range frontier {
                                if f < minFrontier {
                                        minFrontier = f
                                }
                        }
                        if minFrontier > 0 {
                                pastFrontier = true
                        }
                }

		// TODO make states an array of len j.
		if inf, ok := states[j][shapeIndex]; ok { 
//		if m, ok := states[j]; ok {
//			if inf, okk := m[shapeIndex]; okk {
				fmt.Printf("Found same state: turn=%d, height=%d vs turn=%d, height=%d\n", s, currentHeight, inf.lastShapeNum, inf.currentHeight)
				turnDiff := s - inf.lastShapeNum
				heightDiff := currentHeight - inf.currentHeight

//				if s + turnDiff > numShapes {
//					continue
//				}

				for ss := s + turnDiff; ss < numShapes; ss += turnDiff {
					s += turnDiff
//					currentHeight += heightDiff
					bonusHeight += heightDiff
//					fmt.Printf("Turn %d, height %d\n", s, currentHeight)
				}

/*				remainingTurns := numShapes - s
				jTarget := j + remainingTurns
				newShapeIdx := (shapeIndex + remainingTurns) % len(shapeOrder)

				leftHeight := states[jTarget][newShapeIdx].currentHeight - inf.currentHeight
				bonusHeight += leftHeight
				s += remainingTurns*/
		//	}
		} else if pastFrontier {
			if states[j] == nil {
				states[j] = map[int]info{}
			}
			states[j][shapeIndex] = info{
					lastShapeNum: s,
					currentHeight: currentHeight,
				}
			
		}


//		printChamber(chamber, shape, posX, posY, currentHeight)

	}

	return currentHeight + bonusHeight + 1
}

func part1(jets string) int {
        return tetris(jets, 2022, 2022 * 4)
}

func part2(jets string) int {
	return tetris(jets, 1000000000000, 1000000)
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	jets := strings.TrimSuffix(string(data), "\n")

	p1 := part1(jets)
	fmt.Println("Part 1:", p1)

	p2 := part2(jets)
	fmt.Println("Part 2:", p2)
}
