// Advent of Code 2018 - Day 1 - Tests
package day1

import (
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {
	var tests = []struct {
		changes []int
		want    int
	}{
		{[]int{1, -2, 3, 1}, 3},
		{[]int{1, 1, 1}, 3},
		{[]int{1, 1, -2}, 0},
		{[]int{-1, -2, -3}, -6},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%d", tt.changes, tt.want)
		t.Run(testname, func(t *testing.T) {
			got := part1(tt.changes)
			if got != tt.want {
				t.Errorf("Got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	var tests = []struct {
		changes []int
		want    int
	}{
		{[]int{1, -2, 3, 1}, 2},
		{[]int{1, -1}, 0},
		{[]int{3, 3, 4, -2, -4}, 10},
		{[]int{-6, 3, 8, 5, -6}, 5},
		{[]int{7, 7, -2, -7, -4}, 14},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%d", tt.changes, tt.want)
		t.Run(testname, func(t *testing.T) {
			got := part2(tt.changes)
			if got != tt.want {
				t.Errorf("Got %d, want %d", got, tt.want)
			}
		})
	}
}
