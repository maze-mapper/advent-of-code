// Advent of Code 2015 - Day 3
package day3

import (
	"fmt"
	"io/ioutil"
	"log"
)

type position struct {
	x, y int
}

func (pos *position) move(dir rune) {
	switch dir {
	case '^':
		pos.y++
	case 'v':
		pos.y--
	case '>':
		pos.x++
	case '<':
		pos.x--
	}
}

// visit marks a postion as having been visited in a map
func visit(visited map[position]bool, pos position) {
	visited[pos] = true
}

func part1(text *string) {
	pos := position{0, 0}
	visited := map[position]bool{
		pos: true,
	}

	for _, c := range *text {
		pos.move(c)
		visit(visited, pos)
	}

	fmt.Println(len(visited))
}

func part2(text *string) {
	pos1 := position{0, 0}
	pos2 := position{0, 0}

	visited := map[position]bool{
		pos1: true,
	}

	for i, c := range *text {
		if i%2 == 0 {
			pos1.move(c)
			visit(visited, pos1)
		} else {
			pos2.move(c)
			visit(visited, pos2)
		}
	}

	fmt.Println(len(visited))
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	text := string(data)

	part1(&text)
	part2(&text)
}
