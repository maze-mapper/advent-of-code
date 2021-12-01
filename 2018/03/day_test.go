// Advent of Code 2018 - Day 3 - Tests
package day3

import (
	"reflect"
	"testing"
)

var claims = []claim{
	{id: 1, x: 1, y: 3, width: 4, height: 4},
	{id: 2, x: 3, y: 1, width: 4, height: 4},
	{id: 3, x: 5, y: 5, width: 2, height: 2},
}

func TestParseData(t *testing.T) {
	data := []byte(`#1 @ 1,3: 4x4
#2 @ 3,1: 4x4
#3 @ 5,5: 2x2`)

	want := claims

	got := parseData(data)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v, want %v", got, want)
	}
}

func TestSolve(t *testing.T) {
	input := claims
	wantPart1 := 4
	wantPart2 := 3

	gotPart1, gotPart2 := solve(input)
	if gotPart1 != wantPart1 {
		t.Errorf("Got %d, want %d", gotPart1, wantPart1)
	}
	if gotPart2 != wantPart2 {
		t.Errorf("Got %d, want %d", gotPart2, wantPart2)
	}
}
