package day1

import (
	"testing"
)

var input = []byte(`3   4
4   3
2   5
1   3
3   9
3   3`)

func TestPart1(t *testing.T) {
	want := 11
	a, b, err := parseData(input)
	if err != nil {
		t.Fatal(err)
	}
	got := part1(a, b)
	if got != want {
		t.Errorf("part1(%v, %v) = %d, want %d", a, b, got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 31
	a, b, err := parseData(input)
	if err != nil {
		t.Fatal(err)
	}
	got := part2(a, b)
	if got != want {
		t.Errorf("part2(%v, %v) = %d, want %d", a, b, got, want)
	}
}
