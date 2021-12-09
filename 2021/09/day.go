// Advent of Code 2021 - Day 9
package day9

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"

	"adventofcode/coordinates"
)

func parseData(data []byte) [][]int {
	lines := strings.Split(
		strings.TrimSuffix(string(data), "\n"), "\n",
	)

	heightmap := make([][]int, len(lines))
	for i, line := range lines {
		characters := strings.Split(line, "")
		heightmap[i] = make([]int, len(characters))
		for j, s := range characters {
			if val, err := strconv.Atoi(s); err == nil {
				heightmap[i][j] = val
			} else {
				log.Fatal(err)
			}
		}
	}
	return heightmap
}

func getNeighbours(c coordinates.Coord, heightmap [][]int) []coordinates.Coord {
	neighbours := []coordinates.Coord{}
	if c.X != 0 {
		neighbours = append(neighbours, coordinates.Coord{X: c.X - 1, Y: c.Y})
	}
	if c.X != len(heightmap[c.Y])-1 {
		neighbours = append(neighbours, coordinates.Coord{X: c.X + 1, Y: c.Y})
	}
	if c.Y != 0 {
		neighbours = append(neighbours, coordinates.Coord{X: c.X, Y: c.Y - 1})
	}
	if c.Y != len(heightmap)-1 {
		neighbours = append(neighbours, coordinates.Coord{X: c.X, Y: c.Y + 1})
	}
	return neighbours
}

func findLowPoints(heightmap [][]int) []coordinates.Coord {
	lowPoints := []coordinates.Coord{}
	for i := 0; i < len(heightmap); i++ {
		for j := 0; j < len(heightmap[i]); j++ {
			c := coordinates.Coord{X: j, Y: i}
			neighbours := getNeighbours(c, heightmap)
			lowPoint := true
			for _, n := range neighbours {
				if heightmap[i][j] >= heightmap[n.Y][n.X] {
					lowPoint = false
					break
				}
			}
			if lowPoint == true {
				lowPoints = append(lowPoints, c)
			}
		}
	}
	return lowPoints
}

func growBasin(edgePoints []coordinates.Coord, basin map[coordinates.Coord]struct{}, heightmap [][]int) []coordinates.Coord {
	newEdge := []coordinates.Coord{}
	for _, p := range edgePoints {
		neighbours := getNeighbours(p, heightmap)
		for _, n := range neighbours {
			if heightmap[n.Y][n.X] == 9 {
				continue
			}
			if _, ok := basin[n]; !ok {
				basin[n] = struct{}{}
				newEdge = append(newEdge, n)
			}
		}

	}
	return newEdge
}

func findBasin(lowPoint coordinates.Coord, heightmap [][]int) int {
	edgePoints := []coordinates.Coord{lowPoint}
	basin := map[coordinates.Coord]struct{}{
		lowPoint: struct{}{},
	}
	for len(edgePoints) != 0 {
		edgePoints = growBasin(edgePoints, basin, heightmap)
	}

	return len(basin)
}

func part1(heightmap [][]int) int {
	riskLevelSum := 0
	lowPoints := findLowPoints(heightmap)
	for _, p := range lowPoints {
		riskLevelSum += 1 + heightmap[p.Y][p.X]
	}
	return riskLevelSum
}

func part2(heightmap [][]int) int {
	lowPoints := findLowPoints(heightmap)
	basinSizes := []int{}
	for _, p := range lowPoints {
		size := findBasin(p, heightmap)
		basinSizes = append(basinSizes, size)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))
	return basinSizes[0] * basinSizes[1] * basinSizes[2]
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	heightmap := parseData(data)

	p1 := part1(heightmap)
	fmt.Println("Part 1:", p1)

	p2 := part2(heightmap)
	fmt.Println("Part 2:", p2)
}
