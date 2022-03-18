// Advent of Code 2021 - Day 19
package day19

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/maze-mapper/advent-of-code/coordinates"
)

// parseData returns the reports as a slice of coordinate slices
func parseData(data []byte) [][]coordinates.Coord {
	parts := strings.Split(
		strings.TrimSuffix(string(data), "\n"), "\n\n",
	)

	reports := make([][]coordinates.Coord, len(parts))

	for i, part := range parts {
		lines := strings.Split(part, "\n")
		reports[i] = make([]coordinates.Coord, len(lines)-1)
		for j, line := range lines[1:] {
			values := strings.Split(line, ",")
			c := coordinates.Coord{}
			for i, p := range []*int{&c.X, &c.Y, &c.Z} {
				if val, err := strconv.Atoi(values[i]); err == nil {
					*p = val
				} else {
					log.Fatal(err)
				}
			}
			reports[i][j] = c
		}
	}

	return reports
}

// combine returns a slice containing the unique coordinates in "a" and "b"
func combine(a, b []coordinates.Coord, v coordinates.Coord) []coordinates.Coord {
	// Construct map of unique coordinates in "a"
	m := map[coordinates.Coord]struct{}{}
	for i := 0; i < len(a); i++ {
		m[a[i]] = struct{}{}
	}
	// Transform all coordinates in "b" by vector "v" and add to map
	for i := 0; i < len(b); i++ {
		b[i].Transform(v)
		m[b[i]] = struct{}{}

	}
	// Return a slice of the map keys
	out := []coordinates.Coord{}
	for k, _ := range m {
		out = append(out, k)
	}
	return out
}

// minMatchingBeacons is the minimum number of beacons that must match for two scanner regions to overlap
var minMatchingBeacons int = 12

// correlate checks if there exists a vector than when applied to "b" results in at least "minMatchingBeacons" beacons matching
// Returns transformation vector, slice of combined beacons and bool to represent if a match was found
func correlate(a, b []coordinates.Coord) (coordinates.Coord, []coordinates.Coord, bool) {
	// Build map of all possible transformation vectors from "b" to "a"
	vectors := map[coordinates.Coord]int{}
	for _, aa := range a {
		for _, bb := range b {
			vector := coordinates.Coord{
				X: aa.X - bb.X,
				Y: aa.Y - bb.Y,
				Z: aa.Z - bb.Z,
			}
			vectors[vector] += 1
		}
	}

	// Check if a single transformation vector results in at least "minMatchingBeacons" overlapping beacons
	for vector, count := range vectors {
		if count >= minMatchingBeacons {
			c := combine(a, b, vector)
			return vector, c, true
		}
	}
	return coordinates.Coord{}, a, false
}

// doRotateX rotates a slice of coordinates 90 degrees around the x axis
func doRotateX(a []coordinates.Coord) {
	for i := 0; i < len(a); i++ {
		a[i].RotateX90()
	}
}

// doRotateY rotates a slice of coordinates 90 degrees around the y axis
func doRotateY(a []coordinates.Coord) {
	for i := 0; i < len(a); i++ {
		a[i].RotateY90()
	}
}

// doRotateZ rotates a slice of coordinates 90 degrees around the z axis
func doRotateZ(a []coordinates.Coord) {
	for i := 0; i < len(a); i++ {
		a[i].RotateZ90()
	}
}

// doCorrelate attempts to match two scanner regions by trying all possible rotations of the second region
// Returns transformation vector, slice of combined beacons and bool to represent if a match was found
func doCorrelate(a, b []coordinates.Coord) (coordinates.Coord, []coordinates.Coord, bool) {
	for rotX := 0; rotX < 4; rotX++ {
		for rotY := 0; rotY < 4; rotY++ {
			for rotZ := 0; rotZ < 4; rotZ++ {
				v, points, match := correlate(a, b)
				if match {
					return v, points, true
				}
				doRotateZ(b)
			}
			doRotateY(b)
		}
		doRotateX(b)
	}

	return coordinates.Coord{}, a, false
}

func part1(reports [][]coordinates.Coord) (int, []coordinates.Coord) {
	points := reports[0]
	var b []coordinates.Coord
	// Vectors to hold positions of scanners relative to first scanner at origin
	vectors := []coordinates.Coord{coordinates.Coord{}}

	// Create a queue to hold scanner data
	queue := list.New()
	for _, r := range reports[1:] {
		queue.PushBack(r)
	}

	// Not all scanner regions will match but they will form a complete region
	// Take scanner regions off queue and attempt to match to current region
	for queue.Len() > 0 {
		elem := queue.Front()
		queue.Remove(elem)
		b = elem.Value.([]coordinates.Coord)

		if v, newPoints, match := doCorrelate(points, b); match {
			points = newPoints
			vectors = append(vectors, v)
		} else {
			// Scanner data does not yet match, push to back of queue
			queue.PushBack(b)
		}
	}
	return len(points), vectors
}

func part2(vectors []coordinates.Coord) int {
	max := 0
	for i := 0; i < len(vectors)-1; i++ {
		for j := i + 1; j < len(vectors); j++ {
			d := coordinates.ManhattanDistance(vectors[i], vectors[j])
			if d > max {
				max = d
			}
		}
	}
	return max
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	reports := parseData(data)

	p1, vectors := part1(reports)
	fmt.Println("Part 1:", p1)

	p2 := part2(vectors)
	fmt.Println("Part 2:", p2)
}
