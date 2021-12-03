// Advent of Code 2021 - Day 3
package day3

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// parseInput converts the input binary format numbers in to a slice of strings
func parseInput(data []byte) []string {
	lines := strings.Split(
		strings.TrimSuffix(string(data), "\n"), "\n",
	)
	return lines
}

func part1(diagnostics []string) int {
	l := len(diagnostics[0])
	zeros := make([]int, l)
	ones := make([]int, l)

	for _, d := range diagnostics {
		for i, r := range d {
			switch r {
			case '0':
				zeros[i] += 1
			case '1':
				ones[i] += 1
			}
		}
	}

	var gamma, epsilon strings.Builder
	for i := 0; i < l; i++ {
		if zeros[i] > ones[i] {
			gamma.WriteString("0")
			epsilon.WriteString("1")
		} else {
			gamma.WriteString("1")
			epsilon.WriteString("0")
		}
	}
	g := gamma.String()
	e := epsilon.String()

	gg, err := strconv.ParseUint(g, 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	ee, err := strconv.ParseUint(e, 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	return int(gg) * int(ee)
}

// bitCounts returns the number of zeros and ones in a slice of strings
func bitCounts(diagnostics []string, pos int) (int, int) {
	var zeros, ones int
	for _, d := range diagnostics {
		switch d[pos] {
		case '0':
			zeros += 1
		case '1':
			ones += 1
		}
	}
	return zeros, ones
}

func reduce(diagnostics []string, pos int, val rune) []string {
	newDiagnostics := []string{}
	for _, d := range diagnostics {
		if rune(d[pos]) == val {
			newDiagnostics = append(newDiagnostics, d)
		}
	}
	return newDiagnostics
}

func part2(diagnostics []string) int {
	l := len(diagnostics[0])

	// Find oxygen generator rating
	oDiagnostics := make([]string, len(diagnostics))
	copy(oDiagnostics, diagnostics)
	for i := 0; i < l; i++ {
		zeros, ones := bitCounts(oDiagnostics, i)
		switch {
		case zeros > ones:
			oDiagnostics = reduce(oDiagnostics, i, '0')
		case zeros < ones:
			oDiagnostics = reduce(oDiagnostics, i, '1')
		default:
			oDiagnostics = reduce(oDiagnostics, i, '1')
		}
		if len(oDiagnostics) == 1 {
			break
		}
	}
	oxygenRating := oDiagnostics[0]

	// Find CO2 scrubber rating
	co2Diagnostics := make([]string, len(diagnostics))
	copy(co2Diagnostics, diagnostics)
	for i := 0; i < l; i++ {
		zeros, ones := bitCounts(co2Diagnostics, i)
		switch {
		case zeros > ones:
			co2Diagnostics = reduce(co2Diagnostics, i, '1')
		case zeros < ones:
			co2Diagnostics = reduce(co2Diagnostics, i, '0')
		default:
			co2Diagnostics = reduce(co2Diagnostics, i, '0')
		}
		if len(co2Diagnostics) == 1 {
			break
		}
	}
	co2Rating := co2Diagnostics[0]

	o, err := strconv.ParseUint(oxygenRating, 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	co2, err := strconv.ParseUint(co2Rating, 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	return int(o) * int(co2)
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	diagnostics := parseInput(data)

	p1 := part1(diagnostics)
	fmt.Println("Part 1:", p1)

	p2 := part2(diagnostics)
	fmt.Println("Part 2:", p2)
}
