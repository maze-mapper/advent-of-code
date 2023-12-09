package day9

import (
	"testing"
)

var input = [][]int{
	{0, 3, 6, 9, 12, 15},
	{1, 3, 6, 10, 15, 21},
	{10, 13, 16, 21, 30, 45},
}

func TestPart1(t *testing.T) {
	want := 114
	got := part1(input)
	if got != want {
		t.Errorf("part1(%#v) = %d, want %d", input, got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 2
	got := part2(input)
	if got != want {
		t.Errorf("part2(%#v) = %d, want %d", input, got, want)
	}
}
