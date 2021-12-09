// Advent of Code 2021
package aoc2021

import (
	"log"

	"adventofcode/2021/01"
	"adventofcode/2021/02"
	"adventofcode/2021/03"
	"adventofcode/2021/04"
	"adventofcode/2021/05"
	"adventofcode/2021/06"
	"adventofcode/2021/07"
	"adventofcode/2021/08"
	"adventofcode/2021/09"
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
	case "9":
		f = day9.Run
	default:
		log.Fatal(day, " is not a valid day")
	}
	f(inputFile)
}
