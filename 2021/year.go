// Advent of Code 2021
package aoc2021

import (
	"log"

	"github.com/maze-mapper/advent-of-code/2021/01"
	"github.com/maze-mapper/advent-of-code/2021/02"
	"github.com/maze-mapper/advent-of-code/2021/03"
	"github.com/maze-mapper/advent-of-code/2021/04"
	"github.com/maze-mapper/advent-of-code/2021/05"
	"github.com/maze-mapper/advent-of-code/2021/06"
	"github.com/maze-mapper/advent-of-code/2021/07"
	"github.com/maze-mapper/advent-of-code/2021/08"
	"github.com/maze-mapper/advent-of-code/2021/09"
	"github.com/maze-mapper/advent-of-code/2021/10"
	"github.com/maze-mapper/advent-of-code/2021/11"
	"github.com/maze-mapper/advent-of-code/2021/12"
	"github.com/maze-mapper/advent-of-code/2021/13"
	"github.com/maze-mapper/advent-of-code/2021/14"
	"github.com/maze-mapper/advent-of-code/2021/15"
	"github.com/maze-mapper/advent-of-code/2021/16"
	"github.com/maze-mapper/advent-of-code/2021/17"
	"github.com/maze-mapper/advent-of-code/2021/18"
	"github.com/maze-mapper/advent-of-code/2021/19"
	"github.com/maze-mapper/advent-of-code/2021/20"
	"github.com/maze-mapper/advent-of-code/2021/21"
	"github.com/maze-mapper/advent-of-code/2021/22"
	"github.com/maze-mapper/advent-of-code/2021/23"
	"github.com/maze-mapper/advent-of-code/2021/24"
	"github.com/maze-mapper/advent-of-code/2021/25"
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
	case "24":
		f = day24.Run
	case "25":
		f = day25.Run
	default:
		log.Fatal(day, " is not a valid day")
	}
	f(inputFile)
}
