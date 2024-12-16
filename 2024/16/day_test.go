package day16

import (
	"testing"
)

var input1 = []byte(`###############
#.......#....E#
#.#.###.#.###.#
#.....#.#...#.#
#.###.#####.#.#
#.#.#.......#.#
#.#.#####.###.#
#...........#.#
###.#.#####.#.#
#...#.....#.#.#
#.#.#.###.#.#.#
#.....#...#.#.#
#.###.#.#.#.#.#
#S..#.....#...#
###############`)

var input2 = []byte(`#################
#...#...#...#..E#
#.#.#.#.#.#.#.#.#
#.#.#.#...#...#.#
#.#.#.#.###.#.#.#
#...#.#.#.....#.#
#.#.#.#.#.#####.#
#.#...#.#.#.....#
#.#.#####.#.###.#
#.#.#.......#...#
#.#.###.#####.###
#.#.#...#.....#.#
#.#.#.#####.###.#
#.#.#.........#.#
#.#.#.#########.#
#S#.............#
#################`)

func TestPart1(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
		want  int
	}{
		{
			name:  "Example_1",
			input: input1,
			want:  7036,
		},
		{
			name:  "Example_2",
			input: input2,
			want:  11048,
		},
	}
	for _, tc := range tests {
		t.Run(string(tc.name), func(t *testing.T) {
			maze, start, end := parseData(tc.input)
			got := part1(maze, start, end)
			if got != tc.want {
				t.Errorf("part1(%s) = %d, want %d", tc.input, got, tc.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
		want  int
	}{
		{
			name:  "Example_1",
			input: input1,
			want:  45,
		},
		{
			name:  "Example_2",
			input: input2,
			want:  64,
		},
	}
	for _, tc := range tests {
		t.Run(string(tc.name), func(t *testing.T) {
			maze, start, end := parseData(tc.input)
			got := part2(maze, start, end)
			if got != tc.want {
				t.Errorf("part2(%s) = %d, want %d", tc.input, got, tc.want)
			}
		})
	}
}
