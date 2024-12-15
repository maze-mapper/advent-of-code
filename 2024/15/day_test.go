package day15

import (
	"testing"
)

var input1 = []byte(`########
#..O.O.#
##@.O..#
#...O..#
#.#.O..#
#...O..#
#......#
########

<^^>>>vv<v>>v<<`)

var input2 = []byte(`##########
#..O..O.O#
#......O.#
#.OO..O.O#
#..O@..O.#
#O#..O...#
#O..O..O.#
#.OO.O.OO#
#....O...#
##########

<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^
vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v
><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<
<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^
^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><
^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^
>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^
<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>
^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>
v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^`)

var input3 = []byte(`#######
#...#.#
#.....#
#..OO@#
#..O..#
#.....#
#######

<vv<<^^<<^^`)

func TestPart1(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
		want  int
	}{
		{
			name:  "Small_Example",
			input: input1,
			want:  2028,
		},
		{
			name:  "Larger_Example",
			input: input2,
			want:  10092,
		},
	}
	for _, tc := range tests {
		t.Run(string(tc.name), func(t *testing.T) {
			positions, start, directions := parseData(tc.input)
			got := part1(positions, start, directions)
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
			name:  "Small_Example",
			input: input3,
			want:  618,
		},
		{
			name:  "Larger_Example",
			input: input2,
			want:  9021,
		},
	}
	for _, tc := range tests {
		t.Run(string(tc.name), func(t *testing.T) {
			positions, start, directions := parseData(tc.input)
			positions, start = scaleWidth(positions, start)
			got := part2(positions, start, directions)
			if got != tc.want {
				t.Errorf("part2(%s) = %d, want %d", tc.input, got, tc.want)
			}
		})
	}
}
