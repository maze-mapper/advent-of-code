// Advent of Code 2019
package aoc2019

import (
	"log"

//	"adventofcode/2019/01"
	"adventofcode/2019/02"
//	"adventofcode/2019/03"
//	"adventofcode/2019/04"
	"adventofcode/2019/05"
//	"adventofcode/2019/06"
	"adventofcode/2019/07"
//	"adventofcode/2019/08"
	"adventofcode/2019/09"
//	"adventofcode/2019/10"
	"adventofcode/2019/11"
//	"adventofcode/2019/12"
	"adventofcode/2019/13"
//	"adventofcode/2019/14"
	"adventofcode/2019/15"
//	"adventofcode/2019/16"
	"adventofcode/2019/17"
//	"adventofcode/2019/18"
	"adventofcode/2019/19"
//	"adventofcode/2019/20"
	"adventofcode/2019/21"
//	"adventofcode/2019/22"
	"adventofcode/2019/23"
//	"adventofcode/2019/24"
//	"adventofcode/2019/25"
)

func Run(day, inputFile string) {
	f := func(s string){}
	switch day {
//	case "1":
//		f = day1.Run
	case "2":
		f = day2.Run
//	case "3":
//		f = day3.Run
//	case "4":
//		f = day4.Run
	case "5":
		f = day5.Run
//	case "6":
//		f = day6.Run
	case "7":
		f = day7.Run
//	case "8":
//		f = day8.Run
	case "9":
		f = day9.Run
//	case "10":
//		f = day10.Run
	case "11":
		f = day11.Run
//	case "12":
//		f = day12.Run
	case "13":
		f = day13.Run
//	case "14":
//		f = day14.Run
	case "15":
		f = day15.Run
//	case "16":
//		f = day16.Run
	case "17":
		f = day17.Run
//	case "18":
//		f = day18.Run
	case "19":
		f = day19.Run
//	case "20":
//		f = day20.Run
	case "21":
		f = day21.Run
//	case "22":
//		f = day22.Run
	case "23":
		f = day23.Run
//	case "24":
//		f = day24.Run
//	case "25":
//		f = day25.Run
	default:
		log.Fatal(day, " is not a valid day")
	}
	f(inputFile)
}
