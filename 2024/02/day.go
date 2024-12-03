// Advent of Code 2024 - Day 2.
package day2

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

func parseData(data []byte) ([][]int, error) {
	lines := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
	reports := make([][]int, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " ")
		levels := make([]int, len(parts))
		for j, part := range parts {
			n, err := strconv.Atoi(part)
			if err != nil {
				return nil, err
			}
			levels[j] = n
		}
		reports[i] = levels
	}
	return reports, nil
}

func part1(reports [][]int) int {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var count int
	for _, report := range reports {
		r := report
		wg.Add(1)

		go func() {
			defer wg.Done()
			if isSafe(r) {
				mu.Lock()
				count += 1
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	return count
}

func part2(reports [][]int) int {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var count int
	for _, report := range reports {
		r := report
		wg.Add(1)

		go func() {
			defer wg.Done()
			if isSafeWithDampener(r) {
				mu.Lock()
				count += 1
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	return count
}

const (
	levelChangeUnknown = iota
	levelChangeIncreasing
	levelChangeDecreasing
)

func isSafe(levels []int) bool {
	change := levelChangeUnknown
	for i := 0; i < len(levels)-1; i++ {
		diff := levels[i+1] - levels[i]
		switch {
		case diff > 0 && diff <= 3:
			if i == 0 {
				change = levelChangeIncreasing
			}
			if change != levelChangeIncreasing {
				return false
			}
		case diff < 0 && diff >= -3:
			if i == 0 {
				change = levelChangeDecreasing
			}
			if change != levelChangeDecreasing {
				return false
			}
		default:
			return false
		}
	}
	return true
}

func isSafeWithDampener(levels []int) bool {
	if isSafe(levels) {
		return true
	}
	for i := 0; i < len(levels); i++ {
		dampenedLevels := make([]int, len(levels))
		copy(dampenedLevels, levels)
		dampenedLevels = append(dampenedLevels[:i], dampenedLevels[i+1:]...)
		if isSafe(dampenedLevels) {
			return true
		}
	}
	return false
}

func Run(inputFile string) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	reports, err := parseData(data)
	if err != nil {
		log.Fatal(err)
	}

	p1 := part1(reports)
	fmt.Println("Part 1:", p1)

	p2 := part2(reports)
	fmt.Println("Part 2:", p2)
}
