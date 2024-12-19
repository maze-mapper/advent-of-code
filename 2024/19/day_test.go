package day19

import (
	"testing"
)

var input = []byte(`r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb`)

func TestPart1(t *testing.T) {
	want := 6
	patterns, designs := parseData(input)
	got := part1(patterns, designs)
	if got != want {
		t.Errorf("part1(%s) = %d, want %d", input, got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 16
	patterns, designs := parseData(input)
	got := part2(patterns, designs)
	if got != want {
		t.Errorf("part2(%s) = %d, want %d", input, got, want)
	}
}
