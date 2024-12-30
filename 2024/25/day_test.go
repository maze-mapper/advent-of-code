package day25

import (
	"testing"
)

var input = []byte(`#####
.####
.####
.####
.#.#.
.#...
.....

#####
##.##
.#.##
...##
...#.
...#.
.....

.....
#....
#....
#...#
#.#.#
#.###
#####

.....
.....
#.#..
###..
###.#
###.#
#####

.....
.....
.....
#....
#.#..
#.#.#
#####`)

func TestPart1(t *testing.T) {
	want := 3
	locks, keys := parseData(input)
	got := part1(locks, keys)
	if got != want {
		t.Errorf("part1(%s) = %d, want %d", input, got, want)
	}
}
