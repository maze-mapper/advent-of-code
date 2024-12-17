package day17

import (
	"testing"
)

var input1 = []byte(`Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`)

var input2 = []byte(`Register A: 2024
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0`)

func TestPart1(t *testing.T) {
	want := "4,6,3,5,6,3,5,2,1,0"
	registers, program, err := parseData(input1)
	if err != nil {
		t.Fatal(err)
	}
	got := part1(registers, program)
	if got != want {
		t.Errorf("part1(%s) = %s, want %s", input1, got, want)
	}
}

// func TestPart2(t *testing.T) {
// 	want := 117440
// 	registers, program, err := parseData(input2)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	got := part2(registers, program)
// 	if got != want {
// 		t.Errorf("part2(%s) = %d, want %d", input1, got, want)
// 	}
// }
