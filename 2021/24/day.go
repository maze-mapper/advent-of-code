package day24

import (
	"fmt"
	//	"strings"
)

/*
type Instruction struct {
	name, a, b string
}

func parseData(data []byte) []Instruction {
	lines := strings.Split(
		strings.TrimSuffix(string(data), "\n"), "\n",
	)

	instructions := make([]Instruction, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " ")
		instructions[i] = Instruction{
			name: parts[0],
			a: parts[1],
		}
		if parts[0] != "inp" {
			instructions[i].b = parts[2]
		}
	}
	return instructions
}

// getDigits extracts the digits from a number
func getDigits(n int, digits *[]int) {
        if n < 10 {
                *digits = append(*digits, n)
        } else {
                getDigits(n / 10, digits)
                *digits = append(*digits, n % 10)
        }
}

// modelNumberToSlice converts an integer in to a slice of ints and also returns a bool to indicate if it is valid
func modelNumberToSlice(model int) ([]int, bool) {
	digits := []int{}
	getDigits(model, &digits)
	for _, v := range digits {
		if v == 0 {
			return digits, false
		}
	}
	return digits, true
}

// recurse is a brute force approach that skips zeros
func recurse(z, depth int, numbers []int) (int, bool) {
	fmt.Println(z, depth, numbers)
	if depth >= 14 {
		if z == 0 {
			fmt.Println(numbers)
			return true
		}
		return z, false
	}
	for i := 9; i > 0; i-- {
		newZ := f(i, z, nums[depth][0], nums[depth][1], nums[depth][2])
		newNumbers := make([]int, len(numbers))
		copy(newNumbers, numbers)
		newNumbers = append(newNumbers, i)
		recurse(newZ, depth + 1, newNumbers)
	}
	return z, false
}

func part12() {
	recurse(0, 0, []int{})
}

func part1() int {
	for model := 99999999999999; model >= 11111111111111; model-- {
		if digits, ok := modelNumberToSlice(model); ok {
			fmt.Println(model)
			z := 0
			for i, d := range digits {
				z = f(d, z, nums[i][0], nums[i][1], nums[i][2])
			}
			fmt.Println(z)
			if z == 0 {
				return model
			}
		}
	}
	return 0
}*/

// Numbers extracted from input data
// Correspond to values (a, b, c) in below function
var nums = [][3]int{
	{1, 12, 15},
	{1, 14, 12},
	{1, 11, 15},
	{26, -9, 12},
	{26, -7, 15},
	{1, 11, 2},
	{26, -1, 11},
	{26, -16, 15},
	{1, 11, 10},
	{26, -15, 2},
	{1, 10, 0},
	{1, 12, 0},
	{26, -4, 15},
	{26, 0, 15},
}

// f is the reverse engineered function for the ALU program
//   w is the input digit
//   z is a base 26 stack
//   a determines if we push (a=1) or pop (a=26) from the stack
// If pushing we store w + c on the stack
// To pop we require that the last item on the stack + b = w
// This gives a relationship between the pushes and pops that can be used to infer the differences between digits pushed and popped.
// Example using above data:
//   Consider the pair {1, 11, 15} and {26, -9, 12}
//   First pushes w_1 + 15 on to the stack
//   Second pops if w_2 = last item on stack - 9
//   So w_2 = w_1 + 15 - 9
//      w_2 = w_1 + 6
func f(w, z, a, b, c int) int {
	// Pop from stack z and store result in x with b added to it if a = 26
	x := (z % 26) + b
	z /= a

	if x != w {
		// Push w + c to stack z
		z *= 26
		z += w + c
	}

	return z
}

func Run(inputFile string) {
	fmt.Println("Solved manually for my specific input data")
	fmt.Println("Part 1:", 94399898949959)
	fmt.Println("Part 2:", 21176121611511)
}
