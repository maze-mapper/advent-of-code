// Advent of Code 2023 - Day 3.
package day3

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1(lines []string) int {
	numbers, _ := partNumbers(lines)
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func isNumber(b byte) bool {
	switch b {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return true
	default:
		return false
	}
}

func isEnginePart(b byte) bool {
	if b == '.' || isNumber(b) {
		return false
	}
	return true
}

func isGear(b byte) bool {
	return b == '*'
}

func partNumbers(lines []string) ([]int, map[[2]int][]int) {
	var numbers []int
	allGears := map[[2]int][]int{}
	for i, line := range lines {
		var s string
		var isPart bool
		gears := map[[2]int]bool{}
		for j, r := range line {
			switch r {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				s += string(r)

				if (j > 0 && isEnginePart(line[j-1])) ||
					(j < len(line)-1 && isEnginePart(line[j+1])) ||
					(i > 0 && isEnginePart(lines[i-1][j])) ||
					(i < len(lines)-1 && isEnginePart(lines[i+1][j])) ||
					(i > 0 && j > 0 && isEnginePart(lines[i-1][j-1])) ||
					(i > 0 && j < len(line)-1 && isEnginePart(lines[i-1][j+1])) ||
					(i < len(lines)-1 && j > 0 && isEnginePart(lines[i+1][j-1])) ||
					(i < len(lines)-1 && j < len(line)-1 && isEnginePart(lines[i+1][j+1])) {
					isPart = true
				}

				if j > 0 && isGear(line[j-1]) {
					gears[[2]int{i, j - 1}] = true
				}
				if j < len(line)-1 && isGear(line[j+1]) {
					gears[[2]int{i, j + 1}] = true
				}
				if i > 0 && isGear(lines[i-1][j]) {
					gears[[2]int{i - 1, j}] = true
				}
				if i < len(lines)-1 && isGear(lines[i+1][j]) {
					gears[[2]int{i + 1, j}] = true
				}
				if i > 0 && j > 0 && isGear(lines[i-1][j-1]) {
					gears[[2]int{i - 1, j - 1}] = true
				}
				if i > 0 && j < len(line)-1 && isGear(lines[i-1][j+1]) {
					gears[[2]int{i - 1, j + 1}] = true
				}
				if i < len(lines)-1 && j > 0 && isGear(lines[i+1][j-1]) {
					gears[[2]int{i + 1, j - 1}] = true
				}
				if i < len(lines)-1 && j < len(line)-1 && isGear(lines[i+1][j+1]) {
					gears[[2]int{i + 1, j + 1}] = true
				}

			default:
				if isPart {
					n, err := strconv.Atoi(s)
					if err != nil {
						log.Fatal(err)
					}
					numbers = append(numbers, n)

					for gear := range gears {
						if _, ok := allGears[gear]; !ok {
							allGears[gear] = nil
						}
						allGears[gear] = append(allGears[gear], n)
					}
				}
				s = ""
				isPart = false
				gears = map[[2]int]bool{}
			}

			if j == len(line)-1 && isPart {
				n, err := strconv.Atoi(s)
				if err != nil {
					log.Fatal(err)
				}
				numbers = append(numbers, n)

				for gear := range gears {
					if _, ok := allGears[gear]; !ok {
						allGears[gear] = nil
					}
					allGears[gear] = append(allGears[gear], n)
				}
			}
		}
	}
	return numbers, allGears
}

func part2(lines []string) int {
	_, gears := partNumbers(lines)
	sum := 0
	for _, vals := range gears {
		if len(vals) == 2 {
			sum += vals[0] * vals[1]
		}
	}
	return sum
}

func Run(inputFile string) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")

	p1 := part1(lines)
	fmt.Println("Part 1:", p1)

	p2 := part2(lines)
	fmt.Println("Part 2:", p2)
}
