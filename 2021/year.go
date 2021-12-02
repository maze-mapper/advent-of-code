// Advent of Code 2021
package aoc2021

import (
	"log"

	"adventofcode/2021/01"
	"adventofcode/2021/02"
)

func Run(day, inputFile string) {
	f := func(s string) {}
	switch day {
	case "1":
		f = day1.Run
	case "2":
		f = day2.Run
	default:
		log.Fatal(day, " is not a valid day")
	}
	f(inputFile)
}
