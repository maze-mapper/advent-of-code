// Advent of Code 2024 - Day 7.
package day7

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type equation struct {
	value   int
	numbers []int
}

func parseData(data []byte) ([]equation, error) {
	lines := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
	equations := make([]equation, len(lines))
	for i, line := range lines {
		before, after, found := strings.Cut(line, ": ")
		if !found {
			return nil, fmt.Errorf("unable to parse line %q", line)
		}

		value, err := strconv.Atoi(before)
		if err != nil {
			return nil, err
		}
		e := equation{value: value}

		for _, s := range strings.Split(after, " ") {
			n, err := strconv.Atoi(s)
			if err != nil {
				return nil, err
			}
			e.numbers = append(e.numbers, n)
		}

		equations[i] = e
	}
	return equations, nil
}

func part1(equations []equation) int {
	return calibrationResult(equations, []opFunc{add, mul})
}

func part2(equations []equation) int {
	return calibrationResult(equations, []opFunc{add, mul, concat})
}

func calibrationResult(equations []equation, opFuncs []opFunc) int {
	var result int
	for _, eq := range equations {
		if isSolvable(eq.numbers, eq.value, 0, eq.numbers[0], opFuncs) {
			result += eq.value
		}
	}
	return result
}

func isSolvable(numbers []int, target, idx, current int, opFuncs []opFunc) bool {
	if idx == len(numbers)-1 {
		return target == current
	}
	idx++
	n := numbers[idx]
	for _, f := range opFuncs {
		if isSolvable(numbers, target, idx, f(current, n), opFuncs) {
			return true
		}
	}
	return false
}

type opFunc func(a, b int) int

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

func concat(a, b int) int {
	shift := 10 // There will always be at least one digit, even if b is 0.
	for shift <= b {
		shift *= 10
	}
	return a*shift + b
}

func Run(inputFile string) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	equations, err := parseData(data)
	if err != nil {
		log.Fatal(err)
	}

	p1 := part1(equations)
	fmt.Println("Part 1:", p1)

	p2 := part2(equations)
	fmt.Println("Part 2:", p2)
}
