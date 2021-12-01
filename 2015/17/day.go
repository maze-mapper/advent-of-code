// Advent of Code 2015 - Day 177
package day17

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

// recurse finds all combinations of containers that will exactly hold a volume
func recurse(remainingVolume int, remainingContainers, usedContainers []int, combinations *[][]int) {
	if remainingVolume < 0 {
		return
	} else if remainingVolume == 0 {
		newUsedContainers := make([]int, len(usedContainers))
		copy(newUsedContainers, usedContainers)
		*combinations = append(*combinations, newUsedContainers)
	} else {
		for i, c := range remainingContainers {
			// Only want smaller containers
			newRemainingContainers := make([]int, len(remainingContainers[i+1:]))
			copy(newRemainingContainers, remainingContainers[i+1:])

			newUsedContainers := append(usedContainers, c)
			recurse(remainingVolume-c, newRemainingContainers, newUsedContainers, combinations)
		}
	}
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(
		strings.TrimSuffix(string(data), "\n"), "\n",
	)

	containers := []int{}
	for _, line := range lines {
		if size, err := strconv.Atoi(line); err == nil {
			containers = append(containers, size)
		} else {
			log.Fatal(err)
		}
	}
	// Sort containers in descending order
	sort.Sort(sort.Reverse(sort.IntSlice(containers)))
	fmt.Println(containers)

	volume := 150
	combinations := [][]int{}
	recurse(volume, containers, []int{}, &combinations)
	fmt.Println("Part 1:", volume, "litres of eggnog can be fit in the containers", len(combinations), "different ways")

	minNumber := len(combinations)
	count := 0

	for _, c := range combinations {
		switch {
		case len(c) < minNumber:
			minNumber = len(c)
			count = 1
		case len(c) == minNumber:
			count += 1
		}
	}
	fmt.Println("Part 2: The minimum number of containers is", minNumber, "and there are", count, "ways to do this")
}
