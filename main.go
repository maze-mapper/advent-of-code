package main

import (
	"flag"
	"log"

	"github.com/maze-mapper/advent-of-code/2015"
	"github.com/maze-mapper/advent-of-code/2018"
	"github.com/maze-mapper/advent-of-code/2019"
	"github.com/maze-mapper/advent-of-code/2021"
)

func main() {
	flag.Parse()
	if flag.NArg() != 3 {
		log.Fatal("Usage: <year> <day> <inputFile>")
	}
	year := flag.Arg(0)
	day := flag.Arg(1)
	inputFile := flag.Arg(2)

	f := func(s, ss string) {}
	switch year {
	case "2015":
		f = aoc2015.Run
	case "2018":
		f = aoc2018.Run
	case "2019":
		f = aoc2019.Run
	case "2021":
		f = aoc2021.Run
	default:
		log.Fatal(year, " is not a valid year")
	}
	f(day, inputFile)
}
