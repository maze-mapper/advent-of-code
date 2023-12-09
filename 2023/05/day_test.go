package day5

import (
	"testing"
)

var input = almanac{
	seeds: []int{79, 14, 55, 13},
	mappings: [][]mappingFunc{
		{
			makeMapping(50, 98, 2),
			makeMapping(52, 50, 48),
		},
		{
			makeMapping(0, 15, 37),
			makeMapping(37, 52, 2),
			makeMapping(39, 0, 15),
		},
		{
			makeMapping(49, 53, 8),
			makeMapping(0, 11, 42),
			makeMapping(42, 0, 7),
			makeMapping(57, 7, 4),
		},
		{
			makeMapping(88, 18, 7),
			makeMapping(18, 25, 70),
		},
		{
			makeMapping(45, 77, 23),
			makeMapping(81, 45, 19),
			makeMapping(68, 64, 13),
		},
		{
			makeMapping(0, 69, 1),
			makeMapping(1, 0, 69),
		},
		{
			makeMapping(60, 56, 37),
			makeMapping(56, 93, 4),
		},
	},
}

func TestPart1(t *testing.T) {
	want := 35
	got := part1(input)
	if got != want {
		t.Errorf("part1(%#v) = %d, want %d", input, got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 46
	got := part2(input)
	if got != want {
		t.Errorf("part2(%#v) = %d, want %d", input, got, want)
	}
}
