// Advent of Code 2023 - Day 9.
package day9

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseData(data []byte) [][]int {
	lines := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
	histories := make([][]int, len(lines))
	for i, line := range lines {
		points := strings.Split(line, " ")
		numbers := make([]int, len(points))
		for j, p := range points {
			n, err := strconv.Atoi(p)
			if err != nil {
				log.Fatal(err)
			}
			numbers[j] = n
		}
		histories[i] = numbers
	}
	return histories
}

func extrapolate(numbers []int, backwards bool) int {
	allZeros := true
	for _, n := range numbers {
		if n != 0 {
			allZeros = false
			break
		}
	}
	if allZeros {
		return 0
	}

	var diffs []int
	for i := 0; i < len(numbers)-1; i++ {
		d := numbers[i+1] - numbers[i]
		diffs = append(diffs, d)
	}
	if backwards {
		return numbers[0] - extrapolate(diffs, backwards)
	} else {
		return numbers[len(numbers)-1] + extrapolate(diffs, backwards)
	}
}

func part1(histories [][]int) int {
	sum := 0
	for _, history := range histories {
		sum += extrapolate(history, false)
	}
	return sum
}

func part2(histories [][]int) int {
	sum := 0
	for _, history := range histories {
		sum += extrapolate(history, true)
	}
	return sum
}

func Run(inputFile string) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	histories := parseData(data)

	p1 := part1(histories)
	fmt.Println("Part 1:", p1)

	p2 := part2(histories)
	fmt.Println("Part 2:", p2)
}
