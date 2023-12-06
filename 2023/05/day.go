// Advent of Code 2023 - Day 5.
package day5

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type almanac struct {
	seeds []int
	// mappings []map[int]int
	mappings [][]mappingFunc
}

type mappingFunc func(n int) (int, bool)

func makeMapping(dst, src, rng int) mappingFunc {
	return func(n int) (int, bool) {
		if n < src || n >= src+rng {
			return 0, false
		}
		offset := n - src
		return dst + offset, true
	}
}

func parseData(data []byte) almanac {
	sections := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n\n")

	a := almanac{}

	seedLine := strings.TrimPrefix(sections[0], "seeds: ")
	for _, s := range strings.Split(seedLine, " ") {
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		a.seeds = append(a.seeds, n)
	}

	for _, section := range sections[1:] {
		lines := strings.Split(section, "\n")
		var m []mappingFunc
		for _, line := range lines[1:] {
			var dst, src, rng int
			fmt.Sscanf(line, "%d %d %d", &dst, &src, &rng)
			f := makeMapping(dst, src, rng)
			m = append(m, f)
		}
		a.mappings = append(a.mappings, m)
	}

	return a
}

func part1(a almanac) int {
	minLocation := math.MaxInt
	for _, input := range a.seeds {
		for _, m := range a.mappings {
			for _, f := range m {
				if output, ok := f(input); ok {
					input = output
					break
				}
			}
		}
		minLocation = min(minLocation, input)
	}
	return minLocation
}

func part2(a almanac) int {
	minLocation := math.MaxInt
	for i := 0; i < len(a.seeds); i += 2 {
		seedStart := a.seeds[i]
		seedRange := a.seeds[i+1]

		for seed := seedStart; seed < seedStart+seedRange; seed++ {
			input := seed
			for _, m := range a.mappings {
				for _, f := range m {
					if output, ok := f(input); ok {
						input = output
						break
					}
				}
			}
			// fmt.Println("Seed:", seed, "location:", input)
			minLocation = min(minLocation, input)
		}
	}
	return minLocation
}

func Run(inputFile string) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	a := parseData(data)

	p1 := part1(a)
	fmt.Println("Part 1:", p1)

	p2 := part2(a)
	fmt.Println("Part 2:", p2)
}
