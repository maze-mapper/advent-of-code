// Advent of Code 2019 - Day 21
package day21

import (
	"github.com/maze-mapper/advent-of-code/2019/intcode"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// runProgram runs the given program with ASCII input and returns the output
func runProgram(program []int, input []byte) []int {
	computer := intcode.New(program)
	chanIn := make(chan int)
	chanOut := make(chan int)
	computer.SetChanIn(chanIn)
	computer.SetChanOut(chanOut)

	go computer.Run()

	go func() {
		for _, in := range input {
			chanIn <- int(in)
		}
	}()

	output := []int{}
	for out := range chanOut {
		output = append(output, out)
	}

	return output
}

// makeInput converts a slice of strings in to a byte slice
func makeInput(parts []string) []byte {
	if len(parts) > 15 {
		log.Fatal("Too many springscript instructions")
	}

	var builder strings.Builder
	for _, part := range parts {
		builder.WriteString(part)
		builder.WriteString("\n")
	}

	return []byte(builder.String())
}

// printOutput will print the output from the Intcode computer
func printOutput(output []int) {
	b := make([]byte, len(output))
	for i, o := range output {
		// Unsafe if o overflows byte
		b[i] = byte(o)
	}
	fmt.Println(string(b))
}

// runSpringdroid runs the springscript instructions and returns the hull damage
func runSpringdroid(program []int, instructions []string) int {
	input := makeInput(instructions)
	output := runProgram(program, input)

	var damage int
	lastIndex := len(output) - 1
	if output[lastIndex] > 256 {
		damage = output[lastIndex]
		output = output[:lastIndex]
	} else {
		printOutput(output)
	}
	return damage
}

func part1(program []int) int {
	// Jump if D is ground and any of A, B or C are holes
	// D & (!A | !B | !C)
	// D & !(A & B & C)
	instructions := []string{
		"OR A J",  // J = A (since J starts as false)
		"AND B J", // J = A & B
		"AND C J", // J = A & B & C
		"NOT J J", // J = !(A & B & C)
		"AND D J", // J = D & !(A & B & C)
		"WALK",
	}
	return runSpringdroid(program, instructions)
}

func part2(program []int) int {
	// Jump if D is ground and either A is a hole or B is a hole or C is a hole and H is ground
	// D & (!A | !B | (!C & H))
	// D & (!(A & B) | (!C & H))
	instructions := []string{
		"OR A J",  // J = A (since J starts as false)
		"AND B J", // J = A & B
		"NOT J J", // J = !(A & B)
		"NOT C T", // T = !C
		"AND H T", // T = !C & H
		"OR T J",  // J = !(A & B) | (!C & H)
		"AND D J", // J = D & (!(A & B) | (!C & H))
		"RUN",
	}
	return runSpringdroid(program, instructions)
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
