// Advent of Code 2024 - Day 17.
package day17

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseData(data []byte) ([3]int, []int, error) {
	sections := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n\n")
	var a, b, c int
	n, err := fmt.Sscanf(sections[0], "Register A: %d\nRegister B: %d\nRegister C: %d", &a, &b, &c)
	if err != nil {
		return [3]int{}, nil, err
	}
	if n != 3 {
		return [3]int{}, nil, fmt.Errorf("expected to parse 3 items, instead parsed %d", n)
	}
	registers := [3]int{a, b, c}

	var program []int
	for _, s := range strings.Split(strings.TrimPrefix(sections[1], "Program: "), ",") {
		n, err := strconv.Atoi(s)
		if err != nil {
			return [3]int{}, nil, err
		}
		program = append(program, n)
	}

	return registers, program, nil
}

func part1(registers [3]int, program []int) string {
	out, err := runProgram(registers, program)
	if err != nil {
		log.Fatal(err)
	}
	return out
}

func part2(registers [3]int, program []int) int {
	a, _ := findInitialARegister(program, len(program)-1, 0, 8)
	return a
}

// runProgram returns the program output for the given initial register values.
func runProgram(registers [3]int, program []int) (string, error) {
	var output []string
	for idx := 0; idx < len(program)-1; idx += 2 {
		operator := program[idx]
		operand := program[idx+1]

		switch operator {

		case 0: // adv
			operand, err := combo(registers, operand)
			if err != nil {
				return "", err
			}
			denominator := 1
			for i := 1; i <= operand; i++ {
				denominator *= 2
			}
			registers[0] /= denominator

		case 1: // bxl
			registers[1] ^= operand

		case 2: // bst
			operand, err := combo(registers, operand)
			if err != nil {
				return "", err
			}
			registers[1] = operand % 8

		case 3: // jnz
			if registers[0] != 0 {
				idx = operand
				idx -= 2 // So the usual +2 doesn't happen.
			}

		case 4: // bxc
			registers[1] = registers[1] ^ registers[2]

		case 5: // out
			operand, err := combo(registers, operand)
			if err != nil {
				return "", err
			}
			output = append(output, strconv.Itoa(operand%8))

		case 6: // bdv
			operand, err := combo(registers, operand)
			if err != nil {
				return "", err
			}
			denominator := 1
			for i := 1; i <= operand; i++ {
				denominator *= 2
			}
			registers[1] = registers[0] / denominator

		case 7: // cdv
			operand, err := combo(registers, operand)
			if err != nil {
				return "", err
			}
			denominator := 1
			for i := 1; i <= operand; i++ {
				denominator *= 2
			}
			registers[2] = registers[0] / denominator
		}
	}
	return strings.Join(output, ","), nil
}

// findInitialARegister works backwards through every value output by the
// program to find the initial value of the A register such that the output is
// the program itself.
// Lower is the inclusive minimum value for the A register.
// Upper is the exclusive maximum value for the A register.
func findInitialARegister(program []int, i, lower, upper int) (int, bool) {
	want := program[i]
	for a := lower; a < upper; a++ {
		if reverseARegister(a) == want {
			if i == 0 {
				return a, true
			}
			ans, found := findInitialARegister(program, i-1, a*8, (a+1)*8)
			if found {
				return ans, true
			}
		}
	}
	return 0, false
}

// reverseARegister undoes the effect of one loop of the program on the A register value.
// This function is dependent on the program and will not work for all inputs.
func reverseARegister(a int) int {
	k := (a % 8) ^ 5
	denominator := 1
	for i := 1; i <= k; i++ {
		denominator *= 2
	}
	return (k ^ 6 ^ (a / denominator)) % 8
}

func combo(registers [3]int, operand int) (int, error) {
	switch operand {
	case 0, 1, 2, 3:
		return operand, nil
	case 4:
		return registers[0], nil
	case 5:
		return registers[1], nil
	case 6:
		return registers[2], nil
	default:
		return 0, fmt.Errorf("operand %d is not valid", operand)
	}
}

func Run(inputFile string) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	registers, program, err := parseData(data)
	if err != nil {
		log.Fatal(err)
	}

	p1 := part1(registers, program)
	fmt.Println("Part 1:", p1)

	p2 := part2(registers, program)
	fmt.Println("Part 2:", p2)
}
