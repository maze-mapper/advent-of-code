// Advent of Code 2024 - Day 21.
package day21

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseData(data []byte) []string {
	return strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
}

func part1(codes []string) int {
	return solveCodes(codes, 2)
}

func part2(codes []string) int {
	return solveCodes(codes, 25)
}

func solveCodes(codes []string, keypads int) int {
	numKeypadMoves := numericKeypad().moveMap()
	dirKeypadMoves := directionalKeypad().moveMap()
	var sum int
	for _, code := range codes {
		sum += solve(code, numKeypadMoves, dirKeypadMoves, keypads)
	}
	return sum
}

func solve(code string, numKeypadMoves, dirKeypadMoves map[string]map[string][]string, keypads int) int {
	moves := parentKeypresses(numKeypadMoves, map[string]int{code: 1})
	for j := 0; j < keypads; j++ {
		moves = parentKeypresses(dirKeypadMoves, moves)
	}
	var total int
	for k, v := range moves {
		total += len(k) * v
	}
	nCode := strings.TrimSuffix(code, "A")
	n, _ := strconv.Atoi(nCode)
	return n * total

}

func parentKeypresses(keypadMoves map[string]map[string][]string, moves map[string]int) map[string]int {
	out := map[string]int{}
	for k, v := range moves {
		currentKey := "A"
		for i := 0; i < len(k); i++ {
			nextKey := string(k[i])
			kpm := strings.Join(keypadMoves[currentKey][nextKey], "")
			out[kpm] += v
			currentKey = nextKey
		}
	}
	return out
}

type Coord struct {
	X, Y int
}

type keypad struct {
	keys  map[Coord]string
	blank Coord
}

func (k keypad) moveMap() map[string]map[string][]string {
	m := map[string]map[string][]string{}

	for c1, key1 := range k.keys {
		for c2, key2 := range k.keys {
			var moves []string

			hitsBlankIfMoveXFirst := c1.Y == k.blank.Y && c2.X == k.blank.X
			hitsBlankIfMoveYFirst := c1.X == k.blank.X && c2.Y == k.blank.Y
			// <
			if c2.X < c1.X && !hitsBlankIfMoveXFirst {
				for n := 0; n < c1.X-c2.X; n++ {
					moves = append(moves, "<")
				}
			}
			// v
			if c2.Y > c1.Y && !hitsBlankIfMoveYFirst {
				for n := 0; n < c2.Y-c1.Y; n++ {
					moves = append(moves, "v")
				}
			}
			// ^
			if c2.Y < c1.Y && !hitsBlankIfMoveYFirst {
				for n := 0; n < c1.Y-c2.Y; n++ {
					moves = append(moves, "^")
				}
			}
			// >
			if c2.X > c1.X {
				for n := 0; n < c2.X-c1.X; n++ {
					moves = append(moves, ">")
				}
			}

			// <
			if c2.X < c1.X && hitsBlankIfMoveXFirst {
				for n := 0; n < c1.X-c2.X; n++ {
					moves = append(moves, "<")
				}
			}
			// v
			if c2.Y > c1.Y && hitsBlankIfMoveYFirst {
				for n := 0; n < c2.Y-c1.Y; n++ {
					moves = append(moves, "v")
				}
			}
			// ^
			if c2.Y < c1.Y && hitsBlankIfMoveYFirst {
				for n := 0; n < c1.Y-c2.Y; n++ {
					moves = append(moves, "^")
				}
			}

			moves = append(moves, "A")
			if _, ok := m[key1]; !ok {
				m[key1] = map[string][]string{}
			}
			m[key1][key2] = moves
		}
	}
	return m
}

func directionalKeypad() keypad {
	return keypad{
		keys: map[Coord]string{
			{X: 1, Y: 0}: "^",
			{X: 2, Y: 0}: "A",
			{X: 0, Y: 1}: "<",
			{X: 1, Y: 1}: "v",
			{X: 2, Y: 1}: ">",
		},
		blank: Coord{X: 0, Y: 0},
	}
}

func numericKeypad() keypad {
	return keypad{
		keys: map[Coord]string{
			{X: 0, Y: 0}: "7",
			{X: 1, Y: 0}: "8",
			{X: 2, Y: 0}: "9",
			{X: 0, Y: 1}: "4",
			{X: 1, Y: 1}: "5",
			{X: 2, Y: 1}: "6",
			{X: 0, Y: 2}: "1",
			{X: 1, Y: 2}: "2",
			{X: 2, Y: 2}: "3",
			{X: 1, Y: 3}: "0",
			{X: 2, Y: 3}: "A",
		},
		blank: Coord{X: 0, Y: 3},
	}
}

func Run(inputFile string) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	codes := parseData(data)

	p1 := part1(codes)
	fmt.Println("Part 1:", p1)

	p2 := part2(codes)
	fmt.Println("Part 2:", p2)
}
