// Advent of Code 2022
package aoc2022

import (
	"log"

	"github.com/maze-mapper/advent-of-code/2022/01"
)

func Run(day, inputFile string) {
	f := func(s string) {}
	switch day {
	case "1":
		f = day1.Run
	default:
		log.Fatal(day, " is not a valid day")
	}
	f(inputFile)
}
