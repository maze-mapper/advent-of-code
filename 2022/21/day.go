// Advent of Code 2022 - Day 21.
package day21

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type monkey struct {
	label       string
	left, right *monkey
	number      int
	op          string
}

func (m *monkey) traverse() int {
	if m.left != nil && m.right != nil {
		l := m.left.traverse()
		r := m.right.traverse()
		switch m.op {
		case "+":
			return l + r
		case "-":
			return l - r
		case "*":
			return l * r
		case "/":
			return l / r
		}
	}
	return m.number
}

func (m *monkey) pathToMonkey(s string) ([]string, bool) {
	var currentPath []string
	if m.label == s {
		return currentPath, true
	}
	if m.left != nil {
		if p, ok := m.left.pathToMonkey(s); ok {
			currentPath = append([]string{"L"}, p...)
			return currentPath, true
		}
	}
	if m.right != nil {
		if p, ok := m.right.pathToMonkey(s); ok {
			currentPath = append([]string{"R"}, p...)
			return currentPath, true
		}
	}
	return currentPath, false
}

func parseInput(data []byte) *monkey {
	lines := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
	monkeys := make(map[string]*monkey, len(lines))
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		label := parts[0]
		var newMonkey *monkey
		if m, ok := monkeys[label]; ok {
			newMonkey = m
		} else {
			newMonkey = &monkey{label: label}
			monkeys[label] = newMonkey
		}

		opParts := strings.Split(parts[1], " ")
		if len(opParts) == 1 {
			if v, err := strconv.Atoi(opParts[0]); err == nil {
				newMonkey.number = v
			} else {
				log.Fatal(err)
			}
		} else {
			newMonkey.op = opParts[1]
			if m, ok := monkeys[opParts[0]]; ok {
				newMonkey.left = m
			} else {
				leftMonkey := &monkey{label: opParts[0]}
				newMonkey.left = leftMonkey
				monkeys[opParts[0]] = leftMonkey
			}
			if m, ok := monkeys[opParts[2]]; ok {
				newMonkey.right = m
			} else {
				rightMonkey := &monkey{label: opParts[2]}
				newMonkey.right = rightMonkey
				monkeys[opParts[2]] = rightMonkey
			}
		}
	}
	return monkeys["root"]
}

func solve(m *monkey, path []string, want int) int {
	for _, dir := range path {
		var child *monkey
		var other int
		if dir == "L" {
			child = m.left
			other = m.right.traverse()
		} else {
			child = m.right
			other = m.left.traverse()
		}

		switch m.op {
		case "+":
			want -= other
		case "-":
			if dir == "L" {
				want += other
			} else {
				want = other - want
			}
		case "*":
			want /= other
		case "/":
			if dir == "L" {
				want *= other
			} else {
				want = other / want
			}
		}
		m = child
	}
	return want
}

func part1(m *monkey) int {
	return m.traverse()
}

func part2(m *monkey) int {
	path, ok := m.pathToMonkey("humn")
	if !ok {
		log.Fatal("did not find monkey \"humn\"")
	}
	fmt.Println(path)

	var want int
	if path[0] == "L" {
		want = solve(m.left, path[1:], m.right.traverse())
	} else {
		want = solve(m.right, path[1:], m.left.traverse())
	}

	return want
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	rootMonkey := parseInput(data)
	fmt.Println(rootMonkey)

	p1 := part1(rootMonkey)
	fmt.Println("Part 1:", p1)

	p2 := part2(rootMonkey)
	fmt.Println("Part 2:", p2)
}
