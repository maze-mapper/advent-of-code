package day9

import (
	"testing"
)

var input = []byte("2333133121414131402")

func TestPart1(t *testing.T) {
	want := 1928
	arr, err := parseData(input)
	if err != nil {
		t.Fatal(err)
	}
	got := part1(arr)
	if got != want {
		t.Errorf("part1(%s) = %d, want %d", input, got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 2858
	arr, err := parseData(input)
	if err != nil {
		t.Fatal(err)
	}
	got := part2(arr)
	if got != want {
		t.Errorf("part2(%s) = %d, want %d", input, got, want)
	}
}
