package day22

import (
	"testing"
)

func TestPart1(t *testing.T) {
	want := 37327623
	input := []byte(`1
10
100
2024`)
	numbers, err := parseData(input)
	if err != nil {
		t.Fatal(err)
	}
	got := part1(numbers)
	if got != want {
		t.Errorf("part1(%s) = %d, want %d", input, got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 23
	input := []byte(`1
2
3
2024`)
	numbers, err := parseData(input)
	if err != nil {
		t.Fatal(err)
	}
	got := part2(numbers)
	if got != want {
		t.Errorf("part2(%s) = %d, want %d", input, got, want)
	}
}
