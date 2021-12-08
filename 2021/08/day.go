// Advent of Code 2021 - Day 8
package day8

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/bits"
	"strings"
)

// observation holds the input and output signals
type observation struct {
	signals, outputs []uint8
}

// outputDigits returns the output digits given a mapping of digit to signals
func (ob *observation) outputDigits(m map[int]uint8) int {
	var out int
	l := len(ob.outputs)
	for digit := 0; digit < l; digit++ {
		n, err := findNumber(ob.outputs[digit], m)
		if err != nil {
			log.Fatal(err)
		}
		// Shift digit by the appropriate amount
		for i := 0; i < l-digit-1; i++ {
			n *= 10
		}
		out += n
	}
	return out
}

// findNumber returns the number displayed for a given signal and mapping of digit to signals
func findNumber(b uint8, m map[int]uint8) (int, error) {
	for k, v := range m {
		if b == v {
			return k, nil
		}
	}
	return -1, fmt.Errorf("Number %d not found", b)
}

// makeBinary converts a signal string in to a binary number where each letter (a-g) represents a single bit
func makeBinary(s string) uint8 {
	var b uint8
	for _, r := range s {
		switch r {
		case 'a':
			b |= 0b1
		case 'b':
			b |= 0b10
		case 'c':
			b |= 0b100
		case 'd':
			b |= 0b1000
		case 'e':
			b |= 0b10000
		case 'f':
			b |= 0b100000
		case 'g':
			b |= 0b1000000
		}
	}
	return b
}

// makeBinarySlice convers a slice of strings in to a slice of uint8
func makeBinarySlice(sl []string) []uint8 {
	b := make([]uint8, len(sl))
	for i, s := range sl {
		b[i] = makeBinary(s)
	}
	return b
}

// parseData converts the input data in to a slice of observations
func parseData(data []byte) []observation {
	lines := strings.Split(
		strings.TrimSuffix(string(data), "\n"), "\n",
	)

	observations := make([]observation, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " | ")

		observations[i] = observation{
			signals: makeBinarySlice(strings.Split(parts[0], " ")),
			outputs: makeBinarySlice(strings.Split(parts[1], " ")),
		}
	}
	return observations
}

func part1(observations []observation) int {
	count := 0
	for _, ob := range observations {
		for _, out := range ob.outputs {
			l := bits.OnesCount8(out)
			if l == 2 || l == 3 || l == 4 || l == 7 {
				count += 1
			}
		}
	}
	return count
}

// findUniqueSegmentDigits determines the signals for digits where the number of signals uniquely identifies them
func findUniqueSegmentDigits(s []uint8, m map[int]uint8) {
	for _, ss := range s {
		l := bits.OnesCount8(ss)
		switch l {
		case 2:
			m[1] = ss
		case 3:
			m[7] = ss
		case 4:
			m[4] = ss
		case 7:
			m[8] = ss
		}
	}
}

// findOtherSegmentDigits determines the signals for digits that are not uniquely identified by the number of signals
func findOtherSegmentDigits(s []uint8, m map[int]uint8) {
	for _, ss := range s {
		if bits.OnesCount8(ss) == 5 { // Two, three or five
			switch {
			// Two: union of 2 and 4 is 8
			case ss|m[4] == m[8]:
				m[2] = ss

			// Three: union of 1 and 3 is 3
			case ss|m[1] == ss:
				m[3] = ss

			// Five
			default:
				m[5] = ss
			}

		} else if bits.OnesCount8(ss) == 6 { // Zero, six or nine
			switch {
			// Six: union of 1 and six gives 8
			case ss|m[1] == m[8]:
				m[6] = ss

			// Nine: union of 4 and 9 gives 9
			case ss|m[4] == ss:
				m[9] = ss

			// Zero
			default:
				m[0] = ss
			}
		}
	}
}

// solve returns the output digits for a given observation
func solve(ob observation) int {
	solution := map[int]uint8{}
	findUniqueSegmentDigits(ob.signals, solution)
	findOtherSegmentDigits(ob.signals, solution)
	return ob.outputDigits(solution)
}

func part2(observations []observation) int {
	sum := 0
	for _, ob := range observations {
		out := solve(ob)
		sum += out
	}
	return sum
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	observations := parseData(data)

	p1 := part1(observations)
	fmt.Println("Part 1:", p1)

	p2 := part2(observations)
	fmt.Println("Part 2:", p2)
}
