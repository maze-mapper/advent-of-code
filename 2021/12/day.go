// Advent of Code 2021 - Day 12
package day12

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type caveSystem map[string]map[string]struct{}

// addToNestedMap adds b to the key a in map m
func addToNestedMap(m caveSystem, a, b string) {
	if _, ok := m[a]; !ok {
		m[a] = map[string]struct{}{}
	}
	m[a][b] = struct{}{}
}

func parseData(data []byte) caveSystem {
	lines := strings.Split(
		strings.TrimSuffix(string(data), "\n"), "\n",
	)
	passages := caveSystem{}
	for _, line := range lines {
		caves := strings.Split(line, "-")
		// We will never revisit the start or leave the end
		if caves[0] != "end" && caves[1] != "start" {
			addToNestedMap(passages, caves[0], caves[1])
		}
		if caves[0] != "start" && caves[1] != "end" {
			addToNestedMap(passages, caves[1], caves[0])
		}
	}

	return passages
}

// isUpper returns true if a string is uppercase
func isUpper(s string) bool {
	return strings.ToUpper(s) == s
}

func contains(sl []string, s string) bool {
	for _, ss := range sl {
		if ss == s {
			return true
		}
	}
	return false
}

func visitedSingleSmallCaveTwice(sl []string) bool {
	visted := map[string]int{}
	for _, s := range sl {
		if !isUpper(s) {
			visted[s] += 1
			if visted[s] > 1 {
				return true
			}
		}
	}
	return false
}

func findRoute(passages caveSystem, pos string, currentRoute []string, routes *[][]string, mode string) {
	// Check if we have reached the end
	if pos == "end" {
		*routes = append(*routes, currentRoute)
		return
	}

	// Try the next cave
	for k, _ := range passages[pos] {
		switch mode {
		case "part1":
			// Do not revisit small (lowercase) caves
			if !isUpper(k) && contains(currentRoute, k) {
				continue
			}
		case "part2":
			// Allow a single small (lowercase) cave to be visited twice
			if !isUpper(k) && visitedSingleSmallCaveTwice(currentRoute) && contains(currentRoute, k) {
				continue
			}
		}
		findRoute(passages, k, append(currentRoute, k), routes, mode)
	}
}

func part1(passages caveSystem) int {
	routes := [][]string{}
	findRoute(passages, "start", []string{"start"}, &routes, "part1")
	return len(routes)
}

func part2(passages caveSystem) int {
	routes := [][]string{}
	findRoute(passages, "start", []string{"start"}, &routes, "part2")
	return len(routes)
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	passages := parseData(data)

	p1 := part1(passages)
	fmt.Println("Part 1:", p1)

	p2 := part2(passages)
	fmt.Println("Part 2:", p2)
}
