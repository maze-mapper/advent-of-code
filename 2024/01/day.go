// Advent of Code 2024 - Day 1.
package day1

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func parseData(data []byte) ([]int, []int, error) {
	lines := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
	a := make([]int, len(lines))
	b := make([]int, len(lines))
	for i, line := range lines {
		before, after, found := strings.Cut(line, "   ")
		if !found {
			return nil, nil, fmt.Errorf("unable to parse line %q", line)
		}
		n1, err := strconv.Atoi(before)
		if err != nil {
			return nil, nil, err
		}
		a[i] = n1
		n2, err := strconv.Atoi(after)
		if err != nil {
			return nil, nil, err
		}
		b[i] = n2
	}
	return a, b, nil
}

func part1(a, b []int) int {
	if len(a) != len(b) {
		log.Fatal("slices have different lengths")
	}
	slices.Sort(a)
	slices.Sort(b)

	var d int
	for i := range a {
		aVal := a[i]
		bVal := b[i]
		if aVal > bVal {
			d += aVal - bVal
		} else {
			d += bVal - aVal
		}
	}
	return d
}

func part2(a, b []int) int {
	f := frequency(b)
	var s int
	for _, e := range a {
		v := f[e]
		s += e * v
	}
	return s
}

func frequency(a []int) map[int]int {
	m := map[int]int{}
	for _, e := range a {
		m[e] += 1
	}
	return m
}

func Run(inputFile string) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	a, b, err := parseData(data)
	if err != nil {
		log.Fatal(err)
	}

	p1 := part1(a, b)
	fmt.Println("Part 1:", p1)

	p2 := part2(a, b)
	fmt.Println("Part 2:", p2)
}
