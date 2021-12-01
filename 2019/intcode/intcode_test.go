package intcode

import (
	"strconv"
	"strings"
	"testing"
)

func intSliceEqual(a, b []int) bool {
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

func testName(s []int) string {
	a := make([]string, len(s))
	for i, elem := range s {
		a[i] = strconv.Itoa(elem)
	}
	return strings.Join(a, "_")
}

func TestIncodeComputer(t *testing.T) {
	basicTests := []struct {
		input, want []int
	}{
		{input: []int{1, 0, 0, 0, 99}, want: []int{2, 0, 0, 0, 99}},
		{input: []int{2, 3, 0, 3, 99}, want: []int{2, 3, 0, 6, 99}},
		{input: []int{2, 4, 4, 5, 99, 0}, want: []int{2, 4, 4, 5, 99, 9801}},
		{input: []int{1, 1, 1, 4, 99, 5, 6, 0, 99}, want: []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
		{input: []int{1002, 4, 3, 4, 33}, want: []int{1002, 4, 3, 4, 99}},
		{input: []int{1101, 100, -1, 4, 0}, want: []int{1101, 100, -1, 4, 99}},
	}

	for _, tc := range basicTests {
		t.Run(testName(tc.input), func(t *testing.T) {
			computer := New(tc.input)
			computer.Run()
			got := computer.Program()
			if !intSliceEqual(got, tc.want) {
				t.Errorf("Got %v, want %v", got, tc.want)
			}
		})
	}

	ioTests := []struct {
		input, want, inputVals, outputVals []int
	}{
		{input: []int{3, 0, 4, 0, 99}, want: []int{42, 0, 4, 0, 99}, inputVals: []int{42}, outputVals: []int{42}},
		{input: []int{104, 1125899906842624, 99}, want: []int{104, 1125899906842624, 99}, inputVals: []int{0}, outputVals: []int{1125899906842624}},
	}
	for _, tc := range ioTests {
		t.Run(testName(tc.input), func(t *testing.T) {
			chanIn := make(chan int, len(tc.inputVals))
			chanOut := make(chan int, len(tc.outputVals))
			computer := New(tc.input)
			computer.SetChanIn(chanIn)
			computer.SetChanOut(chanOut)
			go computer.Run()
			chanIn <- tc.inputVals[0]
			recv := <-chanOut
			if recv != tc.outputVals[0] {
				t.Errorf("Got %d, want %d", recv, tc.outputVals[0])
			}
			got := computer.Program()
			if !intSliceEqual(got, tc.want) {
				t.Errorf("Got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestReadParameterMode(t *testing.T) {
	tests := []struct {
		input int
		want  []int
		name  string
	}{
		{input: 123, want: []int{3, 2, 1}, name: "three digits"},
		{input: 42, want: []int{2, 4, 0}, name: "two digits"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := readParameterMode(tc.input)
			if !intSliceEqual(got, tc.want) {
				t.Errorf("Got %v, want %v", got, tc.want)
			}
		})
	}
}
