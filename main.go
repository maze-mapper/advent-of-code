package main

import (
	"flag"
	"log"

	aoc2015 "github.com/maze-mapper/advent-of-code/2015"
	aoc2018 "github.com/maze-mapper/advent-of-code/2018"
	aoc2019 "github.com/maze-mapper/advent-of-code/2019"
	aoc2021 "github.com/maze-mapper/advent-of-code/2021"
	aoc2022 "github.com/maze-mapper/advent-of-code/2022"
	aoc2023 "github.com/maze-mapper/advent-of-code/2023"
	aoc2024 "github.com/maze-mapper/advent-of-code/2024"
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
	case "2022":
		f = aoc2022.Run
	case "2023":
		f = aoc2023.Run
	case "2024":
		f = aoc2024.Run
	default:
		log.Fatal(year, " is not a valid year")
	}
	f(day, inputFile)
}
