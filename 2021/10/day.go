// Advent of Code 2021 - Day 10
package day10

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

// parseData returns the data as a slice of strings
func parseData(data []byte) []string {
	lines := strings.Split(
		strings.TrimSuffix(string(data), "\n"), "\n",
	)
	return lines
}

// pairs maps closing characters to opening ones
var pairs = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
	'>': '<',
}

// illegalScores are the scores for each invalid character found
var illegalScores = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

// incompleteScores are the scores for the missing pair character
var incompleteScores = map[rune]int{
	'(': 1,
	'[': 2,
	'{': 3,
	'<': 4,
}

// findScores returns the corrupted and incomplete scores for a line
// At least one of the scores will be zero
func findScores(line string) (int, int) {
	l := list.New()
	for _, r := range line {
		switch r {

		// Add opening character to stack
		case '(', '[', '{', '<':
			l.PushFront(r)

		// Handle closing character
		case ')', ']', '}', '>':
			if l.Len() > 0 {
				lastElem := l.Front()
				if pairs[r] == lastElem.Value.(rune) {
					// Character is valid, continue to next one
					l.Remove(lastElem)
					continue
				}
			}
			// Line is corrupted - either no preceding character or preceding character does not match
			return illegalScores[r], 0

		}
	}

	// Line may be incomplete
	score := 0
	for l.Len() > 0 {
		lastElem := l.Front()
		score *= 5
		score += incompleteScores[lastElem.Value.(rune)]
		l.Remove(lastElem)
	}
	return 0, score
}

// solve returns the scores due to corrupted lines and incomplete lines
func solve(lines []string) (int, int) {
	illegalScore := 0
	incompleteScores := []int{}
	for _, line := range lines {
		s1, s2 := findScores(line)
		illegalScore += s1
		if s2 != 0 {
			incompleteScores = append(incompleteScores, s2)
		}
	}

	sort.Ints(incompleteScores)
	incompleteScore := 0
	if len(incompleteScores) > 0 {
		incompleteScore = incompleteScores[len(incompleteScores)/2]
	}

	return illegalScore, incompleteScore
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	lines := parseData(data)

	p1, p2 := solve(lines)
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}
