package day14

import (
	"testing"
)

var input = []byte(`p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3
`)

func TestPart1(t *testing.T) {
	want := 12
	robots, err := parseData(input)
	if err != nil {
		t.Fatal(err)
	}
	got := part1(robots, 11, 7)
	if got != want {
		t.Errorf("part1(%s) = %d, want %d", input, got, want)
	}
}
