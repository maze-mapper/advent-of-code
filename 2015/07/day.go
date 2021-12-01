// Advent of Code 2015 - Day 7
package day7

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

// instruction is an object representing each instruction line
type instruction struct {
	outWire string
	gate    string
	inWires []string
	value   uint16
}

// addInstruction adds an output wire to the given map if all inputs are already present
func (i *instruction) addInstruction(m map[string]uint16) bool {
	switch i.gate {

	case "AND":
		lhs, ok := m[i.inWires[0]]
		if !ok {
			return false
		}
		var rhs uint16
		if len(i.inWires) > 1 {
			rhs, ok = m[i.inWires[1]]
			if !ok {
				return false
			}
		} else {
			rhs = i.value
		}
		m[i.outWire] = lhs & rhs

	case "OR":
		lhs, ok := m[i.inWires[0]]
		if !ok {
			return false
		}
		rhs, ok := m[i.inWires[1]]
		if !ok {
			return false
		}
		m[i.outWire] = lhs | rhs

	case "NOT":
		val, ok := m[i.inWires[0]]
		if !ok {
			return false
		}
		m[i.outWire] = ^val

	case "LSHIFT":
		val, ok := m[i.inWires[0]]
		if !ok {
			return false
		}
		m[i.outWire] = val << i.value

	case "RSHIFT":
		val, ok := m[i.inWires[0]]
		if !ok {
			return false
		}
		m[i.outWire] = val >> i.value

	default:
		if len(i.inWires) > 0 {
			val, ok := m[i.inWires[0]]
			if !ok {
				return false
			}
			m[i.outWire] = val
		} else {
			m[i.outWire] = i.value
		}

	}

	return true
}

var (
	andGatePattern    = regexp.MustCompile(`([a-z]+|[0-9]+) AND ([a-z]+)`)
	orGatePattern     = regexp.MustCompile(`([a-z]+) OR ([a-z]+)`)
	notGatePattern    = regexp.MustCompile(`NOT ([a-z]+)`)
	lshiftGatePattern = regexp.MustCompile(`([a-z]+) LSHIFT ([0-9]+)`)
	rshiftGatePattern = regexp.MustCompile(`([a-z]+) RSHIFT ([0-9]+)`)
)

// parseInstruction converts the given string in to an instruction object
func parseInstruction(line string) instruction {
	parts := strings.Split(line, " -> ")

	i := instruction{outWire: parts[1]}
	switch {
	case strings.Contains(parts[0], "AND"):
		matches := andGatePattern.FindStringSubmatch(parts[0])
		i.gate = "AND"
		if val, err := strconv.Atoi(matches[1]); err == nil {
			i.value = uint16(val)
		} else {
			i.inWires = []string{matches[1]}
		}
		i.inWires = append(i.inWires, matches[2])

	case strings.Contains(parts[0], "OR"):
		matches := orGatePattern.FindStringSubmatch(parts[0])
		i.gate = "OR"
		i.inWires = []string{matches[1], matches[2]}

	case strings.Contains(parts[0], "NOT"):
		matches := notGatePattern.FindStringSubmatch(parts[0])
		i.gate = "NOT"
		i.inWires = []string{matches[1]}

	case strings.Contains(parts[0], "LSHIFT"):
		matches := lshiftGatePattern.FindStringSubmatch(parts[0])
		i.gate = "LSHIFT"
		i.inWires = []string{matches[1]}
		if val, err := strconv.Atoi(matches[2]); err == nil {
			i.value = uint16(val)
		} else {
			log.Fatal(err)
		}

	case strings.Contains(parts[0], "RSHIFT"):
		matches := rshiftGatePattern.FindStringSubmatch(parts[0])
		i.gate = "RSHIFT"
		i.inWires = []string{matches[1]}
		if val, err := strconv.Atoi(matches[2]); err == nil {
			i.value = uint16(val)
		} else {
			log.Fatal(err)
		}

	default:
		if val, err := strconv.Atoi(parts[0]); err == nil {
			i.value = uint16(val)
		} else {
			i.inWires = []string{parts[0]}
		}
	}

	return i
}

func part1(instructions *list.List) uint16 {
	wireSignals := make(map[string]uint16)

	// Loop over all instructions until there is a signal on wire "a"
	elem := instructions.Front()
	for {

		if _, ok := wireSignals["a"]; ok {
			break
		}

		ins := elem.Value.(instruction)
		elem2 := elem.Next()
		if added := ins.addInstruction(wireSignals); added {
			instructions.Remove(elem)
		} else {
			instructions.MoveToBack(elem)
		}
		elem = elem2

	}

	fmt.Println(wireSignals["a"])
	return wireSignals["a"]
}

func part2(instructions *list.List, wireB uint16) {
	// Override value for wire "b"
	for elem := instructions.Front(); elem != nil; elem = elem.Next() {
		ins := elem.Value.(instruction)
		if ins.outWire == "b" {
			instructions.Remove(elem)
			instructions.PushFront(instruction{"b", "", []string{}, wireB})
			break
		}
	}

	part1(instructions)
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(
		strings.TrimSuffix(string(data), "\n"), "\n",
	)

	instructions := list.New()
	instructions2 := list.New()
	for _, line := range lines {
		instruction := parseInstruction(line)
		instructions.PushBack(instruction)
		instructions2.PushBack(instruction)
	}

	val := part1(instructions)
	part2(instructions2, val)
}
