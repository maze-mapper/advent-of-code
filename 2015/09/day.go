// Advent of Code 2015 - Day 9
package day9

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Graph map[string]map[string]int

// makeAdjacencyList coverts the input data in to a Graph data structure
func makeAdjacencyList(data []byte) Graph {
	graph := Graph{}
	lines := strings.Split(
		strings.TrimSuffix(string(data), "\n"), "\n",
	)

	for _, line := range lines {
		parts := strings.Split(line, " ")
		nodeA, nodeB := parts[0], parts[2]
		weight, err := strconv.Atoi(parts[4])
		if err != nil {
			log.Fatal(err)
		}

		// Initialise nested maps
		if _, ok := graph[nodeA]; !ok {
			graph[nodeA] = make(map[string]int)
		}
		if _, ok := graph[nodeB]; !ok {
			graph[nodeB] = make(map[string]int)
		}

		if _, ok := graph[nodeA][nodeB]; !ok {
			graph[nodeA][nodeB] = weight
		} else {
			log.Fatal(err)
		}

		if _, ok := graph[nodeB][nodeA]; !ok {
			graph[nodeB][nodeA] = weight
		} else {
			log.Fatal(err)
		}
	}
	return graph
}

// isExplored checks if a string is present in a list of strings
func isExplored(explored []string, node string) bool {
	for _, n := range explored {
		if n == node {
			return true
		}
	}
	return false
}

// dfs implements a Depth First Search of a Graph
func dfs(graph Graph, currentNode string, explored []string, currentDistance int, bestDistance *int, shortCircuit bool) {
	if shortCircuit && currentDistance >= *bestDistance {
		return
	}

	if len(graph) == len(explored) {
		if currentDistance < *bestDistance {
			*bestDistance = currentDistance
			fmt.Println("New best distance is", *bestDistance, "via", explored)
		}
		return
	}

	for nextNode, distance := range graph[currentNode] {
		if !isExplored(explored, nextNode) {
			dfs(graph, nextNode, append(explored, nextNode), currentDistance+distance, bestDistance, shortCircuit)
		}
	}
	return
}

func part1(graph Graph) {
	bestDistance := int(^uint(0) >> 1) // Initialise to max value

	for startNode, _ := range graph {
		dfs(graph, startNode, []string{startNode}, 0, &bestDistance, true)
	}
	fmt.Println("Part 1: minimum distance is", bestDistance)
}

func part2(graph Graph) {
	bestDistance := int(^uint(0) >> 1) // Initialise to max value

	// Convert all distances to negative
	// The best distance will now be the longest route which will be the most negative
	for _, subGraph := range graph {
		for subNode, distance := range subGraph {
			subGraph[subNode] = 0 - distance
		}
	}

	for startNode, _ := range graph {
		dfs(graph, startNode, []string{startNode}, 0, &bestDistance, false)
	}
	fmt.Println("Part 2: maximum distance is", -bestDistance)
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	graph := makeAdjacencyList(data)
	part1(graph)
	part2(graph)
}
