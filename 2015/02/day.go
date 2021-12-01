// Advent of Code 2015 - Day 2
package day2

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"sync"
)

var mutex sync.Mutex

// parseDimensionString converts a string like 1x2x3 in to an array of integers
func parseDimensionString(line string) [3]int {
	strDim := strings.Split(line, "x")
	dim := [3]int{}
	for i, c := range strDim {
		dim[i], _ = strconv.Atoi(c)
	}
	return dim
}

// volume calculates the volume of a cuboid
func volume(dim [3]int) int {
	vol := 1
	for _, d := range dim {
		vol *= d
	}
	return vol
}

// solve calculates the answer for both parts of the problem for a single present
func solve(line string, paper *int, ribbon *int, wg *sync.WaitGroup) {
	dim := parseDimensionString(line)
	var area int
	minArea := int(^uint(0) >> 1)   // Initialise to max value
	minLength := int(^uint(0) >> 1) // Initialise to max value

	for i := 0; i < len(dim)-1; i++ {
		for j := i + 1; j < len(dim); j++ {
			sideArea := dim[i] * dim[j]
			if sideArea < minArea {
				minArea = sideArea
			}
			area += sideArea

			sideLength := dim[i] + dim[j]
			if sideLength < minLength {
				minLength = sideLength
			}
		}
	}

	area *= 2
	area += minArea

	length := 2 * minLength
	length += volume(dim)

	mutex.Lock()
	*paper += area
	*ribbon += length
	mutex.Unlock()
	wg.Done()
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	text := string(data)

	var wg sync.WaitGroup
	var paper, ribbon int
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		wg.Add(1)
		go solve(line, &paper, &ribbon, &wg)
	}
	wg.Wait()

	fmt.Println("Part 1 answer:", paper)
	fmt.Println("Part 2 answer:", ribbon)
}
