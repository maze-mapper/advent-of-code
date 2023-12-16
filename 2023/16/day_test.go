package day16

import (
	"strings"
	"testing"
)

var input = [][]string{
	strings.Split(".|...\\....", ""),
	strings.Split("|.-.\\.....", ""),
	strings.Split(".....|-...", ""),
	strings.Split("........|.", ""),
	strings.Split("..........", ""),
	strings.Split(".........\\", ""),
	strings.Split("..../.\\\\..", ""),
	strings.Split(".-.-/..|..", ""),
	strings.Split(".|....-|.\\", ""),
	strings.Split("..//.|....", ""),
}

func TestPart1(t *testing.T) {
	want := 46
	got := part1(input)
	if got != want {
		t.Errorf("part1(%#v) = %d, want %d", input, got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 51
	got := part2(input)
	if got != want {
		t.Errorf("part2(%#v) = %d, want %d", input, got, want)
	}
}
