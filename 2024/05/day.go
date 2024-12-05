// Advent of Code 2024 - Day 5.
package day5

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func parseData(data []byte) ([][2]int, [][]int, error) {
	sections := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n\n")
	if len(sections) != 2 {
		log.Fatalf("expecting 2 sections, got %d", len(sections))
	}

	ruleSectionLines := strings.Split(sections[0], "\n")
	rules := make([][2]int, len(ruleSectionLines))
	for i, line := range ruleSectionLines {
		before, after, found := strings.Cut(line, "|")
		if !found {
			return nil, nil, fmt.Errorf("unable to parse line %q", line)
		}
		n1, err := strconv.Atoi(before)
		if err != nil {
			return nil, nil, err
		}
		n2, err := strconv.Atoi(after)
		if err != nil {
			return nil, nil, err
		}
		rules[i] = [2]int{n1, n2}
	}

	pageNumberSectionLines := strings.Split(sections[1], "\n")
	pageNumbers := make([][]int, len(pageNumberSectionLines))
	for i, line := range pageNumberSectionLines {
		pages := strings.Split(line, ",")
		pageNumbers[i] = make([]int, len(pages))
		for j, s := range pages {
			n, err := strconv.Atoi(s)
			if err != nil {
				return nil, nil, err
			}
			pageNumbers[i][j] = n
		}
	}

	return rules, pageNumbers, nil
}

type dependencyGraph map[int]map[int]bool

func buildDependencyGraph(rules [][2]int) dependencyGraph {
	m := dependencyGraph{}
	for _, r := range rules {
		a := r[0]
		b := r[1]
		if _, ok := m[b]; !ok {
			m[b] = map[int]bool{}
		}
		m[b][a] = true
	}
	return m
}

func (g dependencyGraph) isValid(pages []int) bool {
	for i, page := range pages {
		if _, ok := g[page]; !ok {
			// No dependencies for page.
			continue
		}
		for _, remaining := range pages[i+1:] {
			if g[page][remaining] {
				return false
			}
		}

	}
	return true
}

func part1(graph dependencyGraph, pageNumbers [][]int) int {
	var total int
	for _, pn := range pageNumbers {
		if graph.isValid(pn) {
			total += pn[len(pn)/2]
		}
	}
	return total
}

func part2(graph dependencyGraph, pageNumbers [][]int) int {
	var total int
	for _, pn := range pageNumbers {
		if !graph.isValid(pn) {
			slices.SortFunc(pn, func(a, b int) int {
				if a == b {
					return 0
				}
				if graph[a][b] {
					return 1
				}
				return -1
			})
			fmt.Println(pn)
			total += pn[len(pn)/2]
		}
	}
	return total
}

func Run(inputFile string) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	rules, pageNumbers, err := parseData(data)
	if err != nil {
		log.Fatal(err)
	}
	graph := buildDependencyGraph(rules)

	p1 := part1(graph, pageNumbers)
	fmt.Println("Part 1:", p1)

	p2 := part2(graph, pageNumbers)
	fmt.Println("Part 2:", p2)
}
