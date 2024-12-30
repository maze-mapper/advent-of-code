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
	day9 "github.com/maze-mapper/advent-of-code/2024/09"
	day10 "github.com/maze-mapper/advent-of-code/2024/10"
	day11 "github.com/maze-mapper/advent-of-code/2024/11"
	day12 "github.com/maze-mapper/advent-of-code/2024/12"
	day13 "github.com/maze-mapper/advent-of-code/2024/13"
	day14 "github.com/maze-mapper/advent-of-code/2024/14"
	day15 "github.com/maze-mapper/advent-of-code/2024/15"
	day16 "github.com/maze-mapper/advent-of-code/2024/16"
	day17 "github.com/maze-mapper/advent-of-code/2024/17"
	day18 "github.com/maze-mapper/advent-of-code/2024/18"
	day19 "github.com/maze-mapper/advent-of-code/2024/19"
	day20 "github.com/maze-mapper/advent-of-code/2024/20"
	day21 "github.com/maze-mapper/advent-of-code/2024/21"
	day22 "github.com/maze-mapper/advent-of-code/2024/22"
	day23 "github.com/maze-mapper/advent-of-code/2024/23"
	day24 "github.com/maze-mapper/advent-of-code/2024/24"
	day25 "github.com/maze-mapper/advent-of-code/2024/25"
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
