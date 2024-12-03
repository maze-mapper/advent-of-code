// Advent of Code 2024.
package aoc2024

import (
	"log"

	day1 "github.com/maze-mapper/advent-of-code/2024/01"
	day2 "github.com/maze-mapper/advent-of-code/2024/02"
	day3 "github.com/maze-mapper/advent-of-code/2024/03"
)

func Run(day, inputFile string) {
	f := func(s string) {}
	switch day {
	case "1":
		f = day1.Run
	case "2":
		f = day2.Run
	case "3":
		f = day3.Run
	default:
		log.Fatal(day, " is not a valid day")
	}
	f(inputFile)
}
