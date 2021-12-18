// Advent of Code 2021 - Day 18 - Tests
package day18

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSnailFishNumberFromString(t *testing.T) {
	tests := map[string]*snailFishNumber{}

	// Manually create snailfish numbers
	// [1,2]
	root := &snailFishNumber{}
	root.lhs = &snailFishNumber{parent: root, number: 1}
	root.rhs = &snailFishNumber{parent: root, number: 2}
	tests["[1,2]"] = root

	// [[1,2],3]
	root = &snailFishNumber{}
	sub := &snailFishNumber{parent: root}
	sub.lhs = &snailFishNumber{parent: sub, number: 1}
	sub.rhs = &snailFishNumber{parent: sub, number: 2}
	root.lhs = sub
	root.rhs = &snailFishNumber{parent: root, number: 3}
	tests["[[1,2],3]"] = root

	// [9,[8,7]]
	root = &snailFishNumber{}
	sub = &snailFishNumber{parent: root}
	sub.lhs = &snailFishNumber{parent: sub, number: 8}
	sub.rhs = &snailFishNumber{parent: sub, number: 7}
	root.rhs = sub
	root.lhs = &snailFishNumber{parent: root, number: 9}
	tests["[9,[8,7]]"] = root

	for input, want := range tests {
		t.Run(input, func(t *testing.T) {
			got := FromString(input)
			// Check data structure
			if !reflect.DeepEqual(got, want) {
				t.Errorf("Got %v, want %v", got, want)
			}
			// Check string representation
			if input != got.String() {
				t.Errorf("Got %v, input was %s", got, input)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		a, b, want string
	}{
		{a: "[1,2]", b: "[[3,4],5]", want: "[[1,2],[[3,4],5]]"},
		{a: "[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]", b: "[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]", want: "[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]"},
		{a: "[[[[7,7],[7,7]],[[8,7],[8,7]]],[[[7,0],[7,7]],9]]", b: "[[[[4,2],2],6],[8,7]]", want: "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]"},
	}
	for _, tc := range tests {
		testname := fmt.Sprintf("%s+%s=%s", tc.a, tc.b, tc.want)
		t.Run(testname, func(t *testing.T) {
			got := Add(
				FromString(tc.a),
				FromString(tc.b),
			)
			gotStr := got.String()
			if gotStr != tc.want {
				t.Errorf("Got %s, want %s", gotStr, tc.want)
			}
		})
	}
}

func TestMagnitude(t *testing.T) {
	tests := map[string]int{
		"[9,1]":                             29,
		"[[9,1],[1,9]]":                     129,
		"[[1,2],[[3,4],5]]":                 143,
		"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]": 1384,
		"[[[[1,1],[2,2]],[3,3]],[4,4]]":     445,
		"[[[[3,0],[5,3]],[4,4]],[5,5]]":     791,
		"[[[[5,0],[7,4]],[5,5]],[6,6]]":     1137,
		"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]": 3488,
	}
	for name, want := range tests {
		t.Run(name, func(t *testing.T) {
			number := FromString(name)
			got := number.Magnitude()
			if got != want {
				t.Errorf("Got %d, want %d", got, want)
			}
		})
	}
}

func TestSplit(t *testing.T) {
	tests := []struct {
		input  *snailFishNumber
		result string
		want   bool
	}{
		{input: &snailFishNumber{number: 9}, want: false, result: "9"},
		{input: &snailFishNumber{number: 10}, want: true, result: "[5,5]"},
		{input: &snailFishNumber{number: 11}, want: true, result: "[5,6]"},
	}
	for _, tc := range tests {
		t.Run(tc.input.String(), func(t *testing.T) {
			got := tc.input.split()
			if got != tc.want {
				t.Errorf("Got %t, want %t", got, tc.want)
			}
			s := tc.input.String()
			if s != tc.result {
				t.Errorf("Got %s, want %s", s, tc.result)
			}
		})
	}
}

func TestExplodeSplit(t *testing.T) {
	tests := []struct {
		input, result string
		want          bool
	}{
		{input: "[[[[[9,8],1],2],3],4]", want: true, result: "[[[[0,9],2],3],4]"},
		{input: "[7,[6,[5,[4,[3,2]]]]]", want: true, result: "[7,[6,[5,[7,0]]]]"},
		{input: "[[6,[5,[4,[3,2]]]],1]", want: true, result: "[[6,[5,[7,0]]],3]"},
		{input: "[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]", want: true, result: "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"},
		{input: "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]", want: true, result: "[[3,[2,[8,0]]],[9,[5,[7,0]]]]"},
	}
	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			number := FromString(tc.input)
			got := number.explode(0)
			if got != tc.want {
				t.Errorf("Got %t, want %t", got, tc.want)
			}
			s := number.String()
			if s != tc.result {
				t.Errorf("Got %s, want %s", s, tc.result)
			}
		})
	}
}
