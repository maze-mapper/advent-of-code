// Advent of Code 2018 - Day 2 - Tests
package day2

import (
	"fmt"
	"testing"
)

func TestContainsTwoLetters(t *testing.T) {
	var tests = []struct {
		id   string
		want bool
	}{
		{"abcdef", false},
		{"bababc", true},
		{"abbcde", true},
		{"abcccd", false},
		{"aabcdd", true},
		{"abcdee", true},
		{"ababab", false},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s,%t", tt.id, tt.want)
		t.Run(testname, func(t *testing.T) {
			counts := countChar(tt.id)
			got := containsExactlyNSameCharacters(counts, 2)
			if got != tt.want {
				t.Errorf("Got %t, want %t", got, tt.want)
			}
		})
	}
}

func TestContainsThreeLetters(t *testing.T) {
	var tests = []struct {
		id   string
		want bool
	}{
		{"abcdef", false},
		{"bababc", true},
		{"abbcde", false},
		{"abcccd", true},
		{"aabcdd", false},
		{"abcdee", false},
		{"ababab", true},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%s,%t", tt.id, tt.want)
		t.Run(testname, func(t *testing.T) {
			counts := countChar(tt.id)
			got := containsExactlyNSameCharacters(counts, 3)
			if got != tt.want {
				t.Errorf("Got %t, want %t", got, tt.want)
			}
		})
	}
}

func TestPart1(t *testing.T) {
	ids := []string{
		"abcdef",
		"bababc",
		"abbcde",
		"abcccd",
		"aabcdd",
		"abcdee",
		"ababab",
	}
	want := 12

	got := part1(ids)
	if got != want {
		t.Errorf("Got %d, want %d", got, want)
	}
}

func TestCompare(t *testing.T) {
	var tests = []struct {
		a, b  string
		want  bool
		index int
	}{
		{"abcde", "abcdef", false, -1},
		{"abcde", "axcye", false, 3},
		{"fghij", "fguij", true, 2},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%s,%s,%t", tt.a, tt.b, tt.want)
		t.Run(testname, func(t *testing.T) {
			got, idx := compare(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Got %t, want %t", got, tt.want)
			}
			if idx != tt.index {
				t.Errorf("Got %d, want %d", idx, tt.index)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	ids := []string{
		"abcde",
		"fghij",
		"klmno",
		"pqrst",
		"fguij",
		"axcye",
		"wvxyz",
	}
	want := "fgij"

	got := part2(ids)
	if got != want {
		t.Errorf("Got %s, want %s", got, want)
	}
}
