// Advent of Code 2022 - Day 11.
package day11

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	//        "strconv"
	//        "strings"
)

type operation func(x int) int

type monkey struct {
	items               []int
	op                  operation
	divisor             int
	destTrue, destFalse int
	inspections         int
}

func (m *monkey) inspect(item int) int {
	m.inspections += 1
	return m.op(item)
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

func realMonkeys() []monkey {
	return []monkey{
		monkey{
			items:     []int{84, 72, 58, 51},
			op:        func(x int) int { return x * 3 },
			divisor:   13,
			destTrue:  1,
			destFalse: 7,
		},
		monkey{
			items:     []int{88, 58, 58},
			op:        func(x int) int { return x + 8 },
			divisor:   2,
			destTrue:  7,
			destFalse: 5,
		},
		monkey{
			items:     []int{93, 82, 71, 77, 83, 53, 71, 89},
			op:        func(x int) int { return x * x },
			divisor:   7,
			destTrue:  3,
			destFalse: 4,
		},
		monkey{
			items:     []int{81, 68, 65, 81, 73, 77, 96},
			op:        func(x int) int { return x + 2 },
			divisor:   17,
			destTrue:  4,
			destFalse: 6,
		},
		monkey{
			items:     []int{75, 80, 50, 73, 88},
			op:        func(x int) int { return x + 3 },
			divisor:   5,
			destTrue:  6,
			destFalse: 0,
		},
		monkey{
			items:     []int{59, 72, 99, 87, 91, 81},
			op:        func(x int) int { return x * 17 },
			divisor:   11,
			destTrue:  2,
			destFalse: 3,
		},
		monkey{
			items:     []int{86, 69},
			op:        func(x int) int { return x + 6 },
			divisor:   3,
			destTrue:  1,
			destFalse: 0,
		},
		monkey{
			items:     []int{91},
			op:        func(x int) int { return x + 1 },
			divisor:   19,
			destTrue:  2,
			destFalse: 5,
		},
	}
}

var exampleMonkeys = []monkey{
	monkey{
		items:     []int{79, 98},
		op:        func(x int) int { return x * 19 },
		divisor:   23,
		destTrue:  2,
		destFalse: 3,
	},
	monkey{
		items:     []int{54, 65, 75, 74},
		op:        func(x int) int { return x + 6 },
		divisor:   19,
		destTrue:  2,
		destFalse: 0,
	},
	monkey{
		items:     []int{79, 60, 97},
		op:        func(x int) int { return x * x },
		divisor:   13,
		destTrue:  1,
		destFalse: 3,
	},
	monkey{
		items:     []int{74},
		op:        func(x int) int { return x + 3 },
		divisor:   17,
		destTrue:  0,
		destFalse: 1,
	},
}

func parseMonkey(text string) monkey {
	m := monkey{}
	return m
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
	_, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	//        monkeys := parseInput(data)

	p1 := part1(realMonkeys())
	fmt.Println("Part 1:", p1)

	p2 := part2(realMonkeys())
	fmt.Println("Part 2:", p2)
}
