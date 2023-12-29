// Advent of Code 2023 - Day 25.
package day25

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

type wire struct {
	src, dst                 string
	originalSrc, originalDst string
}

type graph map[string]map[string]bool

func toGraph(wires []wire) graph {
	g := graph{}
	for _, w := range wires {
		if _, ok := g[w.src]; !ok {
			g[w.src] = map[string]bool{}
		}
		if _, ok := g[w.dst]; !ok {
			g[w.dst] = map[string]bool{}
		}
		g[w.src][w.dst] = true
		g[w.dst][w.src] = true
	}
	return g
}

func parseData(data []byte) []wire {
	lines := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
	var wires []wire
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		src := parts[0]
		for _, dst := range strings.Split(parts[1], " ") {
			wires = append(wires, wire{src: src, dst: dst, originalSrc: src, originalDst: dst})
		}
	}
	return wires
}

func kargerMinCut(edges []wire) []wire {
	nodes := map[string]bool{}
	for _, edge := range edges {
		nodes[edge.src] = true
		nodes[edge.dst] = true
	}

	for len(nodes) > 2 {
		// Pick a random edge.
		i := rand.Intn(len(edges))
		edge := edges[i]

		// Contract the edge.
		src := edge.src
		dst := edge.dst
		newNode := edge.src + "-" + edge.dst
		nodes[newNode] = true
		delete(nodes, src)
		delete(nodes, dst)

		var newEdges []wire
		for j, e := range edges {
			// This edge is removed.
			if j == i {
				continue
			}
			// Connect edges to the merged node.
			if e.src == src || e.src == dst {
				e.src = newNode
			}
			if e.dst == src || e.dst == dst {
				e.dst = newNode
			}
			// Remove self loops.
			if e.src == e.dst {
				continue
			}
			newEdges = append(newEdges, e)
		}
		edges = newEdges
	}
	return edges
}

func dfs(p string, g graph, visited map[string]bool) {
	if visited[p] {
		return
	}
	visited[p] = true
	for n := range g[p] {
		dfs(n, g, visited)
	}
}

func part1(wires []wire) int {
	minCut := len(wires)
	var edges []wire
	for minCut != 3 {
		edges = kargerMinCut(wires)
		minCut = len(edges)
	}

	var remainingWires []wire
loop:
	for _, w := range wires {
		for _, e := range edges {
			if w.src == e.originalSrc && w.dst == e.originalDst {
				continue loop
			}
		}
		remainingWires = append(remainingWires, w)
	}

	g := toGraph(remainingWires)
	visited := map[string]bool{}
	dfs(remainingWires[0].src, g, visited)

	return len(visited) * (len(g) - len(visited))
}

func Run(inputFile string) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	hailstones := parseData(data)

	p1 := part1(hailstones)
	fmt.Println("Part 1:", p1)
}
