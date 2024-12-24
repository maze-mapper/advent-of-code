// Advent of Code 2024 - Day 24.
package day24

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"
)

type logicGate struct {
	inputWire1, inputWire2, outputWire string
	operator                           string
}

func parseData(data []byte) (map[string]bool, []logicGate, error) {
	sections := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n\n")

	inputWires := map[string]bool{}
	for _, line := range strings.Split(sections[0], "\n") {
		wire, value, found := strings.Cut(line, ": ")
		if !found {
			return nil, nil, fmt.Errorf("unable to parse line %q", line)
		}
		switch value {
		case "0":
			inputWires[wire] = false

		case "1":
			inputWires[wire] = true
		}
	}

	lines := strings.Split(sections[1], "\n")
	gates := make([]logicGate, len(lines))
	for i, line := range lines {
		gate := logicGate{}
		n, err := fmt.Sscanf(line, "%s %s %s -> %s", &gate.inputWire1, &gate.operator, &gate.inputWire2, &gate.outputWire)
		if err != nil {
			return nil, nil, err
		}
		if n != 4 {
			return nil, nil, fmt.Errorf("parsed %d items from line %q, expected 4 ", n, line)
		}
		gates[i] = gate
	}

	return inputWires, gates, nil
}

func part1(inputWires map[string]bool, gates []logicGate) int {
	allWires := map[string]bool{}
	for wire, value := range inputWires {
		allWires[wire] = value
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	for _, gate := range gates {
		gate := gate
		wg.Add(1)
		go func() {
			defer wg.Done()
			var input1, input2 bool
			var ok1, ok2 bool
			for !(ok1 && ok2) {
				mu.Lock()
				input1, ok1 = allWires[gate.inputWire1]
				input2, ok2 = allWires[gate.inputWire2]
				mu.Unlock()
			}
			var output bool
			switch gate.operator {
			case "AND":
				output = input1 && input2
			case "OR":
				output = input1 || input2
			case "XOR":
				output = input1 != input2
			}
			mu.Lock()
			allWires[gate.outputWire] = output
			mu.Unlock()

		}()
	}
	wg.Wait()

	return numberFromBits(allWires, "z")
}

func part2(inputWires map[string]bool, gates []logicGate) string {
	// Solved by inspecting graph with GraphViz and then looking for the least
	// significant bit where the expected answer differs from the gate output.

	// fmt.Println(graphViz(gates))
	// x := numberFromBits(inputWires, "x")
	// y := numberFromBits(inputWires, "y")
	// z := part1(inputWires, gates)
	// fmt.Printf("want: %b, %d\n", x+y, x+y)
	// fmt.Printf("got:  %b, %d\n", z, z)

	swaps := map[string]string{
		// z19 should be output from XOR.
		"z19": "vvf",
		"vvf": "z19",
		// z37 should be output from XOR.
		"z37": "nvh",
		"nvh": "z37",
		// Swap AND and XOR outputs near x23 and y23.
		"fgn": "dck",
		"dck": "fgn",
		// z12 should be output from XOR.
		"qdg": "z12",
		"z12": "qdg",
	}
	for i := range gates {
		outputWire := gates[i].outputWire
		if swap, ok := swaps[outputWire]; ok {
			gates[i].outputWire = swap
		}
	}

	var swappedWires []string
	for k := range swaps {
		swappedWires = append(swappedWires, k)
	}
	slices.Sort(swappedWires)
	return strings.Join(swappedWires, ",")
}

func graphViz(gates []logicGate) string {
	var b strings.Builder
	b.WriteString("digraph {\n")
	for i, gate := range gates {
		b.WriteString(fmt.Sprintf("  %d [label=\"%s\"]\n", i, gate.operator))
		b.WriteString(fmt.Sprintf("  %s -> %d\n", gate.inputWire1, i))
		b.WriteString(fmt.Sprintf("  %s -> %d\n", gate.inputWire2, i))
		b.WriteString(fmt.Sprintf("  %d -> %s\n", i, gate.outputWire))
	}
	b.WriteString("}")
	return b.String()
}

func numberFromBits(wires map[string]bool, prefix string) int {
	var result int
	for wire, value := range wires {
		if !strings.HasPrefix(wire, prefix) {
			continue
		}
		n, err := strconv.Atoi(strings.TrimPrefix(wire, prefix))
		if err != nil {
			log.Fatal(err)
		}
		if value {
			result += 1 << n
		}
	}
	return result
}

func Run(inputFile string) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	inputWires, gates, err := parseData(data)
	if err != nil {
		log.Fatal(err)
	}

	p1 := part1(inputWires, gates)
	fmt.Println("Part 1:", p1)

	p2 := part2(inputWires, gates)
	fmt.Println("Part 2:", p2)
}
