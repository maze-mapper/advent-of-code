// Advent of Code 2015 - Day 133
package day13

import (
	"container/ring"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var lineRegExp = regexp.MustCompile(`([a-zA-Z]+) would (gain|lose) ([0-9]+) happiness units by sitting next to ([a-zA-Z]+).`)

type Relationships map[string]map[string]int

// parseInput converts input data in to a Relationships object
func parseInput(input []byte) Relationships {
	lines := strings.Split(
		strings.TrimSuffix(string(input), "\n"), "\n",
	)

	relationships := Relationships{}
	for _, line := range lines {
		matches := lineRegExp.FindStringSubmatch(line)
		personA, personB := matches[1], matches[4]

		// Initialise nested maps
		if _, ok := relationships[personA]; !ok {
			relationships[personA] = make(map[string]int)
		}

		happiness, err := strconv.Atoi(matches[3])
		if err != nil {
			log.Fatal(err)
		}

		switch matches[2] {
		case "gain":
			relationships[personA][personB] = happiness
		case "lose":
			relationships[personA][personB] = -happiness
		}

	}

	return relationships
}

// generatePermutations uses Heap's algorithm to generate all permutations of a slice
func generatePermutations(k int, s []string) [][]string {
	output := [][]string{}

	if k == 1 {
		// Append a copy of the slice to avoid altering this permutation
		sl := make([]string, len(s))
		copy(sl, s)
		output = append(output, sl)
	} else {
		// Generate permutations with the kth element and beyond unaltered
		output = append(output, generatePermutations(k-1, s)...)

		// Generate permutations for kth element swapped with each k - 1
		for i := 0; i < k-1; i++ {
			// Swap choice is dependent on parity of k (even or odd)
			if k%2 == 0 {
				tmp := s[i]
				s[i] = s[k-1]
				s[k-1] = tmp
			} else {
				tmp := s[0]
				s[0] = s[k-1]
				s[k-1] = tmp
			}
			output = append(output, generatePermutations(k-1, s)...)
		}
	}

	return output
}

// happiness calculates the total happiness metric for a particular seating arrangement
func happiness(r *ring.Ring, relationships Relationships) int {
	var total int
	for i := 0; i < r.Len(); i++ {
		p := r.Prev()
		n := r.Next()

		total += relationships[r.Value.(string)][p.Value.(string)]
		total += relationships[r.Value.(string)][n.Value.(string)]

		r = n
	}
	return total
}

// printOrder prints the string values in a Ring
func printOrder(r *ring.Ring) {
	fmt.Printf("Arrangement is:")
	r.Do(func(p interface{}) {
		fmt.Printf(" %s", p.(string))
	})
	fmt.Printf("\n")
}

// solve finds the seating arrangement with the maximum happiness metric
func solve(relationships Relationships) (int, *ring.Ring) {
	people := []string{}
	for person := range relationships {
		people = append(people, person)
	}

	// To avoid generating circular permutations (such as 1234 and 4123) we fix one person and generate the permutations of the others
	// We will still end up with arrangements that are equivalent if the direction is reversed
	permutations := generatePermutations(len(people)-1, people[1:])

	maxHappiness := -int(^uint(0)>>1) - 1 // initialise to min value
	var bestOrder *ring.Ring

	for _, perm := range permutations {

		// Create a Ring for each arrangement
		r := ring.New(len(people))
		r.Value = people[0]
		for _, p := range perm {
			r = r.Next()
			r.Value = p
		}

		// Determine happiness
		if h := happiness(r, relationships); h > maxHappiness {
			maxHappiness = h
			bestOrder = r
		}
	}

	return maxHappiness, bestOrder
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	relationships := parseInput(data)

	maxHappiness, bestOrder := solve(relationships)
	fmt.Println("Part 1: Max happiness is", maxHappiness)
	printOrder(bestOrder)

	// Update relationships to include you
	yourRelationships := map[string]int{}
	for person, rel := range relationships {
		rel["You"] = 0
		yourRelationships[person] = 0
	}
	relationships["You"] = yourRelationships

	maxHappiness, bestOrder = solve(relationships)
	fmt.Println("Part 2: Max happiness is", maxHappiness)
	printOrder(bestOrder)
}
