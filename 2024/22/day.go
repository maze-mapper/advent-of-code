// Advent of Code 2024 - Day 22.
package day22

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseData(data []byte) ([]int, error) {
	lines := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
	numbers := make([]int, len(lines))
	for i, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		numbers[i] = n
	}
	return numbers, nil
}

func part1(numbers []int) int {
	var sum int
	for _, n := range numbers {
		for i := 0; i < 2000; i++ {
			n = evolve(n)
		}
		sum += n
	}
	return sum
}

func part2(numbers []int) int {
	bananasBySequence := map[[4]int]int{}
	for _, n := range numbers {
		var changes []int
		var prices []int
		var lastPrice int
		for j := 0; j < 2000; j++ {
			n = evolve(n)
			price := n % 10
			if j > 0 {
				changes = append(changes, price-lastPrice)
				prices = append(prices, price)
			}
			lastPrice = price
		}

		sequences := map[[4]int]bool{}
		for j := 3; j < len(changes); j++ {
			seq := [4]int{
				changes[j-3],
				changes[j-2],
				changes[j-1],
				changes[j],
			}
			if _, ok := sequences[seq]; !ok {
				sequences[seq] = true
				bananasBySequence[seq] += prices[j]
			}
		}
	}

	var best int
	for _, n := range bananasBySequence {
		if n > best {
			best = n
		}
	}
	return best
}

func evolve(n int) int {
	n = prune(n ^ (n * 64))
	n = prune(n ^ (n / 32))
	n = prune(n ^ (n * 2048))
	return n
}

func prune(n int) int {
	return n % 16777216
}

func Run(inputFile string) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	numbers, err := parseData(data)
	if err != nil {
		log.Fatal(err)
	}

	p1 := part1(numbers)
	fmt.Println("Part 1:", p1)

	p2 := part2(numbers)
	fmt.Println("Part 2:", p2)
}
