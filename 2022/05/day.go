// Advent of Code 2022 - Day 5.
package day5

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type stack []string

func (s stack) Push(c string) stack {
	return append(s, c)

}

func (s stack) Push2(c stack) stack {
	return append(s, c...)
}

func (s stack) Pop() (stack, string) {
	if len(s) == 0 {
		return nil, ""
	}
	i := len(s) - 1
	v := s[i]
	s[i] = ""
	return s[:i], v
}

func (s stack) Pop2(n int) (stack, stack) {
	idx := len(s) - n
	return s[:idx], s[idx:]
}

func (s stack) Peek() string {
	return s[len(s)-1]
}

type move struct {
	quantity, from, to int
}

func parseInput(data []byte) ([]stack, []move) {
	parts := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n\n")

	crateLines := strings.Split(parts[0], "\n")
	fields := strings.Fields(crateLines[len(crateLines)-1])
	crates := make([]stack, len(fields))
	for i := len(crateLines) - 2; i >= 0; i-- {

		for j := 0; j < len(crateLines[i]); j += 3 {
			c := crateLines[i][j : j+3]
			if c != "   " {
				idx := j / 4
				crates[idx] = crates[idx].Push(strings.TrimPrefix(strings.TrimSuffix(c, "]"), "["))
			}
			j += 1
		}
	}

	moveLines := strings.Split(parts[1], "\n")
	moves := make([]move, len(moveLines))
	for i, line := range moveLines {
		var q, f, t int
		n, err := fmt.Sscanf(line, "move %d from %d to %d", &q, &f, &t)
		if err != nil {
			log.Fatal(err)
		}
		if n != 3 {
			log.Fatalf("Only read %d items from move string", n)
		}
		// Adjust stack indices to be zero indexed.
		moves[i] = move{quantity: q, from: f - 1, to: t - 1}
	}

	return crates, moves
}

func part1(crates []stack, moves []move) string {
	for _, m := range moves {
		for i := 0; i < m.quantity; i++ {
			var c string
			crates[m.from], c = crates[m.from].Pop()
			if c == "" {
				panic(m)
			}
			crates[m.to] = crates[m.to].Push(c)
		}
	}

	b := strings.Builder{}
	for _, c := range crates {
		b.WriteString(c.Peek())
	}
	return b.String()
}

func part2(crates []stack, moves []move) string {
	for _, m := range moves {
		var c stack
		crates[m.from], c = crates[m.from].Pop2(m.quantity)
		crates[m.to] = crates[m.to].Push2(c)
	}

	b := strings.Builder{}
	for _, c := range crates {
		b.WriteString(c.Peek())
	}
	return b.String()
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	crates, moves := parseInput(data)
	crates2 := make([]stack, len(crates))
	for i, c := range crates {
		c2 := make(stack, len(crates[i]))
		copy(c2, c)
		crates2[i] = c2
	}

	p1 := part1(crates, moves)
	fmt.Println("Part 1:", p1)

	p2 := part2(crates2, moves)
	fmt.Println("Part 2:", p2)
}
