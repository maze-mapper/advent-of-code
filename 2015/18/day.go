// Advent of Code 2015 - Day 188
package day18

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type Lights [100][100]int

// onNeighbours counts the number of on neighbours for a particular light
func (lights *Lights) onNeighbours(x, y int) int {
	count := 0
	for i := x - 1; i < x+2; i++ {
		for j := y - 1; j < y+2; j++ {
			if i >= 0 && i < len(lights) && j >= 0 && j < len(lights[x]) {
				count += lights[i][j]
			}
		}
	}
	return count
}

// update switches the lights on/off based on the number of on neighbours
func (lights *Lights) update() *Lights {
	newLights := Lights{}
	for i := 0; i < len(lights); i++ {
		for j := 0; j < len(lights[i]); j++ {
			onNeighbours := lights.onNeighbours(i, j)
			if onNeighbours == 2 || onNeighbours == 3 {
				newLights[i][j] = 1
			}
		}
	}
	return &newLights
}

// totalLightsOn determines the number of on lights
func (lights *Lights) totalLightsOn() int {
	count := 0
	for x := range lights {
		for y := range lights[x] {
			count += lights[x][y]
		}
	}
	return count
}

// makeLights returns a Lights object with the given initial configuration
func makeLights(data []byte) Lights {
	lines := strings.Split(
		strings.TrimSuffix(string(data), "\n"), "\n",
	)

	lights := Lights{}
	for i, line := range lines {
		for j, char := range line {
			if char == '#' {
				lights[i][j] = 1
			}
		}
	}
	return lights
}

func part1(lights *Lights, steps int) {
	for s := 0; s < steps; s++ {
		lights = lights.update()
	}
	fmt.Println("Part 1: There are", lights.totalLightsOn(), "lights on after", steps, "steps")
}

func part2(lights *Lights, steps int) {
	iMax := len(lights) - 1
	jMax := len(lights[0]) - 1
	for s := 0; s < steps; s++ {
		lights = lights.update()
		// Fix the corners on
		lights[0][0] = 1
		lights[iMax][0] = 1
		lights[0][jMax] = 1
		lights[iMax][jMax] = 1
	}
	fmt.Println("Part 2: There are", lights.totalLightsOn(), "lights on after", steps, "steps")
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	steps := 100
	lights := makeLights(data)

	part1(&lights, steps)
	part2(&lights, steps)
}
