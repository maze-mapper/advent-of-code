// Advent of Code 2022 - Day 14.
package day14

import (
        "fmt"
        "io/ioutil"
        "log"
        "math"
        "strconv"
        "strings"

	"github.com/maze-mapper/advent-of-code/coordinates"
)

func parseInput(data []byte) ([][]rune, int, int) {
	lines := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
	rocks := make([][]coordinates.Coord, len(lines))

	minX := int(math.MaxInt)
	maxX := 0
	minY := 0
	maxY := 0

	for i, line := range lines {
		parts := strings.Split(line, " -> ")
		c := make([]coordinates.Coord, len(parts))
		for j, p := range parts {
			co := strings.Split(p, ",")
			x, err := strconv.Atoi(co[0])
			if err != nil {
				log.Fatal(err)
			}
			y, err := strconv.Atoi(co[1])
                        if err != nil {
                                log.Fatal(err)
                        }
			c[j] = coordinates.Coord{X: x, Y: y}

			if x < minX {
				minX = x
			}
			if x > maxX {
				maxX = x
			}
			if y < minY {
				minY = y
			}
			if y > maxY {
				maxY = y
			}
		}
		rocks[i] = c
	}

	minX -= maxY
	maxX += maxY

	cave := make([][]rune, maxY + 1)
	for i := range cave {
		cave[i] = make([]rune, maxX - minX + 1)
	}

	for _, r := range rocks {
		var start, end coordinates.Coord
		for i := 0; i < len(r) - 1; i++ {
			start = r[i]
                        end = r[i+1]
			x1 := start.X
			x2 := end.X
			if x1 > x2 {
				x1, x2 = x2, x1
			}
			y1 := start.Y
			y2 := end.Y
			if y1 > y2 {
				y1, y2 = y2, y1
			}
			for y := y1; y <= y2; y++ {
				for x := x1; x <= x2; x++ {
					cave[y][x-minX] = '#'
				}
			}
		}
	}
	sourceX := 500 - minX
	sourceY := 0
	return cave, sourceX, sourceY
}

func simulate(cave [][]rune, sourceX, sourceY int) int {
	sandX := sourceX
	sandY := sourceY
	sandAtRest := 0

	loop:
	for {
		switch {
		case sandX == 0 || sandX == len(cave[sandY]) - 1 || sandY == len(cave) - 1 || cave[sourceY][sourceX] == 'o':
			break loop
		case cave[sandY+1][sandX] == rune(0):
			sandY += 1
		case cave[sandY+1][sandX-1] == rune(0):
			sandY += 1
			sandX -= 1
		case cave[sandY+1][sandX+1] == rune(0):
                        sandY += 1
                        sandX += 1
		default:
			cave[sandY][sandX] = 'o'
			sandAtRest += 1
			sandX = sourceX
			sandY = sourceY
		}
	}
	return sandAtRest
}



func printCave(cave [][]rune) {
	for _, row := range cave {
		for _, col := range row {
			if col == rune(0) {
				fmt.Print(" ")
			} else {
				fmt.Print(string(col))
			}
		}
		fmt.Print("\n")
	}
}

func part1(cave [][]rune, x, y int) int {
	return simulate(cave, x, y)
}

func part2(cave [][]rune, x, y int) int {
	cave = append(cave, make([]rune, len(cave[0])), make([]rune, len(cave[0])))
        for i := 0; i < len(cave[0]); i++ {
                cave[len(cave) - 1][i] = '#'
        }
	return simulate(cave, x, y)
}

func Run(inputFile string) {
        data, err := ioutil.ReadFile(inputFile)
        if err != nil {
                log.Fatal(err)
        }
        cave, x, y := parseInput(data)

        p1 := part1(cave, x, y)
        fmt.Println("Part 1:", p1)

        p2 := part2(cave, x, y) + p1
        fmt.Println("Part 2:", p2)
}

