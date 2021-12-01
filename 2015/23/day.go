// Advent of Code 2015 - Day 233
package day23

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// Instruction holds the information for a single instruction
type Instruction struct {
	name, register string
	offset         int
}

// parseInstructions converts the input lines in to a slice of Instruction objects
func parseInstructions(lines []string) []Instruction {
	instructions := make([]Instruction, len(lines))

	for i, line := range lines {
		parts := strings.Split(line, " ")
		instructions[i].name = parts[0]
		instructions[i].register = strings.TrimSuffix(parts[1], ",")
		if len(parts) == 3 {
			if offset, err := strconv.Atoi(parts[2]); err == nil {
				instructions[i].offset = offset
			}
			instructions[i].register = strings.TrimSuffix(parts[1], ",")
		} else {
			if parts[0] == "jmp" {
				if offset, err := strconv.Atoi(parts[1]); err == nil {
					instructions[i].offset = offset
				}
			} else {
				instructions[i].register = parts[1]
			}
		}
	}

	return instructions
}

// solve follows a set of instuctions to operate on the given registers
func solve(instructions []Instruction, registers map[string]uint) {
	for i := 0; i < len(instructions); {

		switch instructions[i].name {

		case "hlf":
			registers[instructions[i].register] /= 2
			i++

		case "tpl":
			registers[instructions[i].register] *= 3
			i++

		case "inc":
			registers[instructions[i].register] += 1
			i++

		case "jmp":
			i += instructions[i].offset

		case "jie":
			if registers[instructions[i].register]%2 == 0 {
				i += instructions[i].offset
			} else {
				i++
			}

		case "jio":
			if registers[instructions[i].register] == 1 {
				i += instructions[i].offset
			} else {
				i++
			}

		default:
			log.Fatal("Unrecognised instruction")

		}

	}
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(
		strings.TrimSuffix(string(data), "\n"), "\n",
	)
	instructions := parseInstructions(lines)

	registers := map[string]uint{
		"a": 0,
		"b": 0,
	}
	solve(instructions, registers)
	fmt.Println("Part 1: Register b has value", registers["b"])

	registers["a"] = 1
	registers["b"] = 0
	solve(instructions, registers)
	fmt.Println("Part 2: Register b has value", registers["b"])
}
