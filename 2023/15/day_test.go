package day15

import (
	"testing"
)

var input = [][]byte{
	[]byte("rn=1"),
	[]byte("cm-"),
	[]byte("qp=3"),
	[]byte("cm=2"),
	[]byte("qp-"),
	[]byte("pc=4"),
	[]byte("ot=9"),
	[]byte("ab=5"),
	[]byte("pc-"),
	[]byte("pc=6"),
	[]byte("ot=7"),
}

func TestPart1(t *testing.T) {
	want := 1320
	got := part1(input)
	if got != want {
		t.Errorf("part1(%#v) = %d, want %d", input, got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 145
	got := part2(input)
	if got != want {
		t.Errorf("part2(%#v) = %d, want %d", input, got, want)
	}
}
