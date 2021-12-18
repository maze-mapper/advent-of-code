// Advent of Code 2021 - Day 18
package day18

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// snailFishNumber is a node in a binary tree representing a snailfish number
type snailFishNumber struct {
	lhs, rhs, parent *snailFishNumber
	number           int
}

// Copy returns a deep copy of a snailFishNumber
func (s *snailFishNumber) Copy() *snailFishNumber {
	root := &snailFishNumber{
		number: s.number,
	}
	if s.lhs != nil {
		root.lhs = s.lhs.Copy()
		root.lhs.parent = root
	}
	if s.rhs != nil {
		root.rhs = s.rhs.Copy()
		root.rhs.parent = root
	}
	return root
}

// isLeaf returns true if a snailFishNumber has no children
func (s *snailFishNumber) isLeaf() bool {
	return s.lhs == nil && s.rhs == nil
}

// String returns the string representation of a snailFishNumber
func (s *snailFishNumber) String() string {
	var b strings.Builder
	if s.isLeaf() {
		n := strconv.Itoa(s.number)
		b.WriteString(n)
	} else {
		b.WriteString("[")
		b.WriteString(s.lhs.String())
		b.WriteString(",")
		b.WriteString(s.rhs.String())
		b.WriteString("]")
	}
	return b.String()
}

// Magnitude returns the magnitude of a snailFishNumber
func (s *snailFishNumber) Magnitude() int {
	m := 0
	if s.isLeaf() {
		m = s.number
	} else {
		m += 3 * s.lhs.Magnitude()
		m += 2 * s.rhs.Magnitude()
	}
	return m
}

// updatePreviousRegularNumber will add "val" to the previous regular number from the current node if it exists
func (s *snailFishNumber) updatePreviousRegularNumber(val int) {
	lastVisited := s
	p := s.parent
	updated := false
	for p != nil && !updated {
		if p.lhs != lastVisited {
			updated = p.lhs.traverseAndUpdate(val, false)
		}
		lastVisited = p
		p = p.parent
	}
}

// updateNextRegularNumber will add "val" to the next regular number from the current node if it exists
func (s *snailFishNumber) updateNextRegularNumber(val int) {
	lastVisited := s
	p := s.parent
	updated := false
	for p != nil && !updated {
		if p.rhs != lastVisited {
			updated = p.rhs.traverseAndUpdate(val, true)
		}
		lastVisited = p
		p = p.parent
	}
}

// traverseAndUpdate will fully explore a binary tree from the given node and add "val" to the first regular number it finds the return
// inOrder specifies if the traversal should be done left to right or right to left
func (s *snailFishNumber) traverseAndUpdate(val int, inOrder bool) bool {
	if s.isLeaf() {
		s.number += val
		return true
	}
	if s.parent == nil {
		return true
	}

	if inOrder {
		left := s.lhs.traverseAndUpdate(val, inOrder)
		if left {
			return true
		}
		right := s.rhs.traverseAndUpdate(val, inOrder)
		if right {
			return true
		}
	} else {
		right := s.rhs.traverseAndUpdate(val, inOrder)
		if right {
			return true
		}
		left := s.lhs.traverseAndUpdate(val, inOrder)
		if left {
			return true
		}
	}
	return false
}

// explode will perform an explode action on a snailFishNumber
func (s *snailFishNumber) explode(depth int) bool {
	if depth >= 4 && !s.isLeaf() {
		s.updatePreviousRegularNumber(s.lhs.number)
		s.updateNextRegularNumber(s.rhs.number)
		s.lhs = nil
		s.rhs = nil
		s.number = 0
		return true
	} else if !s.isLeaf() {
		left := s.lhs.explode(depth + 1)
		if left {
			return true
		}
		right := s.rhs.explode(depth + 1)
		if right {
			return true
		}
	}
	return false
}

// split will perform a split action on a snailFishNumber
func (s *snailFishNumber) split() bool {
	if s.isLeaf() && s.number >= 10 {
		half := s.number / 2
		rem := s.number % 2

		s.lhs = &snailFishNumber{
			parent: s,
			number: half,
		}
		s.rhs = &snailFishNumber{
			parent: s,
			number: half + rem,
		}
		s.number = 0

		return true
	} else if !s.isLeaf() {
		left := s.lhs.split()
		if left {
			return true
		}
		right := s.rhs.split()
		if right {
			return true
		}
	}
	return false
}

// reduce will reduce a snailFishNumber by performing explode and split actions
func (s *snailFishNumber) reduce() {
	hasExploded := true
	hasSplit := true
	for hasExploded || hasSplit {
		hasExploded = s.explode(0)
		if hasExploded {
			continue
		}
		hasSplit = s.split()
	}
}

// FromString converts a string in to a snailFishNumber
func FromString(s string) *snailFishNumber {
	n, _ := fromString(s, nil)
	return n
}

// fromString converts a string in to a snailFishNumber
func fromString(s string, parent *snailFishNumber) (*snailFishNumber, int) {
	root := &snailFishNumber{parent: parent}
	left := true
	// Ignore starting "["
	for i := 1; i < len(s); i++ {
		switch s[i] {
		case '[':
			n, chars := fromString(s[i:], root)
			i += chars
			if left {
				root.lhs = n
			} else {
				root.rhs = n
			}
		case ']':
			return root, i

		case ',':
			left = false

		// Any integer 0-9
		default:
			if num, err := strconv.Atoi(string(s[i])); err == nil {
				leaf := &snailFishNumber{
					parent: root,
					number: num,
				}
				if left {
					root.lhs = leaf
				} else {
					root.rhs = leaf
				}
			} else {
				log.Fatal(err)
			}
		}
	}
	log.Fatal("Failed to parse snailfish number ", s)
	return root, -1
}

// Add will add two snailFishNumbers together
func Add(a, b *snailFishNumber) *snailFishNumber {
	root := &snailFishNumber{
		lhs: a,
		rhs: b,
	}
	a.parent = root
	b.parent = root

	root.reduce()

	return root
}

// parseData returns the data as a slice of snailFishNumber
func parseData(data []byte) []*snailFishNumber {
	lines := strings.Split(
		strings.TrimSuffix(string(data), "\n"), "\n",
	)
	numbers := make([]*snailFishNumber, len(lines))
	for i, line := range lines {
		numbers[i] = FromString(line)
	}
	return numbers
}

func part1(numbers []*snailFishNumber) int {
	n := numbers[0]
	for _, nn := range numbers[1:] {
		n = Add(n, nn)
	}
	return n.Magnitude()
}

func part2(numbers1, numbers2 []*snailFishNumber) int {
	maxMagnitude := 0
	for i, n1 := range numbers1 {
		for j, n2 := range numbers2 {
			if i != j {
				n1Copy := n1.Copy()
				n2Copy := n2.Copy()
				sum := Add(n1Copy, n2Copy)
				m := sum.Magnitude()
				if m > maxMagnitude {
					maxMagnitude = m
				}
			}
		}
	}
	return maxMagnitude
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	numbers := parseData(data)

	p1 := part1(numbers)
	fmt.Println("Part 1:", p1)

	// Re-read numbers as the existing ones will have been modified
	numbers1 := parseData(data)
	numbers2 := parseData(data)
	p2 := part2(numbers1, numbers2)
	fmt.Println("Part 2:", p2)
}
