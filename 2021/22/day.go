// Advent of Code 2021 - Day 22
package day22

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/maze-mapper/advent-of-code/coordinates"
)

type action struct {
	a, b  coordinates.Coord
	value bool
}

func parseData(data []byte) []action {
	lines := strings.Split(
		strings.TrimSuffix(string(data), "\n"), "\n",
	)

	steps := make([]action, len(lines))
	pattern := regexp.MustCompile(`(on|off) x=(-?[0-9]+)..(-?[0-9]+),y=(-?[0-9]+)..(-?[0-9]+),z=(-?[0-9]+)..(-?[0-9]+)`)
	for i, line := range lines {
		if match := pattern.FindStringSubmatch(line); match != nil {
			var value bool
			if match[1] == "on" {
				value = true
			} else {
				value = false
			}

			xMin, _ := strconv.Atoi(match[2])
			xMax, _ := strconv.Atoi(match[3])
			yMin, _ := strconv.Atoi(match[4])
			yMax, _ := strconv.Atoi(match[5])
			zMin, _ := strconv.Atoi(match[6])
			zMax, _ := strconv.Atoi(match[7])

			a := coordinates.Coord{
				X: xMin,
				Y: yMin,
				Z: zMin,
			}
			b := coordinates.Coord{
				X: xMax,
				Y: yMax,
				Z: zMax,
			}

			steps[i] = action{
				a:     a,
				b:     b,
				value: value,
			}

		} else {
			log.Fatal("Unable to match line", line)
		}
	}
	return steps
}

func countCubes(region [][][]bool) int {
	count := 0
	for i := 0; i < len(region); i++ {
		for j := 0; j < len(region[i]); j++ {
			for k := 0; k < len(region[i][j]); k++ {
				if region[i][j][k] {
					count += 1
				}
			}
		}
	}
	return count
}

func part1(steps []action) int {
	offset := 50

	// Initialise region
	size := offset*2 + 1
	region := make([][][]bool, size)
	for i := 0; i < len(region); i++ {
		region[i] = make([][]bool, size)
		for j := 0; j < len(region[i]); j++ {
			region[i][j] = make([]bool, size)
		}
	}

	for _, step := range steps {
		if step.a.X < -offset || step.a.Y < -offset || step.a.Z < -offset || step.b.X > offset || step.b.Y > offset || step.b.Z > offset {
			continue
		}
		for x := step.a.X + offset; x <= step.b.X+offset; x++ {
			for y := step.a.Y + offset; y <= step.b.Y+offset; y++ {
				for z := step.a.Z + offset; z <= step.b.Z+offset; z++ {
					region[x][y][z] = step.value
				}
			}
		}
	}

	return countCubes(region)
}

type Region struct {
	a, b coordinates.Coord
}

func (r *Region) Volume() int {
	var x, y, z int

	if r.b.X > r.a.X {
		x = r.b.X - r.a.X + 1
	} else {
		x = r.a.X - r.b.X + 1
	}
	if r.b.Y > r.a.Y {
		y = r.b.Y - r.a.Y + 1
	} else {
		y = r.a.Y - r.b.Y + 1
	}
	if r.b.Z > r.a.Z {
		z = r.b.Z - r.a.Z + 1
	} else {
		z = r.a.Z - r.b.Z + 1
	}

	return x * y * z
}

func Intersection(r, s Region) (Region, bool) {
	a := coordinates.Coord{}
	b := coordinates.Coord{}

	// Set "a" to the larger values of the two "a" points
	if r.a.X < s.a.X {
		a.X = s.a.X
	} else {
		a.X = r.a.X
	}
	if r.a.Y < s.a.Y {
		a.Y = s.a.Y
	} else {
		a.Y = r.a.Y
	}
	if r.a.Z < s.a.Z {
		a.Z = s.a.Z
	} else {
		a.Z = r.a.Z
	}

	// Set "b" to the smaller values of the two "b" points
	if r.b.X > s.b.X {
		b.X = s.b.X
	} else {
		b.X = r.b.X
	}
	if r.b.Y > s.b.Y {
		b.Y = s.b.Y
	} else {
		b.Y = r.b.Y
	}
	if r.b.Z > s.b.Z {
		b.Z = s.b.Z
	} else {
		b.Z = r.b.Z
	}

	// Check if there actually is an intersection
	valid := a.X <= b.X && a.Y <= b.Y && a.Z <= b.Z

	return Region{a: a, b: b}, valid
}

func part2(steps []action) int {
	doneActions := []action{}
	for _, step := range steps {
		changes := []action{}
		if step.value {
			changes = append(changes, step)
		}
		r := Region{a: step.a, b: step.b}
		for _, act := range doneActions {
			r2 := Region{a: act.a, b: act.b}
			if rr, ok := Intersection(r, r2); ok {
				changes = append(changes, action{value: !act.value, a: rr.a, b: rr.b})
			}
		}
		doneActions = append(doneActions, changes...)
	}

	onCubes := 0
	for _, act := range doneActions {
		r := Region{a: act.a, b: act.b}
		volume := r.Volume()
		if act.value {
			onCubes += volume
		} else {
			onCubes -= volume
		}
	}
	return onCubes
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	steps := parseData(data)

	p1 := part1(steps)
	fmt.Println("Part 1:", p1)

	p2 := part2(steps)
	fmt.Println("Part 2:", p2)
}
