package day3

import (
	"testing"
)

var input = []string{
	"467..114.",
	"...*.....",
	"..35..633",
	"......#..",
	"617*.....",
	".....+.58",
	"..592....",
	"......755",
	"...$.*...",
	".664.598.",
}

func TestPart1(t *testing.T) {
	want := 4361
	got := part1(input)
	if got != want {
		t.Errorf("part1(%#v) = %d, want %d", input, got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 467835
	got := part2(input)
	if got != want {
		t.Errorf("part2(%#v) = %d, want %d", input, got, want)
	}
}
