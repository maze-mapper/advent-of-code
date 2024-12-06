package day6

import (
	"testing"
)

var input = []byte(`
....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`)

func TestPart1(t *testing.T) {
	want := 41
	grid, start := parseData(input)
	got := part1(grid, start)
	if got != want {
		t.Errorf("part1(%s) = %d, want %d", input, got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 6
	grid, start := parseData(input)
	got := part2(grid, start)
	if got != want {
		t.Errorf("part2(%s) = %d, want %d", input, got, want)
	}
}
