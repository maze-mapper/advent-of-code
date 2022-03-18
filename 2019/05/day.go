// Advent of Code 2019 - Day 5
package day5

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/maze-mapper/advent-of-code/2019/intcode"
)

// runProgram runs an Intcode computer with a single input and returns all outputs
func runProgram(program []int, input int) []int {
	computer := intcode.New(program)
	chanIn := make(chan int)
	chanOut := make(chan int)
	computer.SetChanIn(chanIn)
	computer.SetChanOut(chanOut)

	go computer.Run()
	chanIn <- input

	outputs := []int{}
	for i := range chanOut {
		outputs = append(outputs, i)
	}
	return outputs
}

func part1(program []int) int {
	outputs := runProgram(program, 1)

	// Check all outputs other than the last are zero then return the diagnostic code
	lastIndex := len(outputs) - 1
	for _, output := range outputs[:lastIndex] {
		if output != 0 {
			log.Fatal("Unexpected non-zero output")
		}
	}

	return outputs[lastIndex]
}

func part2(program []int) int {
	outputs := runProgram(program, 5)

	if len(outputs) != 1 {
		log.Fatalf("More numbers outputed than expected: %v", outputs)
	}

	return outputs[0]
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	program := intcode.ReadProgram(data)

	p1 := part1(program)
	fmt.Println("Part 1:", p1)

	p2 := part2(program)
	fmt.Println("Part 2:", p2)
}
