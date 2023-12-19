// Advent of Code 2023 - Day 19.
package day19

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type machinePart struct {
	x, m, a, s int
}

type machinePartRange struct {
	xMin, mMin, aMin, sMin int
	xMax, mMax, aMax, sMax int
}

type comparator int

const (
	always comparator = iota
	lessThan
	greaterThan
)

type workflowRule struct {
	category    string
	comaparison comparator
	value       int
	destination string
}

func (wfr workflowRule) apply(part machinePart) (string, bool) {
	passed := false
	switch wfr.comaparison {
	case always:
		passed = true
	case lessThan:
		switch wfr.category {
		case "x":
			passed = part.x < wfr.value
		case "m":
			passed = part.m < wfr.value
		case "a":
			passed = part.a < wfr.value
		case "s":
			passed = part.s < wfr.value
		}
	case greaterThan:
		switch wfr.category {
		case "x":
			passed = part.x > wfr.value
		case "m":
			passed = part.m > wfr.value
		case "a":
			passed = part.a > wfr.value
		case "s":
			passed = part.s > wfr.value
		}
	}
	if passed {
		return wfr.destination, true
	}
	return "", false
}

func parseData(data []byte) (map[string][]workflowRule, []machinePart) {
	sections := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n\n")

	workflowLines := strings.Split(sections[0], "\n")
	workflows := map[string][]workflowRule{}
	for _, line := range workflowLines {
		s := strings.TrimSuffix(line, "}")
		parts := strings.Split(s, "{")
		var rules []workflowRule
		for _, ss := range strings.Split(parts[1], ",") {
			ruleParts := strings.Split(ss, ":")
			if len(ruleParts) == 1 {
				rules = append(rules, workflowRule{destination: ruleParts[0]})
			} else {
				wfr := workflowRule{
					destination: ruleParts[1],
				}
				var comparisonParts []string
				if strings.Contains(ruleParts[0], "<") {
					wfr.comaparison = lessThan
					comparisonParts = strings.Split(ruleParts[0], "<")
				} else if strings.Contains(ruleParts[0], ">") {
					wfr.comaparison = greaterThan
					comparisonParts = strings.Split(ruleParts[0], ">")
				} else {
					log.Fatal("Unknown comparator")
				}
				wfr.category = comparisonParts[0]
				n, err := strconv.Atoi(comparisonParts[1])
				if err != nil {
					log.Fatal(err)
				}
				wfr.value = n
				rules = append(rules, wfr)
			}
		}
		workflows[parts[0]] = rules
	}

	machinePartLines := strings.Split(sections[1], "\n")
	mParts := make([]machinePart, len(machinePartLines))
	for i, line := range machinePartLines {
		s := strings.ReplaceAll(strings.ReplaceAll(strings.Trim(line, "{}"), ",", " "), "=", " = ")
		mp := machinePart{}
		n, err := fmt.Sscanf(s, "x = %d m = %d a = %d s = %d", &mp.x, &mp.m, &mp.a, &mp.s)
		if err != nil {
			log.Fatal(err)
		}
		if n != 4 {
			log.Fatal("Not enough items parsed")
		}
		mParts[i] = mp
	}

	return workflows, mParts
}

func sortPart(workflows map[string][]workflowRule, part machinePart) bool {
	workflow, ok := workflows["in"]
	if !ok {
		log.Fatal("Workflow not found")
	}
	for {
		for _, rule := range workflow {
			if dest, passed := rule.apply(part); passed {
				if dest == "A" {
					return true
				}
				if dest == "R" {
					return false
				}
				workflow, ok = workflows[dest]
				if !ok {
					log.Fatal("Workflow not found")
				}
				break
			}
		}
	}
}

func part1(workflows map[string][]workflowRule, parts []machinePart) int {
	sum := 0
	for _, p := range parts {
		if sortPart(workflows, p) {
			sum += p.x + p.m + p.a + p.s
		}
	}
	return sum
}

func part2(workflows map[string][]workflowRule) int {
	r := machinePartRange{
		xMin: 1, mMin: 1, aMin: 1, sMin: 1,
		xMax: 4000, mMax: 4000, aMax: 4000, sMax: 4000,
	}
	var allRanges []machinePartRange
	recurse(workflows, "in", r, &allRanges)

	total := 0
	for _, rr := range allRanges {
		total += (1 + rr.xMax - rr.xMin) * (1 + rr.mMax - rr.mMin) * (1 + rr.aMax - rr.aMin) * (1 + rr.sMax - rr.sMin)
	}
	return total
}

func recurse(workflows map[string][]workflowRule, name string, r machinePartRange, allRanges *[]machinePartRange) {
	if name == "A" {
		*allRanges = append(*allRanges, r)
		return
	}
	if name == "R" {
		return
	}
	workflow, ok := workflows[name]
	if !ok {
		log.Fatal("Workflow not found")
	}
	for {
		for _, rule := range workflow {
			switch rule.comaparison {
			case always:
				recurse(workflows, rule.destination, r, allRanges)
				return
			case lessThan:
				switch rule.category {
				case "x":
					if r.xMin < rule.value {
						recurse(workflows, rule.destination, machinePartRange{
							xMin: r.xMin, mMin: r.mMin, aMin: r.aMin, sMin: r.sMin,
							xMax: min(r.xMax, rule.value-1), mMax: r.mMax, aMax: r.aMax, sMax: r.sMax,
						}, allRanges)
						r.xMin = rule.value
					}

				case "m":
					if r.mMin < rule.value {
						recurse(workflows, rule.destination, machinePartRange{
							xMin: r.xMin, mMin: r.mMin, aMin: r.aMin, sMin: r.sMin,
							xMax: r.xMax, mMax: min(r.mMax, rule.value-1), aMax: r.aMax, sMax: r.sMax,
						}, allRanges)
						r.mMin = rule.value
					}

				case "a":
					if r.aMin < rule.value {
						recurse(workflows, rule.destination, machinePartRange{
							xMin: r.xMin, mMin: r.mMin, aMin: r.aMin, sMin: r.sMin,
							xMax: r.xMax, mMax: r.mMax, aMax: min(r.aMax, rule.value-1), sMax: r.sMax,
						}, allRanges)
						r.aMin = rule.value
					}

				case "s":
					if r.sMin < rule.value {
						recurse(workflows, rule.destination, machinePartRange{
							xMin: r.xMin, mMin: r.mMin, aMin: r.aMin, sMin: r.sMin,
							xMax: r.xMax, mMax: r.mMax, aMax: r.aMax, sMax: min(r.sMax, rule.value-1),
						}, allRanges)
						r.sMin = rule.value
					}

				}
			case greaterThan:
				switch rule.category {
				case "x":
					if r.xMax > rule.value {
						recurse(workflows, rule.destination, machinePartRange{
							xMin: max(r.xMin, rule.value+1), mMin: r.mMin, aMin: r.aMin, sMin: r.sMin,
							xMax: r.xMax, mMax: r.mMax, aMax: r.aMax, sMax: r.sMax,
						}, allRanges)
						r.xMax = rule.value
					}

				case "m":
					if r.mMax > rule.value {
						recurse(workflows, rule.destination, machinePartRange{
							xMin: r.xMin, mMin: max(r.mMin, rule.value+1), aMin: r.aMin, sMin: r.sMin,
							xMax: r.xMax, mMax: r.mMax, aMax: r.aMax, sMax: r.sMax,
						}, allRanges)
						r.mMax = rule.value
					}

				case "a":
					if r.aMax > rule.value {
						recurse(workflows, rule.destination, machinePartRange{
							xMin: r.xMin, mMin: r.mMin, aMin: max(r.aMin, rule.value+1), sMin: r.sMin,
							xMax: r.xMax, mMax: r.mMax, aMax: r.aMax, sMax: r.sMax,
						}, allRanges)
						r.aMax = rule.value
					}

				case "s":
					if r.sMax > rule.value {
						recurse(workflows, rule.destination, machinePartRange{
							xMin: r.xMin, mMin: r.mMin, aMin: r.aMin, sMin: max(r.sMin, rule.value+1),
							xMax: r.xMax, mMax: r.mMax, aMax: r.aMax, sMax: r.sMax,
						}, allRanges)
						r.sMax = rule.value
					}

				}
			}
		}
	}
}

func Run(inputFile string) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	workflows, parts := parseData(data)

	p1 := part1(workflows, parts)
	fmt.Println("Part 1:", p1)

	p2 := part2(workflows)
	fmt.Println("Part 2:", p2)
}
