// Advent of Code 2018 - Day 17
package day17

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

// Define constants to represent elements in the reservoir
const (
	spring       = '+'
	clay         = '#'
	sand         = '.'
	runningWater = '|'
	stillWater   = '~'
)

// clayVein represents a straight vein of clay
type clayVein struct {
	orientation                  string
	fixCoord, minCoord, maxCoord int
}

// parseData reads the input text file and returns a data structure for the reservoir
func parseData(file string) ([][]rune, int) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	// Regular expression for each line
	// Not strictly accurate as coordinates should differ but makes submatches simpler
	pattern := regexp.MustCompile(`(x|y)=([0-9]+), (x|y)=([0-9]+)\.\.([0-9]+)`)

	lines := strings.Split(
		strings.TrimSuffix(string(data), "\n"), "\n",
	)
	clayVeins := make([]*clayVein, len(lines))

	// Bounds for reservoir
	minX := int(^uint(0)>>1) - 1
	maxX := 0
	minY := int(^uint(0)>>1) - 1
	maxY := 0

	for i, line := range lines {
		if match := pattern.FindStringSubmatch(line); match != nil {
			// Extract numeric coordinate values
			fixCoord, _ := strconv.Atoi(match[2])
			minCoord, _ := strconv.Atoi(match[4])
			maxCoord, _ := strconv.Atoi(match[5])

			// Update potential size of reservoir region
			switch match[1] {
			case "x":
				if fixCoord < minX {
					minX = fixCoord
				}
				if fixCoord > maxX {
					maxX = fixCoord
				}
				if minCoord < minY {
					minY = minCoord
				}
				if maxCoord > maxY {
					maxY = maxCoord
				}

			case "y":
				if fixCoord < minY {
					minY = fixCoord
				}
				if fixCoord > maxY {
					maxY = fixCoord
				}
				if minCoord < minX {
					minX = minCoord
				}
				if maxCoord > maxX {
					maxX = maxCoord
				}
			}

			clayVeins[i] = &clayVein{
				orientation: match[1],
				fixCoord:    fixCoord,
				minCoord:    minCoord,
				maxCoord:    maxCoord,
			}
		} else {
			log.Fatal("Unable to match regular expression to line", line)
		}
	}

	// Extend x range by one in each direction to allow for any spill over
	minX -= 1
	maxX += 1
	width := maxX - minX

	// Initialise reservoir
	reservoir := make([][]rune, maxY+1)
	for i := range reservoir {
		reservoir[i] = make([]rune, width+1)
		for j := range reservoir[i] {
			reservoir[i][j] = sand
		}
	}

	// Add spring at shifted x coordinate from (x=500, y=0)
	reservoir[0][500-minX] = spring

	// Add clay veins with x coordinate shift
	for _, cv := range clayVeins {
		switch cv.orientation {
		case "x":
			for i := cv.minCoord; i <= cv.maxCoord; i++ {
				reservoir[i][cv.fixCoord-minX] = clay
			}
		case "y":
			for j := cv.minCoord; j <= cv.maxCoord; j++ {
				reservoir[cv.fixCoord][j-minX] = clay
			}
		}
	}

	return reservoir, minY
}

// printReservoir prints the reservoir
func printReservoir(reservoir [][]rune) {
	ANSIGrey := "\033[30m\033[40m"
	ANSIYellow := "\033[33m\033[43m"
	ANSIBlue := "\033[34m\033[44m"
	ANSICyan := "\033[36m\033[46m"
	ANSIClear := "\033[0m"
	for _, row := range reservoir {
		for _, col := range row {
			switch col {
			case runningWater, spring:
				fmt.Print(ANSICyan)
			case stillWater:
				fmt.Print(ANSIBlue)
			case clay:
				fmt.Print(ANSIGrey)
			case sand:
				fmt.Print(ANSIYellow)
			}
			fmt.Print(string(col))
			fmt.Print(ANSIClear)
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

// canFlowDown checks if water can flow downwards from the given coordinates and updates the reservoir
func canFlowDown(reservoir [][]rune, x, y int) bool {
	switch reservoir[y+1][x] {
	case sand:
		reservoir[y+1][x] = runningWater
		return true
	case runningWater:
		return true
	default:
		return false
	}
}

// isContained checks if water is contained in a reservoir and updates the reservoir
func isContained(reservoir [][]rune, x, y int) bool {
	contained := true

	// Check to the left
	xLeft := x
	for i := xLeft - 1; i >= 0; i-- {
		// Check if we reach a clay wall
		if reservoir[y][i] == clay {
			break
		}
		xLeft -= 1

		// Check if water can flow downwards
		if canFlowDown(reservoir, i, y) {
			contained = false
			break
		}
	}

	// Check to the right
	xRight := x
	for i := xRight + 1; i < len(reservoir[y]); i++ {
		// Check if we reach a clay wall
		if reservoir[y][i] == clay {
			break
		}
		xRight += 1

		// Check if water can flow downwards
		if canFlowDown(reservoir, i, y) {
			contained = false
			break
		}
	}

	// Fill in either still or running water
	fill := runningWater
	if contained {
		fill = stillWater
	}
	for i := xLeft; i <= xRight; i++ {
		reservoir[y][i] = fill
	}

	return contained
}

// simulate models the flow of water in a reservoir region from the source
func simulate(reservoir [][]rune) {
	for i := 0; i < len(reservoir)-1; {
		jump := 1
		for j := 0; j < len(reservoir[i]); j++ {
			if reservoir[i][j] == runningWater || reservoir[i][j] == spring {
				if canFlowDown(reservoir, j, i) {
					// Space below will now be filled with water
				} else if isContained(reservoir, j, i) {
					// Spaces to the side will now be filled with still or running water
					jump = -1
				}
			}
		}
		i += jump
	}
}

// countElement returns the count of a particular rune from the reservoir
func countElement(reservoir [][]rune, r rune, fromRow int) int {
	count := 0
	for i := fromRow; i < len(reservoir); i++ {
		for j := 0; j < len(reservoir[i]); j++ {
			if reservoir[i][j] == r {
				count += 1
			}
		}
	}
	return count
}

func Run(inputFile string) {
	reservoir, minY := parseData(inputFile)
	simulate(reservoir)
	//	printReservoir(reservoir)
	runningWaterCount := countElement(reservoir, runningWater, minY)
	stillWaterCount := countElement(reservoir, stillWater, minY)
	fmt.Println("Part 1:", runningWaterCount+stillWaterCount)
	fmt.Println("Part 2:", stillWaterCount)
}
