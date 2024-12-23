// Advent of Code 2024 - Day 23.
package day23

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func parseData(data []byte) (map[string]map[string]bool, error) {
	lines := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
	m := map[string]map[string]bool{}
	for _, line := range lines {
		before, after, found := strings.Cut(line, "-")
		if !found {
			return nil, fmt.Errorf("line %q did not match expected format", line)
		}
		if _, ok := m[before]; !ok {
			m[before] = map[string]bool{}
		}
		if _, ok := m[after]; !ok {
			m[after] = map[string]bool{}
		}
		m[before][after] = true
		m[after][before] = true
	}
	return m, nil
}

func part1(graph map[string]map[string]bool) int {
	loops := map[string]bool{}
	for n0, v := range graph {
		if !strings.HasPrefix(n0, "t") {
			continue
		}
		for n1 := range v {
			for n2 := range graph[n1] {
				for n3 := range graph[n2] {
					if n3 == n0 {
						loop := []string{n0, n1, n2}
						slices.Sort(loop)
						key := strings.Join(loop, ",")
						loops[key] = true
					}
				}
			}
		}
	}
	return len(loops)
}

func part2(graph map[string]map[string]bool) string {
	// https://en.wikipedia.org/wiki/Clique_problem
	var maximalClique map[string]bool
	for initialNode := range graph {
		clique := map[string]bool{initialNode: true}
		for node, neighbours := range graph {
			if node == initialNode {
				continue
			}
			inClique := true
			for cliqueNode := range clique {
				if !neighbours[cliqueNode] {
					inClique = false
					break
				}
			}
			if inClique {
				clique[node] = true
			}
		}
		if len(clique) > len(maximalClique) {
			maximalClique = clique
		}
	}
	var maximalCliqueNodes []string
	for node := range maximalClique {
		maximalCliqueNodes = append(maximalCliqueNodes, node)
	}
	slices.Sort(maximalCliqueNodes)
	return strings.Join(maximalCliqueNodes, ",")
}

func Run(inputFile string) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	graph, err := parseData(data)
	if err != nil {
		log.Fatal(err)
	}

	p1 := part1(graph)
	fmt.Println("Part 1:", p1)

	p2 := part2(graph)
	fmt.Println("Part 2:", p2)
}
