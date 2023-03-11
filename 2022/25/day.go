// Advent of Code 2022 - Day 25.
package day25

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strings"
)

func parseInput(data []byte) []string {
	return strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
}

func snafuToInt(s string) int {
	var n int
	for j, r := range s {
		i := len(s) - j - 1
		p := int(math.Pow(float64(5), float64(i)))
		switch r {
		case '2':
			n += 2 * p
		case '1':
			n += p
		case '0':
			//
		case '-':
			n += -1 * p
		case '=':
			n += -2 * p
		}
	}
	return n
}

func intToSnafu(n int) string {
	maxPower := 0
	for {
		p := int(math.Pow(float64(5), float64(maxPower)))
		if n/p == 0 {
			break
		}
		maxPower += 1
	}

	s := ""
	parts := make([]int, maxPower+1)

	for pow := maxPower; pow >= 0; pow-- {
		p := int(math.Pow(float64(5), float64(pow)))
		v := n / p
		parts[maxPower-pow] = v
		n = n % p
	}
	for j := len(parts) - 1; j >= 1; j-- {
		switch parts[j] {
		case 2, 1, 0:
		case 3:
			parts[j] = -2
			parts[j-1] += 1
		case 4:
			parts[j] = -1
			parts[j-1] += 1
		case 5:
			parts[j] = 0
			parts[j-1] += 1
		}
	}

	for _, p := range parts {
		switch p {
		case 2:
			s += "2"
		case 1:
			s += "1"
		case 0:
			s += "0"
		case -1:
			s += "-"
		case -2:
			s += "="
		default:
			log.Fatalf("invalid number %d", p)
		}
	}
	return s
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	snafuNumbers := parseInput(data)

	total := 0
	for _, n := range snafuNumbers {
		in := snafuToInt(n)
		total += in
	}
	fmt.Println(intToSnafu(total))
}
