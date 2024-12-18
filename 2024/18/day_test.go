package day18

import (
	"testing"
)

var input = []byte(`5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0`)

func TestPart1(t *testing.T) {
	want := 22
	fallingBytes, err := parseData(input)
	if err != nil {
		t.Fatal(err)
	}
	got := part1(fallingBytes, 6, 6, 12)
	if got != want {
		t.Errorf("part1(%s) = %d, want %d", input, got, want)
	}
}

func TestPart2(t *testing.T) {
	want := "6,1"
	fallingBytes, err := parseData(input)
	if err != nil {
		t.Fatal(err)
	}
	got := part2(fallingBytes, 6, 6, 12)
	if got != want {
		t.Errorf("part2(%s) = %s, want %s", input, got, want)
	}
}
