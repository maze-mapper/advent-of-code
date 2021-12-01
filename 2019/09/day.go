// Advent of Code 2019 - Day 9
package day9

import (
	"fmt"
	"io/ioutil"
	"log"

	"adventofcode/2019/intcode"
)

// runProgram runs an Intcode computer with a single input and returns the single output
func runProgram(program []int, input int) int {
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
	if len(outputs) != 1 {
		log.Fatalf("More numbers outputed than expected: %v", outputs)
	}
	return outputs[0]
}

func part1(program []int) int {
	return runProgram(program, 1)
}

func part2(program []int) int {
	return runProgram(program, 2)
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
