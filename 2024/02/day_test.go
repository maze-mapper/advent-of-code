package day2

import (
	"testing"
)

var input = []byte(`7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`)

func TestPart1(t *testing.T) {
	want := 2
	reports, err := parseData(input)
	if err != nil {
		t.Fatal(err)
	}
	got := part1(reports)
	if got != want {
		t.Errorf("part1(%v) = %d, want %d", reports, got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 4
	reports, err := parseData(input)
	if err != nil {
		t.Fatal(err)
	}
	got := part2(reports)
	if got != want {
		t.Errorf("part2(%v) = %d, want %d", reports, got, want)
	}
}
