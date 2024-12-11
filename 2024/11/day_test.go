package day11

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := []int{125, 17}
	want := 55312
	got := part1(input)
	if got != want {
		t.Errorf("part1(%v) = %d, want %d", input, got, want)
	}
}
