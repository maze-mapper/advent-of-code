// Advent of Code 2021 - Day 6
package day6

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// parseInput converts the input data in to a slice of ints
func parseInput(data []byte) []int {
	fishStr := strings.Split(
		strings.TrimSuffix(string(data), "\n"), ",",
	)

	fish := make([]int, len(fishStr))
	for i, f := range fishStr {
		if val, err := strconv.Atoi(f); err == nil {
			fish[i] = val
		} else {
			log.Fatal(err)
		}
	}
	return fish
}

// Max timer constants for fish spawning
var (
	fishMaxTimer    int = 6
	newFishMaxTimer int = 8
)

// countSpawnTimers returns a map of spawn timers to the number of fish with that timer value
func countSpawnTimers(fish []int) map[int]int {
	counts := map[int]int{}
	for _, f := range fish {
		counts[f] += 1
	}
	return counts
}

// countFish counts the total numer of fish in a map of spawn timer to number of fish
func countFish(fish map[int]int) int {
	count := 0
	for _, v := range fish {
		count += v
	}
	return count
}

// simulateSpawn spawns the fish for the given number of days and returns the total number of fish after that time
func simulateSpawn(fish map[int]int, days int) int {
	for day := 0; day < days; day++ {
		newFish := map[int]int{}
		// Decrease the timer for fish that are not spawning
		for i := 1; i <= newFishMaxTimer; i++ {
			newFish[i-1] = fish[i]
		}
		// Add newly spawned fish
		newFish[newFishMaxTimer] = fish[0]
		// Reset parent fish that have just spawned children
		newFish[fishMaxTimer] += fish[0]
		fish = newFish
	}
	return countFish(fish)
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	fish := parseInput(data)
	fishByTimer := countSpawnTimers(fish)

	p1 := simulateSpawn(fishByTimer, 80)
	fmt.Println("Part 1:", p1)

	p2 := simulateSpawn(fishByTimer, 256)
	fmt.Println("Part 2:", p2)
}
