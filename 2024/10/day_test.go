package day10

import (
	"testing"
)

var input = []byte(`89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`)

func TestPart1(t *testing.T) {
	want := 36
	region, err := parseData(input)
	if err != nil {
		t.Fatal(err)
	}
	got := part1(region)
	if got != want {
		t.Errorf("part1(%s) = %d, want %d", input, got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 81
	region, err := parseData(input)
	if err != nil {
		t.Fatal(err)
	}
	got := part2(region)
	if got != want {
		t.Errorf("part2(%s) = %d, want %d", input, got, want)
	}
}
