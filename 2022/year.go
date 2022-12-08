// Advent of Code 2022.
package aoc2022

import (
	"log"

	"github.com/maze-mapper/advent-of-code/2022/01"
	"github.com/maze-mapper/advent-of-code/2022/02"
	"github.com/maze-mapper/advent-of-code/2022/03"
	"github.com/maze-mapper/advent-of-code/2022/04"
	"github.com/maze-mapper/advent-of-code/2022/05"
	"github.com/maze-mapper/advent-of-code/2022/06"
	"github.com/maze-mapper/advent-of-code/2022/07"
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
	case "4":
		f = day4.Run
	case "5":
		f = day5.Run
	case "6":
		f = day6.Run
	case "7":
		f = day7.Run
	default:
		log.Fatal(day, " is not a valid day")
	}
	f(inputFile)
}