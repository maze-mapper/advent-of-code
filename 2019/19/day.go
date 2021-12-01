// Advent of Code 2019 - Day 19
package day19

import (
	"fmt"
	"io/ioutil"
	"log"

	"adventofcode/2019/intcode"
)

// Constants for whether or not beam is affecting drone
const (
	stationary = 0
	pulled     = 1
)

// checkPoint runs the Intcode program to check if a point is pulled by the beam
func checkPoint(program []int, x, y int) int {
	computer := intcode.New(program)
	chanIn := make(chan int)
	chanOut := make(chan int)
	computer.SetChanIn(chanIn)
	computer.SetChanOut(chanOut)

	go computer.Run()

	chanIn <- x
	chanIn <- y
	return <-chanOut
}

func part1(program []int, size int) int {
	pulledByBeam := 0
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			if checkPoint(program, x, y) == pulled {
				pulledByBeam += 1
			}
		}
	}
	return pulledByBeam
}

func part2(program []int, size int) int {
	lhs := 0
	// First few rows may not have the beam in them
	for y := 10; ; y++ {
		reachedBeam := false
		for x := lhs; ; x++ {
			if checkPoint(program, x, y) == pulled {
				// Update x position of left hand side of beam when we first reach it
				if reachedBeam == false {
					lhs = x
					reachedBeam = true
				}

				// Check if point to the right is outside the beam
				if checkPoint(program, x+size-1, y) == stationary {
					break
				}

				// Check if point downwards is within the beam
				if checkPoint(program, x, y+size-1) == pulled {
					return x*10000 + y
				}
			}
		}
	}
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	program := intcode.ReadProgram(data)

	p1 := part1(program, 50)
	fmt.Println("Part 1:", p1)

	p2 := part2(program, 100)
	fmt.Println("Part 2:", p2)
}
