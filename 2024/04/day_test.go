package day4

import (
	"testing"
)

var input = []byte(`MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`)

func TestPart1(t *testing.T) {
	want := 18
	ws := parseData(input)
	got := part1(ws)
	if got != want {
		t.Errorf("part1(%v) = %d, want %d", input, got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 9
	ws := parseData(input)
	got := part2(ws)
	if got != want {
		t.Errorf("part2(%v) = %d, want %d", input, got, want)
	}
}
