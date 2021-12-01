// Advent of Code 2018 - Day 18
package day18

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// Define constants to represent elements in the lumber collection area
const (
	openGround = '.'
	trees      = '|'
	lumberyard = '#'
)

// parseData reads the input text file and returns a data structure for the lumber collection area
func parseData(file string) [][]rune {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(
		strings.TrimSuffix(string(data), "\n"), "\n",
	)

	region := make([][]rune, len(lines))
	for i, line := range lines {
		region[i] = make([]rune, len(line))
		for j, r := range line {
			region[i][j] = r
		}
	}

	return region
}

// getNeighbours finds counts of each neighbouring terrain type
func getNeighbours(region [][]rune, x, y int) map[rune]int {
	neighbours := map[rune]int{}

	for i := y - 1; i <= y+1; i++ {
		// Handle edges
		if i < 0 || i >= len(region) {
			continue
		}

		for j := x - 1; j <= x+1; j++ {
			// Handle edges
			if j < 0 || j >= len(region[i]) {
				continue
			}

			// Skip centre
			if i == y && j == x {
				continue
			}

			n := region[i][j]
			neighbours[n] += 1
		}
	}

	return neighbours
}

// tick simulates the lumber collection area after one minute
func tick(region [][]rune) [][]rune {
	newRegion := make([][]rune, len(region))
	for i, row := range region {
		newRegion[i] = make([]rune, len(row))
		for j, col := range row {
			neighbours := getNeighbours(region, j, i)
			switch col {
			case openGround:
				if neighbours[trees] >= 3 {
					newRegion[i][j] = trees
				} else {
					newRegion[i][j] = openGround
				}

			case trees:
				if neighbours[lumberyard] >= 3 {
					newRegion[i][j] = lumberyard
				} else {
					newRegion[i][j] = trees
				}

			case lumberyard:
				if neighbours[lumberyard] >= 1 && neighbours[trees] >= 1 {
					newRegion[i][j] = lumberyard
				} else {
					newRegion[i][j] = openGround
				}

			default:
				log.Fatal("Unknown terrain type")
			}

		}
	}
	return newRegion
}

// printRegion prints the lumber collection area
func printRegion(region [][]rune) {
	ANSIGrey := "\033[30m\033[40m"
	ANSIGreen := "\033[32m\033[42m"
	ANSIYellow := "\033[33m\033[43m"
	ANSIClear := "\033[0m"
	for _, row := range region {
		for _, col := range row {
			switch col {
			case openGround:
				fmt.Print(ANSIYellow)
			case trees:
				fmt.Print(ANSIGreen)
			case lumberyard:
				fmt.Print(ANSIGrey)
			}
			fmt.Print(string(col))
			fmt.Print(ANSIClear)
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

// countElement returns the count of a particular rune from the lumber collection area
func countElement(region [][]rune, r rune) int {
	count := 0
	for _, row := range region {
		for _, col := range row {
			if col == r {
				count += 1
			}
		}
	}
	return count
}

// resourceValue returns the resource value of the lumber collection area
func resourceValue(region [][]rune) int {
	treesCount := countElement(region, trees)
	lumberyardCount := countElement(region, lumberyard)
	return treesCount * lumberyardCount
}

func Run(inputFile string) {
	region := parseData(inputFile)
	//	printRegion(region)
	part1Time := 10
	for i := 1; i <= part1Time; i++ {
		region = tick(region)
		//		printRegion(region)
	}
	fmt.Println("Part 1: Resource value after", part1Time, "minutes is", resourceValue(region))

	part2Time := 1000000000
	resourceValuesSeen := map[int]int{}
	for i := part1Time + 1; i <= part2Time; i++ {
		region = tick(region)
		//		printRegion(region)
		rv := resourceValue(region)
		if lastSeen, ok := resourceValuesSeen[rv]; ok {
			// Hopefully the region has reached a stable cycle
			period := i - lastSeen
			remainingTime := part2Time - i
			// Break when we are a complete number of cyles from our end goal
			if remainingTime%period == 0 {
				break
			}
		} else {
			resourceValuesSeen[rv] = i
		}
	}
	fmt.Println("Part 2: Resource value after", part2Time, "minutes is", resourceValue(region))
}
