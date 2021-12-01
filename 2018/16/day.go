// Advent of Code 2018 - Day 16
package day16

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type registers [4]int

// Opcode holds unformation on a named opcode
type Opcode struct {
	name    string
	a, b, c int
}

// uOpcode holds information on a numbered opcode
type uOpcode struct {
	number, a, b, c int
}

// stringToUOpcode converts a string "1 2 3 4" in to a uOpcode
func stringToUOpcode(s string) uOpcode {
	op := uOpcode{}
	for i, ss := range strings.Split(s, " ") {
		if n, err := strconv.Atoi(ss); err == nil {
			switch i {
			case 0:
				op.number = n
			case 1:
				op.a = n
			case 2:
				op.b = n
			case 3:
				op.c = n
			}
		} else {
			log.Fatal(err)
		}
	}
	return op
}

// process will apply an opcode to the given registers
func (o *Opcode) process(r *registers) {
	switch o.name {

	case "addr":
		r[o.c] = r[o.a] + r[o.b]

	case "addi":
		r[o.c] = r[o.a] + o.b

	case "mulr":
		r[o.c] = r[o.a] * r[o.b]

	case "muli":
		r[o.c] = r[o.a] * o.b

	case "banr":
		r[o.c] = r[o.a] & r[o.b]

	case "bani":
		r[o.c] = r[o.a] & o.b

	case "borr":
		r[o.c] = r[o.a] | r[o.b]

	case "bori":
		r[o.c] = r[o.a] | o.b

	case "setr":
		r[o.c] = r[o.a]

	case "seti":
		r[o.c] = o.a

	case "gtir":
		if o.a > r[o.b] {
			r[o.c] = 1
		} else {
			r[o.c] = 0
		}

	case "gtri":
		if r[o.a] > o.b {
			r[o.c] = 1
		} else {
			r[o.c] = 0
		}

	case "gtrr":
		if r[o.a] > r[o.b] {
			r[o.c] = 1
		} else {
			r[o.c] = 0
		}

	case "eqir":
		if o.a == r[o.b] {
			r[o.c] = 1
		} else {
			r[o.c] = 0
		}

	case "eqri":
		if r[o.a] == o.b {
			r[o.c] = 1
		} else {
			r[o.c] = 0
		}

	case "eqrr":
		if r[o.a] == r[o.b] {
			r[o.c] = 1
		} else {
			r[o.c] = 0
		}

	}
}

// Names of all opcodes
var opcodeNames = []string{
	"addr",
	"addi",
	"mulr",
	"muli",
	"banr",
	"bani",
	"borr",
	"bori",
	"setr",
	"seti",
	"gtir",
	"gtri",
	"gtrr",
	"eqir",
	"eqri",
	"eqrr",
}

// sample holds before and after states of registers after an unnamed opcode is appled
type sample struct {
	before, after registers
	opcode        uOpcode
}

func part1(samples []sample, threshold int) {
	total := 0
	for _, s := range samples {
		count := 0
		for _, name := range opcodeNames {
			var r registers
			for i, v := range s.before {
				r[i] = v
			}

			o := Opcode{
				name: name,
				a:    s.opcode.a,
				b:    s.opcode.b,
				c:    s.opcode.c,
			}

			o.process(&r)

			if r == s.after {
				count += 1
			}
		}

		if count >= threshold {
			total += 1
			continue
		}
	}
	fmt.Println("Part 1:", total, "samples behave like", threshold, "or more opcodes")
}

// solve reduces the list of possible opcodes to the only viable mapping
func solve(opcodeNumbers map[int][]string) map[int]string {
	// Create a slice of the available opcode names
	availableOpcodeNames := make([]string, len(opcodeNames))
	copy(availableOpcodeNames, opcodeNames)
	solution := map[int]string{}

	for len(availableOpcodeNames) > 0 {
		for num, possibleOpcodes := range opcodeNumbers {
			valid := 0
			chosenOpcode := ""
			// Check how many of the opcode names remain available
			for _, possibleOpcode := range possibleOpcodes {
				for _, name := range availableOpcodeNames {
					if possibleOpcode == name {
						valid += 1
						chosenOpcode = name
					}
				}
			}
			// If there is only one available opcode name we add it to our solution
			if valid == 1 {
				solution[num] = chosenOpcode
				// Remove the opcode name from the available names
				for i, name := range availableOpcodeNames {
					if name == chosenOpcode {
						availableOpcodeNames = append(availableOpcodeNames[:i], availableOpcodeNames[i+1:]...)
						break
					}
				}
			}
		}
	}

	return solution
}

func part2(samples []sample, program []uOpcode) {
	opcodeNumbers := map[int][]string{}

	for _, s := range samples {
		// Find all possible opcodes for this sample
		possibleOpcodes := []string{}
		for _, name := range opcodeNames {
			var r registers
			for i, v := range s.before {
				r[i] = v
			}

			o := Opcode{
				name: name,
				a:    s.opcode.a,
				b:    s.opcode.b,
				c:    s.opcode.c,
			}

			o.process(&r)

			if r == s.after {
				possibleOpcodes = append(possibleOpcodes, name)
			}
		}

		// If opcode number already has some possibilities, find the intersection between the previous and current values
		if currentOpcodes, ok := opcodeNumbers[s.opcode.number]; ok {
			revisedOpcodes := []string{}
			for _, possible := range possibleOpcodes {
				for _, c := range currentOpcodes {
					if c == possible {
						revisedOpcodes = append(revisedOpcodes, c)
					}
				}
			}
			possibleOpcodes = revisedOpcodes
		}
		// Add or update the possibilities for this opcode number
		opcodeNumbers[s.opcode.number] = possibleOpcodes
	}

	solved := solve(opcodeNumbers)

	r := registers{}
	for _, uop := range program {
		op := Opcode{
			name: solved[uop.number],
			a:    uop.a,
			b:    uop.b,
			c:    uop.c,
		}
		op.process(&r)
	}
	fmt.Println("Part 2:", r[0])
}

// parseData reads the input text file and returns data structures
func parseData(file string) ([]sample, []uOpcode) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	parts := strings.Split(
		strings.TrimSuffix(string(data), "\n"), "\n\n",
	)

	part1 := parts[0 : len(parts)-2]
	part2 := parts[len(parts)-1]

	// Parse part 1 data
	samples := make([]sample, len(part1))
	for i, s := range part1 {
		lines := strings.Split(s, "\n")

		// Before
		b := strings.Split(
			strings.TrimPrefix(
				strings.TrimSuffix(
					lines[0],
					"]",
				),
				"Before: [",
			),
			", ",
		)
		br := registers{}
		for j, v := range b {
			if n, err := strconv.Atoi(v); err == nil {
				br[j] = n
			} else {
				log.Fatal(err)
			}
		}

		// Opcode
		op := stringToUOpcode(lines[1])

		// After
		a := strings.Split(
			strings.TrimPrefix(
				strings.TrimSuffix(
					lines[2],
					"]",
				),
				"After:  [",
			),
			", ",
		)
		ar := registers{}
		for j, v := range a {
			if n, err := strconv.Atoi(v); err == nil {
				ar[j] = n
			} else {
				log.Fatal(err)
			}
		}

		samples[i] = sample{
			before: br,
			after:  ar,
			opcode: op,
		}
	}

	// Parse part 2 data
	lines := strings.Split(part2, "\n")
	program := make([]uOpcode, len(lines))
	for i, line := range lines {
		op := stringToUOpcode(line)
		program[i] = op
	}

	return samples, program
}

func Run(inputFile string) {
	samples, program := parseData(inputFile)
	part1(samples, 3)
	part2(samples, program)
}
