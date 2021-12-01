// Advent of Code 2015 - Day 6
package day6

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var instructionRegexp = regexp.MustCompile(`(turn on|turn off|toggle) ([0-9]+),([0-9]+) through ([0-9]+),([0-9]+)`)

type Lights [1000][1000]int

// turnOn switches on lights
func (lights *Lights) turnOn(x1, y1, x2, y2 int) {
	for ; x1 <= x2; x1++ {
		for ; y1 <= y2; y1++ {
			lights[x1][y1] = 1
		}
	}
}

// turnOff switches off lights
func (lights *Lights) turnOff(x1, y1, x2, y2 int) {
	for ; x1 <= x2; x1++ {
		for ; y1 <= y2; y1++ {
			lights[x1][y1] = 0
		}
	}
}

// toggle switches the state of lights
func (lights *Lights) toggle(x1, y1, x2, y2 int) {
	for ; x1 <= x2; x1++ {
		for ; y1 <= y2; y1++ {
			if lights[x1][y1] == 0 {
				lights[x1][y1] = 1
			} else {
				lights[x1][y1] = 0
			}
		}
	}
}

// totalLightsOn determines the brightness of all the lights
func (lights *Lights) totalLightsOn() int {
	brightness := 0
	for x := range lights {
		for y := range lights[x] {
			brightness += lights[x][y]
		}
	}
	return brightness
}

// updateBrightness changes the brightness of the lights for part 2
func (lights *Lights) updateBrightness(action string, x1, y1, x2, y2 int) {
	for ; x1 <= x2; x1++ {
		for ; y1 <= y2; y1++ {
			switch action {
			case "turn on":
				lights[x1][y1]++
			case "turn off":
				if lights[x1][y1] > 0 {
					lights[x1][y1]--
				}
			case "toggle":
				lights[x1][y1] += 2
			}
		}
	}
}

// Instruction represents each command for updating lights
type Instruction struct {
	name           string
	x1, y1, x2, y2 int
}

// parseInstruction converts an instruction string in to an Instruction struct
func parseInstruction(line string) Instruction {
	matches := instructionRegexp.FindStringSubmatch(line)
	x1, _ := strconv.Atoi(matches[2])
	y1, _ := strconv.Atoi(matches[3])
	x2, _ := strconv.Atoi(matches[4])
	y2, _ := strconv.Atoi(matches[5])

	return Instruction{
		name: matches[1],
		x1:   x1,
		y1:   y1,
		x2:   x2,
		y2:   y2,
	}
	return Instruction{}
}

func part1(lights Lights, instructions []Instruction) {
	for _, instruction := range instructions {
		switch instruction.name {
		case "turn on":
			lights.turnOn(instruction.x1, instruction.y1, instruction.x2, instruction.y2)
		case "turn off":
			lights.turnOff(instruction.x1, instruction.y1, instruction.x2, instruction.y2)
		case "toggle":
			lights.toggle(instruction.x1, instruction.y1, instruction.x2, instruction.y2)
		}
	}
	fmt.Println(lights.totalLightsOn())
}

func part2(lights Lights, instructions []Instruction) {
	for _, instruction := range instructions {
		lights.updateBrightness(instruction.name, instruction.x1, instruction.y1, instruction.x2, instruction.y2)
	}
	fmt.Println(lights.totalLightsOn())
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
	instructions := []Instruction{}

	for _, line := range lines {
		instructions = append(instructions, parseInstruction(line))
	}

	lights := Lights{}

	part1(lights, instructions)
	part2(lights, instructions)
}
