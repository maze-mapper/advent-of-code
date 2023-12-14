// Advent of Code 2023.
package aoc2023

import (
	"log"

	day1 "github.com/maze-mapper/advent-of-code/2023/01"
	day2 "github.com/maze-mapper/advent-of-code/2023/02"
	day3 "github.com/maze-mapper/advent-of-code/2023/03"
	day4 "github.com/maze-mapper/advent-of-code/2023/04"
	day5 "github.com/maze-mapper/advent-of-code/2023/05"
	day6 "github.com/maze-mapper/advent-of-code/2023/06"
	day7 "github.com/maze-mapper/advent-of-code/2023/07"
	day8 "github.com/maze-mapper/advent-of-code/2023/08"
	day9 "github.com/maze-mapper/advent-of-code/2023/09"
	day10 "github.com/maze-mapper/advent-of-code/2023/10"
	day11 "github.com/maze-mapper/advent-of-code/2023/11"
	day12 "github.com/maze-mapper/advent-of-code/2023/12"
	day13 "github.com/maze-mapper/advent-of-code/2023/13"
	day14 "github.com/maze-mapper/advent-of-code/2023/14"
	// "github.com/maze-mapper/advent-of-code/2023/15"
	// "github.com/maze-mapper/advent-of-code/2023/16"
	// "github.com/maze-mapper/advent-of-code/2023/17"
	// "github.com/maze-mapper/advent-of-code/2023/18"
	// "github.com/maze-mapper/advent-of-code/2023/19"
	// "github.com/maze-mapper/advent-of-code/2023/20"
	// "github.com/maze-mapper/advent-of-code/2023/21"
	// "github.com/maze-mapper/advent-of-code/2023/22"
	// "github.com/maze-mapper/advent-of-code/2023/23"
	// "github.com/maze-mapper/advent-of-code/2023/24"
	// "github.com/maze-mapper/advent-of-code/2023/25"
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
	// case "15":
	// 	f = day15.Run
	// case "16":
	// 	f = day16.Run
	// case "17":
	// 	f = day17.Run
	// case "18":
	// 	f = day18.Run
	// case "19":
	// 	f = day19.Run
	// case "20":
	// 	f = day20.Run
	// case "21":
	// 	f = day21.Run
	// case "22":
	// 	f = day22.Run
	// case "23":
	// 	f = day23.Run
	// case "24":
	// 	f = day24.Run
	// case "25":
	// 	f = day25.Run
	default:
		log.Fatal(day, " is not a valid day")
	}
	f(inputFile)
}
