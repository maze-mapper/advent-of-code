package day5

import (
	"testing"
)

var input = []byte(`47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`)

func TestPart1(t *testing.T) {
	want := 143
	rules, pageNumbers, err := parseData(input)
	if err != nil {
		t.Fatal(err)
	}
	graph := buildDependencyGraph(rules)
	got := part1(graph, pageNumbers)
	if got != want {
		t.Errorf("part1(%v, %v) = %d, want %d", graph, pageNumbers, got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 123
	rules, pageNumbers, err := parseData(input)
	if err != nil {
		t.Fatal(err)
	}
	graph := buildDependencyGraph(rules)
	got := part2(graph, pageNumbers)
	if got != want {
		t.Errorf("part2(%v, %v) = %d, want %d", graph, pageNumbers, got, want)
	}
}
