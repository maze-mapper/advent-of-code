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
	"github.com/maze-mapper/advent-of-code/2022/08"
	"github.com/maze-mapper/advent-of-code/2022/09"
	"github.com/maze-mapper/advent-of-code/2022/10"
	"github.com/maze-mapper/advent-of-code/2022/11"
	"github.com/maze-mapper/advent-of-code/2022/12"
	"github.com/maze-mapper/advent-of-code/2022/13"
	"github.com/maze-mapper/advent-of-code/2022/14"
	"github.com/maze-mapper/advent-of-code/2022/15"
	"github.com/maze-mapper/advent-of-code/2022/16"
	"github.com/maze-mapper/advent-of-code/2022/17"
	"github.com/maze-mapper/advent-of-code/2022/18"
	"github.com/maze-mapper/advent-of-code/2022/19"
	"github.com/maze-mapper/advent-of-code/2022/20"
	"github.com/maze-mapper/advent-of-code/2022/21"
	"github.com/maze-mapper/advent-of-code/2022/22"
	"github.com/maze-mapper/advent-of-code/2022/23"
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
	case "10":
		f = day10.Run
	case "11":
		f = day11.Run
	case "12":
		f = day12.Run
	case "13":
		f = day13.Run
	case "14":
		f = day14.Run
	case "15":
		f = day15.Run
	case "16":
		f = day16.Run
	case "17":
		f = day17.Run
	case "18":
		f = day18.Run
	case "19":
		f = day19.Run
	case "20":
		f = day20.Run
	case "21":
		f = day21.Run
	case "22":
		f = day22.Run
	case "23":
		f = day23.Run
	default:
		log.Fatal(day, " is not a valid day")
	}
	f(inputFile)
}
