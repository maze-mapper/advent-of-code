package day1

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}
	want := 142
	got := part1(input)
	if got != want {
		t.Errorf("part1(%q) = %d, want %d", input, got, want)
	}
}

func TestPart2(t *testing.T) {
	tests := map[string]struct {
		input []string
		want  int
	}{
		"example": {
			input: []string{
				"two1nine",
				"eightwothree",
				"abcone2threexyz",
				"xtwone3four",
				"4nineeightseven2",
				"zoneight234",
				"7pqrstsixteen",
			},
			want: 281,
		},
		"overlap_at_start": {
			input: []string{"eightwothree"},
			want:  83,
		},
		"overlap_at_end": {
			input: []string{"5twone"},
			want:  51,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := part2(tc.input)
			if got != tc.want {
				t.Errorf("part2(%q) = %d, want %d", tc.input, got, tc.want)
			}
		})
	}
}
