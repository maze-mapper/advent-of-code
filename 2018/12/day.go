// Advent of Code 2018 - Day 12
package day12

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// Rule holds the information for whether or not a a plant will exist in the next generation
type Rule struct {
	input  [5]bool
	output bool
}

// thingToBool converts a string or rune to a bool. Return true if it is "#" else false.
func thingToBool(i interface{}) bool {
	switch v := i.(type) {
	case string:
		if v == "#" {
			return true
		}
	case rune:
		if v == '#' {
			return true
		}
	}
	return false
}

// stringToBoolSlice converts a string representing pots to a bool slice
func stringToBoolSlice(s string) []bool {
	out := make([]bool, len(s))
	for i, v := range s {
		out[i] = thingToBool(v)
	}
	return out
}

// makeRule returns a Rule object from an input string
func makeRule(s string) Rule {
	parts := strings.Split(s, " => ")

	sl := stringToBoolSlice(parts[0])
	input := [5]bool{}
	for i, v := range sl {
		input[i] = v
	}

	rule := Rule{
		input:  input,
		output: thingToBool(parts[1]),
	}

	return rule
}

// parseData reads the input text file and returns data structures
func parseData(file string) ([]bool, []Rule) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	parts := strings.Split(
		strings.TrimSuffix(string(data), "\n"), "\n\n",
	)

	initial := strings.TrimPrefix(parts[0], "initial state: ")
	start := stringToBoolSlice(initial)

	lines := strings.Split(parts[1], "\n")
	rules := make([]Rule, len(lines))
	for i, line := range lines {
		rules[i] = makeRule(line)
	}

	return start, rules
}

// checkRule will check a slice of pots against a rule to determine if there should be plant or not
func checkRule(rule *Rule, pots []bool) bool {
	match := true
	for i, v := range rule.input {
		if pots[i] != v {
			match = false
			break
		}
	}
	return match
}

// advanceGeneration moves the selection of pots through one generation of changes
func advanceGeneration(pots []bool, rules []Rule, zeroIndex *int) []bool {
	// We need to consider pots beyond the current selection
	// Rules consider five pots so go a further four pots each end
	// This moves our original zero indexed pot
	pots = append(
		make([]bool, 4),
		pots...,
	)
	pots = append(
		pots,
		make([]bool, 4)...,
	)
	*zeroIndex += 4

	newPots := make([]bool, len(pots))
	// Copy current plants in case there are any that are left untouched by the rules
	copy(newPots, pots)

	for i := 2; i < len(newPots)-2; i++ {
		// TODO: input is guaranteed to have all combinations so could only check rules that result in a plant
		for _, rule := range rules {
			if match := checkRule(&rule, pots[i-2:i+3]); match {
				newPots[i] = rule.output
				break
			}
		}
	}

	return newPots
}

// sumPots sums the indices of all pots with a plant, factoring in where the original zero index pot now is
func sumPots(pots []bool, zeroIndex int) int {
	sum := 0
	for i, v := range pots {
		if v {
			sum += i - zeroIndex
		}
	}
	return sum
}

// potsToString converts the pots to a string representation
func potsToString(pots []bool) string {
	s := ""
	for _, v := range pots {
		if v {
			s += "#"
		} else {
			s += "."
		}
	}
	s = strings.Trim(s, ".")
	return s
}

// patternInfo holds generation and sum information for when a pattern is encountered
type patternInfo struct {
	generation, sum int
}

// solve will advance from startGeneration to finalGeneration but will skip iteration if a state is encountered again
func solve(startGeneration, finalGeneration int, pots []bool, rules []Rule, zeroIndex *int, patterns map[string]patternInfo) (int, []bool) {
	sum := 0
	for i := startGeneration; i <= finalGeneration; i++ {
		pots = advanceGeneration(pots, rules, zeroIndex)
		pattern := potsToString(pots)
		sum = sumPots(pots, *zeroIndex)
		if v, ok := patterns[pattern]; ok {
			// We've had the same pattern before
			genGap := i - v.generation
			scoreGap := sum - v.sum

			// Check if we will reach the final generation on this pattern
			if (finalGeneration-i)%genGap == 0 {
				cycles := (finalGeneration - i) / genGap
				finalSum := sum + scoreGap*cycles
				return finalSum, pots
			} else {
				// Update the generation
				patterns[pattern] = patternInfo{i, sum}
			}
		} else {
			patterns[pattern] = patternInfo{i, sum}
		}

	}
	return sum, pots
}

func Run(inputFile string) {
	pots, rules := parseData(inputFile)
	zeroIndex := 0
	sum := sumPots(pots, zeroIndex)

	// Part 1
	patterns := map[string]patternInfo{}
	patterns[potsToString(pots)] = patternInfo{0, sum}
	generations := 20
	sum, pots = solve(1, generations, pots, rules, &zeroIndex, patterns)
	fmt.Println("Part 1:", sum)

	// Part 2
	sum, _ = solve(generations+1, 50000000000, pots, rules, &zeroIndex, patterns)
	fmt.Println("Part 2:", sum)
}
