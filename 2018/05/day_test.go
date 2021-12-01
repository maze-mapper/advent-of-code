// Advent of Code 2018 - Day 5 - Tests
package day5

import (
	"testing"
)

func byteSliceEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestAbs(t *testing.T) {
	tests := map[string]struct {
		input, want int
	}{
		"postive":  {input: 42, want: 42},
		"negative": {input: -42, want: 42},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := abs(tc.input)
			if got != tc.want {
				t.Errorf("Got %d, want %d", got, tc.want)
			}
		})
	}
}

func TestCheckOffset(t *testing.T) {
	tests := map[string]struct {
		a, b byte
		want bool
	}{
		"Same type, opposite polarity":         {a: byte('a'), b: byte('A'), want: true},
		"Same type, opposite polarity reverse": {a: byte('A'), b: byte('a'), want: true},
		"Same type, same polarity":             {a: byte('a'), b: byte('a'), want: false},
		"Different type, same polarity":        {a: byte('a'), b: byte('b'), want: false},
		"Different type, different polarity":   {a: byte('a'), b: byte('B'), want: false},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := checkOffset(tc.a, tc.b)
			if got != tc.want {
				t.Errorf("Got %t, want %t", got, tc.want)
			}
		})
	}
}

func TestReduce(t *testing.T) {
	tests := []struct {
		input, want []byte
		ok          bool
	}{
		{input: []byte("aA"), want: []byte(""), ok: true},
		{input: []byte("abBA"), want: []byte("aA"), ok: true},
		{input: []byte("abAB"), want: []byte("abAB"), ok: false},
		{input: []byte("aabAAB"), want: []byte("aabAAB"), ok: false},
	}

	for _, tc := range tests {
		t.Run(string(tc.input), func(t *testing.T) {
			got, ok := reduce(tc.input)
			if !byteSliceEqual(got, tc.want) {
				t.Errorf("Got %s, want %s", got, tc.want)
			}
			if ok != tc.ok {
				t.Errorf("Got %t, want %t", ok, tc.ok)
			}
		})
	}
}

func TestPart1(t *testing.T) {
	input := []byte("dabAcCaCBAcCcaDA")
	want := 10

	got := part1(input)
	if got != want {
		t.Errorf("Got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := []byte("dabCBAcaDA")
	want := 4

	got := part2(input)
	if got != want {
		t.Errorf("Got %d, want %d", got, want)
	}
}
