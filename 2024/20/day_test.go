package day20

import (
	"testing"
)

var input = []byte(`###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############`)

func TestPart1(t *testing.T) {
	want := 44
	region, start, end := parseData(input)
	got := part1(region, start, end, 1)
	if got != want {
		t.Errorf("part1(%s) = %d, want %d", input, got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 285
	region, start, end := parseData(input)
	got := part2(region, start, end, 50)
	if got != want {
		t.Errorf("part2(%s) = %d, want %d", input, got, want)
	}
}
