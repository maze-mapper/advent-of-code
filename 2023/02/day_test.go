package day2

import (
	"testing"
)

var games [][]cubes = [][]cubes{
	{
		{r: 4, b: 3},
		{r: 1, g: 2, b: 6},
		{g: 2},
	},
	{
		{g: 2, b: 1},
		{r: 1, g: 3, b: 4},
		{g: 1, b: 1},
	},
	{
		{r: 20, g: 8, b: 6},
		{r: 4, b: 5},
		{g: 13},
		{r: 1, g: 5},
	},
	{
		{r: 3, g: 1, b: 6},
		{r: 6, g: 3},
		{r: 14, g: 3, b: 15},
	},
	{
		{r: 6, g: 3, b: 1},
		{r: 1, g: 2, b: 2},
	},
}

func TestPart1(t *testing.T) {
	want := 8
	got := part1(games)
	if got != want {
		t.Errorf("part1(%#v) = %d, want %d", games, got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 2286
	got := part2(games)
	if got != want {
		t.Errorf("part2(%#v) = %d, want %d", games, got, want)
	}
}
