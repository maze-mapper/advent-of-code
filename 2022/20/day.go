// Advent of Code 2022 - Day 20.
package day20

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func parseInput(data []byte) []int {
	lines := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
	numbers := make([]int, len(lines))
	for i, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		numbers[i] = n
	}
	return numbers
}

// mix mixes the provided numbers with the number movement order specified by positions.
// The index of positions is the starting number and the value is its current position in the numbers slice.
func mix(numbers, positions []int) {
	l := len(numbers)
	for i := 0; i < l; i++ {
		currentIdx := positions[i]

		// Remove number.
		n := numbers[currentIdx]
		numbers = append(numbers[:currentIdx], numbers[currentIdx+1:]...)

		// Determine new index. List is now one element shorter.
		newIdx := currentIdx + n
		newIdx = newIdx % (l - 1)
		for newIdx < 0 {
			newIdx += l - 1
		}

		// Insert number at new index.
		numbers = append(numbers[:newIdx], append([]int{n}, numbers[newIdx:]...)...)

		// Update the current positions of the starting numbers.
		for j := 0; j < l; j++ {
			if positions[j] > currentIdx && positions[j] <= newIdx {
				positions[j] -= 1
			}
			if positions[j] < currentIdx && positions[j] >= newIdx {
				positions[j] += 1
			}
		}
		positions[i] = newIdx
	}
}

// initialPositions returns a slice where each element value is its index.
func initialPositions(n int) []int {
	positions := make([]int, n)
	for i := 0; i < n; i++ {
		positions[i] = i
	}
	return positions
}

// getIndexOfZero returns the index of the first zero in the slice of numbers.
func getIndexOfZero(numbers []int) (int, bool) {
	for i, n := range numbers {
		if n == 0 {
			return i, true
		}
	}
	return 0, false
}

// score returns the sum of the 1000th, 2000th and 3000th elements after zero in the slice of numbers.
func score(numbers []int) int {
	zeroIdx, ok := getIndexOfZero(numbers)
	if !ok {
		log.Fatal("did not find zero in numbers")
	}

	sum := 0
	l := len(numbers)
	for _, n := range []int{1000, 2000, 3000} {
		idx := (zeroIdx + n) % l
		sum += numbers[idx]
	}
	return sum
}

func part1(numbers []int) int {
	l := len(numbers)
	nums := make([]int, l)
	copy(nums, numbers)
	positions := initialPositions(l)
	mix(nums, positions)
	return score(nums)
}

func part2(numbers []int) int {
	l := len(numbers)
	nums := make([]int, l)
	copy(nums, numbers)
	for i := 0; i < l; i++ {
		nums[i] *= 811589153
	}
	positions := initialPositions(l)
	for i := 0; i < 10; i++ {
		mix(nums, positions)
	}
	return score(nums)
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	numbers := parseInput(data)

	p1 := part1(numbers)
	fmt.Println("Part 1:", p1)

	p2 := part2(numbers)
	fmt.Println("Part 2:", p2)
}
