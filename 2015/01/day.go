// Advent of Code 2015 - Day 1
package day1

import (
	"fmt"
	"io/ioutil"
	"log"
)

func part1(text string) {
	floor := 0
	for _, c := range text {
		switch c {
		case '(':
			floor++
		case ')':
			floor--
		}
	}
	fmt.Println(floor)
}

func part2(text string) {
	floor := 0
	for i, c := range text {
		switch c {
		case '(':
			floor++
		case ')':
			floor--
		}

		if floor == -1 {
			fmt.Println(i + 1)
			break
		}
	}
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	text := string(data)

	part1(text)
	part2(text)
}
