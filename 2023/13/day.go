// Advent of Code 2023 - Day 13.
package day13

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func parseData(data []byte) [][]string {
	chunks := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n\n")
	output := make([][]string, len(chunks))
	for i, chunk := range chunks {
		lines := strings.Split(chunk, "\n")
		output[i] = lines
	}
	return output
}

func findHorizontalMirror(lines []string) (int, bool) {
loop:
	for i := 0; i < len(lines)-1; i++ {
		j := i
		k := i + 1
		for j >= 0 && k < len(lines) {
			if lines[j] != lines[k] {
				continue loop
			}
			j -= 1
			k += 1
		}
		return i + 1, true
	}
	return 0, false
}

func findVerticalMirror(lines []string) (int, bool) {
loop:
	for i := 0; i < len(lines[0])-1; i++ {
		j := i
		k := i + 1
		for j >= 0 && k < len(lines[0]) {
			for row := range lines {
				if lines[row][j] != lines[row][k] {
					continue loop
				}
			}
			j -= 1
			k += 1
		}
		return i + 1, true
	}
	return 0, false
}

func findHorizontalMirrorWithSmudge(lines []string) (int, bool) {
	for i := 0; i < len(lines)-1; i++ {
		j := i
		k := i + 1
		smudges := 0
		for j >= 0 && k < len(lines) && smudges <= 1 {
			for col := 0; col < len(lines[0]); col++ {
				if lines[j][col] != lines[k][col] {
					smudges += 1
				}
			}
			j -= 1
			k += 1
		}
		if smudges == 1 {
			return i + 1, true
		}
	}
	return 0, false
}

func findVerticalMirrorWithSmudge(lines []string) (int, bool) {
	for i := 0; i < len(lines[0])-1; i++ {
		j := i
		k := i + 1
		smudges := 0
		for j >= 0 && k < len(lines[0]) && smudges <= 1 {
			for row := range lines {
				if lines[row][j] != lines[row][k] {
					smudges += 1
				}
			}
			j -= 1
			k += 1
		}
		if smudges == 1 {
			return i + 1, true
		}
	}
	return 0, false
}

func part1(input [][]string) int {
	sum := 0
	for _, lines := range input {
		if n, ok := findHorizontalMirror(lines); ok {
			sum += 100 * n
		}
		if n, ok := findVerticalMirror(lines); ok {
			sum += n
		}
	}
	return sum
}

func part2(input [][]string) int {
	sum := 0
	for _, lines := range input {
		if n, ok := findHorizontalMirrorWithSmudge(lines); ok {
			sum += 100 * n
		}
		if n, ok := findVerticalMirrorWithSmudge(lines); ok {
			sum += n
		}
	}
	return sum
}

func Run(inputFile string) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	input := parseData(data)

	p1 := part1(input)
	fmt.Println("Part 1:", p1)

	p2 := part2(input)
	fmt.Println("Part 2:", p2)
}
