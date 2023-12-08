// Advent of Code 2023 - Day 8.
package day8

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type network struct {
	instructions  string
	lefts, rights map[string]string
}

func (n network) step(currentNode, direction string) string {
	var m map[string]string
	switch direction {
	case "L":
		m = n.lefts
	case "R":
		m = n.rights
	default:
		log.Fatalf("Unrecognised direction %s", direction)
	}
	nextNode, ok := m[currentNode]
	if !ok {
		log.Fatalf("No node from %s for instruction %s", currentNode, direction)
	}
	return nextNode
}

func parseData(data []byte) network {
	parts := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n\n")
	n := network{
		instructions: parts[0],
		lefts:        map[string]string{},
		rights:       map[string]string{},
	}
	for _, line := range strings.Split(parts[1], "\n") {
		var start, left, right string
		line = strings.TrimSuffix(strings.ReplaceAll(line, ",", ""), ")")
		fmt.Sscanf(line, "%s = (%s %s", &start, &left, &right)
		n.lefts[start] = left
		n.rights[start] = right
	}
	return n
}

func part1(input network) int {
	current := "AAA"
	steps := 0
	i := 0
	for current != "ZZZ" {
		direction := string(input.instructions[i])
		current = input.step(current, direction)
		steps += 1
		i += 1
		i %= len(input.instructions)
	}
	return steps
}

func isTerminal(nodes []string) bool {
	for _, node := range nodes {
		if !strings.HasSuffix(node, "Z") {
			return false
		}
	}
	return true
}

// findLoop returns the cycle length and the first times terminal nodes are visited.
func findLoop(input network, node string) (int, map[string]int) {
	steps := 0
	i := 0
	visited := map[string]int{}
	key := fmt.Sprintf("%s-%d", node, i)
	for {
		if _, ok := visited[key]; ok {
			break
		}
		visited[key] = steps
		direction := string(input.instructions[i])
		node = input.step(node, direction)
		steps += 1
		i += 1
		i %= len(input.instructions)
		key = fmt.Sprintf("%s-%d", node, i)
	}
	cycleStart := visited[key]
	cycleLength := steps - cycleStart

	terminalNodes := map[string]int{}
	for k, v := range visited {
		parts := strings.Split(k, "-")
		if strings.HasSuffix(parts[0], "Z") && v >= cycleStart {
			terminalNodes[parts[0]] = v
		}
	}
	return cycleLength, terminalNodes
}

func part2(input network) int {
	var curentNodes []string
	for node := range input.lefts {
		if strings.HasSuffix(node, "A") {
			curentNodes = append(curentNodes, node)
		}
	}

	var numbers []int
	for _, node := range curentNodes {
		s, m := findLoop(input, node)
		fmt.Println(node, s, m)
		numbers = append(numbers, s)
	}

	n := numbers[0]
loop:
	for ; ; n += numbers[0] {
		for _, nn := range numbers {
			// fmt.Println(n, nn, n%nn)
			if n%nn != 0 {
				continue loop
			}
		}
		break
	}

	return n
}

func Run(inputFile string) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	n := parseData(data)

	p1 := part1(n)
	fmt.Println("Part 1:", p1)

	p2 := part2(n)
	fmt.Println("Part 2:", p2)
}
