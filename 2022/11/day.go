// Advent of Code 2022 - Day 11.
package day11

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	items               []int
	a, b, c             int // Coefficients for polynomial inspection operation.
	divisor             int
	destTrue, destFalse int
	inspections         int
}

func (m *monkey) inspect(item int) int {
	m.inspections += 1
	// ax^2 + bx + c
	return (m.a * item * item) + (m.b * item) + m.c
}

func (m *monkey) test(item int) int {
	if item%m.divisor == 0 {
		return m.destTrue
	} else {
		return m.destFalse
	}
}

func (m *monkey) giveItem(item int) {
	m.items = append(m.items, item)
}

func parseMonkey(text string) monkey {
	m := monkey{}

	lines := strings.Split(text, "\n")

	var err error
	_, err = fmt.Sscanf(lines[3], "  Test: divisible by %d", &(m.divisor))
	if err != nil {
		log.Fatal(err)
	}
	_, err = fmt.Sscanf(lines[4], "    If true: throw to monkey %d", &(m.destTrue))
	if err != nil {
		log.Fatal(err)
	}
	_, err = fmt.Sscanf(lines[5], "    If false: throw to monkey %d", &(m.destFalse))
	if err != nil {
		log.Fatal(err)
	}

	itemLine := strings.TrimPrefix(lines[1], "  Starting items: ")
	items := strings.Split(itemLine, ", ")
	for _, item := range items {
		n, err := strconv.Atoi(item)
		if err != nil {
			log.Fatal(err)
		}
		m.items = append(m.items, n)
	}

	op := strings.TrimPrefix(lines[2], "  Operation: new = ")
	parts := strings.Split(op, " ")
	switch {
	case op == "old * old":
		m.a = 1
	case parts[1] == "*":
		v, err := strconv.Atoi(parts[2])
		if err != nil {
			log.Fatal(err)
		}
		m.b = v
	case parts[1] == "+":
		v, err := strconv.Atoi(parts[2])
		if err != nil {
			log.Fatal(err)
		}
		m.b = 1
		m.c = v
	default:
		log.Fatal("Unkown operation.")
	}

	return m
}

func parseInput(data []byte) []monkey {
	groups := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n\n")
	monkeys := make([]monkey, len(groups))
	for i, group := range groups {
		monkeys[i] = parseMonkey(group)
	}
	return monkeys
}

func round(monkeys []monkey, f func(int) int) {
	for i := 0; i < len(monkeys); i++ {
		m := &monkeys[i]
		toThrow := make([]int, len(m.items))
		copy(toThrow, m.items)
		m.items = nil
		for _, item := range toThrow {
			item = m.inspect(item)
			item = f(item)
			dest := m.test(item)
			monkeys[dest].giveItem(item)
		}
	}
}

func monkeyBusiness(monkeys []monkey) int {
	inspections := make([]int, len(monkeys))
	for i, m := range monkeys {
		inspections[i] = m.inspections
	}
	sort.Sort(sort.Reverse(sort.IntSlice(inspections)))
	return inspections[0] * inspections[1]
}

func part1(monkeys []monkey) int {
	f := func(item int) int {
		return item / 3
	}
	for r := 0; r < 20; r++ {
		round(monkeys, f)
	}
	return monkeyBusiness(monkeys)
}

func part2(monkeys []monkey) int {
	lcm := 1
	for _, m := range monkeys {
		lcm *= m.divisor
	}

	f := func(item int) int {
		return item % lcm
	}
	for r := 0; r < 10000; r++ {
		round(monkeys, f)
	}
	return monkeyBusiness(monkeys)
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	monkeys := parseInput(data)
	monkeys2 := parseInput(data)

	p1 := part1(monkeys)
	fmt.Println("Part 1:", p1)

	p2 := part2(monkeys2)
	fmt.Println("Part 2:", p2)
}
