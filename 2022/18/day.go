// Advent of Code 2022 - Day 18.
package day18

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/maze-mapper/advent-of-code/coordinates"
)

func parseInput(data []byte) map[coordinates.Coord]struct{} {
	lines := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
	out := make(map[coordinates.Coord]struct{}, len(lines))
	for _, line := range lines {
		c := coordinates.Coord{}
		n, err := fmt.Sscanf(line, "%d,%d,%d", &c.X, &c.Y, &c.Z)
		if err != nil {
			log.Fatal(err)
		}
		if n != 3 {
			log.Fatal("Expecting three coordinates.")
		}
		out[c] = struct{}{}
	}
	return out
}

func adjacentCubes(cube coordinates.Coord) []coordinates.Coord {
	adjacent := make([]coordinates.Coord, 6)
	adjacent[0] = coordinates.Coord{X: cube.X + 1, Y: cube.Y, Z: cube.Z}
	adjacent[1] = coordinates.Coord{X: cube.X - 1, Y: cube.Y, Z: cube.Z}
	adjacent[2] = coordinates.Coord{X: cube.X, Y: cube.Y + 1, Z: cube.Z}
	adjacent[3] = coordinates.Coord{X: cube.X, Y: cube.Y - 1, Z: cube.Z}
	adjacent[4] = coordinates.Coord{X: cube.X, Y: cube.Y, Z: cube.Z + 1}
	adjacent[5] = coordinates.Coord{X: cube.X, Y: cube.Y, Z: cube.Z - 1}
	return adjacent
}

func getNeighbours(c coordinates.Coord, r int) []coordinates.Coord {
	neighbours := []coordinates.Coord{}
	for x := -r; x <= r; x++ {
		for y := -r; y <= r; y++ {
			for z := -r; z <= r; z++ {
				neighbours = append(neighbours, coordinates.Coord{
					X: c.X + x,
					Y: c.Y + y,
					Z: c.Z + z,
				})
			}
		}
	}
	return neighbours
}

func part1(cubes map[coordinates.Coord]struct{}) int {
	area := 0
	for cube := range cubes {
		for _, adjacent := range adjacentCubes(cube) {
			if _, ok := cubes[adjacent]; !ok {
				area += 1
			}
		}
	}
	return area
}

// Returns true if cube is next to part of the droplet.
func testNeighbours(cubes map[coordinates.Coord]struct{}, c coordinates.Coord) bool {
	neighbours := adjacentCubes(c)
	for _, n := range neighbours {
		if _, ok := cubes[n]; ok {
			return true
		}
	}
	return false
}

func countDropletNeighbours(cubes map[coordinates.Coord]struct{}, c coordinates.Coord) int {
	neighbours := adjacentCubes(c)
	count := 0
	for _, n := range neighbours {
		if _, ok := cubes[n]; ok {
			count += 1
		}
	}
	return count
}

func findOuterSpace(cubes map[coordinates.Coord]struct{}) coordinates.Coord {
	all := make([]coordinates.Coord, len(cubes))
	i := 0
	for c := range cubes {
		all[i] = c
		i += 1
	}
	minCube, _ := coordinates.Range(all)
	minCube.Transform(coordinates.Coord{X: -1})

	toCheck := []coordinates.Coord{minCube}

	for len(toCheck) > 0 {
		c := toCheck[0]

		next := []coordinates.Coord{
			coordinates.Coord{X: c.X + 1, Y: c.Y, Z: c.Z},
			coordinates.Coord{X: c.X, Y: c.Y + 1, Z: c.Z},
			coordinates.Coord{X: c.X, Y: c.Y, Z: c.Z + 1},
		}
		for _, n := range next {
			if _, ok := cubes[n]; ok {
				return c
			}
		}

		toCheck = append(toCheck, next...)
		toCheck = toCheck[1:]
	}
	return coordinates.Coord{}
}

func part2(cubes map[coordinates.Coord]struct{}) int {
	start := findOuterSpace(cubes)
	visited := map[coordinates.Coord]int{}
	toConsider := []coordinates.Coord{start}

	for len(toConsider) > 0 {
		c := toConsider[0]
		toConsider = toConsider[1:]

		tooFar := true
		for cc := range cubes {
			if coordinates.ManhattanDistance(c, cc) <= 2 {
				tooFar = false
				break
			}
		}
		if tooFar {
			continue
		}

		neighbours := adjacentCubes(c)
		for _, n := range neighbours {
			// Already visited
			if _, ok := visited[n]; ok {
				continue
			}
			// Inside droplet
			if _, ok := cubes[n]; ok {
				continue
			}
			count := countDropletNeighbours(cubes, n)
			visited[n] = count
			toConsider = append(toConsider, n)
		}
	}

	area := 0
	for _, v := range visited {
		area += v
	}
	return area
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	cubes := parseInput(data)

	p1 := part1(cubes)
	fmt.Println("Part 1:", p1)

	p2 := part2(cubes)
	fmt.Println("Part 2:", p2)
}
