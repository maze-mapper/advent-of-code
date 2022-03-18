// Advent of Code 2021 - Day 13
package day13

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/maze-mapper/advent-of-code/coordinates"
)

type fold struct {
	axis  string
	value int
}

func parseData(data []byte) (map[coordinates.Coord]struct{}, []fold) {
	parts := strings.Split(
		strings.TrimSuffix(string(data), "\n"), "\n\n",
	)

	points := map[coordinates.Coord]struct{}{}
	for _, line := range strings.Split(parts[0], "\n") {
		var x, y int
		fmt.Sscanf(line, "%d,%d", &x, &y)
		c := coordinates.Coord{X: x, Y: y}
		points[c] = struct{}{}
	}

	foldLines := strings.Split(parts[1], "\n")
	folds := make([]fold, len(foldLines))
	for i, line := range foldLines {
		line = strings.Replace(line, "=", " ", 1)
		var axis string
		var value int
		fmt.Sscanf(line, "fold along %s %d", &axis, &value)
		folds[i] = fold{axis: axis, value: value}
	}

	return points, folds
}

func doFold(points *map[coordinates.Coord]struct{}, f fold) {
	tempPoints := map[coordinates.Coord]struct{}{}
	for p, _ := range *points {
		switch f.axis {
		case "x":
			if p.X > f.value {
				p.X -= f.value
				p.X = -p.X
				p.X += f.value
			}
		case "y":
			if p.Y > f.value {
				p.Y -= f.value
				p.Y = -p.Y
				p.Y += f.value
			}
		}
		tempPoints[p] = struct{}{}
	}
	*points = tempPoints
}

func part1(points *map[coordinates.Coord]struct{}, f fold) int {
	doFold(points, f)
	return len(*points)
}

func part2(points *map[coordinates.Coord]struct{}, folds []fold) {
	for _, f := range folds {
		doFold(points, f)
	}

	finalPoints := []coordinates.Coord{}
	for k, _ := range *points {
		finalPoints = append(finalPoints, k)
	}
	min, max := coordinates.Range(finalPoints)

	for i := min.Y; i <= max.Y; i++ {
		for j := min.X; j <= max.X; j++ {
			c := coordinates.Coord{X: j, Y: i}
			if _, ok := (*points)[c]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	points, folds := parseData(data)

	p1 := part1(&points, folds[0])
	fmt.Println("Part 1:", p1)

	fmt.Println("Part 2:")
	part2(&points, folds[1:])
}
