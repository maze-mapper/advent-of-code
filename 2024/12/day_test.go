package day12

import (
	"testing"
)

var (
	input1 = []byte(`AAAA
BBCD
BBCC
EEEC`)
	input2 = []byte(`OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`)
	input3 = []byte(`RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`)
	input4 = []byte(`EEEEE
EXXXX
EEEEE
EXXXX
EEEEE`)
	input5 = []byte(`AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA`)
)

func TestPart1(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
		want  int
	}{
		{
			name:  "Example 1",
			input: input1,
			want:  140,
		},
		{
			name:  "Example 2",
			input: input2,
			want:  772,
		},
		{
			name:  "Example 3",
			input: input3,
			want:  1930,
		},
	}
	for _, tc := range tests {
		t.Run(string(tc.name), func(t *testing.T) {
			region := parseData(tc.input)
			got := part1(region)
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
			name:  "Example 1",
			input: input1,
			want:  80,
		},
		{
			name:  "Example 2",
			input: input2,
			want:  436,
		},
		{
			name:  "Example 3",
			input: input3,
			want:  1206,
		},
		{
			name:  "Example 4",
			input: input4,
			want:  236,
		},
		{
			name:  "Example 5",
			input: input5,
			want:  368,
		},
	}
	for _, tc := range tests {
		t.Run(string(tc.name), func(t *testing.T) {
			region := parseData(tc.input)
			got := part2(region)
			if got != tc.want {
				t.Errorf("part2(%s) = %d, want %d", tc.input, got, tc.want)
			}
		})
	}
}
