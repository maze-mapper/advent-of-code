// Advent of Code 2024.
package aoc2024

import (
	"log"

	day1 "github.com/maze-mapper/advent-of-code/2024/01"
	day2 "github.com/maze-mapper/advent-of-code/2024/02"
	day3 "github.com/maze-mapper/advent-of-code/2024/03"
	day4 "github.com/maze-mapper/advent-of-code/2024/04"
	day5 "github.com/maze-mapper/advent-of-code/2024/05"
	day6 "github.com/maze-mapper/advent-of-code/2024/06"
	day7 "github.com/maze-mapper/advent-of-code/2024/07"
	day8 "github.com/maze-mapper/advent-of-code/2024/08"
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
	case "8":
		f = day8.Run
	default:
		log.Fatal(day, " is not a valid day")
	}
	f(inputFile)
}
