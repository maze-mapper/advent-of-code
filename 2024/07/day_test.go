package day7

import (
	"testing"
)

var input = []byte(`190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`)

func TestPart1(t *testing.T) {
	want := 3749
	equations, err := parseData(input)
	if err != nil {
		t.Fatal(err)
	}
	got := part1(equations)
	if got != want {
		t.Errorf("part1(%s) = %d, want %d", input, got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 11387
	equations, err := parseData(input)
	if err != nil {
		t.Fatal(err)
	}
	got := part2(equations)
	if got != want {
		t.Errorf("part2(%s) = %d, want %d", input, got, want)
	}
}
