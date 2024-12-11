// Advent of Code 2024 - Day 11.
package day11

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseData(data []byte) ([]int, error) {
	line := strings.TrimSuffix(string(data), "\n")
	parts := strings.Split(line, " ")
	stones := make([]int, len(parts))
	for i, s := range parts {
		n, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		stones[i] = n
	}
	return stones, nil
}

func part1(stones []int) int {
	return lengthAfterBlinks(stones, 25)
}

func part2(stones []int) int {
	return lengthAfterBlinks(stones, 75)
}

func lengthAfterBlinks(stones []int, blinks int) int {
	freq := map[int]int{}
	for _, stone := range stones {
		freq[stone]++
	}

	for i := 0; i < blinks; i++ {
		newFreq := map[int]int{}
		for k, v := range freq {
			for _, newStone := range blink(k) {
				newFreq[newStone] += v
			}
		}
		freq = newFreq
	}

	var length int
	for _, v := range freq {
		length += v
	}
	return length
}

func blink(stone int) []int {
	if stone == 0 {
		return []int{1}
	}
	if newStones, ok := splitIfEvenNumberOfDigits(stone); ok {
		return newStones
	}
	return []int{stone * 2024}
}

func splitIfEvenNumberOfDigits(n int) ([]int, bool) {
	s := strconv.Itoa(n)
	if len(s)%2 != 0 {
		return nil, false
	}
	mid := len(s) / 2
	n1, _ := strconv.Atoi(s[:mid])
	n2, _ := strconv.Atoi(s[mid:])
	return []int{n1, n2}, true
}

func Run(inputFile string) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	stones, err := parseData(data)
	if err != nil {
		log.Fatal(err)
	}

	p1 := part1(stones)
	fmt.Println("Part 1:", p1)

	p2 := part2(stones)
	fmt.Println("Part 2:", p2)
}
