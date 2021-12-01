// Advent of Code 2015 - Day 244
package day24

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func parseData(data []byte) ([]int, int) {
	lines := strings.Split(
		strings.TrimSuffix(string(data), "\n"), "\n",
	)

	packages := make([]int, len(lines))
	var totalWeight int
	for i, line := range lines {
		if weight, err := strconv.Atoi(line); err == nil {
			packages[i] = weight
			totalWeight += weight
		}
	}

	return packages, totalWeight
}

type Distribution [][]int

func sum(sl []int) int {
	total := 0
	for _, val := range sl {
		total += val
	}
	return total
}

func recurse(packages []int, currentWeight, targetWeight int, currentPackages, otherPackages []int, currentResult Distribution, results *[]Distribution, minPackages, minQE *int, breakOnFound bool) bool {

	// Return if the current number of packages is greater than the minimum found so far
	if len(currentPackages) > *minPackages {
		return false
	}

	switch {

	case currentWeight > targetWeight:
		return false

	case currentWeight == targetWeight:
		qe := product(currentPackages)
		// Return if the quantum entanglement is greater than the minimum found so far
		if qe >= *minQE {
			return false
		}

		// Add current selection of packages to distribution
		newCurrentPackages := make([]int, len(currentPackages))
		copy(newCurrentPackages, currentPackages)
		currentResult = append(currentResult, newCurrentPackages)

		// Recombine to get slice of remaining packages
		remainingPackages := make([]int, len(packages))
		copy(remainingPackages, packages)

		newOtherPackages := make([]int, len(otherPackages))
		copy(newOtherPackages, otherPackages)

		remainingPackages = append(remainingPackages, newOtherPackages...)
		sort.Sort(sort.IntSlice(remainingPackages))

		// Check if all remaining packages will make up one group
		if sum(remainingPackages) == targetWeight {
			currentResult = append(currentResult, remainingPackages)
			*results = append(*results, currentResult)
			*minPackages = len(currentPackages)
			*minQE = qe
			return true
		} else {
			// Check if remaining packages can be split up in to groups of the correct weight
			newMinPackages := len(remainingPackages)
			newMinQE := int(^uint(0) >> 1)
			success := recurse(remainingPackages, 0, targetWeight, []int{}, []int{}, currentResult, results, &newMinPackages, &newMinQE, true)
			if success {
				*minPackages = len(currentPackages)
				*minQE = qe
			}
			return success
		}

	default:
		// packages are in ascending order
		for i := len(packages) - 1; i >= 0; i-- {
			// Only take smaller weights to avoid duplicates
			newPackages := make([]int, len(packages[:i]))
			copy(newPackages, packages[:i])
			// Store the larger unused packages for later use
			extraPackages := make([]int, len(packages[i+1:]))
			copy(extraPackages, packages[i+1:])

			success := recurse(newPackages, currentWeight+packages[i], targetWeight, append(currentPackages, packages[i]), append(otherPackages, extraPackages...), currentResult, results, minPackages, minQE, breakOnFound)
			if success && breakOnFound {
				return true
			}
		}
	}
	return false
}

func product(sl []int) int {
	prod := 1
	for _, v := range sl {
		prod *= v
	}
	return prod
}

func solve(packages []int, totalWeight, groups int) {
	if totalWeight%groups != 0 {
		log.Fatal("Total weight is not divisible by", groups)
	}
	targetWeight := totalWeight / groups

	fmt.Println(totalWeight, targetWeight)

	combinations := []Distribution{}
	minPackages := len(packages)
	minQuantumEntanglement := int(^uint(0) >> 1)

	recurse(packages, 0, targetWeight, []int{}, []int{}, Distribution{}, &combinations, &minPackages, &minQuantumEntanglement, false)
	fmt.Println(combinations[len(combinations)-1])
	fmt.Println("QE:", minQuantumEntanglement)
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	packages, totalWeight := parseData(data)

	fmt.Println(packages)
	solve(packages, totalWeight, 3)
	solve(packages, totalWeight, 4)
}
