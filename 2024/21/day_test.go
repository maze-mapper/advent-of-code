package day21

import (
	"testing"
)

func TestPart1(t *testing.T) {
	want := 126384
	input := []byte(`029A
980A
179A
456A
379A`)
	codes := parseData(input)
	got := part1(codes)
	if got != want {
		t.Errorf("part1(%s) = %d, want %d", input, got, want)
	}
}
