package day6

import (
	"testing"
)

func TestDay6(t *testing.T) {
	tests := []struct {
		input                string
		part1Want, part2Want int
	}{
		{input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb", part1Want: 7, part2Want: 19},
		{input: "bvwbjplbgvbhsrlpgdmjqwftvncz", part1Want: 5, part2Want: 23},
		{input: "nppdvjthqldpwncqszvftbrmjlhg", part1Want: 6, part2Want: 23},
		{input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", part1Want: 10, part2Want: 29},
		{input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", part1Want: 11, part2Want: 26},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			part1Got := part1(tc.input)
			if part1Got != tc.part1Want {
				t.Errorf("part1(%q) = %d, want %d.", tc.input, part1Got, tc.part1Want)
			}

			part2Got := part2(tc.input)
			if part2Got != tc.part2Want {
				t.Errorf("part2(%q) = %d, want %d.", tc.input, part2Got, tc.part2Want)
			}
		})
	}
}
