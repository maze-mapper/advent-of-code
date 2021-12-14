// Advent of Code 2021 - Day 14
package day14

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func parseData(data []byte) (string, map[string]string) {
	parts := strings.Split(
		strings.TrimSuffix(string(data), "\n"), "\n\n",
	)

	insertionRules := map[string]string{}
	for _, line := range strings.Split(parts[1], "\n") {
		var k, v string
		fmt.Sscanf(line, "%s -> %s", &k, &v)
		insertionRules[k] = v
	}

	return parts[0], insertionRules
}

func doInsertion(polymerTemplate string, insertionRules map[string]string) string {
	var b strings.Builder

	b.WriteString(string(polymerTemplate[0]))
	for i := 1; i < len(polymerTemplate); i++ {
		pair := polymerTemplate[i-1 : i+1]
		insert := insertionRules[pair]
		b.WriteString(insert)
		b.WriteString(string(polymerTemplate[i]))
	}

	return b.String()
}

func elementsByCount(polymerTemplate string) map[rune]int {
	counts := map[rune]int{}
	for _, r := range polymerTemplate {
		counts[r] += 1
	}
	return counts
}

func score(polymerTemplate string) int {
	counts := elementsByCount(polymerTemplate)
	min := int(^uint(0) >> 1)
	max := 0

	for _, v := range counts {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return max - min
}

func doNInsertions(polymerTemplate string, insertionRules map[string]string, n int) int {
	for i := 0; i < n; i++ {
		polymerTemplate = doInsertion(polymerTemplate, insertionRules)
	}

	return score(polymerTemplate)
}

func makePairMap(polymerTemplate string) map[string]int {
	pairs := map[string]int{}
	for i := 1; i < len(polymerTemplate); i++ {
		pair := polymerTemplate[i-1 : i+1]
		pairs[pair] += 1
	}
	return pairs
}

func superDoInsertion(pairs map[string]int, insertionRules map[string]string) map[string]int {
	newPairs := map[string]int{}
	for k, v := range pairs {
		insert := insertionRules[k]
		parts := strings.Split(k, "")
		newPairs[parts[0]+insert] += v
		newPairs[insert+parts[1]] += v
	}
	return newPairs
}

func superScore(pairs map[string]int, first, last string) int {
	totals := map[string]int{}
	totals[first] = 1
	totals[last] = 1
	for k, v := range pairs {
		parts := strings.Split(k, "")
		totals[parts[0]] += v
		totals[parts[1]] += v
	}

	for k, v := range totals {
		totals[k] = v / 2
	}

	min := int(^uint(0) >> 1)
	max := 0
	for _, v := range totals {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return max - min

}

func superDoNInsertions(polymerTemplate string, insertionRules map[string]string, n int) int {
	polymerPairs := makePairMap(polymerTemplate)
	for i := 0; i < n; i++ {
		polymerPairs = superDoInsertion(polymerPairs, insertionRules)
		fmt.Println(polymerPairs)
	}

	l := len(polymerTemplate)
	return superScore(polymerPairs, polymerTemplate[0:1], polymerTemplate[l-1:l])
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	polymerTemplate, insertionRules := parseData(data)

	p1 := doNInsertions(polymerTemplate, insertionRules, 10)
	fmt.Println("Part 1:", p1)

	p2 := superDoNInsertions(polymerTemplate, insertionRules, 40)
	fmt.Println("Part 2:", p2)

}
