package day8

import (
	"testing"
)

var input = []byte(`............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`)

func TestPart1(t *testing.T) {
	want := 14
	antennas, minC, maxC := parseData(input)
	got := part1(antennas, minC, maxC)
	if got != want {
		t.Errorf("part1(%s) = %d, want %d", input, got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 34
	antennas, minC, maxC := parseData(input)
	got := part2(antennas, minC, maxC)
	if got != want {
		t.Errorf("part2(%s) = %d, want %d", input, got, want)
	}
}
