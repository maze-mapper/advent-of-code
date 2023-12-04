package day4

import (
	"testing"
)

var input = []scratchcard{
	{
		winningNumbers: map[int]bool{41: true, 48: true, 83: true, 86: true, 17: true},
		haveNumbers:    map[int]bool{83: true, 86: true, 6: true, 31: true, 17: true, 9: true, 48: true, 53: true},
	},
	{
		winningNumbers: map[int]bool{13: true, 32: true, 20: true, 16: true, 61: true},
		haveNumbers:    map[int]bool{61: true, 30: true, 68: true, 82: true, 17: true, 32: true, 24: true, 19: true},
	},
	{
		winningNumbers: map[int]bool{1: true, 21: true, 53: true, 59: true, 44: true},
		haveNumbers:    map[int]bool{69: true, 82: true, 63: true, 72: true, 16: true, 21: true, 14: true, 1: true},
	},
	{
		winningNumbers: map[int]bool{41: true, 92: true, 73: true, 84: true, 69: true},
		haveNumbers:    map[int]bool{59: true, 84: true, 76: true, 51: true, 58: true, 5: true, 54: true, 83: true},
	},
	{
		winningNumbers: map[int]bool{87: true, 83: true, 26: true, 28: true, 32: true},
		haveNumbers:    map[int]bool{88: true, 30: true, 70: true, 12: true, 93: true, 22: true, 82: true, 36: true},
	},
	{
		winningNumbers: map[int]bool{31: true, 18: true, 13: true, 56: true, 72: true},
		haveNumbers:    map[int]bool{74: true, 77: true, 10: true, 23: true, 35: true, 67: true, 36: true, 11: true},
	},
}

func TestPart1(t *testing.T) {
	want := 13
	got := part1(input)
	if got != want {
		t.Errorf("part1(%#v) = %d, want %d", input, got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 30
	got := part2(input)
	if got != want {
		t.Errorf("part2(%#v) = %d, want %d", input, got, want)
	}
}
